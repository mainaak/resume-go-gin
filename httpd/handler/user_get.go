package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Designation string `json:"designation"`
	Location string `json:"location"`
	Email string `json:"email_id"`
	LinkedIn string `json:"linked_in_id"`
	Github string `json:"github_id"`
}

func UserInformation() gin.HandlerFunc {

	return func(context *gin.Context) {
		user := User{
			Name:        "Mainaak Aanand",
			Designation: "Software Quality Engineer",
			Location:    "Gurgaon, Haryana",
			Email:       "mainaakaanand@gmail.com",
			LinkedIn:    "https://linkedin.com/in/mainaakaanand",
			Github:      "https://github.com/mainaak",
		}
		context.JSON(http.StatusOK, user)
	}
}
