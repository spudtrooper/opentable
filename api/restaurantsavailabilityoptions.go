// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"
	"time"
)

type RestaurantsAvailabilityOption struct {
	f func(*restaurantsAvailabilityOptionImpl)
	s string
}

func (o RestaurantsAvailabilityOption) String() string { return o.s }

type RestaurantsAvailabilityOptions interface {
	DatabaseRegion() string
	HasDatabaseRegion() bool
	Date() time.Time
	HasDate() bool
	ForwardDays() int
	HasForwardDays() bool
	OnlyPop() bool
	HasOnlyPop() bool
	PartySize() int
	HasPartySize() bool
	RequestNewAvailability() bool
	HasRequestNewAvailability() bool
	RequireTimes() bool
	HasRequireTimes() bool
	RestaurantIDs() []int
	HasRestaurantIDs() bool
	Verbose() bool
	HasVerbose() bool
}

func RestaurantsAvailabilityDatabaseRegion(databaseRegion string) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_databaseRegion = true
		opts.databaseRegion = databaseRegion
	}, fmt.Sprintf("api.RestaurantsAvailabilityDatabaseRegion(string %+v)}", databaseRegion)}
}
func RestaurantsAvailabilityDatabaseRegionFlag(databaseRegion *string) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		if databaseRegion == nil {
			return
		}
		opts.has_databaseRegion = true
		opts.databaseRegion = *databaseRegion
	}, fmt.Sprintf("api.RestaurantsAvailabilityDatabaseRegion(string %+v)}", databaseRegion)}
}

func RestaurantsAvailabilityDate(date time.Time) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_date = true
		opts.date = date
	}, fmt.Sprintf("api.RestaurantsAvailabilityDate(time.Time %+v)}", date)}
}
func RestaurantsAvailabilityDateFlag(date *time.Time) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		if date == nil {
			return
		}
		opts.has_date = true
		opts.date = *date
	}, fmt.Sprintf("api.RestaurantsAvailabilityDate(time.Time %+v)}", date)}
}

func RestaurantsAvailabilityForwardDays(forwardDays int) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_forwardDays = true
		opts.forwardDays = forwardDays
	}, fmt.Sprintf("api.RestaurantsAvailabilityForwardDays(int %+v)}", forwardDays)}
}
func RestaurantsAvailabilityForwardDaysFlag(forwardDays *int) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		if forwardDays == nil {
			return
		}
		opts.has_forwardDays = true
		opts.forwardDays = *forwardDays
	}, fmt.Sprintf("api.RestaurantsAvailabilityForwardDays(int %+v)}", forwardDays)}
}

func RestaurantsAvailabilityOnlyPop(onlyPop bool) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_onlyPop = true
		opts.onlyPop = onlyPop
	}, fmt.Sprintf("api.RestaurantsAvailabilityOnlyPop(bool %+v)}", onlyPop)}
}
func RestaurantsAvailabilityOnlyPopFlag(onlyPop *bool) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		if onlyPop == nil {
			return
		}
		opts.has_onlyPop = true
		opts.onlyPop = *onlyPop
	}, fmt.Sprintf("api.RestaurantsAvailabilityOnlyPop(bool %+v)}", onlyPop)}
}

func RestaurantsAvailabilityPartySize(partySize int) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_partySize = true
		opts.partySize = partySize
	}, fmt.Sprintf("api.RestaurantsAvailabilityPartySize(int %+v)}", partySize)}
}
func RestaurantsAvailabilityPartySizeFlag(partySize *int) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		if partySize == nil {
			return
		}
		opts.has_partySize = true
		opts.partySize = *partySize
	}, fmt.Sprintf("api.RestaurantsAvailabilityPartySize(int %+v)}", partySize)}
}

func RestaurantsAvailabilityRequestNewAvailability(requestNewAvailability bool) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_requestNewAvailability = true
		opts.requestNewAvailability = requestNewAvailability
	}, fmt.Sprintf("api.RestaurantsAvailabilityRequestNewAvailability(bool %+v)}", requestNewAvailability)}
}
func RestaurantsAvailabilityRequestNewAvailabilityFlag(requestNewAvailability *bool) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		if requestNewAvailability == nil {
			return
		}
		opts.has_requestNewAvailability = true
		opts.requestNewAvailability = *requestNewAvailability
	}, fmt.Sprintf("api.RestaurantsAvailabilityRequestNewAvailability(bool %+v)}", requestNewAvailability)}
}

func RestaurantsAvailabilityRequireTimes(requireTimes bool) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_requireTimes = true
		opts.requireTimes = requireTimes
	}, fmt.Sprintf("api.RestaurantsAvailabilityRequireTimes(bool %+v)}", requireTimes)}
}
func RestaurantsAvailabilityRequireTimesFlag(requireTimes *bool) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		if requireTimes == nil {
			return
		}
		opts.has_requireTimes = true
		opts.requireTimes = *requireTimes
	}, fmt.Sprintf("api.RestaurantsAvailabilityRequireTimes(bool %+v)}", requireTimes)}
}

