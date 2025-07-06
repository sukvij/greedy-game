package campaign

import (
	"fmt"

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

func (campaignRepository *CampaignRepository) UpdateCampaign() (*Campaign, error) {
	var existing Campaign
	err := campaignRepository.Db.Where("cid=?", campaignRepository.Campaign.CampaignID).First(&existing).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return campaignRepository.CreateCampaign()
		} else {
			return nil, fmt.Errorf("cannot update error some problem - err is %v", err)
		}
	}
	err = (campaignRepository.Db).Model(&existing).Updates(campaignRepository.Campaign).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update targeting rule: %v", err.Error)
	}
	return campaignRepository.Campaign, nil
}
