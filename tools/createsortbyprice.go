package main

import (
	"context"
	"flag"
	"fmt"

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

	fmt.Println(`
	<table>
		<thead>
			<tr>
				<th>opentable URI</th>
				<th>URI</th>
				<th>Menu Item</th>
				<th>Price</th>
			</tr>
		</thead>
		<tbody>
`)
	td := func(s string) {
		fmt.Printf("<td>%s</td>\n", s)
	}
	commas := message.NewPrinter(language.AmericanEnglish)
	for _, s := range ss {
		fmt.Println("<tr>")
		td(fmt.Sprintf("[%s](%s)", s.RestaurantName, s.OpentableLink))
		td(fmt.Sprintf("[Web](%s)", s.RestaurantWebsite))
		td(s.Title)
		td(commas.Sprintf("%d", int(s.Price)))
		fmt.Println("</tr>")
	}
	fmt.Println(`
	<tbody>
	</table>`)

	return nil
}

func main() {
	flag.Parse()
	check.Err(createSortByPrice(context.Background()))
}
