package api

type RawLolzViewAllLinksInfo struct {
	LolzViewAll struct {
		SearchResults struct {
			Restaurants []struct {
				Urls struct {
					ProfileLink struct {
						Link     string `json:"link"`
						Typename string `json:"__typename"`
					} `json:"profileLink"`
					Typename string `json:"__typename"`
				} `json:"urls"`
			} `json:"restaurants"`
		} `json:"searchResults"`
	} `json:"lolzViewAll"`
}

func (payload RawLolzViewAllLinksInfo) Convert() LolzViewAllLinksInfo {
	var profileLinks []string
	for _, r := range payload.LolzViewAll.SearchResults.Restaurants {
		profileLinks = append(profileLinks, r.Urls.ProfileLink.Link)
	}
	return LolzViewAllLinksInfo{
		ProfileLinks: profileLinks,
	}
}

type LolzViewAllLinksInfo struct {
	ProfileLinks []string
}
