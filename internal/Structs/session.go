package Structs

import "time"

type Session struct {
	UserId       int
	ServiceId    int
	CreateDate   time.Time
	ValidUntil   time.Time
	SessionToken string
	SessionId    int
}
