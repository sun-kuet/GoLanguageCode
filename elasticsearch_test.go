package main

import (
	"gopkg.in/olivere/elastic.v3"
	"fmt"
	"bufio"
	"os"
	"strconv"
	"gopkg.in/square/go-jose.v1/json"
)

type Tweet struct {
	User    string        `json:"user"`
	Message string        `json:"message"`
}

func main() {
	reader:= bufio.NewReader(os.Stdin)

	client, err := elastic.NewClient()

	if err != nil {
		panic("client create hoy nai")
	}
	_, err = client.DeleteIndex("_all").Do()

	exists, err := client.IndexExists("twitter").Do()
	if err != nil {
		panic(err)
	}

	if (!exists) {
		_, err = client.CreateIndex("twitter").Do()

		if err != nil {
			panic(err)
		}
	}



	for i := 1; i <= 3; i++ {
		name, _ := reader.ReadString('\n')
		mesg, _ := reader.ReadString('\n')
		name = name[0:len(name)-1]
		mesg = mesg[0:len(mesg)-1]

		tweet := Tweet{User:name, Message:mesg }
		ret, err := client.Index().Index("twitter").Type("tweet").Id(strconv.Itoa(i)).BodyJson(tweet).Do()
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s -> %s -> %s\n", ret.Index, ret.Type, ret.Id)
	}

	fmt.Println("\n\n\nTerm searching query")

	retget , err := client.Search().Index("twitter").Query(elastic.NewTermQuery("user", "sun")).Sort("user",true).Pretty(true).Do()
	if err!=nil{
		panic(err)
	}

	if retget.Hits.TotalHits != 0 {

		for _, hit := range retget.Hits.Hits{

			var t Tweet
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil{
				panic(err)
			}

			fmt.Printf("%s : %s \n", t.User, t.Message)
		}

	}else {
		fmt.Printf("No term query result found\n")
	}

	fmt.Println("\n\n\nPrefix searching query")

	partialSearch, err := client.Search().Index("twitter").Query(elastic.NewPrefixQuery("user","sun")).Pretty(true).Do()
	if err !=nil{
		panic(err)
	}

	if partialSearch.Hits.TotalHits != 0 {

		for _, hit := range partialSearch.Hits.Hits{

			var t Tweet
			prefixSearchError := json.Unmarshal(*hit.Source, &t)
			if prefixSearchError != nil{
				panic(prefixSearchError)
			}

			fmt.Printf("%s : %s \n", t.User, t.Message)
		}

	}else {
		fmt.Printf("No prefix query result found\n")
	}


	fmt.Println("\n\n\nPhrase searching query")
	phraseMatchingSearch, err := client.Search().Index("twitter").Query(elastic.NewMatchPhraseQuery("name","sun tamanna")).Do()

	if phraseMatchingSearch.Hits.TotalHits != 0 {

		for _, hit := range phraseMatchingSearch.Hits.Hits{

			var t Tweet
			phraseMatchError := json.Unmarshal(*hit.Source, &t)
			if phraseMatchError != nil{
				panic(phraseMatchError)
			}

			fmt.Printf("%s : %s \n", t.User, t.Message)
		}

	}else {
		fmt.Printf("No phrase matching result found\n")
	}

	fmt.Println("\n\n\nMulti-Match searching query")

	multisearchFields := make([]string, 0)
	multisearchFields = append(multisearchFields,"user")
	multisearchFields = append(multisearchFields, "message")
	multisearchQ := elastic.NewMultiMatchQuery("sun kuet" ,multisearchFields...)

	multisearchQuery, err := client.Search().Index("twitter").Query(multisearchQ).Do()
	if err!= nil{
		panic(err)
	}

	if multisearchQuery.Hits.TotalHits != 0 {

		for _, hit := range multisearchQuery.Hits.Hits{

			var t Tweet
			multiMatchError := json.Unmarshal(*hit.Source, &t)
			if multiMatchError != nil{
				panic(multiMatchError)
			}

			fmt.Printf("%s : %s \n", t.User, t.Message)
		}

	}else {
		fmt.Printf("No multi-search result found\n")
	}

	// bool query
	fmt.Println("\n\n\nBool Search Query")

	boolSearchQ := elastic.NewBoolQuery()
	boolSearchQ = boolSearchQ.Should(elastic.NewMatchQuery("user", "sun"),elastic.NewMatchQuery("user","tamanna"))
	boolSearchQ = boolSearchQ.Must(elastic.NewMatchQuery("message","facebook"))

	boolSearchQuery, err := client.Search().Index("twitter").Query(boolSearchQ).Do()

	if boolSearchQuery.Hits.TotalHits != 0 {

		for _, hit := range boolSearchQuery.Hits.Hits{

			var t Tweet
			boolSearchError := json.Unmarshal(*hit.Source, &t)
			if boolSearchError != nil{
				panic(boolSearchError)
			}

			fmt.Printf("%s : %s \n", t.User, t.Message)
		}

	}else {
		fmt.Printf("No bool-search result found\n")
	}

	// Id query
	fmt.Println("\n\n\nId result query")

	ids := make([]string , 0)
	ids = append(ids,"1")
	ids = append(ids, "3")

	idSearchQ := elastic.NewIdsQuery("tweet").Ids("1","3","2")
	idSearchQuery , err := client.Search().Index("twitter").Query(idSearchQ).Do()

	if idSearchQuery.Hits.TotalHits != 0{

		for _, hit := range idSearchQuery.Hits.Hits{

			var t Tweet
			idSearchError := json.Unmarshal(*hit.Source, &t)
			if idSearchError != nil{
				panic(idSearchError)
			}

			fmt.Printf("%s : %s\n", t.User, t.Message)
		}

	}else {
		fmt.Printf("No id-search result found\n")
	}


}
