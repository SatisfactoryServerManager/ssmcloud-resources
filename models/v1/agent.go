package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/SatisfactoryServerManager/ssmcloud-resources/models"
	"github.com/SatisfactoryServerManager/ssmcloud-resources/utils"
	"github.com/mrhid6/go-mongoose/mongoose"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Agents struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	AgentName string             `json:"agentName" bson:"agentName"`
	APIKey    string             `json:"apiKey" bson:"apiKey"`
	Status    AgentStatus        `json:"status" bson:"status"`

	Config       AgentConfig       `json:"config" bson:"config"`
	ServerConfig AgentServerConfig `json:"serverConfig" bson:"serverConfig"`

	MapData AgentMapData `json:"mapData" bson:"mapData"`

	Saves   []AgentSave   `json:"saves" bson:"saves"`
	Backups []AgentBackup `json:"backups" bson:"backups"`

	Tasks []AgentTask `json:"tasks" bson:"tasks"`

	Logs       primitive.A `json:"-" bson:"logs" mson:"collection=agentlogs"`
	LogObjects []AgentLogs `json:"logs" bson:"-"`

	Stats       primitive.A `json:"-" bson:"stats" mson:"collection=agentstats"`
	StatObjects []AgentStat `json:"stats" bson:"-"`

	ModConfig AgentModConfig `json:"modConfig" bson:"modConfig"`

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

// Mod Data

type AgentModConfig struct {
	SelectedMods []AgentModConfigSelectedMod `json:"selectedMods" bson:"selectedMods"`
}

type AgentModConfigSelectedMod struct {
	Mod              primitive.ObjectID `json:"-" bson:"mod" mson:"collection=mods"`
	ModObject        models.Mods        `json:"mod" bson:"-"`
	DesiredVersion   string             `json:"desiredVersion" bson:"desiredVersion"`
	InstalledVersion string             `json:"installedVersion" bson:"installedVersion"`
	Installed        bool               `json:"installed" bson:"installed"`
	NeedsUpdate      bool               `json:"needsUpdate" bson:"needsUpdate"`
	Config           string             `json:"config" bson:"config"`
}

// Task Data

type AgentTask struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Action    string             `json:"action" bson:"action"`
	Data      interface{}        `json:"data" bson:"data"`
	Completed bool               `json:"completed" bson:"completed"`
	Retries   int                `json:"retries" bson:"retries"`
}

type AgentLogs struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	FileName  string             `json:"fileName" bson:"fileName"`
	Type      string             `json:"type" bson:"type"`
	Snippet   string             `json:"snippet" bson:"snippet"`
	FileURL   string             `json:"fileUrl" bson:"fileUrl"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type AgentStat struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Running   bool               `json:"running" bson:"running"`
	CPU       float64            `json:"cpu" bson:"cpu"`
	MEM       float32            `json:"mem" bson:"mem"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

func (obj *Agents) AtomicDelete() error {

	if err := obj.PopulateLogs(); err != nil {
		return err
	}

	if err := obj.PopulateStats(); err != nil {
		return err
	}

	for i := range obj.LogObjects {
		log := &obj.LogObjects[i]

		fmt.Printf("** deleting agent log: %s\n", log.Type)
		if err := log.AtomicDelete(); err != nil {
			return err
		}
	}

	for i := range obj.StatObjects {
		stat := &obj.StatObjects[i]

		fmt.Printf("** deleting agent stat: %s\n", stat.ID.Hex())
		if err := stat.AtomicDelete(); err != nil {
			return err
		}
	}

	if _, err := mongoose.DeleteOne(bson.M{"_id": obj.ID}, Agents{}); err != nil {
		return err
	}

	fmt.Printf("deleted agent: %s\n", obj.AgentName)

	return nil
}

func (obj *AgentLogs) AtomicDelete() error {

	if _, err := mongoose.DeleteOne(bson.M{"_id": obj.ID}, AgentLogs{}); err != nil {
		return err
	}

	return nil
}

func (obj *AgentStat) AtomicDelete() error {

	if _, err := mongoose.DeleteOne(bson.M{"_id": obj.ID}, AgentStat{}); err != nil {
		return err
	}

	return nil
}

