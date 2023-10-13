package services

import (
	"log"
	"synapsis_test/entities"
	"synapsis_test/inputs"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AuthService interface {
	Register(payload *inputs.CreateUserInput) error
}

type AuthServiceImpl struct {
	DB          *gorm.DB
	UtilService UtilService
}

func (service *AuthServiceImpl) Register(payload *inputs.CreateUserInput) error {
	// hash the password first!
	hashed, err := service.UtilService.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	// begin database transactions
	tx := service.DB.Begin()
	if tx.Error != nil {
		log.Println("TX Error:", tx.Error)
		return tx.Error
	}

	user := entities.User{
		Username: payload.Username,
		Password: hashed,
	}

	// start create tx with returning values.
	// return values will automatically bind into user variable.
	if err := tx.Clauses(clause.Returning{}).Create(&user).Error; err != nil {
		// cancel the transaction
		tx.Rollback()
		log.Println("Error Creating User:", err)
		return err
	}

	cart := entities.Cart{
		UserID: user.ID,
	}

	if err := tx.Create(&cart).Error; err != nil {
		// cancel the transaction
		tx.Rollback()
		log.Println("Error Creating Cart:", err)
		return err
	}

	tx.Commit()

	return nil
}
