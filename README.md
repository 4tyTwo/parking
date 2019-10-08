# Parking project
Программная часть, проекта макета автоматизированной парковочной системы

## API
HTTP API - является единственной точкой доступа для управления системой.
На данный момент API предоставляет 3 публичных метода.  
В ответ на любой вызов, система возвращает http статус код и тело в формате JSON.
---

### GetFreePlaces
Параметр | Значение
---------|------------
`endpoint` | `/parkingLot`  
`method`   | `GET`

#### Пример
```
curl -i http://127.0.0.1:8080/parkingLot
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 15 Aug 2019 17:55:23 GMT
Content-Length: 17

{
    "FreePlaces":15
}
```

### PlaceCar

Параметр | Значение
---------|------------
`endpoint` | `/parkingLot`  
`method`   | `POST`

В случае успеха, в теле ответа будет содержаться пятизначный код, этот код необходимо указать при обращении за машиной.

#### Пример
```
curl -i --request POST --url http://127.0.0.1:8080/parkingLot \
    --header 'content-type: application/json' \
    --data '{}'
HTTP/1.1 202 Accepted
Content-Type: application/json; charset=utf-8
Date: Tue, 15 Aug 2019 17:58:47 GMT
Content-Length: 21

{"PlaceCode":"GZIEY"}
```

### GetCar
Параметр | Значение
---------|------------
`endpoint` | `/parkingLot/{code}`  
`method`   | `GET`
`code`     |  Пятизначный код, полученный при вызове метода PlaceCar

#### Пример
```
curl -i http://127.0.0.1:8080/parkingLot/GZIEY
HTTP/1.1 202 Accepted
Content-Type: application/json; charset=utf-8
Date: Tue, 15 Aug 2019 18:02:20 GMT
Content-Length: 2

{}
```
