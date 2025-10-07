package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    nats "github.com/nats-io/nats.go"
)

type TaskRequest struct {
    Input string `json:"input"`
}

func main() {
    natsURL := os.Getenv("NATS_URL")
    if natsURL == "" {
        natsURL = "nats://localhost:4222"
    }

    nc, err := nats.Connect(natsURL, nats.Name("api-go"))
    if err != nil {
        log.Fatalf("connect nats: %v", err)
    }
    defer nc.Drain()

    // Subscribe to results to log outcomes
    _, err = nc.Subscribe("results", func(m *nats.Msg) {
        log.Printf("Result: %s", string(m.Data))
    })
    if err != nil {
        log.Fatalf("subscribe results: %v", err)
    }

    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    r.POST("/tasks", func(c *gin.Context) {
        var req TaskRequest
        if err := c.BindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
            return
        }
        payload, _ := json.Marshal(req)
        if err := nc.Publish("tasks", payload); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "publish failed"})
            return
        }
        _ = nc.Flush()
        c.JSON(http.StatusAccepted, gin.H{"queued": true, "at": time.Now()})
    })

    if err := r.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}

