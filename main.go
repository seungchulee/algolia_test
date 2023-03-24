package main

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

type Record struct {
	ObjectID string `json:"objectID"`
	Name     string `json:"name"`
}

func main() {
	client := search.NewClient("", "")

	index := client.InitIndex("")
	resSave, err := index.SaveObjects(Record{ObjectID: "", Name: ""})
	if err != nil {
		panic(err)
	}
	resSave.Wait()
}
