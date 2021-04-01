package api

import (
	"context"

	"go.viam.com/robotcore/board"
	"go.viam.com/robotcore/lidar"
	pb "go.viam.com/robotcore/proto/api/v1"

	"github.com/edaniels/golog"
	"github.com/edaniels/gostream"
)

type Robot interface {
	// providers are for singletons for a whole model
	ProviderByModel(model string) Provider
	AddProvider(p Provider, c Component)

	RemoteByName(name string) Robot
	ArmByName(name string) Arm
	BaseByName(name string) Base
	GripperByName(name string) Gripper
	CameraByName(name string) gostream.ImageSource
	LidarDeviceByName(name string) lidar.Device
	BoardByName(name string) board.Board

	RemoteNames() []string
	ArmNames() []string
	GripperNames() []string
	CameraNames() []string
	LidarDeviceNames() []string
	BaseNames() []string
	BoardNames() []string

	// this is allowed to be partial or empty
	GetConfig(ctx context.Context) (Config, error)

	// use CreateStatus helper in most cases
	Status(ctx context.Context) (*pb.Status, error)

	Logger() golog.Logger
	Close(ctx context.Context) error
}

type Provider interface {
	Ready(r Robot) error
}
