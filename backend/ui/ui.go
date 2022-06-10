package ui

import (
	"github.com/ariopri/rg-km/final-project-engineering-66/backend/repository"
	"github.com/rivo/tview"
)

type UI struct {
	usersRepo repository.UserRepository
	app       *tview.Application
	pages     *tview.Pages
}

func NewUI(usersRepo repository.UserRepository, productsRepo repository.ProductRepository, cartItemRepo repository.CartItemRepository, transactionRepo repository.TransactionRepository) UI {
	return UI{
		usersRepo, nil, nil,
	}
}

func (ui *UI) Start() {
	app := tview.NewApplication()
	ui.app = app

	pages := tview.NewPages()
	pages.AddPage("login", ui.LoginPage(), true, true)

	ui.pages = pages
	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}
