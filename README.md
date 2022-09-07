# 植物大战僵尸Online 数据包解析
由于需要对上下文的部分信息进行引用，所以设定语法：
  - 使用`$markname`对某段信息进行标记，标记内容请自行联系上下文理解，标记名称在全文章中只能用于对同一信息含义进行解释
  - 使用`#markname#`在当前位置填充上对应内容

---

- 首页 <http://pvz.youkia.com/>
- 选区 <http://pvz.youkia.com/xf/index.php>
- 服务器区号（1-46），`$sid`
- 游戏服务器域名，`$host`
  - 区号大于等于12 <http://s#sid#.youkia.pvz.youkia.com>
  - 区号小于等于11 <http://pvz-s#sid#.youkia.com>

浏览器客户端使用swf+xml文件构建，http协议，无加密，无混淆……



## 账号cookie获取
通过明文的账号(`$user`)密码(`$password`)，获取youkia账号cookie，最后提取pvz_youkia这个cookie

1. 获取youkia账号cookie
    ```
    GET http://www.youkia.com/index.php/logins/login?   account=#user#&password=#password#
    ```  
    在回复标头中查找，可获得`Set-Cookie: youkia=$youkia`

2. 获取游戏播放文件(.html)  
    ```
    GET http://www.youkia.com/index.php/pvz/s#sid#
    Cookie: youkia=#youkia#
    ```
    回复正文是html文件，在正文中匹配到`playAddress`这个标记
    ```
    <iframe name="iframe_canvas" id="iframe_canvas" src="$playAddress" frameborder="0" style="margin: 0px;padding: 0px;width: 100%; height:750px;"></iframe> 
    ```

3. 获取游戏账号cookie
    ```
    GET #playAddress#
    Cookie: youkia=#youkia#
    ```
    在回复标头中查找，可获得
    ```
    Set-Cookie: pvz_youkia=$pvz_youkia
    ```

## XML信息文件
`GET #host#/youkia/main.swf `
该游戏的主程序(.swf)，结合抓包数据提取到部分无上层索引的url，以下路径均使用主机域名(`#host#`)：
```
----基础UI索引----
youkia/config/config.xml             // 基础素材（含主界面）以及版本号信息，标记版本号$pvz_version与$loading_version
youkia/config/ui_config.xml          // 部分轻量子界面(与其它ui文件模板不同)
youkia/config/load/world/world.xml   // 副本、奖励、排行榜三个子界面

----加载UI界面----
youkia/UILibs/loading#loading_version#.swf  // 加载读条界面
youkia/UILibs/pvz#pvz_version#.swf          // 主界面界面基础图标

----子界面UI索引----
youkia/config/load/tree.xml          // 世界树
youkia/config/load/arena.xml         // 斗技场
youkia/config/load/garden.xml        // 花园
youkia/config/load/genius.xml        // 宝石系统
youkia/config/load/hunting.xml       // 狩猎场
youkia/config/load/possession.xml    // 领地
youkia/config/load/shakeTree.xml     // 宝石狩猎场

----副本UI索引----
youkia/config/load/copy/limit.xml    // 限时副本
youkia/config/load/copy/stone.xml    // 宝石副本
youkia/config/load/world/insideWorld_ + param1 + .xml // 世界副本（参数：1、2、3、4、5）五张地图
youkia/config/load/copy/panels/panel_ + this.m_chapterId + .xml // 似乎是剧情副本（参数：1、2、3、4、5、6、7、8、9）九张地图

----信息----
pvz/php_xml/tool.xml                 // 道具
pvz/php_xml/talent.xml               // 天赋
pvz/php_xml/awards.xml               // 狩猎场关卡奖励
pvz/php_xml/organism.xml             // 生物（包括僵尸）
youkia/config/quality.xml            // 品质
pvz/php_xml/gemexchange.xml          // 宝石兑换
youkia/config/lang/" + param2 + param1 + ".xml // 语言文件（参数：param2="language_"、"possession_"、"world_"、"severBattle_"、"genius_"、"copy_"、shakeTree_"，param1="cn"）

----其它----
youkia/config/MenuBtnConfig.xml      // 主界面按钮的ID
youkia/config/helpGardenInfo.xml     // 新手帮助花园部分
youkia/config/helpBattleInfo.xml     // 新手帮助战斗选择部分
youkia/config/helpBattleInfoSec.xml  // 新手帮助战斗动画部分
```

