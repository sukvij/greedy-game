package delivery

type DeliveryResponse struct {
	CampaignID   string `gorm:"column:cid" json:"cid"`
	CampaignName string `gorm:"column:campaign_name" json:"campaign_name"`
	Image        string `gorm:"column:img" json:"img"`
	CTA          string `gorm:"column:cta" json:"cta"`
}
