package main
import ( "fmt" 
		
		"strconv"
		"strings"
		"flag"
		"bufio"
		"os"
			)
			var Reset = "\033[0m"
// var Red = "\033[31m"
var Blue = "\033[34m"
var Bold = "\033[1m"
// var Yellow = "\033[33m"
// var art string

func main() {
		
	if len(os.Args) !=2 || os.Args[1] == "-h" {
		fmt.Println(Blue + Bold +"itinerary usage:\n go run . (Your code for Art)" + Reset)
			return
	}

	multiLine := flag.Bool("m", false, "enable multiline mode")

	flag.Parse()
	// Multiline Mode
	if *multiLine {
		scanner := bufio.NewScanner(os.Stdin)
		results := []string{}

		for scanner.Scan() {
			line := scanner.Text()

			art, err := displayArt( line)
		
			if err != nil {
				fmt.Println(err)
				return
			}
			
			results = append(results,art)
			
			
			
		}
		fmt.Println()
		for _, result := range results {
			fmt.Println(result)
		}
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
	art := ""
	
	for i:= 0 ; i < len(code); i++ {
		if code[i] == '[' {
			closingIndex :=strings.Index(code[i:], "]")		
			if closingIndex == -1 {
				// Error:error in close bracket not found
				return "", fmt.Errorf("Error\n")
			}
			closingIndex += i
			text, err := printArt(code[i+1 :closingIndex])
			if err != nil {
				// Error: error in text
				return "", fmt.Errorf("Error\n")
			}
			art = art + text
			i = closingIndex
		} else {
			art = art+string(code[i])
		}
	}
	return art , nil
}

func printArt(artPattern string) (string, error) {
	before, after, found := strings.Cut(artPattern, " ")
	if found == false {
		// Error: no valid values inside brackerts
		return "", fmt.Errorf("Error\n")
	}
	repeats, err := strconv.Atoi(before)
	if err != nil {
		// Error:first value  is not number
		return "", fmt.Errorf("Error\n")
	}

	output := strings.Repeat(after, repeats)
	return output, nil	
}

