# PB

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/go-petstore
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	pb "PB/v2"
	"PB/v2/pkg/models/operations"
	"PB/v2/pkg/models/shared"
	"context"
	"log"
)

func main() {
	s := pb.New(
		pb.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Animals.CreateAnimal(ctx, &operations.CreateAnimalRequestBody{
		Color: "white",
		ID:    "<ID>",
		Name:  "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Animals != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [.Animals](docs/sdks/animals/README.md)

* [CreateAnimal](docs/sdks/animals/README.md#createanimal) - create an animal
* [CreateLivingThings](docs/sdks/animals/README.md#createlivingthings) - create a living thing
* [DeleteAnimalsByID](docs/sdks/animals/README.md#deleteanimalsbyid) - Delete Animal Object
* [GetAllAnimals](docs/sdks/animals/README.md#getallanimals) - Your GET endpoint
* [GetAllLivingThings](docs/sdks/animals/README.md#getalllivingthings) - Get All living things
* [GetAnimalsByID](docs/sdks/animals/README.md#getanimalsbyid) - Get Animal
* [UpdateAnimalsByID](docs/sdks/animals/README.md#updateanimalsbyid) - Update Animal

### [.Birds](docs/sdks/birds/README.md)

* [CreateLivingThings](docs/sdks/birds/README.md#createlivingthings) - create a living thing
* [CreateNewBird](docs/sdks/birds/README.md#createnewbird) - Create new Bird
* [GetAllBirds](docs/sdks/birds/README.md#getallbirds) - Get Birds
* [GetAllLivingThings](docs/sdks/birds/README.md#getalllivingthings) - Get All living things
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


## Example

```go
package main

import (
	pb "PB/v2"
	"PB/v2/pkg/models/operations"
	"PB/v2/pkg/models/shared"
	"context"
	"log"
)

func main() {
	s := pb.New(
		pb.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Animals.CreateAnimal(ctx, &operations.CreateAnimalRequestBody{
		Color: "white",
		ID:    "<ID>",
		Name:  "string",
	})
	if err != nil {

		var e *sdkerrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.petstore.com` | None |
| 1 | `https://sandbox-api.petstore.com` | None |

For example:

```go
package main

import (
	pb "PB/v2"
	"PB/v2/pkg/models/operations"
	"PB/v2/pkg/models/shared"
	"context"
	"log"
)

func main() {
	s := pb.New(
		pb.WithServerIndex(1),
		pb.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Animals.CreateAnimal(ctx, &operations.CreateAnimalRequestBody{
		Color: "white",
		ID:    "<ID>",
		Name:  "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Animals != nil {
		// handle response
	}
}

```


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:

```go
package main

import (
	pb "PB/v2"
	"PB/v2/pkg/models/operations"
	"PB/v2/pkg/models/shared"
	"context"
	"log"
)

func main() {
	s := pb.New(
		pb.WithServerURL("https://api.petstore.com"),
		pb.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Animals.CreateAnimal(ctx, &operations.CreateAnimalRequestBody{
		Color: "white",
		ID:    "<ID>",
		Name:  "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Animals != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->



<!-- Start Authentication -->

# Authentication

## Per-Client Security Schemes

Your SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `Key1`       | oauth2       | OAuth2 token |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:

```go
package main

import (
	pb "PB/v2"
	"PB/v2/pkg/models/operations"
	"PB/v2/pkg/models/shared"
	"context"
	"log"
)

func main() {
	s := pb.New(
		pb.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Animals.CreateAnimal(ctx, &operations.CreateAnimalRequestBody{
		Color: "white",
		ID:    "<ID>",
		Name:  "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Animals != nil {
		// handle response
	}
}

```
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
