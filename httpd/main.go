package main

import (
	"curriculumvitae/httpd/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/api/go", handler.Ping())
	r.GET("/api/go/user", handler.UserInformation())


	e := r.Run(":8082")

	if e != nil{
		panic("Port not available\n" + e.Error())
	}
}
