# Animals
(*Animals*)

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
	pb "PB"
	"PB/pkg/models/shared"
	"PB/pkg/models/operations"
)

func main() {
    s := pb.New(
        pb.WithSecurity(shared.Security{
            Key1: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Animals.CreateAnimal(ctx, operations.CreateAnimalRequestBody{
        Age: pb.Int64(943749),
        Color: "saepe",
        ID: "a7596eb1-0faa-4a23-92c5-955907aff1a3",
        Name: "Carlos Ziemann",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateAnimalRequestBody](../../models/operations/createanimalrequestbody.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


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
	pb "PB"
	"PB/pkg/models/shared"
)

func main() {
    s := pb.New(
        pb.WithSecurity(shared.Security{
            Key1: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Animals.CreateLivingThings(ctx, shared.ComplexObject{
        Data: &shared.ComplexObjectData{
            Animal: []shared.Animals{
                shared.Animals{
                    Age: pb.Int64(253291),
                    Color: pb.String("commodi"),
                    ID: pb.String("7739251a-a52c-43f5-ad01-9da1ffe78f09"),
                    Name: pb.String("Ms. Karla Aufderhar"),
                },
            },
            Birds: &shared.ComplexObjectDataBirds{
                Food: []string{
                    "maiores",
                },
                ID: "15471b5e-6e13-4b99-9488-e1e91e450ad2",
                Name: "Rudy Spencer",
            },
            CreatedDate: &shared.ComplexObjectDataCreatedDate{},
            UpdatedDate: &shared.ComplexObjectDataUpdatedDate{},
        },
        Meta: &shared.ComplexObjectMeta{},
        Name: pb.String("Beth McGlynn Sr."),
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
	pb "PB"
	"PB/pkg/models/shared"
	"PB/pkg/models/operations"
)

func main() {
    s := pb.New(
        pb.WithSecurity(shared.Security{
            Key1: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Animals.DeleteAnimalsByID(ctx, operations.DeleteAnimalsByIDRequest{
        ID: "d502a94b-b4f6-43c9-a9e9-a3efa77dfb14",
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
	pb "PB"
	"PB/pkg/models/shared"
	"PB/pkg/models/operations"
)

func main() {
    s := pb.New(
        pb.WithSecurity(shared.Security{
            Key1: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Animals.GetAllAnimals(ctx, operations.GetAllAnimalsRequest{
        Age: pb.String("cumque"),
        Color: pb.String("facere"),
        ID: pb.String("66ae395e-fb9b-4a88-b3a6-6997074ba446"),
        Name: pb.String("Robin Keebler"),
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
	pb "PB"
	"PB/pkg/models/shared"
	"PB/pkg/models/operations"
)

func main() {
    s := pb.New(
        pb.WithSecurity(shared.Security{
            Key1: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Animals.GetAllLivingThings(ctx, operations.GetAllLivingThingsRequest{
        Filter: []interface{}{
            "architecto",
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
	pb "PB"
	"PB/pkg/models/shared"
	"PB/pkg/models/operations"
)

func main() {
    s := pb.New(
        pb.WithSecurity(shared.Security{
            Key1: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Animals.GetAnimalsByID(ctx, operations.GetAnimalsByIDRequest{
        Animals: &shared.Animals{
            Age: pb.Int64(298282),
            Color: pb.String("et"),
            ID: pb.String("959890af-a563-4e25-96fe-4c8b711e5b7f"),
            Name: pb.String("Louis Turner Sr."),
        },
        ID: "8921cddc-6926-401f-b576-b0d5f0d30c5f",
        PerPage: pb.Int64(729991),
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetAnimalsByIDRequest](../../models/operations/getanimalsbyidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
	pb "PB"
	"PB/pkg/models/shared"
	"PB/pkg/models/operations"
)

func main() {
    s := pb.New(
        pb.WithSecurity(shared.Security{
            Key1: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Animals.UpdateAnimalsByID(ctx, operations.UpdateAnimalsByIDRequest{
        Animals: &shared.Animals{
            Age: pb.Int64(749999),
            Color: pb.String("dolores"),
            ID: pb.String("58705320-2c73-4d5f-a9b9-0c28909b3fe4"),
            Name: pb.String("Omar Leuschke"),
        },
        ID: "cbf48633-323f-49b7-bf3a-4100674ebf69",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateAnimalsByIDRequest](../../models/operations/updateanimalsbyidrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateAnimalsByIDResponse](../../models/operations/updateanimalsbyidresponse.md), error**

