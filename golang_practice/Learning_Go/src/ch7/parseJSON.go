package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"math/big"
)

type Tour struct {
	Name, Price string
}

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)

	tours := toursFromJson(content)
	// fmt.Println(tours)
	
	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		fmt.Printf("%v ($%.2f)\n", tour.Name, price)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {
	
	resp, err := http.Get(url)
	checkError(err)
	
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)
	
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)
	
	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}
	
	return tours
}

/*

 % go run parseJSON.go 
2 Days Adrift the Salton Sea ($256.00)
A Week of Wine ($768.00)
Amgen Tour of California Special ($4096.00)
Avila Beach Hot springs ($768.00)
Big Sur Retreat ($512.00)
Channel Islands Excursion ($128.00)
Coastal Experience ($1024.00)
Cycle California: My Way ($1024.00)
Day Spa Package ($512.00)
Endangered Species Expedition ($512.00)
Fossil Tour ($384.00)
Hot Salsa Tour ($384.00)
Huntington Library and Pasadena Retreat Tour ($192.00)
In the Steps of John Muir ($512.00)
Joshua Tree: Best of the West Tour ($128.00)
Kids L.A. Tour ($192.00)
Mammoth Mountain Adventure ($768.00)
Matilija Hot springs ($768.00)
Mojave to Malibu ($192.00)
Monterey to Santa Barbara Tour ($2048.00)
Mountain High Lift-off ($768.00)
Olive Garden Tour ($64.00)
Oranges & Apples Tour ($256.00)
Restoration Package ($768.00)
The Death Valley Survivor's Trek ($192.00)
The Mt. Whitney Climbers Tour ($512.00)

*/