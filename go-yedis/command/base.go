package command

import "Monica/go-yedis/core"

//通过key查询db中data，获取value值
func lookupKey(data core.Dict, userKey *core.YedisObject) (ret *core.YedisObject) {

	//TODO 无法dicts["key"]访问，不知道什么鬼
	for key, val := range data {
		if key == userKey.Ptr.(string) {
			return val
		}
	}
	return nil

}