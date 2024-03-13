package core

type Artist struct {
	Id           int
	Image        string
	Nom          string
	Members      []Member
	CreationDate string
	FirstAlbum   Date
	Concerts     []Concert
	Relations    string
}

type Member struct {
	Surname string
	Name    string
}

type Concert struct {
	Date     Date
	Location string
}

type Date struct {
	Day   int
	Month int
	Year  int
}
