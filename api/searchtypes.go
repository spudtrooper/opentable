package api

type RawListByURIInfo struct {
	// AuthModal struct {
	// 	IsAuthModalOpen                 bool        `json:"isAuthModalOpen"`
	// 	HideCloseButton                 bool        `json:"hideCloseButton"`
	// 	ShouldCloseAuthModalForAuthUser bool        `json:"shouldCloseAuthModalForAuthUser"`
	// 	AuthModalType                   interface{} `json:"authModalType"`
	// 	AuthIframeSrc                   interface{} `json:"authIframeSrc"`
	// 	IsFromBookingFlow               bool        `json:"isFromBookingFlow"`
	// } `json:"authModal"`
	// Authentication struct {
	// 	IsAuthenticated   bool `json:"isAuthenticated"`
	// 	UserType          int  `json:"userType"`
	// 	AuthFinished      bool `json:"authFinished"`
	// 	UseClientSideAuth bool `json:"useClientSideAuth"`
	// } `json:"authentication"`
	// DagAnalytics struct {
	// 	PageName      string      `json:"pageName"`
	// 	CorrelationID interface{} `json:"correlationId"`
	// 	DeviceType    string      `json:"deviceType"`
	// 	RequestTypes  struct {
	// 		MultiSearch struct {
	// 			CorrelationID string `json:"correlationId"`
	// 		} `json:"multiSearch"`
	// 		PopTable struct {
	// 			CorrelationID string `json:"correlationId"`
	// 		} `json:"popTable"`
	// 	} `json:"requestTypes"`
	// 	OriginalRequestID string `json:"originalRequestId"`
	// 	UserAgent         string `json:"userAgent"`
	// } `json:"dagAnalytics"`
	// Domain struct {
	// 	DomainID            int    `json:"domainId"`
	// 	Tld                 string `json:"tld"`
	// 	OtEnv               string `json:"otEnv"`
	// 	DeploymentRegionKey string `json:"deploymentRegionKey"`
	// } `json:"domain"`
	// Experiments struct {
	// 	BUX1978                          int `json:"BUX-1978"`
	// 	Sfx5536                          int `json:"sfx-5536"`
	// 	WEBX255                          int `json:"WEBX-255"`
	// 	WEBX312                          int `json:"WEBX-312"`
	// 	ShowReservationUnavailableBanner int `json:"showReservationUnavailableBanner"`
	// } `json:"experiments"`
	// FacebookTracking struct {
	// 	SessionIDHash interface{} `json:"sessionIdHash"`
	// } `json:"facebookTracking"`
	// FeatureToggles struct {
	// 	ShowPrideLogo         bool `json:"showPrideLogo"`
	// 	IsUploadPhotosEnabled bool `json:"isUploadPhotosEnabled"`
	// } `json:"featureToggles"`
	// Footer struct {
	// 	FormIsSubmitted bool `json:"formIsSubmitted"`
	// 	RequestDidError bool `json:"requestDidError"`
	// 	DiscoverLinks   []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"discoverLinks"`
	// 	OpentableLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"opentableLinks"`
	// 	MoreLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"moreLinks"`
	// 	SiteLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"siteLinks"`
	// 	SocialLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"socialLinks"`
	// 	RestaurateursLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"restaurateursLinks"`
	// 	LegalLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"legalLinks"`
	// } `json:"footer"`
	// Header struct {
	// 	Crumbs []struct {
	// 		Name string `json:"name"`
	// 		Href string `json:"href,omitempty"`
	// 	} `json:"crumbs"`
	// 	MetroDomainID                   int           `json:"metroDomainId"`
	// 	FeaturedMetros                  []interface{} `json:"featuredMetros"`
	// 	MetroID                         int           `json:"metroId"`
	// 	MacroID                         interface{}   `json:"macroId"`
	// 	NeighborhoodID                  interface{}   `json:"neighborhoodId"`
	// 	CurrentMetro                    interface{}   `json:"currentMetro"`
	// 	IsLocationPickerLoading         bool          `json:"isLocationPickerLoading"`
	// 	HideLocationPicker              bool          `json:"hideLocationPicker"`
	// 	TimezoneOffset                  int           `json:"timezoneOffset"`
	// 	ShouldDisplayOverlay            bool          `json:"shouldDisplayOverlay"`
	// 	UseHeaderAuth                   bool          `json:"useHeaderAuth"`
	// 	IsUserProfileLoading            bool          `json:"isUserProfileLoading"`
	// 	UserProfile                     interface{}   `json:"userProfile"`
	// 	UserTransactions                []interface{} `json:"userTransactions"`
	// 	UserInvites                     []interface{} `json:"userInvites"`
	// 	UserNotifications               []interface{} `json:"userNotifications"`
	// 	HasUnreadNotifications          bool          `json:"hasUnreadNotifications"`
	// 	UserNotificationsInitialLoading bool          `json:"userNotificationsInitialLoading"`
	// 	UserNewsfeedStory               interface{}   `json:"userNewsfeedStory"`
	// 	HasUnreadNewsfeedStory          bool          `json:"hasUnreadNewsfeedStory"`
	// 	UserNewsfeedInitialLoading      bool          `json:"userNewsfeedInitialLoading"`
	// 	ShowCovid19Banner               bool          `json:"showCovid19Banner"`
	// 	CovidBannerCTALink              string        `json:"covidBannerCTALink"`
	// 	HideLabel                       bool          `json:"hideLabel"`
	// } `json:"header"`
	// Language struct {
	// 	Locale          string `json:"locale"`
	// 	AcceptLanguage  string `json:"acceptLanguage"`
	// 	DomainLanguages struct {
	// 		En string `json:"en"`
	// 		Fr string `json:"fr"`
	// 		Es string `json:"es"`
	// 		De string `json:"de"`
	// 		Ja string `json:"ja"`
	// 		Nl string `json:"nl"`
	// 		It string `json:"it"`
	// 	} `json:"domainLanguages"`
	// 	CurrencySymbol string `json:"currencySymbol"`
	// } `json:"language"`
	// SearchAutocomplete struct {
	// 	AutocompleteResults   interface{} `json:"autocompleteResults"`
	// 	CorrelationID         interface{} `json:"correlationId"`
	// 	InitialMetroID        int         `json:"initialMetroId"`
	// 	InitialMacroID        interface{} `json:"initialMacroId"`
	// 	InitialNeighborhoodID interface{} `json:"initialNeighborhoodId"`
	// 	Latitude              float64     `json:"latitude"`
	// 	Longitude             float64     `json:"longitude"`
	// 	SelectedItem          interface{} `json:"selectedItem"`
	// 	InputValue            string      `json:"inputValue"`
	// 	RecentSearchItems     interface{} `json:"recentSearchItems"`
	// 	CurrentLocationName   string      `json:"currentLocationName"`
	// 	CachedResults         struct {
	// 	} `json:"cachedResults"`
	// } `json:"searchAutocomplete"`
	// SignalTracking struct {
	// 	Page struct {
	// 		Name          string `json:"name"`
	// 		SeoPageType   string `json:"seoPageType"`
	// 		SeoVertical   string `json:"seoVertical"`
	// 		MarketingName string `json:"marketingName"`
	// 	} `json:"page"`
	// 	Reservation struct {
	// 		Hash     interface{} `json:"hash"`
	// 		ViewInfo interface{} `json:"viewInfo"`
	// 	} `json:"reservation"`
	// } `json:"signalTracking"`
	// SojernTracking struct {
	// 	SessionIDHash interface{} `json:"sessionIdHash"`
	// } `json:"sojernTracking"`
	// UserTracking struct {
	// 	IsUserOptedIn         bool `json:"isUserOptedIn"`
	// 	ConsentLoaded         bool `json:"consentLoaded"`
	// 	ConsentPerformance    bool `json:"consentPerformance"`
	// 	ConsentTargeting      bool `json:"consentTargeting"`
	// 	SessionIDHashesLoaded bool `json:"sessionIdHashesLoaded"`
	// } `json:"userTracking"`
	// Availability struct {
	// 	Loading                 bool `json:"loading"`
	// 	RestaurantsAvailability struct {
	// 	} `json:"restaurantsAvailability"`
	// 	RestaurantsMultiDayAvailability struct {
	// 	} `json:"restaurantsMultiDayAvailability"`
	// 	PopRestaurantsAvailability struct {
	// 	} `json:"popRestaurantsAvailability"`
	// 	AvailabilityTime      string `json:"availabilityTime"`
	// 	AvailabilityDate      string `json:"availabilityDate"`
	// 	AvailabilityPartySize int    `json:"availabilityPartySize"`
	// 	Scrollers             struct {
	// 	} `json:"scrollers"`
	// 	ScrollersLoading struct {
	// 	} `json:"scrollersLoading"`
	// 	AttributionToken string `json:"attributionToken"`
	// } `json:"availability"`
	// Geolocation struct {
	// 	Lat         float64     `json:"lat"`
	// 	Long        float64     `json:"long"`
	// 	ServerLat   interface{} `json:"serverLat"`
	// 	ServerLong  interface{} `json:"serverLong"`
	// 	Metro       interface{} `json:"metro"`
	// 	CountryCode string      `json:"countryCode"`
	// } `json:"geolocation"`
	// Map struct {
	// 	IsFullScreenMode          bool   `json:"isFullScreenMode"`
	// 	IsMapLoaded               bool   `json:"isMapLoaded"`
	// 	IsSearchOnMoveActive      bool   `json:"isSearchOnMoveActive"`
	// 	IsDistanceSelectOpen      bool   `json:"isDistanceSelectOpen"`
	// 	Zoom                      int    `json:"zoom"`
	// 	HasUserInteractedWithMap  bool   `json:"hasUserInteractedWithMap"`
	// 	IsUserLocationLoading     bool   `json:"isUserLocationLoading"`
	// 	GtmPrefix                 string `json:"gtmPrefix"`
	// 	ExpandedRestaurantDetails struct {
	// 		IsLoading  bool        `json:"isLoading"`
	// 		Restaurant interface{} `json:"restaurant"`
	// 	} `json:"expandedRestaurantDetails"`
	// } `json:"map"`
	// MultiDayAvailabilityModal struct {
	// 	IsLoading bool `json:"isLoading"`
	// 	IsOpen    bool `json:"isOpen"`
	// } `json:"multiDayAvailabilityModal"`
	MultiSearch struct {
		// PageType  string      `json:"pageType"`
		// MetroID   int         `json:"metroId"`
		// Latitude  interface{} `json:"latitude"`
		// Longitude interface{} `json:"longitude"`
		// Metro     struct {
		// 	MetroID        int    `json:"metroId"`
		// 	Name           string `json:"name"`
		// 	DomainID       int    `json:"domainId"`
		// 	GeographicData struct {
		// 		Center struct {
		// 			Latitude  float64 `json:"latitude"`
		// 			Longitude float64 `json:"longitude"`
		// 			Typename  string  `json:"__typename"`
		// 		} `json:"center"`
		// 		Typename string `json:"__typename"`
		// 	} `json:"geographicData"`
		// 	Timezone struct {
		// 		OffsetInMinutes int    `json:"offsetInMinutes"`
		// 		Typename        string `json:"__typename"`
		// 	} `json:"timezone"`
		// 	Urls struct {
		// 		StartPageLink struct {
		// 			Link     string `json:"link"`
		// 			Typename string `json:"__typename"`
		// 		} `json:"startPageLink"`
		// 		ListingsPageLink struct {
		// 			Link     string `json:"link"`
		// 			Typename string `json:"__typename"`
		// 		} `json:"listingsPageLink"`
		// 		Typename string `json:"__typename"`
		// 	} `json:"urls"`
		// 	Typename string `json:"__typename"`
		// } `json:"metro"`
		// IsNotFoundError bool `json:"isNotFoundError"`
		// IsServerError   bool `json:"isServerError"`
		// IsLoading       bool `json:"isLoading"`
		// IsFacetsLoading bool `json:"isFacetsLoading"`
		// FacetsLoaded    bool `json:"facetsLoaded"`
		// Facets          struct {
		// 	Prices []int `json:"prices"`
		// 	Times  []struct {
		// 		Key string `json:"key"`
		// 	} `json:"times"`
		// 	DiningTypes []struct {
		// 		Key string `json:"key"`
		// 	} `json:"diningTypes"`
		// 	LegacyAreas []struct {
		// 		ID       int    `json:"id"`
		// 		Name     string `json:"name"`
		// 		Count    int    `json:"count"`
		// 		Type     string `json:"type"`
		// 		Typename string `json:"__typename"`
		// 	} `json:"legacyAreas"`
		// 	Cuisines []struct {
		// 		CuisineID string `json:"cuisineId"`
		// 		Name      string `json:"name"`
		// 		Count     int    `json:"count"`
		// 		Typename  string `json:"__typename"`
		// 	} `json:"cuisines"`
		// 	Tags []struct {
		// 		TagID    string `json:"tagId"`
		// 		Name     string `json:"name"`
		// 		Count    int    `json:"count"`
		// 		Typename string `json:"__typename"`
		// 	} `json:"tags"`
		// 	TableCategories []struct {
		// 		TableCategoryID string `json:"tableCategoryId"`
		// 		Count           int    `json:"count"`
		// 		Typename        string `json:"__typename"`
		// 	} `json:"tableCategories"`
		// 	ExperienceTypes []struct {
		// 		TypeID   int    `json:"typeId"`
		// 		Name     string `json:"name"`
		// 		Typename string `json:"__typename"`
		// 	} `json:"experienceTypes"`
		// 	ProofOfVaccinationRequired bool `json:"proofOfVaccinationRequired"`
		// 	AdditionalDetails          []struct {
		// 		AdditionalDetailID string `json:"additionalDetailId"`
		// 		Name               string `json:"name"`
		// 		Count              int    `json:"count"`
		// 		Typename           string `json:"__typename"`
		// 	} `json:"additionalDetails"`
		// 	DniTags []struct {
		// 		DniTagID string `json:"dniTagId"`
		// 		Count    int    `json:"count"`
		// 		Typename string `json:"__typename"`
		// 	} `json:"dniTags"`
		// 	Typename string `json:"__typename"`
		// } `json:"facets"`
		// Filters struct {
		// 	AdditionalDetailIds        []interface{} `json:"additionalDetailIds"`
		// 	CuisineIds                 []interface{} `json:"cuisineIds"`
		// 	DiningType                 string        `json:"diningType"`
		// 	DniTagIds                  []interface{} `json:"dniTagIds"`
		// 	ExperienceTypeIds          []interface{} `json:"experienceTypeIds"`
		// 	LoyaltyRedemptionTiers     []interface{} `json:"loyaltyRedemptionTiers"`
		// 	NeighborhoodIds            []interface{} `json:"neighborhoodIds"`
		// 	OnlyPopTimes               bool          `json:"onlyPopTimes"`
		// 	PriceBandIds               []interface{} `json:"priceBandIds"`
		// 	ProofOfVaccinationRequired bool          `json:"proofOfVaccinationRequired"`
		// 	RegionIds                  []interface{} `json:"regionIds"`
		// 	TableCategoryIds           []interface{} `json:"tableCategoryIds"`
		// 	TagIds                     []interface{} `json:"tagIds"`
		// } `json:"filters"`
		// FiltersHistory []interface{} `json:"filtersHistory"`
		Restaurants []struct {
			RestaurantID  int    `json:"restaurantId"`
			Name          string `json:"name"`
			Type          string `json:"type"`
			SlotDiscovery string `json:"slotDiscovery"`
			Urls          struct {
				ProfileLink struct {
					Link     string `json:"link"`
					Typename string `json:"__typename"`
				} `json:"profileLink"`
				Typename string `json:"__typename"`
			} `json:"urls"`
			PriceBand struct {
				PriceBandID    int    `json:"priceBandId"`
				CurrencySymbol string `json:"currencySymbol"`
				Name           string `json:"name"`
				Typename       string `json:"__typename"`
			} `json:"priceBand"`
			Neighborhood struct {
				Name     string `json:"name"`
				Typename string `json:"__typename"`
			} `json:"neighborhood"`
			Statistics struct {
				RecentReservationCount int `json:"recentReservationCount"`
				Reviews                struct {
					AllTimeTextReviewCount int `json:"allTimeTextReviewCount"`
					Ratings                struct {
						Overall struct {
							Rating   float64 `json:"rating"`
							Typename string  `json:"__typename"`
						} `json:"overall"`
						Typename string `json:"__typename"`
					} `json:"ratings"`
					Typename string `json:"__typename"`
				} `json:"reviews"`
				Typename string `json:"__typename"`
			} `json:"statistics"`
			PrimaryCuisine struct {
				Name     string `json:"name"`
				Typename string `json:"__typename"`
			} `json:"primaryCuisine"`
			CampaignID int  `json:"campaignId"`
			IsPromoted bool `json:"isPromoted"`
			Features   struct {
				Subtype       interface{} `json:"subtype"`
				Bar           bool        `json:"bar"`
				Counter       bool        `json:"counter"`
				HighTop       bool        `json:"highTop"`
				Outdoor       bool        `json:"outdoor"`
				PickupEnabled interface{} `json:"pickupEnabled"`
				Typename      string      `json:"__typename"`
			} `json:"features"`
			RestaurantAvailabilityToken string `json:"restaurantAvailabilityToken"`
			Groups                      []struct {
				ID                        int         `json:"id"`
				Rids                      interface{} `json:"rids"`
				OffDiscoveryRestaurantIds interface{} `json:"offDiscoveryRestaurantIds"`
				Typename                  string      `json:"__typename"`
			} `json:"groups"`
			Typename string `json:"__typename"`
			IsPinned bool   `json:"isPinned"`
			Photos   struct {
				Profile struct {
					Xsmall struct {
						URL      string `json:"url"`
						Width    int    `json:"width"`
						Typename string `json:"__typename"`
					} `json:"xsmall"`
					Small struct {
						URL      string `json:"url"`
						Width    int    `json:"width"`
						Typename string `json:"__typename"`
					} `json:"small"`
					Legacy struct {
						URL      string `json:"url"`
						Width    int    `json:"width"`
						Typename string `json:"__typename"`
					} `json:"legacy"`
					Medium struct {
						URL      string `json:"url"`
						Width    int    `json:"width"`
						Typename string `json:"__typename"`
					} `json:"medium"`
					Typename string `json:"__typename"`
				} `json:"profile"`
				Typename string `json:"__typename"`
			} `json:"photos"`
			JustAddedDetails struct {
				JustAdded bool   `json:"justAdded"`
				Typename  string `json:"__typename"`
			} `json:"justAddedDetails"`
			Offers      []interface{} `json:"offers"`
			DiningStyle string        `json:"diningStyle"`
			Coordinates struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				Typename  string  `json:"__typename"`
			} `json:"coordinates"`
			Address struct {
				Line1    string `json:"line1"`
				Line2    string `json:"line2"`
				City     string `json:"city"`
				State    string `json:"state"`
				PostCode string `json:"postCode"`
				Typename string `json:"__typename"`
			} `json:"address"`
			Description string `json:"description"`
			TopReview   struct {
				HighlightedText string `json:"highlightedText"`
				Typename        string `json:"__typename"`
			} `json:"topReview"`
			MatchRelevance     interface{} `json:"matchRelevance"`
			HasTakeout         bool        `json:"hasTakeout"`
			ContactInformation struct {
				PhoneNumber          string `json:"phoneNumber"`
				FormattedPhoneNumber string `json:"formattedPhoneNumber"`
				Typename             string `json:"__typename"`
			} `json:"contactInformation"`
			DeliveryPartners []struct {
				DeliveryPartnerID   int    `json:"deliveryPartnerId"`
				DeliveryPartnerName string `json:"deliveryPartnerName"`
				MenuURL             string `json:"menuUrl"`
				Active              bool   `json:"active"`
				Typename            string `json:"__typename"`
			} `json:"deliveryPartners"`
			OrderOnlineLink string `json:"orderOnlineLink"`
		} `json:"restaurants"`
		TotalRestaurantCount int `json:"totalRestaurantCount"`
		PopRestaurants       []struct {
			RestaurantID  int    `json:"restaurantId"`
			Name          string `json:"name"`
			Type          string `json:"type"`
			SlotDiscovery string `json:"slotDiscovery"`
			Urls          struct {
				ProfileLink struct {
					Link     string `json:"link"`
					Typename string `json:"__typename"`
				} `json:"profileLink"`
				Typename string `json:"__typename"`
			} `json:"urls"`
			PriceBand struct {
				PriceBandID    int    `json:"priceBandId"`
				CurrencySymbol string `json:"currencySymbol"`
				Name           string `json:"name"`
				Typename       string `json:"__typename"`
			} `json:"priceBand"`
			Neighborhood struct {
				Name     string `json:"name"`
				Typename string `json:"__typename"`
			} `json:"neighborhood"`
			// Statistics struct {
			// 	RecentReservationCount int `json:"recentReservationCount"`
			// 	Reviews                struct {
			// 		AllTimeTextReviewCount int `json:"allTimeTextReviewCount"`
			// 		Ratings                struct {
			// 			Overall struct {
			// 				Rating   float64 `json:"rating"`
			// 				Typename string  `json:"__typename"`
			// 			} `json:"overall"`
			// 			Typename string `json:"__typename"`
			// 		} `json:"ratings"`
			// 		Typename string `json:"__typename"`
			// 	} `json:"reviews"`
			// 	Typename string `json:"__typename"`
			// } `json:"statistics"`
			PrimaryCuisine struct {
				Name     string `json:"name"`
				Typename string `json:"__typename"`
			} `json:"primaryCuisine"`
			CampaignID interface{} `json:"campaignId"`
			IsPromoted bool        `json:"isPromoted"`
			Features   struct {
				Subtype       interface{} `json:"subtype"`
				Bar           bool        `json:"bar"`
				Counter       bool        `json:"counter"`
				HighTop       bool        `json:"highTop"`
				Outdoor       bool        `json:"outdoor"`
				PickupEnabled interface{} `json:"pickupEnabled"`
				Typename      string      `json:"__typename"`
			} `json:"features"`
			RestaurantAvailabilityToken string        `json:"restaurantAvailabilityToken"`
			Groups                      []interface{} `json:"groups"`
			Typename                    string        `json:"__typename"`
		} `json:"popRestaurants"`
		// TotalPopRestaurantCount   int         `json:"totalPopRestaurantCount"`
		// PageNumber                int         `json:"pageNumber"`
		// Limit                     int         `json:"limit"`
		// PopLimit                  int         `json:"popLimit"`
		// MapLimit                  int         `json:"mapLimit"`
		// PinnedRid                 interface{} `json:"pinnedRid"`
		// FreetextTerm              interface{} `json:"freetextTerm"`
		// OriginalTerm              interface{} `json:"originalTerm"`
		// IntentModifiedTerm        interface{} `json:"intentModifiedTerm"`
		// SortBy                    string      `json:"sortBy"`
		// QueryUnderstandingType    string      `json:"queryUnderstandingType"`
		// IsValidGiftPage           bool        `json:"isValidGiftPage"`
		// DisableMapView            bool        `json:"disableMapView"`
		// IsMobileSearch            bool        `json:"isMobileSearch"`
		// IsMapView                 bool        `json:"isMapView"`
		// ShowCollapsibleDtp        bool        `json:"showCollapsibleDtp"`
		// ShowCollapsibleDtpHeading bool        `json:"showCollapsibleDtpHeading"`
		// OnlyWithOffers            bool        `json:"onlyWithOffers"`
		// Debug                     bool        `json:"debug"`
		// ShouldUseLatLongSearch    bool        `json:"shouldUseLatLongSearch"`
		// AnytimeAvailability       bool        `json:"anytimeAvailability"`
	} `json:"multiSearch"`
}

