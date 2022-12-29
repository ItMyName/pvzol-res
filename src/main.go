package main

import (
	"PvzCapture/src/packet"
	capturedfiles "PvzCapture/src/packet/captured_files"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

func main() {
	capturedfiles.MainSwf.Get()

	keep := make(chan *packet.Packet, 2)
	go func() {
		for pack := range keep {
			pack.Get()
		}
	}()
	var wait sync.WaitGroup
	capturedfiles.InfoFileAndOrgIcon(keep, &wait)
	capturedfiles.UIIndexFileAndSwf(keep, &wait)
	wait.Wait()
	if err := packet.Close(); err != nil {
		fmt.Println("====================\n！！！日志信息未能妥善保存！！！\n====================")
	}
}

func init() {
	// 默认设置：
	// 根目录     "./Resource"  标志：-o
	// 服务器区号  39			 标志：-s
	// 请求间隔    300ms         标志：-d
	// 不对数据文件进行分类			标志：unclass
	// 当根目录文件夹已经存在时不会会尝试清空文件 标志：del
	// 请求间隔不宜过短，推荐在200ms以上
	var (
		rootPath string        = "Resource"
		sid      int           = 39
		duration time.Duration = 300
		del      bool          = false
		unclass  bool          = false
	)
	for _, arg := range os.Args[1:] {
		if arg == "del" {
			del = true
		} else if arg == "unclass" {
			unclass = true
		}
		if len(arg) < 3 {
			continue
		}
		switch arg[:2] {
		case "-o":
			fmt.Println("121")
			rootPath = arg[2:]
		case "-s":
			if id, err := strconv.Atoi(arg[2:]); err == nil {
				sid = id
			}
		case "-d":
			if d, err := strconv.Atoi(arg[2:]); err == nil {
				duration = time.Duration(d)
			}
		}

	}
	if del {
		if _, err := os.Stat(rootPath); !errors.Is(err, fs.ErrNotExist) {
			files, err := os.ReadDir(rootPath)
			if err != nil {
				panic(fmt.Errorf("已存在根目录，但无法读取根目录文件：%s", err.Error()))
			}

			fmt.Println("该目录的以下文件或文件夹将会被删除：")
			for _, f := range files {
				fmt.Println(filepath.Join(rootPath, f.Name()))
			}
			var t string
			fmt.Print("是否继续（yes/y）：")
			fmt.Scanf("%s", &t)
			if t == "yes" || t == "y" {
				var haveErr = false
				for _, f := range files {
					if err := os.RemoveAll(filepath.Join(rootPath, f.Name())); err != nil {
						fmt.Printf("× %s\n-> %s\n", filepath.Join(rootPath, f.Name()), err.Error())
						haveErr = true
					} else {
						fmt.Printf("√ %s\n", filepath.Join(rootPath, f.Name()))
					}
				}
				if haveErr {
					fmt.Print("可能存在未被删除的文件或文件夹，是否继续（yes/y）：")
					fmt.Scanf("%s", &t)
					if t != "yes" && t != "y" {
						os.Exit(0)
					}
				}
			} else {
				os.Exit(0)
			}
		}
	}
	if unclass {
		packet.Unclassified()
	}
	packet.SetRootPath(rootPath)
	packet.SetHostDomain(sid)
	packet.SetRequestInterval(duration)
}
