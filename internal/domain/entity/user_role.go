package entity

type UserRole struct {
	ID       int    `gorm:"primary_key"`
	RoleName string `json:"role_name"`
	Users    []User `gorm:"foreignKey:RoleId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
