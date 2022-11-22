package school

type Guardian struct {
	Id      int    `json:"guadian_id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}
