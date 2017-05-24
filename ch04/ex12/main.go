package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type xkcd struct {
	Title      string `json:"title,omitempty"`
	URLHTML    string `json:"url_html,omitempty"`
	Transcript string `json:"transcript,omitempty"`
}

const xkcdURL = "https://xkcd.com/"
const xkcdInfo = "/info.0.json"

func main() {
	var (
		mode   int
		number int
	)
	flag.IntVar(&mode, "m", 0, "動作モード")
	flag.IntVar(&number, "n", 0, "番号")
	flag.Parse()
	if mode == 1 {
		generateIndex(number)
		return
	}
	var indexList []xkcd
	for i := 614; ; i++ {
		indexStr := "./index" + strconv.Itoa(i) + ".json"
		data, err := ioutil.ReadFile(indexStr)
		if err != nil {
			break
		}
		var index xkcd
		if err := json.Unmarshal(data, &index); err != nil {
			log.Fatalf("%s", err)
		}
		indexList = append(indexList, index)
	}

	for _, v := range indexList {
		if strings.Contains(v.Title, os.Args[1]) {
			fmt.Println("URL:")
			fmt.Printf("%v\n", v.URLHTML)
			fmt.Println("Transcript:")
			fmt.Printf("%v\n", v.Transcript)
		}
	}
}

func generateIndex(number int) {
	qURL := xkcdURL + strconv.Itoa(number) + xkcdInfo
	resp, err := http.Get(qURL)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	defer resp.Body.Close()
	var data xkcd
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("%s\n", err)
	}
	data.URLHTML = qURL
	jsondata, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Printf("%s", jsondata)
	// assert(data)
	// fmt.Println()
}

func assert(data interface{}) {
	switch data.(type) {
	case string:
		fmt.Printf("%v,\n", data.(string))
	case float64:
		fmt.Printf("%v,\n", data.(float64))
	case bool:
		fmt.Printf("%v,\n", data.(bool))
	case nil:
		fmt.Printf("%v,\n", "null")
	case []interface{}:
		fmt.Print("[")
		for _, v := range data.([]interface{}) {
			assert(v)
			fmt.Print(" ")
		}
		fmt.Print("]")
	case map[string]interface{}:
		fmt.Print("{")
		for k, v := range data.(map[string]interface{}) {
			fmt.Print(k + ":")
			assert(v)
			fmt.Print(" ")
		}
		fmt.Print("}")
	default:
	}
}
