package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gl1n0m3c/learning_gin/internal/config"
)

func main() {
	server := gin.Default()

	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Printf("Config error: %v\n", err)
		return
	}
	fmt.Println(*cfg)

	server.Run(":8000")
}
