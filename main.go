package main

import (
    "log"
)

func main() {
    a := App{}
    a.Initialize("root", "", "test")
    address := ":8080"
	log.Print("Listening on " + address)
    a.Run(address)
}
