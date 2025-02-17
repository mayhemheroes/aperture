package components

import (
	"fmt"
	"time"

	"go.uber.org/fx"

	policylangv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/policy/language/v1"
	"github.com/fluxninja/aperture/pkg/alerts"
	"github.com/fluxninja/aperture/pkg/config"
	"github.com/fluxninja/aperture/pkg/info"
	"github.com/fluxninja/aperture/pkg/notifiers"
	"github.com/fluxninja/aperture/pkg/policies/controlplane/iface"
	"github.com/fluxninja/aperture/pkg/policies/controlplane/runtime"
)

// Alerter is a component that monitors signal value and creates alert on true value.
type Alerter struct {
	alerterIface   alerts.Alerter
	policyReadAPI  iface.Policy
	name           string
	severity       string
	alertChannels  []string
	resolveTimeout time.Duration
}

// Name implements runtime.Component.
func (*Alerter) Name() string { return "Alerter" }

// Type implements runtime.Component.
func (*Alerter) Type() runtime.ComponentType { return runtime.ComponentTypeSink }

// ShortDescription implements runtime.Component.
func (a *Alerter) ShortDescription() string { return fmt.Sprintf("%s/%s", a.name, a.severity) }

// Make sure Alerter complies with Component interface.
var _ runtime.Component = (*Alerter)(nil)

// NewAlerterAndOptions creates alerter and its fx options.
func NewAlerterAndOptions(alerterProto *policylangv1.Alerter, _ string, policyReadAPI iface.Policy) (runtime.Component, fx.Option, error) {
	parameters := alerterProto.Parameters
	alerter := &Alerter{
		name:           parameters.AlertName,
		severity:       parameters.Severity,
		resolveTimeout: parameters.ResolveTimeout.AsDuration(),
		alertChannels:  make([]string, 0),
		policyReadAPI:  policyReadAPI,
	}
	alerter.alertChannels = append(alerter.alertChannels, parameters.AlertChannels...)

	return alerter, fx.Options(
		fx.Invoke(
			alerter.setup,
		)), nil
}

func (a *Alerter) setup(alerterIface *alerts.SimpleAlerter) {
	a.alerterIface = alerterIface
}

// Execute implements runtime.Component.Execute.
func (a *Alerter) Execute(inPortReadings runtime.PortToReading, tickInfo runtime.TickInfo) (runtime.PortToReading, error) {
	signalValue := inPortReadings.ReadSingleReadingPort("signal")
	if !signalValue.Valid() {
		return nil, nil
	}

	if signalValue.Value() > 0 {
		a.alerterIface.AddAlert(a.createAlert())
	}

	return nil, nil
}

// DynamicConfigUpdate is a no-op for Alerter.
func (a *Alerter) DynamicConfigUpdate(event notifiers.Event, unmarshaller config.Unmarshaller) {
}

func (a *Alerter) createAlert() *alerts.Alert {
	newAlert := alerts.NewAlert(
		alerts.WithName(a.name),
		alerts.WithSeverity(alerts.ParseSeverity(a.severity)),
		alerts.WithAlertChannels(a.alertChannels),
		alerts.WithLabel("policy_name", a.policyReadAPI.GetPolicyName()),
		alerts.WithLabel("type", "alerter"),
		alerts.WithResolveTimeout(a.resolveTimeout),
		alerts.WithGeneratorURL(
			fmt.Sprintf("http://%s/%s/%s", info.GetHostInfo().Hostname, a.policyReadAPI.GetPolicyName(), a.name),
		),
	)

	return newAlert
}
