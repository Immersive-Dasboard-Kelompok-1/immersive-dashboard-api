package service

import (
	//"be17/main/feature/user"

	"alta/immersive-dashboard-api/features/classes"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type classService struct {
	classData classes.ClassDataInterface
	validate  *validator.Validate
}

// GetAll implements classes.ClassServiceInterface
func (service *classService) GetAll() ([]classes.Core, error) {
	dataClass, err := service.classData.SelectAll()
	return dataClass,err
}

// Delete implements classes.ClassServiceInterface
func (service *classService) Deleted(id int, UserId int) error {
	err := service.classData.Deleted(id, UserId)
	return err
}

// Edit implements classes.ClassServiceInterface
func (service *classService) Edit(id int, UserId int, input classes.Core) error {
	err := service.classData.Update(id, UserId, input)
	if err != nil {
		return fmt.Errorf("failed to update classses with ID %d:%w", id, err)
	}
	return nil
}

// Create implements classes.ClassServiceInterface
func (service *classService) Create(input classes.Core, UserId int) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	errInsert := service.classData.Insert(input, UserId)
	return errInsert
}

func New(repo classes.ClassDataInterface) classes.ClassServiceInterface {
	return &classService{
		classData: repo,
		validate:  validator.New(),
	}
}
