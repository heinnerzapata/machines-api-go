package controller

import (
	"context"
	"fmt"
	"machines-api-go/connect"
	"machines-api-go/db"
	"machines-api-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	body := models.CreateUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	var user = db.CreateUserParams{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	}

	userResult, err := connect.DataQueries.CreateUser(context.Background(), user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	ctx.JSON(http.StatusCreated, &userResult)
}

func Login(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{"user": "heinner Login"})
}
