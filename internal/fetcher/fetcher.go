package fetcher

import (
	"context"
	"errors"
	"log"

	"github.com/hi20160616/ms-kabar/configs"
)

// Fetch fetch and storage all stuffs to `db/articles.json`
func Fetch() error {
	defer log.Printf("[%s] Fetch Done.", configs.Data.MS.Title)
	log.Printf("[%s] Fetching ...", configs.Data.MS.Title)
	as, err := fetch(context.Background())
	if err != nil {
		return err
	}
	return storage(as)
}

// fetch fetch all articles by url set in config.json
func fetch(ctx context.Context) (as []*Article, err error) {
	links, err := fetchLinks()
	if err != nil {
		return
	}
	for _, link := range links {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			a := NewArticle()
			a, err = a.fetchArticle(link)
			if err != nil {
				if !errors.Is(err, ErrTimeOverDays) {
					log.Printf("[%s] fetch error: %v, link: %s",
						configs.Data.MS.Title, err, link)
				}
				err = nil
				continue
			}
			// ignore redundant articles
			exist := false
			for _, _a := range as {
				if a.Title == _a.Title {
					exist = true
				}
			}
			if !exist {
				as = append(as, a)
			}
		}
	}
	return
}
