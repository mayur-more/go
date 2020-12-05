package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	fmt.Println(states)
	
	states["MH"] = "Maharashtra"
	states["MP"] = "Madhyapradesh"
	fmt.Println(states)
	
	state_name := states["MP"]
	fmt.Println(state_name)
	
	delete(states, "MP")
	fmt.Println(states)
	
	states["RJ"] = "Rajstan"
	states["GJ"] = "Gujrat"
	
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}
	
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	
	fmt.Println("\n Sorted: ")
	for i:= range keys {
		fmt.Println(states[keys[i]])
	}
}

/*

 % go run maps.go
map[]
map[MH:Maharashtra MP:Madhyapradesh]
Madhyapradesh
map[MH:Maharashtra]
MH: Maharashtra
RJ: Rajstan
GJ: Gujrat
[GJ MH RJ]

 Sorted: 
Gujrat
Maharashtra
Rajstan

*/