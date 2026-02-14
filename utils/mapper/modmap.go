package mapper

import (
	models "github.com/SatisfactoryServerManager/ssmcloud-resources/models"
	pb "github.com/SatisfactoryServerManager/ssmcloud-resources/proto/generated/models"
)

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
