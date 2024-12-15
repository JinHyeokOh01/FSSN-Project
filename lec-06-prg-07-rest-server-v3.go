package main

import (
	"github.com/gin-gonic/gin"
)

type MembershipHandler struct {
	database map[string]string
}

func (h *MembershipHandler) Create(c *gin.Context) {
	id := c.Param("member_id")
	value := c.PostForm(id)

	if _, exists := h.database[id]; exists {
		c.JSON(200, gin.H{id: "None"})
		return
	}
	h.database[id] = value
	c.JSON(200, gin.H{id: h.database[id]})
}

func (h *MembershipHandler) Read(c *gin.Context) {
	id := c.Param("member_id")

	if value, exists := h.database[id]; exists {
		c.JSON(200, gin.H{id: value})
		return
	}
	c.JSON(200, gin.H{id: "None"})
}

func (h *MembershipHandler) Update(c *gin.Context) {
	id := c.Param("member_id")
	value := c.PostForm(id)

	if _, exists := h.database[id]; !exists {
		c.JSON(200, gin.H{id: "None"})
		return
	}
	h.database[id] = value
	c.JSON(200, gin.H{id: h.database[id]})
}

func (h *MembershipHandler) Delete(c *gin.Context) {
	id := c.Param("member_id")

	if _, exists := h.database[id]; !exists {
		c.JSON(200, gin.H{id: "None"})
		return
	}
	delete(h.database, id)
	c.JSON(200, gin.H{id: "Removed"})
}

func main() {
	router := gin.Default()
	
	handler := &MembershipHandler{
		database: make(map[string]string),
	}

	api := router.Group("/membership_api")
	{
		api.POST("/:member_id", handler.Create)
		api.GET("/:member_id", handler.Read)
		api.PUT("/:member_id", handler.Update)
		api.DELETE("/:member_id", handler.Delete)
	}

	router.Run(":5000")
}