package dto

/*
厦门大学计算机专业 | 前华为工程师
专注《零基础学编程系列》  http://lblbc.cn/blog
包含：Java | 安卓 | 前端 | Flutter | iOS | 小程序 | 鸿蒙
公众号：蓝不蓝编程
*/
type RegisterDTORequest struct {
	Name     string `json:"name" form:"name" binding:"required,min=1,max=100"`
	Password string `json:"password" form:"password" binding:"required,max=100"`
}
