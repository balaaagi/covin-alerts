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

* If you want to search by pincode ``./bin/main-osx 1 03-05-2021 45 560078 true 5m``

* If you want to search by district id ``./bin/main-osx 2 03-05-2021 45 265 false``

* first Argument - 1- search by pincode 2 seach by district id
* second argument - - date dd-mm-yyyy
* minimum age limit - 18/45
* Last argument - pincode/district id 

* For State ID - https://cdn-api.co-vin.in/api/v2/admin/location/states
* For Distric ID - https://cdn-api.co-vin.in/api/v2/admin/location/districts/`stateid`  

## Output
You would get console output as below
```
Center Name Manipal Clinic
No Slots Available for given date | pincode | district !
No Slots Available for given date | pincode | district !
--------------------
Center Name Jarganahalli Corporator Office
Date 03-05-2021
Slot Information [09:00AM-11:00AM 11:00AM-01:00PM 01:00PM-03:00PM 03:00PM-06:00PM]
Count Available: 6 Minimum Age: 45
--------------------
```

Also you will get OS level alerts