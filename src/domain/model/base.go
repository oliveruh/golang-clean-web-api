package model

type Country struct {
	BaseModel
	Name      string `gorm:"size:15;type:string;not null;"`
	Cities    []City
	Companies []Company
}

type City struct {
	BaseModel
	Name      string `gorm:"size:10;type:string;not null;"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}

type Company struct {
	BaseModel
	Name      string `gorm:"size:20;type:string;not null,unique"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}

type Color struct {
	BaseModel
	Name    string `gorm:"size:15;type:string;not null,unique"`
	HexCode string `gorm:"size:7;type:string;not null,unique"`
}

type File struct {
	BaseModel
	Name        string `gorm:"size:100;type:string;not null"`
	Directory   string `gorm:"size:100;type:string;not null"`
	Description string `gorm:"size:500;type:string;not null"`
	MimeType    string `gorm:"size:20;type:string;not null"`
}
