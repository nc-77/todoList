package request

type Register struct {
	Username string `json:"username" binding:"required,min=2,max=16"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}

type Login struct {
	Register
}
