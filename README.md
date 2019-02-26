# trawlingyoutube
Official Trawlingweb Youtube client for Go Language

[https://trawlingweb.com](https://trawlingweb.com)


### Example
```go
package main

import (
	"fmt"

	trwyt "github.com/anpro21/trawlingyoutube"
)

func main() {
	request := trwyt.TrwYtResponse{Token: "ea58ad77426816b16f2cd3c950de07886bc64472", Query: "debian AND redhat", Ts: "1533459254944"}
	ret, err := trwyt.Query(request)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Articles: ", len(ret.Response.Data))
		for _, article := range ret.Response.Data {
			fmt.Println(article.URL)
			fmt.Println("---------------------")
		}
	}

	urlnext := ret.Response.Next
	for urlnext != "" {
		ret, err = trwyt.Next(urlnext)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Articles: ", len(ret.Response.Data))
		fmt.Println("Rest results: ", ret.Response.RestResults)
		fmt.Println("---------------------")
		urlnext = ret.Response.Next
	}

	fmt.Println("End Query")

}
```

### Links:
[Register](https://dashboard.trawlingweb.com/register)

[Control Panel](https://dashboard.trawlingweb.com)

[API documentation](https://dashboard.trawlingweb.com/dashboard)

[Github Repository](https://github.com/anpro21)

### Functions:
```
Query(params TrwRequest)
Next(nexturl string)
```

### Input Params:
* Token: string with service key
* Query: String with trawling query
* Ts: unixtimestamp in ms
* Tsi: unixtimestamp in ms
* Sort: "published"/"crawled"
* Order: "asc"/"desc"


### Input Params struct:
```go
type TrwYtRequest struct {
	Token string
	Query string
	Ts    string
	Tsi   string
}
```


### Return struct:
```go
type TrwYtResponse struct {
	Response struct {
		Data         []TrwYtArticle `json:"data"`
		RequestLeft  int            `json:"requestLeft"`
		TotalResults int            `json:"totalResults"`
		Next         string         `json:"next"`
	} `json:"response"`
}

type TrwYtArticle struct {
	ID                   string          `json:"id"`
	URL                  string          `json:"url"`
	URLShort             string          `json:"urlShort"`
	URLEmbed             string          `json:"urlEmbed"`
	Published            time.Time       `json:"published"`
	ChannelID            string          `json:"channelId"`
	ChannelTitle         string          `json:"channelTitle"`
	Description          string          `json:"description"`
	Thumbnails           TrwYtThumbnails `json:"thumbnails"`
	Tags                 []string        `json:"tags"`
	CategoryID           string          `json:"categoryId"`
	LiveBroadcastContent string          `json:"liveBroadcastContent"`
	ViewCount            string          `json:"viewCount"`
	LikeCont             string          `json:"likeCont"`
	DislikeCount         string          `json:"dislikeCount"`
	FavoriteCount        string          `json:"favoriteCount"`
	CommentCount         string          `json:"commentCount"`
}

type TrwYtThumbnails struct {
	Default  TrwYtThumbnail `json:"default"`
	Medium   TrwYtThumbnail `json:"width"`
	High     TrwYtThumbnail `json:"high"`
	Standard TrwYtThumbnail `json:"standard"`
	Maxres   TrwYtThumbnail `json:"maxres"`
}

type TrwYtThumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

```


### MIT License

Copyright (c) 2018 Anpro21

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
