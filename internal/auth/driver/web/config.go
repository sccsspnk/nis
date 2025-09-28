package web

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/qjebbs/go-jsons"
)

type WebAuthDriverConfig struct {
	config interface{}
}

func NewWebAuthDriverConfig() *WebAuthDriverConfig {
	config := &WebAuthDriverConfig{
		config: map[string]string{},
	}
	config.Load()
	return config
}

func (w *WebAuthDriverConfig) Load() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fileData, err := os.ReadFile(filepath.Join(path, "config.dat"))
	if err != nil {
		file, err := os.Create(filepath.Join(path, "config.dat"))
		if err != nil {
			log.Fatal(err)
		}
		file.Write([]byte("{}"))
		file.Close()
		w.Load()
		return
	}
	err = json.NewDecoder(bytes.NewBuffer(fileData)).Decode(&w.config)
	if err != nil {
		err := os.WriteFile(filepath.Join(path, "config.dat"), []byte("{}"), os.FileMode(os.O_WRONLY))
		if err != nil {
			log.Fatal(err)
		}
		w.Load()
	}
}

func (w *WebAuthDriverConfig) GetRaw() string {
	mJson, err := json.Marshal(w.config)
	if err != nil {
		log.Fatal(err)
	}
	return string(mJson)
}

func (w *WebAuthDriverConfig) Save(data ...[]byte) {
	mJson, err := json.Marshal(w.config)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := jsons.Merge(mJson, data)
	if err != nil {
		log.Fatal(err)
	}
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(filepath.Join(path, "config.dat"), bytes, os.FileMode(os.O_WRONLY))
	w.Load()
}
