package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "The service is running on address")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJsonApiServer(*listenAddr, svc)
	server.Run()

	// price, err := svc.FetchPrice(context.Background(), "BTC")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	//fmt.Println(price)
}
