// package main

// import (
// 	"fmt"
// 	// "math"
// )

// 返回a、b中最大值.
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// //返回 A+B 和 A*B
// func SumAndProduct(A, B int) (int, int) {
// 	return A+B, A*B
// }

//简单的一个函数，实现了参数+1的操作
// func add1(a int) int {
// 	a = a+1 // 我们改变了a的值
// 	return a //返回一个新值
// }

//简单的一个函数，实现了参数+1的操作
// func add1(a *int) int { // 请注意，
// 	*a = *a+1 // 修改了a的值
// 	return *a // 返回新值
// }

// type testInt func(int) bool // 声明了一个函数类型
// func isOdd(integer int) bool {
// 	if integer%2 == 0 {
// 		return false
// 	}
// 	return true
// }
// func isEven(integer int) bool {
// 	if integer%2 == 0 {
// 		return true
// 	}
// 	return false
// }
// 声明的函数类型在这个地方当做了一个参数
// func filter(slice []int, f testInt) []int {
// 	var result []int
// 	for _, value := range slice {
// 		if f(value) {
// 			result = append(result, value)
// 		}
// 	}
// 	return result
// }

// // 声明一个新的类型
// type person struct {
// 	name string
// 	age int
// }
// // 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// // struct也是传值的
// func Older(p1, p2 person) (person, int) {
// 	if p1.age>p2.age {  // 比较p1和p2这两个人的年龄
// 		return p1, p1.age-p2.age
// 	}
// 	return p2, p2.age-p1.age
// }


// func printSlice(x []int) {
// 	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
// }

// type Human struct {
// 	name string
// 	age int
// 	weight int
// }
// type Student struct {
// 	Human
// 	spec string
// }

// type 	Rectangle struct {
// 	width, height float64
// }
// type Circle struct {
// 	radius float64
// }
// func (r Rectangle) area() float64 {
// 	return r.width*r.height
// }
// func (c Circle) area() float64 {
// 	return c.radius * c.radius * math.Pi
// }


// const(
// 	WHITE = iota
// 	BLACK
// 	BLUE
// 	RED
// 	YELLOW
// )
// type Color byte
// type Box struct {
// 	width, height, depth float64
// 	color Color
// }
// type BoxList []Box //a slice of boxes
// func (b Box) Volume() float64 {
// 	return b.width * b.height * b.depth
// }
// func (b *Box) SetColor(c Color) {
// 	b.color = c
// }
// func (bl BoxList) BiggestColor() Color {
// 	v := 0.00
// 	k := Color(WHITE)
// 	for _, b := range bl {
// 		if bv := b.Volume(); bv > v {
// 			v = bv
// 			k = b.color
// 		}
// 	}
// 	return k
// }
// func (bl BoxList) PaintItBlack() {
// 	for i := range bl {
// 		bl[i].SetColor(BLACK)
// 	}
// }
// func (c Color) String() string {
// 	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
// 	return strings[c]
// }


