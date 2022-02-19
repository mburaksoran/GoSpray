package main

import (
	// "fmt"
	"strings"
	"strconv"
)

type targetStruct struct{
	host string
	port int
}

func parseTarget(targetString string) *targetStruct {
	target := new(targetStruct)
	s := strings.Split(targetString, ":")
	target.host = s[0]
	target.port,_ = strconv.Atoi(s[1])
	return target
}