### 信息文件

#### organism.xml
该xml格式文件，记录有所有生物（植物、僵尸）的相关信息，包含以下信息：
```
root
  organisms
    item
    |id                   生物ID
    |name                 生物名称
    |type                 生物类型
    |attack_amount        攻击次数或其子弹个数
    |attack_type          攻击类型
    |attribute            属性
    |use_condition        说明
    |use_result           描述
    |expl                 生物级别信息（星级、成长值范围）
    |height               占用高度
    |width                占用宽度
    |output               光合作用产出金币
    |sell_price           售价
    |img_id               图像ID，$o_img
    |photosynthesis_time  光合作用时间
      evolutions
        item
        |id       这条进化路线的ID
        |grade    进化所需等级
        |target   进化后的植物ID
        |tool_id  进化所需素材
        |money    进化所需金币
```



**区分植物与僵尸的方法：判断photosynthesis_time是否为0**

- 获取生物图标文件.swf（只有植物有图标）
  ```
  GET #%host%#/youkia/IconRes/IconOrg/#o_img#.swf
  ```

- 获取生物的立绘文件.swf
  ```
  GET #%host%#/youkia/ORGLibs/org_#o_img#.swf
  ```

#### tool.xml
该xml格式文件，记录有所有道具的相关信息，包含以下信息：
```
root
  tools
    item
    |id             道具ID
    |name           道具名称
    |type           未知
    |use_level      使用所需等级
    |img_id         图像ID，`$t_img`
    |type_name      类型名称
    |sell_price     售价
    |use_condition  说明
    |use_result     使用效果
    |describe       使用方法
    |rare           品质
    |lottery_name   未知，似乎是对具有抽奖性质的道具标识
```

- 获取道具图标文件.swf
  ```
  GET #host#/youkia/IconRes/IconTool/#t_img#.swf
  ```

#### talent.xml 
该xml格式文件，记录有天赋以及灵魂各等级的相关信息，包含以下信息：
```
root
  talents
    talent
    |id  天赋ID
      level
      |grade       天赋等级
      |org_grade   升级天赋所需的植物等级
      |user_grade  升级天赋所需的账号等级
      |possi       释放该天赋的几率（百分比）
      |num         未知
      |describe    说明
        tools
          tool
          |id   升级天赋的道具（宝石）ID
          |num  所需个数
  souls
    soul
    |id  灵魂ID
      tals
        tal
        |id     该灵魂等级需要达标的天赋ID
        |level  该灵魂等级需要达标的天赋等级
      add_prop
        prop
        |name   提供加成名称
        |vaule  提供的加成数值（百分比）
```

#### awards.xml
该xml格式文件，记录有狩猎场各个关卡奖励的相关信息，包含以下信息：
```
root
  hunting
    h
    |id  狩猎场关卡ID
      a
        ad  奖励物品
          t  奖励类型(tool、organism)
          v  奖励物品的ID
```

#### quality.xml
该xml格式文件，记录有植物每种品质的相关信息，包含以下信息：
```
root
  qualitys
    item
    |id         品质ID
    |name       品质名称
    |skill_num  未知
    |up_ratio   升级几率（百分比）
```

#### gemexchange.xml 
该xml格式文件，记录有宝石兑换的相关信息，包含以下信息：
```
root
  gemexchanges
    item
    |target_tool_id  合成的目标道具ID（宝石道具ID）
    |cost_tool       花费的道具的ID（宝石道具ID）
    |cost_num        花费的道具数目
    |cost_money      花费的金币
```

