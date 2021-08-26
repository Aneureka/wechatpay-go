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

func ExampleMerchantsApiService_QueryMerchantRatio() {
	var (
		ctx    context.Context
		client *core.Client
	)
	// 假设已获得初始化后的 core.Client

	svc := profitsharing.MerchantsApiService{Client: client}
	resp, result, err := svc.QueryMerchantRatio(ctx,
		profitsharing.QueryMerchantRatioRequest{
			SubMchid: core.String("1900000109"),
		},
	)

	// TODO: 处理返回结果
	_, _, _ = resp, result, err
}
