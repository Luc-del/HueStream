package main

import (
	"HueStream/auth"
	"fmt"
	"github.com/amimof/huego"
)

var (
	configFile = "config.yml"
)

func main() {


	//bridge, err := huego.Discover()
	//if err != nil {
	//	panic(err)
	//}
	//
	//
	//user, err := bridge.CreateUser("Lululuc") // Link button needs to be pressed
	//if err != nil {
	//	panic(err)
	//}
	//bridge = bridge.Login(user)
	//
	//fmt.Printf("%+v", bridge)
	//
	//
	//light, _ := bridge.GetLight(3)
	//err = light.Off()
	//if err != nil {
	//	panic(err)
	//}
	//

	cfg := auth.LoadConfig(configFile)
	bridge := huego.New(cfg.Ip, cfg.User)
	lights, err := bridge.GetLights()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %d lights\n", len(lights))
	fmt.Println(lights)

	fmt.Printf("\n%+v\n", lights[0])
}
