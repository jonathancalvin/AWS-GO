package database

type Users struct {
	UserID   uint64 `gorm:"type:int;primaryKey;autoIncrement:true"`
	UserName string `gorm:"type:varchar(50);not null"`
	Email    string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"type:varchar(100);not null"`
}


