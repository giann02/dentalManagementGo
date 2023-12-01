package dentistas

import (
	"net/http"
	"strconv"

	"github.com/giann02/finalBackEndGo/internal/domain"
	"github.com/giann02/finalBackEndGo/internal/dentistas"
	"github.com/giann02/finalBackEndGo/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service dentistas.Service
}

func NewControladorDentista(service dentistas.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}

// Dentista godoc
// @Summary dentista example
// @Description Create a new dentista
// @Tags dentista
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentistas [post]
func (c *Controlador) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.Dentista

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		dentista, err := c.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, dentista)

	}
}

// Dentista godoc
// @Summary dentista example
// @Description Get all dentistas
// @Tags dentista
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /dentistas [get]
func (c *Controlador) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dentistas, err := c.service.GetAll(ctx)

		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, dentistas)
	}
}

// Dentista godoc
// @Summary dentista example
// @Description Get dentista by id
// @Tags dentista
// @Param id path int true "id del dentista"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentista/:id [get]
func (c *Controlador) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		dentista, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, dentista)
	}
}

// Dentista godoc
// @Summary dentista example
// @Description Update dentista by id
// @Tags dentista
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentistas/:id [put]
func (c *Controlador) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.Dentista

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

		dentista, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, dentista)

	}
}

// Dentista godoc
// @Summary dentista example
// @Description Delete dentista by id
// @Tags dentista
// @Param id path int true "id del dentista"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentista/:id [delete]
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
			"mensaje": "dentista eliminado",
		})
	}
}

// Dentista godoc
// @Summary dentista example
// @Description Patch dentista
// @Tags dentista
// @Param id path int true "id del dentista"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentista/:id [patch]
func (c *Controlador) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		var request domain.Dentista

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		dentista, err := c.service.Patch(ctx, request, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, dentista)
	}
}