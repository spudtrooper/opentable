package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type RawRestaurantDetails struct {
	AuthModal struct {
		IsAuthModalOpen                 bool        `json:"isAuthModalOpen"`
		HideCloseButton                 bool        `json:"hideCloseButton"`
		ShouldCloseAuthModalForAuthUser bool        `json:"shouldCloseAuthModalForAuthUser"`
		AuthModalType                   interface{} `json:"authModalType"`
		AuthIframeSrc                   interface{} `json:"authIframeSrc"`
		IsFromBookingFlow               bool        `json:"isFromBookingFlow"`
	} `json:"authModal"`
	Authentication struct {
		IsAuthenticated   bool   `json:"isAuthenticated"`
		UserType          int    `json:"userType"`
		UserCountryID     string `json:"userCountryId"`
		AuthFinished      bool   `json:"authFinished"`
		UseClientSideAuth bool   `json:"useClientSideAuth"`
	} `json:"authentication"`
	DagAnalytics struct {
		PageName      string      `json:"pageName"`
		CorrelationID interface{} `json:"correlationId"`
		OriginID      interface{} `json:"originId"`
		RequestTypes  struct {
			RestaurantProfile struct {
				CorrelationID string `json:"correlationId"`
			} `json:"restaurantProfile"`
			FeaturedIn struct {
				CorrelationID string `json:"correlationId"`
			} `json:"featured_in"`
		} `json:"requestTypes"`
		OriginalRequestID string `json:"originalRequestId"`
		UserAgent         string `json:"userAgent"`
	} `json:"dagAnalytics"`
	Domain struct {
		DomainID            int    `json:"domainId"`
		Tld                 string `json:"tld"`
		OtEnv               string `json:"otEnv"`
		DeploymentRegionKey string `json:"deploymentRegionKey"`
	} `json:"domain"`
	// Experiments struct {
	// 	BUX1978                          int `json:"BUX-1978"`
	// 	GRO1164                          int `json:"GRO-1164"`
	// 	CT3337                           int `json:"CT3-337"`
	// 	ShowReservationUnavailableBanner int `json:"showReservationUnavailableBanner"`
	// } `json:"experiments"`
	FacebookTracking struct {
		SessionIDHash interface{} `json:"sessionIdHash"`
		Fbp           interface{} `json:"fbp"`
		Fbc           interface{} `json:"fbc"`
	} `json:"facebookTracking"`
	FeatureToggles struct {
		ShowPrideLogo         bool `json:"showPrideLogo"`
		IsUploadPhotosEnabled bool `json:"isUploadPhotosEnabled"`
	} `json:"featureToggles"`
	Footer struct {
		FormIsSubmitted bool `json:"formIsSubmitted"`
		RequestDidError bool `json:"requestDidError"`
		DiscoverLinks   []struct {
			URL         string `json:"url"`
			ID          string `json:"id"`
			IsCrawlable bool   `json:"isCrawlable"`
			Typename    string `json:"__typename"`
		} `json:"discoverLinks"`
		OpentableLinks []struct {
			URL         string `json:"url"`
			ID          string `json:"id"`
			IsCrawlable bool   `json:"isCrawlable"`
			Typename    string `json:"__typename"`
		} `json:"opentableLinks"`
		MoreLinks []struct {
			URL         string `json:"url"`
			ID          string `json:"id"`
			IsCrawlable bool   `json:"isCrawlable"`
			Typename    string `json:"__typename"`
		} `json:"moreLinks"`
		SiteLinks []struct {
			URL         string `json:"url"`
			ID          string `json:"id"`
			IsCrawlable bool   `json:"isCrawlable"`
			Typename    string `json:"__typename"`
		} `json:"siteLinks"`
		SocialLinks []struct {
			URL         string `json:"url"`
			ID          string `json:"id"`
			IsCrawlable bool   `json:"isCrawlable"`
			Typename    string `json:"__typename"`
		} `json:"socialLinks"`
		RestaurateursLinks []struct {
			URL         string `json:"url"`
			ID          string `json:"id"`
			IsCrawlable bool   `json:"isCrawlable"`
			Typename    string `json:"__typename"`
		} `json:"restaurateursLinks"`
		LegalLinks []struct {
			URL         string `json:"url"`
			ID          string `json:"id"`
			IsCrawlable bool   `json:"isCrawlable"`
			Typename    string `json:"__typename"`
		} `json:"legalLinks"`
	} `json:"footer"`
	Header struct {
		Crumbs []struct {
			Name string `json:"name"`
			Href string `json:"href,omitempty"`
		} `json:"crumbs"`
		MetroDomainID                   int           `json:"metroDomainId"`
		FeaturedMetros                  []interface{} `json:"featuredMetros"`
		MetroID                         int           `json:"metroId"`
		MacroID                         int           `json:"macroId"`
		NeighborhoodID                  int           `json:"neighborhoodId"`
		CurrentMetro                    interface{}   `json:"currentMetro"`
		IsLocationPickerLoading         bool          `json:"isLocationPickerLoading"`
		HideLocationPicker              bool          `json:"hideLocationPicker"`
		TimezoneOffset                  int           `json:"timezoneOffset"`
		ShouldDisplayOverlay            bool          `json:"shouldDisplayOverlay"`
		UseHeaderAuth                   bool          `json:"useHeaderAuth"`
		IsUserProfileLoading            bool          `json:"isUserProfileLoading"`
		UserProfile                     interface{}   `json:"userProfile"`
		UserTransactions                []interface{} `json:"userTransactions"`
		UserInvites                     []interface{} `json:"userInvites"`
		UserNotifications               []interface{} `json:"userNotifications"`
		HasUnreadNotifications          bool          `json:"hasUnreadNotifications"`
		UserNotificationsInitialLoading bool          `json:"userNotificationsInitialLoading"`
		UserNewsfeedStory               interface{}   `json:"userNewsfeedStory"`
		HasUnreadNewsfeedStory          bool          `json:"hasUnreadNewsfeedStory"`
		UserNewsfeedInitialLoading      bool          `json:"userNewsfeedInitialLoading"`
		ShowCovid19Banner               bool          `json:"showCovid19Banner"`
		CovidBannerCTALink              string        `json:"covidBannerCTALink"`
		HideLabel                       bool          `json:"hideLabel"`
	} `json:"header"`
	Language struct {
		Locale          string `json:"locale"`
		AcceptLanguage  string `json:"acceptLanguage"`
		DomainLanguages struct {
			En string `json:"en"`
			Fr string `json:"fr"`
			Es string `json:"es"`
			De string `json:"de"`
			Ja string `json:"ja"`
			Nl string `json:"nl"`
			It string `json:"it"`
		} `json:"domainLanguages"`
		CurrencySymbol interface{} `json:"currencySymbol"`
	} `json:"language"`
	SearchAutocomplete struct {
		AutocompleteResults   interface{} `json:"autocompleteResults"`
		CorrelationID         interface{} `json:"correlationId"`
		InitialMetroID        int         `json:"initialMetroId"`
		InitialMacroID        interface{} `json:"initialMacroId"`
		InitialNeighborhoodID int         `json:"initialNeighborhoodId"`
		Latitude              float64     `json:"latitude"`
		Longitude             float64     `json:"longitude"`
		SelectedItem          interface{} `json:"selectedItem"`
		InputValue            string      `json:"inputValue"`
		RecentSearchItems     interface{} `json:"recentSearchItems"`
		CurrentLocationName   interface{} `json:"currentLocationName"`
		CachedResults         struct {
		} `json:"cachedResults"`
	} `json:"searchAutocomplete"`
	SignalTracking struct {
		Page struct {
			Name          string `json:"name"`
			SeoPageType   string `json:"seoPageType"`
			SeoVertical   string `json:"seoVertical"`
			MarketingName string `json:"marketingName"`
		} `json:"page"`
		Reservation struct {
			Hash     interface{} `json:"hash"`
			ViewInfo interface{} `json:"viewInfo"`
		} `json:"reservation"`
	} `json:"signalTracking"`
	SojernTracking struct {
		SessionIDHash interface{} `json:"sessionIdHash"`
	} `json:"sojernTracking"`
	UserTracking struct {
		IsUserOptedIn         bool `json:"isUserOptedIn"`
		ConsentLoaded         bool `json:"consentLoaded"`
		ConsentPerformance    bool `json:"consentPerformance"`
		ConsentTargeting      bool `json:"consentTargeting"`
		SessionIDHashesLoaded bool `json:"sessionIdHashesLoaded"`
	} `json:"userTracking"`
	Availability struct {
		Loading                 bool `json:"loading"`
		RestaurantsAvailability struct {
		} `json:"restaurantsAvailability"`
		RestaurantsMultiDayAvailability struct {
		} `json:"restaurantsMultiDayAvailability"`
		PopRestaurantsAvailability struct {
		} `json:"popRestaurantsAvailability"`
		AvailabilityTime      string `json:"availabilityTime"`
		AvailabilityDate      string `json:"availabilityDate"`
		AvailabilityPartySize int    `json:"availabilityPartySize"`
		Scrollers             struct {
		} `json:"scrollers"`
		ScrollersLoading struct {
		} `json:"scrollersLoading"`
		AttributionToken string `json:"attributionToken"`
	} `json:"availability"`
	Faqs struct {
		FaqsData []struct {
			Question     string `json:"question"`
			Answer       string `json:"answer"`
			GtmEventData struct {
				Category string `json:"category"`
				Action   string `json:"action"`
				Label    string `json:"label"`
			} `json:"gtmEventData"`
		} `json:"faqsData"`
	} `json:"faqs"`
	MultiDayAvailabilityModal struct {
		IsLoading bool `json:"isLoading"`
		IsOpen    bool `json:"isOpen"`
	} `json:"multiDayAvailabilityModal"`
	PickupAvailabilityData struct {
		Loading   bool `json:"loading"`
		Variables struct {
			PickupAvailabilityParams struct {
				RestaurantID int    `json:"restaurantId"`
				DateTime     string `json:"dateTime"`
				ForwardDays  int    `json:"forwardDays"`
			} `json:"pickupAvailabilityParams"`
		} `json:"variables"`
		NetworkError interface{} `json:"networkError"`
		Data         struct {
		} `json:"data"`
		// Availability struct {
		// 	Two0220917T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-09-17T00:00:00"`
		// 	Two0220918T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-09-18T00:00:00"`
		// 	Two0220919T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-09-19T00:00:00"`
		// 	Two0220920T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-09-20T00:00:00"`
		// 	Two0220921T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-09-21T00:00:00"`
		// 	Two0220922T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-09-22T00:00:00"`
		// 	Two0220923T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-09-23T00:00:00"`
		// 	Two0220924T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-09-24T00:00:00"`
		// 	Two0220925T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-09-25T00:00:00"`
		// 	Two0220926T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-09-26T00:00:00"`
		// 	Two0220927T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-09-27T00:00:00"`
		// 	Two0220928T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-09-28T00:00:00"`
		// 	Two0220929T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-09-29T00:00:00"`
		// 	Two0220930T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-09-30T00:00:00"`
		// 	Two0221001T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-01T00:00:00"`
		// 	Two0221002T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-02T00:00:00"`
		// 	Two0221003T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-03T00:00:00"`
		// 	Two0221004T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-04T00:00:00"`
		// 	Two0221005T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-05T00:00:00"`
		// 	Two0221006T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-06T00:00:00"`
		// 	Two0221007T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-07T00:00:00"`
		// 	Two0221008T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-08T00:00:00"`
		// 	Two0221009T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-09T00:00:00"`
		// 	Two0221010T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-10T00:00:00"`
		// 	Two0221011T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-11T00:00:00"`
		// 	Two0221012T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-12T00:00:00"`
		// 	Two0221013T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-13T00:00:00"`
		// 	Two0221014T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-14T00:00:00"`
		// 	Two0221015T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-15T00:00:00"`
		// 	Two0221016T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-16T00:00:00"`
		// 	Two0221017T000000 struct {
		// 		PickupTimes       []interface{} `json:"pickupTimes"`
		// 		PickupAvailable   bool          `json:"pickupAvailable"`
		// 		LastPickupTime    string        `json:"lastPickupTime"`
		// 		LeadTimeInMinutes int           `json:"leadTimeInMinutes"`
		// 		NoTimesReasons    []interface{} `json:"noTimesReasons"`
		// 		DayOffset         int           `json:"dayOffset"`
		// 	} `json:"2022-10-17T00:00:00"`
		// } `json:"availability"`
		AvailableMenus struct {
		} `json:"availableMenus"`
		AvailableMenusByDateTime struct {
		} `json:"availableMenusByDateTime"`
	} `json:"pickupAvailabilityData"`
	RestaurantProfile struct {
		Lolz struct {
			IsLoading bool          `json:"isLoading"`
			LolzList  []interface{} `json:"lolzList"`
		} `json:"lolz"`
		UserWishlist struct {
		} `json:"userWishlist"`
		GalleryImageLoading bool `json:"galleryImageLoading"`
		GalleryImages       struct {
			Images struct {
				Photos []struct {
					Title      string    `json:"title"`
					Pid        int       `json:"pid"`
					Creator    string    `json:"creator"`
					Source     string    `json:"source"`
					Category   string    `json:"category"`
					IsProfile  bool      `json:"isProfile"`
					Width      int       `json:"width"`
					Height     int       `json:"height"`
					Timestamp  time.Time `json:"timestamp"`
					Thumbnails []struct {
						Width    int    `json:"width"`
						Height   int    `json:"height"`
						Label    string `json:"label"`
						URL      string `json:"url"`
						Typename string `json:"__typename"`
					} `json:"thumbnails"`
					Typename string `json:"__typename"`
				} `json:"photos"`
				Categories []struct {
					Name     string `json:"name"`
					Count    int    `json:"count"`
					Typename string `json:"__typename"`
				} `json:"categories"`
				CategoryPhotos struct {
					IsLoading bool `json:"isLoading"`
				} `json:"categoryPhotos"`
				CurrentCategory string `json:"currentCategory"`
				TotalCount      int    `json:"totalCount"`
				PageNumber      int    `json:"pageNumber"`
			} `json:"images"`
		} `json:"galleryImages"`
		GalleryPhotoModal struct {
			Show       bool `json:"show"`
			IsBanner   bool `json:"isBanner"`
			PhotoIndex int  `json:"photoIndex"`
		} `json:"galleryPhotoModal"`
		Restaurant struct {
			RestaurantID      int           `json:"restaurantId"`
			Name              string        `json:"name"`
			Description       string        `json:"description"`
			DiningStyle       string        `json:"diningStyle"`
			PublicTransit     string        `json:"publicTransit"`
			ExecutiveChef     interface{}   `json:"executiveChef"`
			Entertainment     interface{}   `json:"entertainment"`
			DniTags           []interface{} `json:"dniTags"`
			CateringDetails   string        `json:"cateringDetails"`
			Website           string        `json:"website"`
			HoursOfOperation  string        `json:"hoursOfOperation"`
			ParkingDetails    interface{}   `json:"parkingDetails"`
			ParkingInfo       string        `json:"parkingInfo"`
			DressCode         string        `json:"dressCode"`
			CrossStreet       interface{}   `json:"crossStreet"`
			AdditionalDetails []string      `json:"additionalDetails"`
			PaymentOptions    []interface{} `json:"paymentOptions"`
			Type              string        `json:"type"`
			StateID           int           `json:"stateId"`
			MaxAdvanceDays    int           `json:"maxAdvanceDays"`
			EditorialLists    struct {
				CorrelationID string `json:"correlationId"`
				Lists         []struct {
					CorrelationID string `json:"correlationId"`
					Name          string `json:"name"`
					TopicLink     struct {
						Link     string `json:"link"`
						Typename string `json:"__typename"`
					} `json:"topicLink"`
					Topic struct {
						TopicID           string      `json:"topicId"`
						ThumbnailMetadata interface{} `json:"thumbnailMetadata"`
						Typename          string      `json:"__typename"`
					} `json:"topic"`
					Typename string `json:"__typename"`
				} `json:"lists"`
				Typename string `json:"__typename"`
			} `json:"editorialLists"`
			MostRecentReview struct {
				Reviews []struct {
					Text     string `json:"text"`
					Typename string `json:"__typename"`
				} `json:"reviews"`
				Typename string `json:"__typename"`
			} `json:"mostRecentReview"`
			Photos struct {
				Gallery struct {
					Photos []struct {
						Title      string    `json:"title"`
						Pid        int       `json:"pid"`
						Creator    string    `json:"creator"`
						Source     string    `json:"source"`
						Category   string    `json:"category"`
						IsProfile  bool      `json:"isProfile"`
						Width      int       `json:"width"`
						Height     int       `json:"height"`
						Timestamp  time.Time `json:"timestamp"`
						Thumbnails []struct {
							Width    int    `json:"width"`
							Height   int    `json:"height"`
							Label    string `json:"label"`
							URL      string `json:"url"`
							Typename string `json:"__typename"`
						} `json:"thumbnails"`
						Typename string `json:"__typename"`
					} `json:"photos"`
					Typename   string `json:"__typename"`
					TotalCount int    `json:"totalCount"`
					Categories []struct {
						Name     string `json:"name"`
						Count    int    `json:"count"`
						Typename string `json:"__typename"`
					} `json:"categories"`
				} `json:"gallery"`
				Profile struct {
					Large struct {
						URL      string `json:"url"`
						Width    int    `json:"width"`
						Height   int    `json:"height"`
						Typename string `json:"__typename"`
					} `json:"large"`
					Typename string `json:"__typename"`
				} `json:"profile"`
				Typename string `json:"__typename"`
			} `json:"photos"`
			Address struct {
				Line1    string `json:"line1"`
				Line2    string `json:"line2"`
				State    string `json:"state"`
				City     string `json:"city"`
				PostCode string `json:"postCode"`
				Country  string `json:"country"`
				Typename string `json:"__typename"`
			} `json:"address"`
			ContactInformation struct {
				FormattedPhoneNumber string `json:"formattedPhoneNumber"`
				Typename             string `json:"__typename"`
			} `json:"contactInformation"`
			Metro struct {
				MetroID    int `json:"metroId"`
				Statistics struct {
					RestaurantCount int    `json:"restaurantCount"`
					Typename        string `json:"__typename"`
				} `json:"statistics"`
				Typename string `json:"__typename"`
			} `json:"metro"`
			TimeZone struct {
				OffsetInMinutes int    `json:"offsetInMinutes"`
				Typename        string `json:"__typename"`
			} `json:"timeZone"`
			Coordinates struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				MapURL    string  `json:"mapUrl"`
				Typename  string  `json:"__typename"`
			} `json:"coordinates"`
			PriceBand struct {
				PriceBandID int    `json:"priceBandId"`
				Name        string `json:"name"`
				Typename    string `json:"__typename"`
			} `json:"priceBand"`
			Statistics struct {
				RecentReservationCount int `json:"recentReservationCount"`
				Reviews                struct {
					AllTimeTextReviewCount int    `json:"allTimeTextReviewCount"`
					ShowReviews            bool   `json:"showReviews"`
					TotalNumberOfReviews   int    `json:"totalNumberOfReviews"`
					VerifiedReviews        bool   `json:"verifiedReviews"`
					RatingBasedOn          string `json:"ratingBasedOn"`
					Awards                 []struct {
						Location string `json:"location"`
						Name     string `json:"name"`
						Typename string `json:"__typename"`
					} `json:"awards"`
					Ratings struct {
						Count   int `json:"count"`
						Overall struct {
							Rating   float64 `json:"rating"`
							Typename string  `json:"__typename"`
						} `json:"overall"`
						Ambience struct {
							Rating   float64 `json:"rating"`
							Count    int     `json:"count"`
							Typename string  `json:"__typename"`
						} `json:"ambience"`
						Food struct {
							Rating   float64 `json:"rating"`
							Count    int     `json:"count"`
							Typename string  `json:"__typename"`
						} `json:"food"`
						Service struct {
							Rating   float64 `json:"rating"`
							Count    int     `json:"count"`
							Typename string  `json:"__typename"`
						} `json:"service"`
						Value struct {
							Rating   float64 `json:"rating"`
							Count    int     `json:"count"`
							Typename string  `json:"__typename"`
						} `json:"value"`
						Noise struct {
							Rating   int    `json:"rating"`
							Count    int    `json:"count"`
							Typename string `json:"__typename"`
						} `json:"noise"`
						OverallRatingsDistribution struct {
							One      int    `json:"one"`
							Two      int    `json:"two"`
							Three    int    `json:"three"`
							Four     int    `json:"four"`
							Five     int    `json:"five"`
							Typename string `json:"__typename"`
						} `json:"overallRatingsDistribution"`
						Typename string `json:"__typename"`
					} `json:"ratings"`
					ChangeOfConceptType        interface{} `json:"changeOfConceptType"`
					ChangeOfConceptDateTimeUtc interface{} `json:"changeOfConceptDateTimeUtc"`
					Categories                 []struct {
						Category struct {
							Name     string `json:"name"`
							Key      string `json:"key"`
							Typename string `json:"__typename"`
						} `json:"category"`
						Typename string `json:"__typename"`
					} `json:"categories"`
					Typename string `json:"__typename"`
				} `json:"reviews"`
				Typename string `json:"__typename"`
			} `json:"statistics"`
			PrimaryCuisine struct {
				CuisineID string `json:"cuisineId"`
				Name      string `json:"name"`
				Typename  string `json:"__typename"`
			} `json:"primaryCuisine"`
			Experiences []interface{} `json:"experiences"`
			Cuisines    []struct {
				Name     string `json:"name"`
				Typename string `json:"__typename"`
			} `json:"cuisines"`
			Macro struct {
				MacroID  int    `json:"macroId"`
				Typename string `json:"__typename"`
			} `json:"macro"`
			Features struct {
				PermanentlyClosed           interface{} `json:"permanentlyClosed"`
				WaitlistV2Only              bool        `json:"waitlistV2Only"`
				WaitlistV2Enabled           bool        `json:"waitlistV2Enabled"`
				NetworkNonBookable          interface{} `json:"networkNonBookable"`
				PrioritySeatingEnabled      interface{} `json:"prioritySeatingEnabled"`
				Bar                         bool        `json:"bar"`
				Counter                     bool        `json:"counter"`
				HighTop                     bool        `json:"highTop"`
				Outdoor                     bool        `json:"outdoor"`
				OptsOutOfLoyaltyRedemptions interface{} `json:"optsOutOfLoyaltyRedemptions"`
				PickupEnabled               interface{} `json:"pickupEnabled"`
				GiftsURLs                   interface{} `json:"giftsURLs"`
				GiftsEnabled                interface{} `json:"giftsEnabled"`
				Typename                    string      `json:"__typename"`
			} `json:"features"`
			SafetyPrecautions struct {
				SanitizedSurfaces             interface{} `json:"sanitizedSurfaces"`
				CommonAreaCleaning            interface{} `json:"commonAreaCleaning"`
				CleanMenus                    interface{} `json:"cleanMenus"`
				SanitizerProvidedForCustomers interface{} `json:"sanitizerProvidedForCustomers"`
				ContactlessPayment            interface{} `json:"contactlessPayment"`
				SealedUtensils                interface{} `json:"sealedUtensils"`
				LimitedSeating                interface{} `json:"limitedSeating"`
				CommonAreaDistancing          interface{} `json:"commonAreaDistancing"`
				TableLayoutWithExtraSpace     interface{} `json:"tableLayoutWithExtraSpace"`
				RequireWaitstaffMasks         interface{} `json:"requireWaitstaffMasks"`
				RequireDinerMasks             interface{} `json:"requireDinerMasks"`
				ProhibitSickStaff             interface{} `json:"prohibitSickStaff"`
				StaffTempChecksRequired       interface{} `json:"staffTempChecksRequired"`
				DinerTempChecksRequired       interface{} `json:"dinerTempChecksRequired"`
				ContactTracingCollected       interface{} `json:"contactTracingCollected"`
				StaffIsVaccinated             interface{} `json:"staffIsVaccinated"`
				ProofOfVaccinationRequired    bool        `json:"proofOfVaccinationRequired"`
				ProofOfVaccinationOutdoor     interface{} `json:"proofOfVaccinationOutdoor"`
				Typename                      string      `json:"__typename"`
			} `json:"safetyPrecautions"`
			Neighborhood struct {
				NeighborhoodID int    `json:"neighborhoodId"`
				Name           string `json:"name"`
				Coordinates    struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
					Typename  string  `json:"__typename"`
				} `json:"coordinates"`
				Typename string `json:"__typename"`
			} `json:"neighborhood"`
			HasTakeout            bool          `json:"hasTakeout"`
			HasDeliveryPartners   bool          `json:"hasDeliveryPartners"`
			HasDeliveryDirect     bool          `json:"hasDeliveryDirect"`
			HasWhiteLabelDelivery bool          `json:"hasWhiteLabelDelivery"`
			OrderOnlineLink       interface{}   `json:"orderOnlineLink"`
			WhiteLabelDelivery    interface{}   `json:"whiteLabelDelivery"`
			DeliveryPartners      []interface{} `json:"deliveryPartners"`
			FacebookURL           interface{}   `json:"facebookUrl"`
			Localization          struct {
				Address  string `json:"address"`
				Typename string `json:"__typename"`
			} `json:"localization"`
			Country struct {
				CountryID string `json:"countryId"`
				Typename  string `json:"__typename"`
			} `json:"country"`
			PopularDishes struct {
				CorrelationID interface{}   `json:"correlationId"`
				Dishes        []interface{} `json:"dishes"`
				Typename      string        `json:"__typename"`
			} `json:"popularDishes"`
			DiningAreas []struct {
				Name         string `json:"name"`
				DiningAreaID int    `json:"diningAreaId"`
				Description  string `json:"description"`
				Tag          string `json:"tag"`
				LocaleMatch  bool   `json:"localeMatch"`
				Active       bool   `json:"active"`
				Typename     string `json:"__typename"`
			} `json:"diningAreas"`
			PrivateDining struct {
				HasPrivateDining         bool        `json:"hasPrivateDining"`
				HasEnhancedPrivateDining bool        `json:"hasEnhancedPrivateDining"`
				Description              string      `json:"description"`
				FormattedPhone           interface{} `json:"formattedPhone"`
				Contact                  string      `json:"contact"`
				Typename                 string      `json:"__typename"`
			} `json:"privateDining"`
			Typename string `json:"__typename"`
		} `json:"restaurant"`
		Menus struct {
			MenuData []struct {
				Title       string `json:"title"`
				Description string `json:"description"`
				Currency    struct {
					Code     string `json:"code"`
					Typename string `json:"__typename"`
				} `json:"currency"`
				Sections []struct {
					Title       string `json:"title"`
					Description string `json:"description"`
					Items       []struct {
						Title           string        `json:"title"`
						Description     string        `json:"description"`
						Price           interface{}   `json:"price"`
						VariationGroups []interface{} `json:"variationGroups"`
						Typename        string        `json:"__typename"`
					} `json:"items"`
					Typename string `json:"__typename"`
				} `json:"sections"`
				Updated  string `json:"updated"`
				Provider struct {
					Name  string `json:"name"`
					Image struct {
						AltText  string `json:"altText"`
						URL      string `json:"url"`
						Height   int    `json:"height"`
						Typename string `json:"__typename"`
					} `json:"image"`
					Typename string `json:"__typename"`
				} `json:"provider"`
				Typename string `json:"__typename"`
			} `json:"menuData"`
			MenuInfo struct {
				ShowThirdPartyMenu bool   `json:"showThirdPartyMenu"`
				URL                string `json:"url"`
				Typename           string `json:"__typename"`
			} `json:"menuInfo"`
		} `json:"menus"`
		IsHuman   bool `json:"isHuman"`
		MetaLinks struct {
			IsIndexed bool `json:"isIndexed"`
			HrefLangs []struct {
				Href     string `json:"href"`
				HrefLang string `json:"hrefLang"`
				Typename string `json:"__typename"`
			} `json:"hrefLangs"`
			Canonical string `json:"canonical"`
			Typename  string `json:"__typename"`
		} `json:"metaLinks"`
		ReportReviewModal struct {
			IsOpen    bool        `json:"isOpen"`
			ReviewID  interface{} `json:"reviewId"`
			Success   bool        `json:"success"`
			Error     bool        `json:"error"`
			Submitted bool        `json:"submitted"`
		} `json:"reportReviewModal"`
		ReviewSearchResults struct {
			Page        int         `json:"page"`
			SortBy      string      `json:"sortBy"`
			FilterIndex interface{} `json:"filterIndex"`
			Reviews     []struct {
				ReviewID string `json:"reviewId"`
				Text     string `json:"text"`
				Rating   struct {
					Food     int    `json:"food"`
					Service  int    `json:"service"`
					Ambience int    `json:"ambience"`
					Value    int    `json:"value"`
					Noise    string `json:"noise"`
					Overall  int    `json:"overall"`
					Typename string `json:"__typename"`
				} `json:"rating"`
				User struct {
					NumOfApprovedReviews int         `json:"numOfApprovedReviews"`
					Nickname             interface{} `json:"nickname"`
					Metro                interface{} `json:"metro"`
					Initials             interface{} `json:"initials"`
					IsVip                bool        `json:"isVip"`
					ProfileColor         string      `json:"profileColor"`
					Typename             string      `json:"__typename"`
				} `json:"user"`
				PublicRestaurantReply interface{} `json:"publicRestaurantReply"`
				SubmittedDateTime     time.Time   `json:"submittedDateTime"`
				DinedDateTime         string      `json:"dinedDateTime"`
				Type                  string      `json:"type"`
				FromPreviousConcept   bool        `json:"fromPreviousConcept"`
				Highlight             interface{} `json:"highlight"`
				Typename              string      `json:"__typename"`
			} `json:"reviews"`
			IsLoading              bool `json:"isLoading"`
			PrioritiseUserLanguage bool `json:"prioritiseUserLanguage"`
			Aggregations           struct {
				Dishes []struct {
					ID       string `json:"id"`
					Label    string `json:"label"`
					Count    int    `json:"count"`
					Typename string `json:"__typename"`
				} `json:"dishes"`
				Typename string `json:"__typename"`
			} `json:"aggregations"`
			TotalPages int    `json:"totalPages"`
			TotalCount int    `json:"totalCount"`
			Typename   string `json:"__typename"`
		} `json:"reviewSearchResults"`
		AvailabilityToken              string      `json:"availabilityToken"`
		NoAvailabilityList             interface{} `json:"noAvailabilityList"`
		MultiDayAvailabilityModalOpen  bool        `json:"multiDayAvailabilityModalOpen"`
		ShowDiningReward               bool        `json:"showDiningReward"`
		InitialDTPDate                 interface{} `json:"initialDTPDate"`
		InitialDTPTime                 interface{} `json:"initialDTPTime"`
		InitialDTPPartySize            interface{} `json:"initialDTPPartySize"`
		FetchAvailabilityOnLoad        bool        `json:"fetchAvailabilityOnLoad"`
		FindATableLoadingStateOverride bool        `json:"findATableLoadingStateOverride"`
		PointRedemptionRewards         interface{} `json:"pointRedemptionRewards"`
		PrivateDining                  struct {
		} `json:"privateDining"`
		PageAnchor interface{} `json:"pageAnchor"`
	} `json:"restaurantProfile"`
}

