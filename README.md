# ItineraryNoteFormatter
A command-line application that processes admin-generated flight itineraries and transforms them into clear, readable formats for customers.

## Source Code and Configuration Files

ItineraryNoteFormatter/
├── main.go            → main program logic
├── go.mod             → Go module configuration
├── input.txt          → sample input file
├── output.txt         → generated output file
└── airport-lookup.csv → airport codes database

## Project Overview 

"Anywhere Holidays" is a brand new online travel agent, finding cheap holidays for their customers. When this company books the flight for customers their system generates the itinerary which is formatted for administrators, and the information is not customer-friendly.

This command line tool, reads the text-based itinerary from a file named input.txt, processes the text to make it customer-friendly, and writes the result to a new file output.txt . The processing includes translating the airport code (based on airport lookup code data), date and time from originally generated text while booking flights for itinary.

## How to use
To run this tool for given original itinerary in input.txt file you will need
- input.txt file 
- [airport lookup], https://intra.hive.fi/api/file?gitpath=content/coding-fundamentals/coding-fundamentals-go/itinerary/prettifier/_config/resources/airport-lookup.csv
- file name for output file (output.txt) 

## Setup and Installation
Installation 
1.Clone the repository:
```bash  
git clone 
cd 
```
2.Initialize the Go module:
```bash  
go mod init ItineraryNoteFormatter
```

3.Verify Go is installed:
```bash
go version
```

### Usage Guide
```bash
go run . ./input.txt ./output.txt ./airport-lookup.csv
```

## Arguments
- input.txt the path to admin generated itinerary file
- output.txt path with output file name where custumer friendly itinerary will be saved
- airort-look-up path to the airport codes lookup file

## Help Flag
```bash
go run -h
```

## Airport Lookup CSV Format 
The airport lookup file must be a CSV with exactly 6 columns:
```bash
name, iso_country, municipality, icao_code, iata_code, coordinates
```
## Error Handaling

| Situation  | Output |
|----------|----------|
| Wrong number of arguments  | Usage instructions  | 
| Input fie not found   | Input not found  | 
| Airport lookup not found   | Airport lookup not found  | 
| Airport lookup malformed  | Airport lookup malformed  | 
In all error cases the output file is never created or overwritten.

## Features and Bonus Functionality

### Airport Code Replacement

Converts IATA and ICAO airport codes into full airport names.

IATA codes are preceded by a single #:
```bash
#LAX → Los Angeles International Airport
```

ICAO codes are preceded by ##:
```bash
##EGLL → London Heathrow Airport
```
If an airport code is not found in the lookup file it is left unchanged.

## Date and Time Formatting

Converts ISO 8601 dates and times into customer friendly formats.
Input	Output
D(2022-05-09T08:07Z)	09 May 2022
T12(2022-05-09T08:07Z)	08:07AM (+00:00)
T24(2022-05-09T14:30Z)	14:30 (+00:00)

Zulu time Z is displayed as (+00:00). Malformed dates are left unchanged.

## Whitespace Trimming
- Converts \v, \f, \r characters to newlines
- Removes trailing spaces from each line
- Reduces consecutive blank lines to a maximum of one

<br/>

## Colored Output (Bonus)
- Welcome Message displayed on yellow in the terminal
- Usage instructions displayed in blue 
- Updated itinerary text is also displayed on terminal along with output.txt file

## City Names (Bonus)
Along with airport name on the updated itinerary file there are city names also.

## Info
You can find test itinerary text with raw airport codes and date and time in ISO 8601 format. Output fie has the updated custumer friendly itinerary details.