package mapper

import (
	models "github.com/SatisfactoryServerManager/ssmcloud-resources/models"
	v2 "github.com/SatisfactoryServerManager/ssmcloud-resources/models/v2"
	pb "github.com/SatisfactoryServerManager/ssmcloud-resources/proto/generated/models"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func MapAgentToProto(agent *v2.AgentSchema) *pb.Agent {
	agentProto := &pb.Agent{
		Id:                 agent.ID.Hex(),
		AgentName:          agent.AgentName,
		ApiKey:             agent.APIKey,
		LatestAgentVersion: agent.LatestAgentVersion,

		CreatedAt: timestamppb.New(agent.CreatedAt),
		UpdatedAt: timestamppb.New(agent.UpdatedAt),
	}

	agentProto.Status = MapAgentStatusToProto(&agent.Status)
	agentProto.Config = MapAgentConfigToProto(&agent.Config)
	agentProto.ServerConfig = MapAgentServerConfigToProto(&agent.ServerConfig)
	agentProto.ModConfig = MapAgentModConfigToProto(&agent.ModConfig)
	agentProto.Logs = MapAgentLogsToProto(agent.Logs)

	agentProto.Saves = MapAgentSavesToProto(agent.Saves)
	agentProto.Backups = MapAgentBackupsToProto(agent.Backups)
	agentProto.MapData = MapAgentMapDataToProto(&agent.MapData)

	return agentProto
}

func MapAgentStatusToProto(agentStatus *v2.AgentStatus) *pb.AgentStatus {
	return &pb.AgentStatus{
		Online:             agentStatus.Online,
		Installed:          agentStatus.Installed,
		Running:            agentStatus.Running,
		Cpu:                float32(agentStatus.CPU),
		Ram:                float32(agentStatus.RAM),
		InstalledSfVersion: agentStatus.InstalledSFVersion,
		LatestSfVersion:    agentStatus.LatestSFVersion,
		LastCommDate:       timestamppb.New(agentStatus.LastCommDate),
	}
}

func MapAgentConfigToProto(agentConfig *v2.AgentConfig) *pb.AgentConfig {
	return &pb.AgentConfig{
		Version:          agentConfig.Version,
		Port:             int32(agentConfig.Port),
		Memory:           agentConfig.Memory,
		IpAddress:        agentConfig.IP,
		BackupKeepAmount: int32(agentConfig.BackupKeepAmount),
		BackupInterval:   int32(agentConfig.BackupInterval),
	}
}

func MapAgentServerConfigToProto(agentServerConfig *v2.AgentServerConfig) *pb.AgentServerConfig {
	return &pb.AgentServerConfig{
		UpdateOnStart:         wrapperspb.Bool(agentServerConfig.UpdateOnStart),
		Branch:                agentServerConfig.Branch,
		WorkerThreads:         int32(agentServerConfig.WorkerThreads),
		AutoRestart:           wrapperspb.Bool(agentServerConfig.AutoRestart),
		MaxPlayers:            int32(agentServerConfig.MaxPlayers),
		AutoPause:             wrapperspb.Bool(agentServerConfig.AutoPause),
		AutoSaveOnDisconnect:  wrapperspb.Bool(agentServerConfig.AutoSaveOnDisconnect),
		AutoSaveInterval:      int32(agentServerConfig.AutoSaveInterval),
		DisableSeasonalEvents: wrapperspb.Bool(agentServerConfig.DisableSeasonalEvents),
	}
}

func MapAgentModConfigToProto(agentModConfig *v2.AgentModConfig) *pb.ModConfig {

	pdSelectedMods := make([]*pb.SelectedMod, 0, len(agentModConfig.SelectedMods))

	for i := range agentModConfig.SelectedMods {
		pdSelectedMods = append(pdSelectedMods, MapSelectedModToProto(&agentModConfig.SelectedMods[i]))
	}

	return &pb.ModConfig{
		SelectedMods: pdSelectedMods,
	}
}

func MapSelectedModToProto(selectedMod *v2.AgentModConfigSelectedModSchema) *pb.SelectedMod {
	return &pb.SelectedMod{
		Mod:              MapModToProto(&selectedMod.Mod),
		DesiredVersion:   selectedMod.DesiredVersion,
		InstalledVersion: selectedMod.InstalledVersion,
		Installed:        selectedMod.Installed,
		NeedsUpdate:      selectedMod.NeedsUpdate,
		Config:           selectedMod.Config,
	}
}

func MapAgentLogsToProto(logs []v2.AgentLogSchema) []*pb.AgentLog {
	pbLogs := make([]*pb.AgentLog, 0, len(logs))

	for i := range logs {
		pbLogs = append(pbLogs, MapAgentLogToProto(&logs[i]))
	}

	return pbLogs
}

func MapAgentLogToProto(log *v2.AgentLogSchema) *pb.AgentLog {
	return &pb.AgentLog{
		Id:            log.ID.Hex(),
		FileName:      log.FileName,
		Type:          log.Type,
		LogLines:      log.LogLines,
		FileUrl:       log.FileURL,
		PendingUpload: log.PendingUpload,
		CreatedAt:     timestamppb.New(log.CreatedAt),
		UpdatedAt:     timestamppb.New(log.UpdatedAt),
	}
}

func MapAgentSavesToProto(saves []v2.AgentSave) []*pb.AgentSave {
	pbSaves := make([]*pb.AgentSave, 0, len(saves))

	for i := range saves {
		pbSaves = append(pbSaves, MapAgentSaveToProto(&saves[i]))
	}

	return pbSaves
}

func MapAgentSaveToProto(save *v2.AgentSave) *pb.AgentSave {
	return &pb.AgentSave{
		Uuid:      save.UUID,
		FileName:  save.FileName,
		Size:      save.Size,
		FileUrl:   save.FileUrl,
		ModTime:   timestamppb.New(save.ModTime),
		CreatedAt: timestamppb.New(save.CreatedAt),
		UpdatedAt: timestamppb.New(save.UpdatedAt),
	}
}

func MapAgentBackupsToProto(backups []v2.AgentBackup) []*pb.AgentBackup {
	pbBackups := make([]*pb.AgentBackup, 0, len(backups))

	for i := range backups {
		pbBackups = append(pbBackups, MapAgentBackupToProto(&backups[i]))
	}

	return pbBackups
}

func MapAgentBackupToProto(backup *v2.AgentBackup) *pb.AgentBackup {
	return &pb.AgentBackup{
		Uuid:      backup.UUID,
		FileName:  backup.FileName,
		Size:      backup.Size,
		FileUrl:   backup.FileUrl,
		CreatedAt: timestamppb.New(backup.CreatedAt),
		UpdatedAt: timestamppb.New(backup.UpdatedAt),
	}
}

func MapAgentMapDataToProto(mapData *v2.AgentMapData) *pb.AgentMapData {
	if mapData == nil {
		return &pb.AgentMapData{}
	}

	pbPlayers := make([]*pb.AgentMapDataPlayer, 0, len(mapData.Players))
	for i := range mapData.Players {
		pbPlayers = append(pbPlayers, MapAgentMapDataPlayerToProto(&mapData.Players[i]))
	}

	pbBuildings := make([]*pb.AgentMapDataBuilding, 0, len(mapData.Buildings))
	for i := range mapData.Buildings {
		pbBuildings = append(pbBuildings, MapAgentMapDataBuildingToProto(&mapData.Buildings[i]))
	}

	return &pb.AgentMapData{
		Players:   pbPlayers,
		Buildings: pbBuildings,
	}
}

func MapAgentMapDataPlayerToProto(p *v2.AgentMapDataPlayer) *pb.AgentMapDataPlayer {
	return &pb.AgentMapDataPlayer{
		Username: p.Username,
		Location: MapVector3FToProto(p.Location),
		Online:   p.Online,
	}
}

func MapAgentMapDataBuildingToProto(b *v2.AgentMapDataBuilding) *pb.AgentMapDataBuilding {
	return &pb.AgentMapDataBuilding{
		Name:        b.Name,
		Class:       b.Class,
		Location:    MapVector3FToProto(b.Location),
		Rotation:    b.Rotation,
		BoundingBox: MapBoundingBoxToProto(b.BoundingBox),
	}
}

func MapVector3FToProto(v models.Vector3F) *pb.Vector3F {
	return &pb.Vector3F{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
	}
}

func MapBoundingBoxToProto(bb models.BoundingBox) *pb.BoundingBox {
	return &pb.BoundingBox{
		Min: MapVector3FToProto(bb.Min),
		Max: MapVector3FToProto(bb.Max),
	}
}

func MapAgentStatToProto(stat *v2.AgentStatSchema) *pb.AgentStat {
	return &pb.AgentStat{
		Id:        stat.ID.Hex(),
		AgentId:   stat.AgentId.Hex(),
		Running:   wrapperspb.Bool(stat.Running),
		Cpu:       float32(stat.CPU),
		Mem:       float32(stat.MEM),
		CreatedAt: timestamppb.New(stat.CreatedAt),
	}
}
