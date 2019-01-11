package app

var (
	SUCCESS = BaseMicroServiceCode{"0000","","成功"}
	FAIL    = BaseMicroServiceCode{"9999","","失败"}
	BUSINESS_ERROR    = BaseMicroServiceCode{"0001","business_error","业务错误"}
	VALIDATE_ERROR    = BaseMicroServiceCode{"0002","validate_error","验证错误"}
	PARAMATER_ERROR    = BaseMicroServiceCode{"0002","paramater_error","参数错误"}
	ERROR_AUTH_CHECK_TOKEN_FAIL =     BaseMicroServiceCode{"","error_auth_check_token_fail","Token鉴权失败"}
    ERROR_AUTH_CHECK_TOKEN_TIMEOUT =   BaseMicroServiceCode{"","error_auth_check_token_timeout","Token已超时"}
    ERROR_AUTH_TOKEN =     BaseMicroServiceCode{"","error_auth_token","Token生成失败"}
    ERROR_AUTH =         BaseMicroServiceCode{"","error_auth","Token错误"}
)