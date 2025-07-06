package targetingrule

import (
	"fmt"

	"gorm.io/gorm"
)

type TargetingRuleRepository struct {
	Db            *gorm.DB
	TargetingRule *TargetingRule
}

func NewTargetingRuleRepository(db *gorm.DB, targetingRule *TargetingRule) *TargetingRuleRepository {
	return &TargetingRuleRepository{Db: db, TargetingRule: targetingRule}
}

func (repository *TargetingRuleRepository) GetAllTargetingRule() (*[]TargetingRule, error) {
	targetingRule := []TargetingRule{}
	err := repository.Db.Find(&targetingRule).Error
	if err != nil {
		return nil, err
	}
	return &targetingRule, nil
}

func (repository *TargetingRuleRepository) CreateTargetingRule() (*TargetingRule, error) {
	err := repository.Db.Create(repository.TargetingRule).Error
	if err != nil {
		return nil, err
	}
	// fetch creaetd campaign from database to check
	return repository.TargetingRule, nil
}

func (repository *TargetingRuleRepository) UpdateTargetingRule() (*TargetingRule, error) {

	var existing TargetingRule
	result := repository.Db.Where("cid = ?", repository.TargetingRule.CampaignId).First(&existing)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Create if not found record bro
			return repository.CreateTargetingRule()
		}
		return nil, fmt.Errorf("failed to check existing rule: %v", result.Error)
	}

	// Update existing rule
	result = repository.Db.Model(&existing).Updates(repository.TargetingRule)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to update targeting rule: %v", result.Error)
	}
	return nil, nil
}
