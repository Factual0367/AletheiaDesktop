package book

import (
	"AletheiaDesktop/search"
	"AletheiaDesktop/util/shared"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

// cache images after first run
var coverImageCache = make(map[string]*fyne.Container)

func CreateBookLibraryContainer(book search.Book) *fyne.Container {
	bookDetailsString := fmt.Sprintf(
		"Title: %s\nAuthor: %s\nFiletype: %s\nFilesize: %s\nLanguage: %s\nPages: %s\nPublisher: %s",
		book.Title, book.Author, book.Extension, book.Size, book.Language, book.Pages, book.Publisher,
	)

	bookDetailsLabel := widget.NewLabelWithStyle(bookDetailsString, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	bookDetailsLabel.Wrapping = fyne.TextWrapWord

	openButton := widget.NewButtonWithIcon("", theme.FileIcon(), func() {
		go func() {
			err := shared.OpenBookWithDefaultApp(book.Filepath)
			if err != nil {
				log.Fatalln("Could not open book with default application.")
			}
		}()
	})

	convertButton := widget.NewButtonWithIcon("", theme.ContentRedoIcon(), func() {})
	deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {})

	buttonContainer := container.NewHBox(openButton, convertButton, deleteButton, layout.NewSpacer())

	border := canvas.NewRectangle(&color.NRGBA{R: 97, G: 97, B: 97, A: 50})
	border.StrokeColor = color.NRGBA{R: 97, G: 97, B: 97, A: 50}
	border.StrokeWidth = 2
	border.CornerRadius = 10

	bookCover, ok := coverImageCache[book.ID]
	if !ok {
		bookCover = createBookDetailsTopView(book)
		coverImageCache[book.ID] = bookCover
	}

	borderedContainer := container.NewStack(border, container.NewVBox(bookDetailsLabel, buttonContainer))
	borderedContainerWithCover := container.NewHSplit(bookCover, borderedContainer)
	borderedContainerWithCover.SetOffset(0.10)

	return container.NewVBox(borderedContainerWithCover)
}
