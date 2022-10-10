package main

import (
	"github.com/gin-gonic/gin"
	"go.h4n.io/openschool/class/handlers/classes"
	"go.h4n.io/openschool/class/models"
	"go.h4n.io/openschool/class/repos/class"
)

func main() {
	e := gin.Default()

	classRepo := class.InMemoryClassRepository{
		Items: []models.Class{},
	}
	classHandler := classes.NewClassesHandler(&classRepo)
	classHandler.RegisterRoutes(e)

	e.Run("0.0.0.0:8081")
}
