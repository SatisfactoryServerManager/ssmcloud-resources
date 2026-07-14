package v2

import (
	"strings"
	"time"

	"github.com/SatisfactoryServerManager/ssmcloud-resources/models"
	"github.com/SatisfactoryServerManager/ssmcloud-resources/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type AgentSchema struct {
	ID        bson.ObjectID `json:"_id" bson:"_id"`
	AgentName string        `json:"agentName" bson:"agentName"`
	APIKey    string        `json:"apiKey" bson:"apiKey"`
	Status    AgentStatus   `json:"status" bson:"status"`

	Config       AgentConfig       `json:"config" bson:"config"`
	ServerConfig AgentServerConfig `json:"serverConfig" bson:"serverConfig"`

	MapData AgentMapData `json:"mapData" bson:"mapData"`

	Saves   []AgentSave   `json:"saves" bson:"saves"`
	Backups []AgentBackup `json:"backups" bson:"backups"`

	// ConnectedTo names the backend replica currently holding this agent's task
	// stream. ConnectionID is minted by that replica per stream, not supplied by
	// the agent, so a slow teardown of an old stream cannot detach a fresh one.
	ConnectedTo  string `json:"connectedTo,omitempty" bson:"connectedTo,omitempty"`
	ConnectionID string `json:"connectionId,omitempty" bson:"connectionId,omitempty"`

	LogIds bson.A           `json:"-" bson:"logs" mson:"collection=agentlogs"`
	Logs   []AgentLogSchema `json:"logs" bson:"-"`

	LatestAgentVersion string `json:"latestAgentVersion" bson:"latestAgentVersion"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type AgentStatus struct {
	Online    bool `json:"online" bson:"online"`
	Running   bool `json:"running" bson:"running"`
	Installed bool `json:"installed" bson:"installed"`

	CPU float64 `json:"cpu" bson:"cpu"`
	RAM float64 `json:"ram" bson:"ram"`

	LastCommDate time.Time `json:"lastCommDate" bson:"lastCommDate"`

	InstalledSFVersion int64 `json:"installedSFVersion"`
	LatestSFVersion    int64 `json:"latestSFVersion"`
}

type AgentConfig struct {
	Version          string  `json:"version" bson:"version"`
	Port             int     `json:"port" bson:"port"`
	Memory           int64   `json:"memory" bson:"memory"`
	IP               string  `json:"ip" bson:"ip"`
	BackupKeepAmount int     `json:"backupKeepAmount" bson:"backupKeepAmount"`
	BackupInterval   float32 `json:"backupInterval" bson:"backupInterval"`

	// Platform is the agent's mod target name: "WindowsServer" or "LinuxServer".
	// The backend needs it to pin the right build in a lockfile; the agent used to
	// make this choice itself, and the backend has no other way to know.
	Platform string `json:"platform" bson:"platform"`
}

type AgentServerConfig struct {
	UpdateOnStart bool   `json:"updateOnStart" bson:"updateOnStart"`
	Branch        string `json:"branch" bson:"branch"`
	WorkerThreads int    `json:"workerThreads" bson:"workerThreads"`

	AutoRestart bool `json:"autoRestart" bson:"autoRestart"`

	// Settings for Server Ini files
	MaxPlayers            int  `json:"maxPlayers" bson:"maxPlayers"`
	AutoPause             bool `json:"autoPause" bson:"autoPause"`
	AutoSaveOnDisconnect  bool `json:"autoSaveOnDisconnect" bson:"autoSaveOnDisconnect"`
	AutoSaveInterval      int  `json:"autoSaveInterval" bson:"autoSaveInterval"`
	DisableSeasonalEvents bool `json:"disableSeasonalEvents" bson:"disableSeasonalEvents"`
}

// Map Data

type AgentMapData struct {
	Players   []AgentMapDataPlayer   `json:"players" bson:"players"`
	Buildings []AgentMapDataBuilding `json:"buildings" bson:"buildings"`
}

type AgentMapDataPlayer struct {
	Username string          `json:"username" bson:"username"`
	Location models.Vector3F `json:"location" bson:"location"`
	Online   bool            `json:"online" bson:"online"`
}

type AgentMapDataBuilding struct {
	Name        string             `json:"name" bson:"name"`
	Class       string             `json:"class" bson:"class"`
	Location    models.Vector3F    `json:"location" bson:"location"`
	Rotation    float32            `json:"rotation" bson:"rotation"`
	BoundingBox models.BoundingBox `json:"boundingBox" bson:"boundingBox"`
}

// Save Data

type AgentSave struct {
	UUID     string    `json:"uuid" bson:"uuid"`
	FileName string    `json:"fileName" bson:"fileName"`
	Size     int64     `json:"size" bson:"size"`
	FileUrl  string    `json:"fileUrl" bson:"fileUrl"`
	ModTime  time.Time `json:"modTime" bson:"modTime"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

// Backup Data

type AgentBackup struct {
	UUID      string    `json:"uuid" bson:"uuid"`
	FileName  string    `json:"fileName" bson:"fileName"`
	Size      int64     `json:"size" bson:"size"`
	FileUrl   string    `json:"fileUrl" bson:"fileUrl"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type AgentLogSchema struct {
	ID            bson.ObjectID `json:"_id" bson:"_id"`
	FileName      string        `json:"fileName" bson:"fileName"`
	Type          string        `json:"type" bson:"type"`
	LogLines      []string      `json:"lines" bson:"lines"`
	FileURL       string        `json:"fileUrl" bson:"fileUrl"`
	PendingUpload bool          `json:"pendingUpload" bson:"pendingUpload"`
	CreatedAt     time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt" bson:"updatedAt"`
}

type AgentStatSchema struct {
	ID        bson.ObjectID `json:"_id" bson:"_id"`
	AgentId   bson.ObjectID `json:"agentId" bson:"agentId"`
	Running   bool          `json:"running" bson:"running"`
	CPU       float64       `json:"cpu" bson:"cpu"`
	MEM       float32       `json:"mem" bson:"mem"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
}

func NewAgent(agentName string, port int, memory int64, apiKey string) AgentSchema {

	if apiKey == "" {
		apiKey = "API-AGT-" + strings.ToUpper(utils.RandStringBytes(24))
	}

	newAgent := AgentSchema{
		ID:        bson.NewObjectID(),
		AgentName: agentName,
		APIKey:    apiKey,
		LogIds:    make(bson.A, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newAgent.Config.Port = port
	newAgent.Config.Memory = memory

	newAgent.Config.BackupKeepAmount = 24
	newAgent.Config.BackupInterval = 1.0

	newAgent.ServerConfig.MaxPlayers = 4
	newAgent.ServerConfig.WorkerThreads = 20
	newAgent.ServerConfig.Branch = "public"
	newAgent.ServerConfig.UpdateOnStart = true
	newAgent.ServerConfig.AutoSaveInterval = 300
	newAgent.ServerConfig.AutoSaveOnDisconnect = true

	newAgent.MapData.Players = make([]AgentMapDataPlayer, 0)
	newAgent.MapData.Buildings = make([]AgentMapDataBuilding, 0)

	newAgent.Saves = make([]AgentSave, 0)
	newAgent.Backups = make([]AgentBackup, 0)

	return newAgent
}

func NewAgentStat(theAgent *AgentSchema, running bool, cpu float64, memory float32) *AgentStatSchema {
	return &AgentStatSchema{
		ID:        bson.NewObjectID(),
		AgentId:   theAgent.ID,
		Running:   running,
		CPU:       cpu,
		MEM:       memory,
		CreatedAt: time.Now(),
	}
}
