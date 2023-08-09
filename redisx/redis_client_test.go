package redisx

import (
	"fmt"
	"testing"
)

func TestSetGet(t *testing.T) {
	err := client.Set("wodoma", "nihao2", 100)
	if err != nil {
		t.Error(err)
	}

	m, err := client.IsExist("333")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(m)

	dd, err := client.IsExist("wodoma")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(dd)

	// value, err := client.Get("sxf")
	// if err != nil {
	// 	t.Error(err)
	// }
	// fmt.Println(value)
	//
	// if value != "nihao" {
	// 	t.Error("value is not nihao")
	// }
}

func TestMGetMSet(t *testing.T) {
	ar := NewArrayReq(4)
	err := ar.Add("sxf2", "nihao")
	if err != nil {
		t.Error(err)
	}
	err = ar.Add("sxf3", "nihao1")
	if err != nil {
		t.Error(err)
	}
	err = client.MSet(ar)
	if err != nil {
		t.Error(err)
	}

	values, err := client.MGet("sxf2", "sxf3")
	if err != nil {
		t.Error(err)
	}

	if values[0] != "nihao" || values[1] != "nihao1" {
		t.Error("values is not nihao nihao2")
	}
}

func TestDel(t *testing.T) {
	value, err := client.Get("sxf")
	if err != nil {
		t.Error(err)
	}

	if value != "nihao" {
		t.Error("value is not nihao")
	}

	err = client.Del("sxf")
	if err != nil {
		t.Error(err)
	}
	_, err = client.Get("sxf")
	if err == nil {
		t.Error(err)
	}
}

func TestHSetHGet(t *testing.T) {
	err := client.HSet("mykey", "name", "sxf")
	if err != nil {
		t.Error(err)
	}

	value, err := client.HGet("mykey", "name")
	if err != nil {
		t.Error(err)
	}

	if value != "sxf" {
		t.Error("value is not sxf")
	}
}

func TestHMGetHMSet(t *testing.T) {
	ar := NewArrayReq(4)
	err := ar.Add("name", "sxf")
	if err != nil {
		t.Error(err)
	}
	err = ar.Add("age", "18")
	if err != nil {
		t.Error(err)
	}
	err = client.HMSet("mykey2", ar)
	if err != nil {
		t.Error(err)
	}

	values, err := client.HMGet("mykey2", "name", "age")
	if err != nil {
		t.Error(err)
	}

	if values[0] != "sxf" || values[1] != "18" {
		t.Error("values is not sxf 18")
	}
}

func TestHGetAll(t *testing.T) {
	ar := NewArrayReq(4)
	err := ar.Add("quality", 1)
	if err != nil {
		t.Error(err)
	}
	err = ar.Add("level", 4)
	if err != nil {
		t.Error(err)
	}
	err = ar.Add("token_id", 3007)
	if err != nil {
		t.Error(err)
	}
	err = ar.Add("uid", 3768890)
	if err != nil {
		t.Error(err)
	}
	err = client.HMSet("microphone_3007", ar)
	if err != nil {
		t.Error(err)
	}
	bb := NewArrayReq(4)
	err = bb.Add("quality", 2)
	if err != nil {
		t.Error(err)
	}
	err = bb.Add("level", 30)
	if err != nil {
		t.Error(err)
	}
	err = bb.Add("token_id", 3008)
	if err != nil {
		t.Error(err)
	}
	err = bb.Add("uid", 3768890)
	if err != nil {
		t.Error(err)
	}
	err = client.HMSet("microphone_3008", bb)
	if err != nil {
		t.Error(err)
	}
	cc := NewArrayReq(4)
	err = cc.Add("3008", 1)
	if err != nil {
		t.Error(err)
	}
	err = cc.Add("3009", 1)
	if err != nil {
		t.Error(err)
	}
	client.HMSet("microphone_user_3765567", cc)
	// values, err := client.HGetAll("mykey3")
	// if err != nil {
	// 	t.Error(err)
	// }
	//
	// if values["name"] != "sxf" || values["age"] != "18" {
	// 	t.Error("values is not sxf 18")
	// }

}

func TestHDel(t *testing.T) {
	value, err := client.HGet("mykey3", "name")
	if err != nil {
		t.Error(err)
	}

	if value != "sxf" {
		t.Error("value is not sxf")
	}

	err = client.HDel("mykey3", "name")
	if err != nil {
		t.Error(err)
	}
	_, err = client.HGet("mykey3", "name")
	if err == nil {
		t.Error(err)
	}
}

func TestGoRedis_HisExist(t *testing.T) {
	err := client.HSet("xiaoming", "age", 10)
	if err != nil {
		t.Error(err)
		return
	}
	err = client.HSet("xiaoming", "sex", 1)
	if err != nil {
		t.Error(err)
		return
	}
	ok, err := client.HisExist("xiaoming", "nihao")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ok)
}

func TestKeys(t *testing.T) {
	res, err := client.Keys("cache*")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}

func TestZAdd(t *testing.T) {
	ar := NewZAddReq()
	ar.Add(9, "xioaming")
	ar.Add(2, "xiaohong")
	ar.Add(6, "xiaojun")
	val, err := client.ZAdd("mykey4", ar)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(val)
	res, err := client.ZRange("mykey4", 0, -1)
	fmt.Println(res)
}

func TestZRevRange(t *testing.T) {
	ar := NewZAddReq()
	ar.Add(9, "xioaming")
	ar.Add(2, "xiaohong")
	ar.Add(6, "xiaojun")
	val, err := client.ZAdd("zkey1", ar)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(val)
	res, err := client.ZRevRangeWithScores("mykey4", 0, -1)
	fmt.Println(res)
}
