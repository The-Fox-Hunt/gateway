module github.com/The-Fox-Hunt/gateway

go 1.23.3

require (
	github.com/The-Fox-Hunt/auth v0.0.3
	google.golang.org/grpc v1.71.0
)

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250303144028-a0af3efb3deb // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)

//Почему auth не подтягивается, например, при обновлении прото файла
//replace github.com/The-Fox-Hunt/auth => ../auth
