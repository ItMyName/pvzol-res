package packet

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	rootPath   string
	hostDomain string
	ticker     *time.Ticker
	log        *os.File
)

// 设置请求包捕获文件存放到根目录 path 下，如果该文件夹以及文件夹下的log.txt文件创建失败，将会将导致恐慌。
// 可以将 clear 标志设置为 true，它会试图清空根目录下的所有文件，但它不会保证执行结果。
func SetRootPath(path string) {
	var err error
	if err = os.MkdirAll(path, 0777); err != nil {
		panic(errors.New("创建根目录时发生错误： " + err.Error()))
	}
	if log, err = os.OpenFile(filepath.Join(path, "log.txt"), os.O_CREATE|os.O_APPEND, 0666); err != nil {
		panic(errors.New("创建日志文件时发生错误：" + err.Error()))
	}

	rootPath = path
}

// 设置抓包的主机域名，该游戏拥有46个服务器（id：1~46），如果设置范围之外的数字将强制指向39。
func SetHostDomain(serverID int) {
	if serverID < 1 || serverID > 46 {
		serverID = 39
	}
	if serverID <= 11 {
		hostDomain = "http://pvz-s" + strconv.Itoa(serverID) + ".youkia.com/"
	} else {
		hostDomain = "http://s" + strconv.Itoa(serverID) + ".youkia.pvz.youkia.com/"
	}
}

// 设置请求间隔，单位毫秒
func SetRequestInterval(duration time.Duration) {
	if ticker == nil {
		ticker = time.NewTicker(duration * time.Millisecond)
	} else {
		ticker.Reset(duration * time.Millisecond)
	}
}

// 注册一个表示类型名称的文件夹，产生的错误信息会被发送到日志
func SignTypeDir(typeName string) (err error) {
	if err = os.MkdirAll(filepath.Join(rootPath, typeName), 0777); err != nil {
		err = errors.New("创建类型文件夹时发生错误： " + err.Error())
		HaveError(err)
	}
	return
}

// 便捷保存文件，会自动注册类型文件夹，存储状态会被记录
func KeepFile(name string, data []byte) (err error) {
	if err = os.MkdirAll(filepath.Join(rootPath, filepath.Dir(name)), 0777); err == nil {
		if err = writeFile(filepath.Join(rootPath, name), data); err == nil {
			fmt.Printf("√ %s\n", name)
			return
		}
	}
	fmt.Printf("\n× %s\n-> %s\n\n", name, err.Error())
	log.WriteString(fmt.Sprintf("× %s\n-> %s\n\n", name, err.Error()))
	return
}

// 添加一项错误解释，它不能保证成功添加词条
func HaveError(what error) {
	fmt.Printf("\n错误：%s\n\n", what.Error())
	if _, what = log.WriteString(fmt.Sprintf("\n错误：%s\n\n", what.Error())); what != nil {
		fmt.Println("-> 它未能被添加进日志文件！", what.Error())
	}
}

// 停止使用时调用，进行善后工作
func Close() (err error) {
	log.Close()
	return
}
