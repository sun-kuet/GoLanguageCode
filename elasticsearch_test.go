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

	for i := 1; i <= 5; i++ {
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

	fmt.Println("start searching")

	retget , err := client.Search().Index("twitter").Query(elastic.NewTermQuery("user", "s")).Sort("user",true).Pretty(true).Do()
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
		fmt.Printf("No tweet found\n")
	}

}
