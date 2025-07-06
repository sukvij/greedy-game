package targetingrule

import "gorm.io/gorm"

type TargetingRuleService struct {
	Db            *gorm.DB
	Targetingrule *TargetingRule
}

func NewTargetingRuleService(db *gorm.DB, targetingRule *TargetingRule) *TargetingRuleService {
	return &TargetingRuleService{Db: db, Targetingrule: targetingRule}
}

func (service *TargetingRuleService) GetAllTargetingRule() (*[]TargetingRule, error) {
	targetingRuleRepo := NewTargetingRuleRepository(service.Db, service.Targetingrule)
	return targetingRuleRepo.GetAllTargetingRule()
}

func (service *TargetingRuleService) CreateTargetingRule() (*TargetingRule, error) {
	targetingRuleRepo := NewTargetingRuleRepository(service.Db, service.Targetingrule)
	return targetingRuleRepo.CreateTargetingRule()
}

func (service *TargetingRuleService) UpdateTargetingRule() (*TargetingRule, error) {
	targetingRuleRepo := NewTargetingRuleRepository(service.Db, service.Targetingrule)
	return targetingRuleRepo.UpdateTargetingRule()
}
