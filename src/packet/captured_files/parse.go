package capturedfiles

import (
	"PvzCapture/src/packet"
	"encoding/json"
	"encoding/xml"
	"errors"
	"path/filepath"
)

// 适用于organism.xml的解析结构
type OrganismRoot struct {
	All []*organism `json:"organisms" xml:"organisms>item"`
}
type organism struct {
	Id                  string      `json:"id" xml:"id,attr"`
	Name                string      `json:"name" xml:"name,attr"`
	Type                string      `json:"type" xml:"type,attr"`
	Attack_amount       string      `json:"attack_amount" xml:"attack_amount,attr"`
	Attack_type         string      `json:"attack_type" xml:"attack_type,attr"`
	Attribute           string      `json:"attribute" xml:"attribute,attr"`
	Use_condition       string      `json:"use_condition" xml:"use_condition,attr"`
	Use_result          string      `json:"use_result" xml:"use_result,attr"`
	Expl                string      `json:"expl" xml:"expl,attr"`
	Height              string      `json:"height" xml:"height,attr"`
	Width               string      `json:"width" xml:"width,attr"`
	Output              string      `json:"output" xml:"output,attr"`
	Sell_price          string      `json:"sell_price" xml:"sell_price,attr"`
	Img_id              string      `json:"img_id" xml:"img_id,attr"`
	Photosynthesis_time string      `json:"photosynthesis_time" xml:"photosynthesis_time,attr"`
	Evolutions          []evolution `json:"evolutions" xml:"evolutions>item,omitempty"`
}
type evolution struct {
	Id      string `json:"id" xml:"id,attr"`
	Grade   string `json:"grade" xml:"grade,attr"`
	Target  string `json:"target" xml:"target,attr"`
	Tool_id string `json:"tool_id" xml:"tool_id,attr"`
	Money   string `json:"money" xml:"money,attr"`
}

// 适用于tool.xml文件的解析结构
type ToolRoot struct {
	All []tool `json:"tools" xml:"tools>item"`
}
type tool struct {
	Id            string `json:"id" xml:"id,attr"`
	Name          string `json:"name" xml:"name,attr"`
	Type          string `json:"type" xml:"type,attr"`
	Use_level     string `json:"use_level" xml:"use_level,attr"`
	Img_id        string `json:"img_id" xml:"img_id,attr"`
	Type_name     string `json:"type_name" xml:"type_name,attr"`
	Sell_price    string `json:"sell_price" xml:"sell_price,attr"`
	Use_condition string `json:"use_condition" xml:"use_condition,attr"`
	Use_result    string `json:"use_result" xml:"use_result,attr"`
	Describe      string `json:"describe" xml:"describe,attr"`
	Rare          string `json:"rare" xml:"rare,attr"`
	Lottery_name  string `json:"lottery_name" xml:"lottery_name,attr"`
}

// 适用于talent.xml文件的解析结构
type TalentRoot struct {
	Talents []*talent `json:"talents" xml:"talents>talent"`
	Souls   []*soul   `json:"souls" xml:"souls>soul"`
}
type talent struct {
	Id    string        `json:"id" xml:"id,attr"`
	Level *talent_level `json:"level" xml:"level"`
}

type talent_level struct {
	Grade      string              `json:"grade" xml:"grade,attr"`
	Org_Grade  string              `json:"org_grade" xml:"org_grade,attr"`
	User_Grade string              `json:"user_grade" xml:"user_grade,attr"`
	Possi      string              `json:"possi" xml:"possi,attr"`
	Num        string              `json:"num" xml:"num,attr"`
	Describe   string              `json:"describe" xml:"describe,attr"`
	UseTools   []talent_level_tool `json:"use_tools" xml:"tools>tool"`
}
type talent_level_tool struct {
	Id  string `json:"id" xml:"id,attr"`
	Num string `json:"num" xml:"num,attr"`
}

type soul struct {
	Id      string          `json:"id" xml:"id,attr"`
	Tals    []*soul_tal     `json:"tals" xml:"tals>tal"`
	AddProp []*soul_addProp `json:"add_prop" xml:"add_prop>prop"`
}
type soul_tal struct {
	Id    string `json:"id" xml:"id,attr"`
	Level string `json:"level" xml:"level,attr"`
}

type soul_addProp struct {
	Name  string `json:"name" xml:"name,attr"`
	Vaule string `json:"vaule" xml:"vaule,attr"`
}

// 适用于awards.xml文件的解析结构
type AwardsRoot struct {
	Hunting []*hunting `json:"hunting" xml:"hunting>h"`
}

type hunting struct {
	Id  string `json:"id" xml:"id,attr"`
	Ads []*ad  `json:"ads" xml:"a>ad"`
}
type ad struct {
	ID   string `json:"id" xml:"v"`
	Type string `json:"type" xml:"t"`
}

