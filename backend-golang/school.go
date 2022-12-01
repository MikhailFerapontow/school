package school

type Guardian struct {
	Guardian_id int    `json:"guadian_id"`
	Name        string `json:"name" db:"name"`
	Surname     string `json:"surname" db:"surname"`
	Gender      string `json:"gender" db:"gender"`
	Phone       string `json:"phone" db:"phone"`
	Email       string `json:"email" db:"email"`
}

type RegisterAdmin struct {
	Login    string `json:"login"`
	Password string `json:"password" binding:"required"`
}

type RegisterStudent struct {
	Login    string `json:"login"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type RegisterTeacher struct {
	Login    string `json:"login"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"fullname" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type User struct {
	Login    string `json:"login" db:"login"`
	Password string `json:"password" binding:"required"`
	Role_id  int    `json:"role_id" binding:"required"`
	IsActive bool   `json:"isActive" binding:"required"`
}
