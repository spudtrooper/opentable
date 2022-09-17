// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type LocationPickerOption func(*locationPickerOptionImpl)

type LocationPickerOptions interface {
	Tld() string
	HasTld() bool
	MetroID() int
	HasMetroID() bool
	DomainID() int
	HasDomainID() bool
	Verbose() bool
	HasVerbose() bool
}

func LocationPickerTld(tld string) LocationPickerOption {
	return func(opts *locationPickerOptionImpl) {
		opts.has_tld = true
		opts.tld = tld
	}
}
func LocationPickerTldFlag(tld *string) LocationPickerOption {
	return func(opts *locationPickerOptionImpl) {
		if tld == nil {
			return
		}
		opts.has_tld = true
		opts.tld = *tld
	}
}

func LocationPickerMetroID(metroID int) LocationPickerOption {
	return func(opts *locationPickerOptionImpl) {
		opts.has_metroID = true
		opts.metroID = metroID
	}
}
func LocationPickerMetroIDFlag(metroID *int) LocationPickerOption {
	return func(opts *locationPickerOptionImpl) {
		if metroID == nil {
			return
		}
		opts.has_metroID = true
		opts.metroID = *metroID
	}
}

func LocationPickerDomainID(domainID int) LocationPickerOption {
	return func(opts *locationPickerOptionImpl) {
		opts.has_domainID = true
		opts.domainID = domainID
	}
}
func LocationPickerDomainIDFlag(domainID *int) LocationPickerOption {
	return func(opts *locationPickerOptionImpl) {
		if domainID == nil {
			return
		}
		opts.has_domainID = true
		opts.domainID = *domainID
	}
}

func LocationPickerVerbose(verbose bool) LocationPickerOption {
	return func(opts *locationPickerOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}
}
func LocationPickerVerboseFlag(verbose *bool) LocationPickerOption {
	return func(opts *locationPickerOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}
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
}

func (l *locationPickerOptionImpl) Tld() string       { return l.tld }
func (l *locationPickerOptionImpl) HasTld() bool      { return l.has_tld }
func (l *locationPickerOptionImpl) MetroID() int      { return l.metroID }
func (l *locationPickerOptionImpl) HasMetroID() bool  { return l.has_metroID }
func (l *locationPickerOptionImpl) DomainID() int     { return l.domainID }
func (l *locationPickerOptionImpl) HasDomainID() bool { return l.has_domainID }
func (l *locationPickerOptionImpl) Verbose() bool     { return l.verbose }
func (l *locationPickerOptionImpl) HasVerbose() bool  { return l.has_verbose }

func makeLocationPickerOptionImpl(opts ...LocationPickerOption) *locationPickerOptionImpl {
	res := &locationPickerOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeLocationPickerOptions(opts ...LocationPickerOption) LocationPickerOptions {
	return makeLocationPickerOptionImpl(opts...)
}