type RawSearchByURIInfo struct {
	// AuthModal struct {
	// 	IsAuthModalOpen                 bool        `json:"isAuthModalOpen"`
	// 	HideCloseButton                 bool        `json:"hideCloseButton"`
	// 	ShouldCloseAuthModalForAuthUser bool        `json:"shouldCloseAuthModalForAuthUser"`
	// 	AuthModalType                   interface{} `json:"authModalType"`
	// 	AuthIframeSrc                   interface{} `json:"authIframeSrc"`
	// 	IsFromBookingFlow               bool        `json:"isFromBookingFlow"`
	// } `json:"authModal"`
	// Authentication struct {
	// 	IsAuthenticated   bool `json:"isAuthenticated"`
	// 	UserType          int  `json:"userType"`
	// 	AuthFinished      bool `json:"authFinished"`
	// 	UseClientSideAuth bool `json:"useClientSideAuth"`
	// } `json:"authentication"`
	// DagAnalytics struct {
	// 	PageName      string      `json:"pageName"`
	// 	CorrelationID interface{} `json:"correlationId"`
	// 	OriginID      string      `json:"originId"`
	// 	RequestTypes  struct {
	// 		LolzViewAllDefault struct {
	// 			CorrelationID string `json:"correlationId"`
	// 		} `json:"lolzViewAllDefault"`
	// 	} `json:"requestTypes"`
	// 	OriginalRequestID string `json:"originalRequestId"`
	// 	UserAgent         string `json:"userAgent"`
	// } `json:"dagAnalytics"`
	// Domain struct {
	// 	DomainID            int    `json:"domainId"`
	// 	Tld                 string `json:"tld"`
	// 	OtEnv               string `json:"otEnv"`
	// 	DeploymentRegionKey string `json:"deploymentRegionKey"`
	// } `json:"domain"`
	// Experiments struct {
	// 	ShowReservationUnavailableBanner int `json:"showReservationUnavailableBanner"`
	// } `json:"experiments"`
	// FacebookTracking struct {
	// 	SessionIDHash interface{} `json:"sessionIdHash"`
	// } `json:"facebookTracking"`
	// FeatureToggles struct {
	// 	ShowPrideLogo         bool `json:"showPrideLogo"`
	// 	IsUploadPhotosEnabled bool `json:"isUploadPhotosEnabled"`
	// } `json:"featureToggles"`
	// Footer struct {
	// 	FormIsSubmitted bool `json:"formIsSubmitted"`
	// 	RequestDidError bool `json:"requestDidError"`
	// 	DiscoverLinks   []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"discoverLinks"`
	// 	OpentableLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"opentableLinks"`
	// 	MoreLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"moreLinks"`
	// 	SiteLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"siteLinks"`
	// 	SocialLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"socialLinks"`
	// 	RestaurateursLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"restaurateursLinks"`
	// 	LegalLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"legalLinks"`
	// } `json:"footer"`
	// Header struct {
	// 	Crumbs                          []interface{} `json:"crumbs"`
	// 	MetroDomainID                   int           `json:"metroDomainId"`
	// 	FeaturedMetros                  []interface{} `json:"featuredMetros"`
	// 	MetroID                         int           `json:"metroId"`
	// 	MacroID                         interface{}   `json:"macroId"`
	// 	NeighborhoodID                  interface{}   `json:"neighborhoodId"`
	// 	CurrentMetro                    interface{}   `json:"currentMetro"`
	// 	IsLocationPickerLoading         bool          `json:"isLocationPickerLoading"`
	// 	HideLocationPicker              bool          `json:"hideLocationPicker"`
	// 	TimezoneOffset                  int           `json:"timezoneOffset"`
	// 	ShouldDisplayOverlay            bool          `json:"shouldDisplayOverlay"`
	// 	UseHeaderAuth                   bool          `json:"useHeaderAuth"`
	// 	IsUserProfileLoading            bool          `json:"isUserProfileLoading"`
	// 	UserProfile                     interface{}   `json:"userProfile"`
	// 	UserTransactions                []interface{} `json:"userTransactions"`
	// 	UserInvites                     []interface{} `json:"userInvites"`
	// 	UserNotifications               []interface{} `json:"userNotifications"`
	// 	HasUnreadNotifications          bool          `json:"hasUnreadNotifications"`
	// 	UserNotificationsInitialLoading bool          `json:"userNotificationsInitialLoading"`
	// 	UserNewsfeedStory               interface{}   `json:"userNewsfeedStory"`
	// 	HasUnreadNewsfeedStory          bool          `json:"hasUnreadNewsfeedStory"`
	// 	UserNewsfeedInitialLoading      bool          `json:"userNewsfeedInitialLoading"`
	// 	ShowCovid19Banner               bool          `json:"showCovid19Banner"`
	// 	CovidBannerCTALink              string        `json:"covidBannerCTALink"`
	// 	HideLabel                       bool          `json:"hideLabel"`
	// } `json:"header"`
	// Language struct {
	// 	Locale          string `json:"locale"`
	// 	AcceptLanguage  string `json:"acceptLanguage"`
	// 	DomainLanguages struct {
	// 		En string `json:"en"`
	// 		Fr string `json:"fr"`
	// 		Es string `json:"es"`
	// 		De string `json:"de"`
	// 		Ja string `json:"ja"`
	// 		Nl string `json:"nl"`
	// 		It string `json:"it"`
	// 	} `json:"domainLanguages"`
	// 	CurrencySymbol interface{} `json:"currencySymbol"`
	// } `json:"language"`
	// SearchAutocomplete struct {
	// 	AutocompleteResults   interface{} `json:"autocompleteResults"`
	// 	CorrelationID         interface{} `json:"correlationId"`
	// 	InitialMetroID        int         `json:"initialMetroId"`
	// 	InitialMacroID        interface{} `json:"initialMacroId"`
	// 	InitialNeighborhoodID interface{} `json:"initialNeighborhoodId"`
	// 	Latitude              float64     `json:"latitude"`
	// 	Longitude             float64     `json:"longitude"`
	// 	SelectedItem          interface{} `json:"selectedItem"`
	// 	InputValue            string      `json:"inputValue"`
	// 	RecentSearchItems     interface{} `json:"recentSearchItems"`
	// 	CurrentLocationName   interface{} `json:"currentLocationName"`
	// 	CachedResults         struct {
	// 	} `json:"cachedResults"`
	// } `json:"searchAutocomplete"`
	// SignalTracking struct {
	// 	Page struct {
	// 		Name          string `json:"name"`
	// 		SeoPageType   string `json:"seoPageType"`
	// 		SeoVertical   string `json:"seoVertical"`
	// 		MarketingName string `json:"marketingName"`
	// 	} `json:"page"`
	// 	Reservation struct {
	// 		Hash     interface{} `json:"hash"`
	// 		ViewInfo interface{} `json:"viewInfo"`
	// 	} `json:"reservation"`
	// } `json:"signalTracking"`
	// SojernTracking struct {
	// 	SessionIDHash interface{} `json:"sessionIdHash"`
	// } `json:"sojernTracking"`
	// UserTracking struct {
	// 	IsUserOptedIn         bool `json:"isUserOptedIn"`
	// 	ConsentLoaded         bool `json:"consentLoaded"`
	// 	ConsentPerformance    bool `json:"consentPerformance"`
	// 	ConsentTargeting      bool `json:"consentTargeting"`
	// 	SessionIDHashesLoaded bool `json:"sessionIdHashesLoaded"`
	// } `json:"userTracking"`
	// Availability struct {
	// 	Loading                 bool `json:"loading"`
	// 	RestaurantsAvailability struct {
	// 	} `json:"restaurantsAvailability"`
	// 	RestaurantsMultiDayAvailability struct {
	// 	} `json:"restaurantsMultiDayAvailability"`
	// 	PopRestaurantsAvailability struct {
	// 	} `json:"popRestaurantsAvailability"`
	// 	AvailabilityTime      string `json:"availabilityTime"`
	// 	AvailabilityDate      string `json:"availabilityDate"`
	// 	AvailabilityPartySize int    `json:"availabilityPartySize"`
	// 	Scrollers             struct {
	// 	} `json:"scrollers"`
	// 	ScrollersLoading struct {
	// 	} `json:"scrollersLoading"`
	// 	AttributionToken interface{} `json:"attributionToken"`
	// } `json:"availability"`
	// Geolocation struct {
	// 	Lat        float64     `json:"lat"`
	// 	Long       float64     `json:"long"`
	// 	ServerLat  interface{} `json:"serverLat"`
	// 	ServerLong interface{} `json:"serverLong"`
	// 	Metro      struct {
	// 		Name     string `json:"name"`
	// 		MetroID  int    `json:"metroId"`
	// 		DomainID int    `json:"domainId"`
	// 		Timezone struct {
	// 			OffsetInMinutes int    `json:"offsetInMinutes"`
	// 			Typename        string `json:"__typename"`
	// 		} `json:"timezone"`
	// 		Typename string `json:"__typename"`
	// 	} `json:"metro"`
	// 	CountryCode string `json:"countryCode"`
	// } `json:"geolocation"`
	LolzViewAll struct {
		// CollectionToken string `json:"collectionToken"`
		SearchResults struct {
			// Title                string `json:"title"`
			// StaticMapURL         string `json:"staticMapUrl"`
			// TotalRestaurantCount int    `json:"totalRestaurantCount"`
			// RequestedDate        string `json:"requestedDate"`
			// RequestedTime        string `json:"requestedTime"`
			// CorrelationID        string `json:"correlationId"`
			Restaurants []struct {
				RestaurantID int    `json:"restaurantId"`
				Name         string `json:"name"`
				Coordinates  struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
					Typename  string  `json:"__typename"`
				} `json:"coordinates"`
				Urls struct {
					ProfileLink struct {
						Link     string `json:"link"`
						Typename string `json:"__typename"`
					} `json:"profileLink"`
					Typename string `json:"__typename"`
				} `json:"urls"`
				PriceBand struct {
					PriceBandID    int    `json:"priceBandId"`
					CurrencySymbol string `json:"currencySymbol"`
					Typename       string `json:"__typename"`
				} `json:"priceBand"`
				Photos struct {
					Profile struct {
						Xsmall struct {
							URL      string `json:"url"`
							Width    int    `json:"width"`
							Typename string `json:"__typename"`
						} `json:"xsmall"`
						Small struct {
							URL      string `json:"url"`
							Width    int    `json:"width"`
							Typename string `json:"__typename"`
						} `json:"small"`
						Legacy struct {
							URL      string `json:"url"`
							Width    int    `json:"width"`
							Typename string `json:"__typename"`
						} `json:"legacy"`
						Medium struct {
							URL      string `json:"url"`
							Width    int    `json:"width"`
							Typename string `json:"__typename"`
						} `json:"medium"`
						WideMedium struct {
							URL      string `json:"url"`
							Width    int    `json:"width"`
							Typename string `json:"__typename"`
						} `json:"wideMedium"`
						Typename string `json:"__typename"`
					} `json:"profile"`
					Typename string `json:"__typename"`
				} `json:"photos"`
				Neighborhood struct {
					Name     string `json:"name"`
					Typename string `json:"__typename"`
				} `json:"neighborhood"`
				// Statistics struct {
				// 	RecentReservationCount int `json:"recentReservationCount"`
				// 	Reviews                struct {
				// 		AllTimeTextReviewCount int `json:"allTimeTextReviewCount"`
				// 		Ratings                struct {
				// 			Overall struct {
				// 				Rating   float64 `json:"rating"`
				// 				Typename string  `json:"__typename"`
				// 			} `json:"overall"`
				// 			Typename string `json:"__typename"`
				// 		} `json:"ratings"`
				// 		Typename string `json:"__typename"`
				// 	} `json:"reviews"`
				// 	Typename string `json:"__typename"`
				// } `json:"statistics"`
				JustAddedDetails struct {
					JustAdded bool   `json:"justAdded"`
					Typename  string `json:"__typename"`
				} `json:"justAddedDetails"`
				PrimaryCuisine struct {
					Name     string `json:"name"`
					Typename string `json:"__typename"`
				} `json:"primaryCuisine"`
				Description string `json:"description"`
				TopReview   struct {
					HighlightedText string `json:"highlightedText"`
					Typename        string `json:"__typename"`
				} `json:"topReview"`
				Features struct {
					Subtype interface{} `json:"subtype"`
					// PickupEnabled interface{} `json:"pickupEnabled"`
					// Typename      string      `json:"__typename"`
				} `json:"features"`
				HasTakeout         bool `json:"hasTakeout"`
				ContactInformation struct {
					PhoneNumber          string `json:"phoneNumber"`
					FormattedPhoneNumber string `json:"formattedPhoneNumber"`
					Typename             string `json:"__typename"`
				} `json:"contactInformation"`
				// DeliveryPartners []struct {
				// 	DeliveryPartnerID   int    `json:"deliveryPartnerId"`
				// 	DeliveryPartnerName string `json:"deliveryPartnerName"`
				// 	MenuURL             string `json:"menuUrl"`
				// 	Active              bool   `json:"active"`
				// 	Typename            string `json:"__typename"`
				// } `json:"deliveryPartners"`
				// OrderOnlineLink             string `json:"orderOnlineLink"`
				RestaurantAvailabilityToken string `json:"restaurantAvailabilityToken"`
				Typename                    string `json:"__typename"`
			} `json:"restaurants"`
			Typename string `json:"__typename"`
		} `json:"searchResults"`
		// PageMetaData struct {
		// 	IsIndexed bool          `json:"isIndexed"`
		// 	Canonical string        `json:"canonical"`
		// 	HrefLangs []interface{} `json:"hrefLangs"`
		// 	Typename  string        `json:"__typename"`
		// } `json:"pageMetaData"`
		// PageNumber        int  `json:"pageNumber"`
		// Limit             int  `json:"limit"`
		// IsMobileMapOpened bool `json:"isMobileMapOpened"`
		// IsLoading         bool `json:"isLoading"`
	} `json:"lolzViewAll"`
	// MultiDayAvailabilityModal struct {
	// 	IsLoading bool `json:"isLoading"`
	// 	IsOpen    bool `json:"isOpen"`
	// } `json:"multiDayAvailabilityModal"`
}

