package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/or"
	"github.com/spudtrooper/goutil/request"
)

//go:generate genopts --function LocationPicker tld:string metroID:int domainID:int verbose:bool
func (c *Client) LocationPicker(optss ...LocationPickerOption) (*LocationPickerInfo, error) {
	opts := MakeLocationPickerOptions(optss...)

	tld := or.String(opts.Tld(), "com")
	metroID := or.Int(opts.MetroID(), 8)
	domainID := or.Int(opts.DomainID(), 1)

	var payload locationPickerPayload

	variables := struct {
		Tld      string `json:"tld"`
		MetroID  int    `json:"metroId"`
		DomainID int    `json:"domainId"`
	}{
		Tld:      tld,
		MetroID:  metroID,
		DomainID: domainID,
	}
	if err := c.query("LocationPicker", true, opts.Verbose(), variables,
		"0874f3998d8c3bf2a59575a27850342d1326e071cfd1366957bd1caecb2a2fd8",
		&payload); err != nil {
		return nil, err
	}

	return payload.Convert(), nil
}

//go:generate genopts --function HeaderUserProfile tld:string verbose
func (c *Client) HeaderUserProfile(optss ...HeaderUserProfileOption) (*HeaderUserProfileInfo, error) {
	opts := MakeHeaderUserProfileOptions(optss...)

	tld := or.String(opts.Tld(), "com")

	var payload headerUserProfilePayload

	variables := struct {
		Tld string `json:"tld"`
	}{
		Tld: tld,
	}
	if err := c.query("HeaderUserProfile", true, opts.Verbose(), variables,
		"5584dc2069459f8d615ea1cc9a162b395dd9f4f79086bb456e2d73c4518f7ee2",
		&payload); err != nil {
		return nil, err
	}

	return payload.Convert(), nil
}

type IsMandatoryBySeating struct {
	TableCategory string
	IsMandatory   bool
}

//go:generate genopts --function RestaurantsAvailability verbose "restaurantIDs:[]int" onlyPop requestNewAvailability "forwardDays:int" requireTimes  "partySize:int" "databaseRegion:string" "date:time.Time"
func (c *Client) RestaurantsAvailability(optss ...RestaurantsAvailabilityOption) (*RestaurantsAvailabilityInfo, error) {
	opts := MakeRestaurantsAvailabilityOptions(optss...)

	onlyPop := opts.OnlyPop()
	requestNewAvailability := true
	if opts.HasRequestNewAvailability() {
		requestNewAvailability = opts.RequestNewAvailability()
	}
	forwardDays := or.Int(opts.ForwardDays(), 0)
	requireTimes := opts.RequireTimes()
	restaurantIDs := or.Ints(opts.RestaurantIDs(), []int{115417, 1214794, 5698, 1147627, 99922, 34939, 98185, 1973, 1142380, 65347, 118102, 110224, 87049, 1048807, 1787})
	d := or.Time(opts.Date(), time.Now())
	date := d.Format("2006-01-02")
	time := d.Format("15:04")
	partySize := or.Int(opts.PartySize(), 2)
	databaseRegion := or.String(opts.DatabaseRegion(), "NA")

	var payload restaurantAvailabilityPayload

	variables := struct {
		OnlyPop                bool   `json:"onlyPop"`
		RequestNewAvailability bool   `json:"requestNewAvailability"`
		ForwardDays            int    `json:"forwardDays"`
		RequireTimes           bool   `json:"requireTimes"`
		RestaurantIds          []int  `json:"restaurantIds"`
		Date                   string `json:"date"`
		Time                   string `json:"time"`
		PartySize              int    `json:"partySize"`
		DatabaseRegion         string `json:"databaseRegion"`
		// RequireTypes           []interface{} `json:"requireTypes"`
	}{
		OnlyPop:                onlyPop,
		RequestNewAvailability: requestNewAvailability,
		ForwardDays:            forwardDays,
		RequireTimes:           requireTimes,
		RestaurantIds:          restaurantIDs,
		Date:                   date,
		Time:                   time,
		PartySize:              partySize,
		DatabaseRegion:         databaseRegion,
		// RequireTypes:[],
	}
	if err := c.query("RestaurantsAvailability", false, opts.Verbose(), variables,
		"9d1bf94dd4657ec0c9280d61a097ca12c1782b529f9f98a1cf2b9eefd5607982",
		&payload); err != nil {
		return nil, err
	}

	return payload.Convert(), nil
}

