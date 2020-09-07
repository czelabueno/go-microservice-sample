package gateways

import (
	"github.com/czelabueno/go-microservice-sample/gadgets/smartphones/models"
	"github.com/czelabueno/go-microservice-sample/internal/database"
	"github.com/czelabueno/go-microservice-sample/internal/logs"
)

type SmartphoneStorageGateway interface {
	Add(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error)
}

type SmartphoneStorage struct {
	*database.MySQLClient
}

func (s *SmartphoneStorage) Add(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error) {
	tx, err := s.MySQLClient.Begin()

	if err != nil {
		logs.Log().Error("Cant create transaction")
		return nil, err
	}
	res, err := tx.Exec(`insert into smartphone (name, price, country_origin, os)
	values (?, ?, ?, ?)`, cmd.Name, cmd.Price, cmd.CountryOrigin, cmd.Os)

	if err != nil {
		logs.Log().Error("Cant execute statement")
		_ = tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		logs.Log().Error("Cant fetch last id")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &models.Smartphone{
		Id:            id,
		Name:          cmd.Name,
		Price:         cmd.Price,
		CountryOrigin: cmd.CountryOrigin,
		OS:            cmd.OS,
	}, nil

}
