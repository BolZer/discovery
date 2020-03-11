package discovery

type Game string
type Language string

const (
	MHW  Game = "mhw"
	MHGU Game = "mhgu"
	MH4U Game = "mh4u"
)

type Entity struct {
	Name     string
	Url      string
	Game     Game
	Weakness Weakness
}

type Weakness struct {
	Fire    int
	Water   int
	Thunder int
	Ice     int
	Dragon  int
}
