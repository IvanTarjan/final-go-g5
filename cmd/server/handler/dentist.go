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

// HandlerCreate is a function that creates a dentist
// @Summary Create dentist
// @Description Create a new dentist
// @Tags dentist
// @Accept json
// @Produce json
// @Param dentist body domain.Dentist true "New Dentist"
// @Param token header string true "Token for authentication"
// @Success 201 {object} domain.Dentist "Dentist successfully created"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /dentist [post]
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

// HandlerGetAll retrieves all dentists
// @Summary Retrieve all dentists
// @Description Get a list of all dentists
// @Tags dentist
// @Produce json
// @Success 200 {array} domain.Dentist "List of Dentists"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /dentists [get]
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

// HandlerGetById retrieves a dentist by id
// @Summary Retrieve a dentist by id
// @Description Get a single dentist by their id
// @Tags dentist
// @Produce json
// @Param id path int true "Dentist ID"
// @Success 200 {object} domain.Dentist "Dentist Data"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /dentist/{id} [get]
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

// HandlerUpdate updates a dentist's information
// @Summary Update a dentist
// @Description Update a dentist's information by their id
// @Tags dentist
// @Accept json
// @Produce json
// @Param id path int true "Dentist ID"
// @Param token header string true "Token for authentication"
// @Param dentist body domain.Dentist true "Updated Dentist Data"
// @Success 200 {object} domain.Dentist "Updated Dentist Data"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /dentist/{id} [put]
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

// HandlerPatch partially updates a dentist's information
// @Summary Partially update a dentist
// @Description Partially update a dentist's information by their id
// @Tags dentist
// @Accept json
// @Produce json
// @Param id path int true "Dentist ID"
// @Param token header string true "Token for authentication"
// @Param dentist body domain.Dentist true "Partial Dentist Data"
// @Success 200 {object} domain.Dentist "Patched Dentist Data"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /dentist/{id} [patch]
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

// HandlerDelete deletes a dentist
// @Summary Delete a dentist
// @Description Delete a dentist by their id
// @Tags dentist
// @Produce json
// @Param id path int true "Dentist ID"
// @Param token header string true "Token for authentication"
// @Success 200 {string} string "Deleted dentist with id {id}"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /dentist/{id} [delete]
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
