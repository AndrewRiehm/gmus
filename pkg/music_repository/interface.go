package musicrepository

import "fmt"

// NOTE: This is just a first pass at this...
type Song struct {
	Name string
}

type Artist string

type Album struct {
	Artist Artist
	Name   string
	Songs  []Song
}

type MusicRepository interface {
	GetArtists() ([]Artist, error)
	GetAlbumsForArtist(Artist) ([]Album, error)
}

// Mock implementation
type MockMusicRepo struct {
	artists []Artist
	albums  []Album
}

func (m *MockMusicRepo) GetArtists() ([]Artist, error) {
	return m.artists, nil
}

func (m *MockMusicRepo) GetAlbumsForArtist(artist Artist) ([]Album, error) {

	results := []Album{}
	for _, album := range m.albums {
		if album.Artist == artist {
			results = append(results, album)
		}
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no albums found")
	}

	return results, nil
}

func NewMockMusicRepo() MusicRepository {
	artists := []Artist{
		Artist("Weird Al"),
	}
	albums := []Album{
		{
			Artist: artists[0],
			Name:   "Dare To Be Stupid",
			Songs: []Song{
				{Name: "Slime Creatures From Outter Space"},
			},
		},
		{
			Artist: artists[0],
			Name:   "Even Worse",
			Songs: []Song{
				{Name: "I Think I'm A Clone Now"},
			},
		},
	}
	return &MockMusicRepo{
		artists: artists,
		albums:  albums,
	}
}
