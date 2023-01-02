package models

import "testing"

func TestUpdateUser(t *testing.T) {
	user := User{
		ID:         0,
		Username:   "jason",
		Password:   "",
		Email:      "",
		Name:       "",
		CoverPic:   "",
		ProfilePic: "123",
		City:       "3245dfg",
		WebSite:    "sfd23",
	}
	UpdateUser(user)
}