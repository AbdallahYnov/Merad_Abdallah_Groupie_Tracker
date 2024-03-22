package routeur

import (
	"fmt"
	"net/http"
	"rickandmortyapi/utility"
	"strconv"
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
	http.HandleFunc("/error", errorHandler)

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
	tagstrings := r.FormValue("tag")
	fmt.Println(tagstrings)
	fmt.Println(r.FormValue("tag"))
	fmt.Println(r.URL)

	// Define the link for the API request
	var ListPerso []utility.ResultCharacter
	link := "https://rickandmortyapi.com/api/character"
	for {
		resPerso, res := utility.CharacterList(link)
		ListPerso = append(ListPerso, resPerso...)
		if res.Info.Next == "" {
			break
		}
		link = res.Info.Next
	}

	// If tags were selected, filter characters based on the selected tags
	if len(tagstrings) > 0 && tagstrings != "" {
		ListPerso = utility.FilterByTag(ListPerso, tagstrings)
	}

	page := r.FormValue("page")
	currentPage, errPage := strconv.Atoi(page)
	if page == "" || errPage != nil || currentPage < 1 {
		currentPage = 1
	}
	fmt.Println(currentPage)
	if len(ListPerso) < (currentPage * 10) {
		currentPage = 1 //remplacer par redirec page 404
	}
	ListPerso = ListPerso[(currentPage*10)-10 : (currentPage * 10)]

	// Parse the template and execute it with the character data
	tmpl, err := template.New("characters").ParseFiles("templates/characters.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	dataS := utility.CombinedDataChar{
		Navigation: struct {
			PagePrev string
			PageNext string
		}{
			PageNext: fmt.Sprintf("/characters?page=%v&tag=%v", currentPage+1, tagstrings),
			PagePrev: fmt.Sprintf("/characters?page=%vtag=%v", currentPage-1, tagstrings),
		},
		Data: ListPerso,
	}
	tmpl.ExecuteTemplate(w, "characters", dataS)
}

// locationHandler handles requests to the "/location" endpoint
func locationHandler(w http.ResponseWriter, r *http.Request) {

	// Get the selected tags from the form data
	tagstrings := r.FormValue("tag")
	fmt.Println(tagstrings)
	fmt.Println(r.FormValue("tag"))
	fmt.Println(r.URL)

	// Define the link for the API request
	var ListLoc []utility.ResultLocation
	link := "https://rickandmortyapi.com/api/location"
	for {
		resLoc, res := utility.LocationList(link)
		ListLoc = append(ListLoc, resLoc...)
		if res.Info.Next == "" {
			break
		}
		link = res.Info.Next
	}

	page := r.FormValue("page")
	currentPage, errPage := strconv.Atoi(page)
	if page == "" || errPage != nil || currentPage < 1 {
		currentPage = 1
	}
	fmt.Println(currentPage)
	if len(ListLoc) < (currentPage * 10) {
		currentPage = 1 //remplacer par redirec page 404
	}
	ListLoc = ListLoc[(currentPage*10)-10 : (currentPage * 10)]

	// Parse the template and execute it with the location data
	tmpl, err := template.New("locations").ParseFiles("templates/locations.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	dataS := utility.CombinedDataLoc{
		Navigation: struct {
			PagePrev string
			PageNext string
		}{
			PageNext: fmt.Sprintf("/locations?page=%v", currentPage+1),
			PagePrev: fmt.Sprintf("/locations?page=%v", currentPage-1),
		},
		Data: ListLoc,
	}
	// tmpl.ExecuteTemplate(w, "locations", dataS)
	// // Placeholder for location handler logic
	// var link string
	// page := r.URL.Query().Get("page")
	// if page == "" {
	// 	link = "https://rickandmortyapi.com/api/location"

	// } else {
	// 	link = page
	// }
	// data, info := utility.LocationList(link)

	// tmpl, err := template.New("locations").ParseFiles("templates/locations.html")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// dataS := utility.CombinedDataLoc{
	// 	Response: info,
	// 	Data:     data,
	// }
	tmpl.ExecuteTemplate(w, "locations", dataS)
}

// episodeHandler handles requests to the "/episode" endpoint
func episodeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the selected tags from the form data
	tagstrings := r.FormValue("tag")
	fmt.Println(tagstrings)
	fmt.Println(r.FormValue("tag"))
	fmt.Println(r.URL)

	// Define the link for the API request
	var ListEp []utility.ResultEpisode
	link := "https://rickandmortyapi.com/api/episode"
	for {
		resEp, res := utility.EpisodeList(link)
		ListEp = append(ListEp, resEp...)
		if res.Info.Next == "" {
			break
		}
		link = res.Info.Next
	}

	page := r.FormValue("page")
	currentPage, errPage := strconv.Atoi(page)
	if page == "" || errPage != nil || currentPage < 1 {
		currentPage = 1
	}
	fmt.Println(currentPage)
	if len(ListEp) < (currentPage * 10) {
		currentPage = 1 //remplacer par redirec page 404
	}
	ListEp = ListEp[(currentPage*10)-10 : (currentPage * 10)]

	// Parse the template and execute it with the episode data
	tmpl, err := template.New("episodes").ParseFiles("templates/episodes.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	dataS := utility.CombinedDataEp{
		Navigation: struct {
			PagePrev string
			PageNext string
		}{
			PageNext: fmt.Sprintf("/episodes?page=%v", currentPage+1),
			PagePrev: fmt.Sprintf("/episodes?page=%v", currentPage-1),
		},
		Data: ListEp,
	}
	tmpl.ExecuteTemplate(w, "episodes", dataS)
}

// favoritesHanfler handles requests to the "/favorites" endpoint
func favoritesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("favorites").ParseFiles("templates/favorites.html")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.ExecuteTemplate(w, "favorites", nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("query")
	fmt.Println(query)

	tmpl, err := template.New("search").ParseFiles("templates/search.html")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.ExecuteTemplate(w, "search", nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("error").ParseFiles("templates/error.html")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.ExecuteTemplate(w, "error", nil)
}
