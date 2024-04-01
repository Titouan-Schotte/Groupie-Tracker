/*
Titouan Schotté
App core main
*/
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
	filtersVisible            bool
	filterContainer           *fyne.Container
)

func setupSearchComponents() *widget.Entry {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste")
	// Use the OnChanged event to trigger the search on each keystroke
	searchEntry.OnChanged = func(text string) {
		fmt.Printf("%s", text)
		if text != "" {
			artists = core.SearchInAllStruct(text, artistsRef) // Use artistsRef to search across all artists
		} else {
			artists = artistsRef
		}
		updateGrid()
	}

	return searchEntry
}

func setupFilterComponents() {
	// Generate year options for drop-down menus
	yearOptions := generateYearOptions(1950, 2024)
	//CREATE FILTERS :
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

	memberOptions := generateMemberOptions(1, 20)
	numMembersSelect = widget.NewSelect(memberOptions, func(_ string) {
		is_numMembersSelect = true
	})
	numMembersSelect.PlaceHolder = "Nombre de membres"

	concertLocationSelect = widget.NewSelect(concertsLocations, func(_ string) {
		is_concertLocationSelect = true
	})

	applyFiltersButton := widget.NewButton("Appliquer les filtres", func() {
		applyFilters()
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

	// Hide filters by default
	filterContainer.Hide()
	filtersVisible = false
}

// Function to generate a list of number options between min and max (inclusive)
func generateMemberOptions(min, max int) []string {
	options := make([]string, max-min+1)
	for i := range options {
		options[i] = strconv.Itoa(min + i)
	}
	return options
}

// generateYearOptions generates a list of year options between startYear and endYear
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
	myWindow.Content().Refresh()
}
func setupGrid() {

	// Initialize or reset the grid with a defined number of columns
	grid = container.NewGridWithColumns(5) // Vous pouvez ajuster le nombre de colonnes selon vos besoins

	// Update the grid with current artists
	updateGrid()
}

func setupGridContainer() *container.Scroll {
	grid = container.NewGridWithColumns(5)
	updateGrid()

	// Create a scrollable container using grid directly
	scrollContainer := container.NewVScroll(grid)
	return scrollContainer
}

func updateGrid() {

	// Apply filtering and grid updating
	filteredArtists := artists
	if searchTerm != "" && filtersVisible {
		// Apply filter logic here based on filled fields
		applyFilters()
	}
	showArtistsGrid(filteredArtists)
}

func showArtistsGrid(artists []core.Artist) {
	grid.Objects = nil
	for _, artist := range artists {
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

	grid.Refresh()
}

func loadImageFromURL(urlStr string) *canvas.Image {
	resource, _ := fyne.LoadResourceFromURLString(urlStr)

	image := canvas.NewImageFromResource(resource)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(200, 200))
	return image
}

func applyFilters() {
	var filteredArtists []core.Artist = artistsRef // Start with all artists
	// Apply the filter by creation date if it is specified
	creationDateFrom, creationDateTo := getNumberFromSelect(creationDateFromSelect, is_creationDateToSelect), getNumberFromSelect(creationDateToSelect, is_creationDateFromSelect)
	if creationDateFrom != -1 || creationDateTo != -1 {
		filteredArtists = core.FilterByCreationDate(filteredArtists, creationDateFrom, creationDateTo)
	}

	// Apply the filter by date of the first album if it is specified
	firstAlbumFrom, firstAlbumTo := getNumberFromSelect(firstAlbumFromSelect, is_firstAlbumToSelect), getNumberFromSelect(firstAlbumToSelect, is_firstAlbumFromSelect)
	if firstAlbumFrom != -1 || firstAlbumTo != -1 {
		filteredArtists = core.FilterByFirstAlbumDate(filteredArtists, firstAlbumFrom, firstAlbumTo)
	}

	// Apply the filter by number of members if it is specified
	numMembers := getNumberFromSelect(numMembersSelect, is_numMembersSelect)
	if is_numMembersSelect {
		filteredArtists = core.FilterByNumberOfMembers(filteredArtists, numMembers)
	}

	// Apply the filter by concert venue if it is specified
	if is_concertLocationSelect {
		concertLocation := concertLocationSelect.Selected
		filteredArtists = core.FilterByConcertLocation(filteredArtists, concertLocation)
	}

	// Update grid with filtered artists
	showArtistsGrid(filteredArtists)
}

// getYearFromSelect extracts the selected year from a widget.Select
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

	image := preloaderImagesForPopup[artist.Id]
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
		// Navigation buttons placed on the right
		navigationContainer := container.NewHBox(
			layout.NewSpacer(), // This spacer pushes the buttons to the right
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
		fullContent = container.NewVBox(content, navigationContainer)
	}

	spotifyButton := widget.NewButton("Écouter sur Spotify", func() {
		searchQuery := strings.ReplaceAll(artist.Nom, " ", "%20")
		urlStr := "https://open.spotify.com/search/" + searchQuery
		parsedUrl, _ := url.Parse(urlStr)
		fyne.CurrentApp().OpenURL(parsedUrl)
	})

	contentWithHeaderAndFooter := container.NewVBox(fullContent, spotifyButton)

	// Creation and display of the popup
	popUp := widget.NewPopUp(contentWithHeaderAndFooter, myWindow.Canvas())
	popUp.Show()
	popUp.Resize(fyne.NewSize(600, 400))
	popUp.Move(fyne.NewPos(myWindow.Canvas().Size().Width/2-300, myWindow.Canvas().Size().Height/2-200))

	// Display the first page of concerts
	updateConcertsPage(0)
}

func showMapPopup(location string) {
	backgroundImageRessource := core.GenerateMapImage(location)
	backgroundImage := canvas.NewImageFromResource(backgroundImageRessource)
	backgroundImage.FillMode = canvas.ImageFillStretch
	content := container.NewMax(backgroundImage)

	popUp := widget.NewPopUp(content, myWindow.Canvas())
	popUp.Show()
	popUp.Resize(fyne.NewSize(1200, 800))
	popUp.Move(fyne.NewPos(myWindow.Canvas().Size().Width/2-600, myWindow.Canvas().Size().Height/2-400))
}
