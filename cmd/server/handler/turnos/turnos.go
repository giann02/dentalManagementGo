package turnos

import (
	"net/http"
	"strconv"

	"github.com/giann02/finalBackEndGo/internal/domain"
	"github.com/giann02/finalBackEndGo/internal/turnos"
	"github.com/giann02/finalBackEndGo/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service turnos.Service
}

func NewControladorTurno(service turnos.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}


// Turno godoc
// @Summary turno example
// @Description Get all turnos
// @Tags turno
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /turnos [get]
func (c *Controlador) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		turnos, err := c.service.GetAll(ctx)

		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, turnos)
	}
}

// Turno godoc
// @Summary turno example
// @Description Get turno by id
// @Tags turno
// @Param id path int true "id del turno"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turno/:id [get]
func (c *Controlador) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		turno, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, turno)
	}
}

// Turno godoc
// @Summary turno example
// @Description Update turno by id
// @Tags turno
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/:id [put]
func (c *Controlador) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.Turno

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request param")
			return
		}

		turno, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, turno)

	}
}

// Turno godoc
// @Summary turno example
// @Description Delete turno by id
// @Tags turno
// @Param id path int true "id del turno"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turno/:id [delete]
func (c *Controlador) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"mensaje": "turno eliminado",
		})
	}
}

// Turno godoc
// @Summary turno example
// @Description Patch turno
// @Tags turno
// @Param id path int true "id del turno"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turno/:id [patch]
func (c *Controlador) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		var request domain.Turno

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		turno, err := c.service.Patch(ctx, request, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, turno)
	}
}

// Turno godoc
// @Summary turno example
// @Description Get turno by paciente id
// @Tags turno
// @Query id path int true "id del paciente"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turno/pacienteDni/:id [get]
// Manejador para obtener el detalle de un turno por el DNI del paciente.
func (c *Controlador) GetTurnoByDNIDetalleHandler(ctx *gin.Context) {
	dniPaciente := ctx.Query("dniPaciente")

	if dniPaciente == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere el DNI del paciente en los Query Parameters"})
		return
	}

	// Llama al servicio para obtener el detalle del turno.
	turnoDetalle, err := c.service.GetTurnoByDNIDetalle(ctx.Request.Context(), dniPaciente)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el detalle del turno: " + err.Error()})
		return
	}
	// Devuelve el detalle del turno como respuesta.
	ctx.JSON(http.StatusOK, turnoDetalle)
}

// Turno godoc
// @Summary turno example
// @Description Post turno
// @Tags turno
// @Query id path string true "matricula del dentista"
// @Query id path string true "dni del paciente"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /turnos [post]
// Manejador para la creación de un turno por DNI del paciente y matrícula del dentista.
func (c *Controlador) CreateTurnoByDNIMatriculaHandler(ctx *gin.Context) {
	var turno domain.Turno
	if err := ctx.BindJSON(&turno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar el cuerpo de la solicitud"})
		return
	}

	dniPaciente := ctx.Query("dniPaciente")
	matriculaDentista := ctx.Query("matriculaDentista")

	if dniPaciente == "" || matriculaDentista == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere el DNI del paciente y la matrícula del dentista"})
		return
	}

	turnoID, err := c.service.CreateTurnoByDNIMatricula(ctx.Request.Context(), turno, dniPaciente, matriculaDentista)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el turno: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"turnoID": turnoID})
}