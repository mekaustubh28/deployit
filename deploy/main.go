package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ready to hear from the redis queue")
	for {
		item := redis_pop("deploy")
		if item != "" {
			fmt.Println("Popped item:", item)
			building(item)
			redis_push(item+": Deployment is ready at : "+"http://"+item+".localhost:3000/", "log:"+item)
			redis_push("🙏 Thank you for using Deployit.", "log:"+item)
		}
		time.Sleep(2 * time.Second) // Add a sleep to simulate processing time
	}
}
