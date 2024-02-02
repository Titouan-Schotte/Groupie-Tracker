package core

type artist struct {
	id           int
	image        string
	nom          string
	member       []string
	creaDate     string
	firstalbum   string
	locations    string
	concertdates string
	relations    string
}

type locations struct {
	index     []string // se compose des trois variables ci-dessous
	id        int
	locations []string
	dates     string
}

type dates struct {
	index []string
	id    int
	dates []string
}

type relations struct {
	index         []string
	id            int
	ville         string
	dates         string
	dateslocation [][]string
}
