package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-aula-3/aula3-api/internal/logic"
	"go-zero-aula-3/aula3-api/internal/svc"
	"go-zero-aula-3/aula3-api/internal/types"
)

func ListarUsuarioHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListarUsuarioRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListarUsuarioLogic(r.Context(), svcCtx)
		resp, err := l.ListarUsuario(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
