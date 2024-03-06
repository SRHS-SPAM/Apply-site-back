package controllers

import "github.com/gin-gonic/gin"

func NewController(port string) {
	r := gin.Default()

	err := r.Run(port)
	if err != nil {

	}
}
