package app

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var Logo, _ = fyne.LoadResourceFromPath("./logo.png")

var App = app.New()

func Accueil(a fyne.App) {

	w := a.NewWindow("Groupie Tracker ★")
	w.SetIcon(Logo)

	// Créer un texte avec le contenu "Groupie Tracker"
	text := canvas.NewText("Groupie Tracker", color.White)
	text.TextSize = 36
	text.TextStyle.Bold = true // Appliquer le style de police gras

	// Définit l'alignement du texte au centre
	text.Alignment = fyne.TextAlignCenter

	// Crée un texte de description
	description := canvas.NewText("Bienvenue sur Groupie Tracker ! ", color.White)
	descriptiondeux := canvas.NewText("Cette application vous aide à suivre vos artistes préférés et à en découvrir de nouveaux.", color.White)
	description.TextSize = 16
	description.Alignment = fyne.TextAlignCenter
	descriptiondeux.Alignment = fyne.TextAlignCenter

	// Crée une barre de navigation avec des onglets
	navBar := container.NewHBox(
		widget.NewButton("Mes favoris", func() {
			// Gérer l'onglet Mes favoris clic
		}),
		widget.NewButton("Artiste", func() {
			// Gérer l'onglet Artiste click
		}),
		widget.NewButton("Map", func() {
			Map(a)
			w.Hide()
			// Cliquez sur l'onglet Gérer la carte
		}),
		widget.NewEntry(), // Barre de recherche
	)

	searchEntry := navBar.Objects[3].(*widget.Entry)
	searchEntry.Resize(fyne.NewSize(1000, searchEntry.Size().Width)) // MARCHE PAS

	// Définir la mise en page du contenu
	content := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		text,
		layout.NewSpacer(),
		description,
		descriptiondeux,
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
