// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type FindMenuItemOption func(*findMenuItemOptionImpl)

type FindMenuItemOptions interface {
	Verbose() bool
	HasVerbose() bool
	Latitude() float32
	HasLatitude() bool
	Longitude() float32
	HasLongitude() bool
	MetroID() int
	HasMetroID() bool
}

func FindMenuItemVerbose(verbose bool) FindMenuItemOption {
	return func(opts *findMenuItemOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func FindMenuItemVerboseFlag(verbose *bool) FindMenuItemOption {
	return func(opts *findMenuItemOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
}

func FindMenuItemLatitude(latitude float32) FindMenuItemOption {
	return func(opts *findMenuItemOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}
}
func FindMenuItemLatitudeFlag(latitude *float32) FindMenuItemOption {
	return func(opts *findMenuItemOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}
}

func FindMenuItemLongitude(longitude float32) FindMenuItemOption {
	return func(opts *findMenuItemOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}
}
func FindMenuItemLongitudeFlag(longitude *float32) FindMenuItemOption {
	return func(opts *findMenuItemOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}
}

func FindMenuItemMetroID(metroID int) FindMenuItemOption {
	return func(opts *findMenuItemOptionImpl) {
		opts.has_metroID = true
		opts.metroID = metroID
	}
}
func FindMenuItemMetroIDFlag(metroID *int) FindMenuItemOption {
	return func(opts *findMenuItemOptionImpl) {
		if metroID == nil {
			return
		}
		opts.has_metroID = true
		opts.metroID = *metroID
	}
}

type findMenuItemOptionImpl struct {
	verbose       bool
	has_verbose   bool
	latitude      float32
	has_latitude  bool
	longitude     float32
	has_longitude bool
	metroID       int
	has_metroID   bool
}

func (f *findMenuItemOptionImpl) Verbose() bool      { return f.verbose }
func (f *findMenuItemOptionImpl) HasVerbose() bool   { return f.has_verbose }
func (f *findMenuItemOptionImpl) Latitude() float32  { return f.latitude }
func (f *findMenuItemOptionImpl) HasLatitude() bool  { return f.has_latitude }
func (f *findMenuItemOptionImpl) Longitude() float32 { return f.longitude }
func (f *findMenuItemOptionImpl) HasLongitude() bool { return f.has_longitude }
func (f *findMenuItemOptionImpl) MetroID() int       { return f.metroID }
func (f *findMenuItemOptionImpl) HasMetroID() bool   { return f.has_metroID }

type FindMenuItemParams struct {
	Term      string  `json:"term" required:"true"`
	Verbose   bool    `json:"verbose"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	MetroID   int     `json:"metro_id"`
}

func (o FindMenuItemParams) Options() []FindMenuItemOption {
	return []FindMenuItemOption{
		FindMenuItemVerbose(o.Verbose),
		FindMenuItemLatitude(o.Latitude),
		FindMenuItemLongitude(o.Longitude),
		FindMenuItemMetroID(o.MetroID),
	}
}

func makeFindMenuItemOptionImpl(opts ...FindMenuItemOption) *findMenuItemOptionImpl {
	res := &findMenuItemOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeFindMenuItemOptions(opts ...FindMenuItemOption) FindMenuItemOptions {
	return makeFindMenuItemOptionImpl(opts...)
}
