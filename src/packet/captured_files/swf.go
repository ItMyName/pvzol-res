package capturedfiles

import (
	"PvzCapture/src/packet"
	"path/filepath"
)

var MainSwf = packet.Packet{
	Url:  "youkia/main.swf",
	Name: "main.swf",
}

// 生成所有的生物图标请求包，它们通过out管道传递，需要提供typeName存放目录
func OrganismIcon(orgs *OrganismRoot, typeName string, out chan<- *packet.Packet) {
	exist := make(map[string]bool)
	for _, org := range orgs.All {
		if !exist[org.Img_id] && org.Photosynthesis_time != "0" {
			out <- &packet.Packet{
				Url:  "youkia/IconRes/IconOrg/" + org.Img_id + ".swf",
				Name: filepath.Join(typeName, org.Img_id+".swf"),
			}
			exist[org.Img_id] = true
		}
	}
}

// 生成所有的生物立绘请求包，它们通过out管道传递
// 由于特殊性需要提供RootTypeName根存放目录，它会根据结构自动创建类型文件夹，产生的错误会被发送到日志文件中
func OrganismSwf(orgs *OrganismRoot, RootTypeName string, out chan<- *packet.Packet) {
	if packet.SignTypeDir(filepath.Join(RootTypeName, "植物")) != nil || packet.SignTypeDir(filepath.Join(RootTypeName, "僵尸")) != nil {
		return
	}
	exist := make(map[string]bool)
	var type_ string
	for _, org := range orgs.All {
		if !exist[org.Img_id] {
			if org.Photosynthesis_time != "0" {
				type_ = "植物"
			} else {
				type_ = "僵尸"
			}
			out <- &packet.Packet{
				Name: filepath.Join(RootTypeName, type_, org.Img_id+".swf"),
				Url:  "youkia/ORGLibs/org_" + org.Img_id + ".swf",
			}
			exist[org.Img_id] = true
		}
	}
}

// 生成所有的生物立绘请求包，它们通过out管道传递，需要提供typeName存放目录
func ToolIcon(tools *ToolRoot, typeName string, out chan<- *packet.Packet) {
	exist := make(map[string]bool)
	for _, tool := range tools.All {
		if !exist[tool.Img_id] {
			out <- &packet.Packet{
				Name: filepath.Join(typeName, tool.Img_id+".swf"),
				Url:  "youkia/IconRes/IconTool/" + tool.Img_id + ".swf",
			}
			exist[tool.Img_id] = true
		}
	}
}

// 针对除ui_config.xml的ui索引文件.xml，生成所有ui项的请求包
// 由于特殊性需要提供RootTypeName根存放目录，它会根据结构自动创建类型文件夹，产生的错误会被发送到日志文件中
func UISWF(uis *UIIndexRoot, RootTypeName string, out chan<- *packet.Packet) {
	for _, ui := range uis.UI {
		if packet.SignTypeDir(filepath.Join(RootTypeName, ui.Name)) == nil {
			out <- &packet.Packet{
				Name: filepath.Join(RootTypeName, ui.Name, ui.SwfName+ui.Version+".swf"),
				Url:  "youkia/" + ui.Type + "/" + ui.SwfName + ui.Version + ".swf",
			}
		}
	}
}

// 针对config.xml关于版本号索引的两个ui资源文件，生成ui项的请求包
func BaseUI(config *ConfighRoot, TypeName string, out chan<- *packet.Packet) {
	out <- &packet.Packet{
		Name: filepath.Join(TypeName, "loading"+config.Loading_version+".swf"),
		Url:  "youkia/UILibs/loading" + config.Loading_version + ".swf",
	}
	out <- &packet.Packet{
		Name: filepath.Join(TypeName, "pvz"+config.Pvz_version+".swf"),
		Url:  "youkia/UILibs/pvz" + config.Pvz_version + ".swf",
	}
}

// 针对ui_config.xml的ui索引文件.xml，生成所有ui项的请求包
// 由于特殊性需要提供RootTypeName根存放目录，它会根据结构自动创建类型文件夹，产生的错误会被发送到日志文件中
func UIConfigSWF(uis *UIConfigRoot, RootTypeName string, out chan<- *packet.Packet) {
	for _, ui := range uis.UI {
		if packet.SignTypeDir(filepath.Join(RootTypeName, ui.Name)) == nil {
			for _, item := range ui.Items {
				out <- &packet.Packet{
					Name: filepath.Join(RootTypeName, ui.Name, item.Swfname+item.Version+".swf"),
					Url:  "youkia/" + item.Folder + item.Swfname + item.Version + ".swf",
				}
			}
		}
	}
}
