// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type FindMenuItemOption struct {
	f func(*findMenuItemOptionImpl)
	s string
}

func (o FindMenuItemOption) String() string { return o.s }

type FindMenuItemOptions interface {
	Latitude() float32
	HasLatitude() bool
	Longitude() float32
	HasLongitude() bool
	MetroID() int
	HasMetroID() bool
	Verbose() bool
	HasVerbose() bool
}

func FindMenuItemLatitude(latitude float32) FindMenuItemOption {
	return FindMenuItemOption{func(opts *findMenuItemOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}, fmt.Sprintf("api.FindMenuItemLatitude(float32 %+v)", latitude)}
}
func FindMenuItemLatitudeFlag(latitude *float32) FindMenuItemOption {
	return FindMenuItemOption{func(opts *findMenuItemOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}, fmt.Sprintf("api.FindMenuItemLatitude(float32 %+v)", latitude)}
}

func FindMenuItemLongitude(longitude float32) FindMenuItemOption {
	return FindMenuItemOption{func(opts *findMenuItemOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}, fmt.Sprintf("api.FindMenuItemLongitude(float32 %+v)", longitude)}
}
func FindMenuItemLongitudeFlag(longitude *float32) FindMenuItemOption {
	return FindMenuItemOption{func(opts *findMenuItemOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}, fmt.Sprintf("api.FindMenuItemLongitude(float32 %+v)", longitude)}
}

func FindMenuItemMetroID(metroID int) FindMenuItemOption {
	return FindMenuItemOption{func(opts *findMenuItemOptionImpl) {
		opts.has_metroID = true
		opts.metroID = metroID
	}, fmt.Sprintf("api.FindMenuItemMetroID(int %+v)", metroID)}
}
func FindMenuItemMetroIDFlag(metroID *int) FindMenuItemOption {
	return FindMenuItemOption{func(opts *findMenuItemOptionImpl) {
		if metroID == nil {
			return
		}
		opts.has_metroID = true
		opts.metroID = *metroID
	}, fmt.Sprintf("api.FindMenuItemMetroID(int %+v)", metroID)}
}

func FindMenuItemVerbose(verbose bool) FindMenuItemOption {
	return FindMenuItemOption{func(opts *findMenuItemOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.FindMenuItemVerbose(bool %+v)", verbose)}
}
func FindMenuItemVerboseFlag(verbose *bool) FindMenuItemOption {
	return FindMenuItemOption{func(opts *findMenuItemOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.FindMenuItemVerbose(bool %+v)", verbose)}
}

type findMenuItemOptionImpl struct {
	latitude      float32
	has_latitude  bool
	longitude     float32
	has_longitude bool
	metroID       int
	has_metroID   bool
	verbose       bool
	has_verbose   bool
}

func (f *findMenuItemOptionImpl) Latitude() float32  { return f.latitude }
func (f *findMenuItemOptionImpl) HasLatitude() bool  { return f.has_latitude }
func (f *findMenuItemOptionImpl) Longitude() float32 { return f.longitude }
func (f *findMenuItemOptionImpl) HasLongitude() bool { return f.has_longitude }
func (f *findMenuItemOptionImpl) MetroID() int       { return f.metroID }
func (f *findMenuItemOptionImpl) HasMetroID() bool   { return f.has_metroID }
func (f *findMenuItemOptionImpl) Verbose() bool      { return f.verbose }
func (f *findMenuItemOptionImpl) HasVerbose() bool   { return f.has_verbose }

type FindMenuItemParams struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	MetroID   int     `json:"metro_id"`
	Term      string  `json:"term" required:"true"`
	Verbose   bool    `json:"verbose"`
}

func (o FindMenuItemParams) Options() []FindMenuItemOption {
	return []FindMenuItemOption{
		FindMenuItemLatitude(o.Latitude),
		FindMenuItemLongitude(o.Longitude),
		FindMenuItemMetroID(o.MetroID),
		FindMenuItemVerbose(o.Verbose),
	}
}

func makeFindMenuItemOptionImpl(opts ...FindMenuItemOption) *findMenuItemOptionImpl {
	res := &findMenuItemOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeFindMenuItemOptions(opts ...FindMenuItemOption) FindMenuItemOptions {
	return makeFindMenuItemOptionImpl(opts...)
}
