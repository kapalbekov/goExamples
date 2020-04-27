package model

import "time"

type NotifyInfo struct {
	To          string     `json:"to" db:"to"`
	Subject     string     `json:"subject" db:"subject"`
	MessageBody string     `json:"msg_body" db:"message_body"`
	SendStatus  string     `db:"send_status"` // success, error
	ErrorText   *string    `json:"error_text" db:"error_text"`
	CreateDt    *time.Time `db:"createdt"`
}
