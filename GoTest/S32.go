package GoTest

//导入需要的包	import testing (如果你是goland,那么可以忽略，因为ide就自动帮你加上)
//执行命令	go test file_test.go
//测试文件命名	必须以_test.go结尾
//功能测试的用力函数	必须以Test开头&&后头跟着的函数名不能以小写字母开头，比如：Testcbs 就是不行的，TestCbs就是ok的
//功能测试参数	testing.T
//压力测试用例函数	必须以Benchmark开头&&其后的函数名不能以小写字母开头(例子同上)
//压力测试参数	testing.B
//测试信息	.Log方法，默认情况下是不会显示的，只有在go test -v的时候显示
//测试控制	通过Error/Errorf/FailNow/Fatal等来进行测试是否是失败，或者在失败的情况下的控制
//压力测试命令	go test -test.bench file_test.go
//压力测试的循环体	使用test.B.N

func GetArea(weight int, height int) int {
	return weight*height
}