package models

type CartItem struct {
	UserID int64
	SKU    uint32
	Count  uint16
}

type CartItemWithDetails struct {
	CartItem   CartItem
	Name       string
	Price      uint32
	TotalPrice uint32
}
type StockItemDetails struct {
	Name  string
	Price uint32
}
