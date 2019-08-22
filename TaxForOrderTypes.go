package taxjar

// NexusAddress - TODO (document this)
type NexusAddress struct {
	ID      string `json:"id,omitempty"`
	Country string `json:"country"`
	Zip     string `json:"zip,omitempty"`
	State   string `json:"state"`
	City    string `json:"city,omitempty"`
	Street  string `json:"street,omitempty"`
}

// TaxLineItem - TODO (document this)
type TaxLineItem struct {
	ID             string  `json:"id,omitempty"`
	Quantity       int     `json:"quantity,omitempty"`
	ProductTaxCode string  `json:"product_tax_code,omitempty"`
	UnitPrice      float64 `json:"unit_price,omitempty"`
	Discount       float64 `json:"discount,omitempty"`
}

// TaxForOrderParams - TODO (document this)
type TaxForOrderParams struct {
	FromCountry    string         `json:"from_country,omitempty"`
	FromZip        string         `json:"from_zip,omitempty"`
	FromState      string         `json:"from_state,omitempty"`
	FromCity       string         `json:"from_city,omitempty"`
	FromStreet     string         `json:"from_street,omitempty"`
	ToCountry      string         `json:"to_country"`
	ToZip          string         `json:"to_zip,omitempty"`
	ToState        string         `json:"to_state,omitempty"`
	ToCity         string         `json:"to_city,omitempty"`
	ToStreet       string         `json:"to_street,omitempty"`
	Amount         float64        `json:"amount,omitempty"`
	Shipping       float64        `json:"shipping"`
	CustomerID     string         `json:"customer_id,omitempty"`
	ExemptionType  string         `json:"exemption_type"`
	NexusAddresses []NexusAddress `json:"nexus_addresses,omitempty"`
	LineItems      []TaxLineItem  `json:"line_items,omitempty"`
}

// Jurisdictions - TODO (document this)
type Jurisdictions struct {
	Country string `json:"country"`
	State   string `json:"state,omitempty"`
	County  string `json:"county,omitempty"`
	City    string `json:"city,omitempty"`
}

// Shipping - TODO (document this)
type Shipping struct {
	TaxableAmount         float64 `json:"taxable_amount,omitempty"`
	TaxCollectable        float64 `json:"tax_collectable,omitempty"`
	CombinedTaxRate       float64 `json:"combined_tax_rate,omitempty"`
	StateTaxableAmount    float64 `json:"state_taxable_amount,omitempty"`
	StateSalesTaxRate     float64 `json:"state_sales_tax_rate,omitempty"`
	StateAmount           float64 `json:"state_amount,omitempty"`
	CountyTaxableAmount   float64 `json:"county_taxable_amount,omitempty"`
	CityTaxRate           float64 `json:"city_tax_rate,omitempty"`
	CityAmount            float64 `json:"city_amount,omitempty"`
	SpecialTaxableAmount  float64 `json:"special_taxable_amount,omitempty"`
	SpecialTaxRate        float64 `json:"special_tax_rate,omitempty"`
	SpecialDistrictAmount float64 `json:"special_district_amount,omitempty"`
}

// LineItemBreakdown - TODO (document this)
type LineItemBreakdown struct {
	ID                           string  `json:"id,omitempty"`
	TaxableAmount                float64 `json:"taxable_amount,omitempty"`
	TaxCollectable               float64 `json:"tax_collectable,omitempty"`
	CombinedTaxRate              float64 `json:"combined_tax_rate,omitempty"`
	StateTaxableAmount           float64 `json:"state_taxable_amount,omitempty"`
	StateSalesTaxRate            float64 `json:"state_sales_tax_rate,omitempty"`
	StateAmount                  float64 `json:"state_amount,omitempty"`
	CountyTaxableAmount          float64 `json:"county_taxable_amount,omitempty"`
	CityTaxRate                  float64 `json:"city_tax_rate,omitempty"`
	CityAmount                   float64 `json:"city_amount,omitempty"`
	SpecialDistrictTaxableAmount float64 `json:"special_district_taxable_amount,omitempty"`
	SpecialTaxRate               float64 `json:"special_tax_rate,omitempty"`
	SpecialDistrictAmount        float64 `json:"special_district_amount,omitempty"`
}

// Breakdown - TODO (document this)
type Breakdown struct {
	TaxableAmount                 float64             `json:"taxable_amount,omitempty"`
	TaxCollectable                float64             `json:"tax_collectable,omitempty"`
	CombinedTaxRate               float64             `json:"combined_tax_rate,omitempty"`
	StateTaxableAmount            float64             `json:"state_taxable_amount,omitempty"`
	StateTaxRate                  float64             `json:"state_tax_rate,omitempty"`
	StateTaxCollectable           float64             `json:"state_tax_collectable,omitempty"`
	CountyTaxableAmount           float64             `json:"county_taxable_amount,omitempty"`
	CountyTaxRate                 float64             `json:"county_tax_rate,omitempty"`
	CountyTaxCollectable          float64             `json:"county_tax_collectable,omitempty"`
	CityTaxableAmount             float64             `json:"city_taxable_amount,omitempty"`
	CityTaxRate                   float64             `json:"city_tax_rate,omitempty"`
	CityTaxCollectable            float64             `json:"city_tax_collectable,omitempty"`
	SpecialDistrictTaxableAmount  float64             `json:"special_district_taxable_amount,omitempty"`
	SpecialTaxRate                float64             `json:"special_tax_rate,omitempty"`
	SpecialDistrictTaxCollectable float64             `json:"special_district_tax_collectable,omitempty"`
	Shipping                      Shipping            `json:"shipping,omitempty"`
	LineItems                     []LineItemBreakdown `json:"line_items,omitempty"`
}

// TaxForOrderResponse - TODO (document this)
type TaxForOrderResponse struct {
	Tax struct {
		OrderTotalAmount float64       `json:"order_total_amount"`
		Shipping         float64       `json:"shipping"`
		TaxableAmount    float64       `json:"taxable_amount"`
		AmountToCollect  float64       `json:"amount_to_collect"`
		Rate             float64       `json:"rate"`
		HasNexus         bool          `json:"has_nexus"`
		FreightTaxable   bool          `json:"freight_taxable"`
		TaxSource        string        `json:"tax_source"`
		ExemptionType    string        `json:"exemption_type"`
		Jurisdictions    Jurisdictions `json:"jurisdictions"`
		Breakdown        Breakdown     `json:"breakdown"`
	} `json:"tax"`
}
