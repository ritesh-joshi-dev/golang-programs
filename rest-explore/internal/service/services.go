package service

import (
	"rest-explore/internal/controller"
	"rest-explore/internal/model"
)

func CreateDataRegistration(registrationData model.RegistrationData) (*model.RegistrationData, error) {

	//Additional conditions are applied here

	res, err := controller.CreateDataRegistration(registrationData)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func InsertData(user model.RegistrationData) (*model.RegistrationData, error) {
	return controller.Insert(user)
}

// Get by name
func GetByName(name string) (*model.RegistrationData, error) {
	return controller.GetByName(name)
}

// Get all records
func GetAll() []*model.RegistrationData {
	return controller.GetAll()
}

// Update record
func Update(data model.RegistrationData) (*model.RegistrationData, error) {
	return controller.Update(data)
}

// Delete by ID (string for simplicity)
func Delete(id string) error {
	return controller.Delete(id)
}
