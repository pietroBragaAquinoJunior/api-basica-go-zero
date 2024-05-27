package logic

import (
	"context"
	"go-zero-aula-3/aula3-api/internal/svc"
	"go-zero-aula-3/aula3-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListarUsuarioLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListarUsuarioLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListarUsuarioLogic {
	return &ListarUsuarioLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListarUsuarioLogic) ListarUsuario(req *types.ListarUsuarioRequest) (resp *types.ListarUsuarioResponse, err error) {
	infos := make([]types.UserInfo, 0)
	//if len(infos) == 0 {
	//	return nil, errors.New("ERRO!")
	//}
	infos = append(infos, types.UserInfo{Id: 1, Nome: "Piêtro Braga", Idade: 25})
	infos = append(infos, types.UserInfo{Id: 2, Nome: "Ana Silva", Idade: 30})
	return &types.ListarUsuarioResponse{
		Data: infos,
	}, nil
}