type RawSearchInfo struct {
	// AuthModal struct {
	// 	IsAuthModalOpen                 bool        `json:"isAuthModalOpen"`
	// 	HideCloseButton                 bool        `json:"hideCloseButton"`
	// 	ShouldCloseAuthModalForAuthUser bool        `json:"shouldCloseAuthModalForAuthUser"`
	// 	AuthModalType                   interface{} `json:"authModalType"`
	// 	AuthIframeSrc                   interface{} `json:"authIframeSrc"`
	// 	IsFromBookingFlow               bool        `json:"isFromBookingFlow"`
	// } `json:"authModal"`
	// Authentication struct {
	// 	IsAuthenticated   bool `json:"isAuthenticated"`
	// 	UserType          int  `json:"userType"`
	// 	AuthFinished      bool `json:"authFinished"`
	// 	UseClientSideAuth bool `json:"useClientSideAuth"`
	// } `json:"authentication"`
	// DagAnalytics struct {
	// 	PageName      string      `json:"pageName"`
	// 	CorrelationID interface{} `json:"correlationId"`
	// 	OriginID      string      `json:"originId"`
	// 	DeviceType    string      `json:"deviceType"`
	// 	RequestTypes  struct {
	// 		MultiSearch struct {
	// 			CorrelationID string `json:"correlationId"`
	// 		} `json:"multiSearch"`
	// 	} `json:"requestTypes"`
	// 	OriginalRequestID string `json:"originalRequestId"`
	// 	UserAgent         string `json:"userAgent"`
	// } `json:"dagAnalytics"`
	// Domain struct {
	// 	DomainID            int    `json:"domainId"`
	// 	Tld                 string `json:"tld"`
	// 	OtEnv               string `json:"otEnv"`
	// 	DeploymentRegionKey string `json:"deploymentRegionKey"`
	// } `json:"domain"`
	// Experiments struct {
	// 	BUX1978                          int `json:"BUX-1978"`
	// 	Sfx5536                          int `json:"sfx-5536"`
	// 	WEBX255                          int `json:"WEBX-255"`
	// 	WEBX312                          int `json:"WEBX-312"`
	// 	ShowReservationUnavailableBanner int `json:"showReservationUnavailableBanner"`
	// } `json:"experiments"`
	// FacebookTracking struct {
	// 	SessionIDHash interface{} `json:"sessionIdHash"`
	// } `json:"facebookTracking"`
	// FeatureToggles struct {
	// 	ShowPrideLogo         bool `json:"showPrideLogo"`
	// 	IsUploadPhotosEnabled bool `json:"isUploadPhotosEnabled"`
	// } `json:"featureToggles"`
	// Footer struct {
	// 	FormIsSubmitted bool `json:"formIsSubmitted"`
	// 	RequestDidError bool `json:"requestDidError"`
	// 	DiscoverLinks   []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"discoverLinks"`
	// 	OpentableLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"opentableLinks"`
	// 	MoreLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"moreLinks"`
	// 	SiteLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"siteLinks"`
	// 	SocialLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"socialLinks"`
	// 	RestaurateursLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"restaurateursLinks"`
	// 	LegalLinks []struct {
	// 		URL         string `json:"url"`
	// 		ID          string `json:"id"`
	// 		IsCrawlable bool   `json:"isCrawlable"`
	// 		Typename    string `json:"__typename"`
	// 	} `json:"legalLinks"`
	// } `json:"footer"`
	// Header struct {
	// 	Crumbs []struct {
	// 		Name string `json:"name"`
	// 		Href string `json:"href,omitempty"`
	// 	} `json:"crumbs"`
	// 	MetroDomainID                   int           `json:"metroDomainId"`
	// 	FeaturedMetros                  []interface{} `json:"featuredMetros"`
	// 	MetroID                         int           `json:"metroId"`
	// 	MacroID                         interface{}   `json:"macroId"`
	// 	NeighborhoodID                  interface{}   `json:"neighborhoodId"`
	// 	CurrentMetro                    interface{}   `json:"currentMetro"`
	// 	IsLocationPickerLoading         bool          `json:"isLocationPickerLoading"`
	// 	HideLocationPicker              bool          `json:"hideLocationPicker"`
	// 	TimezoneOffset                  int           `json:"timezoneOffset"`
	// 	ShouldDisplayOverlay            bool          `json:"shouldDisplayOverlay"`
	// 	UseHeaderAuth                   bool          `json:"useHeaderAuth"`
	// 	IsUserProfileLoading            bool          `json:"isUserProfileLoading"`
	// 	UserProfile                     interface{}   `json:"userProfile"`
	// 	UserTransactions                []interface{} `json:"userTransactions"`
	// 	UserInvites                     []interface{} `json:"userInvites"`
	// 	UserNotifications               []interface{} `json:"userNotifications"`
	// 	HasUnreadNotifications          bool          `json:"hasUnreadNotifications"`
	// 	UserNotificationsInitialLoading bool          `json:"userNotificationsInitialLoading"`
	// 	UserNewsfeedStory               interface{}   `json:"userNewsfeedStory"`
	// 	HasUnreadNewsfeedStory          bool          `json:"hasUnreadNewsfeedStory"`
	// 	UserNewsfeedInitialLoading      bool          `json:"userNewsfeedInitialLoading"`
	// 	ShowCovid19Banner               bool          `json:"showCovid19Banner"`
	// 	CovidBannerCTALink              string        `json:"covidBannerCTALink"`
	// 	HideLabel                       bool          `json:"hideLabel"`
	// } `json:"header"`
	// Language struct {
	// 	Locale          string `json:"locale"`
	// 	AcceptLanguage  string `json:"acceptLanguage"`
	// 	DomainLanguages struct {
	// 		En string `json:"en"`
	// 		Fr string `json:"fr"`
	// 		Es string `json:"es"`
	// 		De string `json:"de"`
	// 		Ja string `json:"ja"`
	// 		Nl string `json:"nl"`
	// 		It string `json:"it"`
	// 	} `json:"domainLanguages"`
	// 	CurrencySymbol string `json:"currencySymbol"`
	// } `json:"language"`
	// SearchAutocomplete struct {
	// 	AutocompleteResults   interface{} `json:"autocompleteResults"`
	// 	CorrelationID         interface{} `json:"correlationId"`
	// 	InitialMetroID        int         `json:"initialMetroId"`
	// 	InitialMacroID        interface{} `json:"initialMacroId"`
	// 	InitialNeighborhoodID interface{} `json:"initialNeighborhoodId"`
	// 	Latitude              float64     `json:"latitude"`
	// 	Longitude             float64     `json:"longitude"`
	// 	SelectedItem          interface{} `json:"selectedItem"`
	// 	InputValue            string      `json:"inputValue"`
	// 	RecentSearchItems     interface{} `json:"recentSearchItems"`
	// 	CurrentLocationName   string      `json:"currentLocationName"`
	// 	CachedResults         struct {
	// 	} `json:"cachedResults"`
	// } `json:"searchAutocomplete"`
	// SignalTracking struct {
	// 	Page struct {
	// 		Name          string `json:"name"`
	// 		SeoPageType   string `json:"seoPageType"`
	// 		SeoVertical   string `json:"seoVertical"`
	// 		MarketingName string `json:"marketingName"`
	// 	} `json:"page"`
	// 	Reservation struct {
	// 		Hash     interface{} `json:"hash"`
	// 		ViewInfo interface{} `json:"viewInfo"`
	// 	} `json:"reservation"`
	// } `json:"signalTracking"`
	// SojernTracking struct {
	// 	SessionIDHash interface{} `json:"sessionIdHash"`
	// } `json:"sojernTracking"`
	// UserTracking struct {
	// 	IsUserOptedIn         bool `json:"isUserOptedIn"`
	// 	ConsentLoaded         bool `json:"consentLoaded"`
	// 	ConsentPerformance    bool `json:"consentPerformance"`
	// 	ConsentTargeting      bool `json:"consentTargeting"`
	// 	SessionIDHashesLoaded bool `json:"sessionIdHashesLoaded"`
	// } `json:"userTracking"`
	// Availability struct {
	// 	Loading                 bool `json:"loading"`
	// 	RestaurantsAvailability struct {
	// 	} `json:"restaurantsAvailability"`
	// 	RestaurantsMultiDayAvailability struct {
	// 	} `json:"restaurantsMultiDayAvailability"`
	// 	PopRestaurantsAvailability struct {
	// 	} `json:"popRestaurantsAvailability"`
	// 	AvailabilityTime      string `json:"availabilityTime"`
	// 	AvailabilityDate      string `json:"availabilityDate"`
	// 	AvailabilityPartySize int    `json:"availabilityPartySize"`
	// 	Scrollers             struct {
	// 	} `json:"scrollers"`
	// 	ScrollersLoading struct {
	// 	} `json:"scrollersLoading"`
	// 	AttributionToken string `json:"attributionToken"`
	// } `json:"availability"`
	// Geolocation struct {
	// 	Lat         float64     `json:"lat"`
	// 	Long        float64     `json:"long"`
	// 	ServerLat   interface{} `json:"serverLat"`
	// 	ServerLong  interface{} `json:"serverLong"`
	// 	Metro       interface{} `json:"metro"`
	// 	CountryCode string      `json:"countryCode"`
	// } `json:"geolocation"`
	// Map struct {
	// 	IsFullScreenMode          bool   `json:"isFullScreenMode"`
	// 	IsMapLoaded               bool   `json:"isMapLoaded"`
	// 	IsSearchOnMoveActive      bool   `json:"isSearchOnMoveActive"`
	// 	IsDistanceSelectOpen      bool   `json:"isDistanceSelectOpen"`
	// 	Zoom                      int    `json:"zoom"`
	// 	HasUserInteractedWithMap  bool   `json:"hasUserInteractedWithMap"`
	// 	IsUserLocationLoading     bool   `json:"isUserLocationLoading"`
	// 	GtmPrefix                 string `json:"gtmPrefix"`
	// 	ExpandedRestaurantDetails struct {
	// 		IsLoading  bool        `json:"isLoading"`
	// 		Restaurant interface{} `json:"restaurant"`
	// 	} `json:"expandedRestaurantDetails"`
	// } `json:"map"`
	// MultiDayAvailabilityModal struct {
	// 	IsLoading bool `json:"isLoading"`
	// 	IsOpen    bool `json:"isOpen"`
	// } `json:"multiDayAvailabilityModal"`
	MultiSearch struct {
		// PageType  string  `json:"pageType"`
		// MetroID   int     `json:"metroId"`
		// Latitude  float64 `json:"latitude"`
		// Longitude float64 `json:"longitude"`
		// Metro     struct {
		// 	MetroID        int    `json:"metroId"`
		// 	Name           string `json:"name"`
		// 	DomainID       int    `json:"domainId"`
		// 	GeographicData struct {
		// 		Center struct {
		// 			Latitude  float64 `json:"latitude"`
		// 			Longitude float64 `json:"longitude"`
		// 			Typename  string  `json:"__typename"`
		// 		} `json:"center"`
		// 		Typename string `json:"__typename"`
		// 	} `json:"geographicData"`
		// 	Timezone struct {
		// 		OffsetInMinutes int    `json:"offsetInMinutes"`
		// 		Typename        string `json:"__typename"`
		// 	} `json:"timezone"`
		// 	Urls struct {
		// 		StartPageLink struct {
		// 			Link     string `json:"link"`
		// 			Typename string `json:"__typename"`
		// 		} `json:"startPageLink"`
		// 		ListingsPageLink struct {
		// 			Link     string `json:"link"`
		// 			Typename string `json:"__typename"`
		// 		} `json:"listingsPageLink"`
		// 		Typename string `json:"__typename"`
		// 	} `json:"urls"`
		// 	Typename string `json:"__typename"`
		// } `json:"metro"`
		// IsNotFoundError bool `json:"isNotFoundError"`
		// IsServerError   bool `json:"isServerError"`
		// IsLoading       bool `json:"isLoading"`
		// IsFacetsLoading bool `json:"isFacetsLoading"`
		// FacetsLoaded    bool `json:"facetsLoaded"`
		// Facets          struct {
		// 	Prices []int `json:"prices"`
		// 	Times  []struct {
		// 		Key string `json:"key"`
		// 	} `json:"times"`
		// 	DiningTypes []struct {
		// 		Key string `json:"key"`
		// 	} `json:"diningTypes"`
		// 	LegacyAreas []struct {
		// 		ID       int    `json:"id"`
		// 		Name     string `json:"name"`
		// 		Count    int    `json:"count"`
		// 		Type     string `json:"type"`
		// 		Typename string `json:"__typename"`
		// 	} `json:"legacyAreas"`
		// 	Cuisines []struct {
		// 		CuisineID string `json:"cuisineId"`
		// 		Name      string `json:"name"`
		// 		Count     int    `json:"count"`
		// 		Typename  string `json:"__typename"`
		// 	} `json:"cuisines"`
		// 	Tags []struct {
		// 		TagID    string `json:"tagId"`
		// 		Name     string `json:"name"`
		// 		Count    int    `json:"count"`
		// 		Typename string `json:"__typename"`
		// 	} `json:"tags"`
		// 	TableCategories []struct {
		// 		TableCategoryID string `json:"tableCategoryId"`
		// 		Count           int    `json:"count"`
		// 		Typename        string `json:"__typename"`
		// 	} `json:"tableCategories"`
		// 	ExperienceTypes []struct {
		// 		TypeID   int    `json:"typeId"`
		// 		Name     string `json:"name"`
		// 		Typename string `json:"__typename"`
		// 	} `json:"experienceTypes"`
		// 	ProofOfVaccinationRequired bool `json:"proofOfVaccinationRequired"`
		// 	AdditionalDetails          []struct {
		// 		AdditionalDetailID string `json:"additionalDetailId"`
		// 		Name               string `json:"name"`
		// 		Count              int    `json:"count"`
		// 		Typename           string `json:"__typename"`
		// 	} `json:"additionalDetails"`
		// 	DniTags []struct {
		// 		DniTagID string `json:"dniTagId"`
		// 		Count    int    `json:"count"`
		// 		Typename string `json:"__typename"`
		// 	} `json:"dniTags"`
		// 	Typename string `json:"__typename"`
		// } `json:"facets"`
		// Filters struct {
		// 	AdditionalDetailIds        []interface{} `json:"additionalDetailIds"`
		// 	CuisineIds                 []interface{} `json:"cuisineIds"`
		// 	DiningType                 string        `json:"diningType"`
		// 	DniTagIds                  []interface{} `json:"dniTagIds"`
		// 	ExperienceTypeIds          []interface{} `json:"experienceTypeIds"`
		// 	LoyaltyRedemptionTiers     []interface{} `json:"loyaltyRedemptionTiers"`
		// 	NeighborhoodIds            []interface{} `json:"neighborhoodIds"`
		// 	OnlyPopTimes               bool          `json:"onlyPopTimes"`
		// 	PriceBandIds               []interface{} `json:"priceBandIds"`
		// 	ProofOfVaccinationRequired bool          `json:"proofOfVaccinationRequired"`
		// 	RegionIds                  []interface{} `json:"regionIds"`
		// 	TableCategoryIds           []interface{} `json:"tableCategoryIds"`
		// 	TagIds                     []interface{} `json:"tagIds"`
		// } `json:"filters"`
		// FiltersHistory []interface{} `json:"filtersHistory"`
		Restaurants []struct {
			RestaurantID  int    `json:"restaurantId"`
			Name          string `json:"name"`
			Type          string `json:"type"`
			SlotDiscovery string `json:"slotDiscovery"`
			Urls          struct {
				ProfileLink struct {
					Link     string `json:"link"`
					Typename string `json:"__typename"`
				} `json:"profileLink"`
				Typename string `json:"__typename"`
			} `json:"urls"`
			PriceBand struct {
				PriceBandID    int    `json:"priceBandId"`
				CurrencySymbol string `json:"currencySymbol"`
				Name           string `json:"name"`
				Typename       string `json:"__typename"`
			} `json:"priceBand"`
			Neighborhood struct {
				Name     string `json:"name"`
				Typename string `json:"__typename"`
			} `json:"neighborhood"`
			// Statistics struct {
			// 	RecentReservationCount int `json:"recentReservationCount"`
			// 	Reviews                struct {
			// 		AllTimeTextReviewCount int `json:"allTimeTextReviewCount"`
			// 		Ratings                struct {
			// 			Overall struct {
			// 				Rating   float64 `json:"rating"`
			// 				Typename string  `json:"__typename"`
			// 			} `json:"overall"`
			// 			Typename string `json:"__typename"`
			// 		} `json:"ratings"`
			// 		Typename string `json:"__typename"`
			// 	} `json:"reviews"`
			// 	Typename string `json:"__typename"`
			// } `json:"statistics"`
			PrimaryCuisine struct {
				Name     string `json:"name"`
				Typename string `json:"__typename"`
			} `json:"primaryCuisine"`
			CampaignID int  `json:"campaignId"`
			IsPromoted bool `json:"isPromoted"`
			Features   struct {
				Subtype       interface{} `json:"subtype"`
				Bar           bool        `json:"bar"`
				Counter       bool        `json:"counter"`
				HighTop       bool        `json:"highTop"`
				Outdoor       bool        `json:"outdoor"`
				PickupEnabled interface{} `json:"pickupEnabled"`
				Typename      string      `json:"__typename"`
			} `json:"features"`
			RestaurantAvailabilityToken string `json:"restaurantAvailabilityToken"`
			// Groups                      []struct {
			// 	ID                        int         `json:"id"`
			// 	Rids                      interface{} `json:"rids"`
			// 	OffDiscoveryRestaurantIds interface{} `json:"offDiscoveryRestaurantIds"`
			// 	Typename                  string      `json:"__typename"`
			// } `json:"groups"`
			Typename string `json:"__typename"`
			IsPinned bool   `json:"isPinned"`
			Photos   struct {
				Profile struct {
					Xsmall struct {
						URL      string `json:"url"`
						Width    int    `json:"width"`
						Typename string `json:"__typename"`
					} `json:"xsmall"`
					Small struct {
						URL      string `json:"url"`
						Width    int    `json:"width"`
						Typename string `json:"__typename"`
					} `json:"small"`
					Legacy struct {
						URL      string `json:"url"`
						Width    int    `json:"width"`
						Typename string `json:"__typename"`
					} `json:"legacy"`
					Medium struct {
						URL      string `json:"url"`
						Width    int    `json:"width"`
						Typename string `json:"__typename"`
					} `json:"medium"`
					Typename string `json:"__typename"`
				} `json:"profile"`
				Typename string `json:"__typename"`
			} `json:"photos"`
			JustAddedDetails struct {
				JustAdded bool   `json:"justAdded"`
				Typename  string `json:"__typename"`
			} `json:"justAddedDetails"`
			Offers      []interface{} `json:"offers"`
			DiningStyle string        `json:"diningStyle"`
			Coordinates struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				Typename  string  `json:"__typename"`
			} `json:"coordinates"`
			Address struct {
				Line1    string `json:"line1"`
				Line2    string `json:"line2"`
				City     string `json:"city"`
				State    string `json:"state"`
				PostCode string `json:"postCode"`
				Typename string `json:"__typename"`
			} `json:"address"`
			Description string `json:"description"`
			TopReview   struct {
				HighlightedText string `json:"highlightedText"`
				Typename        string `json:"__typename"`
			} `json:"topReview"`
			MatchRelevance struct {
				Text     string `json:"text"`
				Source   string `json:"source"`
				Typename string `json:"__typename"`
			} `json:"matchRelevance"`
			HasTakeout         bool `json:"hasTakeout"`
			ContactInformation struct {
				PhoneNumber          string `json:"phoneNumber"`
				FormattedPhoneNumber string `json:"formattedPhoneNumber"`
				Typename             string `json:"__typename"`
			} `json:"contactInformation"`
			DeliveryPartners []interface{} `json:"deliveryPartners"`
			OrderOnlineLink  interface{}   `json:"orderOnlineLink"`
		} `json:"restaurants"`
		// TotalRestaurantCount      int           `json:"totalRestaurantCount"`
		// PopRestaurants            []interface{} `json:"popRestaurants"`
		// TotalPopRestaurantCount   int           `json:"totalPopRestaurantCount"`
		// PageNumber                int           `json:"pageNumber"`
		// Limit                     int           `json:"limit"`
		// PopLimit                  int           `json:"popLimit"`
		// MapLimit                  int           `json:"mapLimit"`
		// PinnedRid                 interface{}   `json:"pinnedRid"`
		// FreetextTerm              string        `json:"freetextTerm"`
		// OriginalTerm              string        `json:"originalTerm"`
		// IntentModifiedTerm        string        `json:"intentModifiedTerm"`
		// SortBy                    string        `json:"sortBy"`
		// QueryUnderstandingType    string        `json:"queryUnderstandingType"`
		// DateTime                  string        `json:"dateTime"`
		// PartySize                 int           `json:"partySize"`
		// IsValidGiftPage           bool          `json:"isValidGiftPage"`
		// DisableMapView            bool          `json:"disableMapView"`
		// IsMobileSearch            bool          `json:"isMobileSearch"`
		// IsMapView                 bool          `json:"isMapView"`
		// ShowCollapsibleDtp        bool          `json:"showCollapsibleDtp"`
		// ShowCollapsibleDtpHeading bool          `json:"showCollapsibleDtpHeading"`
		// OnlyWithOffers            bool          `json:"onlyWithOffers"`
		// Debug                     bool          `json:"debug"`
		// ShouldUseLatLongSearch    bool          `json:"shouldUseLatLongSearch"`
		// AnytimeAvailability       bool          `json:"anytimeAvailability"`
	} `json:"multiSearch"`
}

