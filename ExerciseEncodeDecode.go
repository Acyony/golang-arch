package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	First string
}

func main() {

	/// ---- =^.^= encoding and decoding =^.^=-----
	http.HandleFunc("/decode", dec)
	http.HandleFunc("/encode", enc)

	/// ---- =^.^= setting the server =^.^=-----
	fmt.Println("Server running on port 6000")
	fmt.Println("Press CONTROL + C to stop the server")
	http.ListenAndServe(":6000", nil)

}

/// ---- =^.^= encoding - sending data to response writer

func enc(w http.ResponseWriter, r *http.Request) {
	u1 := user{
		First: "Jenny",
	}
	u2 := user{
		First: "Josh",
	}
	users := []user{u1, u2}
	err := json.NewEncoder(w).Encode(&users)
	if err != nil {
		log.Println("Encoded bad data", err)
	}
}

/// ---- =^.^= decoding - receiving requested data from some user

func dec(w http.ResponseWriter, r *http.Request) {
	users := []user{}
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		log.Println("Decode bad data", err)
	}
	log.Println("Users:", users)
}
