package main
import (
	"html/template"
	"net/http"
			)
type templateData struct {
	code string
	art string
}
func encodeHandler (w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("interface.html")
	if err !=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := templateData{}
	if req.Method == http.MethodPost {
		req.ParseForm()
		inputArtCode := req.FormValue("artcode")
		

		//run the action functions
		artdesign, err:= singlelineDecode(inputArtCode)

		if err !=nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
		// set data for html page
		data.code = inputArtCode
		data.art = artdesign 

	} 
		tpl.Execute(w, data)
	
}