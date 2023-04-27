package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	router := echo.New()

	authGroup := router.Group("/auth/")
	noteGroup := router.Group("/note/", authhttp.NewAuthMiddlewareHandler(a.authUseCase))

	authhttp.RegisterEndpoints(authGroup, a.authUseCase)
	notehttp.RegisterEndpoints(noteGroup, a.noteUseCase)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Println("Server is starting...")
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}