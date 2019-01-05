package controllers

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func MainHome(c *gin.Context) {
	hello := os.Getenv("MESSAGE")
	fmt.Println("==>" + hello)
}
