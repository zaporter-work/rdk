// Package main is the work-in-progress robotic land rover from Viam.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"image"
	"time"

	"go.viam.com/robotcore/api"
	"go.viam.com/robotcore/artifact"
	"go.viam.com/robotcore/board"
	"go.viam.com/robotcore/lidar"
	"go.viam.com/robotcore/rimage"
	"go.viam.com/robotcore/robot"
	"go.viam.com/robotcore/robot/action"
	"go.viam.com/robotcore/robot/web"
	"go.viam.com/robotcore/utils"
	"go.viam.com/robotcore/vision/segmentation"

	_ "go.viam.com/robotcore/board/detector"
	_ "go.viam.com/robotcore/rimage/imagesource"

	"github.com/edaniels/golog"
	"github.com/edaniels/gostream"
	"github.com/erh/egoutil"
	"go.opencensus.io/trace"
	"go.uber.org/multierr"
)

const (
	PanCenter  = 94
	TiltCenter = 113
)

var logger = golog.NewDevelopmentLogger("minirover")

func init() {
	action.RegisterAction("dock", func(ctx context.Context, r api.Robot) {
		err := dock(ctx, r)
		if err != nil {
			logger.Errorf("error docking: %s", err)
		}
	})
}

func dock(ctx context.Context, r api.Robot) error {
	logger.Info("docking started")

	cam := r.CameraByName("back")
	if cam == nil {
		return errors.New("no back camera")
	}

	base := r.BaseByName("pierre")
	if base == nil {
		return errors.New("no pierre")
	}

	theLidar := r.LidarByName("lidarOnBase")
	if theLidar == nil {
		return errors.New("no lidar lidarOnBase")
	}

	for tries := 0; tries < 5; tries++ {

		ms, err := theLidar.Scan(ctx, lidar.ScanOptions{})
		if err != nil {
			return err
		}
		back := ms.ClosestToDegree(180)
		logger.Debugf("distance to back: %#v", back)

		if back.Distance() < .55 {
			logger.Info("docking close enough")
			return nil
		}

		x, y, err := dockNextMoveCompute(ctx, cam)
		if err != nil {
			logger.Infof("failed to compute, will try again: %s", err)
			continue
		}
		logger.Debugf("x: %v y: %v\n", x, y)

		angle := x * -15
		logger.Debugf("turning %v degrees", angle)
		_, err = base.Spin(ctx, angle, 10, true)
		if err != nil {
			return err
		}

		amount := 100.0 - (100.0 * y)
		logger.Debugf("moved %v millis", amount)
		_, err = base.MoveStraight(ctx, int(-1*amount), 50, true)
		if err != nil {
			return err
		}

		tries = 0
	}

	return errors.New("failed to dock")
}

// return delta x, delta y, error
func dockNextMoveCompute(ctx context.Context, cam gostream.ImageSource) (float64, float64, error) {
	ctx, span := trace.StartSpan(ctx, "dockNextMoveCompute")
	defer span.End()

	logger.Debug("dock - next")
	img, closer, err := cam.Next(ctx)
	if err != nil {
		return 0, 0, err
	}
	defer closer()

	logger.Debug("dock - convert")
	ri := rimage.ConvertImage(img)

	logger.Debug("dock - findBlack")
	p, _, err := findBlack(ctx, ri, nil)
	if err != nil {
		return 0, 0, err
	}

	logger.Debugf("black: %v", p)

	x := 2 * float64((ri.Width()/2)-p.X) / float64(ri.Width())
	y := 2 * float64((ri.Height()/2)-p.Y) / float64(ri.Height())
	return x, y, nil
}

func findTopInSegment(img *segmentation.SegmentedImage, segment int) image.Point {
	mid := img.Width() / 2
	for y := 0; y < img.Height(); y++ {
		for x := mid; x < img.Width(); x++ {
			p := image.Point{x, y}
			s := img.GetSegment(p)
			if s == segment {
				return p
			}

			p = image.Point{mid - (x - mid), y}
			s = img.GetSegment(p)
			if s == segment {
				return p
			}

		}
	}
	return image.Point{0, 0}
}

func findBlack(ctx context.Context, img *rimage.Image, logger golog.Logger) (image.Point, image.Image, error) {
	_, span := trace.StartSpan(ctx, "findBlack")
	defer span.End()

	for x := 1; x < img.Width(); x += 3 {
		for y := 1; y < img.Height(); y += 3 {
			c := img.GetXY(x, y)
			if c.Distance(rimage.Black) > 1 {
				continue
			}

			x, err := segmentation.ShapeWalk(rimage.ConvertToImageWithDepth(img),
				image.Point{x, y},
				segmentation.ShapeWalkOptions{
					SkipCleaning: true,
					//Debug: true,
				},
				logger,
			)
			if err != nil {
				return image.Point{}, nil, err
			}

			if x.PixelsInSegmemnt(1) > 10000 {
				return findTopInSegment(x, 1), x, nil
			}
		}
	}

	return image.Point{}, nil, errors.New("no black found")
}

