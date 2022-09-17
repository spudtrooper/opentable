package api

type RestaurantsAvailabilityDaySlot struct {
	IsAvailable           bool
	TimeOffsetMinutes     int
	SlotHash              string
	PointsType            string
	PointsValue           int
	ExperienceIds         []interface{}
	SlotAvailabilityToken string
	Attributes            []string
	IsMandatory           bool
	IsMandatoryBySeating  []IsMandatoryBySeating
	ExperiencesBySeating  []interface{}
	RedemptionTier        string
	Type                  string
}

type RestaurantsAvailabilityDay struct {
	NoTimesReasons     []interface{}
	EarlyCutoff        interface{}
	SameDayCutoff      interface{}
	DayOffset          int
	AllowNextAvailable bool
	TopExperience      interface{}
	Slots              []RestaurantsAvailabilityDaySlot
}

type RestaurantsAvailability struct {
	RestaurantID                int
	RestaurantAvailabilityToken string
	AvailabilityDays            []RestaurantsAvailabilityDay
}

type RestaurantsAvailabilityInfo struct {
	RestaurantsAvailabilities []RestaurantsAvailability
}

type restaurantAvailabilityPayload struct {
	Data struct {
		Availability []struct {
			RestaurantID                int    `json:"restaurantId"`
			RestaurantAvailabilityToken string `json:"restaurantAvailabilityToken"`
			AvailabilityDays            []struct {
				NoTimesReasons     []interface{} `json:"noTimesReasons"`
				EarlyCutoff        interface{}   `json:"earlyCutoff"`
				SameDayCutoff      interface{}   `json:"sameDayCutoff"`
				DayOffset          int           `json:"dayOffset"`
				AllowNextAvailable bool          `json:"allowNextAvailable"`
				TopExperience      interface{}   `json:"topExperience"`
				Slots              []struct {
					IsAvailable           bool          `json:"isAvailable"`
					Typename              string        `json:"__typename"`
					TimeOffsetMinutes     int           `json:"timeOffsetMinutes,omitempty"`
					SlotHash              string        `json:"slotHash,omitempty"`
					PointsType            string        `json:"pointsType,omitempty"`
					PointsValue           int           `json:"pointsValue,omitempty"`
					ExperienceIds         []interface{} `json:"experienceIds,omitempty"`
					SlotAvailabilityToken string        `json:"slotAvailabilityToken,omitempty"`
					Attributes            []string      `json:"attributes,omitempty"`
					IsMandatory           bool          `json:"isMandatory,omitempty"`
					IsMandatoryBySeating  []struct {
						TableCategory string `json:"tableCategory"`
						IsMandatory   bool   `json:"isMandatory"`
						Typename      string `json:"__typename"`
					} `json:"isMandatoryBySeating,omitempty"`
					ExperiencesBySeating []interface{} `json:"experiencesBySeating,omitempty"`
					RedemptionTier       string        `json:"redemptionTier,omitempty"`
					Type                 string        `json:"type,omitempty"`
				} `json:"slots"`
				Typename string `json:"__typename"`
			} `json:"availabilityDays"`
			Typename string `json:"__typename"`
		} `json:"availability"`
	} `json:"data"`
	Loading       bool `json:"loading"`
	NetworkStatus int  `json:"networkStatus"`
}

func (payload restaurantAvailabilityPayload) Convert() *RestaurantsAvailabilityInfo {
	var restaurantsAvailabilities []RestaurantsAvailability
	for _, availability := range payload.Data.Availability {
		var ads []RestaurantsAvailabilityDay
		for _, ad := range availability.AvailabilityDays {
			var slots []RestaurantsAvailabilityDaySlot
			for _, slot := range ad.Slots {
				var isMandatoryBySeatings []IsMandatoryBySeating
				for _, ism := range slot.IsMandatoryBySeating {
					isMandatoryBySeatings = append(isMandatoryBySeatings, IsMandatoryBySeating{
						TableCategory: ism.TableCategory,
						IsMandatory:   ism.IsMandatory,
					})
				}
				slots = append(slots, RestaurantsAvailabilityDaySlot{
					IsAvailable:           slot.IsAvailable,
					TimeOffsetMinutes:     slot.TimeOffsetMinutes,
					SlotHash:              slot.SlotHash,
					PointsType:            slot.PointsType,
					PointsValue:           slot.PointsValue,
					ExperienceIds:         slot.ExperienceIds,
					SlotAvailabilityToken: slot.SlotAvailabilityToken,
					Attributes:            slot.Attributes,
					IsMandatory:           slot.IsMandatory,
					IsMandatoryBySeating:  isMandatoryBySeatings,
					ExperiencesBySeating:  slot.ExperiencesBySeating,
					RedemptionTier:        slot.RedemptionTier,
					Type:                  slot.Type,
				})
			}
			ads = append(ads, RestaurantsAvailabilityDay{
				NoTimesReasons:     ad.NoTimesReasons,
				EarlyCutoff:        ad.EarlyCutoff,
				SameDayCutoff:      ad.SameDayCutoff,
				DayOffset:          ad.DayOffset,
				AllowNextAvailable: ad.AllowNextAvailable,
				TopExperience:      ad.TopExperience,
				Slots:              slots,
			})
		}
		ra := RestaurantsAvailability{
			RestaurantID:                availability.RestaurantID,
			RestaurantAvailabilityToken: availability.RestaurantAvailabilityToken,
			AvailabilityDays:            ads,
		}
		restaurantsAvailabilities = append(restaurantsAvailabilities, ra)
	}
	return &RestaurantsAvailabilityInfo{
		RestaurantsAvailabilities: restaurantsAvailabilities,
	}
}
