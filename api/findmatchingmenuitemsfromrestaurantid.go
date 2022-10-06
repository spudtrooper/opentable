package api

//go:generate genopts --function FindMatchingMenuItemsFromRestaurantID --params --required "restID string,term string" --extends RestaurantDetailsFromID
func (e *Extended) FindMatchingMenuItemsFromRestaurantID(restID, term string, optss ...FindMatchingMenuItemsFromRestaurantIDOption) (*FindMatchingMenuItemsInfo, error) {
	opts := MakeFindMatchingMenuItemsFromRestaurantIDOptions(optss...)

	info, err := e.RestaurantDetailsFromID(restID, opts.ToRestaurantDetailsFromIDOptions()...)
	if err != nil {
		return nil, err
	}
	items := FindMatchingMenuItems(info.RestaurantDetails, term)
	return &items, nil
}
