package view

import (
	"customerManage/modle"
	"customerManage/service"
	"fmt"
)

type CustomerView struct {
	//接受用户输入
	key string
	//判断是否循环显示菜单
	loop            bool
	customerService *service.CustomerService
}

func NewCustomerView() *CustomerView {
	customerView := &CustomerView{}
	customerView.key = ""
	customerView.loop = true
	customerView.customerService = service.NewCustomerService()
	return customerView
}

func (this *CustomerView) Mainmeun() {
	for {
		fmt.Println("\n---------客户信息管理软甲-------------")
		fmt.Println("         1.添加客户")
		fmt.Println("         2.修改客户")
		fmt.Println("         3.删除客户")
		fmt.Println("         4.客户列表")
		fmt.Println("         5.退出")
		fmt.Printf("请选择（1-5）：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.Add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			this.loop = true
			fmt.Println("请输入正确的选项..")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("您退出了客户信息管理软甲")
}
func (this *CustomerView) update() {
	fmt.Println("-----------修改客户-----------------")
	fmt.Println("请输入要修改1的客户号（-1退出）：")
	id := -1
	customer := modle.Customer{}
	for {
		fmt.Scanln(&id)
		if id == -1 {
			return
		}
		customer = this.customerService.GetCustomer(id)
		if customer.Id != 0 {
			break
		}
		fmt.Println("没有找到该客户！")
	}
	fmt.Println("-----------请输入修改内容-----------------")
	fmt.Print("姓名(", customer.Name, "):")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别(", customer.Sex, "):")
	sex := ""
	fmt.Scanln(&sex)
	fmt.Print("年龄(", customer.Age, "):")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话(", customer.Phone, "):")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱(", customer.Email, "):")
	email := ""
	fmt.Scanln(&email)
	customer = modle.NewCustomer(customer.Id, name, sex, age, phone, email)
	loop := this.customerService.Update(customer)
	if loop {
		fmt.Println("----------修改成功------------")
	}

}
func (this *CustomerView) delete() {
	fmt.Println("-----------删除客户-----------------")
	fmt.Println("请输入要删除的客户号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Print("你确定要删除该客户嘛？(y/n):")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Print("输入有误，请重新输入！(y/n):")
	}
	if choice == "y" {
		loop := this.customerService.Delete(id)
		if loop {
			fmt.Println("删除成功！")
		} else {
			fmt.Println("没有该客户，删除失败！")
		}
	}

}
func (this *CustomerView) exit() {
	fmt.Print("你确定要退出本系统嘛？(y/n):")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Print("输入有误，请重新输入！(y/n):")
	}
	if choice == "y" {
		this.loop = false
	}

}
func (this *CustomerView) list() {
	customers := this.customerService.List()
	fmt.Println("--------客户列表-------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := range customers {
		customerInfo := customers[i].CustomerToString()
		fmt.Println(customerInfo)
	}
	fmt.Println("--------客户列表完成-------------")
}
func (this *CustomerView) Add() {
	fmt.Println("----------添加客户------------")
	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别：")
	sex := ""
	fmt.Scanln(&sex)
	fmt.Print("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱：")
	email := ""
	fmt.Scanln(&email)
	/*customer:=modle.Customer{}
	customer.Name =name
	customer.Sex =sex
	customer.Age =age
	customer.Phone =phone
	customer.Email =email*/
	customer := modle.NewCustomerNoId(name, sex, age, phone, email)
	loop := this.customerService.Add(customer)
	if loop {
		fmt.Println("----------添加成功------------")
	}
}
