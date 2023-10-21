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
	pb "PB"
	"PB/pkg/models/operations"
	"PB/pkg/models/shared"
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


### [Animals](docs/sdks/animals/README.md)

* [CreateAnimal](docs/sdks/animals/README.md#createanimal) - create an animal
* [CreateLivingThings](docs/sdks/animals/README.md#createlivingthings) - create a living thing
* [DeleteAnimalsByID](docs/sdks/animals/README.md#deleteanimalsbyid) - Delete Animal Object
* [GetAllAnimals](docs/sdks/animals/README.md#getallanimals) - Your GET endpoint
* [GetAllLivingThings](docs/sdks/animals/README.md#getalllivingthings) - Get All living things
* [GetAnimalsByID](docs/sdks/animals/README.md#getanimalsbyid) - Get Animal
* [UpdateAnimalsByID](docs/sdks/animals/README.md#updateanimalsbyid) - Update Animal

### [Birds](docs/sdks/birds/README.md)

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

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
