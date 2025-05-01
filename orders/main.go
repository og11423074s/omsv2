package main

import "context"

func main() {

	store := NewStore()
	svc := NewService(store)

	err := svc.CreateOrder(context.Background())
	if err != nil {
		return
	}
}
