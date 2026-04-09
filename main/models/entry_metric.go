package models

import (
	"errors"
	"slices"
)

type EntryMetric struct {
	BaseModel
	MetricTypeID uint               `gorm:"column:metric_type_id;not null;index:idx_metric_type" json:"metric_type_id"`
	MetricType   DictionaryItem     `gorm:"foreignKey:MetricTypeID" json:"metric_type,omitempty"`
	DiaryEntryID uint               `gorm:"column:diary_entry_id;not null;index:idx_metric_entry" json:"diary_entry_id"`
	DiaryEntry   *DiaryEntry        `gorm:"foreignKey:DiaryEntryID" json:"-"`
	Values       []EntryMetricValue `gorm:"foreignKey:EntryMetricID;constraint:OnDelete:CASCADE;" json:"values,omitempty"`
}

func (EntryMetric) TableName() string {
	return "entry_metric"
}

func NewEntryMetric(entryID, metricTypeID uint) (*EntryMetric, error) {
	if entryID == 0 {
		return nil, errors.New("diary entry is required")
	}

	if metricTypeID == 0 {
		return nil, errors.New("metric type is required")
	}

	return &EntryMetric{
		DiaryEntryID: entryID,
		MetricTypeID: metricTypeID,
	}, nil
}

func (m *EntryMetric) AddValue(unitID uint, value int) error {
	if unitID == 0 {
		return errors.New("unit is required")
	}

	if value <= 0 {
		return errors.New("value must be positive")
	}

	exists := slices.ContainsFunc(m.Values, func(item EntryMetricValue) bool {
		return item.UnitID == unitID
	})
	if exists {
		return errors.New("unit already exists for this metric")
	}

	metricValue, err := NewEntryMetricValue(unitID, value)
	if err != nil {
		return err
	}

	m.Values = append(m.Values, *metricValue)
	return nil
}

func (m *EntryMetric) ChangeMetricType(newTypeID uint) error {
	if newTypeID == 0 {
		return errors.New("metric type is required")
	}

	m.MetricTypeID = newTypeID
	return nil
}

func (m *EntryMetric) AttachTo(entry *DiaryEntry) {
	if entry == nil {
		return
	}

	m.DiaryEntry = entry
	m.DiaryEntryID = entry.ID
}

func (m *EntryMetric) Detach() {
	m.DiaryEntry = nil
	m.DiaryEntryID = 0
}
