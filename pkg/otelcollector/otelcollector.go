package otelcollector

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/converter/expandconverter"
	"go.opentelemetry.io/collector/otelcol"
	logsv1 "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	metricsv1 "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	tracev1 "go.opentelemetry.io/proto/otlp/collector/trace/v1"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/fluxninja/aperture/pkg/config"
	"github.com/fluxninja/aperture/pkg/log"
	"github.com/fluxninja/aperture/pkg/net/grpcgateway"
	otelconfig "github.com/fluxninja/aperture/pkg/otelcollector/config"
	"github.com/fluxninja/aperture/pkg/panichandler"
)

// Module is a fx module that invokes OTEL Collector.
func Module() fx.Option {
	return fx.Options(
		grpcgateway.RegisterHandler{Handler: logsv1.RegisterLogsServiceHandlerFromEndpoint}.Annotate(),
		grpcgateway.RegisterHandler{Handler: tracev1.RegisterTraceServiceHandlerFromEndpoint}.Annotate(),
		grpcgateway.RegisterHandler{Handler: metricsv1.RegisterMetricsServiceHandlerFromEndpoint}.Annotate(),
		fx.Invoke(setup),
	)
}

// ConstructorIn describes parameters passed to create OTEL Collector, server providing the OpenTelemetry Collector service.
type ConstructorIn struct {
	fx.In
	Factories     otelcol.Factories
	Lifecycle     fx.Lifecycle
	Shutdowner    fx.Shutdowner
	Unmarshaller  config.Unmarshaller
	BaseConfig    *otelconfig.OTELConfig `name:"base"`
	Logger        *log.Logger
	PluginConfigs []*otelconfig.OTELConfig `group:"plugin-config"`
}

// setup creates and runs a new instance of OTEL Collector with the passed configuration.
func setup(in ConstructorIn) error {
	uris := []string{"file:main"}
	var otelService *otelcol.Collector
	in.Lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			providers := map[string]confmap.Provider{
				"file": otelconfig.NewOTELConfigUnmarshaler(in.BaseConfig.AsMap()),
			}
			for i, pluginConfig := range in.PluginConfigs {
				scheme := fmt.Sprintf("plugin-%v", i)
				uris = append(uris, fmt.Sprintf("%v:%v", scheme, scheme))
				providers[scheme] = otelconfig.NewOTELConfigUnmarshaler(pluginConfig.AsMap())
			}

			configProvider, err := otelcol.NewConfigProvider(otelcol.ConfigProviderSettings{
				ResolverSettings: confmap.ResolverSettings{
					URIs:      uris,
					Providers: providers,
					Converters: []confmap.Converter{
						expandconverter.New(),
					},
				},
			})
			if err != nil {
				return fmt.Errorf("creating OTEL config provider: %w", err)
			}
			otelService, err = otelcol.NewCollector(
				otelcol.CollectorSettings{
					BuildInfo:               component.NewDefaultBuildInfo(),
					Factories:               in.Factories,
					ConfigProvider:          configProvider,
					DisableGracefulShutdown: true,
					LoggingOptions: []zap.Option{zap.WrapCore(func(zapcore.Core) zapcore.Core {
						return log.NewZapAdapter(in.Logger, "otel-collector")
					})},
					// NOTE: do not remove this becauase it causes a data-race condition.
					SkipSettingGRPCLogger: true,
				},
			)
			if err != nil {
				return fmt.Errorf("constructing OTEL Service: %v", err)
			}

			log.Info().Msg("Starting OTEL Collector")
			panichandler.Go(func() {
				err := otelService.Run(context.Background())
				if err != nil {
					log.Error().Err(err).Msg("Failed to run OTEL Collector")
				}
				_ = in.Shutdowner.Shutdown()
			})
			return nil
		},
		OnStop: func(context.Context) error {
			log.Info().Msg("Stopping OTEL Collector")
			otelService.Shutdown()
			return nil
		},
	})

	return nil
}
