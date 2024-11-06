package game

import (
	"fmt"
	"io"
)

import (
	"gomud/dungeonmap"
)

type Client interface {

	io.Reader
	io.Writer
	Close() error 

}

func HandleConnection(client Client, dungeonMap dungeonmap.DungeonMap) {

	fmt.Println(client)
	fmt.Println(dungeonMap.StartArea)

	error := client.Close()
	if error != nil {
		panic(error) // should be a warning
	}

	client.Write([]byte("Hello Client"))
	fmt.Println("closing")
	client.Close()

}
