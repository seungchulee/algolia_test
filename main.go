package main

import (
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

func main() {
	client := search.NewClient("HKWXA0EXKH", "")

	index := client.InitIndex("test_datium")
	params := []interface{}{
		opt.AttributesToRetrieve("email", "company", "city"),
		opt.HitsPerPage(1),
	}
	res, err := index.Search("Kathleen", params...)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Hits[0])

}
