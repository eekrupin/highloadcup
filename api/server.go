package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/semihalev/gin-stats"
	"gopkg.in/go-playground/validator.v8"

	"github.com/eekrupin/hlc-travels/api/httpHandlers"
	"github.com/eekrupin/hlc-travels/api/middlewares"
	"github.com/eekrupin/hlc-travels/config"
	//"github.com/opentracing/opentracing-go"
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("emailCheck", emailCheck)
	}

	otherEP := r.Group("/")
	{
		//otherEP.GET("/stats", func(c *gin.Context) {
		//	c.JSON(http.StatusOK, stats.Report())
		//})
		otherEP.GET("/health", httpHandlers.Health)
		otherEP.POST("/health", httpHandlers.Health)

		otherEP.GET("/user/:id", httpHandlers.User)

		otherEP.POST("/user/:id", httpHandlers.PostUser)
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

func emailCheck(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if email, ok := field.Interface().(string); ok {
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		return re.MatchString(email)
	}
	return false
}
