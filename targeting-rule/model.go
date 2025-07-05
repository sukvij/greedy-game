package targetingrule

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// Rule defines targeting criteria, stored as part of a JSONB column in TargetingRule
type Rule struct {
	IncludeCountry []string `json:"include_country"`
	ExcludeCountry []string `json:"exclude_country"`
	IncludeOs      []string `json:"include_os"`
	ExcludeOs      []string `json:"exclude_os"`
	IncludeApp     []string `json:"include_app"`
	ExcludeApp     []string `json:"exclude_app"`
}

// Scan implements the sql.Scanner interface to deserialize JSONB into Rule
func (r *Rule) Scan(value interface{}) error {
	if value == nil {
		*r = Rule{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan: unable to scan value into []byte")
	}
	return json.Unmarshal(bytes, r)
}

// Value implements the driver.Valuer interface to serialize Rule to JSONB
func (r Rule) Value() (driver.Value, error) {
	return json.Marshal(r)
}

// TargetingRule defines the table structure with a primary key and a JSONB rules field
type TargetingRule struct {
	CampaignId string    `gorm:"column:cid;primaryKey" json:"cid"`
	Rules      *Rule     `gorm:"type:jsonb;column:rules" json:"rules"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
