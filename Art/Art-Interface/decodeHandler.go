package main
import (
	"html/template"
	"net/http"
			)
type templateData struct {
	Code string
	Art string
	Error string 
}
func decodeHandler (w http.ResponseWriter, req *http.Request) {
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
		artdesign, err:= decoder(inputArtCode)

		if err !=nil {
			w.WriteHeader(http.StatusBadRequest)
		// http.Error(w, err.Error(), http.StatusBadRequest)
		data.Code = inputArtCode
		data.Art = "" 
		data.Error = "Error:" + err.Error() 
		
	}
		// set data for html page
		data.Code = inputArtCode
		data.Art = artdesign 

	} 
		tpl.Execute(w, data)
	
}