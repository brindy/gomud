
package dungeonmap

import (
    "encoding/json"
    "fmt"
    "os"
)

type DungeonMap struct {
	StartArea string `json:"start-area"`
	Areas map[string]DungeonArea `json:"areas"`
}

type DungeonArea struct {
	OnEntry string `json:"on-entry"`
	Look string `json:"look"`
	Exits map[string]DungeonAreaExit `json:"exits"`
}

type DungeonAreaExit struct {
	To string `json:"to"`
}

func ReadDungeonMap(fileName string) (DungeonMap, error) {
	var dungeonMap DungeonMap

    file, err := os.ReadFile(fileName)
    if err != nil {
        return dungeonMap, fmt.Errorf("error reading file: %w", err)
    }

    err = json.Unmarshal(file, &dungeonMap)
    if err != nil {
        return dungeonMap, fmt.Errorf("error unmarshalling JSON: %w", err)
    }

    return dungeonMap, nil
}