// 适用于quality.xml文件的解析结构
type QualityRoot struct {
	Qualitys []*quality `json:"qualitys" xml:"qualitys>item"`
}
type quality struct {
	Id        string `json:"id" xml:"id,attr"`
	Name      string `json:"name" xml:"name,attr"`
	Skill_num string `json:"skill_num" xml:"skill_num,attr"`
	Up_ratio  string `json:"up_ratio" xml:"up_ratio,attr"`
}

// 适用于gemexchange.xml文件的解析结构
type GemexchangeRoot struct {
	Gemexchanges []*gemexchange `json:"gemexchanges" xml:"gemexchanges>item"`
}
type gemexchange struct {
	Target_Tool_Id string `json:"target_tool_id" xml:"target_tool_id,attr"`
	Cost_Tool      string `json:"cost_tool" xml:"cost_tool,attr"`
	Cost_Num       string `json:"cost_num" xml:"cost_num,attr"`
	Cost_Money     string `json:"cost_money" xml:"cost_money,attr"`
}

// 适用于所有youkia/config/lang/*.xml文件的解析结构
type LangRoot struct {
	Langs []*lang `json:"lang" xml:"item"`
}
type lang struct {
	Name string `json:"name" xml:"name,attr"`
	Str  string `json:"str" xml:",innerxml"`
}

// 适用于除ui_config.xml以外的所有ui索引文件的解析结构
type UIIndexRoot struct {
	UI []*ui `json:"ui" xml:"ui>item"`
}

type ui struct {
	Name    string `json:"name" xml:",innerxml"`
	Type    string `json:"type" xml:"type,attr"`
	SwfName string `json:"swf_name" xml:"name,attr"`
	Isnew   string `json:"isnew" xml:"isnew,attr"`
	Version string `json:"version" xml:"version,attr"`
}

// 适用于config.xml文件的解析结构，他保存了版本信息
type ConfighRoot struct {
	Base            base   `json:"base" xml:"base"`
	Version         string `json:"version" xml:"version"`
	Org_version     string `json:"org_version" xml:"org_version"`
	Pvz_version     string `json:"pvz_version" xml:"pvz_version"`
	Loading_version string `json:"loading_version" xml:"loading_version"`
	UIIndexRoot
}
type base struct {
	IsWeb    string `json:"isWeb" xml:"isWeb"`
	Language string `json:"language" xml:"language"`
	Rank     string `json:"rank" xml:"rank"`
}

// 适用于ui_config.xml文件的解析结构
type UIConfigRoot struct {
	UI []*uiconfig_ui `json:"ui" xml:"ui"`
}
type uiconfig_ui struct {
	Mark  string              `json:"mark" xml:"mark,attr"`
	Name  string              `json:"name" xml:"name,attr"`
	Items []*uiconfig_ui_item `json:"items" xml:"item"`
}
type uiconfig_ui_item struct {
	Swfname string `json:"swfname" xml:"swfname,attr"`
	Folder  string `json:"folder" xml:"folder,attr"`
	Version string `json:"version" xml:"version,attr"`
}

// 将请求包的内容反序列化后返回，错误会导致函数提前退出并返回，并且会将其发送到日志文件中
// 该函数会将得到的对象序列化成json格式文件并更改后缀后保存在同目录下，中途如果发生错误会发送到日志文件中而不会返回
func Unmarshal(pack *packet.Packet, data []byte) (v interface{}, err error) {
	switch filepath.Base(pack.Url) {
	case "tool.xml":
		v = &ToolRoot{}
	case "talent.xml":
		v = &TalentRoot{}
	case "awards.xml":
		v = &AwardsRoot{}
	case "organism.xml":
		v = &OrganismRoot{}
	case "quality.xml":
		v = &QualityRoot{}
	case "gemexchange.xml":
		v = &GemexchangeRoot{}
	case "ui_config.xml":
		v = &UIConfigRoot{}
	case "config.xml":
		v = &ConfighRoot{}
	default:
		if ok, _ := filepath.Match("youkia/config/load/*", pack.Url); ok || pack.Url == "youkia/config/config.xml" {
			v = &UIIndexRoot{}
		} else if ok, _ := filepath.Match("youkia/config/lang/*", pack.Url); ok {
			v = &LangRoot{}
		} else {
			err = errors.New("未知类型的.xml文件")
			return
		}
	}
	err = xml.Unmarshal(data, v)
	if err != nil {
		err = errors.New("文件：\"" + pack.Name + "\" 反序列化失败：" + err.Error())
		packet.HaveError(err)
		return
	}
	jf, _ := json.MarshalIndent(v, "", "  ")
	packet.KeepFile(pack.Name[:len(pack.Name)-len(filepath.Ext(pack.Name))]+".json", jf)
	return
}
