// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"container/list"
	"fmt"
	"html/template"
	"log"
	"net/http"
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

var htmlTmpl = template.Must(template.New("tracks").Parse(`
<html>
<body>
<table>
<tr>
<th><a href="?sort=Title">Title</a></th>
<th><a href="?sort=Artist">Artist</a></th>
<th><a href="?sort=Album">Album</a></th>
<th><a href="?sort=Year">Year</a></th>
<th><a href="?sort=Length">Length</a></th>
</tr>
{{range .}}
<tr>
<td>{{.Title}}</td>
<td>{{.Artist}}</td>
<td>{{.Album}}</td>
<td>{{.Year}}</td>
<td>{{.Length}}</td>
</tr>
{{end}}
</table>
</body>
</html>
	`))

func main() {
	history := list.New()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		clicked := r.FormValue("sort")
		history.PushBack(clicked)
		if history.Len() > 5 {
			history.Remove(history.Front())
		}
		sort.Sort(customSort{tracks, func(x, y *Track) bool {
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
		}})

		err := htmlTmpl.Execute(w, tracks)
		if err != nil {
			log.Printf("template error: %s", err)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
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
