package domain

import "time"

type Client struct {
	oid         int64
	Username    string
	Nationality string
	Location    string
	Rating      int
	Age         int
	R15         int
	R30         int
	R45         int
	R100        int
	R150        int
	R200        int
	R250        int
	R300        int
	R350        int
	R400        int
	RON         int
	Telephone   string
	Url         string
	RefreshTime time.Time
	UserId      string
	Region      string
	Gender      string
	MemberSince time.Time
	Height      int
	DSize       int
	HairColor   string
	EyeColor    string
}
