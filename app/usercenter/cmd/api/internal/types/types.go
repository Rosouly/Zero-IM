// Code generated by goctl. DO NOT EDIT.
package types

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResp struct {
	Token string `json:"token"`
}
