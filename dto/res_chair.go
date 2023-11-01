package dto

type ResChair struct {
	Id        int64  `json:"id"`
	CodeRef   string `json:"code_ref"`
	IsBook    int8   `json:"is_book"`
	UserBook  string `json:"user_book"`
	UserPhone string `json:"user_phone"`
	Pay       int8   `json:"pay_status"`
}
