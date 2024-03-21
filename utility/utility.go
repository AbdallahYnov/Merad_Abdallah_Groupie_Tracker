package utility

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	Favorite bool      `json:"favorite"`
}
type ResultLocation struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Dimension string    `json:"dimension"`
	Residents []string  `json:"residents"`
	URL       string    `json:"url"`
	Created   time.Time `json:"created"`
	Favorite  bool      `json:"favorite"`
}
type ResultEpisode struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	AirDate    string    `json:"air_date"`
	Episode    string    `json:"episode"`
	Characters []string  `json:"characters"`
	URL        string    `json:"url"`
	Created    time.Time `json:"created"`
	Favorite   bool      `json:"favorite"`
}
type AllResults struct {
	Characters []ResultCharacter `json:"results"`
	Locations  []ResultLocation  `json:"results"`
	Episodes   []ResultEpisode   `json:"results"`
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
	Response ResponseCharacter
	Data     []ResultCharacter
}
type CombinedDataLoc struct {
	Response ResponseLocation
	Data     []ResultLocation
}
type CombinedDataEp struct {
	Response ResponseEpisode
	Data     []ResultEpisode
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

func Filterbytag(results []ResultCharacter, tags []string) []ResultCharacter {
	var trier []ResultCharacter
	addedChar := make(map[string]bool)

	for _, result := range results {
		for _, tag := range tags {
			if result.Gender == tag && addedChar[result.Name] {
				trier = append(trier, result)
				addedChar[result.Gender] = true
			}
			if result.Species == tag && addedChar[result.Name] {
				trier = append(trier, result)
				addedChar[result.Species] = true
			}
			if result.Status == tag && addedChar[result.Name] {
				trier = append(trier, result)
				addedChar[result.Status] = true
			}
		}
	}
	return trier
}
