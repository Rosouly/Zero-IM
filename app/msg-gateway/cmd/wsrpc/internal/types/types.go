// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Token      string `json:"token" form:"token"`
	SendID     string `json:"sendId" form:"sendId"`
	PlatformID string `json:"platformId" form:"platformId"`
}

type Response struct {
	Uid     string `json:"uid"`
	ErrMsg  string `json:"errMsg"`
	Success bool `json:"success"`
}
