package model

type AuthUser struct {
	PubModel
	Username string `gorm:"type:varchar(32);unique;NOT NULL" json:"username"`
	Password string `gorm:"type:varchar(32);NOT NULL" json:"password"`
	CityAdcode string `gorm:"type:varchar(8)" json:"city_adcode"`
	Role string `gorm:"type:enum('member','admin', 'superuser'); NOT NULL" json:"role"`
}

func (AuthUser) TableName() string {
	return "auth_user"
}
