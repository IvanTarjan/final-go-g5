package handler

import (
	"strconv"

	"github.com/IvanTarjan/final-go-g5/internal/domain"
	"github.com/IvanTarjan/final-go-g5/internal/turn"
	"github.com/IvanTarjan/final-go-g5/pkg/web"
	"github.com/gin-gonic/gin"
)

type TurnHandler struct {
	service turn.ServiceTurn
}

func NewTurnHandler(service *turn.ServiceTurn) *TurnHandler {
	return &TurnHandler{service: *service}
}

// HandlerCreate is a function that creates a turn
// @Summary Create turn
// @Description Create a new turn
// @Tags turn
// @Accept json
// @Produce json
// @Param turn body domain.Turn true "New Turn"
// @Param token header string true "Token for authentication"
// @Success 201 {object} domain.Turn "Turn successfully created"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /turn [post]
func (h *TurnHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var turnTemp domain.Turn
		if err := ctx.BindJSON(&turnTemp); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		createdTurn, err := h.service.Create(turnTemp)
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 201, createdTurn)
	}
}

// HandlerGetAll retrieves all turns
// @Summary Retrieve all turns
// @Description Get a list of all turns
// @Tags turn
// @Produce json
// @Success 200 {array} domain.Turn "List of Turns"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /turns [get]
func (h *TurnHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		turnList, err := h.service.GetAll()
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, turnList)
	}
}

// HandlerGetById retrieves a turn by id
// @Summary Retrieve a turn by id
// @Description Get a single turn by their id
// @Tags turn
// @Produce json
// @Param id path int true "Turn ID"
// @Success 200 {object} domain.Turn "Turn Data"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /turn/{id} [get]
func (h *TurnHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		turnTemp, err := h.service.GetByID(int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, turnTemp)
	}
}

// HandlerGetByPatientDni retrieves turns by patient's DNI
// @Summary Retrieve turns by DNI
// @Description Get turns associated with a patient's DNI
// @Tags turn
// @Produce json
// @Param dni path string true "Patient DNI"
// @Success 200 {array} domain.Turn "List of Turns for Patient"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /turn/patient/{dni} [get]
func (h *TurnHandler) HandlerGetByPatientDni() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dni := ctx.Param("dni")

		turnTemp, err := h.service.GetByPatientDni(dni)
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, turnTemp)
	}
}

// HandlerUpdate updates a turn's information
// @Summary Update a turn
// @Description Update a turn's information by their id
// @Tags turn
// @Accept json
// @Produce json
// @Param id path int true "Turn ID"
// @Param token header string true "Token for authentication"
// @Param turn body domain.Turn true "Updated Turn Data"
// @Success 200 {object} domain.Turn "Updated Turn Data"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /turn/{id} [put]
func (h *TurnHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		var turnTemp domain.Turn
		if err := ctx.BindJSON(&turnTemp); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		updatedTurn, err := h.service.Update(turnTemp, int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, updatedTurn)
	}
}

// HandlerPatch partially updates a turn's information
// @Summary Partially update a turn
// @Description Partially update a turn's information by their id
// @Tags turn
// @Accept json
// @Produce json
// @Param id path int true "Turn ID"
// @Param token header string true "Token for authentication"
// @Param turn body domain.Turn true "Partial Turn Data"
// @Success 200 {object} domain.Turn "Patched Turn Data"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /turn/{id} [patch]
func (h *TurnHandler) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		var turnTemp domain.Turn
		if err := ctx.BindJSON(&turnTemp); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		updatedTurn, err := h.service.Patch(turnTemp, int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, updatedTurn)
	}
}

// HandlerDelete deletes a turn
// @Summary Delete a turn
// @Description Delete a turn by their id
// @Tags turn
// @Produce json
// @Param id path int true "Turn ID"
// @Param token header string true "Token for authentication"
// @Success 200 {string} string "Deleted turn with id {id}"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /turn/{id} [delete]
func (h *TurnHandler) HandlerDelete() gin.HandlerFunc {
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
