package campaign

import (
	"gorm.io/gorm"
)

type CampaignService struct {
	Db       *gorm.DB
	Campaign *Campaign
}

func NewCampaignService(db *gorm.DB, campaign *Campaign) *CampaignService {
	return &CampaignService{Db: db, Campaign: campaign}
}

func (campaignService *CampaignService) GetAllCampaign() (*[]Campaign, error) {
	campaignRepository := NewCampaignRepository(campaignService.Db, campaignService.Campaign)
	return campaignRepository.GetAllCampaign()
}

func (campaignService *CampaignService) CreateCampaign() (*Campaign, error) {
	campaignRepository := NewCampaignRepository(campaignService.Db, campaignService.Campaign)
	return campaignRepository.CreateCampaign()
}
