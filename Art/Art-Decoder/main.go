package main
import ( 
		"fmt" 
		"flag"
		"os"
		)
			
var Reset = "\033[0m"
var Red = "\033[31m"
var Blue = "\033[34m"
var Bold = "\033[1m"
var Yellow = "\033[33m"

func main() {
		
	if len(os.Args) !=2 || os.Args[1] == "-h" {
		fmt.Println(Blue + Bold +"\nItinerary usage:\n" + Reset)
		fmt.Println(Yellow + Bold + "For single line art \n" + Blue + Bold + "go run . (Your code for Art)\n" + Reset)
		fmt.Println(Yellow + Bold + "For multi line art \n" + Blue + Bold + "go run . -m \n" + Yellow + Bold + "Write your multi line art code and then \"Ctl + D\" (for windows \"Ctl + D\")\n " + Reset)
			return
	}

	multiLine := flag.Bool("m", false, "enable multiline mode")

	flag.Parse()
	// Multiline Mode
	if *multiLine {
		multilineDecode()
		return
	}

	inputArtCode := os.Args[1]
	singlelineDecode(inputArtCode)
	return
}
