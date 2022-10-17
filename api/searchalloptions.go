// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"
	"time"
)

type SearchAllOption struct {
	f func(*searchAllOptionImpl)
	s string
}

func (o SearchAllOption) String() string { return o.s }

type SearchAllOptions interface {
	Verbose() bool
	HasVerbose() bool
	DebugFailures() bool
	HasDebugFailures() bool
	OriginalTerm() string
	HasOriginalTerm() bool
	Date() time.Time
	HasDate() bool
	IntentModifiedTerm() string
	HasIntentModifiedTerm() bool
	Covers() int
	HasCovers() bool
	Latitude() float32
	HasLatitude() bool
	Longitude() float32
	HasLongitude() bool
	MetroID() int
	HasMetroID() bool
	Threads() int
	HasThreads() bool
	StartPage() int
	HasStartPage() bool
}

func SearchAllVerbose(verbose bool) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.SearchAllVerbose(bool %+v)}", verbose)}
}
func SearchAllVerboseFlag(verbose *bool) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.SearchAllVerbose(bool %+v)}", verbose)}
}

func SearchAllDebugFailures(debugFailures bool) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}, fmt.Sprintf("api.SearchAllDebugFailures(bool %+v)}", debugFailures)}
}
func SearchAllDebugFailuresFlag(debugFailures *bool) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}, fmt.Sprintf("api.SearchAllDebugFailures(bool %+v)}", debugFailures)}
}

func SearchAllOriginalTerm(originalTerm string) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		opts.has_originalTerm = true
		opts.originalTerm = originalTerm
	}, fmt.Sprintf("api.SearchAllOriginalTerm(string %+v)}", originalTerm)}
}
func SearchAllOriginalTermFlag(originalTerm *string) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		if originalTerm == nil {
			return
		}
		opts.has_originalTerm = true
		opts.originalTerm = *originalTerm
	}, fmt.Sprintf("api.SearchAllOriginalTerm(string %+v)}", originalTerm)}
}

func SearchAllDate(date time.Time) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		opts.has_date = true
		opts.date = date
	}, fmt.Sprintf("api.SearchAllDate(time.Time %+v)}", date)}
}
func SearchAllDateFlag(date *time.Time) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		if date == nil {
			return
		}
		opts.has_date = true
		opts.date = *date
	}, fmt.Sprintf("api.SearchAllDate(time.Time %+v)}", date)}
}

func SearchAllIntentModifiedTerm(intentModifiedTerm string) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		opts.has_intentModifiedTerm = true
		opts.intentModifiedTerm = intentModifiedTerm
	}, fmt.Sprintf("api.SearchAllIntentModifiedTerm(string %+v)}", intentModifiedTerm)}
}
func SearchAllIntentModifiedTermFlag(intentModifiedTerm *string) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		if intentModifiedTerm == nil {
			return
		}
		opts.has_intentModifiedTerm = true
		opts.intentModifiedTerm = *intentModifiedTerm
	}, fmt.Sprintf("api.SearchAllIntentModifiedTerm(string %+v)}", intentModifiedTerm)}
}

func SearchAllCovers(covers int) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		opts.has_covers = true
		opts.covers = covers
	}, fmt.Sprintf("api.SearchAllCovers(int %+v)}", covers)}
}
func SearchAllCoversFlag(covers *int) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		if covers == nil {
			return
		}
		opts.has_covers = true
		opts.covers = *covers
	}, fmt.Sprintf("api.SearchAllCovers(int %+v)}", covers)}
}

func SearchAllLatitude(latitude float32) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}, fmt.Sprintf("api.SearchAllLatitude(float32 %+v)}", latitude)}
}
func SearchAllLatitudeFlag(latitude *float32) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}, fmt.Sprintf("api.SearchAllLatitude(float32 %+v)}", latitude)}
}

func SearchAllLongitude(longitude float32) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}, fmt.Sprintf("api.SearchAllLongitude(float32 %+v)}", longitude)}
}
func SearchAllLongitudeFlag(longitude *float32) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}, fmt.Sprintf("api.SearchAllLongitude(float32 %+v)}", longitude)}
}

func SearchAllMetroID(metroID int) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		opts.has_metroID = true
		opts.metroID = metroID
	}, fmt.Sprintf("api.SearchAllMetroID(int %+v)}", metroID)}
}
func SearchAllMetroIDFlag(metroID *int) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		if metroID == nil {
			return
		}
		opts.has_metroID = true
		opts.metroID = *metroID
	}, fmt.Sprintf("api.SearchAllMetroID(int %+v)}", metroID)}
}

func SearchAllThreads(threads int) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		opts.has_threads = true
		opts.threads = threads
	}, fmt.Sprintf("api.SearchAllThreads(int %+v)}", threads)}
}
func SearchAllThreadsFlag(threads *int) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		if threads == nil {
			return
		}
		opts.has_threads = true
		opts.threads = *threads
	}, fmt.Sprintf("api.SearchAllThreads(int %+v)}", threads)}
}

