package main

import (
	"log"
	"strings"
	"golang.org/x/net/html"

	"github.com/PuerkitoBio/goquery"
)

func dropsau() {
	dropsauURL := "https://dropsau.com/pages/memberships"
	doc, err := goquery.NewDocument(dropsauURL)

	if err != nil {
		log.Fatal(err)
	}

	soldoutText := doc.Find("#sold-out")
	
	if strings.Contains(soldoutText.Text(), "Buy now") {
		log.Println("Buy now button is available")
		log.Println(soldoutText.Text())
		return
	}
	log.Println("No change found")
}