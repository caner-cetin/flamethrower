package main

import (
	"flamethrower/src/db"
	"flamethrower/src/views"
	"fmt"
	"log"
	"os"

	"github.com/rivo/tview"
	"github.com/xuri/excelize/v2"
)

func init() {
	db.InitDB("dnd35.db")

}
func _() {
	f, err := excelize.OpenFile("sheet.xlsx")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			log.Fatal(err)
			return
		}
	}()
	app := tview.NewApplication().EnableMouse(true)
	application := views.Application{}
	application.App = app
	application.Sheet = f
	app.SetRoot(views.ReturnClassView(app), true)
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

}

func main() {
	f, err := excelize.OpenFile("sheet.xlsx")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			log.Fatal(err)
			return
		}
	}()
	err = f.SetCellStr("Sheet1", "A4", "caner")
	if err != nil {
		log.Fatal(err)
		return
	}
	f.SaveAs("sheet.xlsx")
}
