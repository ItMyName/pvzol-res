package capturedfiles

import "PvzCapture/src/packet"

// 获取organism.xml、tool.xml以及其它的xml格式信息文件请求包(不含UI索引文件)，如果类型文件夹未能成功创建则会阻止返回数据
func XmlInfoFilePackets() (organism, tool *packet.Packet, other []*packet.Packet, err error) {
	if err = packet.SignTypeDir("信息/语言文件"); err != nil {
		return
	} else if err = packet.SignTypeDir("信息/其它"); err != nil {
		return
	}
	organism = &packet.Packet{Name: "信息/organism.xml", Url: "pvz/php_xml/organism.xml"}
	tool = &packet.Packet{Name: "信息/tool.xml", Url: "pvz/php_xml/tool.xml"}
	other = []*packet.Packet{
		{Name: "信息/talent.xml", Url: "pvz/php_xml/talent.xml"},
		{Name: "信息/awards.xml", Url: "pvz/php_xml/awards.xml"},
		{Name: "信息/quality.xml", Url: "youkia/config/quality.xml"},
		{Name: "信息/gemexchange.xml", Url: "pvz/php_xml/gemexchange.xml"},

		{Name: "信息/语言文件/language_cn.xml", Url: "youkia/config/lang/language_cn.xml"},
		{Name: "信息/语言文件/possession_cn.xml", Url: "youkia/config/lang/possession_cn.xml"},
		{Name: "信息/语言文件/world_cn.xml", Url: "youkia/config/lang/world_cn.xml"},
		{Name: "信息/语言文件/severBattle_cn.xml", Url: "youkia/config/lang/severBattle_cn.xml"},
		{Name: "信息/语言文件/genius_cn.xml", Url: "youkia/config/lang/genius_cn.xml"},
		{Name: "信息/语言文件/copy_cn.xml", Url: "youkia/config/lang/copy_cn.xml"},
		{Name: "信息/语言文件/shakeTree_cn.xml", Url: "youkia/config/lang/shakeTree_cn.xml"},

		{Name: "信息/其它/MenuBtnConfig.xml", Url: "youkia/config/MenuBtnConfig.xml"},
		{Name: "信息/其它/helpGardenInfo.xml", Url: "youkia/config/helpGardenInfo.xml"},
		{Name: "信息/其它/helpBattleInfo.xml", Url: "youkia/config/helpBattleInfo.xml"},
		{Name: "信息/其它/helpBattleInfoSec.xml", Url: "youkia/config/helpBattleInfoSec.xml"},
	}
	return
}

// 获取ui_config.xml以及其它的UI索引文件，如果类型文件夹未能成功创建则会阻止返回数据
func XmlUIIndexFilePackets() (config, ui_config *packet.Packet, other []*packet.Packet, err error) {
	if err = packet.SignTypeDir("信息/UI索引/子界面"); err != nil {
		return
	} else if err = packet.SignTypeDir("信息/UI索引/副本/世界副本"); err != nil {
		return
	} else if err = packet.SignTypeDir("信息/UI索引/副本/剧情副本"); err != nil {
		return
	}

	config = &packet.Packet{
		Name: "信息/UI索引/config.xml", Url: "youkia/config/config.xml",
	}
	ui_config = &packet.Packet{
		Name: "信息/UI索引/ui_config.xml", Url: "youkia/config/ui_config.xml",
	}
	other = []*packet.Packet{
		{Name: "信息/UI索引/world.xml", Url: "youkia/config/load/world/world.xml"},

		{Name: "信息/UI索引/子界面/tree.xml", Url: "youkia/config/load/tree.xml"},
		{Name: "信息/UI索引/子界面/arena.xml", Url: "youkia/config/load/arena.xml"},
		{Name: "信息/UI索引/子界面/garden.xml", Url: "youkia/config/load/garden.xml"},
		{Name: "信息/UI索引/子界面/genius.xml", Url: "youkia/config/load/genius.xml"},
		{Name: "信息/UI索引/子界面/hunting.xml", Url: "youkia/config/load/hunting.xml"},
		{Name: "信息/UI索引/子界面/possession.xml", Url: "youkia/config/load/possession.xml"},
		{Name: "信息/UI索引/子界面/shakeTree.xml", Url: "youkia/config/load/shakeTree.xml"},

		{Name: "信息/UI索引/副本/limit.xml", Url: "youkia/config/load/copy/limit.xml"},
		{Name: "信息/UI索引/副本/stone.xml", Url: "youkia/config/load/copy/stone.xml"},

		{Name: "信息/UI索引/副本/世界副本/insideWorld_1.xml", Url: "youkia/config/load/world/insideWorld_1.xml"},
		{Name: "信息/UI索引/副本/世界副本/insideWorld_2.xml", Url: "youkia/config/load/world/insideWorld_2.xml"},
		{Name: "信息/UI索引/副本/世界副本/insideWorld_3.xml", Url: "youkia/config/load/world/insideWorld_3.xml"},
		{Name: "信息/UI索引/副本/世界副本/insideWorld_4.xml", Url: "youkia/config/load/world/insideWorld_4.xml"},
		{Name: "信息/UI索引/副本/世界副本/insideWorld_5.xml", Url: "youkia/config/load/world/insideWorld_5.xml"},

		{Name: "信息/UI索引/副本/剧情副本/panel_1.xml", Url: "youkia/config/load/copy/panels/panel_1.xml"},
		{Name: "信息/UI索引/副本/剧情副本/panel_2.xml", Url: "youkia/config/load/copy/panels/panel_2.xml"},
		{Name: "信息/UI索引/副本/剧情副本/panel_3.xml", Url: "youkia/config/load/copy/panels/panel_3.xml"},
		{Name: "信息/UI索引/副本/剧情副本/panel_4.xml", Url: "youkia/config/load/copy/panels/panel_4.xml"},
		{Name: "信息/UI索引/副本/剧情副本/panel_5.xml", Url: "youkia/config/load/copy/panels/panel_5.xml"},
		{Name: "信息/UI索引/副本/剧情副本/panel_6.xml", Url: "youkia/config/load/copy/panels/panel_6.xml"},
		{Name: "信息/UI索引/副本/剧情副本/panel_7.xml", Url: "youkia/config/load/copy/panels/panel_7.xml"},
		{Name: "信息/UI索引/副本/剧情副本/panel_8.xml", Url: "youkia/config/load/copy/panels/panel_8.xml"},
		{Name: "信息/UI索引/副本/剧情副本/panel_9.xml", Url: "youkia/config/load/copy/panels/panel_9.xml"},
	}
	return
}
