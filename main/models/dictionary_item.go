package models

import (
	"errors"
	"strings"
)

type DictionaryItem struct {
	BaseModel
	Type  DictionaryType `gorm:"type:varchar(50);not null;uniqueIndex:udx_dictionary_type_label" json:"type"`
	Label string         `gorm:"not null;uniqueIndex:udx_dictionary_type_label" json:"label"`
}

func (DictionaryItem) TableName() string {
	return "dictionary_item"
}

func NewDictionaryItem(itemType DictionaryType, label string) (*DictionaryItem, error) {
	item := &DictionaryItem{}
	if err := item.UpdateType(itemType); err != nil {
		return nil, err
	}

	if err := item.UpdateLabel(label); err != nil {
		return nil, err
	}

	return item, nil
}

func (d *DictionaryItem) UpdateType(itemType DictionaryType) error {
	if !itemType.IsValid() {
		return errors.New("invalid dictionary type")
	}

	d.Type = itemType
	return nil
}

func (d *DictionaryItem) UpdateLabel(label string) error {
	trimmed := strings.TrimSpace(label)
	if trimmed == "" {
		return errors.New("label is required")
	}

	d.Label = trimmed
	return nil
}
