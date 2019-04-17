package trawlingyoutubego

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

// TrwYtArticle API Structure
type TrwYtArticle struct {
	ID                   string          `json:"id"`
	URL                  string          `json:"url"`
	URLShort             string          `json:"urlShort"`
	URLEmbed             string          `json:"urlEmbed"`
	Published            time.Time       `json:"published"`
	ChannelID            string          `json:"channelId"`
	ChannelTitle         string          `json:"channelTitle"`
	Title                string          `json:"title"`
	Description          string          `json:"description"`
	Thumbnails           TrwYtThumbnails `json:"thumbnails"`
	Tags                 []string        `json:"tags"`
	CategoryID           string          `json:"categoryId"`
	LiveBroadcastContent string          `json:"liveBroadcastContent"`
	ViewCount            string          `json:"viewCount"`
	LikeCount            string          `json:"likeCount"`
	DislikeCount         string          `json:"dislikeCount"`
	FavoriteCount        string          `json:"favoriteCount"`
	CommentCount         string          `json:"commentCount"`
}

// TrwYtResponse API structure
type TrwYtResponse struct {
	Response struct {
		Data         []TrwYtArticle `json:"data"`
		RequestLeft  int            `json:"requestLeft"`
		TotalResults int            `json:"totalResults"`
		Next         string         `json:"next"`
	} `json:"response"`
}

// TrwYtRequest API structure
type TrwYtRequest struct {
	Token string
	Query string
	Ts    string
	Tsi   string
}

// TrwYtThumbnails Thubnails structure
type TrwYtThumbnails struct {
	Default  TrwYtThumbnail `json:"default"`
	Medium   TrwYtThumbnail `json:"width"`
	High     TrwYtThumbnail `json:"high"`
	Standard TrwYtThumbnail `json:"standard"`
	Maxres   TrwYtThumbnail `json:"maxres"`
}

// TrwYtThumbnail Thubnail structure
type TrwYtThumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// TrwYtError for get problems
type TrwYtError struct {
	Response struct {
		Error string `json:"error"`
	} `json:"response"`
}

// Request to https service
func Request(url string) (TrwYtResponse, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	var res TrwYtResponse
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, err
	}

	req.Header.Set("User-Agent", "trawlingyt-cli.go 1.1")
	resp, err2 := client.Do(req)
	if err2 != nil {
		return res, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return res, fmt.Errorf("http status code: %d", resp.StatusCode)
	}

	err3 := json.NewDecoder(resp.Body).Decode(&res)
	return res, err3
}

// Query Initial function
func Query(params TrwYtRequest) (TrwYtResponse, error) {
	values := reflect.ValueOf(params)
	twurl := "https://youtube.trawlingweb.com/search/?"
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).String() != "" {
			if i != 0 {
				twurl += "&"
			}
			if values.Type().Field(i).Name == "Query" {
				sturl := values.Field(i).String()
				encodedPath := url.QueryEscape(sturl)
				twurl += "q=" + encodedPath
			} else {
				twurl += strings.ToLower(values.Type().Field(i).Name) + "=" + values.Field(i).String()
			}
		}
	}
	return Request(twurl)
}

// Next query function
func Next(twurl string) (TrwYtResponse, error) {
	return Request(twurl)
}
