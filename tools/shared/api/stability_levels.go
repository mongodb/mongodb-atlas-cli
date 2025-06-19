// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	versionRegex = regexp.MustCompile(`^((?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})(?P<upcoming>\.upcoming)?|(?P<preview>preview))$`)
)

type StabilityLevel string

const (
	StabilityLevelPreview  StabilityLevel = "preview"
	StabilityLevelUpcoming StabilityLevel = "upcoming"
	StabilityLevelStable   StabilityLevel = "stable"
)

type VersionDate struct {
	Year  int
	Month time.Month
	Day   int
}

// Returns true if v(this) is less than other
func (v VersionDate) Less(other *VersionDate) bool {
	switch {
	case v.Year < other.Year:
		return true
	case v.Year == other.Year && v.Month < other.Month:
		return true
	case v.Year == other.Year && v.Month == other.Month && v.Day < other.Day:
		return true
	default:
		return false
	}
}

// Returns true if v(this) is equal to other
func (v VersionDate) Equal(other *VersionDate) bool {
	return v.Year == other.Year && v.Month == other.Month && v.Day == other.Day
}

// Implement Stringer interface
func (v VersionDate) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", v.Year, v.Month, v.Day)
}

type Version interface {
	StabilityLevel() StabilityLevel
	// Returns true if v(this) is less than other
	Less(other Version) bool
	Equal(other Version) bool
	String() string
}

func ParseVersion(version string) (Version, error) {
	matches := versionRegex.FindStringSubmatch(version)
	if len(matches) == 0 {
		return nil, fmt.Errorf("invalid version: %s", version)
	}

	// Get the named group indexes
	yearIndex := versionRegex.SubexpIndex("year")
	monthIndex := versionRegex.SubexpIndex("month")
	dayIndex := versionRegex.SubexpIndex("day")
	upcomingIndex := versionRegex.SubexpIndex("upcoming")
	previewIndex := versionRegex.SubexpIndex("preview")

	// Get the named group values
	year := matches[yearIndex]
	month := matches[monthIndex]
	day := matches[dayIndex]
	upcoming := matches[upcomingIndex]
	preview := matches[previewIndex]

	// If the version is a preview, return a PreviewVersion
	if preview != "" {
		if year != "" || month != "" || day != "" {
			return nil, errors.New("preview version cannot have a year, month, or day")
		}

		return NewPreviewVersion(), nil
	}

	// For upcoming and stable versions, year, month, and day are required
	if year == "" || month == "" || day == "" {
		return nil, errors.New("upcoming and stable versions must have a year, month, and day")
	}

	// Convert the year, month, and day to ints
	// We know that they're in the correct format (because the regex matched), so we can ignore the error
	yearInt, _ := strconv.Atoi(year)
	monthInt, _ := strconv.Atoi(month)
	dayInt, _ := strconv.Atoi(day)

	// If the version is an upcoming version, return an UpcomingVersion
	if upcoming != "" {
		return NewUpcomingVersion(yearInt, time.Month(monthInt), dayInt), nil
	}

	// If the version is a stable version, return a StableVersion
	return NewStableVersion(yearInt, time.Month(monthInt), dayInt), nil
}

type PreviewVersion struct {
}

func NewPreviewVersion() PreviewVersion {
	return PreviewVersion{}
}

func (PreviewVersion) StabilityLevel() StabilityLevel {
	return StabilityLevelPreview
}

func (PreviewVersion) Less(_ Version) bool {
	// Preview is always last, so it's never less than anything
	// When comparing two preview versions, they're equal, so less is also false
	return false
}

func (PreviewVersion) Equal(other Version) bool {
	// switch cast other Version to preview/upcoming/stable
	switch other.(type) {
	case PreviewVersion:
		// other preview versions are always equal
		return true
	case UpcomingVersion:
		// upcoming versions are never equal to preview versions
		return false
	case StableVersion:
		// stable versions are never equal to preview versions
		return false
	}

	panic("unreachable")
}

func (PreviewVersion) String() string {
	return "preview"
}

type UpcomingVersion struct {
	Date VersionDate
}

func NewUpcomingVersion(year int, month time.Month, day int) UpcomingVersion {
	return UpcomingVersion{Date: VersionDate{Year: year, Month: month, Day: day}}
}

func (UpcomingVersion) StabilityLevel() StabilityLevel {
	return StabilityLevelUpcoming
}

func (v UpcomingVersion) Less(other Version) bool {
	// switch cast other Version to preview/upcoming/stable
	switch o := other.(type) {
	case PreviewVersion:
		// preview versions are always newer (greater) than upcoming versions
		return true
	case UpcomingVersion:
		// for other upcoming versions, compare dates
		return v.Date.Less(&o.Date)
	case StableVersion:
		// for stable versions we compare dates
		// if the date is the same, then the stable version is always older (less) than upcoming versions
		if v.Date.Equal(&o.Date) {
			return false
		}

		return v.Date.Less(&o.Date)
	}

	panic("unreachable")
}

func (v UpcomingVersion) Equal(other Version) bool {
	// switch cast other Version to preview/upcoming/stable
	switch o := other.(type) {
	case PreviewVersion:
		// preview versions are never equal to upcoming versions
		return false
	case UpcomingVersion:
		// for other upcoming versions, compare dates
		return v.Date.Equal(&o.Date)
	case StableVersion:
		// stable versions are never equal to upcoming versions
		return false
	}

	panic("unreachable")
}

func (v UpcomingVersion) String() string {
	return fmt.Sprintf("%s.upcoming", v.Date)
}

type StableVersion struct {
	Date VersionDate
}

func NewStableVersion(year int, month time.Month, day int) StableVersion {
	return StableVersion{Date: VersionDate{Year: year, Month: month, Day: day}}
}

func (StableVersion) StabilityLevel() StabilityLevel {
	return StabilityLevelStable
}

func (v StableVersion) Less(other Version) bool {
	// switch cast other Version to preview/upcoming/stable
	switch o := other.(type) {
	case PreviewVersion:
		// preview versions are always newer (greater) than stable versions
		return true
	case UpcomingVersion:
		// for upcoming versions we compare dates
		// if the date is the same, then the upcoming version is always older (less) than stable versions
		if v.Date.Equal(&o.Date) {
			return true
		}

		return v.Date.Less(&o.Date)
	case StableVersion:
		// for other stable versions, compare dates
		return v.Date.Less(&o.Date)
	}

	panic("unreachable")
}

func (v StableVersion) Equal(other Version) bool {
	// switch cast other Version to preview/upcoming/stable
	switch o := other.(type) {
	case PreviewVersion:
		// preview versions are never equal to stable versions
		return false
	case UpcomingVersion:
		// upcoming versions are never equal to stable versions
		return false
	case StableVersion:
		// for other stable versions, compare dates
		return v.Date.Equal(&o.Date)
	}

	panic("unreachable")
}

func (v StableVersion) String() string {
	return v.Date.String()
}
