package main

import (
	"fmt"
	"./DBUtil"
	"strconv"
)

func main() {
	// 初始化链接
	DBUtil.InitConnectionInfo("centos", "addressengine", "user123", "kini")
	// 调用查询方法 -- 封装过 返回 map对象
	mapResult, err := DBUtil.QueryInfo("SELECT * FROM address_engine LIMIT $1", 10)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	for _, m := range mapResult {
		fmt.Println(m)
	}

	insertSQL := "INSERT INTO public.address_engine(id, address, area_id, area_name, community_id," +
		" community_name, gid, latitude, longitude, num, street_id, street_name) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"

	insertResult, err := DBUtil.InsertMap(insertSQL, []map[string]interface{}{
		{	"id": "202890",
			"address": "福建省福州市鼓楼区中军后6号7座806单元 - t3",
			"area_id": "50104",
			"area_name": "仓山区",
			"community_id": "350102004001",
			"community_name": "军门社区",
			"gid": "136376",
			"latitude": nil,
			"longitude": nil,
			"num": nil,
			"street_id": "350102004000",
			"street_name": "东街街道"},
		{	"id": "202891",
			"address": "福建省福州市鼓楼区中军后6号7座806单元 - t4",
			"area_id": "50104",
			"area_name": "仓山区",
			"community_id": "350102004001",
			"community_name": "军门社区",
			"gid": "136376",
			"latitude": nil,
			"longitude": nil,
			"num": nil,
			"street_id": "350102004000",
			"street_name": "东街街道"},
	})

	if nil != err {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("instert sql -> " + strconv.FormatBool(insertResult))
}
