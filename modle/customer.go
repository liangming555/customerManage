package modle

import "fmt"

type Customer struct {
	Id    int
	Name  string
	Sex   string
	Age   int
	Phone string
	Email string
}

func NewCustomer(id int, name string, sex string, age int, phone string, email string) Customer {
	return Customer{
		Id:    id,
		Name:  name,
		Sex:   sex,
		Age:   age,
		Phone: phone,
		Email: email,
	}
}
func NewCustomerNoId(name string, sex string, age int, phone string, email string) Customer {
	return Customer{
		Name:  name,
		Sex:   sex,
		Age:   age,
		Phone: phone,
		Email: email,
	}
}
func (this *Customer) CustomerToString() string {
	customerInfo := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Sex, this.Age, this.Phone, this.Email)
	return customerInfo
}
