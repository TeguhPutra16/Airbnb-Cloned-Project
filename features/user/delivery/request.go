package delivery

import (
	"be13/project/features/user"
)

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

func UserRequestToUserCore(data UserRequest) user.CoreUser {
	return user.CoreUser{
		Name:     data.Name,
		Password: data.Password,
		Email:    data.Email,
		Address:  data.Address,
		Role:     data.Role,
	}
}
