package main

import (
	"fmt"

	musicrepo "github.com/AndrewRiehm/gmus/pkg/music_repository"
)

func main() {
	var weirdAl = musicrepo.Artist("Weird Al")
	fmt.Printf("hello, %v!\n", weirdAl)
}
