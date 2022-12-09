package constant

const (
	HeaderXParams = "X-Params" //http自定义参数
)

const (
	ContextUID             = "uid"               //用户ID
	ContextXParams         = "x_params"          //上下文中的x_params
	ContextHttpRequestBody = "http_request_body" //http请求body（context元数据）
)

const (
	IdentityTypeWeChat = "wechat" //用户第三方登录-微信
	IdentityTypeApple  = "apple"  //用户第三方登录-苹果
	IdentityTypeQQ     = "qq"     //用户第三方登录-qq
)

// jwt业务类型
const (
	JwtTypeApp   = "App"
	JwtTypeAdmin = "Admin"
)
const (
	SmsPlatformJSms = "jsms" //短信平台类型-极光
)

// 短信发送类型
const (
	SmsTypeLogIn        = "login"         //登录
	SmsTypeRegister     = "register"      //注册
	SmsTypeAccountCheck = "account_check" //账号检查
	SmsTypeAccountBind  = "account_bind"  //账号绑定
)
