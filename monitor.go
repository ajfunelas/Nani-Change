package main

import (
	"http"
	"encoding/json"
	"fmt"
)


resp, err := http.Get("https://laced.com.au/products.json")

fmt.Println(resp)
