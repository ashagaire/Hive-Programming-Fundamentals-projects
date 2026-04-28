package main
import ( "fmt"
		"encoding/csv"
		"reflect"
		"bufio"
		"io"
		"strings"
		"regexp"
		"os"
		"time") 

var Reset = "\033[0m"
var Red = "\033[31m"
var Blue = "\033[34m"
var Bold = "\033[1m"
var Yellow = "\033[33m"

func main(){
	if len(os.Args) !=4 || os.Args[1] == "-h" {
		fmt.Println(Blue + Bold +"itinerary usage:\n go run . ./input.txt ./output.txt ./airport-lookup.csv" + Reset)
 		return
	}

	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]
	lookupFilePath := os.Args[3]

	//check input paths are valid
	err := ScanFiles(inputFilePath, lookupFilePath)
	if err != nil {
		fmt.Println(err)
		return 
	}

	//check lookip file is valid
	checkLookupFile := ScanAirportLookup(lookupFilePath)
	if checkLookupFile == false {
		fmt.Println(Red + Bold +"Airport lookup malformed\n" + Reset)
			return 
	}

	fmt.Println(Yellow + Bold +"Welcome to Itinerary Formator App!!!\n"+ Reset)
	processInputFile(inputFilePath, outputFilePath, lookupFilePath)
}

func processInputFile(inputFilePath string, outputFilePath string, lookupFilePath string) {
	processedText := []string{}
	inputFile, err1 := os.Open(inputFilePath)
	if err1 != nil {
		fmt.Println(Red + Bold +"Input not found\n" + Reset)
		return
	} 
	defer inputFile.Close()

	// read the file and save in reader
	reader := bufio.NewReader(inputFile)

	for {
		// read reader untill new line \n
		textLine, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
			break
			}
			fmt.Println(Red + Bold +"Input file read error" + Reset)
			return 
		}
		textLine = strings.TrimSuffix(textLine,"\n")

		// normalize the textLines for different space types like \v, \f, \r whic may return multiple lines
		normalized := cleanLineBreakChars(textLine)
		normalizedLines := strings.Split(normalized, "\n")

		for _, line := range normalizedLines {
			// Step 1: Transformations for airport code and raw date and time values.
			updatedLines := translateCodes(line, lookupFilePath)
			
			//Step 2: Collect the updated lines
			processedText = append(processedText, updatedLines)
		}
	}

	outputdata := strings.Join(processedText, "\n")
	re := regexp.MustCompile(`\n{2,}`)
	finalResult := re.ReplaceAllString(outputdata, "\n\n")
	//Show formated text in terminal

	fmt.Println(Blue + Bold +"Here is your formated Itinary details\n\n" + Reset + finalResult + "\n")

	err2 := os.WriteFile(outputFilePath, []byte(finalResult), 0644)
	if err2 !=nil {
		fmt.Println(Red + Bold +"Error in opeaning output file" + Reset)
	}
}

func translateCodes(text string, lookupFilePath string) string {
	//if word match with any regex replace the word with new replacement 
	icaoRegex := regexp.MustCompile(`^#([A-Z]{3})([^A-Za-z0-9]|$)`)
	iataRegex := regexp.MustCompile(`^##([A-Z]{4})([^A-Za-z0-9]|$)`)
	dateRegex := regexp.MustCompile(`^D\([0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}Z\)$`)
	timeRegex := regexp.MustCompile(`^T12|T24\([0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}.*\)$`)
	
	words := strings.Fields(text)
	for i , word := range words{
		if icaoRegex.MatchString(word) {
			words[i] =  cleanAirportCode(word, "#", 3, lookupFilePath)
		}
		if iataRegex.MatchString(word) {
			words[i] =  cleanAirportCode(word, "##", 4, lookupFilePath)
		}
		if dateRegex.MatchString(word) {
			words[i] = cleanDateFormat(word)
		}
		if timeRegex.MatchString(word) {
			words[i] = cleanTimeFormat(word)
		}
	}
	return strings.Join(words, " ")
}

