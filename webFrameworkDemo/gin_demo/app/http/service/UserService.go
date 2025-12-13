package service

import "github.com/syyongx/php2go"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// Service层实现
func (s *UserService) Authenticate(username, password string) (string, error) {
	// 1. 获取用户信息
	//user, err := s.userRepo.FindByUsername(username)
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		return "", ErrUserNotFound
	//	}
	//	return "", err
	//}
	//
	//// 2. 验证密码
	//if !checkPassword(password, user.PasswordHash) {
	//	return "", ErrWrongPassword
	//}
	//
	//// 3. 生成JWT
	//token, err := jwt.GenerateToken(user.ID, user.Username)
	//if err != nil {
	//	return "", fmt.Errorf("token generation failed: %w", err)
	//}
	token := php2go.Md5(username + password)

	return token, nil
}
