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

func (h *PatientHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var patientTemp domain.Patient
		if err := ctx.BindJSON(&patientTemp); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		createdPatient, err := h.service.Create(patientTemp)
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 201, createdPatient)
	}
}

func (h *PatientHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		patientList, err := h.service.GetAll()
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, patientList)
	}
}

func (h *PatientHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		patientTemp, err := h.service.GetByID(int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, patientTemp)
	}
}

func (h *PatientHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		var patientTemp domain.Patient
		if err := ctx.BindJSON(&patientTemp); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		updatedPatient, err := h.service.Update(patientTemp, int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, updatedPatient)
	}
}

func (h *PatientHandler) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		var patientTemp domain.Patient
		if err := ctx.BindJSON(&patientTemp); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		patchedPatient, err := h.service.Patch(patientTemp, int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 200, patchedPatient)
	}
}

func (h *PatientHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		err = h.service.Delete(int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, "Deleted patient with id "+strconv.FormatInt(int64(id), 10))
	}
}
