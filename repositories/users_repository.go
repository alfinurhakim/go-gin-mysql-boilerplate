package repositories

import (
	"fmt"
	"go-gin-mysql-boilerplate/helper"
	"go-gin-mysql-boilerplate/models"
	"go-gin-mysql-boilerplate/responses"
	"strconv"

	"gorm.io/gorm"
)

//interface
type UsersRepository interface {
	Save(msg string) (string, error)
	FindAll(*helper.Pagination) (interface{}, error)
	FindById(msg string) (string, error)
	Update(msg string) (string, error)
	Delete(msg string) (string, error)
}

type respositoryUsers struct {
	db *gorm.DB
}

func NewRepositoryUsers(db *gorm.DB) *respositoryUsers {
	return &respositoryUsers{db}
}

func (r *respositoryUsers) FindAll(pagination *helper.Pagination) (interface{}, error) {
	fmt.Println(pagination)
	var responses responses.ResponsesFindAll
	var users []models.Users

	// query data
	result := r.db.Scopes(helper.Paginate(pagination.Limit, pagination.Page)).Order("created_at desc").Find(&users)
	fmt.Println(result.Error)
	if result.Error != nil {
		return nil, result.Error
	}

	// count data
	var users_count []models.Users
	var data_count int64
	r.db.Model(&users_count).Count(&data_count)
	responses.Result = users
	responses.Limit = pagination.Limit
	responses.Page = pagination.Page
	responses.TotalData = strconv.FormatInt(data_count, 10)
	return responses, nil
}

func (r *respositoryUsers) Save(msg string) (string, error) {
	return msg, nil
}

func (r *respositoryUsers) FindById(msg string) (string, error) {
	return msg, nil
}

func (r *respositoryUsers) Update(msg string) (string, error) {
	return msg, nil
}

func (r *respositoryUsers) Delete(msg string) (string, error) {
	return msg, nil
}
