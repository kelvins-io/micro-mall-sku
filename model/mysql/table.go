package mysql

import (
	"time"
)

const (
	TableSkuInventory = "sku_inventory"
	TableSkuProperty  = "sku_property"
)

type SkuInventory struct {
	Id         int64     `xorm:"'id' pk autoincr comment('商品库存ID') BIGINT"`
	SkuCode    string    `xorm:"'sku_code' not null comment('商品编码') unique unique(sku_code_shop_id_index) CHAR(64)"`
	Amount     int64     `xorm:"'amount' comment('库存数量') BIGINT"`
	Price      string    `xorm:"'price' comment('入库单价') DECIMAL(32,16)"`
	ShopId     int64     `xorm:"'shop_id' not null comment('所属店铺ID') index unique(sku_code_shop_id_index) BIGINT"`
	CreateTime time.Time `xorm:"'create_time' not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"'update_time' not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type SkuProperty struct {
	Id            int64     `xorm:"'id' pk autoincr comment('ID') BIGINT"`
	Code          string    `xorm:"'code' not null comment('商品唯一编号') index CHAR(64)"`
	Price         string    `xorm:"'price' comment('商品当前价格') DECIMAL(10,2)"`
	Name          string    `xorm:"'name' comment('商品名称') index VARCHAR(255)"`
	Desc          string    `xorm:"'desc' comment('商品描述') TEXT"`
	Production    string    `xorm:"'production' comment('生产企业') VARCHAR(1024)"`
	Supplier      string    `xorm:"'supplier' comment('供应商') VARCHAR(1024)"`
	Category      int       `xorm:"'category' comment('商品类别') INT"`
	Title         string    `xorm:"'title' comment('商品标题') VARCHAR(255)"`
	SubTitle      string    `xorm:"'sub_title' comment('商品副标题') VARCHAR(255)"`
	Color         string    `xorm:"'color' comment('商品颜色') VARCHAR(64)"`
	ColorCode     int       `xorm:"'color_code' comment('商品颜色代码') INT"`
	Specification string    `xorm:"'specification' comment('商品规格') VARCHAR(255)"`
	DescLink      string    `xorm:"'desc_link' comment('商品介绍链接') VARCHAR(255)"`
	State         int       `xorm:"'state' default 0 comment('商品状态，0-有效，1-无效，2-锁定') TINYINT"`
	CreateTime    time.Time `xorm:"'create_time' not null comment('创建时间') DATETIME"`
	UpdateTime    time.Time `xorm:"'update_time' not null comment('更新时间') DATETIME"`
}
