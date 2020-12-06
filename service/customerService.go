package service

import (
	"customerManage/modle"
)

type CustomerService struct {
	customers  []modle.Customer
	customerNo int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNo = 1
	customer := modle.NewCustomer(1, "张三", "男", 18, "13188888888", "13188888888@163.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (this *CustomerService) List() []modle.Customer {
	return this.customers
}

func (this *CustomerService) Add(customer modle.Customer) bool {
	this.customerNo++
	customer.Id = this.customerNo
	this.customers = append(this.customers, customer)
	return true
}
func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := range this.customers {
		if id == this.customers[i].Id {
			index = i
		}
	}
	return index
}
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	//从切片中删除元素，即从0开始到删除元素位置与删除元素后一个元素到结尾进行拼接
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}
func (this *CustomerService) GetCustomer(id int) modle.Customer {
	index := this.FindById(id)
	customer := modle.Customer{}
	if index == -1 {
		return customer
	}
	return this.customers[index]
}
func (this *CustomerService) Update(customer modle.Customer) bool {
	index := this.FindById(customer.Id)
	/**
	修改切片中某一元素
	1.根据下标先删除对应元素
	2.从下标开始到切片最终建立临时切片
	3.在目标切片从0到下标处（不包含）新增内容
	4.将新增内容后的切片拼接上临时切片
	*/
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	temp := append([]modle.Customer{}, this.customers[index:]...)
	this.customers = append(this.customers[:index], customer)
	this.customers = append(this.customers, temp...)
	return true
}
