package sdk

// Sign Type
const (
	SHA256 = "sha256"
)

// Scopes
const (
	ScopeManagePushNotication    string = "manage_push_notification"
	ScopeManagePayment           string = "manage_payment"
	ScopeManageStore             string = "manage_store"
	ScopeManageSocialMedia       string = "manage_social_media"
	ScopeGetMerchantProfile      string = "get_merchant_profile"
	ScopeGetUserProfile          string = "get_user_profile"
	ScopeGetMerchantSubscription string = "get_merchant_subscription"
)

// private function
// Method
const (
	methodPOST   = "POST"
	methodGET    = "GET"
	methodDELETE = "DELETE"
)

// Default
const (
	defaultAPIVersion   = "v3"
	defaultOAuthVersion = "v1"
	defaultSignType     = SHA256
)

// api domain
const productionAPIURL = "https://open.revenuemonster.my"
const sandboxAPIURL = "https://sb-open.revenuemonster.my"

// oauthdomain
const productionOAuthURL = "https://oauth.revenuemonster.my"
const sandboxOAuthURL = "https://sb-oauth.revenuemonster.my"

// Path
// Oauth path
var pathOAuthPathTokenURL = newPath(methodPOST, "/token")

// API path

// QRCode Endpoint
var pathAPICreatePaymentTransactionQRURL = newPath(methodPOST, "/payment/transaction/qrcode")
var pathAPIGetPaymentTransactionQRURL = newPath(methodGET, "/payment/transaction/qrcode")

// Checkout Endpoint
var pathAPICreatePaymentCheckoutURL = newPath(methodPOST, "/payment/online")
var pathAPIGetQRCodeByCheckoutIDURL = newPath(methodGET, "/payment/online/qrcode")
var pathAPICreatePaymentCheckoutByMethodURL = newPath(methodPOST, "/payment/online/checkout")
var pathAPIGetOnlineTransactionByCheckoutIDURL = newPath(methodGET, "/payment/online")

// Payment Transaction Endpoint
var pathAPIGetPaymentTransactionByIDURL = newPath(methodGET, "/payment/transaction")
var pathAPIGetPaymentTransactionByOrderIDURL = newPath(methodGET, "/payment/transaction/order")
var pathAPIRefundPaymentURL = newPath(methodPOST, "/payment/refund")
var pathAPIReversedPaymentURL = newPath(methodPOST, "/payment/reverse")

// Stores Endpoint
var pathAPIGetStoresURL = newPath(methodGET, "/stores")
var pathAPIGetStoreByIDURL = newPath(methodGET, "/store")

// Merchant Information
var pathAPIGetMerchantURL = newPath(methodGET, "/merchant")

// WeChat
var pathAPIGetRMWeChatUserOAuthURL = newPath(methodPOST, "/socialmedia/rm/wechat-oauth/url")
var pathAPIGetRMWeChatUserInfoByCodeURL = newPath(methodPOST, "/socialmedia/rm/wechat-oauth/code")
var pathAPISendWeChatPageTemplateMessageURL = newPath(methodPOST, "/socialmedia/wechat/template-message")
var pathAPIGetWeChatPageAccessTokenURL = newPath(methodGET, "/socialmedia/wechat/access-token")

// Event
var pathAPIPostSendEventURL = newPath(methodPOST, "/event")
var pathAPIPostSendEventByStoreURL = newPath(methodPOST, "/event/store")

// PushNotification
var pathAPIPushNotificationToStoreURL = newPath(methodPOST, "/push-notification/store")

// DNS
var pathCreateDNSRecordURL = newPath(methodPOST, "/dns/record")
var pathGetDNSRecordsURL = newPath(methodGET, "/dns/records")
var pathDeleteDNSRecordURL = newPath(methodDELETE, "/dns/record")

// SMS
var pathSendSms = newPath(methodPOST, "/sms")

// Delivery
var pathCreateDelivery = newPath(methodPOST, "/delivery")
var parhGetDeliveryByID = newPath(methodGET, "/delivery")
var pathCalculateDeliveryFee = newPath(methodPOST, "/delivery/calculate")
var pathConfirmDelivery = newPath(methodPOST, "/delivery/confirm")
var pathCancelDelivery = newPath(methodPOST, "/delivery/cancel")

// Voucher
var pathAPIVoucherVoidURL = newPath(methodPOST, "/voucher")
