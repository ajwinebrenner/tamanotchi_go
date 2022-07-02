package models

type House struct {
	Name        string `gorm:"PrimaryKey"`
	Price       int32
	Size        int32
	UpgradeName string
	Upgrade     *House `gorm:"foreignKey:UpgradeName"`
	//TODO: some other bonus
}
