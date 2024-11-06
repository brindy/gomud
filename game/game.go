package game

import (
	"fmt"
	"net"
)

import (
	"gomud/dungeonmap"
)

func HandleConnection(client net.Conn, dungeonMap dungeonmap.DungeonMap) {

	fmt.Println(client)
	fmt.Println(dungeonMap)
	
}
