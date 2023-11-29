package handler

import (
	"strconv"

	"github.com/IvanTarjan/final-go-g5/internal/dentists"
	"github.com/IvanTarjan/final-go-g5/internal/domain"
	"github.com/IvanTarjan/final-go-g5/pkg/web"
	"github.com/gin-gonic/gin"
)

type DentistHandler struct {
	service dentists.ServiceDentists
}

func NewDentistHandlerDentist(service dentists.ServiceDentists) *DentistHandler {
	return &DentistHandler{service: service}
}

func (h *DentistHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentist domain.Dentist
		if err := ctx.BindJSON(&dentist); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		createdDentist, err := h.service.Create(dentist)
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 201, createdDentist)
	}
}

func (h *DentistHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		dentist, err := h.service.GetByID(int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, dentist)
	}
}

func (h *DentistHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		var dentist domain.Dentist
		if err := ctx.BindJSON(&dentist); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		updatedDentist, err := h.service.Update(dentist, int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, updatedDentist)
	}
}

func (h *DentistHandler) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		var dentist domain.Dentist
		if err := ctx.BindJSON(&dentist); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		patchedDentist, err := h.service.Patch(dentist, int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 200, patchedDentist)
	}
}

func (h *DentistHandler) HandlerDelete() gin.HandlerFunc {
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
