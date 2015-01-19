// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechatv2 for the canonical source repository
// @license     https://github.com/chanxuehong/wechatv2/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package media

const (
	NewsArticleCountLimit = 10 // 图文消息里文章的个数限制
)

const (
	MediaTypeImage = "image"
	MediaTypeVoice = "voice"
	MediaTypeVideo = "video"
	MediaTypeThumb = "thumb"
	MediaTypeNews  = "news"
)

type MediaInfo struct {
	MediaType string `json:"type"`       // 图片（image）、语音（voice）、视频（video）、缩略图（thumb）和 图文消息（news）
	MediaId   string `json:"media_id"`   // 媒体的唯一标识
	CreatedAt int64  `json:"created_at"` // 媒体创建的时间戳
}

// 图文消息里的 Article
type NewsArticle struct {
	ThumbMediaId     string `json:"thumb_media_id"`               // 图文消息缩略图的 media_id, 可以在上传多媒体文件接口中获得
	Title            string `json:"title"`                        // 图文消息的标题
	Author           string `json:"author,omitempty"`             // 图文消息的作者
	Digest           string `json:"digest,omitempty"`             // 图文消息的摘要
	Content          string `json:"content"`                      // 图文消息页面的内容，支持HTML标签
	ContentSourceURL string `json:"content_source_url,omitempty"` // 在图文消息页面点击“阅读原文”后的页面
	ShowCoverPic     int    `json:"show_cover_pic"`               // 是否显示封面, 1为显示, 0为不显示
}

func (article *NewsArticle) SetShowCoverPic(b bool) {
	if b {
		article.ShowCoverPic = 1
	} else {
		article.ShowCoverPic = 0
	}
}