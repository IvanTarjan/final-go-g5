package handler

import (
	"strconv"

	"github.com/IvanTarjan/final-go-g5/internal/dentists"
	"github.com/IvanTarjan/final-go-g5/internal/domain"
	"github.com/IvanTarjan/final-go-g5/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service dentists.Service
}

func NewControllerDentist(service dentists.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
		}

		dentist, err := c.service.GetByID(int64(idParam))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
		}
		web.Success(ctx, 200, dentist)
	}
}

func (c *Controller) HandlerPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentistRequest domain.Dentist
		err := ctx.Bind(&dentistRequest)
		if err != nil {
			web.Failure(ctx, 400, err.Error())
		}

		dentist, err := c.service.Post(dentistRequest)
		if err != nil {
			web.Failure(ctx, 400, err.Error())
		}

		web.Success(ctx, 201, dentist)
	}
}

func (c *Controller) HandlerPut() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
		}

		var dentistRequest domain.Dentist
		bindErr := ctx.Bind(dentistRequest)
		if bindErr != nil {
			web.Failure(ctx, 400, err.Error())
		}

		dentist, err := c.service.Put(dentistRequest, int64(idParam))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
		}

		web.Success(ctx, 200, dentist)
	}
}

func (c *Controller) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (c *Controller) HandleDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
		}

		deleteError := c.service.Delete(int64(idParam))
		if deleteError != nil {
			web.Failure(ctx, 400, err.Error())
		}
		web.Success(ctx, 200, "Dentista Eliminado")
	}
}
