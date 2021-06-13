package main

import (
	"context"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	stop := make(chan os.Signal)
	g.Go(serveAPP())
}

func serveAPP()

func serve(addr string, handler http.Handler, stop <-chan os.Signal) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}
