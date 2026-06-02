package main
import "fmt"

func decoder(inputArtCode string) (string, error) {
	art , err := readArtCode( inputArtCode)
	if err != nil {
		
		return "", fmt.Errorf("\nError in your art encode text\n")
	}

	return art, nil
}