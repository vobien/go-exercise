package prices

import (
	"fmt"
	"tax-calculator/conversion"
	"tax-calculator/io_manager"
)

type TaxIncludedPrice struct {
	IOManager        io_manager.IOManager `json:"-"`
	Tax              float64              `json:"tax"`
	InputPrices      []float64            `json:"prices"`
	PriceIncludedTax []string             `json:"price_included_tax"`
}

func (t *TaxIncludedPrice) LoadData() error {
	lines, err := t.IOManager.ReadLines()
	if err != nil {
		return err
	}

	t.InputPrices, err = conversion.StringToFloat(lines)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaxIncludedPrice) Process() error {
	err := t.LoadData()
	if err != nil {
		return err
	}

	t.PriceIncludedTax = make([]string, len(t.InputPrices))
	for i, price := range t.InputPrices {
		t.PriceIncludedTax[i] = fmt.Sprintf("%.2f", price*(1+t.Tax))
	}
	
	return t.IOManager.Write(t)
}

func New(tax float64, ioManager io_manager.IOManager) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		IOManager: ioManager,
		Tax:       tax,
	}
}
