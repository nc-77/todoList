package response

type LoginResp struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Token     string `json:"token"`
	ExpiresAt string `json:"expiresAt"`
}
