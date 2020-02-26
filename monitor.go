package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func getProducts() []*Product{
	resp, err := http.Get("https://laced.com.au/products.json")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	
	var products []*Product
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&products)
	if err != nil {
		panic(err)
	}
	fmt.Println(products)
	return products
}

