package dto

type Response struct {
	Code    string      `json:"code"`
	Massage string      `json:"massage"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}
