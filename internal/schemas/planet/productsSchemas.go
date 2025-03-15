package planet

// Product представляет информацию о товаре
type Product struct {
	ID                    string             `json:"id"`
	Article               string             `json:"article"`
	Brand                 string             `json:"brand"`
	Title                 string             `json:"title"`
	Price                 float64            `json:"price"`
	Count                 int64              `json:"count"`
	SupplierID            int64              `json:"supplier_id"`
	SupplierNic           string             `json:"supplier_nic"`
	WarehouseDeliveryTime int64              `json:"warehouse_delivery_time"`
	QuantityInPack        int64              `json:"quantity_in_pack"`
	MeasureRatio          int64              `json:"measure_ratio"`
	MeasureUnit           string             `json:"measure_unit"`
	DeliverySchedules     []DeliverySchedule `json:"delivery_schedules"`
	WarehouseGUID         string             `json:"warehouse_guid"`
	ProductCode           string             `json:"product_code"`
	Images                []string           `json:"images"`
	ReturnDisabled        int64              `json:"return_disabled"`
}

// DeliverySchedule представляет расписание доставки
type DeliverySchedule struct {
	Date          string `json:"date"`
	ScheduleID    int64  `json:"schedule_id"`
	Address       string `json:"address"`
	BeginInterval string `json:"begin_interval"`
	EndInterval   string `json:"end_interval"`
}