type RestaurantDetailsMenu struct {
	Title       string
	Description string
	Currency    string
	Sections    []RestaurantDetailsMenuSection
}

type RestaurantDetailsMenuSection struct {
	Title       string
	Description string
	Items       []RestaurantDetailsMenuSectionItem
}

type CurrencyAmount float32

func (c CurrencyAmount) String() string {
	return fmt.Sprintf("$%0.2f", c)
}

func toCurrencyAmount(x interface{}) (CurrencyAmount, error) {
	if x == nil {
		return 0, nil
	}
	switch x := x.(type) {
	case float64:
		return CurrencyAmount(x), nil
	case float32:
		return CurrencyAmount(x), nil
	case int:
		return CurrencyAmount(x), nil
	case int64:
		return CurrencyAmount(x), nil
	case string:
		f, err := strconv.ParseFloat(x, 32)
		if err != nil {
			return 0, err
		}
		return CurrencyAmount(f), nil
	default:
		return 0, errors.Errorf("unknown type %T", x)
	}
}

type RestaurantDetailsMenuSectionItem struct {
	Title       string
	Description string
	Price       CurrencyAmount
}

type RestaurantDetails struct {
	RestaurantID      int
	Name              string
	Description       string
	DiningStyle       string
	PublicTransit     string
	ExecutiveChef     interface{}
	Entertainment     interface{}
	DniTags           []interface{}
	CateringDetails   string
	Website           string
	HoursOfOperation  string
	ParkingDetails    interface{}
	ParkingInfo       string
	DressCode         string
	CrossStreet       interface{}
	AdditionalDetails []string
	PaymentOptions    []interface{}
	Type              string
	StateID           int
	MaxAdvanceDays    int
	Menus             []RestaurantDetailsMenu
}

