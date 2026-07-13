package v2

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// AgentModSchema is one mod selected for one agent.
//
// It replaces the embedded AgentModConfig.SelectedMods array, whose whole-array
// rewrites by three separate writers lost updates. Every mutation here is a
// targeted findOneAndUpdate keyed on {agentId, modReference}.
//
// AccountID is denormalised because the collection is cross-account: it is the
// authz boundary the embedded array got for free from its parent document.
type AgentModSchema struct {
	ID        bson.ObjectID `json:"_id" bson:"_id"`
	AgentID   bson.ObjectID `json:"agentId" bson:"agentId"`
	AccountID bson.ObjectID `json:"accountId" bson:"accountId"`

	ModID        bson.ObjectID `json:"modId" bson:"modId" mson:"collection=mods"`
	ModReference string        `json:"modReference" bson:"modReference"`

	// DesiredVersion is pinned. Only a user action moves it; the catalogue job
	// never does, so a bad mod release cannot take down a live server unattended.
	DesiredVersion   string `json:"desiredVersion" bson:"desiredVersion"`
	InstalledVersion string `json:"installedVersion" bson:"installedVersion"`
	Installed        bool   `json:"installed" bson:"installed"`

	NeedsUpdate   bool   `json:"needsUpdate" bson:"needsUpdate"`
	LatestVersion string `json:"latestVersion" bson:"latestVersion"`

	Config string `json:"config" bson:"config"`

	// Direct is true when the user chose this mod, false when the resolver pulled
	// it in as a dependency. Removal consults it: a dependency is uninstalled only
	// once no direct mod requires it. The old code deleted anything on disk that
	// was not in the selected list, so a transient database error could empty the
	// Mods directory.
	Direct bool `json:"direct" bson:"direct"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

func NewAgentModDoc(agentID, accountID, modID bson.ObjectID, modReference, desiredVersion string, direct bool) AgentModSchema {
	now := time.Now()

	return AgentModSchema{
		ID:             bson.NewObjectID(),
		AgentID:        agentID,
		AccountID:      accountID,
		ModID:          modID,
		ModReference:   modReference,
		DesiredVersion: desiredVersion,
		Direct:         direct,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}

// ModLock is one fully pinned mod: everything the agent needs to install it
// without asking anything else. The agent resolves nothing.
type ModLock struct {
	ModReference string `json:"modReference"`
	Version      string `json:"version"`
	DownloadURL  string `json:"downloadUrl"`
	Hash         string `json:"hash"`
	Size         int64  `json:"size"`
	Config       string `json:"config"`
	Direct       bool   `json:"direct"`
}

// Lockfile is the syncmods task payload: the complete desired state of the
// agent's Mods directory. Anything on disk and absent from here is removed.
//
// There is deliberately no GameFeature flag. It lives in the .uplugin inside the
// archive and the catalogue does not carry it, so the agent reads it after
// unpacking, as it already does.
type Lockfile struct {
	SFVersion string    `json:"sfVersion"`
	Mods      []ModLock `json:"mods"`
}

// InstalledMod is the agent's report of what is actually on its disk.
type InstalledMod struct {
	ModReference     string `json:"modReference"`
	InstalledVersion string `json:"installedVersion"`
	Installed        bool   `json:"installed"`
}
