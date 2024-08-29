package service

import "github.com/muhammadsaefulr/escommerce/internal/domain/entity"

type CategoryService struct {
	CategoryRepo entity.CategoryRepository
}

func NewCategoryService(repo entity.CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepo: repo}
}

func (s *CategoryService) CreateCategory(Category *entity.CategoryProduct) (*entity.CategoryProduct, error) {
	return s.CategoryRepo.CreateCategory(Category)
}

func (s *CategoryService) GetCategoryById(id int) (*entity.CategoryProduct, error) {
	return s.CategoryRepo.GetCategoryById(id)
}

func (s *CategoryService) GetCategoryByName(categoryname string) (*entity.CategoryProduct, error) {
	return s.CategoryRepo.GetCategoryByName(categoryname)
}

func (s *CategoryService) UpdateCategoryData(id int, Category *entity.CategoryProduct) error {
	return s.CategoryRepo.UpdateCategoryData(id, Category)
}

func (s *CategoryService) DeleteCategoryById(id int) error {
	return s.CategoryRepo.DeleteCategoryById(id)
}
