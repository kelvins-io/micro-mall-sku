package args

type SkuInventoryInfo struct {
	ShopId        int64  `json:"shop_id"`
	SkuCode       string `json:"sku_code"`
	Name          string `json:"name"`
	Price         string `json:"price"`
	Title         string `json:"title"`
	SubTitle      string `json:"sub_title"`
	Desc          string `json:"desc"`
	Production    string `json:"production"`
	Supplier      string `json:"supplier"`
	Category      int32  `json:"category"`
	Color         string `json:"color"`
	ColorCode     int32  `json:"color_code"`
	Specification string `json:"specification"`
	DescLink      string `json:"desc_link"`
}

type SkuPropertyEx struct {
	OpUid             int64  `bson:"op_uid"`
	OpIp              string `bson:"op_ip"`
	ShopId            int64  `bson:"shop_id"`
	SkuCode           string `bson:"sku_code"`
	Name              string `bson:"name"`
	Size              string `bson:"size"`
	Shape             string `bson:"shape"`
	ProductionCountry string `bson:"production_country"`
	ProductionDate    string `bson:"production_date"`
	ShelfLife         string `bson:"shelf_life"`
}

type InventoryState struct {
	ShopId   int64    `json:"shop_id"`
	SkuCodes []string `json:"sku_codes"`
}

type OperationInventoryRsp struct {
	List []InventoryState `json:"list"`
}

const (
	RpcServiceMicroMallUsers  = "micro-mall-users"
	RpcServiceMicroMallSearch = "micro-mall-search"
)

const (
	SkuInventorySearchNoticeTag    = "sku_inventory_search_notice"
	SkuInventorySearchNoticeTagErr = "sku_inventory_search_notice_err"
)

const (
	Unknown                  = 0
	SkuInventorySearchNotice = 1000
)

var MsgFlags = map[int]string{
	Unknown:                  "未知",
	SkuInventorySearchNotice: "商品库存搜索通知",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Unknown]
}

type CommonBusinessMsg struct {
	Type    int    `json:"type"`
	Tag     string `json:"tag"`
	UUID    string `json:"uuid"`
	Content string `json:"content"`
}
