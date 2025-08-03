package api

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"main/traceroute"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"log"
)

var logger *log.Logger
var rdb *redis.Client

func TracerouteHandler(c *gin.Context) {
	dest := c.Query("dest")
	if dest == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing dest query parameter"})

		logger.Printf("[ERROR] Missing dest query parameter")

		return
	}
	result, err := traceroute.Traceroute(dest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"[ERROR]": err.Error()})

		logger.Printf("[ERROR] %v", err)

		return
	}
	c.JSON(http.StatusOK, result)

	timestamp := time.Now().Format(time.RFC3339)
	ctx := context.Background()
	requestData := map[string]any{
		"dest":     dest,
		"result":   result,
		"datetime": timestamp,
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		logger.Printf("[ERROR] JSON marshal failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"[ERROR]": "Internal server error"})
		return
	}
	redisKey := "request:" + timestamp
	err = rdb.Set(ctx, redisKey, jsonData, 0).Err() // 0 means no expiration
	if err != nil {
		logger.Printf("[ERROR] Redis SET failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"[ERROR]": "Failed to save to Redis"})
		return
	}

	logger.Printf("[INFO] %v", result)
}

func Api() {
	gin.SetMode(gin.ReleaseMode)

	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // default Redis port,
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	logFile, err := os.OpenFile("/var/log/myapp.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	logger = log.New(logFile, "", log.LstdFlags|log.Lshortfile)

	gin.DefaultWriter = logFile
	gin.DefaultErrorWriter = logFile

	router := gin.Default()
	router.GET("/traceroute", TracerouteHandler)
	router.Run("0.0.0.0:8080")
}
