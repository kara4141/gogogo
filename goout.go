package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/* turn to one struct:
type SitemapIndex struct {
	Locations []Location `xml:"sitemap"` //slice of locations, not fixed size

}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}
*/

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"` //slice of locations, not fixed size
}
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}
type NewsMap struct {
	Keyword  string
	Location string
}

func main() {
	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	resp.Body.Close()

	xml.Unmarshal(bytes, &s)
	//fmt.Printf("Here %s some %s", "are" "vatiables")

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Keywords {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	for idx, data := range news_map {
		fmt.Println("\n\n\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
	//fmt.Println(s.Locations)

}
