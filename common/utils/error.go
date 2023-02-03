package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type ErrorCode uint32

type ErrorWithMessage struct {
	Code    ErrorCode
	Message string
}

const (
	ErrorOk                 ErrorCode = 0x00000 // 一切 ok
	ErrorClient             ErrorCode = 0xA0001 // 用户端错误
	ErrorRegisration        ErrorCode = 0xA0100 // 用户注册错误
	ErrorA0101              ErrorCode = 0xA0101 // 用户未同意隐私协议
	ErrorA0102              ErrorCode = 0xA0102 // 注册国家或地区受限
	ErrorInvalidUsername    ErrorCode = 0xA0110 // 用户名校验失败
	ErrorUsernameConflict   ErrorCode = 0xA0111 // 用户名已存在
	ErrorBlockedUsername    ErrorCode = 0xA0112 // 用户名包含敏感词
	ErrorUsernameCharset    ErrorCode = 0xA0113 // 用户名包含特殊字符
	ErrorPasswordValidation ErrorCode = 0xA0120 // 密码校验失败
	ErrorPasswordLength     ErrorCode = 0xA0121 // 密码长度不够
	ErrorPasswordStrength   ErrorCode = 0xA0122 // 密码强度不够
	ErrorA0130              ErrorCode = 0xA0130 // 校验码输入错误
	ErrorA0131              ErrorCode = 0xA0131 // 短信校验码输入错误
	ErrorA0132              ErrorCode = 0xA0132 // 邮件校验码输入错误
	ErrorA0133              ErrorCode = 0xA0133 // 语音校验码输入错误
	ErrorA0140              ErrorCode = 0xA0140 // 用户证件异常
	ErrorA0141              ErrorCode = 0xA0141 // 用户证件类型未选择
	ErrorA0142              ErrorCode = 0xA0142 // 大陆身份证编号校验非法
	ErrorA0143              ErrorCode = 0xA0143 // 护照编号校验非法
	ErrorA0144              ErrorCode = 0xA0144 // 军官证编号校验非法
	ErrorA0150              ErrorCode = 0xA0150 // 用户基本信息校验失败
	ErrorA0151              ErrorCode = 0xA0151 // 手机格式校验失败
	ErrorA0152              ErrorCode = 0xA0152 // 地址格式校验失败
	ErrorA0153              ErrorCode = 0xA0153 // 邮箱格式校验失败
	ErrorA0200              ErrorCode = 0xA0200 // 用户登录异常
	ErrorNoSuchUser         ErrorCode = 0xA0201 // 用户账户不存在
	ErrorA0202              ErrorCode = 0xA0202 // 用户账户被冻结
	ErrorA0203              ErrorCode = 0xA0203 // 用户账户已作废
	ErrorWrongPassword      ErrorCode = 0xA0210 // 用户密码错误
	ErrorA0211              ErrorCode = 0xA0211 // 用户输入密码错误次数超限
	ErrorA0220              ErrorCode = 0xA0220 // 用户身份校验失败
	ErrorA0221              ErrorCode = 0xA0221 // 用户指纹识别失败
	ErrorA0222              ErrorCode = 0xA0222 // 用户面容识别失败
	ErrorA0223              ErrorCode = 0xA0223 // 用户未获得第三方登录授权
	ErrorA0230              ErrorCode = 0xA0230 // 用户登录已过期
	ErrorA0240              ErrorCode = 0xA0240 // 用户验证码错误
	ErrorA0241              ErrorCode = 0xA0241 // 用户验证码尝试次数超限
	ErrorA0300              ErrorCode = 0xA0300 // 访问权限异常
	ErrorA0301              ErrorCode = 0xA0301 // 访问未授权
	ErrorA0302              ErrorCode = 0xA0302 // 正在授权中
	ErrorA0303              ErrorCode = 0xA0303 // 用户授权申请被拒绝
	ErrorA0310              ErrorCode = 0xA0310 // 因访问对象隐私设置被拦截
	ErrorA0311              ErrorCode = 0xA0311 // 授权已过期
	ErrorA0312              ErrorCode = 0xA0312 // 无权限使用 API
	ErrorA0320              ErrorCode = 0xA0320 // 用户访问被拦截
	ErrorA0321              ErrorCode = 0xA0321 // 黑名单用户
	ErrorA0322              ErrorCode = 0xA0322 // 账号被冻结
	ErrorA0323              ErrorCode = 0xA0323 // 非法 IP 地址
	ErrorA0324              ErrorCode = 0xA0324 // 网关访问受限
	ErrorA0325              ErrorCode = 0xA0325 // 地域黑名单
	ErrorA0330              ErrorCode = 0xA0330 // 服务已欠费
	ErrorA0340              ErrorCode = 0xA0340 // 用户签名异常
	ErrorA0341              ErrorCode = 0xA0341 // RSA 签名错误
	ErrorA0400              ErrorCode = 0xA0400 // 用户请求参数错误
	ErrorA0401              ErrorCode = 0xA0401 // 包含非法恶意跳转链接
	ErrorA0402              ErrorCode = 0xA0402 // 无效的用户输入
	ErrorA0410              ErrorCode = 0xA0410 // 请求必填参数为空
	ErrorA0411              ErrorCode = 0xA0411 // 用户订单号为空
	ErrorA0412              ErrorCode = 0xA0412 // 订购数量为空
	ErrorA0413              ErrorCode = 0xA0413 // 缺少时间戳参数
	ErrorA0414              ErrorCode = 0xA0414 // 非法的时间戳参数
	ErrorA0420              ErrorCode = 0xA0420 // 请求参数值超出允许的范围
	ErrorA0421              ErrorCode = 0xA0421 // 参数格式不匹配
	ErrorA0422              ErrorCode = 0xA0422 // 地址不在服务范围
	ErrorA0423              ErrorCode = 0xA0423 // 时间不在服务范围
	ErrorA0424              ErrorCode = 0xA0424 // 金额超出限制
	ErrorA0425              ErrorCode = 0xA0425 // 数量超出限制
	ErrorA0426              ErrorCode = 0xA0426 // 请求批量处理总个数超出限制
	ErrorA0427              ErrorCode = 0xA0427 // 请求 JSON 解析失败
	ErrorA0430              ErrorCode = 0xA0430 // 用户输入内容非法
	ErrorA0431              ErrorCode = 0xA0431 // 包含违禁敏感词
	ErrorA0432              ErrorCode = 0xA0432 // 图片包含违禁信息
	ErrorA0433              ErrorCode = 0xA0433 // 文件侵犯版权
	ErrorA0440              ErrorCode = 0xA0440 // 用户操作异常
	ErrorA0441              ErrorCode = 0xA0441 // 用户支付超时
	ErrorA0442              ErrorCode = 0xA0442 // 确认订单超时
	ErrorA0443              ErrorCode = 0xA0443 // 订单已关闭
	ErrorA0500              ErrorCode = 0xA0500 // 用户请求服务异常
	ErrorA0501              ErrorCode = 0xA0501 // 请求次数超出限制
	ErrorA0502              ErrorCode = 0xA0502 // 请求并发数超出限制
	ErrorA0503              ErrorCode = 0xA0503 // 用户操作请等待
	ErrorA0504              ErrorCode = 0xA0504 // WebSocket 连接异常
	ErrorA0505              ErrorCode = 0xA0505 // WebSocket 连接断开
	ErrorA0506              ErrorCode = 0xA0506 // 用户重复请求
	ErrorA0600              ErrorCode = 0xA0600 // 用户资源异常
	ErrorA0601              ErrorCode = 0xA0601 // 账户余额不足
	ErrorA0602              ErrorCode = 0xA0602 // 用户磁盘空间不足
	ErrorA0603              ErrorCode = 0xA0603 // 用户内存空间不足
	ErrorA0604              ErrorCode = 0xA0604 // 用户 OSS 容量不足
	ErrorA0605              ErrorCode = 0xA0605 // 用户配额已用光
	ErrorA0700              ErrorCode = 0xA0700 // 用户上传文件异常
	ErrorA0701              ErrorCode = 0xA0701 // 用户上传文件类型不匹配
	ErrorA0702              ErrorCode = 0xA0702 // 用户上传文件太大
	ErrorA0703              ErrorCode = 0xA0703 // 用户上传图片太大
	ErrorA0704              ErrorCode = 0xA0704 // 用户上传视频太大
	ErrorA0705              ErrorCode = 0xA0705 // 用户上传压缩文件太大
	ErrorA0800              ErrorCode = 0xA0800 // 用户当前版本异常
	ErrorA0801              ErrorCode = 0xA0801 // 用户安装版本与系统不匹配
	ErrorA0802              ErrorCode = 0xA0802 // 用户安装版本过低
	ErrorA0803              ErrorCode = 0xA0803 // 用户安装版本过高
	ErrorA0804              ErrorCode = 0xA0804 // 用户安装版本已过期
	ErrorA0805              ErrorCode = 0xA0805 // 用户 API 请求版本不匹配
	ErrorA0806              ErrorCode = 0xA0806 // 用户 API 请求版本过高
	ErrorA0807              ErrorCode = 0xA0807 // 用户 API 请求版本过低
	ErrorA0900              ErrorCode = 0xA0900 // 用户隐私未授权
	ErrorA0901              ErrorCode = 0xA0901 // 用户隐私未签署
	ErrorA0902              ErrorCode = 0xA0902 // 用户摄像头未授权
	ErrorA0903              ErrorCode = 0xA0903 // 用户相机未授权
	ErrorA0904              ErrorCode = 0xA0904 // 用户图片库未授权
	ErrorA0905              ErrorCode = 0xA0905 // 用户文件未授权
	ErrorA0906              ErrorCode = 0xA0906 // 用户位置信息未授权
	ErrorA0907              ErrorCode = 0xA0907 // 用户通讯录未授权
	ErrorA1000              ErrorCode = 0xA1000 // 用户设备异常
	ErrorA1001              ErrorCode = 0xA1001 // 用户相机异常
	ErrorA1002              ErrorCode = 0xA1002 // 用户麦克风异常
	ErrorA1003              ErrorCode = 0xA1003 // 用户听筒异常
	ErrorA1004              ErrorCode = 0xA1004 // 用户扬声器异常
	ErrorA1005              ErrorCode = 0xA1005 // 用户 GPS 定位异常
	ErrorInternalError      ErrorCode = 0xB0001 // 系统执行出错
	ErrorB0100              ErrorCode = 0xB0100 // 系统执行超时
	ErrorB0101              ErrorCode = 0xB0101 // 系统订单处理超时
	ErrorB0200              ErrorCode = 0xB0200 // 系统容灾功能被触发
	ErrorB0210              ErrorCode = 0xB0210 // 系统限流
	ErrorB0220              ErrorCode = 0xB0220 // 系统功能降级
	ErrorB0300              ErrorCode = 0xB0300 // 系统资源异常
	ErrorB0310              ErrorCode = 0xB0310 // 系统资源耗尽
	ErrorB0311              ErrorCode = 0xB0311 // 系统磁盘空间耗尽
	ErrorB0312              ErrorCode = 0xB0312 // 系统内存耗尽
	ErrorB0313              ErrorCode = 0xB0313 // 文件句柄耗尽
	ErrorB0314              ErrorCode = 0xB0314 // 系统连接池耗尽
	ErrorB0315              ErrorCode = 0xB0315 // 系统线程池耗尽
	ErrorB0320              ErrorCode = 0xB0320 // 系统资源访问异常
	ErrorB0321              ErrorCode = 0xB0321 // 系统读取磁盘文件失败
	ErrorC0001              ErrorCode = 0xC0001 // 调用第三方服务出错
	ErrorC0100              ErrorCode = 0xC0100 // 中间件服务出错
	ErrorC0110              ErrorCode = 0xC0110 // RPC 服务出错
	ErrorC0111              ErrorCode = 0xC0111 // RPC 服务未找到
	ErrorC0112              ErrorCode = 0xC0112 // RPC 服务未注册
	ErrorC0113              ErrorCode = 0xC0113 // 接口不存在
	ErrorC0120              ErrorCode = 0xC0120 // 消息服务出错
	ErrorC0121              ErrorCode = 0xC0121 // 消息投递出错
	ErrorC0122              ErrorCode = 0xC0122 // 消息消费出错
	ErrorC0123              ErrorCode = 0xC0123 // 消息订阅出错
	ErrorC0124              ErrorCode = 0xC0124 // 消息分组未查到
	ErrorC0130              ErrorCode = 0xC0130 // 缓存服务出错
	ErrorC0131              ErrorCode = 0xC0131 // key 长度超过限制
	ErrorC0132              ErrorCode = 0xC0132 // value 长度超过限制
	ErrorC0133              ErrorCode = 0xC0133 // 存储容量已满
	ErrorC0134              ErrorCode = 0xC0134 // 不支持的数据格式
	ErrorC0140              ErrorCode = 0xC0140 // 配置服务出错
	ErrorC0150              ErrorCode = 0xC0150 // 网络资源服务出错
	ErrorC0151              ErrorCode = 0xC0151 // VPN 服务出错
	ErrorC0152              ErrorCode = 0xC0152 // CDN 服务出错
	ErrorC0153              ErrorCode = 0xC0153 // 域名解析服务出错
	ErrorC0154              ErrorCode = 0xC0154 // 网关服务出错
	ErrorC0200              ErrorCode = 0xC0200 // 第三方系统执行超时
	ErrorC0210              ErrorCode = 0xC0210 // RPC 执行超时
	ErrorC0220              ErrorCode = 0xC0220 // 消息投递超时
	ErrorC0230              ErrorCode = 0xC0230 // 缓存服务超时
	ErrorC0240              ErrorCode = 0xC0240 // 配置服务超时
	ErrorC0250              ErrorCode = 0xC0250 // 数据库服务超时
	ErrorC0300              ErrorCode = 0xC0300 // 数据库服务出错
	ErrorC0311              ErrorCode = 0xC0311 // 表不存在
	ErrorC0312              ErrorCode = 0xC0312 // 列不存在
	ErrorC0321              ErrorCode = 0xC0321 // 多表关联中存在多个相同名称的列
	ErrorC0331              ErrorCode = 0xC0331 // 数据库死锁
	ErrorC0341              ErrorCode = 0xC0341 // 主键冲突
	ErrorC0400              ErrorCode = 0xC0400 // 第三方容灾系统被触发
	ErrorC0401              ErrorCode = 0xC0401 // 第三方系统限流
	ErrorC0402              ErrorCode = 0xC0402 // 第三方功能降级
	ErrorC0500              ErrorCode = 0xC0500 // 通知服务出错
	ErrorC0501              ErrorCode = 0xC0501 // 短信提醒服务失败
	ErrorC0502              ErrorCode = 0xC0502 // 语音提醒服务失败
	ErrorC0503              ErrorCode = 0xC0503 // 邮件提醒服务失败
)

