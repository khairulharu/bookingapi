package dto

type ReqChair struct {
	Id        int64  `json:"id"`
	CodeRef   string `json:"code_ref"`
	UserBook  string `json:"user_book"`
	UserPhone string `json:"user_phone"`
}
