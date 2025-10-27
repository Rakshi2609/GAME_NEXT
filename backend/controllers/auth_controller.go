package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"codechef/backend/config"
	"codechef/backend/models"
	"codechef/backend/utils"
)

type RegisterDTO struct {
	Name     string      `json:"name" binding:"required"`
	Email    string      `json:"email" binding:"required,email"`
	Password string      `json:"password" binding:"required,min=6"`
	Role     models.Role `json:"role" binding:"required,oneof=creator attendee"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var dto RegisterDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	col := config.DB().Collection("users")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	// unique email check
	var existing models.User
	err := col.FindOne(ctx, bson.M{"email": dto.Email}).Decode(&existing)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		return
	}

	hash, err := utils.HashPassword(dto.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password"})
		return
	}

	u := models.User{
		ID:       primitive.NewObjectID(),
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hash,
		Role:     dto.Role,
	}
	if _, err := col.InsertOne(ctx, u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateJWT(os.Getenv("JWT_SECRET"), u.ID.Hex(), string(u.Role), 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not issue token"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"token": token, "user": gin.H{"id": u.ID.Hex(), "name": u.Name, "role": u.Role}})
}

func Login(c *gin.Context) {
	var dto LoginDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	col := config.DB().Collection("users")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var u models.User
	if err := col.FindOne(ctx, bson.M{"email": dto.Email}).Decode(&u); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	if !utils.CheckPassword(u.Password, dto.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	token, err := utils.GenerateJWT(os.Getenv("JWT_SECRET"), u.ID.Hex(), string(u.Role), 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not issue token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "user": gin.H{"id": u.ID.Hex(), "name": u.Name, "role": u.Role}})
}
