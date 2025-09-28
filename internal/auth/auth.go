package auth

import (
	"fmt"
	"image/color"
	"log"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/sccsspnk/nis/internal/auth/driver/web"
	"github.com/sccsspnk/nis/internal/navmanager"
)

type LoginPage struct {
	theme     *material.Theme
	navigator *navmanager.NavigationManager

	username widget.Editor
	password widget.Editor
	submit   widget.Clickable
	errorMsg string
}

func NewLoginPage(navigator *navmanager.NavigationManager) *LoginPage {
	theme := material.NewTheme()
	username := widget.Editor{}
	password := widget.Editor{}
	submit := widget.Clickable{}

	username.SingleLine = true
	username.Submit = true
	password.SingleLine = true
	password.Submit = true
	password.Mask = '*'

	return &LoginPage{
		theme:     theme,
		navigator: navigator,
		username:  username,
		password:  password,
		submit:    submit,
		errorMsg:  "",
	}
}

func (lp *LoginPage) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.Middle,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				title := material.H3(th, "Авторизация")
				title.Alignment = text.Middle
				return title.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Vertical,
					Alignment: layout.Middle,
					Spacing:   layout.SpaceEnd,
				}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						ed := material.Editor(th, &lp.username, "Имя пользователя")
						ed.TextSize = unit.Sp(16)
						return ed.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: unit.Dp(15)}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						ed := material.Editor(th, &lp.password, "Пароль")
						ed.TextSize = unit.Sp(16)
						return ed.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &lp.submit, "Войти")
						btn.Inset = layout.Inset{
							Top:    unit.Dp(12),
							Bottom: unit.Dp(12),
							Left:   unit.Dp(24),
							Right:  unit.Dp(24),
						}
						btn.CornerRadius = unit.Dp(8)
						btn.TextSize = unit.Sp(16)
						return btn.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: unit.Dp(15)}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if lp.errorMsg != "" {
							errorText := material.Body1(th, lp.errorMsg)
							errorText.Color = color.NRGBA{R: 200, G: 0, B: 0, A: 255}
							errorText.Alignment = text.Middle
							return errorText.Layout(gtx)
						}
						return layout.Dimensions{}
					}),
				)
			}),
		)
	})
}

func (lp *LoginPage) HandleEvents(gtx layout.Context) {
	if lp.submit.Clicked(gtx) {
		if err := lp.validateForm(); err != nil {
			lp.errorMsg = err.Error()
		} else {
			lp.errorMsg = ""
			fmt.Printf("Успешный вход: %s\n", lp.username.Text())
			webDriver := web.NewWebAuthDriver()
			webDriver.Auth("", "")
			lp.navigator.NavigateTo("home")
		}
	}

	for {
		ev, ok := lp.username.Update(gtx)
		if !ok {
			break
		}
		if _, ok := ev.(widget.SubmitEvent); ok {
			log.Println("Change to password input")
		}
	}

	for {
		ev, ok := lp.password.Update(gtx)
		if !ok {
			break
		}
		if _, ok := ev.(widget.SubmitEvent); ok {
			if err := lp.validateForm(); err != nil {
				lp.errorMsg = err.Error()
			} else {
				lp.errorMsg = ""
				fmt.Printf("Успешный вход: %s\n", lp.username.Text())
				lp.navigator.NavigateTo("home")
			}
		}
	}
}

func (lp *LoginPage) Title() string {
	return "Авторизация"
}

func (lp *LoginPage) ID() string {
	return "login"
}

func (lp *LoginPage) validateForm() error {
	if lp.username.Text() == "" {
		return fmt.Errorf("введите имя пользователя")
	}
	if lp.password.Text() == "" {
		return fmt.Errorf("введите пароль")
	}
	return nil
}
