package tracks

import "fmt"

var TRACKS = make(map[string]string)

func AddTrack(id string, content string) {
	TRACKS[id] = content
}

func RemoveTrack(id string) {
	delete(TRACKS, id)
}

func DumpTracks() {
	fmt.Println(TRACKS)
}
