package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"go.viam.com/robotcore/api"
	"go.viam.com/robotcore/robot"
	"go.viam.com/robotcore/robot/actions"
	"go.viam.com/robotcore/robot/web"
	"go.viam.com/robotcore/utils"

	_ "go.viam.com/robotcore/robots/eva" // load eva

	"github.com/edaniels/golog"
)

var logger = golog.NewDevelopmentLogger("armplay")

func init() {
	actions.RegisterAction("play", func(ctx context.Context, r api.Robot) {
		err := play(ctx, r)
		if err != nil {
			logger.Errorf("error playing: %s", err)
		}
	})
}

func play(ctx context.Context, r api.Robot) error {
	if len(r.ArmNames()) != 1 {
		return fmt.Errorf("need 1 arm name")
	}

	arm := r.ArmByName(r.ArmNames()[0])

	start, err := arm.CurrentJointPositions(ctx)
	if err != nil {
		return err
	}

	for i := 0; i < 180; i += 10 {
		start.Degrees[0] = float64(i)
		err := arm.MoveToJointPositions(ctx, start)
		if err != nil {
			return err
		}

		if !utils.SelectContextOrWait(ctx, time.Second) {
			return ctx.Err()
		}
	}

	return nil
}

func main() {
	utils.ContextualMain(mainWithArgs, logger)
}

func mainWithArgs(ctx context.Context, args []string, logger golog.Logger) (err error) {
	flag.Parse()

	cfg, err := api.ReadConfig(flag.Arg(0))
	if err != nil {
		return err
	}

	myRobot, err := robot.NewRobot(ctx, cfg, logger)
	if err != nil {
		return err
	}
	defer myRobot.Close()

	return web.RunWeb(ctx, myRobot, web.NewOptions(), logger)
}
