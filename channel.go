package main

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}
