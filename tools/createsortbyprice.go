package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/opentable/api"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func createSortByPrice(ctx context.Context) error {
	db, err := api.ConnectToDB(ctx)
	if err != nil {
		return err
	}

	cur, err := db.Collection("sortedByPrice").Find(ctx, bson.D{})
	if err != nil {
		return err
	}

	type storedSorteResult struct {
		RestaurantName    string
		RestaurantWebsite string
		Title             string
		Price             float32
		OpentableLink     string
	}

	const limit = 1000
	var ss []storedSorteResult
	for i := 0; i < limit && cur.Next(ctx); i++ {
		var el storedSorteResult
		if err := cur.Decode(&el); err != nil {
			return err
		}
		if el.Price == 0 {
			break
		}
		ss = append(ss, el)
	}

	fmt.Println(`
# OpenTable - Sorted by Price

Menu items from opentable sorted by price from [github.com/spudtrooper/opentable](https://github.com/spudtrooper/opentable)
`)

	var rows [][]string
	commas := message.NewPrinter(language.AmericanEnglish)
	for _, s := range ss {
		row := []string{
			fmt.Sprintf("[%s](%s)", strings.ReplaceAll(s.RestaurantName, "|", " "), s.OpentableLink),
			fmt.Sprintf("[Web](%s)", s.RestaurantWebsite),
			strings.ReplaceAll(s.Title, "|", " "),
			commas.Sprintf("%d", int(s.Price)),
		}
		rows = append(rows, row)
	}

	tab, err := markdown.
		NewTableFormatterBuilder().
		Build("opentable URI", "URI", "Menu Item", "Price").
		Format(rows)
	if err != nil {
		return err
	}
	fmt.Println(tab)

	return nil
}

func main() {
	flag.Parse()
	check.Err(createSortByPrice(context.Background()))
}
