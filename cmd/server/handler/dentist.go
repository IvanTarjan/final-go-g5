package handler

import (
	"strconv"

	"github.com/IvanTarjan/final-go-g5/internal/dentist"
	"github.com/IvanTarjan/final-go-g5/internal/domain"
	"github.com/IvanTarjan/final-go-g5/pkg/web"
	"github.com/gin-gonic/gin"
)

type DentistHandler struct {
	service dentist.ServiceDentists
}

func NewDentistHandler(service *dentist.ServiceDentists) *DentistHandler {
	return &DentistHandler{service: *service}
}

func (h *DentistHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentistTemp domain.Dentist
		if err := ctx.BindJSON(&dentistTemp); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		createdDentist, err := h.service.Create(dentistTemp)
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 201, createdDentist)
	}
}

func (h *DentistHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dentistList, err := h.service.GetAll()
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, dentistList)
	}
}

func (h *DentistHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		dentistTemp, err := h.service.GetByID(int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, dentistTemp)
	}
}

func (h *DentistHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		var dentistTemp domain.Dentist
		if err := ctx.BindJSON(&dentistTemp); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		updatedDentist, err := h.service.Update(dentistTemp, int64(id))
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

		var dentistTemp domain.Dentist
		if err := ctx.BindJSON(&dentistTemp); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		patchedDentist, err := h.service.Patch(dentistTemp, int64(id))
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
