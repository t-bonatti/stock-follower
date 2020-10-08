package main

import (
	"github.com/apex/log"
	_ "github.com/lib/pq"
	"github.com/t-bonatti/stock-follower/config"
	"github.com/t-bonatti/stock-follower/database"
	statusinvest "github.com/t-bonatti/stock-follower/statusinvest"
	"github.com/t-bonatti/stock-follower/stock"

	"time"
)

func main() {
	var cfg = config.Get()
	var now = time.Now()

	var db = database.Connect(cfg.DatabaseURL)
	defer func() {
		if err := db.Close(); err != nil {
			log.WithError(err).Error("failed to close database connections")
		}
	}()

	repo := statusinvest.NewStatusInvestRepository(db)
	craweler := statusinvest.NewCrawler()

	for _, stock := range stock.GetAll() {
		statusInvestData, err := craweler.Get(stock)
		statusInvestData.CrawledAt = now
		if err == nil {
			err2 := repo.Create(statusInvestData)
			if err2 != nil {
				log.WithError(err2).Error("failed to create")
			}
		}
	}

}
