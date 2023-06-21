package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Student struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Course string `json:"course"`
}

var students = [] Student{
	{ID: 1, Name: "Mariam", Course: "Biology"},
	{ID: 2, Name: "Alhaji", Course: "Chemistry"},
	{ID: 3, Name: "Duff", Course: "Auto-Mechanics"},
	{ID: 4, Name: "Grace", Course: "Life Skills"},
}

func getStudents(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, students)
}

func registerStudent(c *gin.Context){
	var newStudent Student

	if err := c.BindJSON(&newStudent); err !=nil{
		return
	}

	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}

func getStudentById(c *gin.Context){
	id := c.Param("id")
	studentId, _ := strconv.ParseInt(id, 8, 8)

	for _, student := range students {
		if student.ID == int(studentId) {
			c.IndentedJSON(http.StatusFound, student)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func removeStudent(c *gin.Context){
	id := c.Param("id")
	studentId, _ := strconv.ParseInt(id, 8, 8)

	for index, student := range students {
		if student.ID == int(studentId) {
			students = append(students[:index],students[index+1:]... )
			c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%v is deleted", student)})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func updateStudent(c *gin.Context){
	id := c.Param("id")
	studentId, _ := strconv.ParseInt(id, 8, 8)

	var newStudent Student

	if err := c.BindJSON(&newStudent); err !=nil{
		return
	}

	for index, student := range students{
		if student.ID == int(studentId){

			if newStudent.Course != "" {
				students[index].Course = newStudent.Course
			}
			if newStudent.Name != "" {
				students[index].Name = newStudent.Name
			}
			c.IndentedJSON(http.StatusOK, gin.H{"message": "student details updated"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func main(){
	router := gin.Default()
    router.GET("/students", getStudents)
	router.POST("/students", registerStudent)
	router.GET("/students/:id", getStudentById)
	router.DELETE("/students/:id", removeStudent)
	router.PUT("/students/:id", updateStudent)

    router.Run("localhost:8080")	
}


