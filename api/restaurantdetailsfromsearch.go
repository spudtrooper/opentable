package api

//go:generate genopts --function RestaurantDetailsFromSearch --params --required "term string" --extends Search,RestaurantDetails
func (e *Extended) RestaurantDetailsFromSearch(term string, optss ...RestaurantDetailsFromSearchOption) (*RestaurantDetailsInfo, error) {
	opts := MakeRestaurantDetailsFromSearchOptions(optss...)
	searchInfo, err := e.Search(term, opts.ToSearchOptions()...)
	if err != nil {
		return nil, err
	}
	return e.RestaurantDetails(searchInfo.Restaurants[0], opts.ToRestaurantDetailsOptions()...)
}
