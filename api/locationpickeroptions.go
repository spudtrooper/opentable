// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type LocationPickerOption struct {
	f func(*locationPickerOptionImpl)
	s string
}

func (o LocationPickerOption) String() string { return o.s }

type LocationPickerOptions interface {
	Tld() string
	HasTld() bool
	MetroID() int
	HasMetroID() bool
	DomainID() int
	HasDomainID() bool
	Verbose() bool
	HasVerbose() bool
	AuthCke() string
	HasAuthCke() bool
}

func LocationPickerTld(tld string) LocationPickerOption {
	return LocationPickerOption{func(opts *locationPickerOptionImpl) {
		opts.has_tld = true
		opts.tld = tld
	}, fmt.Sprintf("api.LocationPickerTld(string %+v)}", tld)}
}
func LocationPickerTldFlag(tld *string) LocationPickerOption {
	return LocationPickerOption{func(opts *locationPickerOptionImpl) {
		if tld == nil {
			return
		}
		opts.has_tld = true
		opts.tld = *tld
	}, fmt.Sprintf("api.LocationPickerTld(string %+v)}", tld)}
}

func LocationPickerMetroID(metroID int) LocationPickerOption {
	return LocationPickerOption{func(opts *locationPickerOptionImpl) {
		opts.has_metroID = true
		opts.metroID = metroID
	}, fmt.Sprintf("api.LocationPickerMetroID(int %+v)}", metroID)}
}
func LocationPickerMetroIDFlag(metroID *int) LocationPickerOption {
	return LocationPickerOption{func(opts *locationPickerOptionImpl) {
		if metroID == nil {
			return
		}
		opts.has_metroID = true
		opts.metroID = *metroID
	}, fmt.Sprintf("api.LocationPickerMetroID(int %+v)}", metroID)}
}

func LocationPickerDomainID(domainID int) LocationPickerOption {
	return LocationPickerOption{func(opts *locationPickerOptionImpl) {
		opts.has_domainID = true
		opts.domainID = domainID
	}, fmt.Sprintf("api.LocationPickerDomainID(int %+v)}", domainID)}
}
func LocationPickerDomainIDFlag(domainID *int) LocationPickerOption {
	return LocationPickerOption{func(opts *locationPickerOptionImpl) {
		if domainID == nil {
			return
		}
		opts.has_domainID = true
		opts.domainID = *domainID
	}, fmt.Sprintf("api.LocationPickerDomainID(int %+v)}", domainID)}
}

func LocationPickerVerbose(verbose bool) LocationPickerOption {
	return LocationPickerOption{func(opts *locationPickerOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.LocationPickerVerbose(bool %+v)}", verbose)}
}
func LocationPickerVerboseFlag(verbose *bool) LocationPickerOption {
	return LocationPickerOption{func(opts *locationPickerOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.LocationPickerVerbose(bool %+v)}", verbose)}
}

func LocationPickerAuthCke(authCke string) LocationPickerOption {
	return LocationPickerOption{func(opts *locationPickerOptionImpl) {
		opts.has_authCke = true
		opts.authCke = authCke
	}, fmt.Sprintf("api.LocationPickerAuthCke(string %+v)}", authCke)}
}
func LocationPickerAuthCkeFlag(authCke *string) LocationPickerOption {
	return LocationPickerOption{func(opts *locationPickerOptionImpl) {
		if authCke == nil {
			return
		}
		opts.has_authCke = true
		opts.authCke = *authCke
	}, fmt.Sprintf("api.LocationPickerAuthCke(string %+v)}", authCke)}
}

type locationPickerOptionImpl struct {
	tld          string
	has_tld      bool
	metroID      int
	has_metroID  bool
	domainID     int
	has_domainID bool
	verbose      bool
	has_verbose  bool
	authCke      string
	has_authCke  bool
}

func (l *locationPickerOptionImpl) Tld() string       { return l.tld }
func (l *locationPickerOptionImpl) HasTld() bool      { return l.has_tld }
func (l *locationPickerOptionImpl) MetroID() int      { return l.metroID }
func (l *locationPickerOptionImpl) HasMetroID() bool  { return l.has_metroID }
func (l *locationPickerOptionImpl) DomainID() int     { return l.domainID }
func (l *locationPickerOptionImpl) HasDomainID() bool { return l.has_domainID }
func (l *locationPickerOptionImpl) Verbose() bool     { return l.verbose }
func (l *locationPickerOptionImpl) HasVerbose() bool  { return l.has_verbose }
func (l *locationPickerOptionImpl) AuthCke() string   { return l.authCke }
func (l *locationPickerOptionImpl) HasAuthCke() bool  { return l.has_authCke }

type LocationPickerParams struct {
	Tld      string `json:"tld"`
	MetroID  int    `json:"metro_id"`
	DomainID int    `json:"domain_id"`
	Verbose  bool   `json:"verbose"`
	AuthCke  string `json:"auth_cke"`
}

func (o LocationPickerParams) Options() []LocationPickerOption {
	return []LocationPickerOption{
		LocationPickerTld(o.Tld),
		LocationPickerMetroID(o.MetroID),
		LocationPickerDomainID(o.DomainID),
		LocationPickerVerbose(o.Verbose),
		LocationPickerAuthCke(o.AuthCke),
	}
}

func makeLocationPickerOptionImpl(opts ...LocationPickerOption) *locationPickerOptionImpl {
	res := &locationPickerOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeLocationPickerOptions(opts ...LocationPickerOption) LocationPickerOptions {
	return makeLocationPickerOptionImpl(opts...)
}
