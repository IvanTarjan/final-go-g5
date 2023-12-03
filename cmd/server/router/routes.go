package router

import (
	"database/sql"

	"github.com/IvanTarjan/final-go-g5/cmd/server/handler"
	"github.com/IvanTarjan/final-go-g5/internal/dentist"
	"github.com/IvanTarjan/final-go-g5/internal/patient"
	"github.com/IvanTarjan/final-go-g5/internal/turn"
	"github.com/IvanTarjan/final-go-g5/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
	db          *sql.DB
}

func NewRouter(engine *gin.Engine, db *sql.DB) Router {
	return &router{
		engine: engine,
		db:     db,
	}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildPatientRoutes()
	r.buildDentistRoutes()
	r.buildTurnRoutes()
}

func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

func (r *router) buildPatientRoutes() {
	patientRepository := patient.NewPatientsSqlRepository(r.db)
	patientService := patient.NewServicePatient(patientRepository)
	patientHandler := handler.NewPatientHandler(&patientService)

	patientGroup := r.routerGroup.Group("/patient")
	{
		patientGroup.POST("", middleware.Authenticate(), patientHandler.HandlerCreate())
		// patientGroup.GET("", patientHandler.GetAll())
		patientGroup.GET("/:id", middleware.Authenticate(), patientHandler.HandlerGetById())
		patientGroup.PUT("/:id", middleware.Authenticate(), patientHandler.HandlerUpdate())
		patientGroup.PATCH(("/:id"), middleware.Authenticate(), patientHandler.HandlerPatch())
		patientGroup.DELETE("/:id", middleware.Authenticate(), patientHandler.HandlerDelete())
	}
}

func (r *router) buildDentistRoutes() {
	dentistRepository := dentist.NewDentistSqlRepository(r.db)
	dentistService := dentist.NewServiceDentist(dentistRepository)
	dentistHandler := handler.NewDentistHandler(&dentistService)

	dentistGroup := r.routerGroup.Group("/dentist")
	{
		dentistGroup.POST("", middleware.Authenticate(), dentistHandler.HandlerCreate())
		// dentistGroup.GET("", dentistHandler.GetAll())
		dentistGroup.GET("/:id", dentistHandler.HandlerGetById())
		dentistGroup.PUT("/:id", middleware.Authenticate(), dentistHandler.HandlerUpdate())
		dentistGroup.PATCH("/:id", middleware.Authenticate(), dentistHandler.HandlerPatch())
		dentistGroup.DELETE("/:id", middleware.Authenticate(), dentistHandler.HandlerDelete())
	}
}

func (r *router) buildTurnRoutes() {
	turnRepository := turn.NewTurnSqlRepository(r.db)
	turnService := turn.NewServiceTurn(turnRepository)
	turnHandler := handler.NewTurnHandler(&turnService)

	turnGroup := r.routerGroup.Group("/turn")
	{
		turnGroup.POST("", middleware.Authenticate(), turnHandler.HandlerCreate())
		turnGroup.GET("", turnHandler.HandlerGetAll())
		turnGroup.GET("/:id", turnHandler.HandlerGetById())
		turnGroup.PUT("/:id", middleware.Authenticate(), turnHandler.HandlerUpdate())
		turnGroup.PATCH("/:id", middleware.Authenticate(), turnHandler.HandlerPatch())
		turnGroup.DELETE("/:id", middleware.Authenticate(), turnHandler.HandlerDelete())
	}
}