type RestaurantPriceBand struct {
	PriceBandID    int
	CurrencySymbol string
	Name           string
}

type RestaurantFeatures struct {
	Subtype       interface{}
	Bar           bool
	Counter       bool
	HighTop       bool
	Outdoor       bool
	PickupEnabled bool
	Typename      string
}

type RestaurantPhoto struct {
	URL   string
	Width int
}

type RestaurantCoordinates struct {
	Latitude  float64
	Longitude float64
}

type RestaurantPhotos struct {
	Xsmall RestaurantPhoto
	Small  RestaurantPhoto
	Legacy RestaurantPhoto
	Medium RestaurantPhoto
}

type RestaurantAddress struct {
	Line1    string
	Line2    string
	City     string
	State    string
	PostCode string
	Typename string
}

type RestaurantContactInformation struct {
	PhoneNumber          string
	FormattedPhoneNumber string
}

type Restaurant struct {
	RestaurantID  int
	Name          string
	Type          string
	SlotDiscovery string
	ProfileLink   string
	PriceBand     RestaurantPriceBand
	Neighborhood  string
	// Statistics struct {
	// 	RecentReservationCount int
	// 	Reviews                struct {
	// 		AllTimeTextReviewCount int
	// 		Ratings                struct {
	// 			Overall struct {
	// 				Rating   float64
	// 				Typename string
	// 			}
	// 			Typename string
	// 		}
	// 		Typename string
	// 	}
	// 	Typename string
	// }
	PrimaryCuisine              string
	CampaignID                  int
	IsPromoted                  bool
	Features                    RestaurantFeatures
	RestaurantAvailabilityToken string
	// Groups                      []struct {
	// 	ID                        int
	// 	Rids                      interface{}
	// 	OffDiscoveryRestaurantIds interface{}
	// 	Typename                  string
	// }
	IsPinned  bool
	Photos    RestaurantPhotos
	JustAdded bool
	// Offers      []interface{}
	DiningStyle string
	Coordinates RestaurantCoordinates
	Address     RestaurantAddress
	Description string
	TopReview   string
	// MatchRelevance struct {
	// 	Text     string
	// 	Source   string
	// 	Typename string
	// }
	HasTakeout         bool
	ContactInformation RestaurantContactInformation
	// DeliveryPartners []interface{}
	// OrderOnlineLink  interface{}
}