func RestaurantsAvailabilityRestaurantIDs(restaurantIDs []int) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_restaurantIDs = true
		opts.restaurantIDs = restaurantIDs
	}, fmt.Sprintf("api.RestaurantsAvailabilityRestaurantIDs([]int %+v)}", restaurantIDs)}
}
func RestaurantsAvailabilityRestaurantIDsFlag(restaurantIDs *[]int) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		if restaurantIDs == nil {
			return
		}
		opts.has_restaurantIDs = true
		opts.restaurantIDs = *restaurantIDs
	}, fmt.Sprintf("api.RestaurantsAvailabilityRestaurantIDs([]int %+v)}", restaurantIDs)}
}

func RestaurantsAvailabilityVerbose(verbose bool) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.RestaurantsAvailabilityVerbose(bool %+v)}", verbose)}
}
func RestaurantsAvailabilityVerboseFlag(verbose *bool) RestaurantsAvailabilityOption {
	return RestaurantsAvailabilityOption{func(opts *restaurantsAvailabilityOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.RestaurantsAvailabilityVerbose(bool %+v)}", verbose)}
}

type restaurantsAvailabilityOptionImpl struct {
	verbose                    bool
	has_verbose                bool
	restaurantIDs              []int
	has_restaurantIDs          bool
	onlyPop                    bool
	has_onlyPop                bool
	requestNewAvailability     bool
	has_requestNewAvailability bool
	forwardDays                int
	has_forwardDays            bool
	requireTimes               bool
	has_requireTimes           bool
	partySize                  int
	has_partySize              bool
	databaseRegion             string
	has_databaseRegion         bool
	date                       time.Time
	has_date                   bool
}

func (r *restaurantsAvailabilityOptionImpl) DatabaseRegion() string  { return r.databaseRegion }
func (r *restaurantsAvailabilityOptionImpl) HasDatabaseRegion() bool { return r.has_databaseRegion }
func (r *restaurantsAvailabilityOptionImpl) Date() time.Time         { return r.date }
func (r *restaurantsAvailabilityOptionImpl) HasDate() bool           { return r.has_date }
func (r *restaurantsAvailabilityOptionImpl) ForwardDays() int        { return r.forwardDays }
func (r *restaurantsAvailabilityOptionImpl) HasForwardDays() bool    { return r.has_forwardDays }
func (r *restaurantsAvailabilityOptionImpl) OnlyPop() bool           { return r.onlyPop }
func (r *restaurantsAvailabilityOptionImpl) HasOnlyPop() bool        { return r.has_onlyPop }
func (r *restaurantsAvailabilityOptionImpl) PartySize() int          { return r.partySize }
func (r *restaurantsAvailabilityOptionImpl) HasPartySize() bool      { return r.has_partySize }
func (r *restaurantsAvailabilityOptionImpl) RequestNewAvailability() bool {
	return r.requestNewAvailability
}
func (r *restaurantsAvailabilityOptionImpl) HasRequestNewAvailability() bool {
	return r.has_requestNewAvailability
}
func (r *restaurantsAvailabilityOptionImpl) RequireTimes() bool     { return r.requireTimes }
func (r *restaurantsAvailabilityOptionImpl) HasRequireTimes() bool  { return r.has_requireTimes }
func (r *restaurantsAvailabilityOptionImpl) RestaurantIDs() []int   { return r.restaurantIDs }
func (r *restaurantsAvailabilityOptionImpl) HasRestaurantIDs() bool { return r.has_restaurantIDs }
func (r *restaurantsAvailabilityOptionImpl) Verbose() bool          { return r.verbose }
func (r *restaurantsAvailabilityOptionImpl) HasVerbose() bool       { return r.has_verbose }

type RestaurantsAvailabilityParams struct {
	DatabaseRegion         string    `json:"database_region"`
	Date                   time.Time `json:"date"`
	ForwardDays            int       `json:"forward_days"`
	OnlyPop                bool      `json:"only_pop"`
	PartySize              int       `json:"party_size"`
	RequestNewAvailability bool      `json:"request_new_availability"`
	RequireTimes           bool      `json:"require_times"`
	RestaurantIDs          []int     `json:"restaurant_i_ds"`
	Verbose                bool      `json:"verbose"`
}

func (o RestaurantsAvailabilityParams) Options() []RestaurantsAvailabilityOption {
	return []RestaurantsAvailabilityOption{
		RestaurantsAvailabilityDatabaseRegion(o.DatabaseRegion),
		RestaurantsAvailabilityDate(o.Date),
		RestaurantsAvailabilityForwardDays(o.ForwardDays),
		RestaurantsAvailabilityOnlyPop(o.OnlyPop),
		RestaurantsAvailabilityPartySize(o.PartySize),
		RestaurantsAvailabilityRequestNewAvailability(o.RequestNewAvailability),
		RestaurantsAvailabilityRequireTimes(o.RequireTimes),
		RestaurantsAvailabilityRestaurantIDs(o.RestaurantIDs),
		RestaurantsAvailabilityVerbose(o.Verbose),
	}
}

func makeRestaurantsAvailabilityOptionImpl(opts ...RestaurantsAvailabilityOption) *restaurantsAvailabilityOptionImpl {
	res := &restaurantsAvailabilityOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeRestaurantsAvailabilityOptions(opts ...RestaurantsAvailabilityOption) RestaurantsAvailabilityOptions {
	return makeRestaurantsAvailabilityOptionImpl(opts...)
}
