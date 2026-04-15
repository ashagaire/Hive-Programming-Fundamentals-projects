package main
import ( "fmt"
"os") 

func main(){

	if len(os.Args) !=4 || os.Args[1] == "-h" {
		fmt.Println("itinerary usage:\n go run . ./input.txt ./output.txt ./airport-lookup.csv")
 		return
	}
	name := os.Args[1]

	fmt.Println("Welcome to Itinerary Formator App!!!\n")
	fmt.Println(name)


}