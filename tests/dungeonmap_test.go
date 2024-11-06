
package tests

import "gomud/dungeonmap"
import "testing"

func TestReadDungeonMap(t *testing.T) {
	 dungeonMap, error := dungeonmap.ReadDungeonMap("dungeonmap_test.json")
	 if error != nil {
	 	t.Errorf("Failed to load test map")
	 }

	 if dungeonMap.StartArea != "start" {
	 	t.Errorf("Missing expected start-area property")
	 }

	 // TODO more extensive testing
}
