package api

import (
	"context"
	"fmt"
)

//go:generate genopts --function SaveRawRestaurantDetailsFromID --params --required "restID string" --extends RawRestaurantDetailsFromID,SaveRestaurant
func (e *Extended) SaveRawRestaurantDetailsFromID(ctx context.Context, restID string, optss ...SaveRawRestaurantDetailsFromIDOption) (*RawRestaurantDetails, error) {
	opts := MakeSaveRawRestaurantDetailsFromIDOptions(optss...)
	info, err := e.RawRestaurantDetailsFromID(restID, opts.ToRawRestaurantDetailsFromIDOptions()...)
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("https://www.opentable.com/r/%s", restID)
	if err := e.Cache().SaveRestaurant(ctx, uri, *info, opts.ToSaveRestaurantOptions()...); err != nil {
		return nil, err
	}
	return info, nil
}
