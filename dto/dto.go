package dto

const (
	PORING_SERVER string = "4290"
)

type ShopDetailResp struct {
	Data ShopDetailData `json:"dt"`
}

type ShopDetailData []ShopDetailEntry

type ShopDetailEntry struct {
	StoreName string `json:"storeName"` // 商店名稱
	ItemID    int    `json:"itemID"`    // 物品編號
	ItemName  string `json:"itemName"`  // 物品名稱
	ItemCNT   int    `json:"itemCNT"`   // 單價
	ItemPrice int    `json:"itemPrice"` // 數量
	Storetype int    `json:"storetype"` // 銷售(1)/販賣(0)
}

type HistoryPayload struct {
	Server    string `json:"div_svr"`          // 伺服器ID
	Days      string `json:"div_history_days"` // 查詢天數
	Keyword   string `json:"txb_KeyWord"`      // 道具關鍵字
	Recaptcha string `json:"recaptcha"`        // 不管它，值寫死 ""
	SortBy    string `json:"sort_by"`          // 不管它，值寫死 "SumitemCNT"
	SortDesc  string `json:"sort_desc"`        // 不管它，值寫死 "desc"
}

type HistoryResp struct {
	Token string      `json:"token"` // TransactionGroupByDay token
	Data  HistoryData `json:"dt"`
}

type HistoryData []HistoryEntry

type HistoryEntry struct {
	ItemID          int    `json:"itemID"`   // 物品編號
	EncryptedItemID string `json:"itemID_e"` // TransactionGroupByDay itemID
}

type TransactionCountPayload struct {
	Server          string `json:"div_svr"`          // 伺服器ID
	Days            string `json:"div_history_days"` // 查詢天數
	EncryptedItemID string `json:"itemID"`           // TransactionGroupByDay itemID
}

type TransactionPerDayPayload struct {
	Server          string `json:"div_svr"`          // 伺服器ID
	Days            string `json:"div_history_days"` // 查詢天數
	RowStart        string `json:"row_start"`        // 第幾筆紀錄 (pagination)
	EncryptedItemID string `json:"itemID"`           // TransactionGroupByDay itemID
	Token           string `json:"token"`            // TransactionGroupByDay token
	SortBy          string `json:"sort_by"`          // 不管它，值寫死 "regDate_"
	SortDesc        string `json:"sort_desc"`        // 不管它，值寫死 "desc"
}

type TransactionPerDayResp struct {
	Data TransactionPerDayData `json:"dt"`
}

type TransactionPerDayData []TransactionPerDay

type TransactionPerDay struct {
	ItemID          int    `json:"itemID"`   // 物品編號
	EncryptedItemID string `json:"itemID_e"` // TransactionGroupByDay itemID
}
