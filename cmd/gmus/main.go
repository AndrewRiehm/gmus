package main

import (
	"fmt"
	"os"

	musicrepo "github.com/AndrewRiehm/gmus/pkg/music_repository"
)

func main() {
	var weirdAl = musicrepo.Artist("Weird Al")

	fmt.Printf("Hello, %v!\n\n", weirdAl)

	repo := musicrepo.NewMockMusicRepo()

	albums, err := repo.GetAlbumsForArtist(weirdAl)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("%d albums for %v:\n", len(albums), weirdAl)
	for _, album := range albums {
		fmt.Printf("\t%q\n", album.Name)
	}
}
