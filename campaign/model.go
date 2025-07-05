package campaign

type Campaign struct {
	CampaignID   string `gorm:"column:cid;primarykey" json:"cid"`
	CampaignName string `gorm:"column:campaign_name" json:"campaign_name"`
	Image        string `gorm:"column:img" json:"img"`
	CTA          string `gorm:"cta"`
	Status       string `gorm:"status"`
}