// func main() {
	// var numbers map[string]int
	// // 另一种map的声明方式
	// numbers = make(map[string]int)
	// // numbers["one"] = 1  //赋值
	// numbers["ten"] = 10 //赋值
	// numbers["three"] = 3
	// numbers["one"]=11
	// fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据


	// 初始化一个字典
	// rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
	// // map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	// csharpRating, ok := rating["C#"]
	// if ok {
	// 	fmt.Println("C# is in the map and its rating is ", csharpRating)
	// } else {
	// 	fmt.Println("We have no rating associated with C# in the map")
	// }
	// delete(rating, "C")  // 删除key为C的元素
	// fmt.Println(rating)

	// m := make(map[string]string)
	// m["Hello"] = "Bonjour"
	// m1 := m
	// m1["Hello"] = "Salut"  // 现在m["hello"]的值已经是Salut了
	// fmt.Println(m1)

	// var x int = 15
	// if x > 10 {
	// 	fmt.Printf("big")
	// } else {
	// 	fmt.Printf("small")
	// }

	// i := 0
	// Here:   //这行的第一个词，以冒号结束作为标签
	// 	println(i)
	// 	i++
	// 	goto Here   //跳转到Here去

	// sum := 0;
	// for index:=0; index < 10 ; index++ {
	// 	sum += index
	// }

	// sum := 1
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println("sum is equal to ", sum)


	// for index := 10; index>0; index-- {
	// 	if index == 5{
	// 		//或者 break 
	// 		continue
	// 	}
	// 	fmt.Println(index)
	// }
	// break打印出来10、9、8、7、6
	// continue打印出来10、9、8、7、6、4、3、2、1

	// for k,v:=range map {
	// 	fmt.Println("map's key:",k)
	// 	fmt.Println("map's val:",v)
	// }

	// for _, v := range map{
	// 	fmt.Println("map's val:", v)
	// }

	// i := 10
	// switch i {
	// case 1:
	// 	fmt.Println("i is equal to 1")
	// case 2, 3, 4:
	// 	fmt.Println("i is equal to 2, 3 or 4")
	// case 10:
	// 	fmt.Println("i is equal to 10")
	// default:
	// 	fmt.Println("All I know is that i is an integer")
	// }


	// integer := 6
	// switch integer {
	// case 4:
	// 	fmt.Println("The integer was <= 4")
	// 	fallthrough
	// case 5:
	// 	fmt.Println("The integer was <= 5")
	// 	fallthrough
	// case 6:
	// 	fmt.Println("The integer was <= 6")
	// 	fallthrough
	// case 7:
	// 	fmt.Println("The integer was <= 7")
	// 	fallthrough
	// case 8:
	// 	fmt.Println("The integer was <= 8")
	// 	fallthrough
	// default:
	// 	fmt.Println("default case")
	// }

	// x := 3
	// y := 4
	// z := 5

	// max_xy := max(x, y) //调用函数max(x, y)
	// max_xz := max(x, z) //调用函数max(x, z)

	// fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	// fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	// fmt.Printf("max(%d, %d) = %d\n", y, z, max(y,z)) // 也可在这直接调用它

	// x := 3
	// y := 4
	// xPLUSy, xTIMESy := SumAndProduct(x, y)
	// fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	// fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)

	  // x := 3
		// fmt.Println("x = ", x)  // 应该输出 "x = 3"
		// x1 := add1(x)  //调用add1(x)
		// fmt.Println("x+1 = ", x1) // 应该输出"x+1 = 4"
		// fmt.Println("x = ", x)    // 应该输出"x = 3"

		  // x := 3
			// fmt.Println("x = ", x)  // 应该输出 "x = 3"
			// x1 := add1(&x)  // 调用 add1(&x) 传x的地址
			// fmt.Println("x+1 = ", x1) // 应该输出 "x+1 = 4"
			// fmt.Println("x = ", x)    // 应该输出 "x = 4"
	//defer语句
	// func ReadWrite() bool {
	// 	file.Open("file")
	// // 做一些工作
	// 	if failureX {
	// 		file.Close()
	// 		return false
	// 	}
	
	// 	if failureY {
	// 		file.Close()
	// 		return false
	// 	}
	
	// 	file.Close()
	// 	return true
	// } 
	// func ReadWrite() bool {
	// 	file.Open("file")
	// 	defer file.Close()
	// 	if failureX {
	// 		return false
	// 	}
	// 	if failureY {
	// 		return false
	// 	}
	// 	return true
	// }

	// for i := 0; i < 5; i++ {
	// 	defer fmt.Printf("%d ", i)// 4 3 2 1 
	// }

	// slice := []int {1, 2, 3, 4, 5, 7}
	// fmt.Println("slice = ", slice)
	// odd := filter(slice, isOdd)    // 函数当做值来传递了
	// fmt.Println("Odd elements of slice are: ", odd)
	// even := filter(slice, isEven)  // 函数当做值来传递了
	// fmt.Println("Even elements of slice are: ", even)

	// var tom person
	// 	// 赋值初始化
	// 	tom.name, tom.age = "Tom", 18
	// 	// 两个字段都写清楚的初始化
	// 	bob := person{age:25, name:"Bob"}
	// 	// 按照struct定义顺序初始化值
	// 	paul := person{"Paul", 43}
	// 	tb_Older, tb_diff := Older(tom, bob)
	// 	tp_Older, tp_diff := Older(tom, paul)
	// 	bp_Older, bp_diff := Older(bob, paul)
	// 	fmt.Printf("Of %s and %s, %s is older by %d years\n",
	// 		tom.name, bob.name, tb_Older.name, tb_diff)
	// 	fmt.Printf("Of %s and %s, %s is older by %d years\n",
	// 		tom.name, paul.name, tp_Older.name, tp_diff)
	// 	fmt.Printf("Of %s and %s, %s is older by %d years\n",
	// 		bob.name, paul.name, bp_Older.name, bp_diff)

