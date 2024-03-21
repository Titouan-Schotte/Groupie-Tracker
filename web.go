package main

import (
	"Groupie_Tracker/core"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

var (
	myApp                 fyne.App
	myWindow              fyne.Window
	artists               []core.Artist
	grid                  *fyne.Container
	searchTerm            string
	creationDateFromEntry *widget.Entry
	creationDateToEntry   *widget.Entry
	firstAlbumFromEntry   *widget.Entry
	firstAlbumToEntry     *widget.Entry
	numMembersEntry       *widget.Entry
	concertLocationSelect *widget.Select
	ConcertLocations      = []string{"Location 1", "Location 2"}
	artistsRef            = core.Api_artists()
)

func main() {
	myApp = app.New()
	myWindow = myApp.NewWindow("Groupie Tracker")

	artists = artistsRef
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste")

	searchButton := widget.NewButton("Rechercher", func() {
		searchTerm = searchEntry.Text
		artists = core.SearchInAllStruct(searchTerm, artistsRef)
		updateGrid()
	})

	creationDateFromEntry = widget.NewEntry()
	creationDateFromEntry.SetPlaceHolder("Année de début")
	creationDateToEntry = widget.NewEntry()
	creationDateToEntry.SetPlaceHolder("Année de fin")
	firstAlbumFromEntry = widget.NewEntry()
	firstAlbumFromEntry.SetPlaceHolder("Année de début")
	firstAlbumToEntry = widget.NewEntry()
	firstAlbumToEntry.SetPlaceHolder("Année de fin")
	numMembersEntry = widget.NewEntry()
	numMembersEntry.SetPlaceHolder("Nombre de membres")
	concertLocationSelect = widget.NewSelect(ConcertLocations, nil)

	applyFiltersButton := widget.NewButton("Appliquer les filtres", applyFilters)

	filterFields := container.NewVBox(
		widget.NewLabel("Filtre par date de création :"),
		container.NewHBox(widget.NewLabel("De "), creationDateFromEntry, widget.NewLabel(" à "), creationDateToEntry),
		widget.NewLabel("Filtre par date du premier album :"),
		container.NewHBox(widget.NewLabel("De "), firstAlbumFromEntry, widget.NewLabel(" à "), firstAlbumToEntry),
		widget.NewLabel("Filtre par nombre de membres :"),
		numMembersEntry,
		widget.NewLabel("Filtre par lieu de concert :"),
		concertLocationSelect,
		applyFiltersButton,
	)

	searchBox := container.NewHBox(
		widget.NewLabel("Search: "),
		searchEntry, // Occupera automatiquement tout l'espace horizontal disponible
		searchButton,
	)

	topContainer := container.NewVBox(
		searchBox,
		filterFields,
	)

	grid = container.NewGridWithColumns(5)
	updateGrid()
	gridWithScroll := container.NewVScroll(grid)

	mainContainer := container.NewBorder(topContainer, nil, nil, nil, gridWithScroll)

	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}

// Les autres fonctions telles que updateGrid(), showArtistsGrid(), loadImageFromURL(), applyFilters(), et showArtistDetails() restent inchangées.

func updateGrid() {
	filteredArtists := artists // Appliquez votre logique de filtrage ici
	if searchTerm != "" {
		// filteredArtists = core.SearchInAllStruct(searchTerm, artists) // Exemple
	}
	showArtistsGrid(filteredArtists)
}

func showArtistsGrid(artists []core.Artist) {
	grid.Objects = nil
	for _, artist := range artists {
		image := loadImageFromURL(artist.Image)
		// Limit the maximum size of the image to 500x500
		if image.Size().Width > 500 || image.Size().Height > 500 {
			image.Resize(fyne.NewSize(500, 500))
		}
		// Create a container for the image
		imageContainer := container.New(layout.NewMaxLayout(), image)
		fmt.Println(artist.Image)

		// Ajouter le bouton avec le nom de l'artiste en dessous de l'image
		button := widget.NewButton(artist.Nom, func() {
			fmt.Println(artist.Image)
			showArtistDetails(artist)
		})

		// Créer un conteneur pour organiser l'image et le bouton
		imageWithButton := container.New(layout.NewVBoxLayout(), imageContainer, button)

		// Ajouter le conteneur à la grille
		grid.Add(imageWithButton)
	}

	grid.Refresh()
}

func loadImageFromURL(urlStr string) *canvas.Image {
	resource, _ := fyne.LoadResourceFromURLString(urlStr)

	image := canvas.NewImageFromResource(resource)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(200, 200)) // Adjust the size as needed
	return image
}

// Appliquer les filtres et mettre à jour la grille
func applyFilters() {
	var filteredArtists []core.Artist

	// Appliquer le filtre par date de création s'il est renseigné
	creationDateFromStr := creationDateFromEntry.Text
	if creationDateFromStr != "" {
		// Convertir l'année en entier
		creationDateFrom, _ := strconv.Atoi(creationDateFromStr)
		// Filtrer les artistes en fonction de l'année de création
		filteredArtists = core.FilterByCreationDate(artistsRef, creationDateFrom)
	} else {
		// Si le filtre n'est pas renseigné, utiliser tous les artistes non filtrés jusqu'à présent
		filteredArtists = artistsRef
	}

	// Appliquer le filtre par date du premier album s'il est renseigné
	firstAlbumFromStr := firstAlbumFromEntry.Text
	if firstAlbumFromStr != "" {
		// Convertir l'année en entier
		firstAlbumFrom, _ := strconv.Atoi(firstAlbumFromStr)
		// Filtrer les artistes en fonction de l'année du premier album
		filteredArtists = core.FilterByFirstAlbumDate(filteredArtists, firstAlbumFrom)
	}

	// Appliquer le filtre par nombre de membres s'il est renseigné
	numMembersStr := numMembersEntry.Text
	if numMembersStr != "" {
		numMembers, _ := strconv.Atoi(numMembersStr)
		// Filtrer les artistes en fonction du nombre de membres
		filteredArtists = core.FilterByNumberOfMembers(filteredArtists, numMembers)
	}

	// Appliquer le filtre par lieu de concert s'il est renseigné
	concertLocation := concertLocationSelect.Selected
	if concertLocation != "" {
		// Filtrer les artistes en fonction du lieu de concert sélectionné
		filteredArtists = core.FilterByConcertLocation(filteredArtists, concertLocation)
	}

	// Mettre à jour la grille avec les artistes filtrés
	showArtistsGrid(filteredArtists)
}

// Fonction pour afficher les détails de l'artiste
func showArtistDetails(artist core.Artist) {
	fmt.Println("ID =>", artist.Id)
	fmt.Println("Nom =>", artist.Nom)
	fmt.Println("Image =>", artist.Image)
	fmt.Println("First Album =>", artist.FirstAlbum)
	fmt.Println("Concerts liste =>", artist.Concerts)
	fmt.Println("Creation Album =>", artist.CreationDate)
	fmt.Println("Relations =>", artist.Relations)
}
