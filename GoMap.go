package main

import "fmt"

/*
map_variable := make(map[KeyType]ValueType, initialCapacity)
*/
func main() {

	var transMap map[string]string
	transMap = make(map[string]string)

	// 插入 key - value 对
	transMap["Google"] = "谷歌"
	transMap["Baidu"] = "百度"
	transMap["Wiki"] = "维基百科"

	fmt.Println(len(transMap))

	// 使用键输出map值
	for trans := range transMap {
		fmt.Println(trans, "中文是", transMap[trans])
	}
	// range获取k-v
	fmt.Println("=====")
	printMap(transMap)

	/*
		查看元素在集合中是否存在
	*/
	name, ok := transMap["Facebook"] /*如果ok是true,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("Facebook 的中文是", name)
	} else {
		fmt.Println("Facebook 不存在")
	}

	/*
		删除元素
	*/
	fmt.Println("====")
	delete(transMap, "Baidu")
	printMap(transMap)
}

func printMap(transMap map[string]string) {
	for k, v := range transMap {
		fmt.Printf("key:%s: value:%s\n", k, v)
	}
}
