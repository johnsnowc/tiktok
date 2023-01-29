package logic

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"image"
	"image/jpeg"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"genuine_douyin/apps/video/model"
	"genuine_douyin/apps/video/rpc/internal/svc"
	"genuine_douyin/apps/video/rpc/video"
	"github.com/gofrs/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishActionLogic) PublishAction(in *video.DouyinPublishActionRequest) (*video.DouyinPublishActionResponse, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return &video.DouyinPublishActionResponse{}, err
	}
	videoPath := strconv.Itoa(int(in.UserId)) + "/" + u.String() + "." + "mp4"

	//上传视频
	_, err = l.svcCtx.MinioClient.PutObject(l.ctx, l.svcCtx.Config.Minio.VideoBucket, videoPath, bytes.NewReader(in.Data), int64(len(in.Data)), minio.PutObjectOptions{
		ContentType: "application/octet-stream"})
	if err != nil {
		l.Logger.Errorf("upload %s of size %d failed, %s", l.svcCtx.Config.Minio.VideoBucket, int64(len(in.Data)), err)
		return &video.DouyinPublishActionResponse{}, err
	}

	//获得链接
	reqParams := make(url.Values)
	expires := time.Second * 60 * 60 * 24
	videoUrl, err := l.svcCtx.MinioClient.PresignedGetObject(l.ctx, l.svcCtx.Config.Minio.VideoBucket, videoPath, expires, reqParams)
	if err != nil {
		l.Logger.Errorf("get url of file %s from bucket %s failed, %s", videoPath, l.svcCtx.Config.Minio.VideoBucket, err)
	}
	playUrl := strings.Split(videoUrl.String(), "?")[0]
	if err != nil {
		return &video.DouyinPublishActionResponse{}, err
	}

	// 获取封面
	coverPath := strconv.Itoa(int(in.UserId)) + "/" + u.String() + "." + "jpg"
	coverData, err := readFrameAsJpeg(playUrl)
	if err != nil {
		return &video.DouyinPublishActionResponse{}, err
	}

	//上传封面
	_, err = l.svcCtx.MinioClient.PutObject(l.ctx, l.svcCtx.Config.Minio.CoverBucket, coverPath, bytes.NewReader(coverData), int64(len(coverData)), minio.PutObjectOptions{
		ContentType: "application/octet-stream"})
	if err != nil {
		l.Logger.Errorf("upload %s of size %d failed, %s", l.svcCtx.Config.Minio.CoverBucket, int64(len(coverData)), err)
		return &video.DouyinPublishActionResponse{}, err
	}

	//获得链接
	reqParams = make(url.Values)
	coverUrl, err := l.svcCtx.MinioClient.PresignedGetObject(l.ctx, l.svcCtx.Config.Minio.CoverBucket, coverPath, expires, reqParams)
	if err != nil {
		l.Logger.Errorf("get url of file %s from bucket %s failed, %s", coverPath, l.svcCtx.Config.Minio.CoverBucket, err)
	}
	CoverUrl := strings.Split(coverUrl.String(), "?")[0]
	if err != nil {
		return &video.DouyinPublishActionResponse{}, err
	}

	now := time.Now().Unix()
	newVideo := model.Video{
		Uid:        in.UserId,
		PlayUrl:    playUrl,
		CoverUrl:   CoverUrl,
		Title:      in.Title,
		CreateTime: now,
		UpdateTime: now,
	}
	_, err = l.svcCtx.VideoModel.Insert(l.ctx, &newVideo)
	if err != nil {
		return &video.DouyinPublishActionResponse{}, err
	}

	return &video.DouyinPublishActionResponse{}, nil
}

// 从视频流中截取一帧并返回 需要在本地环境中安装ffmpeg并将bin添加到环境变量
func readFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}
