package campaign

type Campaign struct {
	// // gorm.Model
	// ID           uint64
	CampaignID   string `gorm:"column:cid;primaryKey"` // Explicitly map to 'id' column and make it primary key
	CampaignName string `gorm:"column:campaign_name"`
	Image        string `gorm:"column:img"`
	CTA          string `gorm:"column:cta"`
	Status       string `gorm:"type:enum('ACTIVE', 'PASSIVE');not null" json:"role"`
}
