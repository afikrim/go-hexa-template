package api

import (
	"github.com/afikrim/go-hexa-template/core/entity"
	engine_v1 "github.com/afikrim/go-hexa-template/handler/api/pb/engine/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapPbGetBusinessTypesRequest(req *engine_v1.GetBusinessTypesRequest) *entity.GetBusinessTypesIn {
	return &entity.GetBusinessTypesIn{}
}

func mapBusinessType(businessType *entity.BusinessType) *engine_v1.BusinessType {
	if businessType == nil {
		return nil
	}

	return &engine_v1.BusinessType{
		Serial:             businessType.Serial,
		Name:               businessType.Name,
		Description:        businessType.Description,
		BusinessTypeSerial: businessType.BusinessTypeSerial,
		CreatedAt:          timestamppb.New(businessType.CreatedAt),
		UpdatedAt:          timestamppb.New(businessType.UpdatedAt),
	}
}

func mapBusinessTypes(businessTypes entity.BusinessTypes) []*engine_v1.BusinessType {
	var result []*engine_v1.BusinessType
	for _, businessType := range businessTypes {
		if businessType == nil {
			continue
		}

		result = append(result, mapBusinessType(businessType))
	}

	return result
}

func mapGetBusinessTypesOut(out *entity.GetBusinessTypesOut) *engine_v1.GetBusinessTypesResponse {
	return &engine_v1.GetBusinessTypesResponse{
		BusinessTypes: mapBusinessTypes(out.BusinessTypes),
	}
}

func mapError(err error) *status.Status {
	return status.New(codes.Internal, err.Error())
}
