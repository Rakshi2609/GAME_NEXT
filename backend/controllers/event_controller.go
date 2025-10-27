package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"codechef/backend/config"
	"codechef/backend/middleware"
	"codechef/backend/models"
)

type CreateEventDTO struct {
	Name        string   `json:"name" binding:"required"`
	Venue       string   `json:"venue" binding:"required"`
	StartAt     string   `json:"startAt" binding:"required"`
	PrizePool   string   `json:"prizePool"`
	Logo        string   `json:"logo"`
	Sponsors    []string `json:"sponsors"`
	Coordinator []string `json:"coordinator"`
	Speaker     string   `json:"speaker"`
}

func CreateEvent(c *gin.Context) {
	uid, _ := c.Get(string(middleware.CtxUserID))
	var dto CreateEventDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t, err := time.Parse(time.RFC3339, dto.StartAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid startAt (use RFC3339)"})
		return
	}
	e := models.Event{
		ID:          primitive.NewObjectID(),
		Name:        dto.Name,
		Venue:       dto.Venue,
		StartAt:     t,
		PrizePool:   dto.PrizePool,
		Logo:        dto.Logo,
		Sponsors:    dto.Sponsors,
		Coordinator: dto.Coordinator,
		Speaker:     dto.Speaker,
		CreatedBy:   primitive.NilObjectID,
	}
	if s, ok := uid.(string); ok {
		if oid, err := primitive.ObjectIDFromHex(s); err == nil {
			e.CreatedBy = oid
		}
	}
	col := config.DB().Collection("events")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	if _, err := col.InsertOne(ctx, e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, e)
}

func ListEvents(c *gin.Context) {
	col := config.DB().Collection("events")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	cur, err := col.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var out []models.Event
	if err := cur.All(ctx, &out); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}

func GetEvent(c *gin.Context) {
	id := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	col := config.DB().Collection("events")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	var e models.Event
	if err := col.FindOne(ctx, bson.M{"_id": oid}).Decode(&e); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, e)
}

func UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var dto map[string]any
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	col := config.DB().Collection("events")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	if _, err := col.UpdateByID(ctx, oid, bson.M{"$set": dto}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	GetEvent(c)
}

func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	col := config.DB().Collection("events")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	if _, err := col.DeleteOne(ctx, bson.M{"_id": oid}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
