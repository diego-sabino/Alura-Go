package routes

import (
	"github.com/GinAPIRest-go/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.AllStudents)
	r.GET("/:nome", controllers.Home)
	r.GET("/alunos/cpf/:cpf", controllers.StudentByCPF)
	r.GET("/alunos/:id", controllers.StudentByID)
	r.DELETE("/alunos/:id", controllers.DeleteStudent)
	r.PATCH("/alunos/:id", controllers.EditStudent)
	r.POST("/alunos", controllers.CreateNewStudent)
	r.Run()
}
