package tracks

import "fmt"

var VERSION = "v0.0"
var TRACKS = make(map[string]string)

func AddTrack(id string, content string) {
	TRACKS[id] = content
}

func RemoveTrack(id string) {
	delete(TRACKS, id)
}

func DumpTracks() {
	fmt.Printf("VERSION: %s TRACKS:\n", VERSION)
	fmt.Println("-----------------------")
	for k, v := range TRACKS {
		fmt.Printf("%s: %s\n", k, v)
	}

}
