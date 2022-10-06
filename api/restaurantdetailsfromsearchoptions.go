// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "time"

type RestaurantDetailsFromSearchOption func(*restaurantDetailsFromSearchOptionImpl)

type RestaurantDetailsFromSearchOptions interface {
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
	ToSearchOptions() []SearchOption
	ToRestaurantDetailsOptions() []RestaurantDetailsOption
}

func RestaurantDetailsFromSearchVerbose(verbose bool) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func RestaurantDetailsFromSearchVerboseFlag(verbose *bool) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

func RestaurantDetailsFromSearchDebugFailures(debugFailures bool) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		opts.has_debugFailures = true
		opts.debugFailures = debugFailures
	}
}
func RestaurantDetailsFromSearchDebugFailuresFlag(debugFailures *bool) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		if debugFailures == nil {
			return
		}
		opts.has_debugFailures = true
		opts.debugFailures = *debugFailures
	}
}

func RestaurantDetailsFromSearchOriginalTerm(originalTerm string) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		opts.has_originalTerm = true
		opts.originalTerm = originalTerm
	}
}
func RestaurantDetailsFromSearchOriginalTermFlag(originalTerm *string) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		if originalTerm == nil {
			return
		}
		opts.has_originalTerm = true
		opts.originalTerm = *originalTerm
	}
}

func RestaurantDetailsFromSearchDate(date time.Time) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		opts.has_date = true
		opts.date = date
	}
}
func RestaurantDetailsFromSearchDateFlag(date *time.Time) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		if date == nil {
			return
		}
		opts.has_date = true
		opts.date = *date
	}
}

func RestaurantDetailsFromSearchIntentModifiedTerm(intentModifiedTerm string) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		opts.has_intentModifiedTerm = true
		opts.intentModifiedTerm = intentModifiedTerm
	}
}
func RestaurantDetailsFromSearchIntentModifiedTermFlag(intentModifiedTerm *string) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		if intentModifiedTerm == nil {
			return
		}
		opts.has_intentModifiedTerm = true
		opts.intentModifiedTerm = *intentModifiedTerm
	}
}

func RestaurantDetailsFromSearchCovers(covers int) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		opts.has_covers = true
		opts.covers = covers
	}
}
func RestaurantDetailsFromSearchCoversFlag(covers *int) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		if covers == nil {
			return
		}
		opts.has_covers = true
		opts.covers = *covers
	}
}

func RestaurantDetailsFromSearchLatitude(latitude float32) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}
}
func RestaurantDetailsFromSearchLatitudeFlag(latitude *float32) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}
}

func RestaurantDetailsFromSearchLongitude(longitude float32) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}
}
func RestaurantDetailsFromSearchLongitudeFlag(longitude *float32) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}
}

func RestaurantDetailsFromSearchMetroID(metroID int) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		opts.has_metroID = true
		opts.metroID = metroID
	}
}
func RestaurantDetailsFromSearchMetroIDFlag(metroID *int) RestaurantDetailsFromSearchOption {
	return func(opts *restaurantDetailsFromSearchOptionImpl) {
		if metroID == nil {
			return
		}
		opts.has_metroID = true
		opts.metroID = *metroID
	}
}

type restaurantDetailsFromSearchOptionImpl struct {
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
}

