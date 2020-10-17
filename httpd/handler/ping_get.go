package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ping struct {
	Status string `json:"status"`
}

func Ping() gin.HandlerFunc {

	return func(context *gin.Context) {
		res := ping{
			Status: "SUCCESS",
		}
		context.JSON(http.StatusOK, res)
	}
}