// Code generated by goctl. DO NOT EDIT.
package types

type TestRequest struct {
	Username string `form:"username"`
}

type TestResponse struct {
	Exist bool `json:"exist"`
}
