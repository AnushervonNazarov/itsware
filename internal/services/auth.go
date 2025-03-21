package services

import (
	"errors"
)

func (s *User) SignIn(email, password string) (string, error) {
	user, err := s.Repo.GetUserByEmailAndPassword(email, password)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Compare provided password with stored hash
	// if !utils.CheckPasswordHash(password, user.Password) {
	// 	return "", errors.New("invalid email or password")
	// }

	role, err := s.Repo.GetRoleByID(int(user.RoleID))
	if err != nil {
		return "", errors.New("failed to retrieve user role")
	}

	accessToken, err := GenerateToken(int(user.ID), user.Email, role, int(user.TenantID))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return accessToken, nil
}
