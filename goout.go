package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"` //slice of locations, not fixed size

}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)
}
