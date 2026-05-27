package main 
import ( 
		"fmt" 
		"strconv"
		"strings"
		)

func readArtCode( inputArtCode string) (string, error) {
	code := inputArtCode
	art := ""
	
	for i:= 0 ; i < len(code); i++ {
		if code[i] == '[' {
			closingIndex :=strings.Index(code[i:], "]")		
			if closingIndex == -1 {
				// Error:error in close bracket not found
				return "", fmt.Errorf("\nError\n")
			}
			closingIndex += i
			text, err := ConvertToArt(code[i+1 :closingIndex])
			if err != nil {
				// Error: error in text
				return "", fmt.Errorf("\nError\n")
			}
			art = art + text
			i = closingIndex
		} else {
			art = art+string(code[i])
		}
	}
	return art , nil
}

func ConvertToArt(artPattern string) (string, error) {
	before, after, found := strings.Cut(artPattern, " ")
	if found == false {
		// Error: no valid values inside brackerts
		return "", fmt.Errorf("\nError\n")
	}
	repeats, err := strconv.Atoi(before)
	if err != nil {
		// Error:first value  is not number
		return "", fmt.Errorf("\nError\n")
	}
	if after == "" {
		// Error:missing second value inside brackets
		return "", fmt.Errorf("\nError\n")
	}

	output := strings.Repeat(after, repeats)
	return output, nil	
}