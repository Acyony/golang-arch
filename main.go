package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}

/*
type Bird struct {
	Species     string
	Description string
}*/

func main() {
	/*p1 := person{
		First: "Alcione",
	}

	p2 := person{
		First: "Alice",
	}

	xp := []person{p1, p2}
	// NOTE: bs - byte of slice

	// json.Marshal: encoding Go struct into a json string

	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("PRINT JSON", string(bs))

	xp2 := []person{}

	// json.Unmarschal: decode the data from json to a struct
	//&xp2 is the adress of the variable I want to store the parsed data in
	err = json.Unmarshal(bs, &xp2)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Back into a Go data structure", xp2)*/

	/// ---------------------- =^.^= ----------------------

	/*birdJson := `{"species": "pigeon", "description": "likes cookies"}`
	fmt.Println(birdJson)
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)*/

	/// ---- =^.^= encoding and decoding =^.^=-----
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	fmt.Println("Server running on port 8080")
	fmt.Println("Press CONTROL + C to stop the server")
	http.ListenAndServe(":8080", nil)

}

/// ---- =^.^= encoding - Sensing data response.write
func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Jenny",
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data", err)
	}
}

/// ---- =^.^= decoding - User sending me the data
func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decoded bad data", err)
	}

	log.Println("Person:", p1)
}
