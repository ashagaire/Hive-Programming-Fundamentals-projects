package main
import ( "fmt" 
		"os"
		"unicode"
			)
			var Reset = "\033[0m"
// var Red = "\033[31m"
var Blue = "\033[34m"
var Bold = "\033[1m"
// var Yellow = "\033[33m"
var art string

func main() {
	
	if len(os.Args) !=2 || os.Args[1] == "-h" {
		fmt.Println(Blue + Bold +"itinerary usage:\n go run . (Your code for Art)" + Reset)
			return
	}

	inputArtCode := os.Args[1]
	art , err := displayArt( inputArtCode)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(art)
	return

}

func displayArt( inputArtCode string) (string, error) {
	code := inputArtCode
	closingIndex:= 0
	for i:= 0 ;i < len(code); i++ {
		if code[i] == '[' {
			closingIndex = findClosingBracket(code, i)
			if closingIndex == 0 {
				return "", fmt.Errorf("Error:error in close bracket not found")
			}
			text, err := printArt(code[i+1 :closingIndex])
			if err != nil {
				fmt.Println(err)
				return "", fmt.Errorf("Error: error in text")
			}
			art = art + text
			i = closingIndex
		} else {
			art = art+string(code[i])
		}
	}
	return string(art) , nil

}

type DecodeItems struct {
	repeat string 
	artChars string
} 

func printArt(artPattern string) (string, error) {
	 repetation := DecodeItems{repeat : "", artChars: ""}
	 //split with first " " found
	
	 //check if the values are numerical and collect it on a value
	 //if the value is space dump everything in last variable
	//check first is a type number

	if  !unicode.IsNumber(rune(artPattern[0])){
		return "", fmt.Errorf("Error:first value  is not number")
	}


	for i:=0 ; i< len(artPattern); i++ {
		if unicode.IsNumber(rune(artPattern[i])) {
			
			repetation.repeat = repetation.repeat + string(artPattern[i])
			fmt.Println(repetation.repeat)
			continue
		} else if artPattern[i] == ' ' {
			fmt.Println(string(artPattern[i+1]))
			repetation.artChars = string(artPattern[i+1:])
			fmt.Println(repetation.artChars)
			break
		} else {
			return "", fmt.Errorf("Error:error in reading patreern code")
		}
	}

	if (repetation.repeat == "" || repetation.artChars == "") {
		return "", fmt.Errorf("Error:error in struct values")
	} 
	

	return repetation.artChars, nil
}


func findClosingBracket(code string, start int ) int{
	for i:= start ;i < len(code); i++ {
		if code[i] == ']' {
			return i
		}
	}
	return 0
}