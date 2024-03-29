package main

import (
	"Groupie_Tracker/core"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"net/url"
	"strconv"
	"strings"
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
	artistsRef                []core.Artist
	filtersVisible            bool // Nouveau : état de visibilité des filtres
	filterContainer           *fyne.Container
)

func setupSearchComponents() *widget.Entry {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste")
	// Utiliser l'événement OnChanged pour déclencher la recherche à chaque frappe
	searchEntry.OnChanged = func(text string) {
		fmt.Printf("%s", text)
		if text != "" {
			artists = core.SearchInAllStruct(text, artistsRef) // Utilisez artistsRef pour rechercher parmi tous les artistes
		} else {
			artists = artistsRef
		}
		updateGrid()
	}

	return searchEntry
}

func setupFilterComponents() {
	// Générer les options d'années pour les menus déroulants
	yearOptions := generateYearOptions(1950, 2024)
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

	concertLocationSelect = widget.NewSelect(concertsLocations, func(_ string) {
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
		currentArtist := artist // Créer une copie locale de la variable pour la capture
		image := preloaderImages[artist.Id]
		// Limiter la taille maximale de l'image à 500x500
		if image.Size().Width > 500 || image.Size().Height > 500 {
			image.Resize(fyne.NewSize(500, 500))
		}
		// Créer un conteneur pour l'image
		imageContainer := container.New(layout.NewMaxLayout(), image)

		// Ajouter un bouton avec le nom de l'artiste en dessous de l'image
		button := widget.NewButton(currentArtist.Nom, func() {
			showArtistDetails(currentArtist) // Utiliser la copie locale dans la fermeture
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

func showArtistDetails(artist core.Artist) {
	nameLabel := widget.NewLabelWithStyle(artist.Nom, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	image := preloaderImagesForPopup[artist.Id] // Utilisation hypothétique de loadImageFromURL pour l'image
	image.SetMinSize(fyne.NewSize(200, 200))

	firstAlbumLabel := widget.NewLabel("First Album: " + artist.FirstAlbum)
	creationDateLabel := widget.NewLabel(fmt.Sprintf("Creation Album: %d", artist.CreationDate))

	detailsContainer := container.NewVBox(nameLabel, firstAlbumLabel, creationDateLabel, widget.NewLabel("Membre(s) de la formation :"))
	for _, member := range artist.Members {
		detailsContainer.Add(widget.NewLabel(" - " + member))
	}

	concertsContainer := container.NewVBox(widget.NewLabel("Concerts :"))
	content := container.NewHBox(detailsContainer, image, concertsContainer) // Contenu à mettre à jour avec les concerts

	// Pagination
	currentPage := 0
	totalPages := (len(artist.ConcertDates)-1)/10 + 1

	updateConcertsPage := func(page int) {
		concertsContainer.Objects = []fyne.CanvasObject{widget.NewLabel("Concerts :")}
		start := page * 10
		end := start + 10
		if end > len(artist.ConcertDates) {
			end = len(artist.ConcertDates)
		}

		for _, concert := range artist.ConcertDates[start:end] {
			concertIn := concert
			concertsContainer.Add(widget.NewButton(fmt.Sprintf("- à %s le %d-%d-%d", strings.Split(concertIn.Location, "-")[0], concertIn.Date.Day, concertIn.Date.Month, concertIn.Date.Year), func() {
				showMapPopup(strings.Split(concertIn.Location, "-")[0]) // Adapt this to your implementation
			}))
		}

		content.Refresh()
	}

	fullContent := container.NewVBox(content)
	if totalPages > 1 {
		// Boutons de navigation placés à droite
		navigationContainer := container.NewHBox(
			layout.NewSpacer(), // Ce spacer pousse les boutons vers la droite
			widget.NewButton(" Précédent ", func() {
				if currentPage > 0 {
					currentPage--
					updateConcertsPage(currentPage)
				}
			}),
			widget.NewButton("   Suivant    ", func() {
				if currentPage < totalPages-1 {
					currentPage++
					updateConcertsPage(currentPage)
				}
			}),
		)
		// Mise en page avec navigation
		fullContent = container.NewVBox(content, navigationContainer)
	}

	spotifyButton := widget.NewButton("Écouter sur Spotify", func() {
		searchQuery := strings.ReplaceAll(artist.Nom, " ", "%20")
		urlStr := "https://open.spotify.com/search/" + searchQuery
		parsedUrl, _ := url.Parse(urlStr)
		fyne.CurrentApp().OpenURL(parsedUrl)
	})

	contentWithHeaderAndFooter := container.NewVBox(fullContent, spotifyButton)

	// Création et affichage de la popup
	popUp := widget.NewPopUp(contentWithHeaderAndFooter, myWindow.Canvas())
	popUp.Show()
	popUp.Resize(fyne.NewSize(600, 400))
	popUp.Move(fyne.NewPos(myWindow.Canvas().Size().Width/2-300, myWindow.Canvas().Size().Height/2-200))

	// Affichez la première page de concerts
	updateConcertsPage(0)
}

func showMapPopup(location string) {
	backgroundImageRessource := core.GenerateMapImage(location)
	backgroundImage := canvas.NewImageFromResource(backgroundImageRessource)
	backgroundImage.FillMode = canvas.ImageFillStretch // Ajuster pour remplir l'espace
	content := container.NewMax(backgroundImage)

	popUp := widget.NewPopUp(content, myWindow.Canvas())
	popUp.Show()
	popUp.Resize(fyne.NewSize(1200, 800))
	popUp.Move(fyne.NewPos(myWindow.Canvas().Size().Width/2-600, myWindow.Canvas().Size().Height/2-400))
}
