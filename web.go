package main

import (
	"Groupie_Tracker/core"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	myApp                     fyne.App
	myWindow                  fyne.Window
	artists                   []core.Artist
	grid                      *fyne.Container
	searchTerm                string
	creationDateFromSelect    *widget.Select
	creationDateToSelect      *widget.Select
	firstAlbumFromSelect      *widget.Select
	firstAlbumToSelect        *widget.Select
	numMembersSelect          *widget.Select
	concertLocationSelect     *widget.Select
	is_creationDateFromSelect = false
	is_creationDateToSelect   = false
	is_firstAlbumFromSelect   = false
	is_firstAlbumToSelect     = false
	is_numMembersSelect       = false
	is_concertLocationSelect  = false
	ConcertLocations          = []string{"Location 1", "Location 2"}
	artistsRef                = core.Api_artists()
	filtersVisible            bool // Nouveau : état de visibilité des filtres
	filterContainer           *fyne.Container
)

func home() {
	myApp = app.New()
	myWindow = myApp.NewWindow("Groupie Tracker")
	artists = artistsRef

	// Configuration de la barre de recherche
	searchEntry, searchButton := setupSearchComponents()

	// Configuration initiale des filtres (masqués par défaut)
	setupFilterComponents()

	// Bouton pour afficher/masquer les filtres
	toggleFiltersButton := widget.NewButton("Afficher les filtres", func() {
		toggleFiltersVisibility()
	})

	// Configuration de la grille d'artistes
	setupGrid()

	// Agencement principal
	topContainer := container.NewVBox(searchEntry, searchButton, toggleFiltersButton, filterContainer)
	mainContainer := container.NewBorder(topContainer, nil, nil, nil, setupGridContainer())

	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}

func setupSearchComponents() (*widget.Entry, *widget.Button) {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste")

	searchButton := widget.NewButton("Rechercher", func() {
		searchTerm = searchEntry.Text
		updateGrid()
	})

	return searchEntry, searchButton
}

func setupFilterComponents() {
	// Générer les options d'années pour les menus déroulants
	yearOptions := generateYearOptions(1900, 2024)
	// Créer les sélecteurs d'années avec les options générées
	creationDateFromSelect = widget.NewSelect(yearOptions, func(_ string) {
		is_creationDateFromSelect = true
	})
	creationDateFromSelect.PlaceHolder = "Année de début"
	creationDateToSelect = widget.NewSelect(yearOptions, func(_ string) {
		is_creationDateToSelect = true
	})
	creationDateToSelect.PlaceHolder = "Année de fin"

	firstAlbumFromSelect = widget.NewSelect(yearOptions, func(_ string) {
		is_firstAlbumFromSelect = true
	})
	firstAlbumFromSelect.PlaceHolder = "Année de début"
	firstAlbumToSelect = widget.NewSelect(yearOptions, func(_ string) {
		is_firstAlbumToSelect = true
	})
	firstAlbumToSelect.PlaceHolder = "Année de fin"

	// Générer les options pour le nombre de membres et créer le sélecteur
	memberOptions := generateMemberOptions(1, 20)
	numMembersSelect = widget.NewSelect(memberOptions, func(_ string) {
		is_numMembersSelect = true
	})
	numMembersSelect.PlaceHolder = "Nombre de membres"

	concertLocationSelect = widget.NewSelect(ConcertLocations, func(_ string) {
		is_concertLocationSelect = true
	})

	applyFiltersButton := widget.NewButton("Appliquer les filtres", func() {
		applyFilters() // Assurez-vous d'adapter cette fonction en conséquence
	})

	filterContainer = container.NewVBox(
		widget.NewLabel("Filtre par date de création :"),
		container.NewHBox(widget.NewLabel("De "), creationDateFromSelect, widget.NewLabel(" à "), creationDateToSelect),
		widget.NewLabel("Filtre par date du premier album :"),
		container.NewHBox(widget.NewLabel("De "), firstAlbumFromSelect, widget.NewLabel(" à "), firstAlbumToSelect),
		widget.NewLabel("Filtre par nombre de membres :"),
		numMembersSelect,
		widget.NewLabel("Filtre par lieu de concert :"),
		concertLocationSelect,
		applyFiltersButton,
	)

	// Masquer les filtres par défaut
	filterContainer.Hide()
	filtersVisible = false
}

