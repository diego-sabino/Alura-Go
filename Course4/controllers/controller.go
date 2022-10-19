package controllers

import (
	"net/http"

	"github.com/GinAPIRest-go/database"
	"github.com/GinAPIRest-go/models"

	"github.com/gin-gonic/gin"
)

func AllStudents(c *gin.Context) {
	var Student []models.Student
	database.DB.Find(&Student)
	c.JSON(200, Student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var Student models.Student
	database.DB.Delete(&Student, id)
	c.JSON(200, Student)
}

func EditStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var Student models.Student
	database.DB.First(&Student, id)

	if err := c.ShouldBindJSON(&Student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateStudentData(&Student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&Student).UpdateColumns(Student)
	c.JSON(200, Student)
}

func StudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var Student models.Student
	database.DB.First(&Student, id)

	if Student.ID == 0 {
		c.JSON(404, gin.H{"message": "Aluno não encontrado"})
		return
	}
	c.JSON(200, Student)
}

func StudentByCPF(c *gin.Context) {
	cpf := c.Param("cpf")
	var Student models.Student
	database.DB.Where(&models.Student{CPF: cpf}).First(&Student)

	if Student.ID == 0 {
		c.JSON(404, gin.H{"message": "Aluno não encontrado"})
		return
	}
	c.JSON(200, Student)
}

func Home(c *gin.Context) {
	name := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"message": name,
	})
}

func CreateNewStudent(c *gin.Context) {
	var Student models.Student
	if err := c.ShouldBindJSON(&Student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidateStudentData(&Student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&Student)
	c.JSON(http.StatusOK, Student)
}
