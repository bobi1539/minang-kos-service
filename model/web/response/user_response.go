package response

type UserResponse struct {
	Id          int64        `json:"id"`
	Email       string       `json:"email"`
	Name        string       `json:"name"`
	PhoneNumber string       `json:"phoneNumber"`
	Role        RoleResponse `json:"role"`
	BaseDomainResponse
}
