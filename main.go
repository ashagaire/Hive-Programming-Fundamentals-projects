package main
import ( "fmt"
		"encoding/csv"
		"reflect"
		"bufio"
		"io"
		"strings"
		"regexp"
		"strconv"
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

	processInputFile(inputFilePath, outputFilePath, lookupFilePath)
}

func processInputFile(inputFilePath string, outputFilePath string, lookupFilePath string) {
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
			// Step 1: Transformations for airport code.
			line = convertAirportCodes(line, lookupFilePath)

			// Step 2: Transformations foor dates
			// line = formatDates(line)

			// Step 3: Write to output.txt
			fmt.Println(line,"\n")
			writer.WriteString(line + "\n")
		}
	}
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

func convertAirportCodes(text string, lookupFilePath string) string {
	//if word match with any regex find the replace word from csv column
	//replace the word with new replacement 
	icaoRegex := regexp.MustCompile(`^#([A-Z]{3})([^A-Za-z0-9]|$)`)
	iataRegex := regexp.MustCompile(`^##([A-Z]{4})([^A-Za-z0-9]|$)`)
	dateRegex := regexp.MustCompile(`^D\([0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}Z\)$`)
	// time12Regex := regexp.MustCompile(`^#`)
	// time24Regex := regexp.MustCompile(`^#`)
	
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
	}
	return strings.Join(words, " ")
}

func cleanDateFormat(word string) string {
	isValid:= true
	year := word[2:6]
	dayStr:= ""
	shortMonthNames := []string{ "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec", }
	month,err1 := strconv.Atoi(word[7:9])
	if err1!= nil {
		isValid = false
	}
	day,err2 := strconv.Atoi(word[10:12])
	if err2!= nil {
		isValid = false
	}
	if month >12 {
		isValid = false
	}
	if day > 32 {
		isValid = false
	}

	monthName := shortMonthNames[month-1]
	if day < 10 {
		dayStr = fmt.Sprintf("%02d", day)
	} else {
		dayStr = strconv.Itoa(day)
	}

	if !isValid {
		return word
	}
	// fmt.Println("year:", year," month:", monthName," day:",dayStr, "\n")
	date:= dayStr+" "+ monthName +" "+ year
	return (date)
}

func cleanAirportCode(word string, prefix string, length int, lookupFilePath string) string {
	word = strings.TrimPrefix(word, prefix)
	code := word[:length]

	// find airport name from lookupp file
	airportName := GetNameFromAirportLookup(code, lookupFilePath)
	//in case the code is not found and we keep the lookup code s it is in putput file
	if code == airportName {
		return ("##"+code+word[length:])
	}
	return airportName+word[length:]
}

func GetNameFromAirportLookup(code string, lookupFilePath string) string {
	airportLookup, err := os.Open(lookupFilePath)
	if err != nil {
		fmt.Println("Airport lookup malformed\n")
			return code
	}
	//defer to close file and prevent resource leaks, no need to do f.Close() later
	defer airportLookup.Close()

	//csv.NewReader to read csv files
	airportLookupReader := csv.NewReader(airportLookup)

	records, err1 := airportLookupReader.ReadAll()
	if err1 != nil {
		fmt.Println("Airport lookup malformed\n")
			return code
	}

	//loop every records, make a array details with values in each record
	for i:=1; i< len(records)-1; i++ {
		for _, details := range records{
	// details := strings.Split(records[i], ",")
			// compare with each value with current record
			for _, data := range details {
				if data == code {
					return details[0]
				}
			}
		}
		
	}
	return code
}

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