type SearchInfo struct {
	Restaurants []Restaurant
}

func (s RawSearchInfo) Convert() *SearchInfo {
	var ress []Restaurant
	for _, r := range s.MultiSearch.Restaurants {
		ress = append(ress, Restaurant{
			RestaurantID:                r.RestaurantID,
			Name:                        r.Name,
			Type:                        r.Type,
			SlotDiscovery:               r.SlotDiscovery,
			ProfileLink:                 r.Urls.ProfileLink.Link,
			Neighborhood:                r.Neighborhood.Name,
			PrimaryCuisine:              r.PrimaryCuisine.Name,
			CampaignID:                  r.CampaignID,
			IsPromoted:                  r.IsPromoted,
			RestaurantAvailabilityToken: r.RestaurantAvailabilityToken,
			IsPinned:                    r.IsPinned,
			DiningStyle:                 r.DiningStyle,
			Description:                 r.Description,
			PriceBand: RestaurantPriceBand{
				PriceBandID:    r.PriceBand.PriceBandID,
				CurrencySymbol: r.PriceBand.CurrencySymbol,
				Name:           r.PriceBand.Name,
			},
			HasTakeout: r.HasTakeout,
			TopReview:  r.TopReview.HighlightedText,
			ContactInformation: RestaurantContactInformation{
				PhoneNumber:          r.ContactInformation.PhoneNumber,
				FormattedPhoneNumber: r.ContactInformation.FormattedPhoneNumber,
			},
			Features: RestaurantFeatures{
				Subtype: r.Features.Subtype,
				Bar:     r.Features.Bar,
				Counter: r.Features.Counter,
				HighTop: r.Features.HighTop,
				Outdoor: r.Features.Outdoor,
				// PickupEnabled: r.Features.PickupEnabled,
			},
			JustAdded: r.JustAddedDetails.JustAdded,
			Coordinates: RestaurantCoordinates{
				Latitude:  r.Coordinates.Latitude,
				Longitude: r.Coordinates.Longitude,
			},
			Address: RestaurantAddress{
				Line1:    r.Address.Line1,
				Line2:    r.Address.Line2,
				City:     r.Address.City,
				State:    r.Address.State,
				PostCode: r.Address.PostCode,
			},
			Photos: RestaurantPhotos{
				Xsmall: RestaurantPhoto{
					URL:   r.Photos.Profile.Xsmall.URL,
					Width: r.Photos.Profile.Xsmall.Width,
				},
				Small: RestaurantPhoto{
					URL:   r.Photos.Profile.Small.URL,
					Width: r.Photos.Profile.Small.Width,
				},
				Legacy: RestaurantPhoto{
					URL:   r.Photos.Profile.Legacy.URL,
					Width: r.Photos.Profile.Legacy.Width,
				},
				Medium: RestaurantPhoto{
					URL:   r.Photos.Profile.Medium.URL,
					Width: r.Photos.Profile.Medium.Width,
				},
			},
		})
	}
	return &SearchInfo{
		Restaurants: ress,
	}
}

