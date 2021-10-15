package configs

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type configuration struct {
	RootPath string
	DBPath   string                  `json:"dbpath"`
	Debug    bool                    `json:"debug"`
	Verbose  bool                    `json:"verbose"`
	Gist     string                  `json:"gist"`
	MS       map[string]MicroService `json:"microservice"`
}

type MicroService struct {
	Title     string   `json:"title"`
	Domain    string   `json:"domain"`
	URL       []string `json:"url"`
	Addr      string   `json:"addr"`
	Timeout   string   `json:"timeout"`
	Heartbeat string   `json:"heartbeat"`
}

var Data = &configuration{}

func setRootPath() error {
	var root string
	var err error
	if strings.Contains(os.Args[0], ".test") {
		root = "../../" // for test dbmanager
		// root = "../" // for test configs
	} else {
		root, err = os.Getwd()
		if err != nil {
			log.Printf("config Getwd: %#v", err)
		}
	}
	Data.RootPath = root
	return nil
}

func get() error {
	f, err := os.ReadFile(filepath.Join(Data.RootPath, "configs/configs.json"))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(f, Data); err != nil {
		return err
	}

	// test gist
	// Data.Gist = "https://gist.github.com/hi20160616/d932caa9c0c905c07ee4f773fea7c850/raw/configs.json"
	resp, err := http.Get(Data.Gist)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &Data); err != nil {
		return err
	}

	return nil
}
func init() {
	if err := setRootPath(); err != nil {
		log.Printf("config init error: %v", err)
	}
	if err := get(); err != nil {
		log.Printf("config get() error: %v", err)
	}
}

// Reset is for test to reset RootPath and invoke get()
func Reset(pwd string) error {
	Data.RootPath = pwd
	return get()
}