// Fonction pour générer une liste d'options de nombres entre min et max (inclus)
func generateMemberOptions(min, max int) []string {
	options := make([]string, max-min+1)
	for i := range options {
		options[i] = strconv.Itoa(min + i)
	}
	return options
}

// generateYearOptions génère une liste d'options d'années entre startYear et endYear
func generateYearOptions(startYear, endYear int) []string {
	years := make([]string, endYear-startYear+1)
	for i := range years {
		years[i] = strconv.Itoa(startYear + i)
	}
	return years
}

func toggleFiltersVisibility() {
	filtersVisible = !filtersVisible
	if filtersVisible {
		filterContainer.Show()
	} else {
		filterContainer.Hide()
	}
	myWindow.Content().Refresh() // Rafraîchir l'affichage pour appliquer les changements
}
func setupGrid() {
	// Initialiser ou réinitialiser la grille avec un nombre défini de colonnes
	grid = container.NewGridWithColumns(5) // Vous pouvez ajuster le nombre de colonnes selon vos besoins

	// Mettre à jour la grille avec les artistes actuels
	updateGrid()
}

func setupGridContainer() *container.Scroll {
	grid = container.NewGridWithColumns(5) // Ajustez selon vos besoins
	updateGrid()

	// Créez un conteneur scrollable en utilisant grid directement
	scrollContainer := container.NewVScroll(grid)
	return scrollContainer
}

func updateGrid() {
	// Appliquer le filtrage et la mise à jour de la grille
	filteredArtists := artists
	if searchTerm != "" && filtersVisible { // Modifier pour tenir compte de la visibilité des filtres
		// Appliquez ici la logique de filtrage en fonction des champs remplis
		applyFilters() // Vous pouvez ajuster cette partie selon vos besoins
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

func applyFilters() {
	var filteredArtists []core.Artist = artistsRef // Commencer avec tous les artistes

	// Appliquer le filtre par date de création s'il est renseigné

	creationDateFrom, creationDateTo := getNumberFromSelect(creationDateFromSelect, is_creationDateToSelect), getNumberFromSelect(creationDateToSelect, is_creationDateFromSelect)
	if creationDateFrom != -1 || creationDateTo != -1 {
		filteredArtists = core.FilterByCreationDate(filteredArtists, creationDateFrom, creationDateTo)
	}

	// Appliquer le filtre par date du premier album s'il est renseigné
	firstAlbumFrom, firstAlbumTo := getNumberFromSelect(firstAlbumFromSelect, is_firstAlbumToSelect), getNumberFromSelect(firstAlbumToSelect, is_firstAlbumFromSelect)
	if firstAlbumFrom != -1 || firstAlbumTo != -1 {
		filteredArtists = core.FilterByFirstAlbumDate(filteredArtists, firstAlbumFrom, firstAlbumTo)
	}

	// Appliquer le filtre par nombre de membres s'il est renseigné
	numMembers := getNumberFromSelect(numMembersSelect, is_numMembersSelect)
	if is_numMembersSelect {
		filteredArtists = core.FilterByNumberOfMembers(filteredArtists, numMembers)
	}

	// Appliquer le filtre par lieu de concert s'il est renseigné

	if is_concertLocationSelect {
		concertLocation := concertLocationSelect.Selected
		filteredArtists = core.FilterByConcertLocation(filteredArtists, concertLocation)
	}

	// Mettre à jour la grille avec les artistes filtrés
	showArtistsGrid(filteredArtists)
}

// getYearFromSelect extrait l'année sélectionnée à partir d'un widget.Select
func getNumberFromSelect(selectWidget *widget.Select, isSelected bool) int {
	if !isSelected {
		return -1
	}
	if selectWidget.SelectedIndex() == -1 {
		return -1
	}
	year, _ := strconv.Atoi(selectWidget.Options[selectWidget.SelectedIndex()])

	return year
}

// Fonction pour afficher les détails de l'artiste
func showArtistDetails(artist core.Artist) {
	fmt.Println("ID =>", artist.Id)
	fmt.Println("Nom =>", artist.Nom)
	fmt.Println("Image =>", artist.Image)
	fmt.Println("First Album =>", artist.FirstAlbum)
	fmt.Println("Concerts liste =>", artist.ConcertDates)
	fmt.Println("Creation Album =>", artist.CreationDate)
	fmt.Println("Relations =>", artist.Relations)
}
