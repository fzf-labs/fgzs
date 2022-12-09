package errorx

import "net/http"

// @Description: 默认为中文错误码,在NewError()时请传入中文
// @return
var (
	Success = NewError(200, "成功")
	Fail    = NewError(500, "失败")

	//服务级错误
	InternalServerError   = NewError(10001, "服务器发生异常", WithLevel(ErrLevel))
	ErrServiceUnavailable = NewError(10002, "服务不可用", WithLevel(ErrLevel))
	UnknownError          = NewError(10003, "未知错误,请联系管理员", WithLevel(ErrLevel))
	ErrDataException      = NewError(10004, "数据异常", WithLevel(ErrLevel))
)

// 请求相关
var (
	TooManyRequests           = NewError(10101, "请求过多")
	RequestFrequencyIsTooFast = NewError(10105, "请求频率太快了")
	NoAccess                  = NewError(10106, "无访问权限")
)

// 路由权限相关
var (
	RouteNoAccess                     = NewError(10201, "无权限访问")
	RoutePermissionVerificationFailed = NewError(10203, "路由权限校验失败")
	RouteMethodNoAccess               = NewError(10204, "无访问该路由权限")
)

// 参数相关
var (
	ParamBindErr           = NewError(20001, "参数绑定到结构时发生错误", WithLevel(WarnLevel))
	ParamErr               = NewError(20002, "参数有误", WithLevel(WarnLevel))
	ParamValidationErr     = NewError(20003, "参数验证失败")
	ParamNotJsonRequest    = NewError(20004, "请使用JSON请求", WithLevel(WarnLevel))
	ParamMissingChannelID  = NewError(20005, "缺少渠道ID")
	ParamMissingDeviceType = NewError(20006, "缺少设备类型")
)

// 数据查询相关
var (
	DataSqlErr           = NewError(20100, "数据异常(S)", WithLevel(ErrLevel))
	DataRedisErr         = NewError(20101, "数据异常(R)", WithLevel(ErrLevel))
	DataRecordNotFound   = NewError(20102, "数据不存在")
	DataDuplicateRecords = NewError(20103, "记录重复")
	DataFormattingError  = NewError(20104, "数据格式化错误", WithLevel(ErrLevel))
)

// token,签名 ,校验相关
var (
	TokenNotRequest          = NewError(20200, "请求中未携带令牌", WithHttpCode(http.StatusUnauthorized))
	TokenFormatErr           = NewError(20201, "令牌格式化错误", WithHttpCode(http.StatusUnauthorized))
	TokenErr                 = NewError(20202, "错误的token", WithHttpCode(http.StatusUnauthorized))
	TokenInvalidErr          = NewError(20203, "令牌无效", WithHttpCode(http.StatusUnauthorized))
	TokenExpired             = NewError(20205, "令牌过期", WithHttpCode(http.StatusUnauthorized))
	TokenRefreshErr          = NewError(20206, "令牌刷新失败", WithHttpCode(http.StatusUnauthorized))
	TokenVerificationFailed  = NewError(20207, "您的登录状态已失效,或在其他设备登录,请您重新登录", WithHttpCode(http.StatusUnauthorized))
	TokenStorageFailed       = NewError(20208, "令牌储存失败", WithHttpCode(http.StatusUnauthorized))
	TokenErrSignatureParam   = NewError(20209, "签名参数缺失", WithHttpCode(http.StatusUnauthorized))
	TokenWrongTypeOfBusiness = NewError(20210, "错误的业务类型", WithHttpCode(http.StatusUnauthorized))
	TokenGenerationFailed    = NewError(20211, "Token生成失败")
)

// 文件上传
var (
	FileParsingError            = NewError(20301, "文件解析错误")
	FileNotExist                = NewError(20302, "上传文件不存在")
	FileError                   = NewError(20303, "文件错误")
	FileClassificationException = NewError(20304, "文件分类异常")
	FileOSSUploadException      = NewError(20305, "OSS上传异常")
)
var (
	SmsSendOverClock    = NewError(20400, "短信发送超频")
	SmsCodeInvalid      = NewError(20401, "短信验证码无效")
	SmsCodeExpired      = NewError(20402, "短信验证码未发送或已失效,请重新发送")
	SmsCodeVerified     = NewError(20403, "短信验证码已验证")
	SmsRepeatSend       = NewError(20404, "短信重复发送")
	SmsRequestOverClock = NewError(20405, "短信请求超频")
	SmsSendFailed       = NewError(20406, "短信发送失败")
	SmsTimesLimit       = NewError(20407, "同一手机号,一天只能发%s次")
	SmsCodeBeenSent     = NewError(20408, "短信发送频繁，请%s秒后重试")
	SmsTypeErr          = NewError(20409, "短信类型错误")
)

// 用户登录注册相关
var (
	EnterTheCorrectPhoneNumber          = NewError(20909, "请填写正确手机号")
	OneClickLoginFailed                 = NewError(20910, "一键登录失败")
	OneClickLoginAuthFailed             = NewError(20911, "一键登录认证失败")
	GuestAccountHasExpired              = NewError(20912, "游客账号已失效,请重新登录")
	WrongGuestAccount                   = NewError(20913, "错误的游客账号,请重新登录")
	AppleCodeCannotBeEmpty              = NewError(20914, "苹果登录Code不能为空")
	FailedToAppleUserID                 = NewError(20915, "苹果用户ID获取失败")
	LoginTokenDoesNotExist              = NewError(20916, "登录token不存在")
	WeChatCodeCannotBeEmpty             = NewError(20917, "微信登录Code不能为空")
	FailedToGetWeChatUserID             = NewError(20918, "微信用户ID获取失败")
	FailedToObtainWeChatUserInformation = NewError(20919, "微信用户信息获取失败")
	PleaseSignIn                        = NewError(20920, "请登录")
	QQCodeCannotBeEmpty                 = NewError(20922, "QQ登录Code不能为空")
	FailedToObtainQQChatUserInformation = NewError(20923, "QQ用户信息获取失败")
)

// 用户账号
var (
	UsernameSecurityIsLow                = NewError(21001, "用户名安全性低")
	UserPasswordSecurityIsLow            = NewError(21002, "用户密码安全性低")
	UserPasswordNotMatch                 = NewError(21003, "用户密码不匹配")
	UserNotFound                         = NewError(21004, "用户不存在")
	UserFailedToObtainAccountInformation = NewError(21005, "用户信息获取失败")
	UserIsLocked                         = NewError(21006, "用户已锁定.请联系客服")
	UserIsLoggedOut                      = NewError(21007, "用户已注销")
	UserWrongPassword                    = NewError(21008, "用户密码错误")
	UserIsBanned                         = NewError(21009, "用户封禁中")
	UserUpdateFailed                     = NewError(21010, "用户更新失败")
	UsernameDuplicate                    = NewError(21011, "用户名重复")
)

// 通用
var (
	MenuListIsEmpty                = NewError(21100, "菜单列表为空")
	FailedToLoadSensitiveThesaurus = NewError(21101, "敏感词库加载失败")
	CaptchaCodeError               = NewError(21102, "验证码错误")
)

// 用户资产
var (
	AssetException         = NewError(21200, "您的账户资产异常,请联系管理员")
	WrongTypeOfConsumption = NewError(21201, "错误的消费类型")
	DoNotRepeatConsumption = NewError(21202, "请勿重复消费,造成您的资产损失")
	InsufficientBalance    = NewError(21203, "余额不足")
)
