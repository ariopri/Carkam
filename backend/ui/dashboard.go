package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

func (ui *UI) DashboardPage(cash int) {
	username, err := ui.usersRepo.FindLoggedinUser()
	if err != nil {
		panic(err)
	}

	createText := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	fmt.Println(createText)
}
