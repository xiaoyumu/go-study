package main

import (
	"fmt"
	"math"
	"reflect"
)

// PI The const
const PI float32 = 3.14159265

var name = "global variable"

func init() {
	fmt.Println("Init function is called.")
}

func testVariableScope(parameter int) {
	fmt.Println(name)
	// This won't work
	fmt.Printf("%d\r\n", parameter)
}

func main() {

	fmt.Println("Hello world.")
	fmt.Printf("%f\r\n", PI)

	var localVariable = 10
	testVariableScope(localVariable)

	fmt.Println()
	fmt.Println("Basic data types")
	fmt.Println("-----------------------------------------")

	// 声明变量和类型
	// var variableName type

	// 声明变量并用其初始值判断类型
	// variableName := value

	var intUint8 uint8
	intUint8 = 255 // uint8 最大值是 255, 超过这个数值将会溢出

	intUint8_2 := uint8(255)

	// 通过reflect反射包中的TypeOf函数可以获取变量的类型
	fmt.Printf("the type of variable intUint8_2 is '%s' \n", reflect.TypeOf(intUint8_2))
	// 输出 the type of variable intUint8_2 is 'uint8'

	var intUint16 uint16
	intUint16 = 65535 // uint16 max value is 65535

	var intUint32 uint32
	intUint32 = 4294967295 // uint32 max value is 4294967295

	var intUint64 uint64
	intUint64 = 18446744073709551615 // uint64 max value is 18446744073709551615

	fmt.Printf("Max value of intUint8:\t%d\r\n", intUint8)
	fmt.Printf("Max value of intUint16:\t%d\r\n", intUint16)
	fmt.Printf("Max value of intUint32:\t%d\r\n", intUint32)
	fmt.Printf("Max value of intUint64:\t%d\r\n", intUint64)
	fmt.Printf("Max value of intUint64:\t%d\r\n", uint64(math.Pow(2, 64))*2-1)

	// Array and slices
	fmt.Println()
	fmt.Println("Array and Slice")
	fmt.Println("-----------------------------------------")

	// 定义长度为3的uint数组，初始化每个元素值为0
	var uintArray [3]uint

	// 设置数组中特定元素的值
	uintArray[0] = 1
	uintArray[1] = 2
	uintArray[2] = 3
	fmt.Println(uintArray)
	fmt.Printf("Length of uintArray is %d\r\n", len(uintArray))

	// 定义一个长度为3的uint数组，并初始化index为1的元素值为2, index为2的元素值为3
	uintArray2 := [3]uint{1: 2, 2: 3}
	fmt.Println(uintArray2)
	// 输出: [0 2 3]

	uintArray3 := [4]uint{1: 2, 2: 3}
	// uintArray3 = uintArray2 不能将长度不同的数组变量相互复制，会导致编译错误。
	fmt.Println(uintArray3)

	// 指针数组
	var uintPointerArray [3]*uint
	uintPointerArray[0] = new(uint)
	uintPointerArray[1] = new(uint)
	uintPointerArray[2] = new(uint)

	fmt.Println(uintPointerArray)
	*uintPointerArray[0] = 10
	*uintPointerArray[1] = 20
	*uintPointerArray[2] = 30

	// 遍历指针数组获取每个指针元素指向地址中的值
	for _, v := range uintPointerArray {
		fmt.Printf("The value of address:%p is %d\r\n", v, *v)
	}

	// 创建一个长度为3的float32数组, [...] 表示根据初始化数组字面量来确定长度
	floatArray := [...]float32{1.1, 2.99, 3.999}
	fmt.Println(floatArray)

	// 创建一个长度为 3 ，容量为 3 的切片
	floatSlice := []float32{1.1, 2.99, 3.999}
	fmt.Printf("Length of Slice floatSlice (%s) is %d Capacity is %d\r\n", reflect.TypeOf(floatSlice), len(floatSlice), cap(floatSlice))

	// 创建一个长度为 3 ，容量为 3 的切片
	floatSlice2 := make([]float32, 3)
	fmt.Printf("Length of Slice floatSlice2 (%s) is %d Capacity is %d\r\n", reflect.TypeOf(floatSlice2), len(floatSlice2), cap(floatSlice2))

	// 创建一个长度为 5 ，容量为 10 的切片
	floatSlice3 := make([]float32, 5, 10)
	fmt.Printf("Length of Slice floatSlice3 (%s) is %d Capacity is %d\r\n", reflect.TypeOf(floatSlice3), len(floatSlice3), cap(floatSlice3))

	// 不能使用创建一个容量小于长度的切片，会导致编译错误
	// floatSlice3 := make([]float32, 5, 3)

	var intArray [10]int32 // 初始化一个长度为 10 的 int32 数组
	fmt.Println(intArray)
	// 输出: [0 0 0 0 0 0 0 0 0 0]
	intSlice := intArray[0:5] // 在现有的数组上创建切片，start index = 0, length = 5， 容量 = 10 - 0 = 10
	fmt.Printf("Length of Slice is %d Capacity is %d\r\n", len(intSlice), cap(intSlice))

	intSlice[0] = 1 // 修改slice成员的内容
	intSlice[1] = 2
	intSlice[2] = 3
	intSlice[3] = 4
	intSlice[4] = 5
	// intSlice[5] = 6 // 这样会导致索引越界的异常 panic: runtime error: index out of range
	fmt.Printf("intSlice = %v Its address is %p\r\n", intSlice, intSlice)
	// 输出修改后的slice内容
	// 输出: [1 2 3 4 5]
	fmt.Print("intArray = ")
	fmt.Println(intArray) // 输出slice所对应的底层数组的内容
	// 可以看到修改slice切片的元素内容，其实修改了其底层数组对应元素的内容
	// 输出: [1 2 3 4 5 0 0 0 0 0]

	alternativeSlice := intSlice[:] // 由于没有指定起始位置，长度和容量，直接返回intSlic
	// 所以可以看到 alternativeSlice 的地址与 intSlice 地址相同
	fmt.Printf("alternativeSlice[:] = %v Its address is %p\r\n", alternativeSlice, alternativeSlice)
	alternativeSlice2 := intSlice[1:2:3] // 创建一个新的切片, 长度为2 - 1 = 1 容量为3 - 1 = 2
	fmt.Printf("alternativeSlice2[1:2:3] = %v Its address is %p Length: %d Capacity: %d\r\n",
		alternativeSlice2,
		alternativeSlice2,
		len(alternativeSlice2),
		cap(alternativeSlice2))

	newSlice := append(intSlice, 6, 7)
	// 向intSlice切片添加两个元素的空间，初始化对应的数组元素值，并返回一个新的切片
	fmt.Print("newSlice = ")
	fmt.Println(newSlice)
	// 输出 [1 2 3 4 5 6 7 ]
	fmt.Printf("Length of Slice is %d Capacity is %d\r\n", len(newSlice), cap(newSlice))
	fmt.Println(intArray)
	// 输出 [1 2 3 4 5 6 7 0 0 0]

	brandNewSlice := append(newSlice, 8, 9, 10, 11)
	fmt.Print("brandNewSlice = ")
	fmt.Println(brandNewSlice)
	// 输出 brandNewSlice = [1 2 3 4 5 6 7 8 9 10 11]
	fmt.Print("intArray = ")
	fmt.Println(intArray)
	// 输出 intArray = [1 2 3 4 5 6 7 0 0 0]
	// 可以看到newSlice所对应的数组并没有发生改变，原因是因为append函数在为切片添加元素时，
	// 会检查切片的容量(可用空间)，如果空间不够，则会产生一个新的数组，并将当前被操作切片对应的数组元素复制到新的数组中
	// 新的切片长度等于len(newSlice) + n, 其中n为调用append函数时传入的要添加的元素的数量。
	// 新的切片的容量等于cap(newSlice) * 2
	fmt.Printf("Length of new slice is %d\r\n", len(newSlice)+4)
	fmt.Printf("Length of Slice is %d Capacity is %d\r\n", len(brandNewSlice), cap(brandNewSlice))

	// Loops
	// Iterate slice 迭代切片
	for index, value := range brandNewSlice {
		fmt.Printf("Index %d \t Value: %d\r\n", index, value)
	}

	// Iterate slice 迭代切片时，忽略index
	for _, value := range brandNewSlice {
		fmt.Printf("Value: %d\r\n", value)
	}

	// 倒序一个数组中的元素
	var myArray [10]int16
	for index := 0; index < len(myArray); index++ {
		myArray[index] = int16(index + 1)
	}
	fmt.Println(myArray)

	for index := 0; index < len(myArray); index++ {
		var pos = len(myArray) - index - 1
		if pos <= index {
			break
		}
		fmt.Printf("Index: %d POS: %d Value of POS: %d\r\n", index, pos, myArray[pos])
		myArray[index], myArray[pos] = myArray[pos], myArray[index]
	}
	fmt.Println(myArray)

	// Array数组按值传递，如果一个函数的参数是一个数组，传入数组将会被复制一份再赋值给实参。
	// Slice切片按引用传递，如果一个函数的参数是一个切片，传入的切片的地址将被赋值给实参。
	array := [3]uint{1, 2, 3}
	fmt.Printf("The address of array is %p\r\n", &array)
	functionThatTakeArray(array)

	slice := []uint{1, 2, 3}
	fmt.Printf("The address of slice is %p\r\n", slice)
	functionThatTakeSlice(slice)

	// Map
	fmt.Println()
	fmt.Println("Map")
	fmt.Println("-----------------------------------------")

	var nilMap map[string]string
	fmt.Println(nilMap)
	// nilMap["Key1"] = "value1" // panic: assignment to entry in nil map

	// 使用make创建一个map
	numberMap := make(map[int]string)
	numberMap[1] = "One"
	numberMap[2] = "Two"
	numberMap[3] = "Three"

	// 使用for循环遍历map
	for key, value := range numberMap {
		fmt.Printf("Key: %d Value: %s\r\n", key, value)
	}

	// 检查key 是否存在map中, 忽略key对应的值
	var exists bool
	_, exists = numberMap[4]
	if !exists {
		fmt.Println("Value of numberMap[4] does not exist.")
	}

	// 将key对应的值添加到map
	numberMap[4] = "Four"

	var value string
	value, exists = numberMap[4]
	if !exists {
		fmt.Println("Value of numberMap[4] does not exist.")
	} else {
		fmt.Printf("Value is %s \r\n", value)
	}
	_, exists = numberMap[3]
	if exists {
		fmt.Println("The key 3 exists in map.")
	}

	// 将Key对应的值从map中删除
	delete(numberMap, 3)
	_, exists = numberMap[3]
	if !exists {
		fmt.Println("The key 3 has been removed.")
	}

	fmt.Println()
	// 创建并初始化一个Map
	weekdayMap := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wensday",
		4: "Thusday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	fmt.Println(weekdayMap)

}

func functionThatTakeArray(value [3]uint) {
	fmt.Printf("Length of (%s)) is %d Capacity is %d Address is %p\r\n",
		reflect.TypeOf(value),
		len(value),
		cap(value),
		&value)
}

func functionThatTakeSlice(value []uint) {
	fmt.Printf("Length of (%s)) is %d Capacity is %d Address is %p\r\n",
		reflect.TypeOf(value),
		len(value),
		cap(value),
		value)
}
