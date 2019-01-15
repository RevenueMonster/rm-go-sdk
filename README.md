# API-SDK-Go
This is an Go SDK that maps some of the RESTful methods of Open API that are documented at [doc.revenuemonster.my](https://doc.revenuemonster.my/).

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

The golang version 1.11 and above

### Covered Functions
- [x] Signature Algorithm
- [x] Client Credentials (Authentication)
- [x] Refresh Token (Authentication)
- [ ] Get Merchant Profile
- [ ] Get Merchant Subscriptions
- [x] Get Stores
- [ ] Get Stores By ID
- [ ] Create Store
- [ ] Update Store
- [ ] Delete Store
- [ ] Get User Profile
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
- [ ] Give Loyalty Point
- [ ] Get Loyalty Members
- [ ] Get Loyalty Member
- [ ] Get Loyalty Member Point History
- [ ] Issue Voucher
- [ ] Void Voucher
- [ ] Get Voucher By Code
- [ ] Get Voucher Batches
- [ ] Get Voucher Batch By Key
- [ ] Send Notification (Merchant)
- [ ] Send Notification (Store)
- [ ] Send Notification (User)

### Usage
1. "sandbox" is for sandbox environment.
2. "production" is for production environment.
3. Get Client ID and Client Secret from portal.
![ClientIDClientSecret](https://storage.googleapis.com/rm-portal-assets/img/rm-landing/clientIDclientSecret.png)
4. Generate private key and publci key from portal.
![PrivateKeyPublicKey](https://storage.googleapis.com/rm-portal-assets/img/rm-landing/privateKeypublicKey.PNG)
5. Store private key for own use and public key at portal.
![PastePublicKey](https://storage.googleapis.com/rm-portal-assets/img/rm-landing/pastePublicKey.png)
6. Set environment variables at begining of the project before using any of the library functions.
```
Environment environment = new Environment();
environment.setEnvironment(clientId, clientSecret, "sandbox");
```

* Sample to read private key file
```

```

* Client Credentials (Authentication)
    * To get refresh token and access token(expired after 2 hours) with using provided clientId and clientSecret
```

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
```
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
``
```



