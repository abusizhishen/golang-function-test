package main

import "time"

//参见https://frankhitman.github.io/zh-CN/gin-validator/
type Params struct {
	Day time.Time `json:"time" binding:"required" time_format:"2006-01-02"`
	Sex string `json:"sex" binding:"required,oneof=男 女"`
	Age int `json:"age" binding:"gte=18,lt=30"`
}
