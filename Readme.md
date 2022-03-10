# Multipass golang

A Golang implementation of shopify multipass.

> Shopify provides a mechanism for single sign-on known as Multipass. Multipass uses an AES encrypted JSON hash and multipassify provides functions for generating tokens

More details can be found [here](http://docs.shopify.com/api/tutorials/multipass-login).

## Installation

<pre>
    go get github.com/soub4i/multipass
</pre>

## Required

- Secret: can be generated like [this](https://shopify.dev/api/multipass#requirements)

## Example

```go

package main

import (
	"fmt"
	"github.com/soub4i/multipass"
)


type ShopifyPayload struct {
    Email      string `json:"email"`
    RemoteIP  string `json:"remote_ip"`
}

func main() {

    payload := &ShopifyPayload{
        Email:      "cool-user@gmail.com",
        RemoteIP:  "84.811.238.81",
    }
    secret := "multipass secret from shop admin"

    var t,e = GenerateToken(secret, payload)

    if e != nil {
        fmt.Println("token generate failed: %v", e)
        return
    }
    fmt.Println(t) // 3ZuH62FJRcsyiWc7dVcQMJvOakNvxxfwieyRV6DTJNg=

    fmt.Println(GenerateURL(t, "soubai-merch"))
    // http://soubai-merch.myshopify.com/account/login/multipass/3ZuH62FJRcsyiWc7dVcQMJvOakNvxxfwieyRV6DTJNg=
}
```
