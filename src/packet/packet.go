package packet

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// 有的数据包状态码200，但会返回这种文件……
//go:embed static/invalid_file_example.xml
var invalidFile []byte

type Packet struct {
	Url  string
	Name string
	What error
}

// 该请求包资源已经妥善处理处理时调用
func (p *Packet) ok() {
	fmt.Printf("√ %s @%s\n", p.Name, p.Url)
}

// 该请求包资源处理中出现错误时调用
func (p *Packet) err() {
	fmt.Printf("\n× %s @%s\n-> %s\n\n", p.Name, p.Url, p.What.Error())
	log.WriteString(fmt.Sprintf("× %s @%s\n-> %s\n\n", p.Name, p.Url, p.What.Error()))
}

// 使用Get请求获取数据文件并保存
func (p *Packet) Get() (data []byte) {
	var resp *http.Response
	req, _ := http.NewRequest("GET", hostDomain+p.Url, nil)
	<-ticker.C
	resp, p.What = http.DefaultClient.Do(req)
	if p.What != nil {
		p.err()
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		p.What = errors.New("StatusCode " + resp.Status)
		p.err()
		return
	}
	if data, p.What = io.ReadAll(resp.Body); p.What != nil {
		p.err()
		return
	}
	if bytes.Equal(data, invalidFile) {
		p.What = errors.New("is an invalid file, the url may have no file")
		p.err()
		return
	}
	if p.What = writeFile(filepath.Join(rootPath, p.Name), data); p.What != nil {
		p.err()
		return
	}

	p.ok()
	return
}

// 对简单的写入操作进行封装，但是该函数不允许对已经存在的文件进行操作
func writeFile(name string, data []byte) (err error) {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return
}
