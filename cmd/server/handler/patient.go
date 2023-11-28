package handler

import (
	"strconv"
	"github.com/IvanTarjan/final-go-g5/internal/domain"
	"github.com/IvanTarjan/final-go-g5/internal/patient"
	"github.com/IvanTarjan/final-go-g5/pkg/web"
	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	service *patient.Service
}

func NewPatientHandler(service *patient.Service) *PatientHandler {
    return &PatientHandler{service}
}

func (h *PatientHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
        var patient domain.Patient
        if err := ctx.BindJSON(&patient); err!= nil {
			web.Failure(ctx, 400, err)
            return
        }
        createdPatient, err := h.service.Post(patient)
        if err!= nil {
            web.Failure(ctx, 400, err)
            return
        }
		web.Success(ctx, 201, createdPatient)
    }
}

func (h *PatientHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id") )
        if err!= nil {
            web.Failure(ctx, 400, err)
            return
        }
        patient, err := h.service.GetById(int64(id))
        if err!= nil {
            web.Failure(ctx, 400, err)
            return
        }
        web.Success(ctx, 200, patient)
    }
}

func ( h *PatientHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var patient domain.Patient
		if err := ctx.BindJSON(&patient) ; err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		updatedPatient, err := h.service.Put(patient)
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		web.Success(ctx, 200, updatedPatient)
	}
}

func (h *PatientHandler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
        var attributes map[string]string
		queryParams :=ctx.Request.URL.Query()
		for k, v := range queryParams {
			if len(v) >0{
				attributes[k] = v[0]
			}
		}
        patchedPatient, err := h.service.Patch(attributes)
        if err!= nil {
            web.Failure(ctx, 400, err)
            return
        }
        web.Success(ctx, 200, patchedPatient)
    }
}

func (h *PatientHandler) DeleteById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		err = h.service.DeleteById(int64(id))
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		web.Success(ctx, 200, "Deleted patient with id " + strconv.FormatInt(int64(id), 10) )
	}
}