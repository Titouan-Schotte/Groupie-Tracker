package main

import (
	"Groupie_Tracker/core"
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
	artistsRef            = []core.Artist{
		{
			Id:           1,
			Image:        "https://groupietrackers.herokuapp.com/api/images/acdc.jpeg",
			Nom:          "Artiste 1",
			Members:      []core.Member{{Surname: "Doe", Name: "John"}},
			CreationDate: "01/01/2020",
			FirstAlbum:   core.Date{Day: 1, Month: 1, Year: 2021},
			Concerts:     []core.Concert{{Date: core.Date{Day: 15, Month: 2, Year: 2022}, Location: "Location 1"}},
			Relations:    "Relations 1",
		},
		{
			Id:           2,
			Image:        "https://groupietrackers.herokuapp.com/api/images/acdc.jpeg",
			Nom:          "Artiste 2",
			Members:      []core.Member{{Surname: "Smith", Name: "Jane"}},
			CreationDate: "02/02/2019",
			FirstAlbum:   core.Date{Day: 2, Month: 2, Year: 2020},
			Concerts:     []core.Concert{{Date: core.Date{Day: 31, Month: 3, Year: 2024}, Location: "Location 2"}},
			Relations:    "Relations 2",
		},
		{
			Id:           3,
			Image:        "https://groupietrackers.herokuapp.com/api/images/acdc.jpeg",
			Nom:          "Artiste 3",
			Members:      []core.Member{{Surname: "Doe", Name: "Jane"}},
			CreationDate: "03/03/2021",
			FirstAlbum:   core.Date{Day: 3, Month: 3, Year: 2022},
			Concerts:     []core.Concert{{Date: core.Date{Day: 10, Month: 4, Year: 2023}, Location: "Location 3"}},
			Relations:    "Relations 3",
		},
		{
			Id:           4,
			Image:        "https://groupietrackers.herokuapp.com/api/images/acdc.jpeg",
			Nom:          "Artiste 4",
			Members:      []core.Member{{Surname: "Doe", Name: "John"}},
			CreationDate: "01/01/2020",
			FirstAlbum:   core.Date{Day: 1, Month: 1, Year: 2021},
			Concerts:     []core.Concert{{Date: core.Date{Day: 15, Month: 2, Year: 2022}, Location: "Location 1"}},
			Relations:    "Relations 1",
		},
		{
			Id:           5,
			Image:        "https://groupietrackers.herokuapp.com/api/images/acdc.jpeg",
			Nom:          "Artiste 5",
			Members:      []core.Member{{Surname: "Smith", Name: "Jane"}},
			CreationDate: "02/02/2019",
			FirstAlbum:   core.Date{Day: 2, Month: 2, Year: 2020},
			Concerts:     []core.Concert{{Date: core.Date{Day: 31, Month: 3, Year: 2024}, Location: "Location 2"}},
			Relations:    "Relations 2",
		},
		{
			Id:           6,
			Image:        "https://groupietrackers.herokuapp.com/api/images/acdc.jpeg",
			Nom:          "Artiste 6",
			Members:      []core.Member{{Surname: "Ko", Name: "Jane"}},
			CreationDate: "03/03/2021",
			FirstAlbum:   core.Date{Day: 3, Month: 3, Year: 2022},
			Concerts:     []core.Concert{{Date: core.Date{Day: 10, Month: 4, Year: 2023}, Location: "Location 3"}},
			Relations:    "Relations 3",
		},
	}
)

func main() {
	myApp = app.New()
	myWindow = myApp.NewWindow("Groupie Tracker")

	artists = artistsRef
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste")

	// Calculer la largeur minimale nécessaire pour la searchEntry
	placeholderWidth := float32(800)

	// Définir la largeur minimale de la searchEntry
	searchEntry.Resize(fyne.NewSize(placeholderWidth+100, searchEntry.MinSize().Height)) // Ajouter 100 pixels à la largeur

	// Créer un bouton de recherche
	searchButton := widget.NewButton("Rechercher", func() {
		searchTerm = searchEntry.Text
		artists = core.SearchInAllStruct(searchTerm, artistsRef)
		updateGrid()
	})

	// Créer des champs d'entrée pour les filtres de date de création et de premier album
	creationDateFromEntry = widget.NewEntry()
	creationDateFromEntry.SetPlaceHolder("Année de début")

	creationDateToEntry = widget.NewEntry()
	creationDateToEntry.SetPlaceHolder("Année de fin")

	firstAlbumFromEntry = widget.NewEntry()
	firstAlbumFromEntry.SetPlaceHolder("Année de début")

	firstAlbumToEntry = widget.NewEntry()
	firstAlbumToEntry.SetPlaceHolder("Année de fin")

	// Créer un champ d'entrée pour le filtre par nombre de membres
	numMembersEntry = widget.NewEntry()
	numMembersEntry.SetPlaceHolder("Nombre de membres")

	// Créer une liste déroulante pour le filtre par lieu de concert
	concertLocationSelect = widget.NewSelect(ConcertLocations, nil)

	// Créer un bouton pour appliquer les filtres
	applyFiltersButton := widget.NewButton("Appliquer les filtres", applyFilters)

	// Créer un conteneur pour organiser les champs de filtrage
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

	// Créer une disposition personnalisée pour la searchBox avec la searchEntry agrandie et le bouton de recherche
	searchBox := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		widget.NewLabel("Search: "),
		container.New(layout.NewMaxLayout(), searchEntry), // Utilisez un conteneur pour permettre à la searchEntry de s'étendre
		searchButton,
	)

	// Redimensionner la searchBox
	searchBox.Resize(fyne.NewSize(placeholderWidth+100, searchBox.MinSize().Height))
	grid = container.NewGridWithColumns(5)
	//scrollContainer := container.NewVScroll(grid)

	updateGrid()

	content := container.NewBorder(searchBox, nil, nil, nil, grid)
	myWindow.SetContent(container.NewVBox(filterFields, content))
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}

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

		label := widget.NewLabel(artist.Nom)

		// Create a container for the image
		imageContainer := container.New(layout.NewMaxLayout(), image)

		// Create a container for the label and center it at the bottom of the image
		labelContainer := container.New(layout.NewHBoxLayout(),
			layout.NewSpacer(),
			label,
			layout.NewSpacer(),
		)

		// Adjust the size of the label container to match the image size
		labelContainer.Resize(fyne.NewSize(image.Size().Width, label.MinSize().Height))

		// Combine the image container and label container vertically
		card := container.New(layout.NewVBoxLayout(), imageContainer, labelContainer)

		// Add the card to the grid
		grid.Add(card)
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
