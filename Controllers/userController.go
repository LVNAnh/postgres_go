package Controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "postgres_go/Models"
    "postgres_go/config"
)

func GetUsers(c *gin.Context) {
    var users []Models.User
    result := config.DB.Find(&users)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func AddUser(c *gin.Context) {
    var user Models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := config.DB.Create(&user)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "User created successfully!",
        "user":    user,
    })
}