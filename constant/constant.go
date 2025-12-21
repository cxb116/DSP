package constant

// 在redis中存储控制台用户令牌的密钥
const RedisSubscribeMessageChannel = "redis:subscribe-message-channel"

const NOT_WORKER = -1 // 没有可用workerId

const (
	EventAllAdd = 1 // 后台首次发布全部数据
	EventAdd    = 2 // 后台添加数据
	EventRemove = 3 // 后台删除数据
	EventUpdate = 4 // 后台跟新数据

	EventCompayAdd    = 5
	EventCompayRemove = 6
	EventCompayUpdate = 7
)

// 同一个预算下的流量权重
const TrafficValue = 100 // 分流总值

// 响应状态
const (
	RES_CODE_SUC         = iota + 1 // 有填充
	RES_CODE_NIL                    // 无填充
	RES_CODE_SLOT_NO                // 广告位Id 不存在
	RES_CODE_PARAM_ERR              // 请求参数解析错误
	RES_CODE_PARAM_LACK             //
	RES_CODE_PARAM_INVAL            // 参数异常
	RES_CODE_TIMEOUT                //  响应超时

)

const (
	REQ_CODE_FAIL             = 0  // 请求失败，这个是统一失败，用来判断是否成功的标准
	REQ_CODE_SLOT_ID          = -1 // 广告位id不存在
	REQ_CODE_AD_TYPE          = -2 // 广告类型不存在
	REQ_CODE_DEVICE_IP        = -3 // 设备IP
	REQ_CODE_DEVICE_UA        = -4 //
	REQ_CODE_DEVICE_OS        = -5
	REQ_CODE_DEVICE_OSV       = -6
	REQ_CODE_DEVICE_TYPE      = -7
	REQ_CODE_DEVICE_CAID      = -8 //  IOS 设备必穿
	REQ_CODE_DEVICE_CAID_VER  = -9 //
	REQ_CODE_DEVICE_IDFA      = -10
	REQ_CODE_DEVICE_IDFV      = -11
	REQ_CODE_DEVICE_NAME      = -12
	REQ_CODE_NETWOEK_CON_TYPE = -13 // 网路类型
	REQ_CODE_NETWOEK_CARRIER  = -14 //网络运营商不存在
	REQ_CODE_NETWOEK_IMSI     = -15
	REQ_CODE_DEVICE_ANDROID   = -16
	REQ_CODE_DEVICE_OAID      = -17
)

// 结算方式
const (
	PAY_PRICE_FIXED      = 1 // 固价
	PAY_PRICE_PROPORTION = 2 // 分成
	PAY_PRICE_RTB        = 3 // RTB

)

const (
	POST_METHOD = "POST"
	GET_METHOD  = "GET"
)
