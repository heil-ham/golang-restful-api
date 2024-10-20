package web

type WebResponse struct {
	Code int32 `json:"code"`
	Status string `json:"status"`
	Data interface{} `json:"data"`
}