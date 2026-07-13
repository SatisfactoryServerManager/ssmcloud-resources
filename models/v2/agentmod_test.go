package v2

import (
	"encoding/json"
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestNewAgentModDocDefaults(t *testing.T) {
	agentID := bson.NewObjectID()
	accountID := bson.NewObjectID()
	modID := bson.NewObjectID()

	doc := NewAgentModDoc(agentID, accountID, modID, "RefinedPower", "3.3.0", true)

	if doc.ID.IsZero() {
		t.Fatal("expected an id to be minted")
	}
	if doc.AgentID != agentID || doc.AccountID != accountID || doc.ModID != modID {
		t.Fatal("expected the ownership fields to be carried through")
	}
	if doc.ModReference != "RefinedPower" || doc.DesiredVersion != "3.3.0" {
		t.Fatal("expected the mod identity to be carried through")
	}
	if !doc.Direct {
		t.Fatal("expected direct to be carried through")
	}
	// A freshly selected mod is not installed until an agent says so.
	if doc.Installed || doc.InstalledVersion != "" || doc.NeedsUpdate {
		t.Fatal("expected a new doc to claim nothing about the agent's disk")
	}
	if doc.CreatedAt.IsZero() || doc.UpdatedAt.IsZero() {
		t.Fatal("expected timestamps to be set")
	}
}

// The lockfile is the syncmods payload: it crosses the wire as the task's JSON
// data string, so its json tags are load-bearing for the agent.
func TestLockfileJSONTags(t *testing.T) {
	lf := Lockfile{
		SFVersion: "1.1.0.0",
		Mods: []ModLock{{
			ModReference: "RefinedPower",
			Version:      "3.3.0",
			DownloadURL:  "https://api.ficsit.app/v1/version/abc/WindowsServer/download",
			Hash:         "deadbeef",
			Size:         1234,
			Config:       "{}",
			Direct:       true,
		}},
	}

	b, err := json.Marshal(lf)
	if err != nil {
		t.Fatal(err)
	}

	var round Lockfile
	if err := json.Unmarshal(b, &round); err != nil {
		t.Fatal(err)
	}
	if len(round.Mods) != 1 || round.Mods[0].Hash != "deadbeef" || round.Mods[0].DownloadURL == "" {
		t.Fatalf("lockfile did not survive a json round trip: %s", string(b))
	}
}
