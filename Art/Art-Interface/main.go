package main
import (
	"fmt"
	"net/http"
	"log"
		)


func main() {
	//SERVER
	// setting files access from local folder named "static"
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	//register  the handaler function from http for root URL path "/"
	http.HandleFunc("/", decodeHandler)
	fmt.Println("Server is running at http://localhost:8080 ...")

	//start the web server on port 8080
	err := http.ListenAndServe(":8080", nil)
	log.Fatalln(err)
}



