package delivery

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sukvij/greedy-game/database"
	"gorm.io/gorm"
)

type TestCaseStructure struct {
	Input    Request
	Db       *gorm.DB
	Response *[]DeliveryResponse
	Err      error
}

func TestGetDelivery(t *testing.T) {

	testcases := []TestCaseStructure{}
	db, _ := database.Connection()
	testcases = append(testcases, TestCaseStructure{Input: Request{
		Country:         "India",
		OperatingStstem: "Android",
		AppId:           "com.duolingo.ludokinggame1",
	}, Db: nil, Response: nil, Err: fmt.Errorf("database failed")})

	testcases = append(testcases, TestCaseStructure{Input: Request{
		Country:         "India",
		OperatingStstem: "Android",
		// AppId:           "com.duolingo.ludokinggame1",
	}, Db: db, Response: nil, Err: fmt.Errorf("app_id are required")})

	testcases = append(testcases, TestCaseStructure{Input: Request{
		// Country:         "India",
		OperatingStstem: "Android",
		AppId:           "com.duolingo.ludokinggame1",
	}, Db: db, Response: nil, Err: fmt.Errorf("country_id are required")})

	testcases = append(testcases, TestCaseStructure{Input: Request{
		Country: "India",
		// OperatingStstem: "Android",
		AppId: "com.duolingo.ludokinggame1",
	}, Db: db, Response: nil, Err: fmt.Errorf("os_id are required")})

	testcases = append(testcases, TestCaseStructure{Input: Request{
		Country:         "US",
		OperatingStstem: "Android",
		AppId:           "com.duolingo.ludokinggame",
	}, Db: db, Response: &[]DeliveryResponse{
		{
			CampaignID:   "duolingo",
			CampaignName: "Duolingo: Best way to learn",
			Image:        "https://somelink2",
			CTA:          "Install",
		},
	}, Err: nil})

	for _, eachCase := range testcases {
		var serviceMethod DeliveryServiceMethods = NewDeliveryService(eachCase.Db, &eachCase.Input)
		res, err := serviceMethod.GetDelivery()
		fmt.Println(res, err)
		// var haha *[]DeliveryResponse = eachCase.Response
		assert.Equal(t, res, eachCase.Response)
		assert.Equal(t, err, eachCase.Err)
	}
}