//go:generate genopts --function ListByURI verbose debugFailures
func (c *Client) RawListByURI(uri string, optss ...ListByURIOption) (*RawListByURIInfo, error) {
	opts := MakeListByURIOptions(optss...)

	res := &RawListByURIInfo{}
	if err := c.rawSearchByURI(uri, res, opts.Verbose(), opts.DebugFailures()); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) ListByURI(uri string, optss ...ListByURIOption) (*SearchInfo, error) {
	opts := MakeListByURIOptions(optss...)

	res := &RawListByURIInfo{}
	if err := c.rawSearchByURI(uri, res, opts.Verbose(), opts.DebugFailures()); err != nil {
		return nil, err
	}
	return res.Convert(), nil
}

//go:generate genopts --function SearchByURI verbose debugFailures
func (c *Client) RawSearchByURI(uri string, optss ...SearchByURIOption) (*RawSearchByURIInfo, error) {
	opts := MakeSearchByURIOptions(optss...)

	res := &RawSearchByURIInfo{}
	if err := c.rawSearchByURI(uri, res, opts.Verbose(), opts.DebugFailures()); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) SearchByURI(uri string, optss ...SearchByURIOption) (*SearchInfo, error) {
	opts := MakeSearchByURIOptions(optss...)

	res := &RawSearchByURIInfo{}
	if err := c.rawSearchByURI(uri, res, opts.Verbose(), opts.DebugFailures()); err != nil {
		return nil, err
	}
	return res.Convert(), nil
}

//go:generate genopts --function Search verbose debugFailures "originalTerm:string" "date:time.Time" "intentModifiedTerm:string" "covers:int" "latitude:float32" "longitude:float32" "metroID:int"
func (c *Client) RawSearch(term string, optss ...SearchOption) (*RawSearchInfo, error) {
	opts := MakeSearchOptions(optss...)

	originalTerm := or.String(opts.OriginalTerm(), term)
	intentModifiedTerm := or.String(opts.IntentModifiedTerm(), term)
	covers := or.Int(opts.Covers(), 2)
	latitude := or.Float32(opts.Latitude(), 40.7683)
	longitude := or.Float32(opts.Longitude(), -73.9802)
	metroID := or.Int(opts.MetroID(), 8)
	d := or.Time(opts.Date(), time.Now())
	dateTime := d.Format("2006-01-02T15%3A00%3A00")
	dateTime = `2022-09-16T21%3A00%3A00` // TODO

	uri := request.MakeURL("https://www.opentable.com/s",
		request.MakeParam("dateTime", dateTime),
		request.MakeParam("term", term),
		request.MakeParam("originCorrelationId", `a148f506-ab44-41aa-8af7-a990ed75bbb9`),
		request.MakeParam("sortBy", `web_conversion`),
		request.MakeParam("originalTerm", originalTerm),
		request.MakeParam("intentModifiedTerm", intentModifiedTerm),
		request.MakeParam("queryUnderstandingType", `default`),
		request.MakeParam("corrid", `cf7a1115-801f-4eaa-8ed6-5892d8aa74a7`),
		request.MakeParam("covers", covers),
		request.MakeParam("latitude", latitude),
		request.MakeParam("longitude", longitude),
		request.MakeParam("metroId", metroID),
	)

	res := &RawSearchInfo{}
	if err := c.rawSearchByURI(uri, res, opts.Verbose(), opts.DebugFailures()); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) Search(term string, optss ...SearchOption) (*SearchInfo, error) {
	res, err := c.RawSearch(term, optss...)
	if err != nil {
		return nil, err
	}
	return res.Convert(), nil
}

