package classes

import (
	"github.com/gin-gonic/gin"
	"go.h4n.io/openschool/services/class/repos/class"
)

type ClassesHandler struct {
	Repository class.ClassRepository
}

func NewClassesHandler(repository class.ClassRepository) ClassesHandler {
	h := ClassesHandler{
		Repository: repository,
	}
	return h
}

func (h *ClassesHandler) RegisterRoutes(e *gin.Engine) {
	group := e.Group("/classes")
	group.GET("/", func(ctx *gin.Context) {
		req := ClassesIndexRequest{}

		res, err := h.ClassesIndex(req)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, res)
	})
}
