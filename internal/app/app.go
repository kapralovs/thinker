package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kapralovs/thinker/internal/auth"
	authhttp "github.com/kapralovs/thinker/internal/auth/controllers/http"
	authRepo "github.com/kapralovs/thinker/internal/auth/repository/localcache"
	authUC "github.com/kapralovs/thinker/internal/auth/usecase"
	"github.com/kapralovs/thinker/internal/note"
	notehttp "github.com/kapralovs/thinker/internal/note/controllers/http"
	noteRepo "github.com/kapralovs/thinker/internal/note/repository/localcache"
	noteUC "github.com/kapralovs/thinker/internal/note/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/natefinch/lumberjack"
)

type app struct {
	httpServer  *http.Server
	noteUseCase note.UseCase
	authUseCase auth.UseCase
}

func NewApp() *app {
	nRepo := noteRepo.NewLocalRepo()
	authRepo := authRepo.NewLocalRepo()

	return &app{
		noteUseCase: noteUC.NewNoteUseCase(nRepo),
		authUseCase: authUC.NewAuthUseCase(authRepo),
	}
}

func (a *app) Run(port string) error {
	// Initialize router instance
	router := echo.New()

	// TODO: Temporary logger. Rewrite it!!!
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/thinker.log",
		MaxSize:    10,
		MaxBackups: 10,
		MaxAge:     10,
		Compress:   true,
	}

	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: lumberJackLogger,
	}))

	// Set router groups
	authGroup := router.Group("/auth/")
	noteGroup := router.Group("/note/", authhttp.NewAuthMiddlewareHandler(a.authUseCase))

	// RegisterEndpoints
	authhttp.RegisterEndpoints(authGroup, a.authUseCase)
	notehttp.RegisterEndpoints(noteGroup, a.noteUseCase)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB (RFC 2616)
	}

	// Launch HTTP server into separate goroutine
	go func() {
		log.Println("Server is starting...")
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	// Make channel of os.Signals with cap=1
	quit := make(chan os.Signal, 1)

	// Handle os.Interrupt, syscall.SIGTERM signals
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Locked until the first element (signal) is passed to the channel
	<-quit

	// Added a context with timeout for passing it to the a.httpServer.Shutdown(ctx)
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	// Shutdown HTTP server with timeout
	return a.httpServer.Shutdown(ctx)
}