func (c *Client) rawSearchByURI(uri string, res interface{}, verbose, debugFailures bool) error {
	convert := func(data []byte) error {
		re := regexp.MustCompile(`(?m)\s*window.__INITIAL_STATE__=JSON.parse\("(.*)"\);\s*`)
		re2 := regexp.MustCompile(`(?m)\s*window.__INITIAL_STATE__=({.*});\s*`)
		for _, line := range strings.Split(string(data), "\n") {
			if m := re.FindStringSubmatch(line); len(m) == 2 {
				jsonContents := m[1]
				// remove escaped quotes and fix escaped backslashes before quotes
				jsonContents = strings.ReplaceAll(jsonContents, `\"`, `"`)
				jsonContents = strings.ReplaceAll(jsonContents, `\\"`, `\\\"`)
				if err := json.Unmarshal([]byte(jsonContents), res); err != nil {
					if debugFailures {
						const f = ".debug/Search.json"
						if writeErr := ioutil.WriteFile(f, []byte(jsonContents), 0644); writeErr != nil {
							err = errors.Wrap(err, writeErr.Error())
						} else {
							log.Printf("wrote json contents to %q for debugging", f)
						}
					}
					return err
				}
			}
			if m := re2.FindStringSubmatch(line); len(m) == 2 {
				jsonContents := m[1]
				if err := json.Unmarshal([]byte(jsonContents), res); err != nil {
					if debugFailures {
						const f = ".debug/Search.json"
						if writeErr := ioutil.WriteFile(f, []byte(jsonContents), 0644); writeErr != nil {
							err = errors.Wrap(err, writeErr.Error())
						} else {
							log.Printf("wrote json contents to %q for debugging", f)
						}
					}
					return err
				}
			}
		}
		return nil
	}

	data, err := c.get(uri, false, verbose)

	if err != nil {
		return errors.Errorf("c.get: %+v", err)
	}

	if err := convert(data); err != nil {
		return errors.Errorf("convert: %+v", err)
	}
	return nil
}

//go:generate genopts --function RestaurantDetails verbose debugFailures
func (c *Client) RestaurantDetails(rest Restaurant, optss ...RestaurantDetailsOption) (*RestaurantDetailsInfo, error) {
	uri := request.MakeURL(rest.ProfileLink,
		request.MakeParam("avt", rest.RestaurantAvailabilityToken),
	)
	return c.RestaurantDetailsByURI(uri, optss...)

}

func (c *Client) RestaurantDetailsFromID(id string, optss ...RestaurantDetailsOption) (*RestaurantDetailsInfo, error) {
	uri := fmt.Sprintf("https://www.opentable.com/r/%s", id)
	return c.RestaurantDetailsByURI(uri, optss...)
}

func (c *Client) RawRestaurantDetailsFromID(id string, optss ...RestaurantDetailsOption) (*RawRestaurantDetails, error) {
	uri := fmt.Sprintf("https://www.opentable.com/r/%s", id)
	return c.RawRestaurantDetailsByURI(uri, optss...)
}

func (c *Client) RestaurantDetailsByURI(uri string, optss ...RestaurantDetailsOption) (*RestaurantDetailsInfo, error) {
	s, err := c.RawRestaurantDetailsByURI(uri, optss...)
	if err != nil {
		return nil, err
	}
	res, err := s.Convert()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) RawRestaurantDetailsByURI(uri string, optss ...RestaurantDetailsOption) (*RawRestaurantDetails, error) {
	opts := MakeRestaurantDetailsOptions(optss...)

	convert := func(data []byte) (*RawRestaurantDetails, error) {
		re := regexp.MustCompile(`(?m)\s*window.__INITIAL_STATE__=({.*});\s*`)
		for _, line := range strings.Split(string(data), "\n") {
			m := re.FindStringSubmatch(line)
			if len(m) == 2 {
				jsonContents := m[1]
				var s RawRestaurantDetails
				if err := json.Unmarshal([]byte(jsonContents), &s); err != nil {
					if opts.DebugFailures() {
						const f = ".debug/RestaurantDetails.json"
						if writeErr := ioutil.WriteFile(f, []byte(jsonContents), 0644); writeErr != nil {
							err = errors.Wrap(err, writeErr.Error())
						} else {
							log.Printf("wrote json contents to %q for debugging", f)
						}
					}
					return nil, err
				}
				return &s, nil
			}
		}
		return nil, nil
	}

	data, err := c.get(uri, true, opts.Verbose())

	if err != nil {
		return nil, errors.Errorf("c.get: %+v", err)
	}

	ret, err := convert(data)
	if err != nil {
		return nil, errors.Errorf("convert: %+v", err)
	}
	return ret, nil
}
