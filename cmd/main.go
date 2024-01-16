package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gl1n0m3c/learning_gin/internal/config"
	"github.com/gl1n0m3c/learning_gin/pkg/logging"
)

func main() {
	server := gin.Default()

	logging.InitLogger()
	logger, loggerFile := logging.GetLogger()
	defer loggerFile.Close()
	logger.Infoln("Logger initialized")

	cfg, err := config.InitConfig()
	if err != nil {
		logger.Errorf("Config error: %v\n", err)
		return
	}
	logger.Infof("Config initialized: %v", *cfg)
	logger.Infoln("Test")

	server.Run(":8000")
}
