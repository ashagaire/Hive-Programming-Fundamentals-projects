package main
import ( "fmt"
"encoding/csv"
"reflect"
		"os") 

func main(){

	if len(os.Args) !=4 || os.Args[1] == "-h" {
		fmt.Println("itinerary usage:\n go run . ./input.txt ./output.txt ./airport-lookup.csv")
 		return
	}

	inputFilePath := os.Args[1]
	// outputFilePath := os.Args[2]
	lookupFilePath := os.Args[3]

	err := ScanInputs(inputFilePath, lookupFilePath)
	if err != nil {
		fmt.Println(err)
		 
		 return 
	}
	records := ScanAirportLookup(lookupFilePath)
	fmt.Println(records)

	fmt.Println("Welcome to Itinerary Formator App!!!\n")

}

//handeling error with error type
func ScanInputs(inputFile string, lookupFile string) error {

	//os.Stat() retrives information about the file or directory
	_, err1 := os.Stat(inputFile)
	if err1 != nil {
		return nil
		// return fmt.Errorf("Input not found")
	}

	_, err2 := os.Stat(lookupFile)
	if err2 != nil {
		return fmt.Errorf("Airport lookup not found")
	}

	return nil

}

func ScanAirportLookup(lookupFile string) string {
	// in csv file if any column in the lookup is corrupted, missing or blank display "Airport lookup malformed"

	indexes := []string{"name", "iso_country","municipality", "icao_code", "iata_code"}

	airportLookup, err := os.Open(lookupFile)
	if err != nil {
		
		return "Airport lookup malformed"
	}
	//defer to close file and prevent resource leaks, no need to do f.Close() later
	defer airportLookup.Close()

	airportLookupReader := csv.NewReader(airportLookup)
	records, err1 := airportLookupReader.ReadAll()
	if err1 != nil {
		
		return "Airport lookup malformed"
	}

	checkColumns := CheckIndexRow(indexes,records[0])

	if checkColumns == false {
		
		return "Airport lookup malformed"
	}
	

	return "here are records"

}

func CheckIndexRow(indexes []string, firstRow []string) bool{
	return reflect.DeepEqual(indexes, firstRow)
}


//TODO functions

// for airport code in input.txt find row with that code in respetiive code column in csv file, if not found keep as it is
// if D(2007-04-05T12:30−02:00) found in input.txt output date in DD-Mmm-YYYY format
// if T12(2007-04-05T12:30−02:00)  found in input.txt output 12:30PM (-02:00)
// if T24(2007-04-05T12:30−02:00) found in input.txt output 12:30 (-02:00)
//  if there is "Z" after "T" in T12:30−02:00 then  show (+00:00)
// if the line after a blank line is also blank then remove it ans break the reading loop
