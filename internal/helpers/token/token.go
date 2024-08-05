package tokenhelper

type JwtPayload struct {
	Id   int    `json:"user_id"`
	Role string `json:"role"`
}
