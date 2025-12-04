# API-SDK-Go
This is an Go SDK that maps some of the RESTful methods of Open API that are documented at [doc.revenuemonster.my](https://doc.revenuemonster.my/).

## 🚀 Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

The golang version 1.11 and above

### Covered Functions
- [x] Signature Algorithm
- [x] Client Credentials (Authentication)
- [x] Refresh Token (Authentication)
- [x] Get Merchant Profile
- [x] Get Merchant Subscriptions
- [x] Get Merchant Settlement
- [x] Get Merchant Settlement List
- [x] Get Merchant Settlement By ID
- [x] Get Stores
- [x] Get Store By ID
- [ ] Create Store
- [ ] Update Store
- [ ] Delete Store
- [x] Get User Profile
- [x] Payment (Transaction QR) - Create Transaction QRCode/URL
- [] Payment (Transaction QR) - Get Transaction QRCode/URL
- [x] Payment (Transaction QR) - Get Transaction QRCode/URL By Code
- [x] Payment (Transaction QR) - Get Transactions By Code
- [x] Payment (Quick Pay) - Payment
- [x] Payment (Quick Pay) - Refund
- [x] Payment (Quick Pay) - Reverse
- [ ] Payment (Quick Pay) - Get All Payment Transactions
- [x] Payment (Quick Pay) - Get All Payment Transaction By ID
- [] Payment (Quick Pay) - Get All Payment Transaction By OrderID
- [ ] Payment (Quick Pay) - Daily Settlement Report
- [x] Loyalty - Check Member Exist
- [x] Loyalty - Register Member
- [x] Loyalty - Get Loyalty Members
- [x] Loyalty - Get Loyalty Member By ID
- [x] Loyalty - Loyalty Member Top Up Balance Online
- [x] Loyalty - Spend Balance
- [x] Loyalty - Refund Balance
- [ ] Give Loyalty Point
- [ ] Get Loyalty Members
- [ ] Get Loyalty Member Point History
- [x] Voucher - Get Voucher By Code
- [x] Voucher - Get Voucher Batches
- [x] Voucher - Get Voucher Batch By Key
- [x] Voucher - Void Voucher
- [x] Voucher - Reinstate Voucher
- [x] Voucher - Bulk Redeem Voucher
- [ ] Send Notification (Merchant)
- [x] Send Notification (Store)
- [ ] Send Notification (User)
- [x] Sms Sending
- [x] Create Delivery
- [x] Get Delivery By Id
- [x] Calculate Delivery Fee
- [x] Ekyc - MyKad Recognition
- [x] EKyc - Face Comparison
- [x] Ekyc - Liveness Verification

### Usage
### Environment Setup
1. Choose environment:  
   - `"sandbox"` – Sandbox testing  
   - `"production"` – Production  
2. Retrieve **Client ID** & **Client Secret** from the portal. 
![ClientIDClientSecret](https://storage.googleapis.com/rm-portal-assets/img/rm-landing/clientIDclientSecret.png)
3. Generate **private** and **public** keys. 
![PrivateKeyPublicKey](https://storage.googleapis.com/rm-portal-assets/img/rm-landing/privateKeypublicKey.PNG)
4. Store **private key locally** and upload the **public key** to the portal.
![PastePublicKey](https://storage.googleapis.com/rm-portal-assets/img/rm-landing/pastePublicKey.png)
6. Set environment variables at begining of the project before using any of the library functions.
```
Environment environment = new Environment();
environment.setEnvironment(clientId, clientSecret, "sandbox");
```

* Client Credentials (Authentication)
    * To get refresh token and access token(expired after 2 hours) with using provided clientId and clientSecret
```go
token, err := sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
}).GetAccessTokenByClientCredentials()
```

* Authorization Code (Authentication)
    * To get authorization code which a partner wants to request permission to develop an application of a merchant and exchange authorization code into access token and refresh token
```go
  oauthURL := sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
  }).GetAuthorizationCodeURL(
    "123456789",
    "http://google.com",
    "manage_payment",
    "manage_store",
  )
```

* Refresh Token (Authentication)
    * To get new access token(expired after 2 hours) with using provided clientId and clientSecret (recommended to schedule to run this fucntion on every less than 2 hours) in order to avoid expired access token error
```go
token, err := sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).GetAccessTokenByRefreshToken("refresh_token_value")
```

* Get Merchant Profile
    * To get Merchant Profile
```go
response, err := sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).GetMerchantProfile("") 
```

* Get Merchant Settlement List
    * To get list of all merchant settlement accounts
```go
response, err := sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).GetMerchantSettlementList("") // cursor parameter for pagination
```

* Get Merchant Settlement (Default)
    * To get the default active merchant settlement account
```go
response, err := sdk.NewClient(sdk.Client{
   ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).GetMerchantSettlement()
```

* Get Merchant Settlement By ID
    * To get specific merchant settlement account by ID
```go
response, err := sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).GetMerchantSettlementById("settlement_account_id")
```

* Create Transaction QRCode/URL (TransactionQR)
    * To create static/dynamic QR code for user scanning merchant's displayed QR
```
```

* Get Transaction QRCode/URL (TransactionQR)
    * To get all QR Code(s) generated previously in the system
```
```

* Get Transaction QRCode/URL By Code (TransactionQR)
    * To get specific QR Code generated previously in the system, by passing in code in query parameter (/qrcode/...)
```
```

* Get Transactions By Code (TransactionQR)
    * To get all transactions under existing QR code, by passing in code in query parameter (/qrcode/.../transactions)
```
```

* Payment (Quick Pay) - Payment
    * To make payment by scanning barcode/authcode from user
```
```

* Payment (Quick Pay) - Refund
    * To refund the successful transactions
```
```

* Payment (Quick Pay) - Reverse
    * To reverse time-out or problematic transaction
```
```

* Payment (Quick Pay) - Get Payment Transaction By ID
    * To get details of a transaction by using transactionId
```
```

* Payment (Quick Pay) - Get Payment Transaction By Order ID
    * To get details of a transaction by using orderId
```
```

* Push Notification ( Store ) - Send Notificaiton To Store By Store ID
```go
sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).PushNotificationToStore(sdk.RequestPushNotificationToStore{
    StoreID: "123123123",
    Title:   "Notification Title",
    Body:    "Notification Body",
})
```

* Sms Sending - Send Sms To Specified Phone Number
```go
sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).SendSms(sdk.RequestSendSms{
    CountryCode: "60",
    PhoneNumber: "187824152",
    Message: "Some message",
    Type: sdk.MessageTypeTAC,
})
```

* Create Delivery - Create delivery 
```go
pointOne := sdk.DeliveryPoint{
    Address: "",
    EntraceNumber: "",
    FloorNumber: "",
    BuildingNumber: "",
    Remark: "",
    Contact: sdk.DeliveryPointContact{
        Name: "",
        PhoneNumber: "",
    },
}

pointTwo := sdk.DeliveryPoint{
    Address: "",
    EntraceNumber: "",
    FloorNumber: "",
    BuildingNumber: "",
    Remark: "",
    Contact: sdk.DeliveryPointContact{
        Name: "",
        PhoneNumber: "",
    },
}

sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).CreateDelivery(sdk.RequestCreateDelivery{
    DeliveryVendor: sdk.DeliveryVendor{
        Vendor: sdk.VendorTypeMrSpeedy,
        Credential: "",
    },
    VehicleType: sdk.VehicleTypeMotobike,
    Type: sdk.DeliveryTypeDocument,
    IsCashAccount: true,
    Points: []sdk.DeliveryContact{pointOne, pointTwo},
})
```

* Calculate Delivery Fee - Calculate delivery fee before delivery is made 
```go
pointOne := sdk.DeliveryPoint{
    Address: "",
    EntraceNumber: "",
    FloorNumber: "",
    BuildingNumber: "",
    Remark: "",
    Contact: sdk.DeliveryPointContact{
        Name: "",
        PhoneNumber: "",
    },
}

pointTwo := sdk.DeliveryPoint{
    Address: "",
    EntraceNumber: "",
    FloorNumber: "",
    BuildingNumber: "",
    Remark: "",
    Contact: sdk.DeliveryPointContact{
        Name: "",
        PhoneNumber: "",
    },
}

sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).CalculateDeliveryFee(sdk.RequestCalculateDeliveryFee{
    DeliveryVendor: sdk.DeliveryVendor{
        Vendor: sdk.VendorTypeMrSpeedy,
        Credential: "",
    },
    VehicleType: sdk.VehicleTypeMotobike,
    Type: sdk.DeliveryTypeDocument,
    IsCashAccount: true,
    Points: []sdk.DeliveryContact{pointOne, pointTwo},
})
```

* Get Delivery - Get delivery by id
```go
sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).GetDeliveryByID("1")
```

* Loyalty - Check Member Exist
    * To check if a loyalty member exists by phone number
```go
response, err := sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).CheckMemberExist(sdk.RequestCheckMemberExist{
    PhoneNumber: "187824152",
    CountryCode: "60",
})

```

* Loyalty - Register Member
    * To register a new loyalty member
```go
response, err := sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).RegisterLoyaltyMember(sdk.RequestRegisterLoyaltyMember{
    ID:          "member123",
    Name:        "John Doe",
    Email:       "john@example.com",
    CountryCode: "60",
    PhoneNumber: "187824152",
    Gender:      "M",
    State:       "KL",
    NRIC:        "123456789012",
    Status:      "ACTIVE",
    LoyaltyPoint: 0,
    Credit:      0,
})

```

* Loyalty - Get Loyalty Member By ID
    * To retrieve loyalty member details by member ID
```go
response, err := sdk.NewClient(sdk.Client{
   ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).GetLoyaltyMemberByID(sdk.RequestGetLoyaltyMemberByID{
    ID: "member123",
})

```

* Loyalty - Get Loyalty Member By Phone Numeber
    * To retrieve loyalty member details by phone number
```go
response, err := sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).GetLoyaltyMember(sdk.RequestGetLoyaltyMember{
    CountryCode: "60",
    PhoneNumber: "187824152",
})

```

* Loyalty - Check Spend Balance
    * To spend loyalty points from a member's balance
```go
response, err := sdk.NewClient(sdk.Client{
   ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).SpendBalance(sdk.RequestSpendBalance{
    CountryCode: "60",
    PhoneNumber: "187824152",
    Amount:      1000,
    Remarks:     "Purchase transaction",
})

```

* Loyalty - Refund Balance
    * To refund loyalty points to a member's balance
```go
response, err := sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).RefundBalance(sdk.RequestRefundBalance{
    CountryCode: "60",
    PhoneNumber: "187824152",
    Amount:      500,
    Remarks:     "Refund for returned item",
})

```

* Voucher - Get Voucher By Code
    * To retrieve voucher details by code
```go
response, err := sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).GetVoucherByCode("VOUCHER123")
```

* Voucher - Get Voucher Batches
    * To get list of voucher batches
```go
response, err := sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).GetVoucherBatches(sdk.RequestGetVoucherBatches{
    Status:   "ACTIVE",
    IsStatic: true,
    Cursor:   "",
})
```

* Voucher - Get Voucher Batch By Key
    * To get vouchers in a specific batch
```go
response, err := sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).GetVoucherBatchByKey(sdk.RequestGetVoucherBatchByKey{
    Key:    "batch_key_123",
    Cursor: "",
})
```

* Voucher - Bulk Redeem Voucher
    * To redeem multiple vouchers at once
```go
response, err := sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey: []byte(`---private key---`),
    PublicKey:  []byte(`---public key---`),
}).BulkRedeemVoucher(sdk.BulkRedeemVoucherRequest{
    Codes: []string{"VOUCHER001", "VOUCHER002", "VOUCHER003"},
})
```

* Ekyc - MyKad Recognition
```go
sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey:  []byte(`---private key---`),
    AccessToken: `access token`,
}).EkycMyKad(sdk.RequestEkycMykad{
    Image: `base64 image`, // an image that contains MyKad
    NotifyUrl: `https://your-backend-notify-path` // a POST request handler to receive ekyc result
})
```

* Ekyc - Face Comparison
```go
sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey:  []byte(`---private key---`),
    AccessToken: `access token`,
}).EkycMyKad(sdk.RequestEkycFaceCompare{
    Image1: `base64 image`, // an image that contains a face
    Image2: `base64 image`, // an image that contains a face
})
```

* Ekyc - Liveness Verification
```go
sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey:  []byte(`---private key---`),
    AccessToken: `access token`,
}).EkycLiveness(sdk.RequestEkycLiveness{
    Image:          `base64 image`, // an image that contains a face
    MykadRequestID: `123456789`, // request ID from MyKad Recognition
})
```

* Ekyc - Get Mykad Result
```go
sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey:  []byte(`---private key---`),
    AccessToken: `access token`,
}).GetMykadResult(sdk.RequestGetMykadResult{
    ID:          `123456789`, // the mykad request id
})
```

* Ekyc - Get Landmarks Info
```go
sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey:  []byte(`---private key---`),
    AccessToken: `access token`,
}).GetLandmarkInfo(sdk.RequestGetLandmarkInfo{
    ID:          `123456789`, // the mykad request id
    Includes: ["msc_logo", "malaysia_flag"], // optional parameter to specify landmark types to be acquired
})
```

* Ekyc - Get eKYC Result
```go
sdk.NewClient(sdk.Client{
    ID:         "123456789",
    Secret:     "123456789",
    IsSandbox:  true,
    PrivateKey:  []byte(`---private key---`),
    AccessToken: `access token`,
}).GetEkycResult(sdk.RequestGetEkycResult{
    ID:          `123456789`, // the ekycId obtained after liveness verification
    Includes: ["mykadImage", "selfieImage"], // optional parameter to obtain request images
})
```