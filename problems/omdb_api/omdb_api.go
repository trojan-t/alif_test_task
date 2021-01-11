package omdb_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	baseURL       = "http://www.omdbapi.com/?"
	plot          = "full"
	MovieSearch   = "movie"
	SeriesSearch  = "series"
	EpisodeSearch = "episode"
)

// You know what it is :)
type OmdbAPI struct {
	apiKey string
}

// Init is your apiKey initialization function.
// You can get your apiKey from omdbapi.com
// I hope, that you have it :)
// P.S. If you don't have, chat me. I'll give you.
func Init(apiKey string) *OmdbAPI {
	return &OmdbAPI{apiKey: apiKey}
}

type QueryData struct {
	Title      string
	Year       string
	ImdbId     string
	SearchType string
}

type SearchResult struct {
	Title  string
	Year   string
	ImdbID string
	Type   string
}

type SearchResponse struct {
	Search       []SearchResult
	Response     string
	Error        string
	totalResults int
}

type MovieResult struct {
	Title    string
	Year     string
	ImdbID   string
	Type     string
	Response string
	Error    string
}

func (api *OmdbAPI) Search(query *QueryData) (*SearchResponse, error) {
	resp, err := api.requestAPI("search", query.Title, query.Year, query.SearchType)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := new(SearchResponse)
	err = json.NewDecoder(resp.Body).Decode(r)

	if err != nil {
		return nil, err
	}
	if r.Response == "False" {
		return r, errors.New(r.Error)
	}

	return r, nil
}

func (api *OmdbAPI) MovieByTitle(query *QueryData) (*MovieResult, error) {
	resp, err := api.requestAPI("title", query.Title, query.Year, query.SearchType)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := new(MovieResult)
	err = json.NewDecoder(resp.Body).Decode(r)

	if err != nil {
		return nil, err
	}
	if r.Response == "False" {
		return r, errors.New(r.Error)
	}
	return r, nil
}

//MovieByImdbID returns a MovieResult given a ImdbID ex:"tt2015381"
func (api *OmdbAPI) MovieByImdbID(id string) (*MovieResult, error) {
	resp, err := api.requestAPI("id", id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := new(MovieResult)
	err = json.NewDecoder(resp.Body).Decode(r)

	if err != nil {
		return nil, err
	}
	if r.Response == "False" {
		return r, errors.New(r.Error)
	}
	return r, nil
}

func (api *OmdbAPI) requestAPI(apiCategory string, params ...string) (resp *http.Response, err error) {
	var URL *url.URL
	URL, err = url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	if len(params) > 1 && params[2] != "" {
		if params[2] != MovieSearch &&
			params[2] != SeriesSearch &&
			params[2] != EpisodeSearch {
			return nil, errors.New("Invalid search category - " + params[2])
		}
	}
	URL.Path += "/"
	parameters := url.Values{}
	parameters.Add("apikey", api.apiKey)

	switch apiCategory {
	case "search":
		parameters.Add("s", params[0])
		parameters.Add("y", params[1])
		parameters.Add("type", params[2])
	case "title":
		parameters.Add("t", params[0])
		parameters.Add("y", params[1])
		parameters.Add("type", params[2])
		parameters.Add("plot", plot)
	case "id":
		parameters.Add("i", params[0])
		parameters.Add("plot", plot)
	}

	URL.RawQuery = parameters.Encode()
	res, err := http.Get(URL.String())
	err = checkErr(res.StatusCode)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func checkErr(status int) error {
	if status != 200 {
		return fmt.Errorf("Status Code %d received from IMDB", status)
	}
	return nil
}

func (sr SearchResult) String() string {
	return fmt.Sprintf("Title: %s, Year: (%s), ImdbID: #%s, Type: %s", sr.Title, sr.Year, sr.ImdbID, sr.Type)
}
