package common

import (
	"strconv"
	"strings"
	"time"

	"github.com/fahmyabdul/efishery-task/fetch-app/internal/models/exchangerates"
	"github.com/fahmyabdul/efishery-task/fetch-app/internal/requests"
	"github.com/fahmyabdul/golibs"
)

// CheckExchangeRates: Check Exchange Rates, fetch exchange rates from api, create/update exchange rates table
func CheckExchangeRates(sourceCaller interface{}) error {
	fetchTime := time.Now().Unix()

	// Check Exchange Rates table existence
	var tableExchangeRates exchangerates.ExchangeRates
	if err := tableExchangeRates.CheckTableExists(); err != nil {
		// If table not exists then Generate table
		err = tableExchangeRates.GenerateTable()
		if err != nil {
			golibs.Log.Printf("| %s | Check Exchange Rates | Generate Table | Failed, error: %s\n", sourceCaller, err.Error())
			return err
		}
		golibs.Log.Printf("| %s | Check Exchange Rates | Generate Table | Success", sourceCaller)
	} else {
		// Get Exchange Rates From Table
		err := tableExchangeRates.GetExchangeRates("IDR", "USD")
		if err != nil {
			golibs.Log.Printf("| %s | Check Exchange Rates | Get From Table | Failed, error: %s\n", sourceCaller, err.Error())
			return err
		}

		// Set GlobalExchangeRatesData
		exchangerates.GlobExchangeRatesData = tableExchangeRates
	}

	switch sourceCaller.(string) {
	case "InitBefore":
		if tableExchangeRates.ID != "" {
			// If Exchange Rates exists then do nothing
			break
		}

		// If Exchange Rates not exists then...

		// Fetch Exchange Rates from CurrencyConverterAPI
		var requestCurrency requests.CurrencyConverterApi
		exchangeRatesData, err := requestCurrency.GetExchangeRates("IDR", "USD")
		if err != nil {
			golibs.Log.Printf("| %s | Check Exchange Rates | Get | From: IDR, To: USD | Failed, error: %s\n", sourceCaller, err.Error())
			return err
		}

		// Create Exchange Rates
		lastInsertId, err := tableExchangeRates.Create(exchangeRatesData.Base, "USD", exchangeRatesData.Rates["USD"], exchangeRatesData.LimitMonth, fetchTime)
		if err != nil {
			golibs.Log.Printf("| %s | Check Exchange Rates | Create | From: %s, To: %s | Failed, error: %s\n", sourceCaller, exchangeRatesData.Base, "USD", err.Error())
			return err
		}

		// Set GlobExchangeRatesData
		exchangerates.GlobExchangeRatesData = exchangerates.ExchangeRates{
			ID:        strconv.Itoa(lastInsertId),
			From:      "IDR",
			To:        "USD",
			Rates:     exchangeRatesData.Rates["USD"],
			Limit:     exchangeRatesData.LimitMonth,
			LastFetch: fetchTime,
		}

		golibs.Log.Printf("| %s | Check Exchange Rates | Create | From: %s, To: %s | Success\n", sourceCaller, exchangeRatesData.Base, "USD")
	case "CronJob":
		// Fetch Exchange Rates from CurrencyConverterAPI
		var requestCurrency requests.CurrencyConverterApi
		exchangeRatesData, err := requestCurrency.GetExchangeRates("IDR", "USD")
		if err != nil {
			golibs.Log.Printf("| %s | Check Exchange Rates | Get | From: IDR, To: USD | Failed, error: %s\n", sourceCaller, err.Error())
			return err
		}

		lastInsertId := 0
		if tableExchangeRates.ID == "" {
			// If Exchange Rates not exists
			// Create Exchange Rates
			lastInsertId, err = tableExchangeRates.Create(exchangeRatesData.Base, "USD", exchangeRatesData.Rates["USD"], exchangeRatesData.LimitMonth, fetchTime)
			if err != nil {
				golibs.Log.Printf("| %s | Check Exchange Rates | Create | From: %s, To: %s | Failed, error: %s\n", sourceCaller, exchangeRatesData.Base, "USD", err.Error())
				return err
			}

			golibs.Log.Printf("| %s | Check Exchange Rates | Create | From: %s, To: %s | Success\n", sourceCaller, exchangeRatesData.Base, "USD")
		} else {
			// If Exchange Rates exists
			// Update Exchange Rates
			lastInsertId, err = tableExchangeRates.Update(exchangeRatesData.Base, "USD", exchangeRatesData.Rates["USD"], exchangeRatesData.LimitMonth, fetchTime)
			if err != nil {
				golibs.Log.Printf("| %s | Check Exchange Rates | Update | From: %s, To: %s | Failed, error: %s\n", sourceCaller, exchangeRatesData.Base, "USD", err.Error())
				return err
			}

			golibs.Log.Printf("| %s | Check Exchange Rates | Update | From: %s, To: %s | Success\n", sourceCaller, exchangeRatesData.Base, "USD")
		}

		// Set GlobExchangeRatesData
		exchangerates.GlobExchangeRatesData = exchangerates.ExchangeRates{
			ID:        strconv.Itoa(lastInsertId),
			From:      "IDR",
			To:        "USD",
			Rates:     exchangeRatesData.Rates["USD"],
			Limit:     exchangeRatesData.LimitMonth,
			LastFetch: fetchTime,
		}
	}

	golibs.Log.Printf("| %s | Check Exchange Rates | Done | Current Limit: %d", sourceCaller, exchangerates.GlobExchangeRatesData.Limit)

	return nil
}

func GetWeekNumber(inputTime string) (int, error) {
	inputTime = strings.ReplaceAll(inputTime, ".0", "")
	unixTime, err := strconv.ParseInt(inputTime, 10, 64)
	if err != nil {
		return 0, err
	}

	formattedTime := time.Unix(unixTime, 0)

	_, week := formattedTime.ISOWeek()

	return week, nil
}

func GetMedian(arrInts []int) int {
	intsLength := len(arrInts)

	if intsLength == 0 {
		return 0
	} else if intsLength%2 == 0 {
		start := (intsLength / 2) - 1
		end := intsLength / 2

		return (arrInts[start] + arrInts[end]) / 2
	}

	return arrInts[intsLength/2]
}

func GetAverage(arrInts []int) float64 {
	intsLength := len(arrInts)
	arrTotal := 0
	for _, value := range arrInts {
		arrTotal += value
	}

	return float64(arrTotal) / float64(intsLength)
}