func cleanTimeFormat(word string) string {
	length := len(word)
	parsedTime := ""
	timeStamp := word[4:length-1]
	t, err := time.Parse("2006-01-02T15:04Z07:00", timeStamp)
		if err != nil {
		return word
	}

	if word[:3] == "T12" {
		parsedTime = t.Format("03:04PM")
	} else if word[:3] == "T24" { 
		parsedTime = t.Format("15:04")
	} else {
		return word
	}

	if word[length-2] == 'Z'{
		return parsedTime+" "+"(+00:00)"
	} else {
		return parsedTime+" "+"("+ word[20:length-1]+")"
	}
}

func cleanDateFormat(word string) string {
	length := len(word)
	dateStamp := word[2:length-1]
	d, err := time.Parse("2006-01-02T15:04Z07:00", dateStamp)
	if err != nil {
		return word
	}
	return (d.Format("02 Jan 2006"))
}

func cleanAirportCode(word string, prefix string, length int, lookupFilePath string) string {
	word = strings.TrimPrefix(word, prefix)
	code := word[:length] //any sign or delimeters afte the airport code from original text

	// find airport name from lookupp file
	airportName, cityName := GetNameFromAirportLookup(code, lookupFilePath)
	//in case the code is not found and we keep the lookup code as it is in putput file
	if code == airportName {
		return ("##"+code+word[length:])
	}
	return airportName+ " "+ cityName+ word[length:]
}

func GetNameFromAirportLookup(code string, lookupFilePath string) (string, string) {
	airportLookup, err := os.Open(lookupFilePath)
	if err != nil {
		fmt.Println(Red + Bold +"Airport lookup malformed\n" + Reset)
		return code, ""
	}
	//defer to close file and prevent resource leaks, no need to do f.Close() later
	defer airportLookup.Close()

	//csv.NewReader to read csv files
	airportLookupReader := csv.NewReader(airportLookup)

	records, err1 := airportLookupReader.ReadAll()
	if err1 != nil {
		fmt.Println(Red + Bold +"Airport lookup malformed\n" + Reset)
		return code , ""
	}

	//loop every records, make a array details with values in each record
	for i:=1; i< len(records)-1; i++ {
		for _, details := range records{
			// compare with each value with current record
			for _, data := range details {
				if data == code {
					return details[0] , details[2]
				}
			}
		}
		
	}
	return code, ""
}

//handeling error with error type
func ScanFiles(inputFile string, lookupFile string) error {
	//os.Stat() retrives information about the file or directory
	_, err1 := os.Stat(inputFile)
	if err1 != nil {
		return fmt.Errorf(Red + Bold +"Input not found\n" + Reset)
	}

	_, err2 := os.Stat(lookupFile)
	if err2 != nil {
		return fmt.Errorf(Red + Bold +"Airport lookup not found\n" + Reset)
	}

	return nil
}

func ScanAirportLookup(lookupFile string) bool {
	indexes := []string{"name", "iso_country","municipality", "icao_code", "iata_code", "coordinates"}

	airportLookup, err := os.Open(lookupFile)
	if err != nil {
		return false
	}
	//defer to close file and prevent resource leaks, no need to do f.Close() later
	defer airportLookup.Close()

	//csv.NewReader to read csv files
	airportLookupReader := csv.NewReader(airportLookup)

	records, err1 := airportLookupReader.ReadAll()
	if err1 != nil {
		return false
	}

	//comparing two string slices with reflect library
	if !reflect.DeepEqual(indexes, records[0]) {
		return false
	}

	checkValues := CheckCsvFileValues(records)
	if checkValues == false {
		return false
	}

	return true
}

func CheckCsvFileValues( records [][]string) bool {
	//loop through each row if any data missing
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

func cleanLineBreakChars(textLine string) string {
	// replace the different space types like \v, \f, \r
	// Using NewReplacer because it  reads through the text in  single pass and only allocates memory once which is more efficient than
	// calling strings.Replace which will read file searching for for every character type we want to swap seperately, multiple times.

	replacer := strings.NewReplacer("\v", "\n", "\f", "\n", "\r", "\n")
	result := replacer.Replace(textLine)
	return result
}
