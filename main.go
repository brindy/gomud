package main

// Std imports
import "fmt"
import "os"

// App imports
import "gomud/dungeonmap"

// The server version
const version = 1

func main() {
	fmt.Printf("GoMUD v%d\n", version)

	// Load the mud map JSON
	dungeonMap, error := dungeonmap.ReadDungeonMap(os.Args[1])
	if error != nil {
		panic(error)
	}

	// Start a server
	fmt.Println("Starting server")

	// On connect create a new MUD client and share the server state

}
