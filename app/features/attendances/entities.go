package attendances

type Core struct {
	ID      uint
	UserID  uint
	EventID uint
}

type Attendances struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint `gorm:"not null;foreignKey:User"`
	EventID uint `gorm:"not null;foreignKey:Event"`
}
