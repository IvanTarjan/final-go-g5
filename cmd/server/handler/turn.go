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
