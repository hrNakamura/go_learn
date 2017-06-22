package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

//!+main

var m sync.Mutex

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/add", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/remove", db.remove)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) add(w http.ResponseWriter, req *http.Request) {
	m.Lock()
	defer m.Unlock()
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusForbidden) // 400
		fmt.Fprintf(w, "%q exists\n", item)
		return
	}
	priceQuery := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceQuery, 32)
	if err != nil {
		w.WriteHeader(http.StatusForbidden) // 400
		fmt.Fprintf(w, "%q is not a number\n", priceQuery)
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "Add %s : %s\n", item, db[item])
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	m.Lock()
	defer m.Unlock()
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusForbidden) // 400
		fmt.Fprintf(w, "%q not found\n", item)
		return
	}
	priceQuery := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceQuery, 32)
	if err != nil {
		w.WriteHeader(http.StatusForbidden) // 400
		fmt.Fprintf(w, "%q is not a number\n", priceQuery)
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "Update %s : %s\n", item, db[item])
}

func (db database) remove(w http.ResponseWriter, req *http.Request) {
	m.Lock()
	defer m.Unlock()
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusForbidden) // 400
		fmt.Fprintf(w, "%q not found\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "Remove %s \n", item)
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	var templateList = template.Must(template.New("templateList").Parse(`
	  <h1>prices</h1>
	  <table>
	  <tr style='text-align: left'>
	    <th>Item</th>
	    <th>Price</th>
	  </tr>
	  {{range $index, $var := .}}
	  <tr>
	  	<td>{{$index}}</td>
	  	<td>{{$var}}</td>
	  </tr>
	  {{end}}
	  </table>
	  `))
	if err := templateList.Execute(w, db); err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "error request")
	}
	// for item, price := range db {
	// 	fmt.Fprintf(w, "%s: %s\n", item, price)
	// }
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
