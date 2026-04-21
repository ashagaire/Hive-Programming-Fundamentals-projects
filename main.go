package main
import ( "fmt"
"encoding/csv"
"reflect"
// "path/filepath"
"bufio"
"io"
"strings"
"regexp"
		"os") 

func main(){

	if len(os.Args) !=4 || os.Args[1] == "-h" {
		fmt.Println("itinerary usage:\n go run . ./input.txt ./output.txt ./airport-lookup.csv")
 		return
	}

	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]
	lookupFilePath := os.Args[3]

	//check input paths are valid
	err := ScanInputs(inputFilePath, lookupFilePath)
	if err != nil {
		fmt.Println(err)
		 
		 return 
	}

	//check lookip file is valid
	checkLookupFile := ScanAirportLookup(lookupFilePath)
	if checkLookupFile == false {
		fmt.Println("Airport lookup malformed\n")
			return 
		}

	fmt.Println("Welcome to Itinerary Formator App!!!\n")
	// createOutputFile(outputFilePath)
	processInputFile(inputFilePath, outputFilePath)

	//now scan the input file and write output file

}

func processInputFile(inputFilePath string, outputFilePath string) {

	// verify input file
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		 fmt.Println("Input not found")
		 return
	} 
	defer inputFile.Close()

	//create output file 
	//File Permissions 6:(Owner): Can Read and Write, 4:(Group): Can Read only, 4:(Others): Can Read only.
	outputFile, err :=os.OpenFile(outputFilePath, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644 )
	if err != nil {
		fmt.Println("Error in opeaning output file")
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()


	// read the file and save in reader
	reader := bufio.NewReader(inputFile)

	for {
		// read reader untill new line \n
		textLine, err := reader.ReadString('\n')
		
		if err != nil {
			if err == io.EOF {
			break
			}
			fmt.Println("Input file read error")
			return 
		}

		if textLine == "" {
			continue
		}
		
		// normalize the textLines for different space types like \v, \f, \r whic may return multiple lines
		normalized := cleanLineBreakChars(textLine)

		lines := strings.Split(normalized, "\n")

		for _, line := range lines {
			// fmt.Println("unformatedLine:", line)
		// Step 1: Transformations for airport code.
		line = convertAirportCodes(line)

		// Step 2: Transformations foor dates
        // line = formatDates(line)

		// Step 3: Write to output.txt
        writer.WriteString(line + "\n")
		}
		
        
	}
	
}

func convertAirportCodes(text string) string {
	//if word match with any regex find the replace word from csv column
	//replace the word with new replacement 
	icaoRegex := regexp.MustCompile(`^#([A-Z]{3})([^A-Za-z0-9]|$)`)
	iataRegex := regexp.MustCompile(`^##([A-Z]{4})([^A-Za-z0-9]|$)`)
	// dateRegex := regexp.MustCompile(`^#`)
	// time12Regex := regexp.MustCompile(`^#`)
	// time24Regex := regexp.MustCompile(`^#`)
	
	words := strings.Fields(text)
	for i , word := range words{
		if icaoRegex.MatchString(word) {
			words[i] = cleanAirportCode(word, "#", 3)
		}
		if iataRegex.MatchString(word) {
			words[i] = cleanAirportCode(word, "##", 4)
		}

	}
	formatedLine := strings.Join(words, " ")
	return formatedLine
}

func cleanAirportCode(word string, prefix string, length int) string {
	word = strings.TrimPrefix(word, prefix)
	code := word[:length]


	fmt.Println("the only code from IcaoLookup code", code)
	// find airport name from lookup

	return code
}

func cleanLineBreakChars(textLine string) string {
	// replace the different space types like \v, \f, \r
	// Using NewReplacer because it  reads through the text in  single pass and only allocates memory once which is more efficient than
	// calling strings.Replace which will read file searching for for every character type we want to swap seperately, multiple times.

	replacer := strings.NewReplacer("\v", "\n", "\f", "\n", "\r", "\n")
	result := replacer.Replace(textLine)
	return result

}




//TODO functions
// for airport code in input.txt find row with that code in respetiive code column in csv file, if not found keep as it is
// if D(2007-04-05T12:30−02:00) found in input.txt output date in DD-Mmm-YYYY format
// if T12(2007-04-05T12:30−02:00)  found in input.txt output 12:30PM (-02:00)
// if T24(2007-04-05T12:30−02:00) found in input.txt output 12:30 (-02:00)
//  if there is "Z" after "T" in T12:30−02:00 then  show (+00:00)
// if the line after a blank line is also blank then remove it ans break the reading loop

// TODO output file 
// TO replace the different space types like \v, \f, \r
// Using NewReplacer because it  reads through the text in  single pass and only allocates memory once
//  which is more efficient than calling strings.Replace which will read file searching for for every character type we want to swap seperately, multiple times.
// replacer := strings.NewReplacer("\v", "\n", "\f", "\n", "\r", "\n")
// result := replacer.Replace(content)

//TO clean multile empty lines
//  Using Regex to find 2 or more newlines and replace them with just 2.
// syntax is x{n,} find matiching x but n times or more x, in this case x= '\n' and replacing x{n,} with '\n\n'
	// re := regexp.MustCompile(`\n{2,}`)
	// finalResult := re.ReplaceAllString(normalized, "\n\n")



	//handeling error with error type
func ScanInputs(inputFile string, lookupFile string) error {

	//os.Stat() retrives information about the file or directory
	_, err1 := os.Stat(inputFile)
	if err1 != nil {
		return fmt.Errorf("Input not found")
	}

	_, err2 := os.Stat(lookupFile)
	if err2 != nil {
		return fmt.Errorf("Airport lookup not found")
	}

	return nil

}

func ScanAirportLookup(lookupFile string) bool {
	// in csv file if any column in the lookup is corrupted, missing or blank display "Airport lookup malformed"

	indexes := []string{"name", "iso_country","municipality", "icao_code", "iata_code", "coordinates"}

	airportLookup, err := os.Open(lookupFile)
	if err != nil {
		return false
	}
	//defer to close file and prevent resource leaks, no need to do f.Close() later
	defer airportLookup.Close()

	//csv.NewReader to read csv files
	airportLookupReader := csv.NewReader(airportLookup)

	//data seperator
	// reader.Comma = '-'
	records, err1 := airportLookupReader.ReadAll()
	if err1 != nil {
		return false
	}

	checkColumns := CheckIndexRow(indexes,records[0])
	if checkColumns == false {
		return false
	}

	checkValues := CheckCsvFileValues(records)
	if checkValues == false {
		return false
	}

	return true

}

func CheckIndexRow(indexes []string, firstRow []string) bool{
	//comparing two string slices with reflect library
	return reflect.DeepEqual(indexes, firstRow)
}

func CheckCsvFileValues( records [][]string) bool {
	//loop through each row
	for _, row:= range records {
			if len(row) != 6 {
			return false
			}
			for _, r := range row {
				if r == "" {
					return false
				}
			}
	}
	return true
}

// func createOutputFile(outputFilePath string) (*os.File, error ) {
// 	//get oly file name from outputFilePath if we are supposed to always create output file in root folder
// 	// fileName := filepath.Base(outputFilePath)

// 	//File Permissions 6:(Owner): Can Read and Write, 4:(Group): Can Read only, 4:(Others): Can Read only.
// 	file, err := os.OpenFile(outputFilePath, os.O_CREATE | os.O_RDWR, 0644) 

// 	if err != nil {
// 		fmt.Println("Error in creating output file\n")
// 		return nil, err
// 	}

// 	defer file.Close()
// 	fmt.Println("Output file created and closed\n")
// }