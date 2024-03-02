package main

import (
	"context"
	"github.com/ValeryVerkhoturov/chat/config"
	"github.com/ValeryVerkhoturov/chat/handlers"
	"github.com/gin-contrib/graceful"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func port() string {
	port := config.Port
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func createGzipMiddleware() gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression)
}

func newRouter() *graceful.Graceful {
	router, err := graceful.Default()
	if err != nil {
		panic(err)
	}

	router.Use(gin.Recovery())
	router.Use(createGzipMiddleware())

	router.Static("/images/", "./public/images")
	router.StaticFile("/css/output.css", "./public/css/output.css")
	router.StaticFile("/robots.txt", "./public/robots.txt")
	router.LoadHTMLGlob("templates/templates/**/*")

	//router.NoRoute(handlers.NotFound)

	router.GET("/", handlers.Index)
	router.GET("/index.html", handlers.Index)
	router.POST("/process-query", handlers.ProcessQuery)

	v1Router := router.Group("/v1")

	v1Router.GET("/", handlers.Index)

	return router
}

func main() {
	var err error

	// Log init
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	// Graceful termination when shutting down a process init
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Gin init
	log.Info("Starting Server on http://localhost" + port())
	router := newRouter()
	defer router.Close()

	if err = router.RunWithContext(ctx); err != nil && err != context.Canceled {
		log.Fatal("Unable to start:", err)
	}
}
