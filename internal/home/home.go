package home

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/sccsspnk/nis/internal/navmanager"
)

type HomePage struct {
	theme     *material.Theme
	navigator *navmanager.NavigationManager
	logoutBtn widget.Clickable
}

func NewHomePage(navigator *navmanager.NavigationManager) *HomePage {
	return &HomePage{
		theme:     material.NewTheme(),
		navigator: navigator,
	}
}

func (hp *HomePage) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.Middle,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				title := material.H3(th, "Главная страница")
				return title.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Body1(th, "Добро пожаловать в утилиту НИС!").Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: unit.Dp(30)}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, &hp.logoutBtn, "Выйти")
				return btn.Layout(gtx)
			}),
		)
	})
}

func (hp *HomePage) HandleEvents(gtx layout.Context) {
	if hp.logoutBtn.Clicked(gtx) {
		hp.navigator.NavigateTo("login")
	}
}

func (hp *HomePage) Title() string {
	return "Главная"
}

func (hp *HomePage) ID() string {
	return "home"
}
