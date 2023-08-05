package main

import (
	"fmt"
	"golangwork/CustomerManage/model"
	"golangwork/CustomerManage/service"
)

type customerView struct {
	key  string //接受用户输入
	loop bool   //表示是否循环显示菜单
	//增加customerService
	customerService *service.CustomerService
}

func (cv *customerView) mainMenu() {
	for {
		fmt.Println("				========================客户信息管理系统=========================")
		fmt.Println("							1.添加客户")
		fmt.Println("							2.修改客户")
		fmt.Println("							3.删除客户")
		fmt.Println("							4.客户列表")
		fmt.Println("							5.退   出")
		fmt.Println("							请选择（1-5）：")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.Update()

		case "3":
			cv.delete()

		case "4":
			cv.list()
		case "5":
			cv.loop = false
		default:
			fmt.Println("你的输入有误请重新输入")

		}
		if !cv.loop {
			break
		}
	}
	fmt.Println("你已经退出客户管理系统。。。。")
}

func (this *customerView) list() {
	//首先获取到客户的信息
	customers := this.customerService.List()
	fmt.Println("					==============客户列表===================")
	fmt.Println("					编号\t姓名\t性别\t年龄\t电话\t\t邮箱		")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("					==========客户列表完成====================\n\n")

}

// 添加客户
func (this *customerView) add() {
	fmt.Println("\t\t\t====================添加客户=====================")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	var age int
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomerWithoutId(name, gender, age, phone, email)
	//分配id
	if this.customerService.Add(customer) == true {
		fmt.Println("添加成功")
	} else {
		fmt.Println("添加失败")
	}

}
func (this *customerView) delete() {
	var id int
	fmt.Println("请输入要删除的编号：(-1退出)")
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	//	业务逻辑应该由service分担
	/*	index := this.customerService.FindById(id)
		if index == -1 {
			fmt.Println("\t\t\t==========不存在此编号,删除失败==============")
			return

		} else {
			//打印客户信息
			var choose string
			this.customerService.PrintByIndex(index)
			fmt.Println("\t\t\t确认删除吗？（y/n ）")
			fmt.Scanln(&choose)
			if choose == "n" {
				fmt.Println("\t\t\t==========退出删除===============")
				return
			}
			flag := this.customerService.Delete(id)
			if flag {
				fmt.Println("\t\t\t==============删除成功================")

			} else {

				fmt.Println("\t\t\t==============删除失败================")
			}
		}
	*/
	flag := this.customerService.Delete(id)
	if flag {
		fmt.Println("\t\t\t================删除成功===============")
	}
}

func (this *customerView) Update() {
	fmt.Println("\t\t\t请选择你要修改的编号")
	id := 0
	fmt.Scanln(&id)
	index := this.customerService.FindById(id)
	if index == -1 {
		fmt.Println("\t\t\t=========修改编号不存在===========")
		return
	}
	flag := this.customerService.Update(id, index)
	if flag {
		fmt.Println("\t\t\t==========修改成功============")
	}
}

func main() {
	//在函数中创建一个实例
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//完成对结构体的初始化
	customerView.customerService = service.NewCustomerService()
	//显示菜单
	customerView.mainMenu()
}
