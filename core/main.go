package main

import (
	adbt "./pkg"
)

func main() {
	m := adbt.MongoDB{
		URI:  "mongodb://192.168.0.106:27017/discovery",
		Name: "CVPM-TEST",
	}
	m.Backup()
}
