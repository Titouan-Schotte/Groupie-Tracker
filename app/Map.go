package app

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Map(a fyne.App) {
	w := a.NewWindow("Groupie Tracker ★")
	w.SetIcon(Logo)

	// Créer un texte avec le contenu "Groupie Tracker"
	text := canvas.NewText("Groupie Tracker", color.White)
	text.TextSize = 16
	text.TextStyle.Bold = true             // Appliquer le style de police gras
	text.Alignment = fyne.TextAlignLeading // Définir l'alignement à gauche

	//Crée un text avec le contenu "Map"
	text2 := canvas.NewText("Map", color.White)
	text2.TextSize = 26
	text2.TextStyle.Bold = true            // Appliquer le style de police gras
	text2.Alignment = fyne.TextAlignCenter // Définir l'alignement au centre

	// Crée une barre de navigation avec des onglets
	navBar := container.NewHBox(
		widget.NewButton("Mes favoris", func() {
			// Gérer l'onglet Accueil clic
		}),
		widget.NewButton("Accueil", func() {
			Accueil(a)
			w.Hide()
			// Gérer l'onglet Accueil clic
		}),
		widget.NewButton("Artiste", func() {
			// Gérer l'onglet Artiste click
		}),
		widget.NewButton("Map", func() {
			// Cliquez sur l'onglet Gérer la carte
		}),
		widget.NewEntry(), // Barre de recherche
	)

	// searchEntry := navBar.Objects[3].(*widget.Entry)
	// searchEntry.Resize(fyne.NewSize(1000, searchEntry.Size().Width)) // MARCHE PAS

	// Définir la mise en page du contenu
	content := container.New(
		layout.NewVBoxLayout(),
		text,
		text2,
		layout.NewSpacer(),
		navBar,
	)

	// Crée un rectangle pour l'arrière-plan
	backgroundRect := canvas.NewRectangle(color.Black)

	// Place le rectangle noir derrière le contenu
	backgroundContainer := container.New(
		layout.NewBorderLayout(nil, nil, nil, nil),
		backgroundRect, content,
	)

	w.SetContent(backgroundContainer)
	w.Resize(fyne.NewSize(900, 500))
	w.Show()
	w.CenterOnScreen()
	w.SetOnClosed(func() {
		a.Quit()
	})

}
