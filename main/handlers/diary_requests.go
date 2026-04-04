package handlers

import "time"

type diaryEntryRequest struct {
	WhenStarted time.Time            `json:"when_started" binding:"required"`
	WhenEnded   time.Time            `json:"when_ended" binding:"required"`
	Mood        *int16               `json:"mood"`
	Description string               `json:"description" binding:"required"`
	Metrics     []entryMetricRequest `json:"metrics"`
}

type entryMetricRequest struct {
	MetricTypeID uint                      `json:"metric_type_id" binding:"required"`
	Values       []entryMetricValueRequest `json:"values"`
}

type entryMetricValueRequest struct {
	UnitID uint `json:"unit_id" binding:"required"`
	Value  int  `json:"value" binding:"required"`
}
