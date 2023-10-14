package main

type ShopDetailResp struct {
	Data ShopDetailData `json:"dt"`
}

type ShopDetailData []ShopDetailEntry

type ShopDetailEntry struct {
	StoreName string `json:"storeName"` // 商店名稱
	ItemID    int64  `json:"itemID"`    // 物品編號
	ItemName  string `json:"itemName"`  // 物品名稱
	ItemCNT   int64  `json:"itemCNT"`   // 單價
	ItemPrice int64  `json:"itemPrice"` // 數量
	Storetype int    `json:"storetype"` // 銷售(1)/販賣(0)
}

type HistoryResp struct {
	Data HistoryData `json:"dt"`
}

type HistoryData []HistoryEntry

type HistoryEntry struct {
	ItemID   int64  `json:"itemID"`   // 物品編號
	ItemName string `json:"itemName"` // 物品名稱
	Token    string `json:"token"`    // DealPerDay token
}
