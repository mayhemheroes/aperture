package preview

import (
	flowpreviewv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/flowcontrol/preview/v1"
	"github.com/fluxninja/aperture/pkg/agentinfo"
	"github.com/fluxninja/aperture/pkg/policies/flowcontrol/iface"
	classification "github.com/fluxninja/aperture/pkg/policies/flowcontrol/resources/classifier"
)

// Handler implements flowpreview.v1 service.
type Handler struct {
	flowpreviewv1.UnimplementedFlowPreviewServiceServer
	engine     iface.Engine
	classifier *classification.ClassificationEngine
	agentGroup string
}

// NewHandler returns a new Handler.
func NewHandler(engine iface.Engine,
	classifier *classification.ClassificationEngine,
	agentInfo *agentinfo.AgentInfo,
) *Handler {
	return &Handler{
		engine:     engine,
		classifier: classifier,
		agentGroup: agentInfo.GetAgentGroup(),
	}
}
