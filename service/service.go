//@Author: Dennis
//@Date: 2023/5/12

package service

import (
	"time"
)

const duration = 200 * time.Millisecond

var orderService = func(p string, s string) string {
	time.Sleep(duration)
	return p + "+" + s + "->" + "order"
}

var productService = func() string {
	time.Sleep(duration)
	return "product"
}

var shippingService = func() string {
	time.Sleep(duration)
	return "shipping"
}

func RemoteCallAsync() string {
	c := make(chan string)
	go func() {
		c <- productService()
	}()
	go func() {
		c <- shippingService()
	}()
	p := <-c
	s := <-c
	go func() {
		c <- orderService(p, s)
	}()
	return <-c
}

func RemoteCallSync() string {
	p := productService()
	s := shippingService()
	return orderService(p, s)
}