func (obj *Agents) PopulateFromURLQuery(populateStrings []string) error {
	for _, popStr := range populateStrings {
		if popStr == "stats" {
			if err := obj.PopulateStats(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (obj *Agents) PopulateModConfig() {

	for idx := range obj.ModConfig.SelectedMods {
		selectedMod := &obj.ModConfig.SelectedMods[idx]
		selectedMod.PopulateMod()
	}
}

func (obj *Agents) PopulateLogs() error {
	err := mongoose.PopulateObjectArray(obj, "Logs", &obj.LogObjects)

	if err != nil {
		return err
	}

	if obj.LogObjects == nil {
		obj.LogObjects = make([]AgentLogs, 0)
	}

	return nil
}

func (obj *Agents) GetLogOfType(logType string) (*AgentLogs, error) {
	if err := obj.PopulateLogs(); err != nil {
		return nil, err
	}

	var theLog AgentLogs
	for _, log := range obj.LogObjects {
		if log.Type == logType {
			theLog = log
			break
		}
	}

	if theLog.ID.IsZero() {
		return nil, fmt.Errorf("cant find log of type: %s", logType)
	}

	return &theLog, nil
}

func (obj *Agents) PopulateStats() error {
	err := mongoose.PopulateObjectArray(obj, "Stats", &obj.StatObjects)

	if err != nil {
		return err
	}

	if obj.StatObjects == nil {
		obj.StatObjects = make([]AgentStat, 0)
	}

	return nil
}

func (obj *Agents) PurgeTasks() error {

	if len(obj.Tasks) == 0 {
		return nil
	}

	newTaskList := make([]AgentTask, 0)
	for _, task := range obj.Tasks {

		if task.Completed {
			continue
		}

		newTaskList = append(newTaskList, task)
	}

	if len(obj.Tasks) != len(newTaskList) {

		dbUpdate := bson.M{
			"tasks":     newTaskList,
			"updatedAt": time.Now(),
		}

		if err := mongoose.UpdateModelData(*obj, dbUpdate); err != nil {
			return err
		}
	}

	return nil
}

func (obj *AgentModConfigSelectedMod) PopulateMod() {
	mongoose.PopulateObject(obj, "Mod", &obj.ModObject)
}

func (obj *Agents) CreateStat(running bool, cpu float64, mem float32) error {
	newStat := AgentStat{
		ID:        primitive.NewObjectID(),
		CPU:       cpu,
		MEM:       mem,
		Running:   running,
		CreatedAt: time.Now(),
	}

	_, err := mongoose.InsertOne(&newStat)
	if err != nil {
		return err
	}

	obj.Stats = append(obj.Stats, newStat.ID)

	dbUpdate := bson.M{
		"stats": obj.Stats,
	}

	if err := mongoose.UpdateModelData(*obj, dbUpdate); err != nil {
		return err
	}

	return nil
}

func (obj *Agents) PurgeStats() error {
	if err := obj.PopulateStats(); err != nil {
		return err
	}

	if len(obj.Stats) == 0 {
		return nil
	}

	now := time.Now()
	expiry := now.AddDate(0, 0, -3)

	newStats := make(primitive.A, 0)
	deleteStats := make([]AgentStat, 0)

	for idx := range obj.StatObjects {
		stat := obj.StatObjects[idx]

		if stat.CreatedAt.After(expiry) {
			newStats = append(newStats, stat.ID)
		} else {
			deleteStats = append(deleteStats, stat)
		}
	}

	if len(obj.Stats) != len(newStats) || len(deleteStats) > 0 {
		obj.Stats = newStats

		dbUpdate := bson.M{
			"stats": obj.Stats,
		}

		if err := mongoose.UpdateModelData(*obj, dbUpdate); err != nil {
			return err
		}

		for _, stat := range deleteStats {
			if _, err := mongoose.DeleteOne(bson.M{"_id": stat.ID}, stat); err != nil {
				return err
			}
		}
	}

	return nil
}

func NewAgent(agentName string, port int, memory int64, apiKey string) Agents {

	if apiKey == "" {
		apiKey = "API-AGT-" + strings.ToUpper(utils.RandStringBytes(24))
	}

	newAgent := Agents{
		ID:        primitive.NewObjectID(),
		AgentName: agentName,
		APIKey:    apiKey,
		Tasks:     make([]AgentTask, 0),
		Logs:      make(primitive.A, 0),
		Stats:     make(primitive.A, 0),
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
	newAgent.Tasks = make([]AgentTask, 0)

	newAgent.ModConfig.SelectedMods = make([]AgentModConfigSelectedMod, 0)

	return newAgent
}

func NewAgentTask(action string, data interface{}) AgentTask {
	return AgentTask{
		ID:     primitive.NewObjectID(),
		Action: action,
		Data:   data,
	}
}
