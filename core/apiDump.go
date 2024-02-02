package core

type Artist struct {
	id           int
	image        string
	nom          string
	members      []Member
	creationDate string
	firstalbum   string
	locations    string
	concertdates []Concert
	relations    string
}
type Member struct {
	surname string
	name    string
}

type Concert struct {
	date     Date
	location string
}

type Date struct {
	day   int
	month int
	year  int
}
