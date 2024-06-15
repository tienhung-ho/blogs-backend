package tokenhelper

type JwtPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}
