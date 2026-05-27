package main


func singlelineDecode(inputArtCode string) string {
	art , err := readArtCode( inputArtCode)
	if err != nil {
		
		return "Error"
	}

	return art
}