type RestaurantDetailsInfo struct {
	RestaurantDetails RestaurantDetails
}

func (payload RawRestaurantDetails) Convert() (*RestaurantDetailsInfo, error) {
	var menus []RestaurantDetailsMenu
	for _, menu := range payload.RestaurantProfile.Menus.MenuData {
		var sections []RestaurantDetailsMenuSection
		for _, s := range menu.Sections {
			var items []RestaurantDetailsMenuSectionItem
			for _, it := range s.Items {
				price, err := toCurrencyAmount(it.Price)
				if err != nil {
					return nil, err
				}
				items = append(items, RestaurantDetailsMenuSectionItem{
					Title:       it.Title,
					Description: it.Description,
					Price:       price,
				})
			}
			sections = append(sections, RestaurantDetailsMenuSection{
				Title:       s.Title,
				Description: s.Description,
				Items:       items,
			})
		}
		menus = append(menus, RestaurantDetailsMenu{
			Title:       menu.Title,
			Description: menu.Description,
			Currency:    menu.Currency.Code,
			Sections:    sections,
		})
	}
	res := &RestaurantDetailsInfo{
		RestaurantDetails: RestaurantDetails{
			RestaurantID:      payload.RestaurantProfile.Restaurant.RestaurantID,
			Name:              payload.RestaurantProfile.Restaurant.Name,
			Description:       payload.RestaurantProfile.Restaurant.Description,
			DiningStyle:       payload.RestaurantProfile.Restaurant.DiningStyle,
			PublicTransit:     payload.RestaurantProfile.Restaurant.PublicTransit,
			ExecutiveChef:     payload.RestaurantProfile.Restaurant.ExecutiveChef,
			Entertainment:     payload.RestaurantProfile.Restaurant.Entertainment,
			DniTags:           payload.RestaurantProfile.Restaurant.DniTags,
			CateringDetails:   payload.RestaurantProfile.Restaurant.CateringDetails,
			Website:           payload.RestaurantProfile.Restaurant.Website,
			HoursOfOperation:  payload.RestaurantProfile.Restaurant.HoursOfOperation,
			ParkingDetails:    payload.RestaurantProfile.Restaurant.ParkingDetails,
			ParkingInfo:       payload.RestaurantProfile.Restaurant.ParkingInfo,
			DressCode:         payload.RestaurantProfile.Restaurant.DressCode,
			CrossStreet:       payload.RestaurantProfile.Restaurant.CrossStreet,
			AdditionalDetails: payload.RestaurantProfile.Restaurant.AdditionalDetails,
			PaymentOptions:    payload.RestaurantProfile.Restaurant.PaymentOptions,
			Type:              payload.RestaurantProfile.Restaurant.Type,
			StateID:           payload.RestaurantProfile.Restaurant.StateID,
			MaxAdvanceDays:    payload.RestaurantProfile.Restaurant.MaxAdvanceDays,
			Menus:             menus,
		},
	}
	return res, nil
}
