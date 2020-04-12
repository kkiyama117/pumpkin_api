package main

import (
	"log"

	"hintan.jp/pumpkin_api/src/factories"
)

func main() {
	// Passing -router to the program will generate docs for the above
	// router definition. See the `router.json` file in this folder for
	// the output.
	server := factories.NewServer()
	log.Fatal(server.Run())
}
