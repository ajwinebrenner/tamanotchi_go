package models

type House struct {
	Name        string `gorm:"PrimaryKey"`
	Price       int
	Size        int
	UpgradeName string
	Upgrade     *House `gorm:"foreignKey:UpgradeName"`
	//TODO: some other bonus
}
