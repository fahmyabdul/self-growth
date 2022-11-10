package transactions

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/fahmyabdul/self-growth/fetch-app/internal/common"
	"github.com/fahmyabdul/self-growth/fetch-app/internal/models/exchangerates"
	"github.com/fahmyabdul/self-growth/fetch-app/internal/requests"
)

func (p *Transactions) GetAll() (*[]Transactions, error) {
	var requestDataApp requests.DataApp

	responseBody, err := requestDataApp.GetAll()

	// Unmarshal response json body to struct models
	var responseStruct RestTransactionResponse
	err = json.Unmarshal(responseBody, &responseStruct)
	if err != nil {
		return nil, err
	}

	var modifiedListTransactions []Transactions
	for _, transaction := range responseStruct.Data {
		// Ignore empty Transaction Hash
		if transaction.TransactionHash == "" {
			continue
		}

		paid := strconv.Itoa(transaction.PaidAmount)

		priceIDR, err := strconv.ParseFloat(paid[:len(paid)-1], 64)
		if err != nil {
			continue
		}

		// Set Price USD
		transaction.PaidUSD = fmt.Sprintf("%.2f", (priceIDR * exchangerates.GlobExchangeRatesData.Rates))

		// Append current data to modifiedListTransactions
		modifiedListTransactions = append(modifiedListTransactions, transaction)
	}

	return &modifiedListTransactions, nil
}

func (p *Transactions) GetAggregate() (map[string]map[int]TransactionsAggregateContent, error) {
	var requestDataApp requests.DataApp

	responseBody, err := requestDataApp.GetAll()

	// Unmarshal response json body to struct models
	var responseStruct RestTransactionResponse
	err = json.Unmarshal(responseBody, &responseStruct)
	if err != nil {
		return nil, err
	}

	groupTransactions := make(map[string]map[int]TransactionsAggregateContent)
	tempPriceCollection := make(map[string]map[int][]int)

	for _, transaction := range responseStruct.Data {
		// Ignore empty Transaction Hash
		if transaction.TransactionHash == "" {
			continue
		}

		groupBy := transaction.Username

		// Check whether current groupBy map group(the parent map) exists or not
		if _, ok := groupTransactions[groupBy]; !ok {
			// If not, then create the groupBy map group
			groupTransactions[groupBy] = make(map[int]TransactionsAggregateContent)
			tempPriceCollection[groupBy] = make(map[int][]int)
		}
		// Check whether current groupBy => payment_method map group(the groupBy child map) exists or not
		if _, ok := groupTransactions[groupBy][transaction.PaymentMethod]; !ok {
			// Create the temporary PaidA map group that will be used for Median & Average calculation
			tempPriceCollection[groupBy][transaction.PaymentMethod] = []int{transaction.PaidAmount}

			// If not, then create the groupBy => payment_method map group
			groupTransactions[groupBy][transaction.PaymentMethod] = TransactionsAggregateContent{
				PaymentMethod: transaction.PaymentMethod,
				Paid: AggregateGroup{
					Collection: tempPriceCollection[groupBy][transaction.PaymentMethod],
					Min:        transaction.PaidAmount,
					Max:        transaction.PaidAmount,
					Median:     transaction.PaidAmount,
					Avg:        float64(transaction.PaidAmount),
				},
			}

			continue
		}

		// Append current Paid to temporary map group
		tempPriceCollection[groupBy][transaction.PaymentMethod] = append(tempPriceCollection[groupBy][transaction.PaymentMethod], transaction.PaidAmount)
		// Sort the temporary map group
		sort.Ints(tempPriceCollection[groupBy][transaction.PaymentMethod])

		// Get the Price Median & Average in every loop, so we don't have to recalcuate outside of current loop
		usedPriceMedian := common.GetMedian(tempPriceCollection[groupBy][transaction.PaymentMethod])
		usedPriceAvg := common.GetAverage(tempPriceCollection[groupBy][transaction.PaymentMethod])

		// Update the content of current map group with the values from above calculation
		groupTransactions[groupBy][transaction.PaymentMethod] = TransactionsAggregateContent{
			PaymentMethod: transaction.PaymentMethod,
			Paid: AggregateGroup{
				Collection: tempPriceCollection[groupBy][transaction.PaymentMethod],
				Min:        tempPriceCollection[groupBy][transaction.PaymentMethod][0],
				Max:        tempPriceCollection[groupBy][transaction.PaymentMethod][len(tempPriceCollection[groupBy][transaction.PaymentMethod])-1],
				Median:     usedPriceMedian,
				Avg:        usedPriceAvg,
			},
		}
	}

	return groupTransactions, nil
}
