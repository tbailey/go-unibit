package client

type RealTimePriceResult map[string][]RealtimePrice
type RealtimePrice struct {
	Ticker    string `json:"ticker"`
	Date      string `json:"date"`
	Minute    string `json:"minute"`
	Price     float64 `json:"price"`
	Volume    uint64 `json:"volume"`
	Timezone  string `json:"timezone"`
	Timestamp uint64 `json:"timestamp"`
	UtcOffset uint64 `json:"utc_offset"`
}

type SecFilingLinkResult map[string][]SecFilingLink
type SecFilingLink struct {
	Date string `json:"date"`
	Cik uint64 `json:"cik"`
	DocUrl string `json:"docurl"`
	HtmlUrl string `json:"htmlurl"`
	Type string `json:"type"`
	TxtUrl string `json:"txturl"`
	Timestamp uint64 `json:"timestamp"`
}

type StockNewsResult map[string][]StockNews
type StockNews struct {
	Sentiment float64 `json:"sentiment"`
	Ticker string `json:"ticker"`
	Author string `json:"author"`
	EventLocation []string `json:"event_location"`
	EventPerson []string `json:"event_person"`
	Description string `json:"description"`
	EventOrganization []string `json:"event_organization"`
	Source string `json:"source"`
	Published string `json:"published"`
	Title string `json:"title"`
	Url string `json:"url"`
	EventSector []string `json:"event_sector"`
	EventGenre []string `json:"event_genre"`
}

type CorporateSplitsResult []CorporateSplits
type CorporateSplits struct {
	BeforeSplit string `json:"before_split"`
	TickerSymbol string `json:"ticker_symbol"`
	CompanyName string `json:"company_name"`
	AfterSplit string `json:"after_split"`
	PayableOn string `json:"payable_on"`
	Optionable string `json:"optionable"`
	Timestamp uint64 `json:"timestamp"`
}

type CorporateDividendsResult []CorporateDividends
type CorporateDividends struct {
	Country string `json:"country"`
	PaymentTimestamp uint64 `json:"payment_timestamp"`
	TickerSymbol string `json:"ticker_symbol"`
	ExDividendDate string `json:"ex_dividend_date"`
	CompanyName string `json:"company_name"`
	YieldPercent float64 `json:"yield_percent"`
	Dividend float64 `json:"dividend"`
	Type string `json:"type"`
	ExDividendTimestamp uint64 `json:"ex_dividend_timestamp"`
	PaymentDate string `json:"payment_date"`
}

