package CommActionInterface

import (
	"context"

	"SimpleTikTok/external_api/commaction/internal/svc"
	"SimpleTikTok/external_api/commaction/internal/types"
	"SimpleTikTok/oprations/commonerror"
	"SimpleTikTok/oprations/mongodb"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommmentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommmentListLogic {
	return &CommmentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommmentListLogic) CommmentList(req *types.CommmentListHandlerRequest) (resp *types.CommmentListHandlerResponse, err error) {
	resp = new(types.CommmentListHandlerResponse)
	resp.StatusCode = 400
	var comments []types.Comment
	token := req.Token
	flag, _, err := tools.CheckToke(token)
	if !flag {
		logx.Errorf("[pkg]logic [func]CommmentList [msg]parse token failed, [err]%v", err)
		resp.StatusCode = int32(commonerror.CommonErr_PARSE_TOKEN_ERROR)
		resp.StatusMsg = "parse token failed"
		return resp, err
	}
	videoId := req.VideoId
	collection := mongodb.MongoDBCollection
	filter := bson.D{{
		Key:   "video_id",
		Value: videoId,
	}}
	opts := &options.FindOptions{}
	sortOption := bson.D{{
		Key:   "create_date",
		Value: -1,
	}, {
		Key:   "_id",
		Value: -1,
	}}
	opts.Sort = sortOption
	cur, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		logx.Errorf("[pkg]logic [func]CommmentList [msg]find comments failed, [err]%v", err)
		resp.StatusCode = int32(commonerror.CommonErr_DB_ERROR)
		resp.StatusMsg = "find comments failed"
		return resp, err
	}
	for cur.Next(context.Background()) {
		var comment types.Comment
		err = cur.Decode(&comment)
		if err != nil {
			logx.Errorf("[pkg]logic [func]CommmentList [msg]decode comment failed, [err]%v", err)
			resp.StatusCode = int32(commonerror.CommonErr_PARAMETER_FAILED)
			resp.StatusMsg = "decode comment failed"
			return resp, err
		}
		comments = append(comments, comment)

	}
	err = cur.Err()
	if err != nil {
		logx.Errorf("[pkg]logic [func]CommmentList [msg]cur has an error, [err]%v", err)
		resp.StatusCode = int32(commonerror.CommonErr_PARAMETER_FAILED)
		resp.StatusMsg = "cur has an error"
		return resp, err
	}
	cur.Close(context.Background())
	resp.StatusCode = int32(commonerror.CommonErr_STATUS_OK)
	resp.StatusMsg = "get commentList success"
	resp.CommentList = comments
	return resp, nil
}
