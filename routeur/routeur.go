package routeur

import (
	"fmt"
	"net/http"
	"rickandmortyapi/utility"
	"text/template"
)

func InitServer() {
	// Define HTTP routes and handlers
	http.HandleFunc("/home", indexHandler)
	http.HandleFunc("/characters", characterHandler)
	http.HandleFunc("/locations", locationHandler)
	http.HandleFunc("/episodes", episodeHandler)
	http.HandleFunc("/favorites", favoritesHandler)
	http.HandleFunc("/search", searchHandler)

	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on port :8080")
	http.ListenAndServe(":8080", nil)
}

// indexHandler handles requests to the root endpoint "/"
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index").ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.ExecuteTemplate(w, "index", nil)
}

// characterHandler handles requests to the "/character" endpoint
func characterHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the selected tags from the form data
	tagstrings := r.Form["tag[]"]

	// Define the link for the API request
	var link string
	page := r.URL.Query().Get("page")
	if page == "" {
		link = "https://rickandmortyapi.com/api/character"
	} else {
		link = page
	}

	// Fetch the list of characters from the API
	data, info := utility.CharacterList(link)

	// If tags were selected, filter characters based on the selected tags
	if len(tagstrings) > 0 {
		data = utility.FilterByTag(data, tagstrings)
	}

	// Parse the template and execute it with the character data
	tmpl, err := template.New("characters").ParseFiles("templates/characters.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	dataS := utility.CombinedDataChar{
		Response: info,
		Data:     data,
	}
	tmpl.ExecuteTemplate(w, "characters", dataS)
}

// locationHandler handles requests to the "/location" endpoint
func locationHandler(w http.ResponseWriter, r *http.Request) {
	// Placeholder for location handler logic
	var link string
	page := r.URL.Query().Get("page")
	if page == "" {
		link = "https://rickandmortyapi.com/api/location"

	} else {
		link = page
	}
	data, info := utility.LocationList(link)

	tmpl, err := template.New("locations").ParseFiles("templates/locations.html")
	if err != nil {
		fmt.Println(err)
	}
	dataS := utility.CombinedDataLoc{
		Response: info,
		Data:     data,
	}
	tmpl.ExecuteTemplate(w, "locations", dataS)
}

// episodeHandler handles requests to the "/episode" endpoint
func episodeHandler(w http.ResponseWriter, r *http.Request) {
	// Placeholder for location handler logic
	var link string
	page := r.URL.Query().Get("page")
	if page == "" {
		link = "https://rickandmortyapi.com/api/episode"

	} else {
		link = page
	}
	data, info := utility.EpisodeList(link)

	tmpl, err := template.New("episodes").ParseFiles("templates/episodes.html")
	if err != nil {
		fmt.Println(err)
	}
	dataS := utility.CombinedDataEp{
		Response: info,
		Data:     data,
	}
	tmpl.ExecuteTemplate(w, "episodes", dataS)
}

// favoritesHanfler handles requests to the "/favorites" endpoint
func favoritesHandler(w http.ResponseWriter, r *http.Request) {
	// Placeholder for favorites handler logic
}

func searchHandler(w http.ResponseWriter, r *http.Request) {

}
