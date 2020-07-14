package main

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v2"
)

type (
	// Ingredient is needed for Cookie
	Ingredient struct {
		ID   int
		Name string
	}
	// Cookie is delicious treat
	Cookie struct {
		Ingredients []Ingredient `yaml:"ingredients"`
	}
	// Oven interface is requirement for baked cookie
	Oven interface {
		Store(Cookie)
		Bake(time.Duration)
		Get() []Cookie
	}
	// Our fancy selfmade Oven, with private Property
	CookieCutterOven struct {
		Storage         []Cookie
		privateProperty bool
	}
)

// PrepareCookie preapres a delicious cookie
func PrepareCookie(ingredients []Ingredient) Cookie {
	if len(ingredients) < 1 {
		panic(("No ingredients, how to make a cookie?"))
	}
	return Cookie{Ingredients: ingredients}
}

func (cc CookieCutterOven) Store(cookie Cookie) {
	cc.Storage = append(cc.Storage, cookie)
}

func (cc CookieCutterOven) Bake(duration time.Duration) {
	time.Sleep(duration)
	fmt.Println("Ding Ding, cookies are done!")
}

func (cc CookieCutterOven) Get() []Cookie {
	return cc.Storage
}

func main() {
	var delicious = []Ingredient{
		Ingredient{Name: "Sugar", ID: 0},
		Ingredient{Name: "Flour", ID: 1},
		Ingredient{Name: "Milk", ID: 2},
		Ingredient{Name: "Egg", ID: 3},
	}
	//fmt.Printf("%+v\n", delicious)
	//duration := 5 * time.Second

	var cookie1 Cookie
	var cookie2 Cookie
	var cookie3 Cookie
	var cookieBaker Oven

	cookie1 = PrepareCookie(delicious)
	cookie2 = PrepareCookie(delicious)
	cookie3 = PrepareCookie(delicious)
	cookieBaker = CookieCutterOven{}

	//_ = cookieCreator.(Oven)

	cookieBaker.Store(cookie1)
	cookieBaker.Store(cookie2)
	cookieBaker.Store(cookie3)

	ovenBytes, ymlErr := yaml.Marshal(cookieBaker)
	if ymlErr != nil {
		panic(ymlErr)
	}
	fmt.Println(string(ovenBytes))

}
