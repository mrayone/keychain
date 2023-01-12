package webserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	v1 "github.com/mrayone/golang-learn/api/products/v1"
	config "github.com/mrayone/golang-learn/configs"
	"golang.org/x/sync/errgroup"
)

func Serve() {
	ctx, cancel := context.WithCancel(context.Background())

	addr := fmt.Sprint(":", config.GetString("HTTP_PORT"))
	srv := &http.Server{
		ReadTimeout:  1000000,
		WriteTimeout: 1000000,
		Addr:         addr,
		Handler:      Router(),
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

		<-c
		cancel()
	}()

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		fmt.Printf("the server is up %s", addr)
		return srv.ListenAndServe()
	})

	g.Go(func() error {
		<-gCtx.Done()
		return srv.Close()
	})

	if err := g.Wait(); err != nil {
		log.Fatalf("shutdown complete, exit reason: %s", err)
	}
}

func Router() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Add("content-type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods(http.MethodGet)

	postProductAPI := v1.NewPostProductHandler()

	router.HandleFunc(postProductAPI.Route, postProductAPI.Handler).Methods(http.MethodOptions, http.MethodPost)

	return handlers.CompressHandler(router)
}
