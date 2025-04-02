package schedule

import (
	"github.com/robfig/cron"
)

const QuantitativeStocks = "Quantitative|Stocks"

func InitSchedule() {
	var cronInstance = cron.New()
	cronInstance.Start()
}

