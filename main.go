package main

import "log"

type Material struct {
	Label string `json:"label"`
	Name  struct {
		En string `json:"en"`
		Ja string `json:"ja"`
	} `json:"name"`
}

type Operator struct {
	Name struct {
		En string `json:"en"`
		Ja string `json:"ja"`
	} `json:"name"`
	Class      string `json:"class"`
	Promotions struct {
		Elite1 struct {
			Materials []struct {
				Label string `json:"label"`
				Count int    `json:"count"`
			} `json:"materials"`
		} `json:"elite1"`
		Elite2 struct {
			Materials []struct {
				Label string `json:"label"`
				Count int    `json:"count"`
			} `json:"materials"`
		} `json:"elite2"`
	} `json:"promotions"`
	Skills []struct {
		Number int `json:"number"`
		Name   struct {
			En string `json:"en"`
			Ja string `json:"ja"`
		} `json:"name"`
		Upgrades []struct {
			Materials []struct {
				Label string `json:"label"`
				Count int    `json:"count"`
			} `json:"materials"`
		} `json:"upgrades"`
	} `json:"skills"`
}

func main() {
	log.Println("It Works!")
}
