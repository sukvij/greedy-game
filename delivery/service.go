package delivery

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

type DeliveryService struct {
	Db      *gorm.DB
	Request *Request
}

func NewDeliveryService(db *gorm.DB, request *Request) *DeliveryService {
	return &DeliveryService{Db: db, Request: request}
}

type DeliveryServiceMethods interface {
	GetDelivery(ctx context.Context) (*[]DeliveryResponse, error)
}

func (service *DeliveryService) GetDelivery(ctx context.Context) (*[]DeliveryResponse, error) {
	if service.Request.AppId == "" {
		return nil, errors.New("app_id are required")
	}
	if service.Request.Country == "" {
		return nil, errors.New("country_id are required")
	}
	if service.Request.OperatingStstem == "" {
		return nil, errors.New("os_id are required")
	}

	context, span := otel.Tracer("service").Start(ctx, "getDelivery Service")
	defer span.End()
	var deliveryRepository DeliveryRepositoryMethods = NewDeliveryRepository(service.Db, service.Request)
	return deliveryRepository.GetDelivery(context)
}

// func CreateDifferentCondition(request *Request) []string {
// 	condition := []string{}

// 	countryJSON := fmt.Sprintf(`["%s"]`, request.Country)

// 	osJSON := fmt.Sprintf(`["%s"]`, request.OperatingStstem)
// 	appJSON := fmt.Sprintf(`["%s"]`, request.AppId)
// 	condition = append(condition, fmt.Sprintf("(t.rules->'include_country' @> '[]'::jsonb OR t.rules->'include_country' @> %v)", countryJSON))
// 	condition = append(condition, fmt.Sprintf("(t.rules->'exclude_country' @> '[]'::jsonb OR NOT (t.rules->'exclude_country' @> %v))", countryJSON))
// 	condition = append(condition, fmt.Sprintf("(t.rules->'include_os' @> '[]'::jsonb OR t.rules->'include_os' @> %v)", osJSON))
// 	condition = append(condition, fmt.Sprintf("(t.rules->'exclude_os' @> '[]'::jsonb OR NOT t.rules->'exclude_os' @> %v)", osJSON))
// 	condition = append(condition, fmt.Sprintf(`(t.rules->'include_app' @> '[]'::jsonb OR t.rules->'include_app' @> %v)"`, appJSON))
// 	condition = append(condition, fmt.Sprintf(`(t.rules->'exclude_app' @> '[]'::jsonb OR NOT t.rules->'exclude_app' @> %v)"`, appJSON))
// 	// var temp string = ""

// 	// temp = ""
// 	// for i := 0; i < len(request.Country); i++ {
// 	// 	temp = temp + fmt.Sprintf(`targeting_rules.rules -> 'include_country' @> '["%s"]'::jsonb`, request.Country[i])
// 	// 	if i == len(request.Country)-1 {
// 	// 		condition = append(condition, temp)
// 	// 		break
// 	// 	}
// 	// 	temp = temp + " or "
// 	// }

// 	// temp = ""
// 	// for i := 0; i < len(request.Country); i++ {
// 	// 	temp = temp + fmt.Sprintf(`not (targeting_rules.rules -> 'exclude_country' @> '["%s"]'::jsonb)`, request.Country[i])
// 	// 	if i == len(request.Country)-1 {
// 	// 		condition = append(condition, temp)
// 	// 		break
// 	// 	}
// 	// 	temp = temp + " or "
// 	// }

// 	// temp = ""
// 	// for i := 0; i < len(request.AppId); i++ {
// 	// 	condition = append(condition, fmt.Sprintf(` targeting_rules.rules -> 'include_app' @> '["%s"]'::jsonb`, request.AppId[i]))
// 	// 	if i == len(request.AppId)-1 {
// 	// 		condition = append(condition, temp)
// 	// 		break
// 	// 	}
// 	// 	temp += " or "
// 	// }

// 	// temp = ""
// 	// for i := 0; i < len(request.AppId); i++ {
// 	// 	condition = append(condition, fmt.Sprintf(`not (targeting_rules.rules -> 'exclude_app' @> '["%s"]'::jsonb)`, request.AppId[i]))
// 	// 	if i == len(request.AppId)-1 {
// 	// 		condition = append(condition, temp)
// 	// 		break
// 	// 	}
// 	// 	temp += " or "
// 	// }

// 	// temp = ""
// 	// for i := 0; i < len(request.OperatingStstem); i++ {
// 	// 	temp = temp + fmt.Sprintf(` targeting_rules.rules -> 'include_os' @> '["%s"]'::jsonb`, request.OperatingStstem[i])
// 	// 	if i == len(request.OperatingStstem)-1 {
// 	// 		condition = append(condition, temp)
// 	// 		break
// 	// 	}
// 	// 	temp += " or "
// 	// }

// 	// temp = ""
// 	// for i := 0; i < len(request.OperatingStstem); i++ {
// 	// 	temp = temp + fmt.Sprintf(`not (targeting_rules.rules -> 'exclude_os' @> '["%s"]'::jsonb)`, request.OperatingStstem[i])
// 	// 	if i == len(request.OperatingStstem)-1 {
// 	// 		condition = append(condition, temp)
// 	// 		break
// 	// 	}
// 	// 	temp += " or "
// 	// }
// 	fmt.Println(condition)
// 	return condition
// }
