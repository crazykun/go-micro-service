package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

// 写入日志
func (fl fileLog) Write(data []byte) (int, error) {
	// 文件不存在时候创建, 写入, 附加
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

// 写入的文件地址, 自定义log
func Run(destination string) {
	// 前缀go 写入时间
	log = stlog.New(fileLog(destination), "go: ", stlog.LstdFlags)
}

// http请求处理
func RegisterHandler() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := ioutil.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			writer(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

// 调用自定义log写入日志
func writer(message string) {
	log.Printf("%v\n", message)
}
