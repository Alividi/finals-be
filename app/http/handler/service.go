package handler

import (
	dtoService "finals-be/app/services/dto"
	serviceService "finals-be/app/services/service"
	"finals-be/internal/lib/auth"
	"finals-be/internal/lib/helper"
	"net/http"
)

type ServiceHandler struct {
	serviceService *serviceService.ServiceService
}

func NewServiceHandler(serviceService *serviceService.ServiceService) *ServiceHandler {
	return &ServiceHandler{
		serviceService: serviceService,
	}
}

func (s *ServiceHandler) GetServices(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request dtoService.ServicesRequest

	userCtx := auth.GetUserContext(ctx)
	request.UserId = userCtx.ID
	request.Active = helper.GetQueryInt64Pointer(r, "active")

	services, err := s.serviceService.GetServices(ctx, request)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}
	helper.WriteResponse(ctx, w, nil, services)
}

func (s *ServiceHandler) GetServiceById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := helper.GetURLParamInt64(r, "id")

	serviceDetail, err := s.serviceService.GetServiceDetail(ctx, id)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	helper.WriteResponse(ctx, w, nil, serviceDetail)
}

func (s *ServiceHandler) GetServiceStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	serviceId := helper.GetURLParamInt64(r, "id")

	interval := helper.GetQueryInt64Pointer(r, "interval")

	req := dtoService.ServiceTelemetryRequest{
		Interval: interval,
	}

	telemetry, err := s.serviceService.GetServiceTelemetry(ctx, serviceId, req)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	helper.WriteResponse(ctx, w, nil, telemetry)
}

func (s *ServiceHandler) ChangeServiceCoordinates(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := dtoService.ChangeCoordinateRequest{}

	err := helper.ReadRequest(r, &req)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	if req.Latitude == 0 || req.Longitude == 0 {
		helper.WriteResponse(ctx, w, helper.NewErrBadRequest("Latitude and Longitude are required"), nil)
		return
	}

	if req.ServiceId == 0 {
		helper.WriteResponse(ctx, w, helper.NewErrBadRequest("ServiceId is required"), nil)
		return
	}

	err = s.serviceService.ChangeServiceCoordinate(ctx, req)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	helper.WriteResponse(ctx, w, nil, nil)
}

func (s *ServiceHandler) GetServiceTroubleshoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	gangguanId := helper.GetURLParamInt64(r, "gangguanId")

	problem, err := s.serviceService.GetTroubleshootSteps(ctx, gangguanId)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	helper.WriteResponse(ctx, w, nil, problem)
}
