package main
import ( 
		"fmt" 
		)

func singlelineDecode(inputArtCode string) {
	art , err := readArtCode( inputArtCode)
	if err != nil {
		fmt.Println(Red+ Bold + err.Error() + Reset)
		return
	}

	fmt.Println("\n" + art + "\n")
	return
}