/*
Titouan Schotté
Main core launch pages
*/
// package main

// import (
// 	"Groupie_Tracker/core"
// 	"image/color"
// 	"slices"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	//Get Artists by api
// 	artistsRef := core.Api_artists()

// 	artists := artistsRef
// 	for _, artist := range artists {
// 		core.PreloaderImages[artist.Id] = core.LoadImageFromURL(artist.Image)
// 		core.PreloaderImagesForPopup[artist.Id] = core.LoadImageFromURL(artist.Image)
// 		for _, concert := range artist.ConcertDates {
// 			if !slices.Contains(core.ConcertsLocations, concert.Location.Locations[0]) {
// 				core.ConcertsLocations = append(core.ConcertsLocations, concert.Location.Locations...)
// 			}
// 		}
// 	}
// 	artists = artistsRef

// 	//Create new app
// 	myApp := app.New()
// 	showHomePage(myApp)
// }

// // HOME PAGE
// func showHomePage(app fyne.App) {
// 	window := app.NewWindow("Groupie Tracker - Accueil")
// 	window.CenterOnScreen()

// 	backgroundImage := core.LoadImageFromURL("https://t3.ftcdn.net/jpg/02/23/60/54/360_F_223605406_nGKtPp42ZRx4ZxvrcVeT3Ek6V5Uw4ETh.jpg")
// 	backgroundImage.FillMode = canvas.ImageFillStretch
// 	title := canvas.NewText("Groupie Tracker", color.White)
// 	title.Alignment = fyne.TextAlignCenter
// 	title.TextSize = 24
// 	title.TextStyle = fyne.TextStyle{Bold: true}

// 	enterButton := widget.NewButton("Entrer", func() {
// 		showMainPage(app, window)
// 	})

// 	// Stack the background image behind the other widgets
// 	content := container.NewMax(backgroundImage, container.NewCenter(container.NewVBox(title, enterButton)))

// 	window.SetContent(content)
// 	window.Resize(fyne.NewSize(585, 360))
// 	window.SetFixedSize(true)
// 	window.ShowAndRun()
// }

// // MAIN PAGE with artists ... etc
// func showMainPage(app fyne.App, window fyne.Window) {
// 	myWindow := app.NewWindow("Groupie Tracker")
// 	myWindow.CenterOnScreen()
// 	// Configuring the search bar
// 	searchEntry := core.SetupSearchComponents()

// 	// Initial configuration of filters (hidden by default)
// 	core.SetupFilterComponents()

// 	// Button to show/hide filters
// 	toggleFiltersButton := widget.NewButton("Afficher les filtres", func() {
// 		core.ToggleFiltersVisibility()
// 	})

// 	// Configuring the artist grid
// 	core.SetupGrid()

// 	// Main layout
// 	topContainer := container.NewVBox(searchEntry, toggleFiltersButton, core.FilterContainer)
// 	mainContainer := container.NewBorder(topContainer, nil, nil, nil, core.SetupGridContainer())

// 	myWindow.SetContent(mainContainer)
// 	myWindow.Resize(fyne.NewSize(1600, 1000))
// 	myWindow.Show()

// 	window.Close()
// }

package main

import Groupie_Tracker "Groupie_Tracker/app"

func main() {
	Groupie_Tracker.Accueil(Groupie_Tracker.App)
	Groupie_Tracker.App.Run()
}
