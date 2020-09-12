//Package po generated by 'freedom new-po'
package po

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Goods .
type Goods struct {
	changes map[string]interface{}
	ID      int       `gorm:"primary_key;column:id"`
	Name    string    `gorm:"column:name"`  // 商品名称
	Price   int       `gorm:"column:price"` // 价格
	Stock   int       `gorm:"column:stock"` // 库存
	Tag     string    `gorm:"column:tag"`   // 标签
	Created time.Time `gorm:"column:created"`
	Updated time.Time `gorm:"column:updated"`
}

// TableName .
func (obj *Goods) TableName() string {
	return "goods"
}

// TakeChanges .
func (obj *Goods) TakeChanges() map[string]interface{} {
	if obj.changes == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range obj.changes {
		result[k] = v
	}
	obj.changes = nil
	return result
}

// updateChanges .
func (obj *Goods) setChanges(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}

// SetName .
func (obj *Goods) SetName(name string) {
	obj.Name = name
	obj.setChanges("name", name)
}

// SetPrice .
func (obj *Goods) SetPrice(price int) {
	obj.Price = price
	obj.setChanges("price", price)
}

// SetStock .
func (obj *Goods) SetStock(stock int) {
	obj.Stock = stock
	obj.setChanges("stock", stock)
}

// SetTag .
func (obj *Goods) SetTag(tag string) {
	obj.Tag = tag
	obj.setChanges("tag", tag)
}

// SetCreated .
func (obj *Goods) SetCreated(created time.Time) {
	obj.Created = created
	obj.setChanges("created", created)
}

// SetUpdated .
func (obj *Goods) SetUpdated(updated time.Time) {
	obj.Updated = updated
	obj.setChanges("updated", updated)
}

// AddPrice .
func (obj *Goods) AddPrice(price int) {
	obj.Price += price
	obj.setChanges("price", gorm.Expr("price + ?", price))
}

// AddStock .
func (obj *Goods) AddStock(stock int) {
	obj.Stock += stock
	obj.setChanges("stock", gorm.Expr("stock + ?", stock))
}