package utility

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Info struct {
	Count int    `json:"count"`
	Pages int    `json:"pages"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

type Origin struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ResultCharacter struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
	Species  string    `json:"species"`
	Type     string    `json:"type"`
	Gender   string    `json:"gender"`
	Origin   Origin    `json:"origin"`
	Location Location  `json:"location"`
	Image    string    `json:"image"`
	Episode  []string  `json:"episode"`
	URL      string    `json:"url"`
	Created  time.Time `json:"created"`
	Favorite bool
}
type ResultLocation struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Dimension string    `json:"dimension"`
	Residents []string  `json:"residents"`
	URL       string    `json:"url"`
	Created   time.Time `json:"created"`
	Favorite  bool
}
type ResultEpisode struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	AirDate    string    `json:"air_date"`
	Episode    string    `json:"episode"`
	Characters []string  `json:"characters"`
	URL        string    `json:"url"`
	Created    time.Time `json:"created"`
	Favorite   bool
}
type AllResults struct {
	Characters []ResultCharacter `json:"resultschar"`
	Locations  []ResultLocation  `json:"resultsloc"`
	Episodes   []ResultEpisode   `json:"resultsep"`
}

type ResponseCharacter struct {
	Info    Info              `json:"info"`
	Results []ResultCharacter `json:"results"`
}
type ResponseLocation struct {
	Info    Info             `json:"info"`
	Results []ResultLocation `json:"results"`
}
type ResponseEpisode struct {
	Info    Info            `json:"info"`
	Results []ResultEpisode `json:"results"`
}

type CombinedDataChar struct {
	Navigation struct {
		PagePrev string
		PageNext string
	}
	Data []ResultCharacter
}
type CombinedDataLoc struct {
	Navigation struct {
		PagePrev string
		PageNext string
	}
	Data []ResultLocation
}
type CombinedDataEp struct {
	Navigation struct {
		PagePrev string
		PageNext string
	}
	Data []ResultEpisode
}

func CharacterList(link string) ([]ResultCharacter, ResponseCharacter) {
	// Make HTTP request to get the list of characters
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, ResponseCharacter{}
	}
	defer resp.Body.Close()
	var results ResponseCharacter
	// Decode the JSON response into a slice of characters
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, ResponseCharacter{}
	}

	return results.Results, results
}

func LocationList(link string) ([]ResultLocation, ResponseLocation) {
	// Make HTTP request to get the list of locations
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, ResponseLocation{}
	}
	defer resp.Body.Close()
	var results ResponseLocation
	// Decode the JSON response into a slice of locations
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, ResponseLocation{}
	}

	return results.Results, results
}

func EpisodeList(link string) ([]ResultEpisode, ResponseEpisode) {
	// Make HTTP request to get the list of episodes
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, ResponseEpisode{}
	}
	defer resp.Body.Close()
	var results ResponseEpisode
	// Decode the JSON response into a slice of episodes
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, ResponseEpisode{}
	}

	return results.Results, results
}

func FilterByTag(results []ResultCharacter, tag string) []ResultCharacter {
	filter := make([]ResultCharacter, 0)
	//addedChar := make(map[string]bool)

	for _, result := range results {
		// Check if the character matches any of the tags
		if result.Gender == tag || result.Species == tag || result.Status == tag {
			filter = append(filter, result)
			/* // Check if the character has not been added already
			if !addedChar[result.Name] {
				// Add the character to the filter slice
				filter = append(filter, result)
				// Mark the character as added
				addedChar[result.Name] = true
				// Break the inner loop since the character has been added
				break
			} */
		}
	}

	return filter
}
func FilterByTog(results []ResultLocation, tog string) []ResultLocation {
	filter := make([]ResultLocation, 0)
	for _, result := range results {
		if result.Type == tog {
			filter = append(filter, result)
		}
	}
	return filter
}

func SearchCharacters(query string) ([]ResultCharacter, error) {
	// Fetch the list of characters
	characters, _ := CharacterList("https://rickandmortyapi.com/api/character")

	// Perform search based on the query
	var results []ResultCharacter
	for _, character := range characters {
		if strings.Contains(strings.ToLower(character.Name), strings.ToLower(query)) {
			results = append(results, character)
		}
	}

	return results, nil
}

// Function to search locations based on the provided query
func SearchLocations(query string) ([]ResultLocation, error) {
	// Fetch the list of locations
	locations, _ := LocationList("https://rickandmortyapi.com/api/location")

	// Perform search based on the query
	var results []ResultLocation
	for _, location := range locations {
		if strings.Contains(strings.ToLower(location.Name), strings.ToLower(query)) {
			results = append(results, location)
		}
	}

	return results, nil
}

// Function to search episodes based on the provided query
func SearchEpisodes(query string) ([]ResultEpisode, error) {
	// Fetch the list of episodes
	episodes, _ := EpisodeList("https://rickandmortyapi.com/api/episode")

	// Perform search based on the query
	var results []ResultEpisode
	for _, episode := range episodes {
		if strings.Contains(strings.ToLower(episode.Name), strings.ToLower(query)) {
			results = append(results, episode)
		}
	}

	return results, nil
}

// Search function to search through characters, locations, and episodes
func Search() {
}
