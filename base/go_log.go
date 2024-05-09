package base

import (
	"io"
	"log"
	"os"
)

/*
InitLog
Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23，日志记录日期
Ltime                         // the time in the local time zone: 01:23:23，日志记录时间
Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.，日志记录微秒
Llongfile                     // full file name and line number: /a/b/c/d.go:23，绝对路径和行号
Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile，文件和行号
LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone，日期时间转为0时区
Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
LstdFlags     = Ldate | Ltime // initial values for the standard logger
*/
func InitLog() {
	log.SetPrefix("【GO_LOG】")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile | log.Lshortfile)
	log.Println("hello world")
	// fatal表示系统收到致命错误后，需要退出，这时使用fatal记录日志，然后程序退出。fatal = print + os.exit
	log.Fatalln("fatal error")
	// panic 表示先试用print记录日志，然后调用panic()函数抛出一个恐慌，这时候除非使用recover函数，否则程序会打印错误堆栈信息，然后程序终止
	log.Panicln("panic")
	/*
		由源码可以看出，日志默认输出到Stderr设备，对应unix里标准错误警告信息的输出设备，同时被作为默认的日志输出目的地。
		同时还有标准的输出设备Stdout和标准输入设备Stdin
		type Logger struct {
			outMu 	  sync.Mutex
			out       io.Writer 		     // destination for output
			prefix    atomic.Pointer[string] // prefix on each line to identify the logger (but see Lmsgprefix)
			flag      atomic.Int32           // properties
			isDiscard atomic.Bool
		}
		func (l *Logger) output(pc uintptr, calldepth int, appendOutput func([]byte) []byte) error {
			if l.isDiscard.Load() {
				return nil
			}

			now := time.Now() 		// get this early.

			// Load prefix and flag once so that their value is consistent within
			// this call regardless of any concurrent changes to their value.
			prefix := l.Prefix()
			flag := l.Flags()

			var file string
			var line int
			// 判断是否配置了文件和行号
			if flag&(Lshortfile|Llongfile) != 0 {
				// 这里不加锁，是因为runtime.Caller(calldepth)执行效率低，会影响打日志效率
				if pc == 0 {
					var ok bool
					// 获取运行时方法的调用信息, calldepth代表跳过栈帧数，0表示不跳过，1表示再往上一层，表示调用者的调用者，2就表示在什么地方调用的log.Print
					_, file, line, ok = runtime.Caller(calldepth)
					if !ok {
						file = "???"
						line = 0
					}
				} else {
					fs := runtime.CallersFrames([]uintptr{pc})
					f, _ := fs.Next()
					file = f.File
					if file == "" {
						file = "???"
					}
					line = f.Line
				}
			}
			// 将日志信息和设置的日志头进行拼接
			buf := getBuffer()
			defer putBuffer(buf)
			formatHeader(buf, now, prefix, flag, file, line)
			*buf = appendOutput(*buf)
			if len(*buf) == 0 || (*buf)[len(*buf)-1] != '\n' {
				*buf = append(*buf, '\n')
			}

			l.outMu.Lock()
			defer l.outMu.Unlock()
			// 输出拼接好的缓冲buf里的日志信息到目的地
			_, err := l.out.Write(*buf)
			return err
		}
	*/
}

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	WarnLog  *log.Logger
)

func MultiLog() {
	errFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败", err)
	}
	InfoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLog = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(io.MultiWriter(os.Stderr, errFile), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogExecute() {
	InfoLog.Println("InfoLog")
	WarnLog.Println("WarnLog")
	ErrorLog.Println("ErrorLog")
}