func (s RawSearchByURIInfo) Convert() *SearchInfo {
	var ress []Restaurant
	for _, r := range s.LolzViewAll.SearchResults.Restaurants {
		ress = append(ress, Restaurant{
			RestaurantID: r.RestaurantID,
			Name:         r.Name,
			// Type:                        r.Type,
			// SlotDiscovery:               r.SlotDiscovery,
			ProfileLink:    r.Urls.ProfileLink.Link,
			Neighborhood:   r.Neighborhood.Name,
			PrimaryCuisine: r.PrimaryCuisine.Name,
			// CampaignID:                  r.CampaignID,
			// IsPromoted:                  r.IsPromoted,
			RestaurantAvailabilityToken: r.RestaurantAvailabilityToken,
			// IsPinned:                    r.IsPinned,
			// DiningStyle:                 r.DiningStyle,
			Description: r.Description,
			PriceBand: RestaurantPriceBand{
				PriceBandID:    r.PriceBand.PriceBandID,
				CurrencySymbol: r.PriceBand.CurrencySymbol,
				// Name:           r.PriceBand.Name,
			},
			HasTakeout: r.HasTakeout,
			TopReview:  r.TopReview.HighlightedText,
			ContactInformation: RestaurantContactInformation{
				PhoneNumber:          r.ContactInformation.PhoneNumber,
				FormattedPhoneNumber: r.ContactInformation.FormattedPhoneNumber,
			},
			Features: RestaurantFeatures{
				Subtype: r.Features.Subtype,
				// Bar:     r.Features.Bar,
				// Counter: r.Features.Counter,
				// HighTop: r.Features.HighTop,
				// Outdoor: r.Features.Outdoor,
				// PickupEnabled: r.Features.PickupEnabled,
			},
			JustAdded: r.JustAddedDetails.JustAdded,
			Coordinates: RestaurantCoordinates{
				Latitude:  r.Coordinates.Latitude,
				Longitude: r.Coordinates.Longitude,
			},
			// Address: RestaurantAddress{
			// 	Line1:    r.Address.Line1,
			// 	Line2:    r.Address.Line2,
			// 	City:     r.Address.City,
			// 	State:    r.Address.State,
			// 	PostCode: r.Address.PostCode,
			// },
			Photos: RestaurantPhotos{
				Xsmall: RestaurantPhoto{
					URL:   r.Photos.Profile.Xsmall.URL,
					Width: r.Photos.Profile.Xsmall.Width,
				},
				Small: RestaurantPhoto{
					URL:   r.Photos.Profile.Small.URL,
					Width: r.Photos.Profile.Small.Width,
				},
				Legacy: RestaurantPhoto{
					URL:   r.Photos.Profile.Legacy.URL,
					Width: r.Photos.Profile.Legacy.Width,
				},
				Medium: RestaurantPhoto{
					URL:   r.Photos.Profile.Medium.URL,
					Width: r.Photos.Profile.Medium.Width,
				},
			},
		})
	}
	return &SearchInfo{
		Restaurants: ress,
	}
}

