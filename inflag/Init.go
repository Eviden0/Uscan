package inflag

import (
	"flag"
	"fmt"
)

/*
UA头
最大爬取链接数
超时数,默认为5
状态码
代理配置
线程数,默认为5
*/
var (
	U  string
	UA string
	MU int
	TI int
	SC string
	PX string
	TH = 50
)

func Init() {
	//添加使用说明
	flag.Usage = func() {
		fmt.Println("使用方法: [选项]")
		fmt.Println("选项:")
		flag.PrintDefaults() // 打印默认的帮助信息
	}
	flag.StringVar(&U, "url", "", "set the target url")
	flag.StringVar(&UA, "ua", "", "set user-agent")
	flag.IntVar(&MU, "maxurls", 999, "set max urls")
	flag.IntVar(&TI, "timeout", 5, "set time out")
	flag.StringVar(&SC, "statuscode", "", "set code")
	flag.StringVar(&PX, "proxy", "", "set proxy url")
	flag.IntVar(&TH, "thread", 50, "set thread")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░░▒▓███████▓▒░░▒▓██████▓▒░ ░▒▓██████▓▒░░▒▓███████▓▒░  \n░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ \n░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ \n░▒▓█▓▒░░▒▓█▓▒░░▒▓██████▓▒░░▒▓█▓▒░      ░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░ \n░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ \n░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ \n ░▒▓██████▓▒░░▒▓███████▓▒░ ░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ \n                                                         𝔹𝕪:𝔼𝕧𝕚𝕕𝕖𝕟           \n                                                                   ")
	flag.Parse()
}
