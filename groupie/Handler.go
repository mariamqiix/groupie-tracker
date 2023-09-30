package groupie

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	// "strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/style.css" {
		ServeFile(w, r, "style.css")
		return
	} else if r.URL.Path == "/index.html" {
		indexTemplate, _ := template.ParseFiles("./template/index.html")
		artists, _, _ := UnmarshalData()

		pageData := PageData{
			All: artists,
		}

		err := indexTemplate.Execute(w, pageData)

		if err != nil {
			fmt.Print(err)
			ServeFile(w, r, "500.html")
		}

		return

	} else if r.URL.Path == "/home.html" || r.URL.Path == "/" {
		ServeFile(w, r, "home.html")
		return
	} else if r.URL.Path == "/aboutus.html" {
		ServeFile(w, r, "aboutus.html")
		return
	} else if r.URL.Path == "/submit" {
		valueStr := r.URL.Query().Get("value")
		value, _ := strconv.Atoi(valueStr)
		indexTemplate, _ := template.ParseFiles("./template/artice.html")
		artists, _, _ := UnmarshalData()

		if valueStr == "" || value > 52 {
			ServeFile(w, r, "400.html")
			return
		}

		pageDataArtice := PageDataArtice{
			All:                    artists[value-1],
			MergeDatesAndLocations: MergeDatesAndLocations(value - 1),
			// AllLocation:  TheLocations.Index[value-1].TheData,
			// AllDates:     TheDates.Index2[value-1].TheData,
			// Allrelations: TheRelations.Index3[value-1].TheData,
		}

		err := indexTemplate.Execute(w, pageDataArtice)

		if err != nil {
			fmt.Print(err)
			ServeFile(w, r, "500.html")
		}

		return
	} else {
		w.WriteHeader(http.StatusNotFound)
		ServeFile(w, r, "404.html")
		return
	}

}

func ServeFile(w http.ResponseWriter, r *http.Request, page string) {
	http.ServeFile(w, r, "./template/"+page)
}
