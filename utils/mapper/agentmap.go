package mapper

import (
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

	agentProto.Status = &pb.AgentStatus{
		Online:             agent.Status.Online,
		Installed:          agent.Status.Installed,
		Running:            agent.Status.Running,
		Cpu:                float32(agent.Status.CPU),
		Ram:                float32(agent.Status.RAM),
		InstalledSfVersion: agent.Status.InstalledSFVersion,
		LatestSfVersion:    agent.Status.LatestSFVersion,
		LastCommDate:       timestamppb.New(agent.Status.LastCommDate),
	}

	agentProto.Config = &pb.AgentConfig{
		Version:          agent.Config.Version,
		Port:             int32(agent.Config.Port),
		Memory:           agent.Config.Memory,
		IpAddress:        agent.Config.IP,
		BackupKeepAmount: int32(agent.Config.BackupKeepAmount),
		BackupInterval:   int32(agent.Config.BackupInterval),
	}

	agentProto.ServerConfig = &pb.AgentServerConfig{
		UpdateOnStart:         wrapperspb.Bool(agent.ServerConfig.UpdateOnStart),
		Branch:                agent.ServerConfig.Branch,
		WorkerThreads:         int32(agent.ServerConfig.WorkerThreads),
		AutoRestart:           wrapperspb.Bool(agent.ServerConfig.AutoRestart),
		MaxPlayers:            int32(agent.ServerConfig.MaxPlayers),
		AutoPause:             wrapperspb.Bool(agent.ServerConfig.AutoPause),
		AutoSaveOnDisconnect:  wrapperspb.Bool(agent.ServerConfig.AutoSaveOnDisconnect),
		AutoSaveInterval:      int32(agent.ServerConfig.AutoSaveInterval),
		DisableSeasonalEvents: wrapperspb.Bool(agent.ServerConfig.DisableSeasonalEvents),
	}

	return agentProto
}
