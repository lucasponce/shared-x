package tracks

import "testing"

func TestTracks(t *testing.T) {

	AddTrack("k1", "v1")
	AddTrack("k2", "v2")
	AddTrack("k3", "v3")

	RemoveTrack("k3")

	version := GetVersion()

	if version != "v2.0" {
		t.Errorf("Version %s, want %s", version, "v2.0")
	}

	tracks := GetTracks()
	if len(tracks) != 2 {
		t.Errorf("GetTracks returned %d tracks, want 2", len(tracks))
	}

}
