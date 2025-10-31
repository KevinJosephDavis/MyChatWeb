package repository

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	"github.com/kevinjosephdavis/chatroom/internal/repository/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return fmt.Errorf("创建用户失败,err=%v", result.Error)
	}
	return nil
}

func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("根据用户名获取用户失败,err=%v", result.Error)
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email=?", email).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("根据邮箱获取用户失败,err=%v", result.Error)
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(id int64) (*model.User, error) {
	var user model.User
	result := r.db.Where("id=?", id).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("根据ID获取用户失败,err=%v", result.Error)
	}
	return &user, nil
}

// GetUserByAccountID 根据对外账号ID查找用户
func (r *UserRepository) GetUserByAccountID(accountID int64) (*model.User, error) {
	var user model.User
	result := r.db.Where("account_id = ?", accountID).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在
		}
		return nil, fmt.Errorf("根据账号ID获取用户失败: %v", result.Error)
	}
	return &user, nil
}

// 生成5-10位随机数
func (r *UserRepository) GenerateUniqueAccountID() (int64, error) {

	// 随机选择位数：5, 6, 7, 8, 9, 10
	digits := rand.Intn(6) + 5 // 5-10

	// 计算该位数的最小值和最大值
	min := int64(math.Pow10(digits - 1)) // 10^(n-1)
	max := int64(math.Pow10(digits)) - 1 // 10^n - 1

	accountID := rand.Int63n(max-min+1) + min

	// 检查唯一性
	var count int64
	r.db.Model(&model.User{}).Where("account_id = ?", accountID).Count(&count)
	if count > 0 {
		return r.GenerateUniqueAccountID()
	}

	return accountID, nil
}
