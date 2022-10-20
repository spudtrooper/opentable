// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"
	"time"
)

type SearchOption struct {
	f func(*searchOptionImpl)
	s string
}

func (o SearchOption) String() string { return o.s }

type SearchOptions interface {
	Covers() int
	HasCovers() bool
	Date() time.Time
	HasDate() bool
	DebugFailures() bool
	HasDebugFailures() bool
	IntentModifiedTerm() string
	HasIntentModifiedTerm() bool
	Latitude() float32
	HasLatitude() bool
	Longitude() float32
	HasLongitude() bool
	MetroID() int
	HasMetroID() bool
	OriginalTerm() string
	HasOriginalTerm() bool
	Verbose() bool
	HasVerbose() bool
}

func SearchCovers(covers int) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_covers = true
		opts.covers = covers
	}, fmt.Sprintf("api.SearchCovers(int %+v)}", covers)}
}
func SearchCoversFlag(covers *int) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if covers == nil {
			return
		}
		opts.has_covers = true
		opts.covers = *covers
	}, fmt.Sprintf("api.SearchCovers(int %+v)}", covers)}
}

func SearchDate(date time.Time) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_date = true
		opts.date = date
	}, fmt.Sprintf("api.SearchDate(time.Time %+v)}", date)}
}
func SearchDateFlag(date *time.Time) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if date == nil {
			return
		}
		opts.has_date = true
		opts.date = *date
	}, fmt.Sprintf("api.SearchDate(time.Time %+v)}", date)}
}

func SearchDebugFailures(debugFailures bool) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}, fmt.Sprintf("api.SearchDebugFailures(bool %+v)}", debugFailures)}
}
func SearchDebugFailuresFlag(debugFailures *bool) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}, fmt.Sprintf("api.SearchDebugFailures(bool %+v)}", debugFailures)}
}

func SearchIntentModifiedTerm(intentModifiedTerm string) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_intentModifiedTerm = true
		opts.intentModifiedTerm = intentModifiedTerm
	}, fmt.Sprintf("api.SearchIntentModifiedTerm(string %+v)}", intentModifiedTerm)}
}
func SearchIntentModifiedTermFlag(intentModifiedTerm *string) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if intentModifiedTerm == nil {
			return
		}
		opts.has_intentModifiedTerm = true
		opts.intentModifiedTerm = *intentModifiedTerm
	}, fmt.Sprintf("api.SearchIntentModifiedTerm(string %+v)}", intentModifiedTerm)}
}

func SearchLatitude(latitude float32) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}, fmt.Sprintf("api.SearchLatitude(float32 %+v)}", latitude)}
}
func SearchLatitudeFlag(latitude *float32) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}, fmt.Sprintf("api.SearchLatitude(float32 %+v)}", latitude)}
}

func SearchLongitude(longitude float32) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}, fmt.Sprintf("api.SearchLongitude(float32 %+v)}", longitude)}
}
func SearchLongitudeFlag(longitude *float32) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}, fmt.Sprintf("api.SearchLongitude(float32 %+v)}", longitude)}
}

func SearchMetroID(metroID int) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_metroID = true
		opts.metroID = metroID
	}, fmt.Sprintf("api.SearchMetroID(int %+v)}", metroID)}
}
func SearchMetroIDFlag(metroID *int) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if metroID == nil {
			return
		}
		opts.has_metroID = true
		opts.metroID = *metroID
	}, fmt.Sprintf("api.SearchMetroID(int %+v)}", metroID)}
}

func SearchOriginalTerm(originalTerm string) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_originalTerm = true
		opts.originalTerm = originalTerm
	}, fmt.Sprintf("api.SearchOriginalTerm(string %+v)}", originalTerm)}
}
func SearchOriginalTermFlag(originalTerm *string) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if originalTerm == nil {
			return
		}
		opts.has_originalTerm = true
		opts.originalTerm = *originalTerm
	}, fmt.Sprintf("api.SearchOriginalTerm(string %+v)}", originalTerm)}
}

func SearchVerbose(verbose bool) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.SearchVerbose(bool %+v)}", verbose)}
}
func SearchVerboseFlag(verbose *bool) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.SearchVerbose(bool %+v)}", verbose)}
}

type searchOptionImpl struct {
	covers                 int
	has_covers             bool
	date                   time.Time
	has_date               bool
	debugFailures          bool
	has_debugFailures      bool
	intentModifiedTerm     string
	has_intentModifiedTerm bool
	latitude               float32
	has_latitude           bool
	longitude              float32
	has_longitude          bool
	metroID                int
	has_metroID            bool
	originalTerm           string
	has_originalTerm       bool
	verbose                bool
	has_verbose            bool
}

func (s *searchOptionImpl) Covers() int                 { return s.covers }
func (s *searchOptionImpl) HasCovers() bool             { return s.has_covers }
func (s *searchOptionImpl) Date() time.Time             { return s.date }
func (s *searchOptionImpl) HasDate() bool               { return s.has_date }
func (s *searchOptionImpl) DebugFailures() bool         { return s.debugFailures }
func (s *searchOptionImpl) HasDebugFailures() bool      { return s.has_debugFailures }
func (s *searchOptionImpl) IntentModifiedTerm() string  { return s.intentModifiedTerm }
func (s *searchOptionImpl) HasIntentModifiedTerm() bool { return s.has_intentModifiedTerm }
func (s *searchOptionImpl) Latitude() float32           { return s.latitude }
func (s *searchOptionImpl) HasLatitude() bool           { return s.has_latitude }
func (s *searchOptionImpl) Longitude() float32          { return s.longitude }
func (s *searchOptionImpl) HasLongitude() bool          { return s.has_longitude }
func (s *searchOptionImpl) MetroID() int                { return s.metroID }
func (s *searchOptionImpl) HasMetroID() bool            { return s.has_metroID }
func (s *searchOptionImpl) OriginalTerm() string        { return s.originalTerm }
func (s *searchOptionImpl) HasOriginalTerm() bool       { return s.has_originalTerm }
func (s *searchOptionImpl) Verbose() bool               { return s.verbose }
func (s *searchOptionImpl) HasVerbose() bool            { return s.has_verbose }

type SearchParams struct {
	Covers             int       `json:"covers"`
	Date               time.Time `json:"date"`
	DebugFailures      bool      `json:"debug_failures"`
	IntentModifiedTerm string    `json:"intent_modified_term"`
	Latitude           float32   `json:"latitude"`
	Longitude          float32   `json:"longitude"`
	MetroID            int       `json:"metro_id"`
	OriginalTerm       string    `json:"original_term"`
	Term               string    `json:"term" required:"true"`
	Verbose            bool      `json:"verbose"`
}

func (o SearchParams) Options() []SearchOption {
	return []SearchOption{
		SearchCovers(o.Covers),
		SearchDate(o.Date),
		SearchDebugFailures(o.DebugFailures),
		SearchIntentModifiedTerm(o.IntentModifiedTerm),
		SearchLatitude(o.Latitude),
		SearchLongitude(o.Longitude),
		SearchMetroID(o.MetroID),
		SearchOriginalTerm(o.OriginalTerm),
		SearchVerbose(o.Verbose),
	}
}

func makeSearchOptionImpl(opts ...SearchOption) *searchOptionImpl {
	res := &searchOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeSearchOptions(opts ...SearchOption) SearchOptions {
	return makeSearchOptionImpl(opts...)
}
