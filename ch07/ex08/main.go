package main

import (
	"container/list"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"GoGo", "Delilah", "Roots", 2012, length("3m38s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!-main

//!+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//!-printTracks

//!+artistcode
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-artistcode

//!+yearcode
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-yearcode

type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byAlbum []*Track

func (x byAlbum) Len() int           { return len(x) }
func (x byAlbum) Less(i, j int) bool { return x[i].Album < x[j].Album }
func (x byAlbum) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byLength []*Track

func (x byLength) Len() int           { return len(x) }
func (x byLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x byLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	history := list.New()
	if len(os.Args) > 0 {
		for _, clicked := range os.Args[1:] {
			history.PushBack(clicked)
			if history.Len() > 5 {
				history.Remove(history.Front())
			}
		}
	}
	sortFunc := func(x, y *Track) bool {
		if history.Len() == 0 {
			return x.Title < y.Title
		}
		for element := history.Back(); element != nil; element = element.Prev() {
			switch element.Value {
			case "Title":
				if x.Title != y.Title {
					return x.Title < y.Title
				}
			case "Artist":
				if x.Artist != y.Artist {
					return x.Artist < y.Artist
				}
			case "Album":
				if x.Album != y.Album {
					return x.Album < y.Album
				}
			case "Year":
				if x.Year != y.Year {
					return x.Year < y.Year
				}
			case "Length":
				if x.Length != y.Length {
					return x.Length < y.Length
				}
			}
		}
		return false
	}
	//!+customcall
	sort.Sort(customSort{tracks, sortFunc})
	//!-customcall
	printTracks(tracks)

	fmt.Println("")

	for element := history.Back(); history.Len() != 0 && element != nil; element = element.Prev() {
		switch element {
		case "Title":
			sort.Stable(byTitle(tracks))
		case "Artist":
			sort.Stable(byArtist(tracks))
		case "Album":
			sort.Stable(byAlbum(tracks))
		case "Year":
			sort.Stable(byYear(tracks))
		case "Length":
			sort.Stable(byLength(tracks))
		}
	}
	printTracks(tracks)
}

//!+customcode
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

//!-customcode
