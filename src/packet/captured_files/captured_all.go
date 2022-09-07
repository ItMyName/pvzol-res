package capturedfiles

import (
	"PvzCapture/src/packet"
	"path/filepath"
	"sync"
)

// 获取所有的UI索引.xml，以及生成UI资源.swf请求包发送到keep管道中，并且会将xml文件转化成json文件并保存
// xml文件成功转化成json文件后保存，注册类型文件夹失败将放弃该部分文件，所有错误均发送到日志文件中
// 注册类型文件夹失败将放弃该部分文件，所有错误均发送到日志文件中，请完成packet包的初始化后再调用，并在停止时调用Close
func UIIndexFileAndSwf(keep chan<- *packet.Packet, wait *sync.WaitGroup) {
	wait.Add(1)
	if packet.SignTypeDir(filepath.Join("UI")) != nil {
		return
	}
	configPack, uiConfigPack, other, err := XmlUIIndexFilePackets()
	if err != nil {
		return
	}
	if data := configPack.Get(); configPack.What == nil {
		if v, err := Unmarshal(configPack, data); err == nil {
			config := v.(*ConfighRoot)
			wait.Add(1)
			go func() {
				BaseUI(config, "", keep)
				UISWF(&config.UIIndexRoot, "UI", keep)
				wait.Done()
			}()
		}
	}

	if data := uiConfigPack.Get(); uiConfigPack.What == nil {
		if v, err := Unmarshal(uiConfigPack, data); err == nil {
			ui_config := v.(*UIConfigRoot)
			go func() {
				UIConfigSWF(ui_config, "UI", keep)
				wait.Done()
			}()
		}
	}
	wait.Add(1)
	go func() {
		for _, pack := range other {
			if data := pack.Get(); pack.What == nil {
				if v, err := Unmarshal(pack, data); err == nil {
					uiindex := v.(*UIIndexRoot)
					UISWF(uiindex, "UI", keep)
				}
			}
		}
		wait.Done()
	}()
}

// 获取所有信息文件.xml，以及生成organism与tool附带的图标和生物立绘文件.swf请求包发送到keep管道中
// xml文件成功转化成json文件后保存，注册类型文件夹失败将放弃该部分文件，所有错误均发送到日志文件中
// 请完成packet包的初始化后再调用，并在停止时调用Close
func InfoFileAndOrgIcon(keep chan<- *packet.Packet, wait *sync.WaitGroup) {
	orgPack, toolPack, other, err := XmlInfoFilePackets()
	if err != nil {
		return
	}
	wait.Add(1)
	go func() {
		for _, pack := range other {
			if data := pack.Get(); pack.What == nil {
				Unmarshal(pack, data)
			}
		}
		wait.Done()
	}()

	data := orgPack.Get()
	if orgPack.What == nil {
		if v, err := Unmarshal(orgPack, data); err == nil {
			orgs := v.(*OrganismRoot)
			if packet.SignTypeDir(filepath.Join("图标", "生物")) == nil {
				wait.Add(1)
				go func() {
					OrganismIcon(orgs, filepath.Join("图标", "生物"), keep)
					wait.Done()
				}()
			}
			wait.Add(1)
			go func() {
				OrganismSwf(orgs, "生物立绘", keep)
				wait.Done()
			}()
		}
	}
	data = toolPack.Get()
	if toolPack.What == nil {
		if v, err := Unmarshal(toolPack, data); err == nil {
			tools := v.(*ToolRoot)
			if packet.SignTypeDir(filepath.Join("图标", "道具")); err == nil {
				wait.Add(1)
				go func() {
					ToolIcon(tools, filepath.Join("图标", "道具"), keep)
					wait.Done()
				}()
			}
		}
	}
}
