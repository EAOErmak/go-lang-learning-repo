package models

type DictionaryItem struct {
	BaseModel
	Type  DictionaryType `gorm:"type:varchar(50);not null;uniqueIndex:udx_dictionary_type_label" json:"type"`
	Label string         `gorm:"not null;uniqueIndex:udx_dictionary_type_label" json:"label"`
}

func (DictionaryItem) TableName() string {
	return "dictionary_item"
}
