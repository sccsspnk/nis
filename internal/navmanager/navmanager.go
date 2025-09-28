package navmanager

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Page interface {
	Layout(gtx layout.Context, th *material.Theme) layout.Dimensions
	HandleEvents(layout.Context)
	Title() string
	ID() string
}

type NavigationManager struct {
	currentPageID string
	history       []string
	pages         map[string]Page
}

func NewNavigationManager() *NavigationManager {
	return &NavigationManager{
		pages: make(map[string]Page),
	}
}

func (nm *NavigationManager) RegisterPage(page Page) {
	nm.pages[page.ID()] = page
}

func (nm *NavigationManager) NavigateTo(pageID string) {
	if _, exists := nm.pages[pageID]; !exists {
		return
	}

	if nm.currentPageID != "" {
		nm.history = append(nm.history, nm.currentPageID)
	}

	nm.currentPageID = pageID
}

func (nm *NavigationManager) NavigateBack() {
	if len(nm.history) == 0 {
		return
	}

	previousPage := nm.history[len(nm.history)-1]
	nm.history = nm.history[:len(nm.history)-1]
	nm.currentPageID = previousPage
}

func (nm *NavigationManager) CurrentPage() Page {
	return nm.pages[nm.currentPageID]
}

func (nm *NavigationManager) CanGoBack() bool {
	return len(nm.history) > 0
}

func (nm *NavigationManager) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	if currentPage := nm.CurrentPage(); currentPage != nil {
		return currentPage.Layout(gtx, th)
	}
	return layout.Dimensions{}
}

func (nm *NavigationManager) HandleEvents(gtx layout.Context) {
	if currentPage := nm.CurrentPage(); currentPage != nil {
		currentPage.HandleEvents(gtx)
	}
}
