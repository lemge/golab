# golab
the golang lab object,this is a teacher and student in there and test and study




2022-09-26 08:18
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ProductInfo.proto

将未实现的服务关闭 就能解决

或者在结构体中实现
type productInfoServer struct {
	product.UnimplementedProductInfoServer //加上这个接口就可以了。
	productMap map[string]*product.Product
}
