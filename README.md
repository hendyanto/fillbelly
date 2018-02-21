# fillbelly

### Setup

run `vagrant up`

run `vagrant ssh`

run `go get ./...`

run `go build`

run `./fillbelly` 

### Testing (inside vagrant)
run `GO_ENV=test go test -v`

### Benchmark (inside vagrant)
run `GO_ENV=test go test -v -bench=.`

# API
## Nearby

Get restaurants 5 KM from location

**URL** : `/nearby?latitude=-6.115734&longitude=106.7916073`

**Method** : `GET`

**Auth required** : None

**Permissions required** : None

## Success Response

**Code** : `200 OK`

```json
[
  {
    "Id":1,
    "Name":"Puvlic 1",
    "Rating":9,
    "Id_category":2,
    "Longitude":106.7867153,
    "Latitude":-6.115894,
    "Category_name":"Seafood"
  },
  {
    "Id":2,
    "Name":"Puvlic 2",
    "Rating":9,
    "Id_category":3,
    "Longitude":106.7867143,
    "Latitude":-6.115594,
    "Category_name":"Seafood"
   }
]
```
## Reservation

Reserve a restaurant

**URL** : `/reserve`

**Method** : `POST`

**Params** : 
```
date: 2011-01-02T22:04:05Z
id_restaurant: 12
name: Batman
```

**Auth required** : None

**Permissions required** : None

## Success Response

**Code** : `200 OK`
