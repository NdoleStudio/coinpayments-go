# coinpayments-go

[![Build](https://github.com/NdoleStudio/coinpayments-go/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/coinpayments-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/coinpayments-go/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/coinpayments-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/coinpayments-go/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/coinpayments-go/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/coinpayments-go)](https://goreportcard.com/report/github.com/NdoleStudio/coinpayments-go)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/coinpayments-go)](https://github.com/NdoleStudio/coinpayments-go/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/coinpayments-go?color=brightgreen)](https://github.com/NdoleStudio/coinpayments-go/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/coinpayments-go)](https://pkg.go.dev/github.com/NdoleStudio/coinpayments-go)


This package provides a generic `go` client template for the CoinPayments HTTP API

## Installation

`coinpayments-go` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/coinpayments-go
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/coinpayments-go"
```


## Implemented

- [Payments](#payments)
    - `create_transaction`: Create Transaction

## Usage

### Initializing the Client

An instance of the client can be created using `New()`.

```go
package main

import (
	"github.com/NdoleStudio/coinpayments-go"
)

func main()  {
    client := coinpayments.New(
        coinpayments.WithAPIKey(/* API Key */),
        coinpayments.WithAPISecret(/* API Secret */),
    )
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
status, response, err := client.Payments.CreatePayment(context.Background())
if err != nil {
    //handle error
}
```

### Payments

#### `create_transaction`: Create Transaction

```go
transaction, response, err := client.Payments.CreatePayment(context.Background(), &CreatePaymentRequest{
        Amount:           "1.00000000",
        OriginalCurrency: "USD",
        SendingCurrency:  "USD",
        BuyerEmail:       "john@example.com",
    }
)

if err != nil {
    log.Fatal(err)
}

log.Println(transaction.Error) // ok
```

## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
