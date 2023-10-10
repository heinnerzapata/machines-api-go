package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{"user": "heinner Register"})
}

func Login(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{"user": "heinner Login"})
}
