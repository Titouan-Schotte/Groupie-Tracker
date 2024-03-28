package main

import (
	"Groupie_Tracker/core"
	"fmt"
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
	artistsRef = core.Api_artists()

	artists = artistsRef
	for i, artist := range artists {
		preloaderImages[artist.Id] = loadImageFromURL(artist.Image)
		preloaderImagesForPopup[artist.Id] = loadImageFromURL(artist.Image)
		artistsRef[i].ConcertDates = append(artistsRef[i].ConcertDates, core.Concert{Date: core.Date{Year: 2012, Month: 9, Day: 22}, Location: "Dublin"})
		artistsRef[i].ConcertDates = append(artistsRef[i].ConcertDates, core.Concert{Date: core.Date{Year: 2012, Month: 6, Day: 22}, Location: "Berlin"})
		fmt.Println(len(artistsRef[i].Relations))

	}
	artists = artistsRef

	myApp := app.New()
	showHomePage(myApp)
}

func showHomePage(app fyne.App) {
	window := app.NewWindow("Groupie Tracker - Accueil")
	window.CenterOnScreen()

	// Charger l'image de fond prétraitée avec un effet de flou
	backgroundImage := loadImageFromURL("https://t3.ftcdn.net/jpg/02/23/60/54/360_F_223605406_nGKtPp42ZRx4ZxvrcVeT3Ek6V5Uw4ETh.jpg")
	backgroundImage.FillMode = canvas.ImageFillStretch // Ajuster pour remplir l'espace

	title := canvas.NewText("Groupie Tracker", color.White)
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24
	title.TextStyle = fyne.TextStyle{Bold: true}

	enterButton := widget.NewButton("Entrer", func() {
		showMainPage(app, window) // Ouvrir la page principale
	})

	// Empiler l'image de fond derrière les autres widgets
	content := container.NewMax(backgroundImage, container.NewCenter(container.NewVBox(title, enterButton)))

	window.SetContent(content)
	window.Resize(fyne.NewSize(585, 360))
	window.SetFixedSize(true)
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
	myWindow.Resize(fyne.NewSize(1600, 1200))
	myWindow.Show()

	window.Close()
}
