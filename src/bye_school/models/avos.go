package models

type Avos struct {
	CommonModel
	UserId       int    `json:"user_id"`
	Installation string `json:"installation"`
}
