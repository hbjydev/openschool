package classes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.h4n.io/openschool/class/repos/class"
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
		perPageStr := ctx.Query("perPage")
		pageStr := ctx.Query("page")
		req := ClassesIndexRequest{}

		if perPageStr != "" {
			perPage, err := strconv.Atoi(perPageStr)
			if err != nil {
				ctx.Error(err)
				ctx.JSON(400, gin.H{"error": "perPage must be a number"})
				return
			}
			req.PerPage = perPage
		}

		if pageStr != "" {
			page, err := strconv.Atoi(pageStr)
			if err != nil {
				ctx.Error(err)
				ctx.JSON(400, gin.H{"error": "page must be a number"})
				return
			}
			req.Page = page
		}

		res, err := h.ClassesIndex(req)
		if err != nil {
			ctx.Error(err)
			ctx.JSON(500, gin.H{"error": "something went wrong trying to retrieve the list of classes"})
			return
		}

		ctx.JSON(200, res)
	})
}
