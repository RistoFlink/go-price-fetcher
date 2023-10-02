package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "address the service is running on")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJsonApiServer(*listenAddr, svc)
	server.Run()

	// price, err := svc.FetchPrice(context.Background(), "BTC")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	//fmt.Println(price)

	/* 	client := client.New("http://localhost:3000")
	price, err := client.FetchPrice(context.Background(), "ET")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", price) */
}