func (s RawListByURIInfo) Convert() *SearchInfo {
	var ress []Restaurant
	for _, r := range s.MultiSearch.PopRestaurants {
		ress = append(ress, Restaurant{
			RestaurantID:   r.RestaurantID,
			Name:           r.Name,
			Type:           r.Type,
			SlotDiscovery:  r.SlotDiscovery,
			ProfileLink:    r.Urls.ProfileLink.Link,
			Neighborhood:   r.Neighborhood.Name,
			PrimaryCuisine: r.PrimaryCuisine.Name,
			// CampaignID:                  r.CampaignID,
			IsPromoted:                  r.IsPromoted,
			RestaurantAvailabilityToken: r.RestaurantAvailabilityToken,
			// IsPinned:                    r.IsPinned,
			// DiningStyle:                 r.DiningStyle,
			// Description:                 r.Description,
			PriceBand: RestaurantPriceBand{
				PriceBandID:    r.PriceBand.PriceBandID,
				CurrencySymbol: r.PriceBand.CurrencySymbol,
				Name:           r.PriceBand.Name,
			},
			// HasTakeout: r.HasTakeout,
			// TopReview:  r.TopReview.HighlightedText,
			// ContactInformation: RestaurantContactInformation{
			// 	PhoneNumber:          r.ContactInformation.PhoneNumber,
			// 	FormattedPhoneNumber: r.ContactInformation.FormattedPhoneNumber,
			// },
			Features: RestaurantFeatures{
				Subtype: r.Features.Subtype,
				Bar:     r.Features.Bar,
				Counter: r.Features.Counter,
				HighTop: r.Features.HighTop,
				Outdoor: r.Features.Outdoor,
				// PickupEnabled: r.Features.PickupEnabled,
			},
			// JustAdded: r.JustAddedDetails.JustAdded,
			// Coordinates: RestaurantCoordinates{
			// 	Latitude:  r.Coordinates.Latitude,
			// 	Longitude: r.Coordinates.Longitude,
			// },
			// Address: RestaurantAddress{
			// 	Line1:    r.Address.Line1,
			// 	Line2:    r.Address.Line2,
			// 	City:     r.Address.City,
			// 	State:    r.Address.State,
			// 	PostCode: r.Address.PostCode,
			// },
			// Photos: RestaurantPhotos{
			// 	Xsmall: RestaurantPhoto{
			// 		URL:   r.Photos.Profile.Xsmall.URL,
			// 		Width: r.Photos.Profile.Xsmall.Width,
			// 	},
			// 	Small: RestaurantPhoto{
			// 		URL:   r.Photos.Profile.Small.URL,
			// 		Width: r.Photos.Profile.Small.Width,
			// 	},
			// 	Legacy: RestaurantPhoto{
			// 		URL:   r.Photos.Profile.Legacy.URL,
			// 		Width: r.Photos.Profile.Legacy.Width,
			// 	},
			// 	Medium: RestaurantPhoto{
			// 		URL:   r.Photos.Profile.Medium.URL,
			// 		Width: r.Photos.Profile.Medium.Width,
			// 	},
			// },
		})
	}
	for _, r := range s.MultiSearch.Restaurants {
		ress = append(ress, Restaurant{
			RestaurantID:                r.RestaurantID,
			Name:                        r.Name,
			Type:                        r.Type,
			SlotDiscovery:               r.SlotDiscovery,
			ProfileLink:                 r.Urls.ProfileLink.Link,
			Neighborhood:                r.Neighborhood.Name,
			PrimaryCuisine:              r.PrimaryCuisine.Name,
			CampaignID:                  r.CampaignID,
			IsPromoted:                  r.IsPromoted,
			RestaurantAvailabilityToken: r.RestaurantAvailabilityToken,
			IsPinned:                    r.IsPinned,
			DiningStyle:                 r.DiningStyle,
			Description:                 r.Description,
			PriceBand: RestaurantPriceBand{
				PriceBandID:    r.PriceBand.PriceBandID,
				CurrencySymbol: r.PriceBand.CurrencySymbol,
				Name:           r.PriceBand.Name,
			},
			HasTakeout: r.HasTakeout,
			TopReview:  r.TopReview.HighlightedText,
			ContactInformation: RestaurantContactInformation{
				PhoneNumber:          r.ContactInformation.PhoneNumber,
				FormattedPhoneNumber: r.ContactInformation.FormattedPhoneNumber,
			},
			Features: RestaurantFeatures{
				Subtype: r.Features.Subtype,
				Bar:     r.Features.Bar,
				Counter: r.Features.Counter,
				HighTop: r.Features.HighTop,
				Outdoor: r.Features.Outdoor,
				// PickupEnabled: r.Features.PickupEnabled,
			},
			JustAdded: r.JustAddedDetails.JustAdded,
			Coordinates: RestaurantCoordinates{
				Latitude:  r.Coordinates.Latitude,
				Longitude: r.Coordinates.Longitude,
			},
			Address: RestaurantAddress{
				Line1:    r.Address.Line1,
				Line2:    r.Address.Line2,
				City:     r.Address.City,
				State:    r.Address.State,
				PostCode: r.Address.PostCode,
			},
			Photos: RestaurantPhotos{
				Xsmall: RestaurantPhoto{
					URL:   r.Photos.Profile.Xsmall.URL,
					Width: r.Photos.Profile.Xsmall.Width,
				},
				Small: RestaurantPhoto{
					URL:   r.Photos.Profile.Small.URL,
					Width: r.Photos.Profile.Small.Width,
				},
				Legacy: RestaurantPhoto{
					URL:   r.Photos.Profile.Legacy.URL,
					Width: r.Photos.Profile.Legacy.Width,
				},
				Medium: RestaurantPhoto{
					URL:   r.Photos.Profile.Medium.URL,
					Width: r.Photos.Profile.Medium.Width,
				},
			},
		})
	}
	return &SearchInfo{
		Restaurants: ress,
	}
}
