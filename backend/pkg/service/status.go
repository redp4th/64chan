package service

const (
	Unauthorized = 10000 + iota
	AuthenticationFailed
	TokenExpired
)
