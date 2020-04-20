// Code is generated by ucloud-model, DO NOT EDIT IT.

package umedia

/*
TaskIdList - 任务ID列表
*/
type TaskIdList struct {

	// 任务ID
	TaskId string
}

/*
CodecPattenListNode - 转码模板信息
*/
type CodecPattenListNode struct {

	// 转码任务结束后，回调客户的url地址。
	CallbackUrl string

	// 音频码率，单位kbps。
	DestAudioBitrate int

	// 音频声道数量。
	DestAudioChannel int

	// 音频采样率，单位hz。
	DestAudioSample int

	// 目标视频格式，可选值为mp4、flv、mpegts。
	DestFormat string

	// 目标视频的文件名后缀。
	DestSuffix string

	// 视频码率，单位kbps。
	DestVideoBitrate int

	// 视频的编码类型
	DestVideoCodec string

	// 视频分辨率，格式为像素宽度x像素高度，例如1280x720。该字段为空表示保持原始视频大小。
	DestVideoResolution string

	// 模版的ID
	PattenId string

	// 模版名称,长度不查过256个字节
	PattenName string
}

/*
WatermarkDetail - 水印模板详情
*/
type WatermarkDetail struct {

	// 图片水印的url地址，只有图片水印有该字段。
	ImageUrl string

	// 水印离最近的水平边线占整个视频宽度的百分比，取值[0-49]
	PaddingX int

	// 水印离最近的垂直边线占整个视频高度的百分比,取值[0-49]
	PaddingY int

	// 水印的参考位置，topleft、topright、center、bottomleft、bottomright分别表示左上、右上、居中、左下、右下
	Position string

	// 文字水印内容
	TextContent string

	// image表示图片水印，text表示文字水印
	WaterMarkType string

	// 任务ID
	WatermarkId string

	// 水印名称
	WatermarkName string
}

/*
CodecPattenDetail - 解码模板详细信息
*/
type CodecPattenDetail struct {

	// 转码模版ID
	CodecPattenId string

	// 转码模版名称
	CodecPattenName string

	// 输出视频格式
	DestFormat string

	// 输出文件后缀
	DestSuffix string

	// 输出文件视频码率
	DestVideoBitrate int

	// 视频编码类型
	DestVideoCodec string

	// 输出文件分辨率
	DestVideoResolution string
}

/*
CodecPattenInfo - 解码模板信息
*/
type CodecPattenInfo struct {

	// 转码清晰度：superdefinition、highdefinition、mediumdefinition、lowdefinition，分别对应计费的4种规格。只有当任务状态为处理完成时，该参数有意义，其他状态该参数为空字符串。
	CodecLevel string

	// 任务创建时间，单位：Unix时间戳
	CreateTime int

	// 输出文件的存储空间
	DestBucket string

	// 输出视频封装格式
	DestFormat string

	// 输出文件名
	DestVideoName string

	// 视频时长，单位：秒。只有当任务状态为处理完成时，该参数有意义，其他状态该参数为0。
	Duration int

	// 原始视频url地址
	SrcUrl string

	// 任务状态：waiting、processing、finished、failed，分别表示排队中，处理中，处理完成，处理失败。
	Status string

	// 任务ID
	TaskId string
}

/*
CodecStatInfo - 转码状态信息
*/
type CodecStatInfo struct {

	// 超清转码时长，单位：秒
	HighDefinitionTime int

	// 普清转码时长，单位：秒
	LowDefinitionTime int

	// 高清转码时长，单位：秒
	MidiumDefinitionTime int

	// 超高清转码时长，单位：秒
	SuperDefinitionTime int

	// 统计时间，单位：Unix时间戳
	Time int

	// 所有清晰度转码时长总和，单位：秒
	TotalTime int

	// 视频的编码类型
	VideoCodec string
}

/*
FormatTaskInfo - 切片任务信息
*/
type FormatTaskInfo struct {

	// 任务创建时间，单位：Unix时间戳
	CreateTime int

	// 输出文件的存储空间
	DestBucket string

	// 输出视频封装格式
	DestFormat string

	// 输出文件名
	DestVideoName string

	// 视频时长，单位：秒。只有当任务状态为处理完成时，该参数有意义，其他状态该参数为0。
	Duration int

	// 原始视频url地址
	SrcUrl string

	// 任务状态：waiting、processing、finished、failed，分别表示排队中，处理中，处理完成，处理失败。
	Status string

	// 任务ID
	TaskId string
}

/*
FormatStatInfo - 切片状态信息
*/
type FormatStatInfo struct {

	// 转封装时长，单位：秒
	FormatTime int

	// 统计时间，单位：Unix时间戳
	Time int
}

/*
SnapTaskInfo - 截图任务信息
*/
type SnapTaskInfo struct {

	// 任务创建时间，单位：Unix时间戳
	CreateTime int

	// 截图张数
	ImageCount int

	// 图片格式类型
	ImageFormat string

	// 截图输出的文件名。对于周期截图，该字段表示不加后缀的文件名。
	ImageName string

	// 截图模式，single表示确定时间点单张截图，periodic表示周期截图,dynamic表示为gif截图
	SnapType string

	// 原始视频url地址
	SrcUrl string

	// 任务状态：waiting、processing、finished、failed，分别表示排队中，处理中，处理完成，处理失败。
	Status string

	// 任务ID
	TaskId string
}

/*
SnapStatInfo - 截图状态信息
*/
type SnapStatInfo struct {

	// 截图张数，单位：张
	ImageCount int

	// 统计时间，单位：Unix时间戳
	Time int
}

/*
WaterMarkPattenListNode - 水印模板列表信息
*/
type WaterMarkPattenListNode struct {

	// 转码任务结束后，回调客户的url地址。
	CallbackUrl string

	// 字体颜色，只有文字水印有该字段。
	FontColor string

	// 字体大小，单位：磅，只有文字水印有该字段。
	FontSize int

	// 字体类型，只有文字水印有该字段。
	FontType string

	// 图片水印的url地址，只有图片水印有该字段。
	ImageUrl string

	// 水印离最近的水平边线占整个视频宽度的百分比，取值[0-49]
	PaddingX int

	// 水印离最近的垂直边线占整个视频高度的百分比,取值[0-49]
	PaddingY int

	// 模版的ID
	PattenId string

	// 模版名称,长度不超过256个字节
	PattenName string

	// 水印的参考位置，topleft、topright、center、bottomleft、bottomright分别表示左上、右上、居中、左下、右下
	Position string

	// 文字水印的内容，只有文字水印有该字段。
	TextContent string

	// image表示图片水印，text表示文字水印
	WaterMarkType string
}
