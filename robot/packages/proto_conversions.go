package packages

import (
	pb "go.viam.com/api/app/packages/v1"

	"go.viam.com/rdk/config"
)

// PackageTypeToProto converts a config PackageType to its proto equivalent
// This is required be because app/packages uses a PackageType enum but app/PackageConfig uses a string Type.
func PackageTypeToProto(t config.PackageType) *pb.PackageType {
	switch t {
	case "":
		// for backwards compatibility
		// TODO(pre-merge) create ticket for this
		fallthrough
	case config.PackageTypeMlModel:
		return pb.PackageType_PACKAGE_TYPE_ML_MODEL.Enum()
	case config.PackageTypeModule:
		return pb.PackageType_PACKAGE_TYPE_MODULE.Enum()
	default:
		return pb.PackageType_PACKAGE_TYPE_UNSPECIFIED.Enum()
	}
}
