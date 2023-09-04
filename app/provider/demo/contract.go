package demo

// Demo服务的key
const DemoKey = "demo"

// Demo服务的接口
type IService interface {
	GetAllStudent() []Student
}
