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
	methodPOST = "POST"
	methodGET  = "GET"
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

// Payment Transaction Endpoint
var pathAPIGetPaymentTransactionByIDURL = newPath(methodGET, "/payment/transaction")
var pathAPIRefundPaymentURL = newPath(methodPOST, "/payment/refund")
var pathAPIReversedPaymentURL = newPath(methodPOST, "/payment/reverse")

// Stores Endpoint
var pathAPIGetStoresURL = newPath(methodGET, "/stores")

// Merchant Information
var pathAPIGetMerchantURL = newPath(methodGET, "/merchant")

// WeChat
var pathAPIGetRMWeChatUserOAuthURL = newPath(methodPOST, "/socialmedia/rm/wechat-oauth/url")
var pathAPIGetRMWeChatUserInfoByCodeURL = newPath(methodPOST, "/socialmedia/rm/wechat-oauth/code")
var pathAPISendWeChatPageTemplateMessageURL = newPath(methodPOST, "/socialmedia/wechat/template-message")
var pathAPIGetWeChatPageAccessTokenURL = newPath(methodGET, "/socialmedia/wechat/access-token")

// Event
var pathAPIPostSendEventURL = newPath(methodPOST, "/event")
