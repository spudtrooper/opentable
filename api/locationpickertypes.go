package api

type FeaturedMetrosForDomainInfo struct {
	Name            string
	MetroID         int
	RestaurantCount int
	StartPageLink   string
	Macros          []MacroInfo
}

type MacroInfo struct {
	Name            string
	MacroID         int
	Active          bool
	RestaurantCount int
	StartPageLink   string
}

type MetroInfo struct {
	Name            string
	MetroID         int
	StartPageLink   string
	RestaurantCount int
	Macros          []MacroInfo
}

type LocationPickerInfo struct {
	FeaturedMetrosListForDomain []FeaturedMetrosForDomainInfo
	Metro                       MetroInfo
}

type locationPickerPayload struct {
	Data struct {
		FeaturedMetrosListForDomain []struct {
			Name       string `json:"name"`
			MetroID    int    `json:"metroId"`
			Statistics struct {
				RestaurantCount int    `json:"restaurantCount"`
				Typename        string `json:"__typename"`
			} `json:"statistics"`
			Urls struct {
				StartPageLink struct {
					Link     string `json:"link"`
					Typename string `json:"__typename"`
				} `json:"startPageLink"`
				Typename string `json:"__typename"`
			} `json:"urls"`
			Macros []struct {
				Name       string `json:"name"`
				MacroID    int    `json:"macroId"`
				Active     bool   `json:"active"`
				Statistics struct {
					RestaurantCount int    `json:"restaurantCount"`
					Typename        string `json:"__typename"`
				} `json:"statistics"`
				Urls struct {
					StartPageLink struct {
						Link     string `json:"link"`
						Typename string `json:"__typename"`
					} `json:"startPageLink"`
					Typename string `json:"__typename"`
				} `json:"urls"`
				Typename string `json:"__typename"`
			} `json:"macros"`
			Typename string `json:"__typename"`
		} `json:"featuredMetrosListForDomain"`
		Metro struct {
			Name       string `json:"name"`
			MetroID    int    `json:"metroId"`
			Statistics struct {
				RestaurantCount int    `json:"restaurantCount"`
				Typename        string `json:"__typename"`
			} `json:"statistics"`
			Urls struct {
				StartPageLink struct {
					Link     string `json:"link"`
					Typename string `json:"__typename"`
				} `json:"startPageLink"`
				Typename string `json:"__typename"`
			} `json:"urls"`
			Macros []struct {
				Name       string `json:"name"`
				MacroID    int    `json:"macroId"`
				Active     bool   `json:"active"`
				Statistics struct {
					RestaurantCount int    `json:"restaurantCount"`
					Typename        string `json:"__typename"`
				} `json:"statistics"`
				Urls struct {
					StartPageLink struct {
						Link     string `json:"link"`
						Typename string `json:"__typename"`
					} `json:"startPageLink"`
					Typename string `json:"__typename"`
				} `json:"urls"`
				Typename string `json:"__typename"`
			} `json:"macros"`
			Typename string `json:"__typename"`
		} `json:"metro"`
	} `json:"data"`
}

func (payload locationPickerPayload) Convert() *LocationPickerInfo {
	var macros []MacroInfo
	for _, m := range payload.Data.Metro.Macros {
		macros = append(macros, MacroInfo{
			Name:            m.Name,
			MacroID:         m.MacroID,
			Active:          m.Active,
			RestaurantCount: m.Statistics.RestaurantCount,
			StartPageLink:   m.Urls.StartPageLink.Link,
		})
	}
	var featuredMetrosListForDomain []FeaturedMetrosForDomainInfo
	for _, f := range payload.Data.FeaturedMetrosListForDomain {
		var macros []MacroInfo
		for _, m := range f.Macros {
			macros = append(macros, MacroInfo{
				Name:            m.Name,
				MacroID:         m.MacroID,
				Active:          m.Active,
				RestaurantCount: m.Statistics.RestaurantCount,
				StartPageLink:   m.Urls.StartPageLink.Link,
			})
		}
		featuredMetrosListForDomain = append(featuredMetrosListForDomain, FeaturedMetrosForDomainInfo{
			Name:            f.Name,
			MetroID:         f.MetroID,
			RestaurantCount: f.Statistics.RestaurantCount,
			StartPageLink:   f.Urls.StartPageLink.Link,
			Macros:          macros,
		})
	}
	return &LocationPickerInfo{
		FeaturedMetrosListForDomain: featuredMetrosListForDomain,
		Metro: MetroInfo{
			Name:            payload.Data.Metro.Name,
			MetroID:         payload.Data.Metro.MetroID,
			StartPageLink:   payload.Data.Metro.Urls.StartPageLink.Link,
			RestaurantCount: payload.Data.Metro.Statistics.RestaurantCount,
			Macros:          macros,
		},
	}
}
