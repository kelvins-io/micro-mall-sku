package models

import (
	"time"
)

type SkuInventory struct {
	Id         int64     `xorm:"pk autoincr comment('商品库存ID') BIGINT"`
	SkuCode    string    `xorm:"not null comment('商品编码') unique unique(sku_code_shop_id_index) CHAR(64)"`
	Amount     int64     `xorm:"comment('库存数量') BIGINT"`
	Price      string    `xorm:"comment('入库单价') DECIMAL(32,16)"`
	ShopId     int64     `xorm:"not null comment('所属店铺ID') index unique(sku_code_shop_id_index) BIGINT"`
	Version    int       `xorm:"not null default 1 comment('商品版本') INT"`
	CreateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type SkuInventoryRecord struct {
	Id           int64     `xorm:"pk autoincr comment('自责ID') BIGINT"`
	ShopId       int64     `xorm:"comment('店铺ID') BIGINT"`
	SkuCode      string    `xorm:"comment('商品sku') CHAR(40)"`
	OpType       int       `xorm:"default 0 comment('操作类型，0-入库，1-出库，2-冻结') TINYINT"`
	OpUid        int64     `xorm:"comment('操作的用户ID') BIGINT"`
	OpIp         string    `xorm:"comment('操作IP地址') VARCHAR(255)"`
	AmountBefore int64     `xorm:"comment('变化之前数量') BIGINT"`
	Amount       int64     `xorm:"comment('操作数量') BIGINT"`
	OpTxId       string    `xorm:"comment('操作的事务ID') index CHAR(60)"`
	State        int       `xorm:"default 0 comment('状态，0-有效，1-锁定中，2-无效') TINYINT"`
	Verify       int       `xorm:"default 0 comment('是否核实，0-未核实，1-已核实') TINYINT"`
	CreateTime   time.Time `xorm:"default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime   time.Time `xorm:"default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type SkuPriceHistory struct {
	Id         int64     `xorm:"pk autoincr comment('自增ID') BIGINT"`
	ShopId     int64     `xorm:"not null comment('调价的店铺id') index(shop_id_sku_code_index) BIGINT"`
	SkuCode    string    `xorm:"not null comment('商品sku_code') index(shop_id_sku_code_index) index CHAR(40)"`
	Price      string    `xorm:"not null comment('调整后价格') DECIMAL(32,16)"`
	Reason     string    `xorm:"comment('调价说明') TEXT"`
	Version    int       `xorm:"comment('调整版本') INT"`
	OpUid      int64     `xorm:"comment('操作员UID') BIGINT"`
	OpIp       string    `xorm:"comment('操作员IP') CHAR(16)"`
	CreateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
}

type SkuProperty struct {
	Id            int64     `xorm:"pk autoincr comment('ID') BIGINT"`
	Code          string    `xorm:"not null comment('商品唯一编号') index CHAR(64)"`
	Price         string    `xorm:"comment('商品当前价格') DECIMAL(32,16)"`
	Name          string    `xorm:"comment('商品名称') index VARCHAR(255)"`
	Desc          string    `xorm:"comment('商品描述') TEXT"`
	Production    string    `xorm:"comment('生产企业') VARCHAR(1024)"`
	Supplier      string    `xorm:"comment('供应商') VARCHAR(1024)"`
	Category      int       `xorm:"comment('商品类别') INT"`
	Title         string    `xorm:"comment('商品标题') VARCHAR(255)"`
	SubTitle      string    `xorm:"comment('商品副标题') VARCHAR(255)"`
	Color         string    `xorm:"comment('商品颜色') VARCHAR(64)"`
	ColorCode     int       `xorm:"comment('商品颜色代码') INT"`
	Specification string    `xorm:"comment('商品规格') VARCHAR(255)"`
	DescLink      string    `xorm:"comment('商品介绍链接') VARCHAR(255)"`
	State         int       `xorm:"default 0 comment('商品状态，0-有效，1-无效，2-锁定') TINYINT"`
	CreateTime    time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime    time.Time `xorm:"not null comment('更新时间') DATETIME"`
}
