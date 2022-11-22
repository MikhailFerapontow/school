package school

type Guardian struct {
	Id      int    `json:"guadian_id" db:"guadian_id"`
	Name    string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
	Gender  string `json:"gender" db:"gender"`
	Phone   string `json:"phone" db:"phone"`
	Email   string `json:"email" db:"email"`
}
