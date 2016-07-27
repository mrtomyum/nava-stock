package model

import "time"

type Order struct {
	Base
	Number string `json:"number"`
	Date   time.Time `json:"order_status"`
	Status Status `json:"status"`
}

type OrderItemStatus int

const (
	HOLD_ITEM OrderItemStatus = 1 + iota
	PICK_ITEM
	BACK_ORDER_ITEM // สินค้าไม่เพียงพอใน AVL ส่วนเกินจะรอสร้าง PO DRAFT
	PO_DRAFTED_ITEM // เอกสารกำลังรออนุมัติ อาจไม่ใช้ โดยอาจรวมกับ BACKORDER
	PO_APPROVED_ITEM // PO อนุมัติ และส่งอีเมล์แล้ว
	SHIPPING_ITEM // Vendor ส่งสินค้าออกมาแล้ว สินค้าอยู่ระหว่างจัดส่ง
	RECEIVED_ITEM // ได้รับสินค้าแล้ว
)

type OrderItem struct {
	OrderID  uint64
	Date     time.Time `json:"date"`
	Status   OrderItemStatus `json:"status"`
	ItemID   uint64 `json:"item_id"`
	ItemName string `json:"item_name"` // may differ from master data from time to time so it be recorded
	UnitSale Unit `json:"unit_sale`
	Price    uint64
	Discount uint64
	Total    uint64
} // Todo: Next stage Document that generate from this must be refer back to this Order Item
