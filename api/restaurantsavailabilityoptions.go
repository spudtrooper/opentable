// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "time"

type RestaurantsAvailabilityOption func(*restaurantsAvailabilityOptionImpl)

type RestaurantsAvailabilityOptions interface {
	Verbose() bool
	HasVerbose() bool
	RestaurantIDs() []int
	HasRestaurantIDs() bool
	OnlyPop() bool
	HasOnlyPop() bool
	RequestNewAvailability() bool
	HasRequestNewAvailability() bool
	ForwardDays() int
	HasForwardDays() bool
	RequireTimes() bool
	HasRequireTimes() bool
	PartySize() int
	HasPartySize() bool
	DatabaseRegion() string
	HasDatabaseRegion() bool
	Date() time.Time
	HasDate() bool
}

func RestaurantsAvailabilityVerbose(verbose bool) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func RestaurantsAvailabilityVerboseFlag(verbose *bool) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

func RestaurantsAvailabilityRestaurantIDs(restaurantIDs []int) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_restaurantIDs = true
		opts.restaurantIDs = restaurantIDs
	}
}
func RestaurantsAvailabilityRestaurantIDsFlag(restaurantIDs *[]int) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		if restaurantIDs == nil {
			return
		}
		opts.has_restaurantIDs = true
		opts.restaurantIDs = *restaurantIDs
	}
}

func RestaurantsAvailabilityOnlyPop(onlyPop bool) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_onlyPop = true
		opts.onlyPop = onlyPop
	}
}
func RestaurantsAvailabilityOnlyPopFlag(onlyPop *bool) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		if onlyPop == nil {
			return
		}
		opts.has_onlyPop = true
		opts.onlyPop = *onlyPop
	}
}

func RestaurantsAvailabilityRequestNewAvailability(requestNewAvailability bool) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_requestNewAvailability = true
		opts.requestNewAvailability = requestNewAvailability
	}
}
func RestaurantsAvailabilityRequestNewAvailabilityFlag(requestNewAvailability *bool) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		if requestNewAvailability == nil {
			return
		}
		opts.has_requestNewAvailability = true
		opts.requestNewAvailability = *requestNewAvailability
	}
}

func RestaurantsAvailabilityForwardDays(forwardDays int) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_forwardDays = true
		opts.forwardDays = forwardDays
	}
}
func RestaurantsAvailabilityForwardDaysFlag(forwardDays *int) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		if forwardDays == nil {
			return
		}
		opts.has_forwardDays = true
		opts.forwardDays = *forwardDays
	}
}

func RestaurantsAvailabilityRequireTimes(requireTimes bool) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_requireTimes = true
		opts.requireTimes = requireTimes
	}
}
func RestaurantsAvailabilityRequireTimesFlag(requireTimes *bool) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		if requireTimes == nil {
			return
		}
		opts.has_requireTimes = true
		opts.requireTimes = *requireTimes
	}
}

func RestaurantsAvailabilityPartySize(partySize int) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_partySize = true
		opts.partySize = partySize
	}
}
func RestaurantsAvailabilityPartySizeFlag(partySize *int) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		if partySize == nil {
			return
		}
		opts.has_partySize = true
		opts.partySize = *partySize
	}
}

func RestaurantsAvailabilityDatabaseRegion(databaseRegion string) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_databaseRegion = true
		opts.databaseRegion = databaseRegion
	}
}
func RestaurantsAvailabilityDatabaseRegionFlag(databaseRegion *string) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		if databaseRegion == nil {
			return
		}
		opts.has_databaseRegion = true
		opts.databaseRegion = *databaseRegion
	}
}

func RestaurantsAvailabilityDate(date time.Time) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		opts.has_date = true
		opts.date = date
	}
}
func RestaurantsAvailabilityDateFlag(date *time.Time) RestaurantsAvailabilityOption {
	return func(opts *restaurantsAvailabilityOptionImpl) {
		if date == nil {
			return
		}
		opts.has_date = true
		opts.date = *date
	}
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

func (r *restaurantsAvailabilityOptionImpl) Verbose() bool          { return r.verbose }
func (r *restaurantsAvailabilityOptionImpl) HasVerbose() bool       { return r.has_verbose }
func (r *restaurantsAvailabilityOptionImpl) RestaurantIDs() []int   { return r.restaurantIDs }
func (r *restaurantsAvailabilityOptionImpl) HasRestaurantIDs() bool { return r.has_restaurantIDs }
func (r *restaurantsAvailabilityOptionImpl) OnlyPop() bool          { return r.onlyPop }
func (r *restaurantsAvailabilityOptionImpl) HasOnlyPop() bool       { return r.has_onlyPop }
func (r *restaurantsAvailabilityOptionImpl) RequestNewAvailability() bool {
	return r.requestNewAvailability
}
func (r *restaurantsAvailabilityOptionImpl) HasRequestNewAvailability() bool {
	return r.has_requestNewAvailability
}
func (r *restaurantsAvailabilityOptionImpl) ForwardDays() int        { return r.forwardDays }
func (r *restaurantsAvailabilityOptionImpl) HasForwardDays() bool    { return r.has_forwardDays }
func (r *restaurantsAvailabilityOptionImpl) RequireTimes() bool      { return r.requireTimes }
func (r *restaurantsAvailabilityOptionImpl) HasRequireTimes() bool   { return r.has_requireTimes }
func (r *restaurantsAvailabilityOptionImpl) PartySize() int          { return r.partySize }
func (r *restaurantsAvailabilityOptionImpl) HasPartySize() bool      { return r.has_partySize }
func (r *restaurantsAvailabilityOptionImpl) DatabaseRegion() string  { return r.databaseRegion }
func (r *restaurantsAvailabilityOptionImpl) HasDatabaseRegion() bool { return r.has_databaseRegion }
func (r *restaurantsAvailabilityOptionImpl) Date() time.Time         { return r.date }
func (r *restaurantsAvailabilityOptionImpl) HasDate() bool           { return r.has_date }

func makeRestaurantsAvailabilityOptionImpl(opts ...RestaurantsAvailabilityOption) *restaurantsAvailabilityOptionImpl {
	res := &restaurantsAvailabilityOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeRestaurantsAvailabilityOptions(opts ...RestaurantsAvailabilityOption) RestaurantsAvailabilityOptions {
	return makeRestaurantsAvailabilityOptionImpl(opts...)
}
