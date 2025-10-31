package service

import (
	"fmt"

	"github.com/kevinjosephdavis/chatroom/internal/repository"
	"github.com/kevinjosephdavis/chatroom/internal/repository/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// Register 注册新用户
func (s *UserService) Register(username, password, email, nickname string) (*model.User, error) {
	// 1.检查邮箱是否已经存在
	existingUser, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("检查邮箱失败,err=%v", err)
	}
	if existingUser != nil {
		return nil, fmt.Errorf("邮箱已存在")
	}

	//2.分配ID
	accountID, err := s.userRepo.GenerateUniqueAccountID()
	if err != nil {
		return nil, fmt.Errorf("分配用户ID失败,err=%v", err)
	}

	// 3. 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("密码加密失败,err=%v", err)
	}
	//是否需要将处理后的哈希密码存到数据库中？

	// 4. 创建用户
	user := &model.User{
		Username:  username,
		AccountID: accountID,
		Password:  string(hashedPassword),
		Email:     email,
		Nickname:  nickname,
		Status:    1,
	}
	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, fmt.Errorf("创建用户失败,err=%v", err)
	}
	return user, nil
}

// Login 用户登录
func (s *UserService) Login(accountID int64, password string) (*model.User, error) {
	user, err := s.userRepo.GetUserByAccountID(accountID)
	if err != nil {
		return nil, fmt.Errorf("获取用户失败,err=%v", err)
	}
	if user == nil {
		return nil, fmt.Errorf("用户不存在")
		// 回到登录界面
	}
	//验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("密码错误,err=%v", err)
	}
	return user, nil
}
