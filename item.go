package main

type Item struct {
	Title     string    `xml:"title"`
	Enclosure Enclosure `xml:"enclosure"`
}

type ByTitle []Item

func (items ByTitle) Len() int           { return len(items) }
func (items ByTitle) Swap(i, j int)      { items[i], items[j] = items[j], items[i] }
func (items ByTitle) Less(i, j int) bool { return items[i].Title < items[j].Title }
