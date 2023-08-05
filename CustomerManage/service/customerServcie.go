package service

import (
	"fmt"
	"golangwork/CustomerManage/model"
)

// 完成对customer的操作
// 增删改查
type CustomerService struct {
	customers []model.Customer
	//声明变量表示当前切片有多少个客户,后面作为新编号
	customerNum int
}

// 编写返回实例的方法
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// 添加用户切片
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) FindById(id int) int {
	//遍历数组
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}

	}
	return index
}
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		fmt.Println("\t\t\t==========不存在此编号,删除失败==============")
		return false

	} else {

		var choose string
		this.PrintByIndex(index)
		fmt.Println("\t\t\t确认删除吗？（y/n ）")
		fmt.Scanln(&choose)
		if choose == "n" {
			fmt.Println("\t\t\t==========退出删除===============")
			return false
		}
		//如何从切片中删除元素
		this.customers = append(this.customers[:index], this.customers[index+1:]...)
		return true

	}

}
func (this *CustomerService) PrintByIndex(index int) {
	fmt.Printf("\t\t\t%v\t%v\t%v\t%v\t%v\t%v\t", this.customers[index].Id, this.customers[index].Name, this.customers[index].Gender, this.customers[index].Age, this.customers[index].Phone, this.customers[index].Email)
}

func (this *CustomerService) Update(id int, index int) bool {
	//修改编号存在
	fmt.Println("\t\t\t========确认修改该用户吗(y/n)========")
	this.PrintByIndex(index)
	choose := ""
	fmt.Scanln(&choose)
	if choose == "" {
		fmt.Println("\t\t\t================退出修改。。。=================")
		return false
	}
	fmt.Println("\t\t\t=====================修改界面=====================")
	fmt.Println("姓名：")
	fmt.Scanln(&this.customers[index].Name)
	fmt.Println("性别：")
	fmt.Scanln(&this.customers[index].Gender)
	fmt.Println("年龄：")
	fmt.Scanln(&this.customers[index].Age)
	fmt.Println("电话：")
	fmt.Scanln(&this.customers[index].Phone)
	fmt.Println("邮箱：")
	fmt.Scanln(&this.customers[index].Email)
	fmt.Println("\t\t\t=====================修改页面结束==================")
	return true

}
