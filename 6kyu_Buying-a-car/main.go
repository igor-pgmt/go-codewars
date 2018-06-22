// https://www.codewars.com/kata/buying-a-car/train/go
// He thinks he can save $1000 each month but the prices of his old car and of the new one decrease of 1.5 percent per month. Furthermore the percent of loss increases by a fixed 0.5 percent at the end of every two months.
// Can you help him? Our man finds it difficult to make all these calculations.
// How many months will it take him to save up enough money to buy the car he wants, and how much money will he have left over?
// Parameters and return of function:
// parameter (positive int, guaranteed) startPriceOld (Old car price)
// parameter (positive int, guaranteed) startPriceNew (New car price)
// parameter (positive int, guaranteed) savingperMonth
// parameter (positive float or int, guaranteed) percentLossByMonth
// nbMonths(2000, 8000, 1000, 1.5) should return [6, 766] or (6, 766)
// where 6 is the number of months at the end of which he can buy the new car and 766 is the nearest integer to '766.158...' .
// Note: Selling, buying and saving are normally done at end of month. Calculations are processed at the end of each considered month but if, by chance from the start, the value of the old car is bigger than the value of the new one or equal there is no saving to be made, no need to wait so he can at the beginning of the month buy the new car:
// nbMonths(12000, 8000, 1000, 1.5) should return [0, 4000]
// nbMonths(8000, 8000, 1000, 1.5) should return [0, 0]
// We don't take care of a deposit of savings in a bank:-)

package main

import "fmt"

func main() {
	fmt.Println(NbMonths(2000, 8000, 1000, 1.5))
}

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {

	months := 0
	priceOld, priceNew := float64(startPriceOld), float64(startPriceNew)

	for ; priceOld+float64(months*savingperMonth) < priceNew; months++ {
		if months%2 == 1 {
			percentLossByMonth += 0.5
		}
		priceOld -= priceOld * percentLossByMonth / 100.0
		priceNew -= priceNew * percentLossByMonth / 100.0
	}

	return [2]int{months, int(priceOld + float64(months*savingperMonth) - priceNew + 0.5)}
}
