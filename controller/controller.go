package controller

import (
	user "gin_project/Users"
	"gin_project/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, user.List)
}

func PostUser(c *gin.Context) {
	var newUser model.Model

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	user.List = append(user.List, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func GetByID(c *gin.Context) {
	id := c.Param("id")
	var found bool
	for _, value := range user.List {
		if value.ID == id {
			c.IndentedJSON(http.StatusOK, value)
			found = true
		}
	}
	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "user not found"})
	}
}

func DeleteByid(c *gin.Context) {
	id := c.Param("id")

	var filtered []model.Model

	for _, value := range user.List {
		if value.ID != id {
			filtered = append(filtered, value)
		}
	}
	user.List = filtered
	c.IndentedJSON(http.StatusOK, user.List)
}

func UpdateUser(c *gin.Context) {
	var newInfoUser model.Model
	if err := c.BindJSON(&newInfoUser); err != nil {
		return
	}
	id := newInfoUser.ID
	for index, value := range user.List {
		if value.ID == id {
			user.List[index] = newInfoUser
		}
	}
	c.IndentedJSON(http.StatusOK, newInfoUser)
}
