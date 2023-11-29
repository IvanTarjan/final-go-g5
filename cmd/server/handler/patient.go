package handler

import (
	"strconv"

	"github.com/IvanTarjan/final-go-g5/internal/domain"
	"github.com/IvanTarjan/final-go-g5/internal/patient"
	"github.com/IvanTarjan/final-go-g5/pkg/web"
	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	service patient.ServicePatients
}

func NewPatientHandler(service *patient.ServicePatients) *PatientHandler {
	return &PatientHandler{service: *service}
}

func (h *PatientHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var patient domain.Patient
		if err := ctx.BindJSON(&patient); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		createdPatient, err := h.service.Create(ctx, patient)
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 201, createdPatient)
	}
}

func (h *PatientHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		patient, err := h.service.GetByID(ctx, int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, patient)
	}
}

func (h *PatientHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var patient domain.Patient
		if err := ctx.BindJSON(&patient); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		updatedPatient, err := h.service.Update(ctx, patient, int64(patient.Id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, updatedPatient)
	}
}

func (h *PatientHandler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		var attributes map[string]string
		queryParams := ctx.Request.URL.Query()
		for k, v := range queryParams {
			if len(v) > 0 {
				attributes[k] = v[0]
			}
		}
		patchedPatient, err := h.service.Patch(ctx, attributes, int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, patchedPatient)
	}
}

func (h *PatientHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		err = h.service.Delete(ctx, int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, "Deleted patient with id "+strconv.FormatInt(int64(id), 10))
	}
}