#### lang/*.xml
此类该xml格式文件，记录着各种词条与其标记名称的映射关系，包含以下信息：
```
lang
  item  对应词条
  |name  词条标记名称

```

### UI资源索引文件
基础UI索引、子界面UI索引、副本UI索引，此类xml格式文件，记录有各个界面的UI资源路径（只是路径）。

除了ui_config.xml，其它的文件包含以下信息：
```
config
  ui
    item  UI名称
    |type     UI资源文件的类型，$uitype
    |name     UI资源文件.swf的名称，$uiname
    |isnew    是否是新版（0、1）
    |version  UI版本号，$uiversion
```

config.xml，在兼容上述结构的情况下附带有游戏版本信息：
```
config
  base
    isWeb
    language
    rank
  version         主版本号
  org_version     生物版本号？
  pvz_version     游戏版本号，加载pvz.swf文件时指定版本
  loading_version 加载loading.swf文件时指定版本
  ui
    item  UI名称
    |type     UI资源文件的类型，$uitype
    |name     UI资源文件.swf的名称，$uiname
    |isnew    是否是新版（0、1）
    |version  UI版本号，$uiversion
```

ui_config.xml，包含以下信息：
```
data
  ui
  |mark  标记名称
  |name  UI名称
    item
    |swfname  UI资源文件.swf的名称，$uiname
    |folder   UI资源文件的类型，$uitype
    |version  I版本号，$uiversion
```

- 除了ui_config.xml的UI资源文件.swf，都可以使用该请求获取：
  ```
  GET #host#/youkia/#uitype#/#uiname##uiversion#.swf
  ```
- ui_config.xml的UI资源.swf类型部分自带分隔符
  ```
  GET #host#/youkia/#uitype##uiname##uiversion#.swf
  ```
- 此外还有两项无索引的ui.swf文件
  ```
  GET #host#/youkia/UILibs/loading#loading_version#.swf
  ```
  ```
  GET #host#/youkia/UILibs/pvz#pvz_version#.swf  
  ```
## 缺失信息
在这里列举除开xml明文提供的信息文件，但是目前依然缺少的相关信息数据原文。

- 世界树奖励
- 签到奖励
- 每日金币奖励
- 商店物品及其售价
- 充值奖励
- 狩猎场敌人信息（已知狩猎场奖励，僵尸类型信息）
- 宝箱奖励
- 世界副本关卡敌人信息
- 世界副本关卡敌人掉落信息
- 宝石副本关卡敌人信息
- 宝石副本关卡敌人掉落信息
- 植物等级提升以及属性奖励信息
- 生物属性与实际数值转换系数
- 技能等合成道具的使用效果（已知技能的抽奖属性ID？）
- 合成参数信息
- 属性克制信息
- 道具兑换信息
………………



## AMF协议的文件
AMF协议有部分是信息文件，也有部分是操作请求，但是不管是信息还是操作它们都是根据当前账号信息进行返回.换句话说，所有AMF协议的POST请求，必须附带账号cookie：`%pvz_youkia%`.

通用的请求标头：
```
POST #host#/pvz/amf/
Content-Length: xx
Content-Type: application/x-amf
Cookie: pvz_youkia=#pvz_youkia#
```

然后这个AMF请求方法名称……一共找到了132条，反正我列出来了，在下面……
虽然很多，但是像包含："get"、"All"、"info"这种关键词的都属于比较有用的信息文件，像上文有提到的目前确实的信息数据，很有可能从该文件获取。

另外一提这个协议是一个十年前的东西了，官方的规范文件都看不到了。


