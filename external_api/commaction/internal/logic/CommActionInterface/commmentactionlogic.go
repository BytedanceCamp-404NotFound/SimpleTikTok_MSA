package CommActionInterface

import (
	"context"

	"SimpleTikTok/external_api/commaction/internal/svc"
	"SimpleTikTok/external_api/commaction/internal/types"
	"SimpleTikTok/oprations/commonerror"
	"SimpleTikTok/oprations/mongodb"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
)

type CommmentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommmentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommmentActionLogic {
	return &CommmentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommmentActionLogic) CommmentAction(req *types.CommmentActionHandlerRequest) (resp *types.CommmentActionHandlerResponse, err error) {
	//parse token
	resp = new(types.CommmentActionHandlerResponse)
	flag, userId, err := tools.CheckToke(req.Token)
	if !flag {
		logx.Errorf("[pkg]logic [func]CommentAction [msg]parse token failed, [err]%v", err)
		resp.StatusCode = int32(commonerror.CommonErr_PARSE_TOKEN_ERROR)
		resp.StatusMsg = "parse token failed"
		return resp, err
	}
	//get collection from mongodb
	collection := mongodb.MongoDBCollection
	actionType := req.ActionType
	videoId := req.VideoId
	if actionType == 2 {
		//delete comment
		commentId := req.CommentId
		filter := bson.D{{
			Key:   "_id",
			Value: commentId,
		},
			{
				Key:   "video_id",
				Value: videoId,
			}}
		_, err = collection.DeleteOne(context.Background(), filter)
		if err != nil {
			logx.Errorf("[pkg]logic [func]CommentAction [msg]delete comment failed, [err]%v", err)
			resp.StatusCode = int32(commonerror.CommonErr_DB_ERROR)
			resp.StatusMsg = "delete comment failed"
			return resp, err
		}
		resp.StatusCode = int32(commonerror.CommonErr_STATUS_OK)
		resp.StatusMsg = "delete success"
	} else {
		//insert comment
		comUser, err := mysqlconnect.CommentGetUserByUserId(userId)
		if err != nil {
			logx.Errorf("[pkg]logic [func]CommentAction [msg]search user_info failed, [err]%v", err)
			resp.StatusCode = int32(commonerror.CommonErr_DB_ERROR)
			resp.StatusMsg = "search user_info failed"
			return resp, err
		}
		user := types.User(comUser)
		content := req.CommentText
		date := time.Now()
		createDate := fmt.Sprintf("%d-%v", date.Month(), date.Day())
		id, err := mongodb.GetId(collection)
		if err != nil {
			logx.Errorf("[pkg]logic [func]CommentAction [msg]get id failed, [err]%v", err)
			resp.StatusCode = int32(commonerror.CommonErr_DB_ERROR)
			resp.StatusMsg = "get id failed"
			return resp, err
		}
		comment := types.Comment{
			Id:         id,
			VideoId:    videoId,
			User:       user,
			Content:    content,
			CreateDate: createDate,
		}
		_, err = collection.InsertOne(context.Background(), comment)
		if err != nil {
			logx.Errorf("[pkg]logic [func]CommentAction [msg]insert comment failed, [err]%v", err)
			resp.StatusCode = int32(commonerror.CommonErr_DB_ERROR)
			resp.StatusMsg = "insert comment failed"
			return resp, err
		}
		resp.StatusCode = int32(commonerror.CommonErr_STATUS_OK)
		resp.StatusMsg = "insert success"
		resp.Comment = comment
	}
	return
}
