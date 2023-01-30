package RelationFollowInterface

import (
	"net/http"

	"SimpleTikTok/external/relationfollow/internal/logic/RelationFollowInterface"
	"SimpleTikTok/external/relationfollow/internal/svc"
	"SimpleTikTok/external/relationfollow/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RelationFollowListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RelationFollowListHandlerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := RelationFollowInterface.NewRelationFollowListLogic(r.Context(), svcCtx)
		resp, err := l.RelationFollowList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
