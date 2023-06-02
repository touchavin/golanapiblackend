package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	
)

func main() {
	r := gin.New()

	r.GET("/books", listBooksHandler)
	r.POST("/books", createBookHandler)
	r.PUT("/books/:id", updateBookHandler)
	r.DELETE("/books/:id", deleteBookHandler)
	

	r.Run()
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pea string `json:"pea"`
}

var books = []Book{
	// {ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	// {ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	// {ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func listBooksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}



func createBookHandler(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	books = append(books, book)

	c.JSON(http.StatusCreated, book)
}

func updateBookHandler(c *gin.Context) {
	id := c.Param("id")

	for i, a := range books {
		if a.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
	
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	books = append(books, book)

	c.JSON(http.StatusCreated, book)
}

func deleteBookHandler(c *gin.Context) {
	id := c.Param("id")

	for i, a := range books {
		if a.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}




// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// Create a new Gin router
// 	router := gin.Default()

// 	// Define your API endpoints
// 	router.GET("/api/users", getUsers)
// 	router.POST("/api/users", createUser)
// 	router.GET("/api/users/:id", getUser)
// 	router.PUT("/api/users/:id", updateUser)
// 	router.DELETE("/api/users/:id", deleteUser)

// 	// Start the server
// 	router.Run(":8080")
// }

// func getUsers(c *gin.Context) {
// 	// Logic to fetch and return a list of users
// 	c.JSON(200, gin.H{"message": "Get all users"})
// }

// func createUser(c *gin.Context) {
// 	// Logic to create a new user
// 	c.JSON(200, gin.H{"message": "Create user"})
// }

// func getUser(c *gin.Context) {
// 	// Logic to fetch and return a specific user
// 	userID := c.Param("id")
// 	c.JSON(200, gin.H{"message": "Get user " + userID})
// }

// func updateUser(c *gin.Context) {
// 	// Logic to update a specific user
// 	userID := c.Param("id")
// 	c.JSON(200, gin.H{"message": "Update user " + userID})
// }

// func deleteUser(c *gin.Context) {
// 	// Logic to delete a specific user
// 	userID := c.Param("id")
// 	c.JSON(200, gin.H{"message": "Delete user " + userID})
// }
