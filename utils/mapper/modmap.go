package mapper

import (
	"time"

	models "github.com/SatisfactoryServerManager/ssmcloud-resources/models"
	pb "github.com/SatisfactoryServerManager/ssmcloud-resources/proto/generated/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapModToProto(mod *models.ModSchema) *pb.Mod {

	pbModVersions := make([]*pb.ModVersion, 0, len(mod.Versions))

	for i := range mod.Versions {
		pbModVersions = append(pbModVersions, MapModVersionToProto(&mod.Versions[i]))
	}

	return &pb.Mod{
		Id:           mod.ID.Hex(),
		ModId:        mod.ModID,
		ModName:      mod.ModName,
		Hidden:       mod.Hidden,
		LogoUrl:      mod.LogoURL,
		Downloads:    int32(mod.Downloads),
		ModReference: mod.ModReference,
		Versions:     pbModVersions,
	}
}

func MapModVersionToProto(modVersion *models.ModVersion) *pb.ModVersion {

	pbModVersionTargets := make([]*pb.ModVersionTarget, 0, len(modVersion.Targets))

	for i := range modVersion.Targets {
		pbModVersionTargets = append(pbModVersionTargets, MapModVersionTargetToProto(&modVersion.Targets[i]))
	}

	createdAtTime, _ := time.Parse(time.RFC3339, modVersion.CreatedAt)

	pbModVersionDependencies := make([]*pb.ModVersionDependency, 0, len(modVersion.Dependencies))

	for i := range modVersion.Dependencies {
		pbModVersionDependencies = append(pbModVersionDependencies, MapModVersionDependencyToProto(modVersion.Dependencies[i]))
	}

	return &pb.ModVersion{
		Version:      modVersion.Version,
		Targets:      pbModVersionTargets,
		Dependencies: pbModVersionDependencies,
		CreatedAt:    timestamppb.New(createdAtTime),
	}
}

func MapModVersionTargetToProto(modVersionTarget *models.ModVersionTarget) *pb.ModVersionTarget {
	return &pb.ModVersionTarget{
		TargetName: modVersionTarget.TargetName,
		Link:       modVersionTarget.Link,
	}
}

func MapModVersionDependencyToProto(dependency models.ModVersionDependency) *pb.ModVersionDependency {
	return &pb.ModVersionDependency{
		ModReference: dependency.ModReference,
		Condition:    dependency.Condition,
		Optional:     dependency.Optional,
	}
}
