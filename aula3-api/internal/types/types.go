// Code generated by goctl. DO NOT EDIT.
package types

type UserInfo struct {
	Id    int64  `path:"id"`
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

type ListarUsuarioRequest struct {
}

type ListarUsuarioResponse struct {
	Data []UserInfo `json:"data"`
}
