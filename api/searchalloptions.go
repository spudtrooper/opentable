// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "time"

type SearchAllOption func(*searchAllOptionImpl)

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
	return func(opts *searchAllOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func SearchAllVerboseFlag(verbose *bool) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

func SearchAllDebugFailures(debugFailures bool) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}
}
func SearchAllDebugFailuresFlag(debugFailures *bool) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}
}

func SearchAllOriginalTerm(originalTerm string) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		opts.has_originalTerm = true
		opts.originalTerm = originalTerm
	}
}
func SearchAllOriginalTermFlag(originalTerm *string) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		if originalTerm == nil {
			return
		}
		opts.has_originalTerm = true
		opts.originalTerm = *originalTerm
	}
}

func SearchAllDate(date time.Time) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		opts.has_date = true
		opts.date = date
	}
}
func SearchAllDateFlag(date *time.Time) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		if date == nil {
			return
		}
		opts.has_date = true
		opts.date = *date
	}
}

func SearchAllIntentModifiedTerm(intentModifiedTerm string) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		opts.has_intentModifiedTerm = true
		opts.intentModifiedTerm = intentModifiedTerm
	}
}
func SearchAllIntentModifiedTermFlag(intentModifiedTerm *string) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		if intentModifiedTerm == nil {
			return
		}
		opts.has_intentModifiedTerm = true
		opts.intentModifiedTerm = *intentModifiedTerm
	}
}

func SearchAllCovers(covers int) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		opts.has_covers = true
		opts.covers = covers
	}
}
func SearchAllCoversFlag(covers *int) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		if covers == nil {
			return
		}
		opts.has_covers = true
		opts.covers = *covers
	}
}

func SearchAllLatitude(latitude float32) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}
}
func SearchAllLatitudeFlag(latitude *float32) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}
}

func SearchAllLongitude(longitude float32) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}
}
func SearchAllLongitudeFlag(longitude *float32) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}
}

func SearchAllMetroID(metroID int) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		opts.has_metroID = true
		opts.metroID = metroID
	}
}
func SearchAllMetroIDFlag(metroID *int) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		if metroID == nil {
			return
		}
		opts.has_metroID = true
		opts.metroID = *metroID
	}
}

func SearchAllThreads(threads int) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		opts.has_threads = true
		opts.threads = threads
	}
}
func SearchAllThreadsFlag(threads *int) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		if threads == nil {
			return
		}
		opts.has_threads = true
		opts.threads = *threads
	}
}

func SearchAllStartPage(startPage int) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		opts.has_startPage = true
		opts.startPage = startPage
	}
}
func SearchAllStartPageFlag(startPage *int) SearchAllOption {
	return func(opts *searchAllOptionImpl) {
		if startPage == nil {
			return
		}
		opts.has_startPage = true
		opts.startPage = *startPage
	}
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

func makeSearchAllOptionImpl(opts ...SearchAllOption) *searchAllOptionImpl {
	res := &searchAllOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeSearchAllOptions(opts ...SearchAllOption) SearchAllOptions {
	return makeSearchAllOptionImpl(opts...)
}
