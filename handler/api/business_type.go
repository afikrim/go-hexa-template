package api

import (
	"context"

	engine_v1 "github.com/afikrim/go-hexa-template/handler/api/pb/engine/v1"
)

func (h *handler) GetBusinessTypes(ctx context.Context, req *engine_v1.GetBusinessTypesRequest) (*engine_v1.GetBusinessTypesResponse, error) {
	out, err := h.svc.GetBusinessTypeModule().GetBusinessTypes(ctx, mapPbGetBusinessTypesRequest(req))
	if err != nil {
		return nil, err
	}

	return mapGetBusinessTypesOut(out), nil
}
