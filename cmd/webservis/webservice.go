package webservice

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/BintangDwitama/quest-board-be/cmd/webservis/router"
	"github.com/BintangDwitama/quest-board-be/internal/domain/usecase"
)

var server *http.Server

func Start() (stopFn func()) {
	usecase := usecase.NewUsecase()
	router := router.Init(usecase)

	server = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Println("running task-service in web service mode...")
		log.Println("starting web, listening on", server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalln("failed starting web on", server.Addr, err)
		}
	}()

	return func() {
		GracefulStop()
	}
}

func GracefulStop() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("shuting down webservice on", server.Addr)
	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatalln("failed shutdown server", err)
	}
	log.Println("web gracefully stopped")
	return
}
