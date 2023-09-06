# Animals

## Overview

Work with Animals.

### Available Operations

* [CreateAnimal](#createanimal) - create an animal
* [CreateLivingThings](#createlivingthings) - create a living thing
* [DeleteAnimalsByID](#deleteanimalsbyid) - Delete Animal Object
* [GetAllAnimals](#getallanimals) - Your GET endpoint
* [GetAllLivingThings](#getalllivingthings) - Get All living things
* [GetAnimalsByID](#getanimalsbyid) - Get Animal
* [UpdateAnimalsByID](#updateanimalsbyid) - Update Animal

## CreateAnimal

Post animals description

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PB"
	"PB/pkg/models/operations"
)

func main() {
    s := pb.New()
    operationSecurity := operations.CreateAnimalSecurity{
            Key1: "",
        }

    ctx := context.Background()
    res, err := s.Animals.CreateAnimal(ctx, operations.CreateAnimalRequestBody{
        Age: pb.Int64(870013),
        Color: "at",
        ID: "f7cc78ca-1ba9-428f-8816-742cb7392059",
        Name: "Sheryl Fadel",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Animals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateAnimalRequestBody](../../models/operations/createanimalrequestbody.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.CreateAnimalSecurity](../../models/operations/createanimalsecurity.md)       | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.CreateAnimalResponse](../../models/operations/createanimalresponse.md), error**


## CreateLivingThings

Create a living thing

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PB"
	"PB/pkg/models/shared"
)

func main() {
    s := pb.New()

    ctx := context.Background()
    res, err := s.Animals.CreateLivingThings(ctx, shared.ComplexObject{
        Data: &shared.ComplexObjectData{
            Animal: []shared.Animals{
                shared.Animals{
                    Age: pb.Int64(943749),
                    Color: pb.String("saepe"),
                    ID: pb.String("a7596eb1-0faa-4a23-92c5-955907aff1a3"),
                    Name: pb.String("Carlos Ziemann"),
                },
            },
            Birds: &shared.ComplexObjectDataBirds{
                Food: []string{
                    "numquam",
                },
                ID: "67739251-aa52-4c3f-9ad0-19da1ffe78f0",
                Name: "Mr. Jared Ritchie",
            },
            CreatedDate: &shared.ComplexObjectDataCreatedDate{},
            UpdatedDate: pb.Int64(979587),
        },
        Meta: &shared.ComplexObjectMeta{},
        Name: pb.String("Stacy Gulgowski MD"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ComplexObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.ComplexObject](../../models/shared/complexobject.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.CreateLivingThingsResponse](../../models/operations/createlivingthingsresponse.md), error**


## DeleteAnimalsByID

Delete the animal

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PB"
	"PB/pkg/models/operations"
)

func main() {
    s := pb.New()

    ctx := context.Background()
    res, err := s.Animals.DeleteAnimalsByID(ctx, operations.DeleteAnimalsByIDRequest{
        ID: "5e6e13b9-9d48-48e1-a91e-450ad2abd442",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteAnimalsByIDRequest](../../models/operations/deleteanimalsbyidrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteAnimalsByIDResponse](../../models/operations/deleteanimalsbyidresponse.md), error**


## GetAllAnimals

Get Animals Description

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PB"
	"PB/pkg/models/operations"
)

func main() {
    s := pb.New()

    ctx := context.Background()
    res, err := s.Animals.GetAllAnimals(ctx, operations.GetAllAnimalsRequest{
        Age: pb.String("aliquid"),
        Color: pb.String("cupiditate"),
        ID: pb.String("802d502a-94bb-44f6-bc96-9e9a3efa77df"),
        Name: pb.String("Keith Gulgowski"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Animals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetAllAnimalsRequest](../../models/operations/getallanimalsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetAllAnimalsResponse](../../models/operations/getallanimalsresponse.md), error**


## GetAllLivingThings

get All living things data

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PB"
	"PB/pkg/models/operations"
)

func main() {
    s := pb.New()

    ctx := context.Background()
    res, err := s.Animals.GetAllLivingThings(ctx, operations.GetAllLivingThingsRequest{
        Filter: []interface{}{
            "ea",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllLivingThings200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetAllLivingThingsRequest](../../models/operations/getalllivingthingsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetAllLivingThingsResponse](../../models/operations/getalllivingthingsresponse.md), error**


## GetAnimalsByID

Get an animal

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PB"
	"PB/pkg/models/operations"
	"PB/pkg/models/shared"
)

func main() {
    s := pb.New()
    operationSecurity := operations.GetAnimalsByIDSecurity{
            Key1: "",
        }

    ctx := context.Background()
    res, err := s.Animals.GetAnimalsByID(ctx, operations.GetAnimalsByIDRequest{
        Animals: &shared.Animals{
            Age: pb.Int64(396506),
            Color: pb.String("laborum"),
            ID: pb.String("e395efb9-ba88-4f3a-a699-7074ba4469b6"),
            Name: pb.String("Brandon Brakus V"),
        },
        ID: "59890afa-563e-4251-afe4-c8b711e5b7fd",
        PerPage: pb.Int64(149448),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Animals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetAnimalsByIDRequest](../../models/operations/getanimalsbyidrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetAnimalsByIDSecurity](../../models/operations/getanimalsbyidsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.GetAnimalsByIDResponse](../../models/operations/getanimalsbyidresponse.md), error**


## UpdateAnimalsByID

Update the animal object

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PB"
	"PB/pkg/models/operations"
	"PB/pkg/models/shared"
)

func main() {
    s := pb.New()
    operationSecurity := operations.UpdateAnimalsByIDSecurity{
            Key1: "",
        }

    ctx := context.Background()
    res, err := s.Animals.UpdateAnimalsByID(ctx, operations.UpdateAnimalsByIDRequest{
        Animals: &shared.Animals{
            Age: pb.Int64(904648),
            Color: pb.String("pariatur"),
            ID: pb.String("028921cd-dc69-4260-9fb5-76b0d5f0d30c"),
            Name: pb.String("Mindy Renner"),
        },
        ID: "58705320-2c73-4d5f-a9b9-0c28909b3fe4",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Animals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateAnimalsByIDRequest](../../models/operations/updateanimalsbyidrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.UpdateAnimalsByIDSecurity](../../models/operations/updateanimalsbyidsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.UpdateAnimalsByIDResponse](../../models/operations/updateanimalsbyidresponse.md), error**

