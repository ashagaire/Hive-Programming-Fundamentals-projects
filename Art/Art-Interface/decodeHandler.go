package main

import (
	"html/template"
	"net/http"
)

type templateData struct {
	Code     string
	Art      string
	Response string
	Error    string
}

func decodeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" && req.URL.Path != "/decoder" {
		http.Error(w, "Page not found", http.StatusMethodNotAllowed)
		return

	}
	//could be out of function
	tpl, err := template.ParseFiles("decoder.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := templateData{}
	if req.Method == http.MethodPost {
		req.ParseForm()
		inputArtCode := req.FormValue("artcode")

		artdesign, err := decoder(inputArtCode)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			data.Code = inputArtCode
			data.Art = ""
			data.Response = "400"
			data.Error = err.Error()
			
		} else {

			// set data for html page
			data.Code = inputArtCode
			data.Art = artdesign
			data.Response = "202"

		}

	}
	w.WriteHeader(http.StatusAccepted)
	err = tpl.Execute(w, data)
	if err != nil {
		tpl.Execute(w, templateData{Code: "500", Art:"", Response:"shaisse", Error:"Somethng wwetn t wrnng"})
		return
	}

	
}
