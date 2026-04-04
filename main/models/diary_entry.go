package models

import (
	"errors"
	"slices"
	"strings"
	"time"
)

type DiaryEntry struct {
	BaseModel
	Metrics     []EntryMetric `gorm:"foreignKey:DiaryEntryID;constraint:OnDelete:CASCADE;" json:"metrics,omitempty"`
	WhenStarted time.Time     `gorm:"column:when_started;not null;index:idx_diary_started" json:"when_started"`
	WhenEnded   time.Time     `gorm:"column:when_ended;not null" json:"when_ended"`
	Duration    int           `gorm:"not null" json:"duration"`
	Mood        *int16        `json:"mood,omitempty"`
	Description string        `gorm:"size:1000;not null" json:"description"`
	Status      EntryStatus   `gorm:"type:varchar(30);not null;index:idx_diary_status" json:"status"`
}

func (DiaryEntry) TableName() string {
	return "diary_entry"
}

func NewDiaryEntry(started, ended time.Time, mood *int16, description string) (*DiaryEntry, error) {
	if started.IsZero() || ended.IsZero() || !ended.After(started) {
		return nil, errors.New("invalid time range")
	}

	entry := &DiaryEntry{
		WhenStarted: started.UTC(),
		WhenEnded:   ended.UTC(),
		Status:      EntryStatusFailed,
	}

	if err := entry.UpdateMood(mood); err != nil {
		return nil, err
	}

	if err := entry.UpdateDescription(description); err != nil {
		return nil, err
	}

	entry.recalculateDuration()
	entry.AutoUpdateStatusByTime(time.Now().UTC())

	return entry, nil
}

func (d *DiaryEntry) UpdateDescription(description string) error {
	trimmed := strings.TrimSpace(description)
	if trimmed == "" {
		return errors.New("description is required")
	}

	d.Description = trimmed
	return nil
}

func (d *DiaryEntry) UpdateMood(mood *int16) error {
	if mood != nil && (*mood < 1 || *mood > 5) {
		return errors.New("mood must be between 1 and 5")
	}

	d.Mood = mood
	return nil
}

func (d *DiaryEntry) AutoUpdateStatusByTime(now time.Time) {
	if d.Status == EntryStatusDeleted {
		return
	}

	if now.Before(d.WhenStarted) {
		d.Status = EntryStatusScheduled
	} else if !now.After(d.WhenEnded) {
		d.Status = EntryStatusActive
	} else {
		d.Status = EntryStatusFinished
	}
}

func (d *DiaryEntry) ForceStatusWin() error {
	if d.Status == EntryStatusDeleted {
		return errors.New("deleted entry cannot change status")
	}

	d.Status = EntryStatusFinished
	return nil
}

func (d *DiaryEntry) UpdateTime(started, ended time.Time) error {
	if started.IsZero() || ended.IsZero() || !ended.After(started) {
		return errors.New("invalid time range")
	}

	d.WhenStarted = started.UTC()
	d.WhenEnded = ended.UTC()
	d.recalculateDuration()
	d.AutoUpdateStatusByTime(time.Now().UTC())

	return nil
}

func (d *DiaryEntry) ChangeStatus(newStatus EntryStatus) error {
	if d.Status == EntryStatusDeleted {
		return errors.New("deleted entry cannot change status")
	}

	if !newStatus.IsValid() {
		return errors.New("invalid entry status")
	}

	d.Status = newStatus
	return nil
}

func (d *DiaryEntry) MarkDeleted() {
	d.Status = EntryStatusDeleted
}

func (d *DiaryEntry) AddMetric(item *EntryMetric) error {
	if item == nil {
		return errors.New("metric cannot be null")
	}

	item.AttachTo(d)
	d.Metrics = append(d.Metrics, *item)
	return nil
}

func (d *DiaryEntry) RemoveMetric(metricID uint) {
	index := slices.IndexFunc(d.Metrics, func(metric EntryMetric) bool {
		return metric.ID == metricID
	})
	if index == -1 {
		return
	}

	d.Metrics[index].Detach()
	d.Metrics = slices.Delete(d.Metrics, index, index+1)
}

func (d *DiaryEntry) IsWin() bool {
	return d.Status == EntryStatusFinished
}

func (d *DiaryEntry) IsLose() bool {
	return d.Status == EntryStatusFailed
}

func (d *DiaryEntry) recalculateDuration() {
	d.Duration = int(d.WhenEnded.Sub(d.WhenStarted).Minutes())
}
