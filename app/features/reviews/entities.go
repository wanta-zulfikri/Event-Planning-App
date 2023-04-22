package reviews

type Core struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	EventID     uint   `gorm:"not null;foreignKey:EventID"`
	ReviewScore int    `gorm:"not null"`
	ReviewText  string `gorm:"not null"`
}
