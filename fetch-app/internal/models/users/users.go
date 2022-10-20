package users

type Users struct {
	Phone     string  `json:"phone"`
	Name      string  `json:"name"`
	Role      string  `json:"role"`
	CreatedAt string  `json:"created_at"`
	LoggedIn  string  `json:"logged_in"`
	Exp       float64 `json:"exp"`
}

func (p *Users) TableName() string {
	return "t_users"
}

func (p *Users) KeyRedis() string {
	return "data:users"
}
