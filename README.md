## Cowin Alerts

Cowin Alerts

## Build

```
go get -d ./...
go build main/main.go
```

```
make <build_type>

for example
make osx_build


```

## Instruction to Execute

``./bin/main-osx 1 03-05-2021 45 560078``

``./bin/main-osx 2 03-05-2021 45 265``

* first Argument - 1- search by pincode 2 seach by district id
* second argument - - date dd-mm-yyyy
* minimum age limit - 18/45
* Last argument - pincode/district id 

* For State ID - https://cdn-api.co-vin.in/api/v2/admin/location/states
* For Distric ID - https://cdn-api.co-vin.in/api/v2/admin/location/districts/<stateid>  