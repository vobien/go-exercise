package prices

import (
	"fmt"
	"tax-calculator/conversion"
	"tax-calculator/filemanager"
)

type TaxIncludedPrice struct {
	Tax              float64
	InputPrices      []float64
	PriceIncludedTax []string
}

func (t *TaxIncludedPrice) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	t.InputPrices, err = conversion.StringToFloat(lines)
	if err != nil {
		fmt.Println("Error converting to float:", err)
		return
	}
}

func (t *TaxIncludedPrice) Calculate() {
	t.LoadData()

	t.PriceIncludedTax = make([]string, len(t.InputPrices))
	for i, price := range t.InputPrices {
		t.PriceIncludedTax[i] = fmt.Sprintf("%.2f",price * (1 + t.Tax))
	}
}

func New(tax float64) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		Tax:         tax,
	}
}