func (r *restaurantDetailsFromSearchOptionImpl) Verbose() bool          { return r.verbose }
func (r *restaurantDetailsFromSearchOptionImpl) HasVerbose() bool       { return r.has_verbose }
func (r *restaurantDetailsFromSearchOptionImpl) DebugFailures() bool    { return r.debugFailures }
func (r *restaurantDetailsFromSearchOptionImpl) HasDebugFailures() bool { return r.has_debugFailures }
func (r *restaurantDetailsFromSearchOptionImpl) OriginalTerm() string   { return r.originalTerm }
func (r *restaurantDetailsFromSearchOptionImpl) HasOriginalTerm() bool  { return r.has_originalTerm }
func (r *restaurantDetailsFromSearchOptionImpl) Date() time.Time        { return r.date }
func (r *restaurantDetailsFromSearchOptionImpl) HasDate() bool          { return r.has_date }
func (r *restaurantDetailsFromSearchOptionImpl) IntentModifiedTerm() string {
	return r.intentModifiedTerm
}
func (r *restaurantDetailsFromSearchOptionImpl) HasIntentModifiedTerm() bool {
	return r.has_intentModifiedTerm
}
func (r *restaurantDetailsFromSearchOptionImpl) Covers() int        { return r.covers }
func (r *restaurantDetailsFromSearchOptionImpl) HasCovers() bool    { return r.has_covers }
func (r *restaurantDetailsFromSearchOptionImpl) Latitude() float32  { return r.latitude }
func (r *restaurantDetailsFromSearchOptionImpl) HasLatitude() bool  { return r.has_latitude }
func (r *restaurantDetailsFromSearchOptionImpl) Longitude() float32 { return r.longitude }
func (r *restaurantDetailsFromSearchOptionImpl) HasLongitude() bool { return r.has_longitude }
func (r *restaurantDetailsFromSearchOptionImpl) MetroID() int       { return r.metroID }
func (r *restaurantDetailsFromSearchOptionImpl) HasMetroID() bool   { return r.has_metroID }

type RestaurantDetailsFromSearchParams struct {
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
}

func (o RestaurantDetailsFromSearchParams) Options() []RestaurantDetailsFromSearchOption {
	return []RestaurantDetailsFromSearchOption{
		RestaurantDetailsFromSearchVerbose(o.Verbose),
		RestaurantDetailsFromSearchDebugFailures(o.DebugFailures),
		RestaurantDetailsFromSearchOriginalTerm(o.OriginalTerm),
		RestaurantDetailsFromSearchDate(o.Date),
		RestaurantDetailsFromSearchIntentModifiedTerm(o.IntentModifiedTerm),
		RestaurantDetailsFromSearchCovers(o.Covers),
		RestaurantDetailsFromSearchLatitude(o.Latitude),
		RestaurantDetailsFromSearchLongitude(o.Longitude),
		RestaurantDetailsFromSearchMetroID(o.MetroID),
	}
}

// ToSearchOptions converts RestaurantDetailsFromSearchOption to an array of SearchOption
func (o *restaurantDetailsFromSearchOptionImpl) ToSearchOptions() []SearchOption {
	return []SearchOption{
		SearchVerbose(o.Verbose()),
		SearchDebugFailures(o.DebugFailures()),
		SearchOriginalTerm(o.OriginalTerm()),
		SearchDate(o.Date()),
		SearchIntentModifiedTerm(o.IntentModifiedTerm()),
		SearchCovers(o.Covers()),
		SearchLatitude(o.Latitude()),
		SearchLongitude(o.Longitude()),
		SearchMetroID(o.MetroID()),
	}
}

// ToRestaurantDetailsOptions converts RestaurantDetailsFromSearchOption to an array of RestaurantDetailsOption
func (o *restaurantDetailsFromSearchOptionImpl) ToRestaurantDetailsOptions() []RestaurantDetailsOption {
	return []RestaurantDetailsOption{
		RestaurantDetailsVerbose(o.Verbose()),
		RestaurantDetailsDebugFailures(o.DebugFailures()),
	}
}

func makeRestaurantDetailsFromSearchOptionImpl(opts ...RestaurantDetailsFromSearchOption) *restaurantDetailsFromSearchOptionImpl {
	res := &restaurantDetailsFromSearchOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeRestaurantDetailsFromSearchOptions(opts ...RestaurantDetailsFromSearchOption) RestaurantDetailsFromSearchOptions {
	return makeRestaurantDetailsFromSearchOptionImpl(opts...)
}