package musicrepository

// NOTE: This is just a first pass at this...
type Song struct {
	Name string
}

type Artist string

type Album struct {
	Artist Artist
	Songs  []Song
}

type MusicRepository interface {
	GetArtists() ([]Artist, error)
	GetAlbumsForArtist(Artist) ([]Album, error)
}
