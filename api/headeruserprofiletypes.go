package api

type HeaderUserProfilePhoneNumber struct {
	Number    string
	CountryID string
}

type HeaderUserProfileMetroInfo struct {
	Name        string
	DisplayName string
}

type HeaderUserProfile struct {
	FirstName            string
	LastName             string
	SortableFirstName    string
	SortableLastName     string
	UserType             int
	IsVip                bool
	IsConcierge          bool
	CountryID            string
	MetroID              int
	Metro                HeaderUserProfileMetroInfo
	CreateDate           string
	Points               int
	Email                string
	Gpid                 int64
	HomePhoneNumber      HeaderUserProfilePhoneNumber
	MobilePhoneNumber    HeaderUserProfilePhoneNumber
	EligibleToEarnPoints bool
	NumOfApprovedReviews int
}

type HeaderUserProfileInfo struct {
	UserProfile HeaderUserProfile
}

type headerUserProfilePayload struct {
	Data struct {
		UserProfile struct {
			FirstName         string `json:"firstName"`
			LastName          string `json:"lastName"`
			SortableFirstName string `json:"sortableFirstName"`
			SortableLastName  string `json:"sortableLastName"`
			UserType          int    `json:"userType"`
			IsVip             bool   `json:"isVip"`
			IsConcierge       bool   `json:"isConcierge"`
			CountryID         string `json:"countryId"`
			MetroID           int    `json:"metroId"`
			Metro             struct {
				DisplayName string `json:"displayName"`
				Name        string `json:"name"`
				Typename    string `json:"__typename"`
			} `json:"metro"`
			ReviewUser struct {
				NumOfApprovedReviews int    `json:"numOfApprovedReviews"`
				Typename             string `json:"__typename"`
			} `json:"reviewUser"`
			CreateDate      string `json:"createDate"`
			Points          int    `json:"points"`
			Email           string `json:"email"`
			Gpid            int64  `json:"gpid"`
			HomePhoneNumber struct {
				Number    string `json:"number"`
				CountryID string `json:"countryId"`
				Typename  string `json:"__typename"`
			} `json:"homePhoneNumber"`
			MobilePhoneNumber struct {
				Number    string `json:"number"`
				CountryID string `json:"countryId"`
				Typename  string `json:"__typename"`
			} `json:"mobilePhoneNumber"`
			EligibleToEarnPoints bool   `json:"eligibleToEarnPoints"`
			Typename             string `json:"__typename"`
		} `json:"userProfile"`
		UserUpcomingTransactions []interface{} `json:"userUpcomingTransactions"`
		Typename                 string        `json:"__typename"`
		InvitesByGpid            interface{}   `json:"invitesByGpid"`
	} `json:"data"`
	Loading       bool `json:"loading"`
	NetworkStatus int  `json:"networkStatus"`
}

func (payload headerUserProfilePayload) Convert() *HeaderUserProfileInfo {
	userProfile := HeaderUserProfile{
		FirstName:         payload.Data.UserProfile.FirstName,
		LastName:          payload.Data.UserProfile.LastName,
		SortableFirstName: payload.Data.UserProfile.SortableFirstName,
		SortableLastName:  payload.Data.UserProfile.SortableLastName,
		UserType:          payload.Data.UserProfile.UserType,
		IsVip:             payload.Data.UserProfile.IsVip,
		IsConcierge:       payload.Data.UserProfile.IsConcierge,
		CountryID:         payload.Data.UserProfile.CountryID,
		MetroID:           payload.Data.UserProfile.MetroID,
		Metro: HeaderUserProfileMetroInfo{
			Name:        payload.Data.UserProfile.Metro.Name,
			DisplayName: payload.Data.UserProfile.Metro.DisplayName,
		},
		CreateDate: payload.Data.UserProfile.CreateDate,
		Points:     payload.Data.UserProfile.Points,
		Email:      payload.Data.UserProfile.Email,
		Gpid:       payload.Data.UserProfile.Gpid,
		HomePhoneNumber: HeaderUserProfilePhoneNumber{
			Number:    payload.Data.UserProfile.HomePhoneNumber.Number,
			CountryID: payload.Data.UserProfile.HomePhoneNumber.CountryID,
		},
		MobilePhoneNumber: HeaderUserProfilePhoneNumber{
			Number:    payload.Data.UserProfile.MobilePhoneNumber.Number,
			CountryID: payload.Data.UserProfile.MobilePhoneNumber.CountryID,
		},
		NumOfApprovedReviews: payload.Data.UserProfile.ReviewUser.NumOfApprovedReviews,
	}
	return &HeaderUserProfileInfo{
		UserProfile: userProfile,
	}
}
