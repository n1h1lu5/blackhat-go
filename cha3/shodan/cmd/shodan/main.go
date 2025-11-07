package main

import (
	"fmt"
	"log"
	"os"

	"blackhat-go/cha3/shodan/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}
	apikey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apikey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("Query Credits: %d\n Scan Credits: %d\n\n", info.QueryCredits, info.ScanCredits)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
