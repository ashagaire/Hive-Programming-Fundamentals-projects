package main
import (
	"fmt"
	"html/template"
	// "log"
	"net/http"
			)

//to render this variable in html it should be capitalcase
var art int 
func calculateSomething(a,b int) int {
	return a*b
}

func encodeHandler (w http.ResponseWriter, req *http.Request) {
	//run the action functions
	art = calculateSomething(1,2)

	// set data for html page
	data := art 

	// Parse the html file
	tpl, err := template.ParseFiles("interface.html")
	if err !=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//render/execute the template by merging the HTML and data
	tpl.Execute(w, data)

}
func main() {
	//SERVER: register  the handaler function from http for root URL path "/"
	http.HandleFunc("/", encodeHandler)

	fmt.Println("Server is running at http://localhost:8080 ...")

	//SERVER:start the web server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}



