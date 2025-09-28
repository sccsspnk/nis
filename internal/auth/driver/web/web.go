package web

import (
	"github.com/go-rod/rod"
)

type WebAuthDriver struct {
	driver *rod.Browser
	page   *rod.Page
}

func NewWebAuthDriver() *WebAuthDriver {
	return &WebAuthDriver{
		driver: rod.New(),
	}
}

func (d *WebAuthDriver) Auth(username, password string) {

}
