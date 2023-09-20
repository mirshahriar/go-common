// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package api

import (
	"encoding/json"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

const (
	SessionTokenScopes = "sessionToken.Scopes"
)

// Defines values for AverageGlucoseUnits.
const (
	MmolL AverageGlucoseUnits = "mmol/L"
	Mmoll AverageGlucoseUnits = "mmol/l"
)

// AverageGlucose Blood glucose value, in `mmol/L`
type AverageGlucose struct {
	Units AverageGlucoseUnits `json:"units"`

	// Value A floating point value representing a `mmol/L` value.
	Value float32 `json:"value"`
}

// AverageGlucoseUnits defines model for AverageGlucose.Units.
type AverageGlucoseUnits string

// BGMBucketData Series of counters which represent one hour of a users data
type BGMBucketData struct {
	Date           *time.Time `json:"date,omitempty"`
	LastRecordTime *time.Time `json:"lastRecordTime,omitempty"`

	// TimeInHighRecords Counter of records in high glucose range
	TimeInHighRecords *int `json:"timeInHighRecords,omitempty"`

	// TimeInLowRecords Counter of records in low glucose range
	TimeInLowRecords *int `json:"timeInLowRecords,omitempty"`

	// TimeInTargetRecords Counter of records in target glucose range
	TimeInTargetRecords *int `json:"timeInTargetRecords,omitempty"`

	// TimeInVeryHighRecords Counter of records in very high glucose range
	TimeInVeryHighRecords *int `json:"timeInVeryHighRecords,omitempty"`

	// TimeInVeryLowRecords Counter of records in very low glucose range
	TimeInVeryLowRecords *int `json:"timeInVeryLowRecords,omitempty"`

	// TotalGlucose Total value of all glucose records
	TotalGlucose *float64 `json:"totalGlucose,omitempty"`
}

// BGMPeriod Summary of a specific BGM time period (currently: 1d, 7d, 14d, 30d)
type BGMPeriod struct {
	// AverageDailyRecords Average daily readings
	AverageDailyRecords *float64 `json:"averageDailyRecords,omitempty"`

	// AverageGlucose Blood glucose value, in `mmol/L`
	AverageGlucose           *AverageGlucose `json:"averageGlucose,omitempty"`
	HasAverageDailyRecords   *bool           `json:"hasAverageDailyRecords,omitempty"`
	HasAverageGlucose        *bool           `json:"hasAverageGlucose,omitempty"`
	HasTimeInHighPercent     *bool           `json:"hasTimeInHighPercent,omitempty"`
	HasTimeInHighRecords     *bool           `json:"hasTimeInHighRecords,omitempty"`
	HasTimeInLowPercent      *bool           `json:"hasTimeInLowPercent,omitempty"`
	HasTimeInLowRecords      *bool           `json:"hasTimeInLowRecords,omitempty"`
	HasTimeInTargetPercent   *bool           `json:"hasTimeInTargetPercent,omitempty"`
	HasTimeInTargetRecords   *bool           `json:"hasTimeInTargetRecords,omitempty"`
	HasTimeInVeryHighPercent *bool           `json:"hasTimeInVeryHighPercent,omitempty"`
	HasTimeInVeryHighRecords *bool           `json:"hasTimeInVeryHighRecords,omitempty"`
	HasTimeInVeryLowPercent  *bool           `json:"hasTimeInVeryLowPercent,omitempty"`
	HasTimeInVeryLowRecords  *bool           `json:"hasTimeInVeryLowRecords,omitempty"`
	HasTotalRecords          *bool           `json:"hasTotalRecords,omitempty"`

	// TimeInHighPercent Percentage of time spent in high glucose range
	TimeInHighPercent *float64 `json:"timeInHighPercent,omitempty"`

	// TimeInHighRecords Counter of records in high glucose range
	TimeInHighRecords *int `json:"timeInHighRecords,omitempty"`

	// TimeInLowPercent Percentage of time spent in low glucose range
	TimeInLowPercent *float64 `json:"timeInLowPercent,omitempty"`

	// TimeInLowRecords Counter of records in low glucose range
	TimeInLowRecords *int `json:"timeInLowRecords,omitempty"`

	// TimeInTargetPercent Percentage of time spent in target glucose range
	TimeInTargetPercent *float64 `json:"timeInTargetPercent,omitempty"`

	// TimeInTargetRecords Counter of records in target glucose range
	TimeInTargetRecords *int `json:"timeInTargetRecords,omitempty"`

	// TimeInVeryHighPercent Percentage of time spent in very high glucose range
	TimeInVeryHighPercent *float64 `json:"timeInVeryHighPercent,omitempty"`

	// TimeInVeryHighRecords Counter of records in very high glucose range
	TimeInVeryHighRecords *int `json:"timeInVeryHighRecords,omitempty"`

	// TimeInVeryLowPercent Percentage of time spent in very low glucose range
	TimeInVeryLowPercent *float64 `json:"timeInVeryLowPercent,omitempty"`

	// TimeInVeryLowRecords Counter of records in very low glucose range
	TimeInVeryLowRecords *int `json:"timeInVeryLowRecords,omitempty"`

	// TotalRecords Counter of records
	TotalRecords *int `json:"totalRecords,omitempty"`
}

// BGMPeriods A map to each supported BGM summary period
type BGMPeriods struct {
	// N14d Summary of a specific BGM time period (currently: 1d, 7d, 14d, 30d)
	N14d *BGMPeriod `json:"14d,omitempty"`

	// N1d Summary of a specific BGM time period (currently: 1d, 7d, 14d, 30d)
	N1d *BGMPeriod `json:"1d,omitempty"`

	// N30d Summary of a specific BGM time period (currently: 1d, 7d, 14d, 30d)
	N30d *BGMPeriod `json:"30d,omitempty"`

	// N7d Summary of a specific BGM time period (currently: 1d, 7d, 14d, 30d)
	N7d *BGMPeriod `json:"7d,omitempty"`
}

// BGMStats A summary of a users recent BGM glucose values
type BGMStats struct {
	// Buckets Rotating list containing the stats for each currently tracked hour in order
	Buckets *[]Bucket `json:"buckets,omitempty"`

	// Periods A map to each supported BGM summary period
	Periods *BGMPeriods `json:"periods,omitempty"`

	// TotalHours Total hours represented in the hourly stats
	TotalHours *int `json:"totalHours,omitempty"`
}

// Bucket bucket containing an hour of bgm or cgm aggregations
type Bucket struct {
	Data           *Bucket_Data `json:"data,omitempty"`
	Date           *time.Time   `json:"date,omitempty"`
	LastRecordTime *time.Time   `json:"lastRecordTime,omitempty"`
}

// Bucket_Data defines model for Bucket.Data.
type Bucket_Data struct {
	union json.RawMessage
}

// CGMBucketData Series of counters which represent one hour of a users data
type CGMBucketData struct {
	// TimeCGMUseMinutes Counter of minutes using a cgm
	TimeCGMUseMinutes *int `json:"timeCGMUseMinutes,omitempty"`

	// TimeCGMUseRecords Counter of records wearing a cgm
	TimeCGMUseRecords *int `json:"timeCGMUseRecords,omitempty"`

	// TimeInHighMinutes Counter of minutes spent in high glucose range
	TimeInHighMinutes *int `json:"timeInHighMinutes,omitempty"`

	// TimeInHighRecords Counter of records in high glucose range
	TimeInHighRecords *int `json:"timeInHighRecords,omitempty"`

	// TimeInLowMinutes Counter of minutes spent in low glucose range
	TimeInLowMinutes *int `json:"timeInLowMinutes,omitempty"`

	// TimeInLowRecords Counter of records in low glucose range
	TimeInLowRecords *int `json:"timeInLowRecords,omitempty"`

	// TimeInTargetMinutes Counter of minutes spent in target glucose range
	TimeInTargetMinutes *int `json:"timeInTargetMinutes,omitempty"`

	// TimeInTargetRecords Counter of records in target glucose range
	TimeInTargetRecords *int `json:"timeInTargetRecords,omitempty"`

	// TimeInVeryHighMinutes Counter of minutes spent in very high glucose range
	TimeInVeryHighMinutes *int `json:"timeInVeryHighMinutes,omitempty"`

	// TimeInVeryHighRecords Counter of records in very high glucose range
	TimeInVeryHighRecords *int `json:"timeInVeryHighRecords,omitempty"`

	// TimeInVeryLowMinutes Counter of minutes spent in very low glucose range
	TimeInVeryLowMinutes *int `json:"timeInVeryLowMinutes,omitempty"`

	// TimeInVeryLowRecords Counter of records in very low glucose range
	TimeInVeryLowRecords *int `json:"timeInVeryLowRecords,omitempty"`

	// TotalGlucose Total value of all glucose records
	TotalGlucose *float64 `json:"totalGlucose,omitempty"`
}

// CGMPeriod Summary of a specific CGM time period (currently: 1d, 7d, 14d, 30d)
type CGMPeriod struct {
	// AverageDailyRecords Average daily readings
	AverageDailyRecords *float64 `json:"averageDailyRecords,omitempty"`

	// AverageGlucose Blood glucose value, in `mmol/L`
	AverageGlucose *AverageGlucose `json:"averageGlucose,omitempty"`

	// GlucoseManagementIndicator A derived value which emulates A1C
	GlucoseManagementIndicator    *float64 `json:"glucoseManagementIndicator,omitempty"`
	HasAverageDailyRecords        *bool    `json:"hasAverageDailyRecords,omitempty"`
	HasAverageGlucose             *bool    `json:"hasAverageGlucose,omitempty"`
	HasGlucoseManagementIndicator *bool    `json:"hasGlucoseManagementIndicator,omitempty"`
	HasTimeCGMUseMinutes          *bool    `json:"hasTimeCGMUseMinutes,omitempty"`
	HasTimeCGMUsePercent          *bool    `json:"hasTimeCGMUsePercent,omitempty"`
	HasTimeCGMUseRecords          *bool    `json:"hasTimeCGMUseRecords,omitempty"`
	HasTimeInHighMinutes          *bool    `json:"hasTimeInHighMinutes,omitempty"`
	HasTimeInHighPercent          *bool    `json:"hasTimeInHighPercent,omitempty"`
	HasTimeInHighRecords          *bool    `json:"hasTimeInHighRecords,omitempty"`
	HasTimeInLowMinutes           *bool    `json:"hasTimeInLowMinutes,omitempty"`
	HasTimeInLowPercent           *bool    `json:"hasTimeInLowPercent,omitempty"`
	HasTimeInLowRecords           *bool    `json:"hasTimeInLowRecords,omitempty"`
	HasTimeInTargetMinutes        *bool    `json:"hasTimeInTargetMinutes,omitempty"`
	HasTimeInTargetPercent        *bool    `json:"hasTimeInTargetPercent,omitempty"`
	HasTimeInTargetRecords        *bool    `json:"hasTimeInTargetRecords,omitempty"`
	HasTimeInVeryHighMinutes      *bool    `json:"hasTimeInVeryHighMinutes,omitempty"`
	HasTimeInVeryHighPercent      *bool    `json:"hasTimeInVeryHighPercent,omitempty"`
	HasTimeInVeryHighRecords      *bool    `json:"hasTimeInVeryHighRecords,omitempty"`
	HasTimeInVeryLowMinutes       *bool    `json:"hasTimeInVeryLowMinutes,omitempty"`
	HasTimeInVeryLowPercent       *bool    `json:"hasTimeInVeryLowPercent,omitempty"`
	HasTimeInVeryLowRecords       *bool    `json:"hasTimeInVeryLowRecords,omitempty"`
	HasTotalRecords               *bool    `json:"hasTotalRecords,omitempty"`

	// TimeCGMUseMinutes Counter of minutes spent wearing a cgm
	TimeCGMUseMinutes *int `json:"timeCGMUseMinutes,omitempty"`

	// TimeCGMUsePercent Percentage of time spent wearing a cgm
	TimeCGMUsePercent *float64 `json:"timeCGMUsePercent,omitempty"`

	// TimeCGMUseRecords Counter of minutes spent wearing a cgm
	TimeCGMUseRecords *int `json:"timeCGMUseRecords,omitempty"`

	// TimeInHighMinutes Counter of minutes spent in high glucose range
	TimeInHighMinutes *int `json:"timeInHighMinutes,omitempty"`

	// TimeInHighPercent Percentage of time spent in high glucose range
	TimeInHighPercent *float64 `json:"timeInHighPercent,omitempty"`

	// TimeInHighRecords Counter of records in high glucose range
	TimeInHighRecords *int `json:"timeInHighRecords,omitempty"`

	// TimeInLowMinutes Counter of minutes spent in low glucose range
	TimeInLowMinutes *int `json:"timeInLowMinutes,omitempty"`

	// TimeInLowPercent Percentage of time spent in low glucose range
	TimeInLowPercent *float64 `json:"timeInLowPercent,omitempty"`

	// TimeInLowRecords Counter of records in low glucose range
	TimeInLowRecords *int `json:"timeInLowRecords,omitempty"`

	// TimeInTargetMinutes Counter of minutes spent in target glucose range
	TimeInTargetMinutes *int `json:"timeInTargetMinutes,omitempty"`

	// TimeInTargetPercent Percentage of time spent in target glucose range
	TimeInTargetPercent *float64 `json:"timeInTargetPercent,omitempty"`

	// TimeInTargetRecords Counter of records in target glucose range
	TimeInTargetRecords *int `json:"timeInTargetRecords,omitempty"`

	// TimeInVeryHighMinutes Counter of minutes spent in very high glucose range
	TimeInVeryHighMinutes *int `json:"timeInVeryHighMinutes,omitempty"`

	// TimeInVeryHighPercent Percentage of time spent in very high glucose range
	TimeInVeryHighPercent *float64 `json:"timeInVeryHighPercent,omitempty"`

	// TimeInVeryHighRecords Counter of records in very high glucose range
	TimeInVeryHighRecords *int `json:"timeInVeryHighRecords,omitempty"`

	// TimeInVeryLowMinutes Counter of minutes spent in very low glucose range
	TimeInVeryLowMinutes *int `json:"timeInVeryLowMinutes,omitempty"`

	// TimeInVeryLowPercent Percentage of time spent in very low glucose range
	TimeInVeryLowPercent *float64 `json:"timeInVeryLowPercent,omitempty"`

	// TimeInVeryLowRecords Counter of records in very low glucose range
	TimeInVeryLowRecords *int `json:"timeInVeryLowRecords,omitempty"`

	// TotalRecords Counter of records
	TotalRecords *int `json:"totalRecords,omitempty"`
}

// CGMPeriods A map to each supported CGM summary period
type CGMPeriods struct {
	// N14d Summary of a specific CGM time period (currently: 1d, 7d, 14d, 30d)
	N14d *CGMPeriod `json:"14d,omitempty"`

	// N1d Summary of a specific CGM time period (currently: 1d, 7d, 14d, 30d)
	N1d *CGMPeriod `json:"1d,omitempty"`

	// N30d Summary of a specific CGM time period (currently: 1d, 7d, 14d, 30d)
	N30d *CGMPeriod `json:"30d,omitempty"`

	// N7d Summary of a specific CGM time period (currently: 1d, 7d, 14d, 30d)
	N7d *CGMPeriod `json:"7d,omitempty"`
}

// CGMStats A summary of a users recent CGM glucose values
type CGMStats struct {
	// Buckets Rotating list containing the stats for each currently tracked hour in order
	Buckets *[]Bucket `json:"buckets,omitempty"`

	// Periods A map to each supported CGM summary period
	Periods *CGMPeriods `json:"periods,omitempty"`

	// TotalHours Total hours represented in the hourly stats
	TotalHours *int `json:"totalHours,omitempty"`
}

// Config Summary schema version and calculation configuration
type Config struct {
	// HighGlucoseThreshold Threshold used for determining if a value is high
	HighGlucoseThreshold *float64 `json:"highGlucoseThreshold,omitempty"`

	// LowGlucoseThreshold Threshold used for determining if a value is low
	LowGlucoseThreshold *float64 `json:"lowGlucoseThreshold,omitempty"`

	// SchemaVersion Summary schema version
	SchemaVersion *int `json:"schemaVersion,omitempty"`

	// VeryHighGlucoseThreshold Threshold used for determining if a value is very high
	VeryHighGlucoseThreshold *float64 `json:"veryHighGlucoseThreshold,omitempty"`

	// VeryLowGlucoseThreshold Threshold used for determining if a value is very low
	VeryLowGlucoseThreshold *float64 `json:"veryLowGlucoseThreshold,omitempty"`
}

// Dates dates tracked for summary calculation
type Dates struct {
	// FirstData Date of the first included value
	FirstData         *time.Time `json:"firstData,omitempty"`
	HasFirstData      *bool      `json:"hasFirstData,omitempty"`
	HasLastData       *bool      `json:"hasLastData,omitempty"`
	HasLastUploadDate *bool      `json:"hasLastUploadDate,omitempty"`
	HasOutdatedSince  *bool      `json:"hasOutdatedSince,omitempty"`

	// LastData Date of the last calculated value
	LastData *time.Time `json:"lastData,omitempty"`

	// LastUpdatedDate Date of the last calculation
	LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`

	// LastUploadDate Created date of the last calculated value
	LastUploadDate *time.Time `json:"lastUploadDate,omitempty"`

	// OutdatedSince Date of the first user upload after lastData, removed when calculated
	OutdatedSince *time.Time `json:"outdatedSince,omitempty"`
}

// Summary A summary of a users recent data
type Summary struct {
	// Config Summary schema version and calculation configuration
	Config *Config `json:"config,omitempty"`

	// Dates dates tracked for summary calculation
	Dates *Dates         `json:"dates,omitempty"`
	Stats *Summary_Stats `json:"stats,omitempty"`

	// Type Field which contains a summary type string.
	Type *SummaryTypeSchema `json:"type,omitempty"`

	// UserId String representation of a Tidepool User ID. Old style IDs are 10-digit strings consisting of only hexadeximcal digits. New style IDs are 36-digit [UUID v4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random))
	UserId *TidepoolUserId `json:"userId,omitempty"`
}

// Summary_Stats defines model for Summary.Stats.
type Summary_Stats struct {
	union json.RawMessage
}

// SummaryTypeSchema Field which contains a summary type string.
type SummaryTypeSchema = string

// TidepoolUserId String representation of a Tidepool User ID. Old style IDs are 10-digit strings consisting of only hexadeximcal digits. New style IDs are 36-digit [UUID v4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random))
type TidepoolUserId = string

// SummaryType Field which contains a summary type string.
type SummaryType = SummaryTypeSchema

// UserId String representation of a Tidepool User ID. Old style IDs are 10-digit strings consisting of only hexadeximcal digits. New style IDs are 36-digit [UUID v4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random))
type UserId = TidepoolUserId

// AsBGMBucketData returns the union data inside the Bucket_Data as a BGMBucketData
func (t Bucket_Data) AsBGMBucketData() (BGMBucketData, error) {
	var body BGMBucketData
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBGMBucketData overwrites any union data inside the Bucket_Data as the provided BGMBucketData
func (t *Bucket_Data) FromBGMBucketData(v BGMBucketData) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBGMBucketData performs a merge with any union data inside the Bucket_Data, using the provided BGMBucketData
func (t *Bucket_Data) MergeBGMBucketData(v BGMBucketData) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsCGMBucketData returns the union data inside the Bucket_Data as a CGMBucketData
func (t Bucket_Data) AsCGMBucketData() (CGMBucketData, error) {
	var body CGMBucketData
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCGMBucketData overwrites any union data inside the Bucket_Data as the provided CGMBucketData
func (t *Bucket_Data) FromCGMBucketData(v CGMBucketData) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCGMBucketData performs a merge with any union data inside the Bucket_Data, using the provided CGMBucketData
func (t *Bucket_Data) MergeCGMBucketData(v CGMBucketData) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Bucket_Data) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Bucket_Data) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCGMStats returns the union data inside the Summary_Stats as a CGMStats
func (t Summary_Stats) AsCGMStats() (CGMStats, error) {
	var body CGMStats
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCGMStats overwrites any union data inside the Summary_Stats as the provided CGMStats
func (t *Summary_Stats) FromCGMStats(v CGMStats) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCGMStats performs a merge with any union data inside the Summary_Stats, using the provided CGMStats
func (t *Summary_Stats) MergeCGMStats(v CGMStats) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsBGMStats returns the union data inside the Summary_Stats as a BGMStats
func (t Summary_Stats) AsBGMStats() (BGMStats, error) {
	var body BGMStats
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBGMStats overwrites any union data inside the Summary_Stats as the provided BGMStats
func (t *Summary_Stats) FromBGMStats(v BGMStats) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBGMStats performs a merge with any union data inside the Summary_Stats, using the provided BGMStats
func (t *Summary_Stats) MergeBGMStats(v BGMStats) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Summary_Stats) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Summary_Stats) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}