var messages = map[ErrorCode]string{
	0x00000: "一切 ok",
	0xA0001: "用户端错误",
	0xA0100: "用户注册错误",
	0xA0101: "用户未同意隐私协议",
	0xA0102: "注册国家或地区受限",
	0xA0110: "用户名校验失败",
	0xA0111: "用户名已存在",
	0xA0112: "用户名包含敏感词",
	0xA0113: "用户名包含特殊字符",
	0xA0120: "密码校验失败",
	0xA0121: "密码长度不够",
	0xA0122: "密码强度不够",
	0xA0130: "校验码输入错误",
	0xA0131: "短信校验码输入错误",
	0xA0132: "邮件校验码输入错误",
	0xA0133: "语音校验码输入错误",
	0xA0140: "用户证件异常",
	0xA0141: "用户证件类型未选择",
	0xA0142: "大陆身份证编号校验非法",
	0xA0143: "护照编号校验非法",
	0xA0144: "军官证编号校验非法",
	0xA0150: "用户基本信息校验失败",
	0xA0151: "手机格式校验失败",
	0xA0152: "地址格式校验失败",
	0xA0153: "邮箱格式校验失败",
	0xA0200: "用户登录异常",
	0xA0201: "用户账户不存在",
	0xA0202: "用户账户被冻结",
	0xA0203: "用户账户已作废",
	0xA0210: "用户密码错误",
	0xA0211: "用户输入密码错误次数超限",
	0xA0220: "用户身份校验失败",
	0xA0221: "用户指纹识别失败",
	0xA0222: "用户面容识别失败",
	0xA0223: "用户未获得第三方登录授权",
	0xA0230: "用户登录已过期",
	0xA0240: "用户验证码错误",
	0xA0241: "用户验证码尝试次数超限",
	0xA0300: "访问权限异常",
	0xA0301: "访问未授权",
	0xA0302: "正在授权中",
	0xA0303: "用户授权申请被拒绝",
	0xA0310: "因访问对象隐私设置被拦截",
	0xA0311: "授权已过期",
	0xA0312: "无权限使用 API",
	0xA0320: "用户访问被拦截",
	0xA0321: "黑名单用户",
	0xA0322: "账号被冻结",
	0xA0323: "非法 IP 地址",
	0xA0324: "网关访问受限",
	0xA0325: "地域黑名单",
	0xA0330: "服务已欠费",
	0xA0340: "用户签名异常",
	0xA0341: "RSA 签名错误",
	0xA0400: "用户请求参数错误",
	0xA0401: "包含非法恶意跳转链接",
	0xA0402: "无效的用户输入",
	0xA0410: "请求必填参数为空",
	0xA0411: "用户订单号为空",
	0xA0412: "订购数量为空",
	0xA0413: "缺少时间戳参数",
	0xA0414: "非法的时间戳参数",
	0xA0420: "请求参数值超出允许的范围",
	0xA0421: "参数格式不匹配",
	0xA0422: "地址不在服务范围",
	0xA0423: "时间不在服务范围",
	0xA0424: "金额超出限制",
	0xA0425: "数量超出限制",
	0xA0426: "请求批量处理总个数超出限制",
	0xA0427: "请求 JSON 解析失败",
	0xA0430: "用户输入内容非法",
	0xA0431: "包含违禁敏感词",
	0xA0432: "图片包含违禁信息",
	0xA0433: "文件侵犯版权",
	0xA0440: "用户操作异常",
	0xA0441: "用户支付超时",
	0xA0442: "确认订单超时",
	0xA0443: "订单已关闭",
	0xA0500: "用户请求服务异常",
	0xA0501: "请求次数超出限制",
	0xA0502: "请求并发数超出限制",
	0xA0503: "用户操作请等待",
	0xA0504: "WebSocket 连接异常",
	0xA0505: "WebSocket 连接断开",
	0xA0506: "用户重复请求",
	0xA0600: "用户资源异常",
	0xA0601: "账户余额不足",
	0xA0602: "用户磁盘空间不足",
	0xA0603: "用户内存空间不足",
	0xA0604: "用户 OSS 容量不足",
	0xA0605: "用户配额已用光",
	0xA0700: "用户上传文件异常",
	0xA0701: "用户上传文件类型不匹配",
	0xA0702: "用户上传文件太大",
	0xA0703: "用户上传图片太大",
	0xA0704: "用户上传视频太大",
	0xA0705: "用户上传压缩文件太大",
	0xA0800: "用户当前版本异常",
	0xA0801: "用户安装版本与系统不匹配",
	0xA0802: "用户安装版本过低",
	0xA0803: "用户安装版本过高",
	0xA0804: "用户安装版本已过期",
	0xA0805: "用户 API 请求版本不匹配",
	0xA0806: "用户 API 请求版本过高",
	0xA0807: "用户 API 请求版本过低",
	0xA0900: "用户隐私未授权",
	0xA0901: "用户隐私未签署",
	0xA0902: "用户摄像头未授权",
	0xA0903: "用户相机未授权",
	0xA0904: "用户图片库未授权",
	0xA0905: "用户文件未授权",
	0xA0906: "用户位置信息未授权",
	0xA0907: "用户通讯录未授权",
	0xA1000: "用户设备异常",
	0xA1001: "用户相机异常",
	0xA1002: "用户麦克风异常",
	0xA1003: "用户听筒异常",
	0xA1004: "用户扬声器异常",
	0xA1005: "用户 GPS 定位异常",
	0xB0001: "系统执行出错",
	0xB0100: "系统执行超时",
	0xB0101: "系统订单处理超时",
	0xB0200: "系统容灾功能被触发",
	0xB0210: "系统限流",
	0xB0220: "系统功能降级",
	0xB0300: "系统资源异常",
	0xB0310: "系统资源耗尽",
	0xB0311: "系统磁盘空间耗尽",
	0xB0312: "系统内存耗尽",
	0xB0313: "文件句柄耗尽",
	0xB0314: "系统连接池耗尽",
	0xB0315: "系统线程池耗尽",
	0xB0320: "系统资源访问异常",
	0xB0321: "系统读取磁盘文件失败",
	0xC0001: "调用第三方服务出错",
	0xC0100: "中间件服务出错",
	0xC0110: "RPC 服务出错",
	0xC0111: "RPC 服务未找到",
	0xC0112: "RPC 服务未注册",
	0xC0113: "接口不存在",
	0xC0120: "消息服务出错",
	0xC0121: "消息投递出错",
	0xC0122: "消息消费出错",
	0xC0123: "消息订阅出错",
	0xC0124: "消息分组未查到",
	0xC0130: "缓存服务出错",
	0xC0131: "key 长度超过限制",
	0xC0132: "value 长度超过限制",
	0xC0133: "存储容量已满",
	0xC0134: "不支持的数据格式",
	0xC0140: "配置服务出错",
	0xC0150: "网络资源服务出错",
	0xC0151: "VPN 服务出错",
	0xC0152: "CDN 服务出错",
	0xC0153: "域名解析服务出错",
	0xC0154: "网关服务出错",
	0xC0200: "第三方系统执行超时",
	0xC0210: "RPC 执行超时",
	0xC0220: "消息投递超时",
	0xC0230: "缓存服务超时",
	0xC0240: "配置服务超时",
	0xC0250: "数据库服务超时",
	0xC0300: "数据库服务出错",
	0xC0311: "表不存在",
	0xC0312: "列不存在",
	0xC0321: "多表关联中存在多个相同名称的列",
	0xC0331: "数据库死锁",
	0xC0341: "主键冲突",
	0xC0400: "第三方容灾系统被触发",
	0xC0401: "第三方系统限流",
	0xC0402: "第三方功能降级",
	0xC0500: "通知服务出错",
	0xC0501: "短信提醒服务失败",
	0xC0502: "语音提醒服务失败",
	0xC0503: "邮件提醒服务失败",
}

