package handlers

import (
    "net/http"
    "strconv"

    "github.com/Matheus-Armando/go-api/api/clients"
    "github.com/Matheus-Armando/go-api/api/models"
    "github.com/gin-gonic/gin"
)

// UserHandler handles user-related requests
type UserHandler struct {
    jsonClient *clients.JSONServerClient
}

// NewUserHandler creates a new user handler
func NewUserHandler() *UserHandler {
    return &UserHandler{
        jsonClient: clients.NewJSONServerClient(),
    }
}

// GetUsers handles the GET /users request
func (h *UserHandler) GetUsers(c *gin.Context) {
    var users []models.User
    err := h.jsonClient.Get("users", &users)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}

// GetUserByID handles the GET /users/:id request
func (h *UserHandler) GetUserByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

    var users []models.User
    err = h.jsonClient.Get("users", &users)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    for _, user := range users {
        if user.ID == id {
            c.JSON(http.StatusOK, user)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}