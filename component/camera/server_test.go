package camera_test

import (
	"bytes"
	"context"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"testing"

	"go.viam.com/test"

	"go.viam.com/rdk/component/camera"
	"go.viam.com/rdk/pointcloud"
	pb "go.viam.com/rdk/proto/api/component/v1"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/subtype"
	"go.viam.com/rdk/testutils/inject"
)

func newServer() (pb.CameraServiceServer, *inject.Camera, *inject.Camera, error) {
	injectCamera := &inject.Camera{}
	injectCamera2 := &inject.Camera{}
	cameras := map[resource.Name]interface{}{
		camera.Named(testCameraName): injectCamera,
		camera.Named(failCameraName): injectCamera2,
		camera.Named(fakeCameraName): "notCamera",
	}
	cameraSvc, err := subtype.New(cameras)
	if err != nil {
		return nil, nil, nil, err
	}
	return camera.NewServer(cameraSvc), injectCamera, injectCamera2, nil
}

func TestServer(t *testing.T) {
	cameraServer, injectCamera, injectCamera2, err := newServer()
	test.That(t, err, test.ShouldBeNil)

	img := image.NewNRGBA(image.Rect(0, 0, 4, 4))
	var imgBuf bytes.Buffer
	test.That(t, png.Encode(&imgBuf, img), test.ShouldBeNil)
	var imgBufJpeg bytes.Buffer
	test.That(t, jpeg.Encode(&imgBufJpeg, img, nil), test.ShouldBeNil)

	pcA := pointcloud.New()
	err = pcA.Set(pointcloud.NewBasicPoint(5, 5, 5))
	test.That(t, err, test.ShouldBeNil)

	var imageReleased bool
	injectCamera.NextFunc = func(ctx context.Context) (image.Image, func(), error) {
		return img, func() { imageReleased = true }, nil
	}
	injectCamera.NextPointCloudFunc = func(ctx context.Context) (pointcloud.PointCloud, error) {
		return pcA, nil
	}

	injectCamera2.NextFunc = func(ctx context.Context) (image.Image, func(), error) {
		return nil, nil, errors.New("can't generate next frame")
	}
	injectCamera2.NextPointCloudFunc = func(ctx context.Context) (pointcloud.PointCloud, error) {
		return nil, errors.New("can't generate next point cloud")
	}
	t.Run("GetFrame", func(t *testing.T) {
		_, err := cameraServer.GetFrame(context.Background(), &pb.CameraServiceGetFrameRequest{Name: missingCameraName})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "no camera")

		_, err = cameraServer.GetFrame(context.Background(), &pb.CameraServiceGetFrameRequest{Name: fakeCameraName})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "not a camera")

		resp, err := cameraServer.GetFrame(context.Background(), &pb.CameraServiceGetFrameRequest{Name: testCameraName})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, imageReleased, test.ShouldBeTrue)
		test.That(t, resp.MimeType, test.ShouldEqual, "image/raw-rgba")
		test.That(t, resp.Frame, test.ShouldResemble, img.Pix)

		imageReleased = false
		resp, err = cameraServer.GetFrame(context.Background(), &pb.CameraServiceGetFrameRequest{
			Name:     testCameraName,
			MimeType: "image/png",
		})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, imageReleased, test.ShouldBeTrue)
		test.That(t, resp.MimeType, test.ShouldEqual, "image/png")
		test.That(t, resp.Frame, test.ShouldResemble, imgBuf.Bytes())

		imageReleased = false
		_, err = cameraServer.GetFrame(context.Background(), &pb.CameraServiceGetFrameRequest{
			Name:     testCameraName,
			MimeType: "image/who",
		})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "do not know how")
		test.That(t, imageReleased, test.ShouldBeTrue)

		_, err = cameraServer.GetFrame(context.Background(), &pb.CameraServiceGetFrameRequest{Name: failCameraName})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "can't generate next frame")
	})

	t.Run("RenderFrame", func(t *testing.T) {
		_, err := cameraServer.RenderFrame(context.Background(), &pb.CameraServiceRenderFrameRequest{Name: missingCameraName})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "no camera")

		resp, err := cameraServer.RenderFrame(context.Background(), &pb.CameraServiceRenderFrameRequest{
			Name: testCameraName,
		})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, imageReleased, test.ShouldBeTrue)
		test.That(t, resp.ContentType, test.ShouldEqual, "image/jpeg")
		test.That(t, resp.Data, test.ShouldResemble, imgBufJpeg.Bytes())

		imageReleased = false
		resp, err = cameraServer.RenderFrame(context.Background(), &pb.CameraServiceRenderFrameRequest{
			Name:     testCameraName,
			MimeType: "image/png",
		})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, imageReleased, test.ShouldBeTrue)
		test.That(t, resp.ContentType, test.ShouldEqual, "image/png")
		test.That(t, resp.Data, test.ShouldResemble, imgBuf.Bytes())

		imageReleased = false
		_, err = cameraServer.RenderFrame(context.Background(), &pb.CameraServiceRenderFrameRequest{
			Name:     testCameraName,
			MimeType: "image/who",
		})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "do not know how")
		test.That(t, imageReleased, test.ShouldBeTrue)

		_, err = cameraServer.RenderFrame(context.Background(), &pb.CameraServiceRenderFrameRequest{Name: failCameraName})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "can't generate next frame")
	})

	t.Run("GetPointCloud", func(t *testing.T) {
		_, err := cameraServer.GetPointCloud(context.Background(), &pb.CameraServiceGetPointCloudRequest{Name: missingCameraName})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "no camera")

		pcA := pointcloud.New()
		err = pcA.Set(pointcloud.NewBasicPoint(5, 5, 5))
		test.That(t, err, test.ShouldBeNil)

		injectCamera.NextPointCloudFunc = func(ctx context.Context) (pointcloud.PointCloud, error) {
			return pcA, nil
		}
		_, err = cameraServer.GetPointCloud(context.Background(), &pb.CameraServiceGetPointCloudRequest{
			Name: testCameraName,
		})
		test.That(t, err, test.ShouldBeNil)

		_, err = cameraServer.GetPointCloud(context.Background(), &pb.CameraServiceGetPointCloudRequest{
			Name: failCameraName,
		})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "can't generate next point cloud")
	})

	t.Run("GetObjectPointClouds", func(t *testing.T) {
		_, err := cameraServer.GetObjectPointClouds(context.Background(), &pb.CameraServiceGetObjectPointCloudsRequest{Name: missingCameraName})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "no camera")

		// request the two segments in the point cloud
		pcA := pointcloud.New()
		err = pcA.Set(pointcloud.NewBasicPoint(5, 5, 5))
		test.That(t, err, test.ShouldBeNil)
		err = pcA.Set(pointcloud.NewBasicPoint(5, 5, 6))
		test.That(t, err, test.ShouldBeNil)
		err = pcA.Set(pointcloud.NewBasicPoint(5, 5, 4))
		test.That(t, err, test.ShouldBeNil)
		err = pcA.Set(pointcloud.NewBasicPoint(-5, -5, 5))
		test.That(t, err, test.ShouldBeNil)
		err = pcA.Set(pointcloud.NewBasicPoint(-5, -5, 6))
		test.That(t, err, test.ShouldBeNil)
		err = pcA.Set(pointcloud.NewBasicPoint(-5, -5, 4))
		test.That(t, err, test.ShouldBeNil)

		injectCamera.NextPointCloudFunc = func(ctx context.Context) (pointcloud.PointCloud, error) {
			return pcA, nil
		}
		segs, err := cameraServer.GetObjectPointClouds(context.Background(), &pb.CameraServiceGetObjectPointCloudsRequest{
			Name:               testCameraName,
			MinPointsInPlane:   100,
			MinPointsInSegment: 3,
			ClusteringRadiusMm: 5.,
		})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, len(segs.Objects), test.ShouldEqual, 2)
		test.That(t, segs.Objects[0].CenterCoordinatesMm.Z, test.ShouldEqual, 5.)
		test.That(t, segs.Objects[1].CenterCoordinatesMm.Z, test.ShouldEqual, 5.)
		test.That(t, segs.Objects[0].BoundingBoxMm.WidthMm, test.ShouldEqual, 0)
		test.That(t, segs.Objects[0].BoundingBoxMm.LengthMm, test.ShouldEqual, 0)
		test.That(t, segs.Objects[0].BoundingBoxMm.DepthMm, test.ShouldEqual, 2)
		test.That(t, segs.Objects[1].BoundingBoxMm.WidthMm, test.ShouldEqual, 0)
		test.That(t, segs.Objects[1].BoundingBoxMm.LengthMm, test.ShouldEqual, 0)
		test.That(t, segs.Objects[1].BoundingBoxMm.DepthMm, test.ShouldEqual, 2)

		// empty pointcloud
		pcB := pointcloud.New()

		injectCamera.NextPointCloudFunc = func(ctx context.Context) (pointcloud.PointCloud, error) {
			return pcB, nil
		}
		segs, err = cameraServer.GetObjectPointClouds(context.Background(), &pb.CameraServiceGetObjectPointCloudsRequest{
			Name:               testCameraName,
			MinPointsInPlane:   100,
			MinPointsInSegment: 3,
			ClusteringRadiusMm: 5.,
		})
		test.That(t, err, test.ShouldBeNil)
		test.That(t, len(segs.Objects), test.ShouldEqual, 0)

		_, err = cameraServer.GetObjectPointClouds(context.Background(), &pb.CameraServiceGetObjectPointCloudsRequest{
			Name:               failCameraName,
			MinPointsInPlane:   100,
			MinPointsInSegment: 3,
			ClusteringRadiusMm: 5.,
		})
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "can't generate next point cloud")
	})
}
