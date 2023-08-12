package views

import "github.com/rivo/tview"

func returnErrorView(app *tview.Application, errorMessage string) *tview.Flex {
	modal := tview.NewModal().
		SetText(errorMessage).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				app.SetRoot(ReturnClassView(app), true)
			}
		})

	vFlex := tview.NewFlex()

	// vertical centering
	vFlex.AddItem(tview.NewBox(), 0, 1, false)
	vFlex.AddItem(modal, 0, 2, false)
	vFlex.AddItem(tview.NewBox(), 0, 1, false)

	// horizontal centering
	hFlex := tview.NewFlex()
	hFlex.AddItem(tview.NewBox(), 0, 1, false)
	hFlex.AddItem(vFlex, 0, 2, true)
	hFlex.AddItem(tview.NewBox(), 0, 1, false)
	return hFlex
}

func HandleError(err error, app *tview.Application) {
	if err != nil {
		app.SetRoot(returnErrorView(app, err.Error()), true)
	}
}
