package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"rock-n-roll/database"
	"rock-n-roll/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func InsertStudent(c *gin.Context) {
	var student models.Student
	err := c.BindJSON(&student)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
	}

	var exists int
	errRow := database.DB.QueryRow(`SELECT 1 FROM STUDENTS WHERE email = $1`, student.Email).Scan(&exists)
	if errRow == sql.ErrNoRows {
		fmt.Println("No student with this email")
	} else if errRow != nil {
		fmt.Println("Query error:", errRow)
	} else {
		fmt.Println("Student with this email exists")
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Student with this email exists"})
		return
	}

	query := `INSERT INTO STUDENTS (name, email) VALUES($1,$2)`
	_, errDB := database.DB.Exec(query, student.Name, student.Email)
	if errDB != nil {
		fmt.Println(errDB)
		// panic(errDB)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Student Inserted Successfully!"})

}

func GetStudent(c *gin.Context) {
	rows, err := database.DB.Query(`SELECT * FROM STUDENTS`)

	if err != nil {
		fmt.Println("Error whike fetching data", err)
	}

	defer rows.Close()

	var students []models.Student

	for rows.Next() {
		var id int
		var name, email string
		var createdAt time.Time

		err := rows.Scan(&id, &name, &email, &createdAt)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue // or return depending on use case
		}
		newStudent := models.Student{
			ID:        id,
			Name:      name,
			Email:     email,
			CreatedAt: createdAt,
		}
		students = append(students, newStudent)
	}

	c.IndentedJSON(http.StatusOK, students)

}

func GetStudentById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid student ID"})
		return
	}
	var student models.Student
	err = database.DB.QueryRow(`SELECT id,name,email,created_at FROM STUDENTS WHERE id =$1`, id).Scan(&student.ID, &student.Name, &student.Email, &student.CreatedAt)
	if err != nil {
		fmt.Println("Error while fetching the student with id:", id)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid student ID"})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Println("Invalid Id", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	var body models.Student
	err = c.BindJSON(&body)
	if err != nil {
		fmt.Println("Invalid Body", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Body"})
		return
	}
	var student models.Student
	err = database.DB.QueryRow(`SELECT * FROM STUDENTS WHERE id = $1`, id).Scan(&student.ID, &student.Name, &student.Email, &student.CreatedAt)
	if err != nil {
		fmt.Println("Invalid Id", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	fmt.Println(body)
	query := `UPDATE STUDENTS SET name = $1, email=$2 WHERE id = $3`
	results, err := database.DB.Exec(query, body.Name, body.Email, id)
	if err != nil {
		fmt.Println("Error while update", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Failed to Update"})
		return
	}
	rowsAffected, _ := results.RowsAffected()
	fmt.Println(rowsAffected)

	if rowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Failed to Update"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Student updated successfully"})

}

func DeleteStudent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Println("Invalid Id", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	var exists bool
	err = database.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM STUDENTS WHERE id=$1)`, id).Scan(&exists)
	if err != nil || !exists {
		fmt.Println("err:", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid ID"})
		return
	}

	results, err := database.DB.Exec(`DELETE FROM STUDENTS WHERE id=$1`, id)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Invalid ID"})
		return
	}

	rowsAffected, _ := results.RowsAffected()
	if rowsAffected == 0 {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Invalid ID"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Student Deleted successfully"})

}
