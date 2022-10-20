package komoditas

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fahmyabdul/efishery-task/fetch-app/internal/common"
	"github.com/fahmyabdul/efishery-task/fetch-app/internal/models/exchangerates"
	"github.com/fahmyabdul/efishery-task/fetch-app/internal/requests"
)

func (p *Komoditas) GetAll() (*[]Komoditas, error) {
	var requestEfisheryApi requests.EfisheryApi

	responseBody, err := requestEfisheryApi.GetKomoditas()

	// Unmarshal response json body to struct models
	var listKomoditas []Komoditas
	err = json.Unmarshal(responseBody, &listKomoditas)
	if err != nil {
		return nil, err
	}

	var modifiedListKomoditas []Komoditas
	for _, komoditas := range listKomoditas {
		// Ignore empty UUID
		if komoditas.Uuid == "" {
			continue
		}

		// Skip add usd_price if price is empty
		if komoditas.Price == "" {
			modifiedListKomoditas = append(modifiedListKomoditas, komoditas)
			continue
		}

		priceIDR, err := strconv.ParseFloat(komoditas.Price[:len(komoditas.Price)-1], 64)
		if err != nil {
			continue
		}

		// Set Price USD
		komoditas.PriceUSD = fmt.Sprintf("%.2f", (priceIDR * exchangerates.GlobExchangeRatesData.Rates))

		// Append current data to modifiedListKomoditas
		modifiedListKomoditas = append(modifiedListKomoditas, komoditas)
	}

	return &modifiedListKomoditas, nil
}

func (p *Komoditas) GetAggregate() (map[string]map[string]KomoditasAggregateContent, error) {
	var requestEfisheryApi requests.EfisheryApi

	responseBody, err := requestEfisheryApi.GetKomoditas()

	// Unmarshal response json body to struct models
	var listKomoditas []Komoditas
	err = json.Unmarshal(responseBody, &listKomoditas)
	if err != nil {
		return nil, err
	}

	groupKomoditas := make(map[string]map[string]KomoditasAggregateContent)
	tempPriceCollection := make(map[string]map[string][]int)
	tempSizeCollection := make(map[string]map[string][]int)

	for _, komoditas := range listKomoditas {
		// Ignore empty UUID
		if komoditas.Uuid == "" {
			continue
		}

		areaProvinsi := strings.ToUpper(strings.Trim(komoditas.AreaProvinsi, " "))
		weekNumber, err := common.GetWeekNumber(komoditas.Timestamp)
		if err != nil {
			continue
		}
		// Convert Weeknumber to string
		weekNumberString := fmt.Sprintf("%d", weekNumber)

		// Convert Price & Size to Integer
		price, _ := strconv.Atoi(komoditas.Price)
		size, _ := strconv.Atoi(komoditas.Size)

		// Check whether current areaprovinsi map group(the parent map) exists or not
		if _, ok := groupKomoditas[areaProvinsi]; !ok {
			// If not, then create the areaprovinsi map group
			groupKomoditas[areaProvinsi] = make(map[string]KomoditasAggregateContent)
			tempPriceCollection[areaProvinsi] = make(map[string][]int)
			tempSizeCollection[areaProvinsi] = make(map[string][]int)
		}
		// Check whether current areaprovinsi => week number map group(the areaprovinsi child map) exists or not
		if _, ok := groupKomoditas[areaProvinsi][weekNumberString]; !ok {
			// Create the temporary Price & Size map group that will be used for Median & Average calculation
			tempPriceCollection[areaProvinsi][weekNumberString] = []int{price}
			tempSizeCollection[areaProvinsi][weekNumberString] = []int{size}

			// If not, then create the areaprovinsi => week number map group
			groupKomoditas[areaProvinsi][weekNumberString] = KomoditasAggregateContent{
				WeekNumber: weekNumber,
				Price: AggregateGroup{
					Collection: tempPriceCollection[areaProvinsi][weekNumberString],
					Min:        price,
					Max:        price,
					Median:     price,
					Avg:        float64(price),
				},
				Size: AggregateGroup{
					Collection: tempSizeCollection[areaProvinsi][weekNumberString],
					Min:        size,
					Max:        size,
					Median:     size,
					Avg:        float64(size),
				},
			}

			continue
		}

		// Append current Price & Size to temporary map group
		tempPriceCollection[areaProvinsi][weekNumberString] = append(tempPriceCollection[areaProvinsi][weekNumberString], price)
		tempSizeCollection[areaProvinsi][weekNumberString] = append(tempSizeCollection[areaProvinsi][weekNumberString], size)
		// Sort the temporary map group
		sort.Ints(tempPriceCollection[areaProvinsi][weekNumberString])
		sort.Ints(tempSizeCollection[areaProvinsi][weekNumberString])

		// Get the Price Median & Average in every loop, so we don't have to recalcuate outside of current loop
		usedPriceMedian := common.GetMedian(tempPriceCollection[areaProvinsi][weekNumberString])
		usedPriceAvg := common.GetAverage(tempPriceCollection[areaProvinsi][weekNumberString])

		// Get the Size Median & Average in every loop, so we don't have to recalcuate outside of current loop
		usedSizeMedian := common.GetMedian(tempSizeCollection[areaProvinsi][weekNumberString])
		usedSizeAvg := common.GetAverage(tempSizeCollection[areaProvinsi][weekNumberString])

		// Update the content of current map group with the values from above calculation
		groupKomoditas[areaProvinsi][weekNumberString] = KomoditasAggregateContent{
			WeekNumber: weekNumber,
			Price: AggregateGroup{
				Collection: tempPriceCollection[areaProvinsi][weekNumberString],
				Min:        tempPriceCollection[areaProvinsi][weekNumberString][0],
				Max:        tempPriceCollection[areaProvinsi][weekNumberString][len(tempPriceCollection[areaProvinsi][weekNumberString])-1],
				Median:     usedPriceMedian,
				Avg:        usedPriceAvg,
			},
			Size: AggregateGroup{
				Collection: tempSizeCollection[areaProvinsi][weekNumberString],
				Min:        tempSizeCollection[areaProvinsi][weekNumberString][0],
				Max:        tempSizeCollection[areaProvinsi][weekNumberString][len(tempSizeCollection[areaProvinsi][weekNumberString])-1],
				Median:     usedSizeMedian,
				Avg:        usedSizeAvg,
			},
		}
	}

	return groupKomoditas, nil
}
