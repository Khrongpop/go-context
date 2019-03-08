package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("start process context with time out...")
	wg := sync.WaitGroup{}

	// Create Context with timeout
	tmo := 3 * time.Second
	ctxCooking, cancelCooking := context.WithTimeout(context.Background(), tmo)
	defer cancelCooking()

	wg.Add(1)
	// Split cooking cast
	go func() {
		defer wg.Done()
		cooking(ctxCooking,"Fried egg",10,"c")
	}()
	wg.Wait()
}

func cooking(ctx context.Context, food string, m int, chef string) {
	for i:=1; i<=m; i++ {
		select {
			case<-ctx.Done():
				fmt.Printf("Chef %v cooked %v total %v dish \n",chef,food,i-1)
				return
			default: 
				fmt.Printf("Chef %v cooking %v : %v \n",chef,food,i)
				time.Sleep(1*time.Second)
		}
	}
	return
}