// Slice
	/* 创建切片 */
	// 	numbers := [] int{0,1,2,3,4,5,6,7,8}
	// 	printSlice(numbers)
	// 	 /* 打印子切片从索引1(包含) 到索引4(不包含)*/
	// 	 fmt.Println("numbers[1:4]==", numbers[1:4])
	// 	 /* 默认下限为 0*/
	// 	 fmt.Println("numbers[:3]==",numbers[:3])
	// 	 /* 默认上限为 len(s)*/
	// 	 fmt.Println("numbers[4:]==",numbers[4:])
	// 	 numbers1 := make([]int,0,5)
	// 	 printSlice(numbers1)
	// 	 /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	// 	 numbers2 := numbers[:2]
	// 	 printSlice(numbers2)
	// 	 /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	// 	 numbers3 := numbers[2:5]
	// 	 printSlice(numbers3)
	// // append和copy
	// 	 var nub []int
	// 	 printSlice(nub)
	// 	 /* 允许追加空切片 */
	// 	 nub = append(nub, 0)
	// 	 printSlice(nub)
	// 	 /* 向切片添加一个元素 */
	// 	 nub = append(nub, 1)
	// 	 printSlice(nub)
	// 	 /* 同时添加多个元素 */
	// 	 nub = append(nub, 2,3,4)
	// 	 printSlice(nub)
	// 	 /* 创建切片 numbers1 是之前切片的两倍容量*/
	// 	 nub1 := make([]int, len(nub), (cap(nub))*2)
	// 	 printSlice(nub1)
	// 	 /* 拷贝 numbers 的内容到 numbers1 */
	// 	 copy(nub1,nub)
	// 	 printSlice(nub1)


/*struct*/ 
	// // 我们初始化一个学生
	// mark := Student{Human{"Mark",25,120}, "Computer Science"}
	// // 我们访问相应的字段
	// fmt.Println("His name is", mark.name)
	// fmt.Println("His age is", mark.age)
	// fmt.Println("His spec is", mark.spec)
	// // 修改对应的备注信息
	// mark.spec = "IT"
	// fmt.Println("His spec is", mark.spec)
	// // 修改他的年龄信息
	// mark.age = 20
	// fmt.Println("His age is", mark.age)
	// // 修改他的体重信息
	// mark.weight += 80
	// fmt.Println("His weight is", mark.weight)

/*面向对象*/ 
  // 长方形
	// r1 := Rectangle{12, 2}
	// r2 := Rectangle{9, 4}
	// fmt.Println("Area of r1 is：", r1.area())
	// fmt.Println("Area of r2 is：", r2.area())
	// // 圆形
	// c1 := Circle{10}
	// c2 := Circle{25}
	// fmt.Println("Area of c1 is：", c1.area())
	// fmt.Println("Area of c1 is：", c2.area())


/*复杂的method*/ 
	// boxes := BoxList {
	// 	Box{4, 4, 4, RED},
	// 	Box{10, 10, 1, YELLOW},
	// 	Box{1, 1, 20, BLACK},
	// 	Box{10, 10, 1, BLUE},
	// 	Box{10, 30, 1, WHITE},
	// 	Box{20, 20, 20, YELLOW},
	// }

	// fmt.Printf("We have %d boxes in our set\n", len(boxes))
	// fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	// fmt.Println("The color of the last one is",boxes[len(boxes)-1].color.String())
	// fmt.Println("The biggest one is", boxes.BiggestColor().String())

	// fmt.Println("Let's paint them all black")
	// boxes.PaintItBlack()
	// fmt.Println("The color of the second one is", boxes[1].color.String())

	// fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())

	
// }