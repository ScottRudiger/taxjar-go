package taxjar

import "encoding/json"

// ShowOrderParams - TODO (document this)
type ShowOrderParams struct {
	Provider string `url:"provider"`
}

// ShowOrderResponse - TODO (document this)
type ShowOrderResponse struct {
	Order struct {
		TransactionID   string          `json:"transaction_id"`
		UserID          int             `json:"user_id"`
		TransactionDate string          `json:"transaction_date"`
		Provider        string          `json:"provider"`
		ExemptionType   string          `json:"exemption_type"`
		FromCountry     string          `json:"from_country"`
		FromZip         string          `json:"from_zip"`
		FromState       string          `json:"from_state"`
		FromCity        string          `json:"from_city"`
		FromStreet      string          `json:"from_street"`
		ToCountry       string          `json:"to_country"`
		ToZip           string          `json:"to_zip"`
		ToState         string          `json:"to_state"`
		ToCity          string          `json:"to_city"`
		ToStreet        string          `json:"to_street"`
		Amount          float64         `json:"amount"`
		Shipping        float64         `json:"shipping"`
		SalesTax        float64         `json:"sales_tax"`
		LineItems       []OrderLineItem `json:"line_items"`
	} `json:"order"`
}

// OrderLineItem - TODO (document this)
type OrderLineItem struct {
	ID                string  `json:"id"`
	Quantity          int     `json:"quantity"`
	ProductIdentifier string  `json:"product_identifier,omitempty"`
	Description       string  `json:"description,omitempty"`
	ProductTaxCode    string  `json:"product_tax_code"`
	UnitPrice         float64 `json:"unit_price"`
	Discount          float64 `json:"discount,omitempty"`
	SalesTax          float64 `json:"sales_tax"`
}

// ShowOrder - TODO (document this)
func (client *Config) ShowOrder(transactionID string, params ...ShowOrderParams) (*ShowOrderResponse, error) {
	res, err := client.get("transactions/orders/"+transactionID, params)
	if err != nil {
		return nil, err
	}
	order := new(ShowOrderResponse)
	json.Unmarshal(res.([]byte), &order)
	return order, nil
}
