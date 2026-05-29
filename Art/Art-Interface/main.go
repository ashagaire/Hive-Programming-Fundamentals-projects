package main
import (
	"fmt"
	"net/http"
			)

//to render this variable in html it should be capitalcase



func main() {
	//SERVER
	// setting files access from local folder named "static"
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	//register  the handaler function from http for root URL path "/"
	http.HandleFunc("/", encodeHandler)

	fmt.Println("Server is running at http://localhost:8080 ...")

	//start the web server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}



