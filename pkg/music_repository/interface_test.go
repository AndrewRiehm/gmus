package musicrepository_test

import (
	"testing"

	musicrepository "github.com/AndrewRiehm/gmus/pkg/music_repository"
	"github.com/stretchr/testify/assert"
)

func TestStuff(t *testing.T) {
	repo := musicrepository.NewMockMusicRepo()
	t.Run("should be able to make a new mock repo", func(t *testing.T) {
		assert.NotNil(t, repo)
	})
	t.Run("should be able to get all the artists", func(t *testing.T) {
		repo := musicrepository.NewMockMusicRepo()
		got, err := repo.GetArtists()
		assert.NoError(t, err)
		assert.NotEmpty(t, got)
	})
	t.Run("should be able to get albums for an artist", func(t *testing.T) {
		got, err := repo.GetArtists()
		assert.NoError(t, err)
		assert.NotEmpty(t, got)

		for _, artist := range got {
			albums, err := repo.GetAlbumsForArtist(artist)
			assert.NoError(t, err)
			assert.NotEmpty(t, albums)
		}
	})
	t.Run("should complain if there aren't any albums", func(t *testing.T) {
		albums, err := repo.GetAlbumsForArtist("madonna")
		assert.Error(t, err)
		assert.Nil(t, albums)
	})
}
