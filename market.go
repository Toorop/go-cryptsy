package cryptsy

type Market struct {
	Marketid       int64   `json:"marketid"`
	Label          string  `json:"label"`
	LastTradePrice float64 `json:"lasttradeprice"`
	LastTradeTime  string  `json:"lasttradetime"`
	Volume         float64 `json:"volume"`
	PrimaryName    string  `json:"primaryname"`
	PrimaryCode    string  `json:"primarycode"`
	SecondaryName  string  `json:"secondaryname"`
	SecondaryCode  string  `json:"secondarycode"`
	RecentTrades   []trade `json:"recenttrades"`
	SellOrders     []Order `json:"sellorders"`
	BuyOrders      []Order `json:"buyoders"`
}
