package main

import "log"
import "github.com/tbailey/go-unibit/client"

func main()  {
	request := client.Request{
		Tickers: []string{"AAPL","WORK"},
		StartDate: "2019-08-16",
		EndDate: "2019-08-20",
	}

	client := client.NewClient("Demo")
	r, err := client.GetCompanyProfile(request)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", r)
}
