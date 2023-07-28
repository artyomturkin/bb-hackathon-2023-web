package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// u, _ := url.Parse("http://localhost:5173")
	// p := httputil.NewSingleHostReverseProxy(u)
	// http.Handle("/", p)
	http.HandleFunc("/api/parse", demoHandler)
	http.ListenAndServe(":3003", nil)
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	response := Bill{
		Total: 990,
		Items: []Item{
			{Name: "position 1", Price: 100},
			{Name: "position 2", Price: 100},
			{Name: "position 3", Price: 100},
			{Name: "position 4", Price: 100},
		},
	}

	b, _ := json.Marshal(response)
	fmt.Printf("err: %v\nresponse:\n%s\n", err, string(b))

	w.Write(b)
}

type Bill struct {
	Total      float32
	WithoutTip float32
	TipPercent float32
	TipAmount  float32

	Items []Item
}

type Item struct {
	Name     string
	Price    float32
	Quantity int
}
