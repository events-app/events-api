package auth

type JwtToken struct {
	Token   string `json:"token"`
	Expires int64  `json:"expiration_date"`
}

type Exception struct {
	Message string `json:"message"`
}
