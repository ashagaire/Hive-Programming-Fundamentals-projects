package main
import (
	"html/template"
	"net/http"
			)

func encodeHandler (w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("interface.html")
	if err !=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if req.Method == http.MethodPost {
		req.ParseForm()
		inputArtCode := req.FormValue("artcode")
		

		//run the action functions
		art := singlelineDecode(inputArtCode)

		// set data for html page
		data := art 
		
		if err !=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
			//render/execute the template by merging the HTML and data
		tpl.Execute(w, data)
	} else {
		tpl.Execute(w, nil)
	}

}