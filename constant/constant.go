package constant

var (
	Root  string = "/data"      // 工作目录 如果为空  默认/data
	To    string = "single-aac" // 转换到的编码 如果为空  默认vp9
	Level string = "Debug"      //日志的输出等级
)

func GetRoot() string {
	return Root
}
func SetRoot(s string) {
	Root = s
}

func GetTo() string {
	return To
}
func SetTo(s string) {
	To = s
}
func GetLevel() string {
	return Level
}
func SetLevel(s string) {
	Level = s
}
