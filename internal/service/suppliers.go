package service

import (
	"auto_order/internal/adapters"
	"log"
)

type Suppliers struct {
	suppliers     []adapters.Supplier
	supplierNames map[string]int
}

func NewSuppliers() *Suppliers {
	return &Suppliers{suppliers: adapters.NewSuppliers()}
}

func (s *Suppliers) GetSupplierService(id int) adapters.Supplier {
	if id >= len(s.suppliers) {
		log.Println("SuppliersAPI: GetSupplierService: out of range")
		return nil
	}

	sup := s.suppliers[id]
	if sup == nil {
		log.Println("SuppliersAPI: GetSupplierService: supplier not found")
	}

	return sup
}

func (s *Suppliers) GetSupplierServiceByName(name string) adapters.Supplier {
	return nil
}
