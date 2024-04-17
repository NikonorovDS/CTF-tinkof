package model

type UserDTO struct {
	Username string `json:"username"`
	Balance  uint   `json:"balance"`
	Luck     uint   `json:"luck"`
}

func (dto *UserDTO) ToModel() *User {
	return &User{
		Username: dto.Username,
		Balance:  dto.Balance,
		Luck:     dto.Luck,
	}
}

func (m *User) ToDTO() *UserDTO {
	return &UserDTO{
		Username: m.Username,
		Balance:  m.Balance,
		Luck:     m.Luck,
	}
}
