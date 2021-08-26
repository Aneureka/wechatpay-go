// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微信支付分账API
//
// 微信支付分账API
//
// API version: 0.0.3

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package profitsharing_test

import (
	"context"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
)

func ExampleReceiversApiService_AddReceiver() {
	var (
		ctx    context.Context
		client *core.Client
	)
	// 假设已获得初始化后的 core.Client

	svc := profitsharing.ReceiversApiService{Client: client}
	resp, result, err := svc.AddReceiver(ctx,
		profitsharing.AddReceiverRequest{
			Account:        core.String("86693852"),
			Appid:          core.String("wx8888888888888888"),
			CustomRelation: core.String("代理商"),
			Name:           core.String("hu89ohu89ohu89o"),
			RelationType:   profitsharing.RECEIVERRELATIONTYPE_SERVICE_PROVIDER.Ptr(),
			SubAppid:       core.String("wx8888888888888889"),
			SubMchid:       core.String("1900000109"),
			Type:           profitsharing.RECEIVERTYPE_MERCHANT_ID.Ptr(),
		},
	)

	// TODO: 处理返回结果
	_, _, _ = resp, result, err
}

func ExampleReceiversApiService_DeleteReceiver() {
	var (
		ctx    context.Context
		client *core.Client
	)
	// 假设已获得初始化后的 core.Client

	svc := profitsharing.ReceiversApiService{Client: client}
	resp, result, err := svc.DeleteReceiver(ctx,
		profitsharing.DeleteReceiverRequest{
			Account:  core.String("86693852"),
			Appid:    core.String("wx8888888888888888"),
			SubAppid: core.String("wx8888888888888889"),
			SubMchid: core.String("1900000109"),
			Type:     profitsharing.RECEIVERTYPE_MERCHANT_ID.Ptr(),
		},
	)

	// TODO: 处理返回结果
	_, _, _ = resp, result, err
}
