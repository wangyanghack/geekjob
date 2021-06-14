package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	s := http.Server{
		Addr:    "addr",
		Handler: http.DefaultServeMux,
	}
	ctx0, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx0)
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	g.Go(func() error {
		<-ctx.Done()
		return s.Shutdown(context.Background())
	})
	g.Go(func() error {
		return s.ListenAndServe()
	})
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return errors.New("receive cancle context error")
			case <-stop:
				cancel()
			}
		}
	})
	if err := g.Wait(); err != nil {
		fmt.Printf("err is %+v\n", err)
	}
}
