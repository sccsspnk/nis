package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"github.com/sccsspnk/nis/internal/auth"
	"github.com/sccsspnk/nis/internal/home"
	"github.com/sccsspnk/nis/internal/navmanager"
)

type application struct {
	window *app.Window
	theme  *material.Theme

	width, height int
	title         string

	navigator *navmanager.NavigationManager
}

func NewApp(width, height int, title string) *application {
	return &application{
		window: new(app.Window),
		theme:  material.NewTheme(),
		width:  width,
		height: height,
		title:  title,
	}
}

func (a *application) configure() {
	shaper := text.NewShaper(text.WithCollection(gofont.Collection()))
	a.theme.Shaper = shaper
	a.window.Option(
		app.Size(unit.Dp(a.width), unit.Dp(a.height)),
		app.Title(a.title),
	)
}

func (a *application) setupNavigation() {
	a.navigator = navmanager.NewNavigationManager()
	loginPage := auth.NewLoginPage(a.navigator)
	homePage := home.NewHomePage(a.navigator)
	a.navigator.RegisterPage(loginPage)
	a.navigator.RegisterPage(homePage)
	a.navigator.NavigateTo("login")
}

func (a *application) Run() error {
	a.configure()
	a.setupNavigation()

	var ops op.Ops

	for {
		switch e := a.window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			a.navigator.HandleEvents(gtx)
			a.navigator.Layout(gtx, a.theme)
			e.Frame(gtx.Ops)
		}
	}
}
