// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechatv2 for the canonical source repository
// @license     https://github.com/chanxuehong/wechatv2/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package menu

const (
	MenuButtonCountLimit    = 3 // 一级菜单最多包含 3 个按钮
	SubMenuButtonCountLimit = 5 // 二级菜单最多包含 5 个按钮
)

const (
	MenuButtonNameLenLimit    = 16 // 菜单标题不超过16个字节
	SubMenuButtonNameLenLimit = 40 // 子菜单标题不超过40个字节
)

const (
	ButtonKeyLenLimit = 128 // 菜单KEY值不能超过128字节
	ButtonURLLenLimit = 256 // 网页链接不能超过256字节
)

const (
	BUTTON_TYPE_CLICK = "click" // 点击推事件
	BUTTON_TYPE_VIEW  = "view"  // 跳转URL

	// 请注意, 下面的事件, 仅支持微信iPhone5.4.1以上版本, 和Android5.4以上版本的微信用户,
	// 旧版本微信用户点击后将没有回应, 开发者也不能正常接收到事件推送.
	BUTTON_TYPE_SCANCODE_PUSH      = "scancode_push"      // 扫码推事件
	BUTTON_TYPE_SCANCODE_WAITMSG   = "scancode_waitmsg"   // 扫码带提示
	BUTTON_TYPE_PIC_SYSPHOTO       = "pic_sysphoto"       // 系统拍照发图
	BUTTON_TYPE_PIC_PHOTO_OR_ALBUM = "pic_photo_or_album" // 拍照或者相册发图
	BUTTON_TYPE_PIC_WEIXIN         = "pic_weixin"         // 微信相册发图
	BUTTON_TYPE_LOCATION_SELECT    = "location_select"    // 发送位置
)

type Menu struct {
	Buttons []Button `json:"button,omitempty"` // 一级菜单数组，个数应为1~3个
}

// 菜单的按钮
type Button struct {
	Type       string   `json:"type,omitempty"`       // 非必须; 菜单的响应动作类型
	Name       string   `json:"name"`                 // 必须;  菜单标题，不超过16个字节，子菜单不超过40个字节
	Key        string   `json:"key,omitempty"`        // 非必须; 菜单KEY值，用于消息接口推送，不超过128字节
	URL        string   `json:"url,omitempty"`        // 非必须; 网页链接，用户点击菜单可打开链接，不超过256字节
	SubButtons []Button `json:"sub_button,omitempty"` // 非必须; 二级菜单数组，个数应为1~5个
}

// 初始化 btn 指向的 Button 为 子菜单 类型按钮
func (btn *Button) InitToSubMenuButton(name string, subButtons []Button) {
	btn.Name = name
	btn.SubButtons = subButtons

	// 容错性考虑, 清除其他字段
	btn.Type = ""
	btn.Key = ""
	btn.URL = ""
}

// 初始化 btn 指向的 Button 为 click 类型按钮
func (btn *Button) InitToClickButton(name, key string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_CLICK
	btn.Key = key

	// 容错性考虑, 清除其他字段
	btn.URL = ""
	btn.SubButtons = nil
}

// 初始化 btn 指向的 Button 为 view 类型按钮
func (btn *Button) InitToViewButton(name, url string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_VIEW
	btn.URL = url

	// 容错性考虑, 清除其他字段
	btn.Key = ""
	btn.SubButtons = nil
}

// 初始化 btn 指向的 Button 为 扫码推事件 类型按钮
func (btn *Button) InitToScanCodePushButton(name, key string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_SCANCODE_PUSH
	btn.Key = key

	// 容错性考虑, 清除其他字段
	btn.URL = ""
	btn.SubButtons = nil
}

// 初始化 btn 指向的 Button 为 扫码推事件且弹出“消息接收中”提示框 类型按钮
func (btn *Button) InitToScanCodeWaitMsgButton(name, key string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_SCANCODE_WAITMSG
	btn.Key = key

	// 容错性考虑, 清除其他字段
	btn.URL = ""
	btn.SubButtons = nil
}

// 初始化 btn 指向的 Button 为 弹出系统拍照发图 类型按钮
func (btn *Button) InitToPicSysPhotoButton(name, key string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_PIC_SYSPHOTO
	btn.Key = key

	// 容错性考虑, 清除其他字段
	btn.URL = ""
	btn.SubButtons = nil
}

// 初始化 btn 指向的 Button 为 弹出拍照或者相册发图 类型按钮
func (btn *Button) InitToPicPhotoOrAlbumButton(name, key string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_PIC_PHOTO_OR_ALBUM
	btn.Key = key

	// 容错性考虑, 清除其他字段
	btn.URL = ""
	btn.SubButtons = nil
}

// 初始化 btn 指向的 Button 为 弹出微信相册发图器 类型按钮
func (btn *Button) InitToPicWeixinButton(name, key string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_PIC_WEIXIN
	btn.Key = key

	// 容错性考虑, 清除其他字段
	btn.URL = ""
	btn.SubButtons = nil
}

// 初始化 btn 指向的 Button 为 弹出地理位置选择器 类型按钮
func (btn *Button) InitToLocationSelectButton(name, key string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_LOCATION_SELECT
	btn.Key = key

	// 容错性考虑, 清除其他字段
	btn.URL = ""
	btn.SubButtons = nil
}