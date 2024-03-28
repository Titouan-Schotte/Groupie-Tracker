package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

var preloaderImages = map[int]*canvas.Image{}
var preloaderImagesForPopup = map[int]*canvas.Image{}

func main() {
	artists = artistsRef
	for _, artist := range artists {
		preloaderImages[artist.Id] = loadImageFromURL(artist.Image)
		preloaderImagesForPopup[artist.Id] = loadImageFromURL(artist.Image)
	}
	myApp := app.New()
	showHomePage(myApp)
}

func showHomePage(app fyne.App) {
	window := app.NewWindow("Groupie Tracker - Accueil")
	window.CenterOnScreen()
	title := canvas.NewText("Groupie Tracker", color.White)
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24
	title.TextStyle = fyne.TextStyle{Bold: true}

	enterButton := widget.NewButton("Entrer", func() {
		showMainPage(app, window) // Ouvrir la page principale
	})

	vbox := container.NewVBox(
		title,
		enterButton,
	)

	window.SetContent(container.NewCenter(vbox))
	window.Resize(fyne.NewSize(400, 200))
	window.ShowAndRun()
}

func showMainPage(app fyne.App, window fyne.Window) {
	myWindow = app.NewWindow("Groupie Tracker")
	myWindow.CenterOnScreen()
	// Configuration de la barre de recherche
	searchEntry := setupSearchComponents()

	// Configuration initiale des filtres (masqués par défaut)
	setupFilterComponents()

	// Bouton pour afficher/masquer les filtres
	toggleFiltersButton := widget.NewButton("Afficher les filtres", func() {
		toggleFiltersVisibility()
	})

	// Configuration de la grille d'artistes
	setupGrid()

	// Agencement principal
	topContainer := container.NewVBox(searchEntry, toggleFiltersButton, filterContainer)
	mainContainer := container.NewBorder(topContainer, nil, nil, nil, setupGridContainer())

	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
	window.Close()
}
