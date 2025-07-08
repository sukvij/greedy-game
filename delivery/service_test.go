package delivery

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestCaseStructure struct {
	Input         Request
	Db            *gorm.DB
	Response      *[]DeliveryResponse
	Err           error
	ResponseEqual bool
	ErrorEqual    bool
}

// go test -coverprofile cover.out

// go tool cover -html cover.out

func TestGetDelivery(t *testing.T) {

	testcases := []TestCaseStructure{}
	sqlDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close() // Ensure the mock SQL DB is closed

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB, // Inject the mock SQL DB connection
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open GORM with mock PostgreSQL DB: %v", err)
	}

	testcases = append(testcases, TestCaseStructure{Input: Request{
		Country:         "India",
		OperatingStstem: "Android",
		AppId:           "com.duolingo.ludokinggame1",
	}, Db: nil, Response: nil, Err: fmt.Errorf("database failed"), ResponseEqual: true, ErrorEqual: true})

	testcases = append(testcases, TestCaseStructure{Input: Request{
		Country:         "India",
		OperatingStstem: "Android",
		// AppId:           "com.duolingo.ludokinggame1",
	}, Db: db, Response: nil, Err: fmt.Errorf("app_id are required"), ResponseEqual: true, ErrorEqual: true})

	testcases = append(testcases, TestCaseStructure{Input: Request{
		// Country:         "India",
		OperatingStstem: "Android",
		AppId:           "com.duolingo.ludokinggame1",
	}, Db: db, Response: nil, Err: fmt.Errorf("country_id are required"), ResponseEqual: true, ErrorEqual: true})

	testcases = append(testcases, TestCaseStructure{Input: Request{
		Country: "India",
		// OperatingStstem: "Android",
		AppId: "com.duolingo.ludokinggame1",
	}, Db: db, Response: nil, Err: fmt.Errorf("os_id are required"), ResponseEqual: true, ErrorEqual: true})

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
	}, Err: errors.New("database failed"), ResponseEqual: false, ErrorEqual: false})

	for _, eachCase := range testcases {
		var serviceMethod DeliveryServiceMethods = NewDeliveryService(eachCase.Db, &eachCase.Input)
		res, err := serviceMethod.GetDelivery(context.Background())
		fmt.Println(res, err)
		// var haha *[]DeliveryResponse = eachCase.Response
		if eachCase.ResponseEqual {
			assert.Equal(t, res, eachCase.Response)
		} else {
			assert.NotEqual(t, res, eachCase.Response)
		}
		if eachCase.ErrorEqual {
			assert.Equal(t, err.Error(), eachCase.Err.Error())
		} else {
			assert.NotEqual(t, err.Error(), eachCase.Err.Error())
		}
	}
}
