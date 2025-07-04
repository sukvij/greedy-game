package campaign

import (
	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db       *gorm.DB
	Campaign *Campaign
}

func NewCampaignRepository(db *gorm.DB, campaign *Campaign) *CampaignRepository {
	return &CampaignRepository{Db: db, Campaign: campaign}
}

func (campaignRepository *CampaignRepository) GetAllCampaign() (*[]Campaign, error) {
	campaigns := []Campaign{}
	err := campaignRepository.Db.Find(&campaigns).Error
	if err != nil {
		return nil, err
	}
	return &campaigns, nil
}

func (campaignRepository *CampaignRepository) CreateCampaign() (*Campaign, error) {
	err := campaignRepository.Db.Create(campaignRepository.Campaign).Error
	if err != nil {
		return nil, err
	}
	// fetch creaetd campaign from database to check
	return campaignRepository.Campaign, nil
}
