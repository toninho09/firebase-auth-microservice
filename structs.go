package main

type BaseResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type VerifyStruct struct{
	Token string `json:"token" binding:"required"`
}