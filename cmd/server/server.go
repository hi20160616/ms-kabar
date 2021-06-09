package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hi20160616/ms-kabar/configs"
	"github.com/hi20160616/ms-kabar/internal/job"
	"github.com/hi20160616/ms-kabar/internal/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// MS
	g.Go(func() error {
		log.Printf("[%s] MS start at: %s", configs.Data.MS.Title, configs.Data.MS.Addr)
		return server.Start(ctx)
	})
	g.Go(func() error {
		<-ctx.Done() // wait for stop signal
		return server.Stop(ctx)
	})

	// Job
	g.Go(func() error {
		log.Printf("[%s] Job is working", configs.Data.MS.Title)
		return job.Crawl(ctx)
	})
	g.Go(func() error {
		<-ctx.Done() // wait for stop signal
		return job.Stop(ctx)
	})

	// Elegant stop
	c := make(chan os.Signal, 1)
	sigs := []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	signal.Notify(c, sigs...)
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-c:
			log.Printf("signal caught: %s ready to quit...", sig.String())
			cancel()
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		if !errors.Is(err, context.Canceled) {
			log.Printf("not canceled by context: %s", err)
		} else {
			log.Println(err)
		}
	}
}
