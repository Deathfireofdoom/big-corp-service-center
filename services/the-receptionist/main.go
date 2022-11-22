package main

import "github.com/Deathfireofdoom/big-corp-service-center/the-receptionist/internal/router"

func main() {
	var router_ = router.Configure()
	router_.Run()
}