```
"api.apiorganism.activities"
"api.arena.awardWeek"
"api.arena.getArenaList"
"api.arena.getRankListOld"
"api.arena.getAwardWeekInfo"
"api.arena.getRankList"
"api.arena.setOrganism"
"api.arena.challenge"
"api.reward.lottery"
"api.garden.outAndStealAll"
"api.garden.add"
"api.garden.removeStateAll"
"api.cave.challenge"
"api.cave.openCave"
"api.cave.openGrid"
"api.cave.useTimesands"
"api.tool.synthesis"
"api.apiorganism.refreshHp"
"api.apiorganism.matureRecompute"
"api.apiorganism.qualityUp"
"api.apiorganism.quality12Up"
"api.apiorganism.skillLearn"
"api.apiorganism.skillUp"
"api.shop.buy"
"api.shop.getMerchandises"
"api.shop.init"
"api.shop.reset"
"api.apiskill.getAllSkills"
"api.garden.organismReturnHome"
"api.shop.sell"
"api.reward.openbox"
"api.duty.getAward"
"api.duty.acceptedDuty"
"api.duty.getDuty"
"api.tool.useOf"
"api.apiuser.lock"
"api.territory.getTerritory"
"api.territory.init"
"api.territory.challenge"
"api.territory.getMsg"
"api.territory.recommen"
"api.territory.quit"
"api.territory.getAward"
"api.apiactivity.info"
"api.apiactivity.getActTool"
"api.apiactivity.exchange"
"api.guide.getGuideInfo"
"api.guide.setCusReward"
"api.guide.getCurAccSmall"
"api.guide.setAccSmall"
"api.guide.getCurAccBig"
"api.guide.setAccBig"
"api.guide.getsumtimeact"
"api.guide.getsumtimereward"
"api.guide.getDailyReward"
"api.apiorganism.strengthen"
"api.fuben.display"
"api.fuben.caveInfo"
"api.fuben.openCave"
"api.fuben.challenge"
"api.fuben.addChallengeCount"
"api.fuben.addCaveChallengeCount"
"api.fuben.top"
"api.fuben.reward"
"api.fuben.award"
"api.guide.getFirstReward"
"api.guide.setFirstReward"
"api.vip.rewards"
"api.vip.awards"
"api.vip.startAuto"
"api.vip.stopAuto"
"api.vip.autoRewards"
"api.gift.get"
"api.shop.getMerchandises"
"api.shop.buy"
"api.serverbattle.knockoutReward"
"api.serverbattle.knockoutAward"
"api.serverbattle.qualifying"
"api.serverbattle.getOpponent"
"api.serverbattle.addChallengeCount"
"api.serverbattle.challenge"
"api.serverbattle.knockout"
"api.serverbattle.fightLog"
"api.serverbattle.getGroup"
"api.serverbattle.knockout"
"api.serverbattle.guess"
"api.serverbattle.ruleDescription"
"api.serverbattle.signUp"
"api.serverbattle.qualifyingReward"
"api.serverbattle.qualifyingAward"
"api.serverbattle.allGuess"
"api.serverbattle.guessAward"
"api.serverbattle.getIntegralTop"
"api.serverbattle.guessTop"
"api.serverbattle.message"
"api.banner.get"
"api.apiorganism.getEvolutionOrgs"
"api.apiorganism.getEvolutionCost"
"api.shop.gemExchange"
"api.apiorganism.upgradeTalent"
"api.apiorganism.restTalent"
"api.zombie.getInfo"
"api.zombie.beat"
"api.zombie.addcount"
"api.message.gets"
"api.stone.getChapInfo"
"api.stone.getCaveInfo"
"api.stone.getRewardInfo"
"api.stone.reward"
"api.stone.getRankByCid"
"api.stone.challenge"
"api.stone.addCountByMoney"
"api.stone.getCaveThrougInfo"
"api.active.getCopyAllChapters"
"api.active.getAllLevels"
"api.active.challenge"
"api.active.addCount"
"api.active.getCopy"
"api.active.getState"
"api.active.getSignInfo"
"api.active.rewardTimes"
"api.active.rewardPoint"
"api.active.sign"
"api.apiskill.getSpecSkillAll"
"api.apiorganism.specSkillUp"
"api.apiorganism.getAllExchangeInfo"
"api.apiorganism.getOneExchangeInfo"
"api.apiorganism.exchangeOne"
"api.apiorganism.exchangeAll"
"api.duty.getAll"
"api.duty.reward"
"api.garden.challenge"
```