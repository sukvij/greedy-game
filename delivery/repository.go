package delivery

import (
	"fmt"

	"gorm.io/gorm"
)

type DeliveryRepository struct {
	Db      *gorm.DB
	Request *Request
}

func NewDeliveryRepository(db *gorm.DB, request *Request) *DeliveryRepository {
	return &DeliveryRepository{Db: db, Request: request}
}

func (repository *DeliveryRepository) GetDelivery() (interface{}, error) {
	results := []DeliveryResponse{}
	dsn := fmt.Sprintf(`inner join targeting_rules on 
				campaigns.cid = targeting_rules.cid 
				where targeting_rules.rules -> 'include_country' @> '["%s"]'::jsonb`, repository.Request.Country[0])
	repository.Db.Table("campaigns").Select("*").Joins(dsn).Scan(&results)
	return &results, nil
}
