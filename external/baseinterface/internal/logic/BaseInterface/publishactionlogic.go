package BaseInterface

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"SimpleTikTok/external/baseinterface/internal/svc"
	"SimpleTikTok/external/baseinterface/internal/types"
	"SimpleTikTok/oprations/commonerror"
	minio "SimpleTikTok/oprations/minioconnect"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishActionLogic) PublishAction(req *types.PublishActionHandlerRequest, r *http.Request) (resp *types.PublishActionHandlerResponse, err error) {
	ok, userId, err := tools.CheckToke(req.Token)
	if err != nil {
		return &types.PublishActionHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "Token校验出错",
		}, nil
	}
	if !ok {
		logx.Infof("[pkg]logic [func]PublishAction [msg]feedUserInfo.Name is nuil ")
		return &types.PublishActionHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARAMETER_FAILED),
			StatusMsg:  "登录过期，请重新登陆",
		}, nil
	}

	minioVideoUrl, minioPictureUrl, err := minioUpDate(r) // Minio 上传文件
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishAction [msg]minioUpDate is fail [err]%v", err)
		return &types.PublishActionHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PAGE_NOT_EXIT),
			StatusMsg:  "没有收到视频文件或者出现其他错误",
		}, err
	}

	VideoInfo := &mysqlconnect.PublishActionVideoInfo{
		Video_id:       int32(uuid.New().ID()),
		Author_id:      int64(userId),
		Play_url:       minioVideoUrl,
		Cover_url:      minioPictureUrl,
		Favorite_count: 0,
		Comment_count:  0,
		Video_title:    req.Title,
	}

	// gorm创建一条信息
	err = mysqlconnect.CreatePublishActionViedeInfo(VideoInfo)
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishAction [msg]CreatePublishActionViedeInfo is err [err]%v", err)
		return nil, err
	}
	return &types.PublishActionHandlerResponse{
		StatusCode: int32(commonerror.CommonErr_STATUS_OK),
		StatusMsg:  "上传成功",
	}, err
}

// 图片和视频上传到minio
func minioUpDate(r *http.Request) (string, string, error) {
	// TOTEMP
	exePath, _ := os.Executable()
	sourceFile := filepath.Dir(filepath.Dir(exePath))
	//vidoeFile := fmt.Sprintf("%s/source/video/video_test1.mp4", sourceFile)
	pictureFile := fmt.Sprintf("%s/source/pic/pic_test2.jpg", sourceFile)
	// content, err := os.ReadFile(vidoeFile)

	bucketName := "test-minio"
	minioClient, err := minio.MinioConnect()
	if err != nil {
		logx.Errorf("[pkg]logic [func]minioUpDate [msg]MinioConnect fail [err]%v", err)
		return "", "", err
	}

	file, FileHeader, err := r.FormFile("data") //可以优化
	if err != nil {
		logx.Errorf("[pkg]logic [func]minioUpDate [msg]r.FormFile fail [err]%v", err)
		return "", "", err
	}

	//minioVideoUrl, err := minio.MinioFileUploader(minioClient, bucketName, "vidoeFile/", vidoeFile) //上传本地文件
	minioVideoUrl, err := minio.MinioFileUploader_byte(minioClient, bucketName, "vidoeFile/", FileHeader.Filename, file, FileHeader.Size) // 上传字节类型文件
	if err != nil {
		logx.Errorf("[pkg]logic [func]minioUpDate [msg]minio video upload to fail [err]%v", err)
		return "", "", err // TODO: 上传失败暂时中断接口
	}

	minioPictureUrl, err := minio.MinioFileUploader(minioClient, bucketName, "pictureFile/", pictureFile)
	if err != nil {
		logx.Errorf("[pkg]logic [func]minioUpDate [msg]minio picture upload to fail [err]%v", err)
		// return "", "", err // TODO: 上传图片失败不中断接口
	}
	return minioVideoUrl, minioPictureUrl, err
}
