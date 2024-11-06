package main

// Std imports
import (
	"fmt"
	"os"
	"net"
)

// App imports
import (
	"gomud/dungeonmap"
	"gomud/game"
)

// The server version
const version = 1

func main() {
	fmt.Printf("GoMUD v%d\n", version)

	if len(os.Args) != 3 {
		fmt.Println("Invalid number of args")
		os.Exit(1)
	}

	// Load the mud map JSON
	dungeonMap, error := dungeonmap.ReadDungeonMap(os.Args[1])
	if error != nil {
		panic(error)
	}
	fmt.Println(dungeonMap)

	// Set up a server
	listener, error := net.Listen("tcp", ":" + os.Args[2])
	if error != nil {
		panic(error)
	}

	// List for connections until something breaks
	for {
		client, error := listener.Accept()
		if error != nil {
			panic(error)
		}		

		// Handle connections asynchronously
		go game.HandleConnection(client, dungeonMap)
	}
}
