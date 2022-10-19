package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/GinAPIRest-go/controllers"
	"github.com/GinAPIRest-go/database"
	"github.com/GinAPIRest-go/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var id int

func SetupRoutesTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{Nome: "Júpiter", CPF: "123456789", RG: "12345678912"}
	database.DB.Create(&student)
	id = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student)
}

func TestHomeStatusCode(t *testing.T) {
	r := SetupRoutesTest()
	r.GET("/:nome", controllers.Home)
	req, _ := http.NewRequest("GET", "/zephyr", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code,
		"Status error: ---Expected %d +++Actual %d", http.StatusOK, res.Code)
	mockRes := `{"message":"zephyr"}`
	bodyRes, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, mockRes, string(bodyRes))
}

func TestAllStudent(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.GET("/alunos", controllers.AllStudents)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestStudentByCPF(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.GET("/alunos/cpf/:cpf", controllers.StudentByCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/123456789", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestStudentByID(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.GET("/alunos/:id", controllers.StudentByID)
	path := "/alunos/" + strconv.Itoa(id)
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var studentMock models.Student
	json.Unmarshal(res.Body.Bytes(), &studentMock)
	assert.Equal(t, "Júpiter", studentMock.Nome)
	assert.Equal(t, "123456789", studentMock.CPF)
	assert.Equal(t, "12345678912", studentMock.RG)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()
	r := SetupRoutesTest()
	r.DELETE("/alunos/:id", controllers.DeleteStudent)
	path := "/alunos/" + strconv.Itoa(id)
	req, _ := http.NewRequest("DELETE", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestEditStudent(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.PATCH("/alunos/:id", controllers.EditStudent)
	student := models.Student{Nome: "Júpite", CPF: "123456780", RG: "12345678911"}
	studentMockJson, _ := json.Marshal(student)
	path := "/alunos/" + strconv.Itoa(id)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(studentMockJson))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var studentMockUpdated models.Student
	json.Unmarshal(res.Body.Bytes(), &studentMockUpdated)
	assert.Equal(t, "123456780", studentMockUpdated.CPF)
	assert.Equal(t, "12345678911", studentMockUpdated.RG)
	assert.Equal(t, "Júpite", studentMockUpdated.Nome)
}