// ------

type Rover struct {
	theBoard  board.Board
	pan, tilt board.Servo
}

func (r *Rover) neckCenter(ctx context.Context) error {
	return r.neckPosition(ctx, PanCenter, TiltCenter)
}

func (r *Rover) neckOffset(ctx context.Context, left int) error {
	return r.neckPosition(ctx, uint8(PanCenter+(left*-30)), uint8(TiltCenter-20))
}

func (r *Rover) neckPosition(ctx context.Context, pan, tilt uint8) error {
	logger.Debugf("neckPosition to %v %v", pan, tilt)
	return multierr.Combine(r.pan.Move(ctx, pan), r.tilt.Move(ctx, tilt))
}

func (r *Rover) Ready(ctx context.Context, theRobot api.Robot) error {
	logger.Debug("minirover2 Ready called")
	cam := theRobot.CameraByName("front")
	if cam == nil {
		return errors.New("no camera named front")
	}

	// doing this in a goroutine so i can see camera and servo data in web ui, but probably not right long term
	if false {
		utils.PanicCapturingGo(func() {
			for {
				if !utils.SelectContextOrWait(ctx, time.Second) {
					return
				}
				var depthErr bool
				func() {
					img, release, err := cam.Next(ctx)
					if err != nil {
						logger.Debugf("error from camera %s", err)
						return
					}
					defer release()
					pc := rimage.ConvertToImageWithDepth(img)
					if pc.Depth == nil {
						logger.Warn("no depth data")
						depthErr = true
						return
					}
					err = pc.WriteTo(artifact.MustNewPath(fmt.Sprintf("samples/minirover/rover-centering-%d.both.gz", time.Now().Unix())))
					if err != nil {
						logger.Debugf("error writing %s", err)
					}
				}()
				if depthErr {
					return
				}
			}
		})
	}

	return nil
}

func NewRover(ctx context.Context, r api.Robot, theBoard board.Board) (*Rover, error) {
	rover := &Rover{theBoard: theBoard}
	rover.pan = theBoard.Servo("pan")
	rover.tilt = theBoard.Servo("tilt")

	if false {
		utils.PanicCapturingGo(func() {
			for {
				if !utils.SelectContextOrWait(ctx, 1500*time.Millisecond) {
					return
				}
				err := rover.neckCenter(ctx)
				if err != nil {
					panic(err)
				}

				if !utils.SelectContextOrWait(ctx, 1500*time.Millisecond) {
					return
				}

				err = rover.neckOffset(ctx, -1)
				if err != nil {
					panic(err)
				}

				if !utils.SelectContextOrWait(ctx, 1500*time.Millisecond) {
					return
				}

				err = rover.neckOffset(ctx, 1)
				if err != nil {
					panic(err)
				}

			}
		})
	} else {
		err := rover.neckCenter(ctx)
		if err != nil {
			return nil, err
		}
	}

	theLidar := r.LidarByName("lidarBase")
	if false && theLidar != nil {
		utils.PanicCapturingGo(func() {
			for {
				if !utils.SelectContextOrWait(ctx, time.Second) {
					return
				}

				ms, err := theLidar.Scan(ctx, lidar.ScanOptions{})
				if err != nil {
					logger.Infof("theLidar scan failed: %s", err)
					continue
				}

				logger.Debugf("fowrad distance %#v", ms[0])
			}
		})
	}

	logger.Debug("rover ready")

	return rover, nil
}

func main() {
	utils.ContextualMain(mainWithArgs, logger)
}

func mainWithArgs(ctx context.Context, args []string, logger golog.Logger) (err error) {
	flag.Parse()

	exp := egoutil.NewNiceLoggingSpanExporter()
	trace.RegisterExporter(exp)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	cfg, err := api.ReadConfig("samples/minirover/config.json")
	if err != nil {
		return err
	}

	myRobot, err := robot.NewRobot(ctx, cfg, logger)
	if err != nil {
		return err
	}
	defer myRobot.Close()

	rover, err := NewRover(ctx, myRobot, myRobot.BoardByName("local"))
	if err != nil {
		return err
	}
	err = rover.Ready(ctx, myRobot)
	if err != nil {
		return err
	}

	options := web.NewOptions()
	options.AutoTile = false
	options.Pprof = true
	return web.RunWeb(ctx, myRobot, options, logger)
}
