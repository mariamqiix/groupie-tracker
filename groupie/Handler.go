package groupie

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	// "strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	artists, _, _, flag := UnmarshalData()
	if !flag {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "./template/500.html")
		return

	} else if r.URL.Path == "/style.css" {
		http.ServeFile(w, r, "./template/style.css")
		return

	} else if r.URL.Path == "/index.html" {
		indexTemplate, _ := template.ParseFiles("./template/index.html")

		pageData := PageData{
			All: artists,
		}

		err := indexTemplate.Execute(w, pageData)
		if err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "./template/500.html")
		}

		return

	} else if r.URL.Path == "/home.html" || r.URL.Path == "/" {
		http.ServeFile(w, r, "./template/home.html")
		return

	} else if r.URL.Path == "/aboutus.html" {
		http.ServeFile(w, r, "./template/aboutus.html")
		return

	} else if r.URL.Path == "/submit" {
		valueStr := r.URL.Query().Get("value")
		value, _ := strconv.Atoi(valueStr)
		indexTemplate, _ := template.ParseFiles("./template/artice.html")

		if valueStr == "" || value > 52 {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "./template/400.html")
			return
		}

		pageDataArtice := PageDataArtice{
			All:                    artists[value-1],
			MergeDatesAndLocations: MergeDatesAndLocations(value - 1),
		}

		err := indexTemplate.Execute(w, pageDataArtice)
		if err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "./template/500.html")
		}

		return
	} else {

		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "./template/404.html")
		return
	}

}
