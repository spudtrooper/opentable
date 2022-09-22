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

func createMenuItemHistogram(ctx context.Context) error {
	db, err := api.ConnectToDB(ctx)
	if err != nil {
		return err
	}

	cur, err := db.Collection("menuItemHistogram").Find(ctx, bson.D{})
	if err != nil {
		return err
	}

	type stored struct {
		MenuItem string
		Count    int
	}

	const limit = 1000
	var ss []stored
	for i := 0; i < limit && cur.Next(ctx); i++ {
		var el stored
		if err := cur.Decode(&el); err != nil {
			return err
		}
		ss = append(ss, el)
	}

	fmt.Println(`
# OpenTable - Menu Item Histogram

A histogram of New York City opentable menu items.
`)

	var rows [][]string
	commas := message.NewPrinter(language.AmericanEnglish)
	for _, s := range ss {
		row := []string{
			strings.ReplaceAll(s.MenuItem, "|", " "),
			commas.Sprintf("%d", int(s.Count)),
		}
		rows = append(rows, row)
	}

	tab, err := markdown.
		NewTableFormatterBuilder().
		Build("Menu Item", "Count").
		Format(rows)
	if err != nil {
		return err
	}
	fmt.Println(tab)

	return nil
}

func main() {
	flag.Parse()
	check.Err(createMenuItemHistogram(context.Background()))
}
