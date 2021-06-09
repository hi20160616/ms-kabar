package job

import (
	"context"
	"log"
	"time"

	"github.com/hi20160616/ms-kabar/configs"
	"github.com/hi20160616/ms-kabar/internal/fetcher"
	"github.com/pkg/errors"
)

func Crawl(ctx context.Context) error {
	f := func() {
		if err := fetcher.Fetch(); err != nil {
			if !errors.Is(err, fetcher.ErrTimeOverDays) {
				log.Printf("%#v", err)
			}
		}
	}
	f() // fetch init while start up
	t, err := time.ParseDuration(configs.Data.MS.Heartbeat)
	if err != nil {
		return err
	}
	for {
		select {
		case <-time.Tick(t):
			f()
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// Stop is nil now
func Stop(ctx context.Context) error {
	log.Println("Job gracefully stopping.")
	// return error can define here, so it will display on frontend
	return ctx.Err()
}
