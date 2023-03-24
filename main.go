package main

import (
	"fmt"
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/personalization"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/recommend"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/region"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

type Record struct {
	ObjectID string `json:"objectID"`
	Name     string `json:"name"`
	NewCol   string `json:"newCol"`
}

func algoSearch(apiKey string) {
	client := search.NewClient("HKWXA0EXKH", apiKey)

	index := client.InitIndex("test_datium")
	// write
	resSave, err := index.SaveObject(Record{ObjectID: "test", Name: "test", NewCol: "test"})
	if err != nil {
		panic(err)
	}
	resSave.Wait()

	// search
	params := []interface{}{
		opt.AttributesToRetrieve("email", "company", "city"),
		opt.HitsPerPage(10),
		opt.UserToken("stephen-user-test"),
	}
	res, err := index.Search("Kathleen", params...)
	if err != nil {
		panic(err)
	}
	fmt.Println("QUERY ID")
	fmt.Println(res.QueryID, res.Hits[0])

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

	settings, _ := index.GetSettings()
	fmt.Println(settings)
}

func algoRecommendation(apiKey string) {
	rec := recommend.NewClient("HKWXA0EXKH", apiKey)
	options := recommend.NewRelatedProductsOptions("test_datium", "saas-sample-data-69", 0, nil, nil, nil)
	res, err := rec.GetRelatedProducts([]recommend.RelatedProductsOptions{options})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Results[0].Hits[0])

	options2 := recommend.RecommendationsOptions{IndexName: "test_datium", ObjectID: "saas-sample-data-69", Model: recommend.BoughtTogether}
	res2, err := rec.GetRecommendations([]recommend.RecommendationsOptions{options2})
	if err != nil {
		panic(err)
	}
	fmt.Println(res2)
}

func algoPersonalization(apiKey string) {
	client := personalization.NewClientWithConfig(
		personalization.Configuration{
			AppID:  "HKWXA0EXKH",
			APIKey: apiKey,
			Region: region.US,
		})
	strategy := personalization.Strategy{
		EventsScoring: []personalization.EventsScoring{
			{"Add to cart", "conversion", 50},
			{"Purchase", "conversion", 100},
		},
		FacetsScoring: []personalization.FacetsScoring{
			{"brand", 100},
			{"categories", 10},
		},
		PersonalizationImpact: opt.PersonalizationImpact(50),
	}
	res, _ := client.SetPersonalizationStrategy(strategy, true)
	fmt.Println(res)

	res2, _ := client.GetPersonalizationStrategy()
	fmt.Println(res2)
}

func main() {
	// need search/write key
	apiKey := os.Getenv("key")
	algoSearch(apiKey)
	algoRecommendation(apiKey)
	algoPersonalization(apiKey)
}
