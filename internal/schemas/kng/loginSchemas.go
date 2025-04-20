package planet

// LoginForm представляет данные для авторизации
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
