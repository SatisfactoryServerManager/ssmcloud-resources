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
		Id:        agent.ID.Hex(),
		AgentName: agent.AgentName,
		ApiKey:    agent.APIKey,
		CreatedAt: timestamppb.New(agent.CreatedAt),
		UpdatedAt: timestamppb.New(agent.UpdatedAt),
	}

	agentProto.Status = MapAgentStatusToProto(&agent.Status)
	agentProto.Config = MapAgentConfigToProto(&agent.Config)
	agentProto.ServerConfig = MapAgentServerConfigToProto(&agent.ServerConfig)
	agentProto.ModConfig = MapAgentModConfigToProto(&agent.ModConfig)

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

func MapModToProto(mod *models.ModSchema) *pb.Mod {

	pbModVersions := make([]*pb.ModVersion, 0, len(mod.Versions))

	for i := range mod.Versions {
		pbModVersions = append(pbModVersions, MapModVersionToProto(&mod.Versions[i]))
	}

	return &pb.Mod{
		Id:           mod.ID.Hex(),
		ModId:        mod.ModID,
		ModReference: mod.ModReference,
		Versions:     pbModVersions,
	}
}

func MapModVersionToProto(modVersion *models.ModVersion) *pb.ModVersion {

	pbModVersionTargets := make([]*pb.ModVersionTarget, 0, len(modVersion.Targets))

	for i := range modVersion.Targets {
		pbModVersionTargets = append(pbModVersionTargets, MapModVersionTargetToProto(&modVersion.Targets[i]))
	}

	return &pb.ModVersion{
		Version: modVersion.Version,
		Targets: pbModVersionTargets,
	}
}

func MapModVersionTargetToProto(modVersionTarget *models.ModVersionTarget) *pb.ModVersionTarget {
	return &pb.ModVersionTarget{
		TargetName: modVersionTarget.TargetName,
		Link:       modVersionTarget.Link,
	}
}
