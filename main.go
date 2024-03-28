package main

import (
<<<<<<< HEAD
	"Groupie_Tracker/core"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

//func main() {
//	artists := []core.Artist{
//		{
//			Id:           1,
//			Image:        "image1.jpg",
//			Nom:          "Artiste 1",
//			Members:      []core.Member{{Surname: "Doe", Name: "John"}},
//			CreationDate: "01/01/2020",
//			FirstAlbum:   core.Date{Day: 1, Month: 1, Year: 2021},
//			Concerts:     []core.Concert{{Date: core.Date{Day: 15, Month: 2, Year: 2022}, Location: "Location 1"}},
//			Relations:    "Relations 1",
//		},
//		{
//			Id:           2,
//			Image:        "image2.jpg",
//			Nom:          "Artiste 2",
//			Members:      []core.Member{{Surname: "Smith", Name: "Jane"}, {Surname: "Smith", Name: "Jane"}},
//			CreationDate: "02/02/2019",
//			FirstAlbum:   core.Date{Day: 2, Month: 2, Year: 2020},
//			Concerts:     []core.Concert{{Date: core.Date{Day: 31, Month: 3, Year: 2024}, Location: "Location 2"}},
//			Relations:    "Relations 2",
//		},
//	}
//	print(core.SearchInAllStruct("3", artists)[0].Id)
//
//	// Exemple d'utilisation des filtres
//	filteredArtists := core.FilterByCreationDate(artists, 2020)
//	fmt.Println("Filtered by creation date:")
//	for _, artist := range filteredArtists {
//		fmt.Println(artist.Nom)
//	}
//
//	filteredArtists = core.FilterByFirstAlbumDate(artists, 2021)
//	fmt.Println("\nFiltered by first album date:")
//	for _, artist := range filteredArtists {
//		fmt.Println(artist.Nom)
//	}
//
//	filteredArtists = core.FilterByNumberOfMembers(artists, 1)
//	fmt.Println("\nFiltered by number of members:")
//	for _, artist := range filteredArtists {
//		fmt.Println(artist.Nom)
//	}
//
//	filteredArtists = core.FilterByConcertLocation(artists, "Location 2")
//	fmt.Println("\nFiltered by concert location:")
//	for _, artist := range filteredArtists {
//		fmt.Println(artist.Nom)
//	}
//}

func main() {
	// Créer une nouvelle application
	app := app.New()

	// Créer une nouvelle fenêtre
	window := app.NewWindow("My Fyne App")

	// Créer une image à partir de la ressource générée (on ne peut par image que par ressource)
	source, _ := core.GenerateMapImage("Paris")
	image := canvas.NewImageFromResource(source)
	image.FillMode = canvas.ImageFillContain
	window.SetContent(image)

	// Afficher et exécuter la fenêtre
	window.ShowAndRun()

=======
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
>>>>>>> 3b8c4f2bc45c2184e4456e5249116f7fbd68ea0b
}
