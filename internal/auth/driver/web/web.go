package web

import (
	"github.com/go-rod/rod"
)

type WebAuthDriver struct {
	driver *rod.Browser
}

func NewWebAuthDriver() *WebAuthDriver {
	driver := rod.New().MustConnect()
	return &WebAuthDriver{
		driver: driver,
	}
}

func (d *WebAuthDriver) Auth(username, password string) {
}