func SearchAllStartPage(startPage int) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		opts.has_startPage = true
		opts.startPage = startPage
	}, fmt.Sprintf("api.SearchAllStartPage(int %+v)}", startPage)}
}
func SearchAllStartPageFlag(startPage *int) SearchAllOption {
	return SearchAllOption{func(opts *searchAllOptionImpl) {
		if startPage == nil {
			return
		}
		opts.has_startPage = true
		opts.startPage = *startPage
	}, fmt.Sprintf("api.SearchAllStartPage(int %+v)}", startPage)}
}

type searchAllOptionImpl struct {
	verbose                bool
	has_verbose            bool
	debugFailures          bool
	has_debugFailures      bool
	originalTerm           string
	has_originalTerm       bool
	date                   time.Time
	has_date               bool
	intentModifiedTerm     string
	has_intentModifiedTerm bool
	covers                 int
	has_covers             bool
	latitude               float32
	has_latitude           bool
	longitude              float32
	has_longitude          bool
	metroID                int
	has_metroID            bool
	threads                int
	has_threads            bool
	startPage              int
	has_startPage          bool
}

func (s *searchAllOptionImpl) Verbose() bool               { return s.verbose }
func (s *searchAllOptionImpl) HasVerbose() bool            { return s.has_verbose }
func (s *searchAllOptionImpl) DebugFailures() bool         { return s.debugFailures }
func (s *searchAllOptionImpl) HasDebugFailures() bool      { return s.has_debugFailures }
func (s *searchAllOptionImpl) OriginalTerm() string        { return s.originalTerm }
func (s *searchAllOptionImpl) HasOriginalTerm() bool       { return s.has_originalTerm }
func (s *searchAllOptionImpl) Date() time.Time             { return s.date }
func (s *searchAllOptionImpl) HasDate() bool               { return s.has_date }
func (s *searchAllOptionImpl) IntentModifiedTerm() string  { return s.intentModifiedTerm }
func (s *searchAllOptionImpl) HasIntentModifiedTerm() bool { return s.has_intentModifiedTerm }
func (s *searchAllOptionImpl) Covers() int                 { return s.covers }
func (s *searchAllOptionImpl) HasCovers() bool             { return s.has_covers }
func (s *searchAllOptionImpl) Latitude() float32           { return s.latitude }
func (s *searchAllOptionImpl) HasLatitude() bool           { return s.has_latitude }
func (s *searchAllOptionImpl) Longitude() float32          { return s.longitude }
func (s *searchAllOptionImpl) HasLongitude() bool          { return s.has_longitude }
func (s *searchAllOptionImpl) MetroID() int                { return s.metroID }
func (s *searchAllOptionImpl) HasMetroID() bool            { return s.has_metroID }
func (s *searchAllOptionImpl) Threads() int                { return s.threads }
func (s *searchAllOptionImpl) HasThreads() bool            { return s.has_threads }
func (s *searchAllOptionImpl) StartPage() int              { return s.startPage }
func (s *searchAllOptionImpl) HasStartPage() bool          { return s.has_startPage }

type SearchAllParams struct {
	Term               string    `json:"term" required:"true"`
	Verbose            bool      `json:"verbose"`
	DebugFailures      bool      `json:"debug_failures"`
	OriginalTerm       string    `json:"original_term"`
	Date               time.Time `json:"date"`
	IntentModifiedTerm string    `json:"intent_modified_term"`
	Covers             int       `json:"covers"`
	Latitude           float32   `json:"latitude"`
	Longitude          float32   `json:"longitude"`
	MetroID            int       `json:"metro_id"`
	Threads            int       `json:"threads"`
	StartPage          int       `json:"start_page"`
}

func (o SearchAllParams) Options() []SearchAllOption {
	return []SearchAllOption{
		SearchAllVerbose(o.Verbose),
		SearchAllDebugFailures(o.DebugFailures),
		SearchAllOriginalTerm(o.OriginalTerm),
		SearchAllDate(o.Date),
		SearchAllIntentModifiedTerm(o.IntentModifiedTerm),
		SearchAllCovers(o.Covers),
		SearchAllLatitude(o.Latitude),
		SearchAllLongitude(o.Longitude),
		SearchAllMetroID(o.MetroID),
		SearchAllThreads(o.Threads),
		SearchAllStartPage(o.StartPage),
	}
}

func makeSearchAllOptionImpl(opts ...SearchAllOption) *searchAllOptionImpl {
	res := &searchAllOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeSearchAllOptions(opts ...SearchAllOption) SearchAllOptions {
	return makeSearchAllOptionImpl(opts...)
}
