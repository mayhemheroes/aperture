package components

import (
	"fmt"

	"go.uber.org/fx"

	policylangv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/policy/language/v1"
	"github.com/fluxninja/aperture/pkg/config"
	"github.com/fluxninja/aperture/pkg/notifiers"
	"github.com/fluxninja/aperture/pkg/policies/controlplane/constraints"
	"github.com/fluxninja/aperture/pkg/policies/controlplane/iface"
	"github.com/fluxninja/aperture/pkg/policies/controlplane/runtime"
)

// Integrator is a component that accumulates sum of signal every tick.
type Integrator struct {
	sum    float64
	minMax *constraints.MinMaxConstraints
}

// Name implements runtime.Component.
func (*Integrator) Name() string { return "Integrator" }

// Type implements runtime.Component.
func (*Integrator) Type() runtime.ComponentType { return runtime.ComponentTypeSignalProcessor }

// ShortDescription implements runtime.Component.
func (in *Integrator) ShortDescription() string {
	return fmt.Sprintf("min: %f, max: %f", in.minMax.Min, in.minMax.Max)
}

// NewIntegrator creates an integrator component.
func NewIntegrator() runtime.Component {
	integrator := &Integrator{
		sum:    0,
		minMax: constraints.NewMinMaxConstraints(),
	}
	return integrator
}

// NewIntegratorAndOptions creates an integrator component and its fx options.
func NewIntegratorAndOptions(_ *policylangv1.Integrator, _ string, _ iface.Policy) (runtime.Component, fx.Option, error) {
	return NewIntegrator(), fx.Options(), nil
}

// Execute implements runtime.Component.Execute.
func (in *Integrator) Execute(inPortReadings runtime.PortToReading, tickInfo runtime.TickInfo) (runtime.PortToReading, error) {
	inputVal := inPortReadings.ReadSingleReadingPort("input")
	resetVal := inPortReadings.ReadSingleReadingPort("reset")
	if resetVal.Valid() && resetVal.Value() != 0 {
		in.sum = 0
	} else if inputVal.Valid() {
		minVal := inPortReadings.ReadSingleReadingPort("min")
		maxVal := inPortReadings.ReadSingleReadingPort("max")

		if minVal.Valid() && maxVal.Valid() {
			in.minMax.Max = maxVal.Value()
			in.minMax.Min = minVal.Value()

			value, _ := in.minMax.Constrain(inputVal.Value())
			in.sum += value
		}
	}

	return runtime.PortToReading{
		"output": []runtime.Reading{runtime.NewReading(in.sum)},
	}, nil
}

// DynamicConfigUpdate is a no-op for Integrator.
func (in *Integrator) DynamicConfigUpdate(event notifiers.Event, unmarshaller config.Unmarshaller) {}