func ClientError(c *app.RequestContext, err error) {
	ErrorClient.With(err.Error()).Write(c)
}

func Error(c *app.RequestContext, err error) {
	switch e := err.(type) {
	case ErrorCode:
		e.Write(c)
	case ErrorWithMessage:
		e.Write(c)
	default:
		hlog.Warn("unrecognized error", err)
		c.SetContentType("application/json")
		c.String(
			consts.StatusOK, `{"status_code":0x%x,"status_msg":"%s"}`,
			ErrorInternalError, err.Error(),
		)
	}
}

func (code ErrorCode) Write(c *app.RequestContext) {
	if message, ok := messages[code]; ok {
		c.SetContentType("application/json")
		c.String(consts.StatusOK, `{"status_code":0x%x,"status_msg":"%s"}`, code, message)
	} else {
		hlog.Error("unrecognized error code", code)
		ErrorInternalError.Write(c)
	}
}

func (code ErrorCode) Wrap(err error) ErrorWithMessage {
	return code.With(err.Error())
}

func (code ErrorCode) With(message string) ErrorWithMessage {
	return ErrorWithMessage{
		Code:    code,
		Message: message,
	}
}

func (err ErrorWithMessage) Write(c *app.RequestContext) {
	c.SetContentType("application/json")
	c.String(
		consts.StatusOK, `{"status_code":0x%x,"status_msg":"%s"}`,
		err.Code, err.Message,
	)
}

func (code ErrorCode) Error() string {
	return messages[code]
}

func (err ErrorWithMessage) Error() string {
	return err.Message
}
