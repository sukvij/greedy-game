package targetingrule

import "gorm.io/gorm"

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
