package repositories

import (
	"fmt"
	"go-gin-mysql-boilerplate/helper"
	"go-gin-mysql-boilerplate/models"
	"go-gin-mysql-boilerplate/responses"
	validations "go-gin-mysql-boilerplate/validations"
	"strconv"

	"gorm.io/gorm"
)

//interface
type BooksRepository interface {
	Save(books_validation *validations.BooksValidation, created_by int64) (interface{}, error)
	FindAll(pagination *helper.Pagination) (interface{}, error)
	FindById(id int64) (interface{}, error)
	Update(id int64, books_validation *validations.BooksValidation, updated_by int64) (interface{}, error)
	Delete(id int64) (interface{}, error)
}

type respositoryBooks struct {
	db *gorm.DB
}

func NewRepositoryBooks(db *gorm.DB) *respositoryBooks {
	return &respositoryBooks{db}
}

func (r *respositoryBooks) FindAll(pagination *helper.Pagination) (interface{}, error) {
	fmt.Println(pagination)
	var responses responses.ResponsesFindAll
	var books []models.Books

	// query data
	result := r.db.Scopes(helper.Paginate(pagination.Limit, pagination.Page)).Order("created_at desc").Find(&books)
	fmt.Println(result.Error)
	if result.Error != nil {
		return nil, result.Error
	}

	// count data
	var books_count []models.Books
	var data_count int64
	r.db.Model(&books_count).Count(&data_count)

	//response
	responses.Result = books
	responses.Limit = pagination.Limit
	responses.Page = pagination.Page
	responses.TotalData = strconv.FormatInt(data_count, 10)
	return responses, nil
}

func (r *respositoryBooks) Save(books_validation *validations.BooksValidation, created_by int64) (interface{}, error) {

	//add data books to database
	var books models.Books
	books.BookNumber = books_validation.BookNumber
	books.BookTitle = books_validation.BookTitle
	books.Author = books_validation.Author
	books.PublicationYear = books_validation.PublicationYear
	books.Publisher = books_validation.Publisher
	books.CreatedBy = created_by

	//save db
	r.db.Create(&books)
	return books, nil
}

func (r *respositoryBooks) FindById(id int64) (interface{}, error) {
	var books models.Books
	if err := r.db.Where("id = ?", id).First(&books).Error; err != nil {
		return nil, err
	} else {
		return books, nil
	}
}

func (r *respositoryBooks) Update(id int64, books_validation *validations.BooksValidation, updated_by int64) (interface{}, error) {

	// check by id
	var books models.Books
	if err := r.db.Where("id = ?", id).First(&books).Error; err != nil {
		return nil, err
	} else {
		//add data books to database
		books.BookNumber = books_validation.BookNumber
		books.BookTitle = books_validation.BookTitle
		books.Author = books_validation.Author
		books.PublicationYear = books_validation.PublicationYear
		books.Publisher = books_validation.Publisher
		books.UpdatedBy = updated_by

		//save db
		r.db.Save(&books)
		return books, nil
	}
}

func (r *respositoryBooks) Delete(id int64) (interface{}, error) {
	var books models.Books
	if err := r.db.Where("id = ?", id).First(&books).Error; err != nil {
		return nil, err
	} else {

		//delete data
		d := r.db.Where("id = ?", id).Delete(&books)
		fmt.Println(d)

		return nil, nil
	}
}
