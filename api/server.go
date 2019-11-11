package api

import (
	"context"
	"github.com/gin-gonic/gin"
	//"github.com/opentracing/opentracing-go"
	"github.com/semihalev/gin-stats"
	"highloadcup/travels/api/httpHandlers"
	"highloadcup/travels/api/middlewares"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"highloadcup/travels/config"
)

func Run() {
	gin.DisableConsoleColor()
	if !config.Config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(stats.RequestStats())
	r.Use(middlewares.RequestMiddleware())
	r.Use(middlewares.ResponseMiddleware())
	r.Use(middlewares.RecoveryMiddleware())
	//r.Use(middlewares.TracerMiddleware(tracer))

	/*jsonEP := r.Group("/market/v1/monitoring")
	{
		//jsonEP.GET("/deliveryRegistration", httpHandlers.DeliveryRegistration)
		//jsonEP.POST("/review/list", httpHandlers.ReviewListRequestHandler)
	}*/

	otherEP := r.Group("/")
	{
		//otherEP.GET("/stats", func(c *gin.Context) {
		//	c.JSON(http.StatusOK, stats.Report())
		//})
		otherEP.GET("/health", httpHandlers.Health)
		otherEP.POST("/health", httpHandlers.Health)

		otherEP.GET("/user/:id", httpHandlers.User)
		otherEP.POST("/user/:id", httpHandlers.CreateUser)
	}

	srv := &http.Server{
		Addr:         config.Config.HTTPServer.Host + ":" + strconv.Itoa(int(config.Config.HTTPServer.InternalPort)),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		//TLSConfig: tlsConfig,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}
