package sentry

import (
	"encoding/json"
	"io"
	"time"
	"unsafe"

	"github.com/buger/jsonparser"
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"

	"github.com/fluxninja/aperture/pkg/info"
	"github.com/fluxninja/aperture/pkg/log"
	"github.com/fluxninja/aperture/pkg/panichandler"
	"github.com/fluxninja/aperture/pkg/status"
)

var _ = io.WriteCloser(new(SentryWriter))

const (
	// SentryFlushWait is the duration to wait for the sentry client to flush its queue.
	SentryFlushWait = "5000ms"
)

var zerologToSentryLevel = map[zerolog.Level]sentry.Level{
	log.DebugLevel: sentry.LevelDebug,
	log.InfoLevel:  sentry.LevelInfo,
	log.WarnLevel:  sentry.LevelWarning,
	log.ErrorLevel: sentry.LevelError,
	log.FatalLevel: sentry.LevelFatal,
	log.PanicLevel: sentry.LevelFatal,
}

// SentryWriter is a writer that forwards the data to the sentry client and the CrashWriter.
type SentryWriter struct {
	client         *sentry.Client
	levels         map[zerolog.Level]struct{}
	crashWriter    *CrashWriter
	statusRegistry status.Registry
}

// Write implements io.Writer and forwards the data to CrashWriter buffer.
func (s *SentryWriter) Write(data []byte) (int, error) {
	level, ok := s.parseLogLevel(data)

	if ok {
		if level == sentry.LevelFatal {
			// Parse log event only when it is fatal level and the event needs to be captured.
			event, err := s.parseLogEvent(level, data)
			if err != nil {
				log.Error().Err(err).Msg("Failed to parse log event")
			} else {
				s.client.CaptureEvent(event, nil, nil)
			}
			_ = s.Close()
		} else {
			_, _ = s.crashWriter.Write(data)
		}
	}
	return len(data), nil
}

func (s *SentryWriter) parseLogLevel(data []byte) (sentry.Level, bool) {
	levelStr, err := jsonparser.GetUnsafeString(data, zerolog.LevelFieldName)
	if err != nil {
		return "", false
	}

	level, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		return "", false
	}

	_, enabled := s.levels[level]
	if !enabled {
		return "", false
	}

	sentryLevel, ok := zerologToSentryLevel[level]
	if !ok {
		return "", false
	}

	return sentryLevel, true
}

func (s *SentryWriter) parseLogEvent(level sentry.Level, data []byte) (*sentry.Event, error) {
	event := sentry.Event{
		Timestamp: time.Now(),
		Level:     level,
		Logger:    "zerolog",
	}

	err := jsonparser.ObjectEach(data, func(key, value []byte, _ jsonparser.ValueType, _ int) error {
		switch string(key) {
		case zerolog.MessageFieldName:
			event.Message = bytesToStrUnsafe(value)
		case zerolog.ErrorFieldName:
			event.Exception = append(event.Exception, sentry.Exception{
				Value: bytesToStrUnsafe(value),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// Close implements io.Closer and wait for the sentry client to flush its queue.
func (s *SentryWriter) Close() error {
	duration, _ := time.ParseDuration(SentryFlushWait)
	s.client.Flush(duration)
	return nil
}

func bytesToStrUnsafe(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

// sentryPanicHandler is a panic handler that sends the fatal level event to Sentry with diagnostic information.
func (s *SentryWriter) sentryPanicHandler(e interface{}, _ panichandler.Callstack) {
	duration, _ := time.ParseDuration(SentryFlushWait)

	// Crash Log
	crashLogs := s.crashWriter.GetCrashLogs()
	for _, crashLog := range crashLogs {
		levelStr, ok := crashLog["level"].(string)
		if !ok {
			levelStr = "info"
		}
		level, err := zerolog.ParseLevel(levelStr)
		if err != nil {
			level = zerolog.InfoLevel
		}
		sentryLevel := zerologToSentryLevel[level]
		delete(crashLog, "level")

		msg, ok := crashLog["message"].(string)
		if !ok {
			msg = ""
		}
		delete(crashLog, "message")

		timestamp, ok := crashLog["timestamp"].(time.Time)
		if !ok {
			timestamp = time.Now()
		}
		delete(crashLog, "timestamp")

		sentry.AddBreadcrumb(&sentry.Breadcrumb{
			Type:      "debug",
			Category:  "log",
			Level:     sentryLevel,
			Data:      crashLog,
			Message:   msg,
			Timestamp: timestamp,
		})
	}

	if s.statusRegistry != nil {
		// Dump Status Registry
		if status := s.statusRegistry.GetGroupStatus(); status != nil {
			groupStatus, err := json.Marshal(status)
			if err != nil {
				log.Error().Err(err).Msg("Failed to marshal status registry")
			}

			statusData := make(map[string]interface{})
			err = json.Unmarshal(groupStatus, &statusData)
			if err != nil {
				log.Error().Err(err).Msg("Failed to unmarshal status registry")
			}

			sentry.AddBreadcrumb(&sentry.Breadcrumb{
				Category: "Status Registry",
				Level:    sentry.LevelInfo,
				Data:     statusData,
			})
		} else {
			sentry.AddBreadcrumb(&sentry.Breadcrumb{
				Category: "Status Registry",
				Level:    sentry.LevelInfo,
				Message:  "No Status Registry found",
			})
		}
	}

	// Service Version Information
	versionInfo := info.GetVersionInfo()
	if versionInfo != nil {
		vInfo, err := json.Marshal(versionInfo)
		if err != nil {
			log.Error().Err(err).Msg("Failed to marshal version information")
		}

		versionData := make(map[string]interface{})
		err = json.Unmarshal(vInfo, &versionData)
		if err != nil {
			log.Error().Err(err).Msg("Failed to unmarshal version information")
		}

		sentry.AddBreadcrumb(&sentry.Breadcrumb{
			Category: "Version Information",
			Level:    sentry.LevelInfo,
			Data:     versionData,
		})
	} else {
		sentry.AddBreadcrumb(&sentry.Breadcrumb{
			Category: "Version Information",
			Level:    sentry.LevelInfo,
			Message:  "No Version Information found",
		})
	}

	sentry.CurrentHub().Recover(e)
	sentry.Flush(duration)
}
