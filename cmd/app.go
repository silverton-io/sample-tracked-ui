package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var VERSION string

var STATIC_FILE = "./sample.html"

type App struct {
	engine *gin.Engine
}

func (a *App) initializeEngine() {
	a.engine = gin.New()
	a.engine.SetTrustedProxies(nil)
	a.engine.RedirectTrailingSlash = false
}

func (a *App) initializeStaticRoutes() {
	a.engine.StaticFile("/", STATIC_FILE)
	a.engine.StaticFile("/test", STATIC_FILE)
}

func (a *App) Run() {
	log.Info().Interface("version", VERSION).Msgf("starting server")
	a.initializeEngine()
	a.initializeStaticRoutes()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: a.engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Info().Msgf("server shut down")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Stack().Msg("server forced to shut down")
	}
}

func main() {
	app := App{}
	app.Run()
}
