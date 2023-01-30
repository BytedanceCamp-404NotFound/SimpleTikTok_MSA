package BaseInterface

import (
	"net/http"

	"SimpleTikTok/external/baseinterface/internal/logic/BaseInterface"
	"SimpleTikTok/external/baseinterface/internal/svc"
	"SimpleTikTok/external/baseinterface/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserHandlerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := BaseInterface.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.User(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}