type HistoricalPriceResult map[string][]HistoricalPrice
type HistoricalPrice struct {
	Date     string  `json:"date"`
	Volume   uint64  `json:"volume"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	AdjClose float64 `json:"adj_close"`
	Close    float64 `json:"close"`
	Open     float64 `json:"open"`
}

type CompanyFinancialsResult map[string][]CompanyFinancials
type CompanyFinancials map[string]string

type CompanyProfileResult map[string]CompanyProfile
type CompanyProfile struct {
	Ticker              string `json:"ticker"`
	CompanyName         string `json:"company_name"`
	Address             string `json:"address"`
	Phone               string `json:"phone"`
	Website             string `json:"website"`
	Sector              string `json:"sector"`
	Industry            string `json:"industry"`
	EmployeeNumber      string `json:"employee_number"`
	Executive1Name      string `json:"executive1_name"`
	Executive1Title     string `json:"executive1_title"`
	Executive1Pay       string `json:"executive1_pay"`
	Executive1Exercised string `json:"executive1_exercised"`
	Executive1Yearborn  string `json:"executive1_yearborn"`
	Executive2Name      string `json:"executive2_name"`
	Executive2Title     string `json:"executive2_title"`
	Executive2Pay       string `json:"executive2_pay"`
	Executive2Exercised string `json:"executive2_exercised"`
	Executive2Yearborn  string `json:"executive2_yearborn"`
	Executive3Name      string `json:"executive3_name"`
	Executive3Title     string `json:"executive3_title"`
	Executive3Pay       string `json:"executive3_pay"`
	Executive3Exercised string `json:"executive3_exercised"`
	Executive3Yearborn  string `json:"executive3_yearborn"`
	Executive4Name      string `json:"executive4_name"`
	Executive4Title     string `json:"executive4_title"`
	Executive4Pay       string `json:"executive4_pay"`
	Executive4Exercised string `json:"executive4_exercised"`
	Executive4Yearborn  string `json:"executive4_yearborn"`
	Executive5Name      string `json:"executive5_name"`
	Executive5Title     string `json:"executive5_title"`
	Executive5Pay       string `json:"executive5_pay"`
	Executive5Exercised string `json:"executive5_exercised"`
	Executive5Yearborn  string `json:"executive5_yearborn"`
	CompanyDecription   string `json:"company_description"`
}

type CompanyFinancialSummaryResult map[string]CompanyFinancialSummary
type CompanyFinancialSummary struct {
	Open                 string `json:"open"`
	PreviousClose        string `json:"previous_close"`
	MarketCap            string `json:"market_cap"`
	ForwardDividendYield string `json:"forward_dividend_yield"`
	AvgVolume            string `json:"avg_volume"`
	ExDividendDate       string `json:"ex_dividend_date"`
	Beta                 string `json:"beta"`
	PeRatio              string `json:"pe_ratio"`
	EPS                  string `json:"eps"`
	Week52Range          string `json:"week_52_range"`
	EarningsDate         string `json:"earnings_date"`
	Sentiment            float64 `json:"sentiment"`
}

type CompanyOwnershipStructureResult map[string]CompanyOwnershipStructure
type CompanyOwnershipStructure struct {
	PercentSharesHeldByInsiders float64 `json:"percent_of_shares_held_by_all_insider"`
	DateTime string `json:"datetime"`
	PercentFloatHeldByInstitutions float64 `json:"percent_of_float_held_by_institutions"`
	PercentSharesHeldByInstitutions float64 `json:"percent_of_shares_held_by_institutions"`
	NumberOfInstitutions uint64 `json:"number_of_institutions"`
	Timestamp uint64 `json:"timestamp"`
}

type CompanyInsiderTransactionResult map[string][]CompanyInsiderTransaction
type CompanyInsiderTransaction struct {
	ChangeOwn string `json:"change_own"`
	InsiderName string `json:"insider_name"`
	Price string `json:"price"`
	Owned string `json:"owned"`
	FillingDate string `json:"filing_date"`
	Qty string `json:"qty"`
	InsiderTitle string `json:"insider_title"`
	TradeType string `json:"trade_type"`
	Type string `json:"type"`
	TradeDate string `json:"trade_date"`
	Value string `json:"value"`
	FilingTimestamp uint64 `json:"filing_timestamp"`
}

func (c *Client) GetRealtimeStockPrice(request Request) (*RealTimePriceResult, error) {
	var results RealTimePriceResult
	err := c.Get("stock/realtime", request.GetUrlParams(), &results)

	return &results, err
}

func (c *Client) GetHistoricalStockPrice(request Request) (*HistoricalPriceResult, error) {
	var results HistoricalPriceResult
	err := c.Get("stock/historical", request.GetUrlParams(), &results)

	return &results, err
}

func (c *Client) GetCompanyFinancials(request Request) (*CompanyFinancialsResult, error) {
	var results CompanyFinancialsResult
	err := c.Get("company/financials", request.GetUrlParams(), &results)

	return &results, err
}

func (c *Client) GetCompanyProfile(request Request) (*CompanyProfileResult, error) {
	var results CompanyProfileResult
	err := c.Get("company/profile", request.GetUrlParams(), &results)

	return &results, err
}

func (c *Client) GetCompanyFinancialSummary(request Request) (*CompanyFinancialSummaryResult, error) {
	var results CompanyFinancialSummaryResult
	err := c.Get("company/financialSummary", request.GetUrlParams(), &results)

	return &results, err
}

func (c *Client) GetCompanyOwnershipStructure(request Request) (*CompanyOwnershipStructure, error) {
	var results CompanyOwnershipStructure
	err := c.Get("company/ownership", request.GetUrlParams(), &results)

	return &results, err
}

func (c *Client) GetCompanyInsiderTransactions(request Request) (*CompanyInsiderTransactionResult, error) {
	var results CompanyInsiderTransactionResult
	err := c.Get("company/insiderTransaction", request.GetUrlParams(), &results)

	return &results, err
}


func (c *Client) GetSecFilingLink(request Request) (*SecFilingLinkResult, error) {
	var results SecFilingLinkResult
	err := c.Get("company/secFilingLink", request.GetUrlParams(), &results)

	return &results, err
}

func (c *Client) GetStockNews(request Request) (*StockNewsResult, error) {
	var results StockNewsResult
	err := c.Get("company/news", request.GetUrlParams(), &results)

	return &results, err
}

func (c *Client) GetCorporateSplits(request Request) (*CorporateSplitsResult, error) {
	var results CorporateSplitsResult
	err := c.Get("company/actions/splits", request.GetUrlParams(), &results)

	return &results, err
}

func (c *Client) GetCorporateDividends(request Request) (*CorporateDividendsResult, error) {
	var results CorporateDividendsResult
	err := c.Get("company/actions/dividends", request.GetUrlParams(), &results)

	return &results, err
}
