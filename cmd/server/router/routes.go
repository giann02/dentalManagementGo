package routes

import (
	"database/sql"

	"github.com/giann02/finalBackEndGo/pkg/middleware"

	"github.com/giann02/finalBackEndGo/cmd/server/handler/ping"
	handlerDentista "github.com/giann02/finalBackEndGo/cmd/server/handler/dentistas"
	handlerPaciente "github.com/giann02/finalBackEndGo/cmd/server/handler/pacientes"
	handlerTurno "github.com/giann02/finalBackEndGo/cmd/server/handler/turnos"
	dentista "github.com/giann02/finalBackEndGo/internal/dentistas"
	paciente "github.com/giann02/finalBackEndGo/internal/pacientes"
	turno "github.com/giann02/finalBackEndGo/internal/turnos"
	"github.com/gin-gonic/gin"
)

// Router interface defines the methods that any router must implement.
type Router interface {
	MapRoutes()
}

// router is the Gin router.
type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
	db          *sql.DB
}

// NewRouter creates a new Gin router.
func NewRouter(engine *gin.Engine, db *sql.DB) Router {
	return &router{
		engine: engine,
		db:     db,
	}
}

// MapRoutes maps all routes.
func (r *router) MapRoutes() {
	r.setGroup()
	r.buildDentistaRoutes()
	r.buildPacienteRoutes()
	r.buildTurnoRoutes()
	r.buildPingRoutes()
}

// setGroup sets the router group.
func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

// buildDentistaRoutes maps all routes for the dentista domain.
func (r *router) buildDentistaRoutes() {
	// Create a new dentista controller.
	repository := dentista.NewMySqlRepository(r.db)
	service := dentista.NewServiceDentista(repository)
	controlador := handlerDentista.NewControladorDentista(service)

	grupoDentista := r.routerGroup.Group("/dentista")
	{
		grupoDentista.POST("", middleware.Authenticate(), controlador.HandlerCreate())
		grupoDentista.GET("", middleware.Authenticate(), controlador.HandlerGetAll())
		grupoDentista.GET("/:id", controlador.HandlerGetByID())
		grupoDentista.PUT("/:id", middleware.Authenticate(), controlador.HandlerUpdate())
		grupoDentista.DELETE("/:id", middleware.Authenticate(), controlador.HandlerDelete())
		grupoDentista.PATCH("/:id", middleware.Authenticate(), controlador.HandlerPatch())
	}


}

// buildPacienteRoutes maps all routes for the paciente domain.
func (r *router) buildPacienteRoutes() {
		// Create a new paciente controller.
		repository := paciente.NewMySqlRepository(r.db)
		service := paciente.NewServicePaciente(repository)
		controlador := handlerPaciente.NewControladorPaciente(service)

	grupoPaciente := r.routerGroup.Group("/paciente")
	{
		grupoPaciente.POST("", middleware.Authenticate(), controlador.HandlerCreate())
		grupoPaciente.GET("", middleware.Authenticate(), controlador.HandlerGetAll())
		grupoPaciente.GET("/:id", controlador.HandlerGetByID())
		grupoPaciente.PUT("/:id", middleware.Authenticate(), controlador.HandlerUpdate())
		grupoPaciente.DELETE("/:id", middleware.Authenticate(), controlador.HandlerDelete())
		grupoPaciente.PATCH("/:id", middleware.Authenticate(), controlador.HandlerPatch())
	}

}

// buildTurnoRoutes maps all routes for the turno domain.
func (r *router) buildTurnoRoutes() {
	// Create a new turno controller.
	repository := turno.NewMySqlRepository(r.db)
	service := turno.NewServiceTurno(repository)
	controlador := handlerTurno.NewControladorTurno(service)

grupoTurno := r.routerGroup.Group("/turno")
{
	grupoTurno.GET("", middleware.Authenticate(), controlador.HandlerGetAll())
	grupoTurno.GET("/:id", controlador.HandlerGetByID())
	grupoTurno.PUT("/:id", middleware.Authenticate(), controlador.HandlerUpdate())
	grupoTurno.DELETE("/:id", middleware.Authenticate(), controlador.HandlerDelete())
	grupoTurno.PATCH("/:id", middleware.Authenticate(), controlador.HandlerPatch())
	grupoTurno.GET("/getByDni", middleware.Authenticate(), controlador.GetTurnoByDNIDetalleHandler)
	grupoTurno.POST("", middleware.Authenticate(), controlador.CreateTurnoByDNIMatriculaHandler)
}

}

// buildPingRoutes maps all routes for the ping domain.
func (r *router) buildPingRoutes() {
	// Create a new ping controller.
	pingController := ping.NewControllerPing()
	r.routerGroup.GET("/ping", pingController.HandlerPing())

}