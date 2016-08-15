package model

import (
	sys "github.com/mrtomyum/nava-sys/model"
)

type machineType int

const (
	CAN machineType = 1 + iota
	CUP
	CUP_FRESH_COFFEE
	CUP_NOODLE
	SEE_THROUGH
)

type MachineBrand int

const (
	NATIONAL MachineBrand = 1 + iota
	SANDEN
	FUJI_ELECTRIC
	CIRBOX
)

type Machine struct {
	sys.Base
	LocID        uint64       `json:"loc_id" db:"loc_id"`
	Code         string       `json:"code"`
	Type         machineType  `json:"type"`
	Brand        MachineBrand `json:"brand"`
	Model        string       `json:"model"`
	SerialNumber string
	Selection    int          `json:"selection"`  //จำนวน Column หรือช่องเก็บสินค้า
											   //LocRow int  	//จำนวนแถว และคอลัมน์ไว้ทำ Schematic Profile  หน้าตู้
											   //LocCol int  //ควรจะเป็น 2 Dimension Array
}

type ColumnType int

const (
	FREE ColumnType = iota //สินค้าไม่มีตัวตน หรือต้องส่งข้อมูลสั่งขายไปยังระบบอื่น
	TICKET                   // สินค้าที่ต้องพิมพ์ตั๋ว
	CAN_S                    // กระป๋องหรือขวดสั้น
	CAN_L                    // กระป๋องหรือขวดยาว
	SPRING_S
	SPRING_M
	SPRING_L
)

type ColumnStatus int

const (
	OK ColumnStatus = iota
	FAIL
)

// MachineColumn เก็บยอด Counter ล่าสุดของแต่ละ column ในแต่ละ Machine
type MachineColumn struct {
	sys.Base
	MachineID uint64       `json:"machine_id"`
	ColumnNo  int          `json:"column_no"`
	SaleCount int          `json:"sale_count"`
	Type      ColumnType   `json:"type"`
	Status    ColumnStatus `json:"status"`
}

// Design this struct for data from VMC telemetry system.
//type SaleStatus int
//const (
//	COMPLETED SaleStatus = iota
//	INCOMPLETED
//)
type MachineRealTimeSale struct {
	sys.Base
	MachineID uint64 `json:"machine_id"`
	ColumnNo  int    `json:"column_no"`
	ItemID    uint64 `json:"item_id"`
	Price     int `json:"price"`
	//Status SaleStatus `json:"status"`
}

// Receive Batch data from mobile app daily.
type MachineBatchSale struct {
	sys.Base
	MachineID uint64 `json:"machine_id"`
	ColumnNo  int    `json:"column_no"`
	Counters  int    `json:"counters"`
	SalePrice Price
}

// เก็บ Transaction ที่มีความผิดปกติทั้งหมด เช่น  ข้อมูลที่ส่งมาหา Column ไม่เจอ ไปจนถึง Error ที่แจ้งจาก Machine
type MachineErrType int

const (
	X MachineErrType = iota // UNIDENTIFIED ERROR
	COLUMN_NOT_FOUND
	COUNTER_OVER_SALE
)

type MachineErrLog struct {
	sys.Base
	MachineID uint64         `json:"machine_id"`
	ColumnNo  int            `json:"column_no"`
	Type      MachineErrType `json:"type"`
	Message   string         `json:"message"`
}
