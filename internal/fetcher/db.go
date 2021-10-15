package fetcher

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/hi20160616/gears"
	"github.com/hi20160616/ms-kabar/configs"
	"github.com/pkg/errors"
)

var dbfile = filepath.Join(configs.Data.RootPath, configs.Data.DBPath, "articles.json")

func storage(as []*Article) error {
	defer log.Printf("[%s] Storage Done.", configs.Data.MS["kabar"].Title)
	log.Printf("[%s] Storage ...", configs.Data.MS["kabar"].Title)
	data, err := json.Marshal(as)
	if err != nil {
		return errors.WithMessagef(err, "[%s] storage marshal error:",
			configs.Data.MS["kabar"].Title)
	}
	gears.MakeDirAll(filepath.Join(configs.Data.RootPath, configs.Data.DBPath))
	if err := os.WriteFile(dbfile, data, 0755); err != nil {
		return errors.WithMessagef(err, "[%s] storage WriteFile error:",
			configs.Data.MS["kabar"].Title)
	}
	return nil
}

func load() (as []*Article, err error) {
	data, err := os.ReadFile(dbfile)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &as); err != nil {
		return
	}
	return
}
