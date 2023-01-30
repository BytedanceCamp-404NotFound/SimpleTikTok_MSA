package CommActionInterface

import (
	"net/http"

	"SimpleTikTok/external/commaction/internal/logic/CommActionInterface"
	"SimpleTikTok/external/commaction/internal/svc"
	"SimpleTikTok/external/commaction/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteListRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteListRegisterHandlerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := CommActionInterface.NewFavoriteListRegisterLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteListRegister(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
