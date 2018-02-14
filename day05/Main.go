package main

import (
	"fmt"
	"./DBUtil"
	"reflect"
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
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) returning id"

	insertResult, err := DBUtil.InsertMap(insertSQL, []interface{}{
		[]interface{}{202891, "福建省福州市鼓楼区中军后6号7座806单元 - t4", "350104", "仓山区", "350102004001", "军门社区", "136376", nil, nil, nil, "350102004000", "东街街道"},
		[]interface{}{202892, "福建省福州市鼓楼区中军后6号7座806单元 - t5", "350104", "仓山区", "350102004001", "军门社区", "136376", nil, nil, nil, "350102004000", "东街街道"},
		[]interface{}{202893, "福建省福州市鼓楼区中军后6号7座806单元 - t6", "350104", "仓山区", "350102004001", "军门社区", "136376", nil, nil, nil, "350102004000", "东街街道"},
		[]interface{}{202894, "福建省福州市鼓楼区中军后6号7座806单元 - t7", "350104", "仓山区", "350102004001", "军门社区", "136376", nil, nil, nil, "350102004000", "东街街道"},
	})
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	for _, id := range insertResult {
		val := reflect.ValueOf(id)
		fmt.Println(val.MethodByName("LastInsertId").Call(nil)[0].Interface())
	}

	updateSQL := "UPDATE public.address_engine set address = $1 WHERE id = $2"
	update, err := DBUtil.UpdateMap(updateSQL, []interface{}{
		[]interface{}{"测试地址", 202891},
		[]interface{}{"测试地址", 202892},
		[]interface{}{"测试地址", 202893},
		[]interface{}{"测试地址", 202894},
	})
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(update)

	deleteSQL := "DELETE FROM public.address_engine WHERE id = $1"
	del, err := DBUtil.UpdateMap(deleteSQL, []interface{}{
		[]interface{}{202891},
		[]interface{}{202892},
		[]interface{}{202893},
		[]interface{}{202894},
	})
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(del)

}
