/*
Titouan SchottÃ©
Main core launch pages
*/
package main

import (
	"Groupie_Tracker/core"
	"image/color"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var preloaderImages = map[int]*canvas.Image{}
var preloaderImagesForPopup = map[int]*canvas.Image{}
var concertsLocations = []string{}

func main() {
	//Load the favorite datas
	core.LoadJson()

	//Get Artists by api
	artistsRef = core.Api_artists()

	artists = artistsRef
	for _, artist := range artists {
		preloaderImages[artist.Id] = loadImageFromURL(artist.Image)
		preloaderImagesForPopup[artist.Id] = loadImageFromURL(artist.Image)
		for _, concert := range artist.ConcertDates {
			if !slices.Contains(concertsLocations, concert.Location) {
				concertsLocations = append(concertsLocations, concert.Location)
			}
		}
	}
	artists = artistsRef

	//Create new app
	myApp := app.New()
	showHomePage(myApp)
}

// HOME PAGE
func showHomePage(app fyne.App) {
	window := app.NewWindow("Groupie Tracker - Accueil")
	window.CenterOnScreen()

	backgroundImage := loadImageFromURL("https://t3.ftcdn.net/jpg/02/23/60/54/360_F_223605406_nGKtPp42ZRx4ZxvrcVeT3Ek6V5Uw4ETh.jpg")
	backgroundImage.FillMode = canvas.ImageFillStretch
	title := canvas.NewText("Groupie Tracker", color.White)
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24
	title.TextStyle = fyne.TextStyle{Bold: true}

	enterButton := widget.NewButton("Entrer", func() {
		showMainPage(app, window)
	})

	// Stack the background image behind the other widgets
	content := container.NewMax(backgroundImage, container.NewCenter(container.NewVBox(title, enterButton)))

	window.SetContent(content)
	window.Resize(fyne.NewSize(585, 360))
	window.SetFixedSize(true)
	window.ShowAndRun()
}

// MAIN PAGE with artists ... etc
func showMainPage(app fyne.App, window fyne.Window) {
	myWindow = app.NewWindow("Groupie Tracker")
	myWindow.CenterOnScreen()
	// Configuring the search bar
	searchEntry := setupSearchComponents()

	// Initial configuration of filters (hidden by default)
	setupFilterComponents()

	// Button to show/hide filters
	toggleFiltersButton := widget.NewButton("Afficher les filtres", func() {
		toggleFiltersVisibility()
	})

	// ===================JAI AJOUTER CE BOUTON==================

	//Button to go to the favorite page
	favoriteButton := widget.NewButton("Afficher les favoris", func() {
		FavorisPage(app, window)
	})
	// ===========================================================

	// Configuring the artist grid
	setupGrid()

	// Main layout
	topContainer := container.NewVBox(favoriteButton, searchEntry, toggleFiltersButton, filterContainer)
	mainContainer := container.NewBorder(topContainer, nil, nil, nil, setupGridContainer())

	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(1600, 1000))
	myWindow.Show()

	window.Close()
}

// ===================JAI AJOUTER CETTE PAGE==================

func FavorisPage(app fyne.App, window fyne.Window) {
	myWindow := app.NewWindow("Groupie Tracker")
	myWindow.CenterOnScreen()

	homeButton := widget.NewButton("Accueil", func() {
		//showMainPage(app, window)
		myWindow.Hide()
	})

	grid := container.NewGridWithColumns(5)

	favArtist := core.GetFavorite()

	for _, artist := range favArtist {
		currentArtist := artist // Create a local copy of the variable for capturing
		image := preloaderImages[artist.Id]
		if image.Size().Width > 500 || image.Size().Height > 500 {
			image.Resize(fyne.NewSize(500, 500))
		}

		imageContainer := container.New(layout.NewMaxLayout(), image)

		button := widget.NewButton(currentArtist.Nom, func() {
			showArtistDetails(currentArtist) // Utiliser la copie locale dans la fermeture
		})

		imageWithButton := container.New(layout.NewVBoxLayout(), imageContainer, button)

		grid.Add(imageWithButton)
	}

	scrollContainer := container.NewVScroll(grid)

	mainContainer := container.NewBorder(homeButton, nil, nil, nil, scrollContainer)

	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(1600, 1000))
	myWindow.Show()

	window.Close()
}

// ===========================================================
