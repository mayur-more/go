package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()
	
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	content := string(bytes)
	fmt.Print(content) 

}

/*

 % go run webRequest.go 
Response type: *http.Response
[{"tourId":"14","packageId":"5","packageTitle":"From Desert to Sea","name":"2 Days Adrift the Salton Sea","blurb":"The Salton Sea, 25% saltier than the Pacific, has been a tourist destination since the 1920s. See what attracts people to this desert oasis.","description":"The Salton Sea is saltier than the Pacific, an unusual feat for inland body of water. And even though its salinity has risen over the years, due in part to lack of outflows and pollution from agricultural runoff, it has attracted a small, but dedicated population. The sea itself offers recreational opportunities including boating, camping, off-roading, hiking, use of personal watercraft, photography and bird 
...

*/