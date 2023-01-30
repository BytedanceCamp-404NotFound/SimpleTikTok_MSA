package BaseInterface

import (
	"net/http"

	"SimpleTikTok/external_api/baseinterface/internal/logic/BaseInterface"
	"SimpleTikTok/external_api/baseinterface/internal/svc"
	"SimpleTikTok/external_api/baseinterface/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserloginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserloginHandlerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := BaseInterface.NewUserloginLogic(r.Context(), svcCtx)
		resp, err := l.Userlogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
