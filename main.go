package main

import (
	"fmt"
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

type Record struct {
	ObjectID string `json:"objectID"`
	Name     string `json:"name"`
}

func main() {
	apiKey := os.Getenv("key")
	client := search.NewClient("HKWXA0EXKH", apiKey)

	index := client.InitIndex("test_datium")
	// write
	resSave, err := index.SaveObject(Record{ObjectID: "test", Name: "test"})
	if err != nil {
		panic(err)
	}
	resSave.Wait()

	// search
	params := []interface{}{
		opt.AttributesToRetrieve("email", "company", "city"),
		opt.HitsPerPage(10),
	}
	res, err := index.Search("Kathleen", params...)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Hits[0])

	r, err := index.SearchForFacetValues("company", "Associates")
	if err != nil {
		panic(err)
	}
	fmt.Println(r.FacetHits)

	f := func(o map[string]interface{}) bool {
		itf, ok := o["objectID"]
		if !ok {
			return false
		}
		hitObjectID, ok := itf.(string)
		return ok && hitObjectID == "test"
	}
	rr, _ := index.FindObject(f, "hello", false)
	fmt.Println(rr)
	position := res.GetObjectPosition("saas-sample-data-98")
	fmt.Println(position)
}
