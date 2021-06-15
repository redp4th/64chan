package service

var message = map[int]string{
	Unauthorized:         "Need token to authorize request",
	AuthenticationFailed: "Authentication Failed",
	TokenExpired:         "Token expired",
}

func ErrorMessage(code int) string {
	m, ok := message[code]
	if !ok {
		return "unknown code"
	}
	return m
}
