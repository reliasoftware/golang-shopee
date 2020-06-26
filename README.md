# Official document
https://open.shopee.com/documents

# Setup package

```bash
$ go get github.com/reliasoftware/golang-shopee
```

# Usage

```golang
client := shopeego.NewClient(&ClientOptions{
	Secret: "0c2c7b3bd59c2503f49a307fcf49dc985df43a1214821d5453e9ba54ca8e2e44",
})

// Get data base on request
shop := client.GetShopInfo(&GetShopInfoRequest{
	PartnerID: 841237,
	ShopID:    307271,
	Timestamp: int(time.Now().Unix()),
})

fmt.Println(shop.ShopName)
```

# Note

The following problems were discovered during the development of this API suite.

Some "required fields" are marked as "optional" .
Called namethe name of the field but type is float64orint .
Field obviously float64but will receive an empty string as a zero value (this part Shopeego been through the replacement string will ""be changed "0"in order to resolve a).
These problems you might encounter while using, when faced with API does not work properly or if there were error_param, consider the Issue in return.

# Sign

This kit still does this for you, but still wants you to know what happened. In response to the official request of Shrimp Skin, all requests must be signed with HMAC-SHA256. Obviously this is a somewhat magical security method (after all, HTTPS is already available). If you do not understand, here is an example.

```
POST /api/v1/orders/detail HTTP/1.1
Host: partner.shopeemobile.com
Authorization: b37c061daf2fcfa2baffe7539110938be5b7525041c147e78ad8afa78cc1a72d

{"ordersn": "160726152598865", "shopid": 61299, "partner_id": 1, "timestamp": 1470198856}
```

In the request header Authorizationfield, the hash code from "Request URL|Content" and through your API key as a secret (Secret) calculation is made.