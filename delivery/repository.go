package delivery

import (
	"fmt"

	"gorm.io/gorm"
)

type DeliveryRepository struct {
	Db      *gorm.DB
	Request *Request
}

type DeliveryRepositoryMethods interface {
	GetDelivery() (*[]DeliveryResponse, error)
}

func NewDeliveryRepository(db *gorm.DB, request *Request) *DeliveryRepository {
	return &DeliveryRepository{Db: db, Request: request}
}

func (repository *DeliveryRepository) GetDelivery() (*[]DeliveryResponse, error) {

	// dsn := "inner join targeting_rules on campaigns.cid = targeting_rules.cid"

	// ans := repository.Db.Table("campaigns").Joins(dsn)
	// for i := 0; i < len(conditions); i++ {
	// 	ans = ans.Where(conditions[i])
	// }
	// err := ans.Select("*").Scan(&results).Error

	countryJSON := fmt.Sprintf(`["%s"]`, repository.Request.Country)

	osJSON := fmt.Sprintf(`["%s"]`, repository.Request.OperatingStstem)
	appJSON := fmt.Sprintf(`["%s"]`, repository.Request.AppId)
	// err := repository.Db.Table("targeting_rules t").
	// 	Joins("INNER JOIN campaigns c ON t.cid = c.cid").
	// 	Where("(t.rules->'include_country' @> '[]'::jsonb OR t.rules->'include_country' @> ?)", countryJSON).
	// 	Or("(t.rules->'include_country' @> '[]'::jsonb OR NOT (t.rules->'exclude_country' @> ?))", countryJSON).
	// 	Where("(t.rules->'include_os' @> '[]'::jsonb OR t.rules->'include_os' @> ?)", osJSON).
	// 	Or("(t.rules->'include_os' @> '[]'::jsonb OR NOT (t.rules->'exclude_os' @> ?))", osJSON).
	// 	Where("(t.rules->'include_app' @> '[]'::jsonb OR t.rules->'include_app' @> ?)", appJSON).
	// 	Or("(t.rules->'include_app' @> '[]'::jsonb OR NOT (t.rules->'exclude_app' @> ?))", appJSON).
	// 	Scan(&results).Error
	// var results []DeliveryResponse
	// query := fmt.Sprintf(`SELECT
	// 						c.cid,
	// 						c.campaign_name,
	// 						c.img,
	// 						c.cta
	// 					FROM targeting_rules t
	// 					INNER JOIN campaigns c ON t.cid = c.cid
	// 					WHERE
	// 						(
	// 							(t.rules->'include_country' IS NOT NULL AND t.rules->'include_country' @> '%v'::jsonb)
	// 							OR
	// 							(t.rules->'include_country' IS NULL AND NOT (t.rules->'exclude_country' @> '%v'::jsonb))
	// 						)
	// 						AND (
	// 							(t.rules->'include_os' IS NOT NULL AND t.rules->'include_os' @> '%v'::jsonb)
	// 							OR
	// 							(t.rules->'include_os' IS NULL AND NOT (t.rules->'exclude_os' @> '%v'::jsonb))
	// 						)
	// 						AND (
	// 							(t.rules->'include_app' IS NOT NULL AND t.rules->'include_app' @> '%v'::jsonb)
	// 							OR
	// 							(t.rules->'include_app' IS NULL AND NOT (t.rules->'exclude_app' @> '%v'::jsonb))
	// 						);`, countryJSON, countryJSON, osJSON, osJSON, appJSON, appJSON)
	// err := repository.Db.Raw(query, &results).Error

	var results []DeliveryResponse
	response := repository.Db.Table("targeting_rules t").
		Joins("INNER JOIN campaigns c ON t.cid = c.cid and c.status = 'ACTIVE'").
		Where("(t.rules->'include_country' IS NOT NULL AND t.rules->'include_country' @> ?) OR (t.rules->'include_country' IS NULL AND NOT (t.rules->'exclude_country' @> ?))", countryJSON, countryJSON).
		Where("(t.rules->'include_os' IS NOT NULL AND t.rules->'include_os' @> ?) OR (t.rules->'include_os' IS NULL AND NOT (t.rules->'exclude_os' @> ?))", osJSON, osJSON).
		Where("(t.rules->'include_app' IS NOT NULL AND t.rules->'include_app' @> ?) OR (t.rules->'include_app' IS NULL AND NOT (t.rules->'exclude_app' @> ?))", appJSON, appJSON).
		Select("c.cid, c.campaign_name, c.img, c.cta").
		Scan(&results)
	// fmt.Println("bhai ", response.Error, response.RowsAffected, response)
	if response.Error == nil {

		if response.RowsAffected == 0 || response == nil {
			// fmt.Println("bhai ", response.Error, response.RowsAffected, response)
			return nil, gorm.ErrRecordNotFound
		} else {
			return &results, nil
		}
	}
	return nil, response.Error
}
