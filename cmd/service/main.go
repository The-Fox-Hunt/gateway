package main

import (
	"fmt"
	"github.com/The-Fox-Hunt/gateway/internal/api"
)

func main() {
	handler := api.NewHandler()
	res := handler.Handle("a", "b")
	fmt.Println(res)
}
