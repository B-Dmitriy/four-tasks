## Текническое задание GOLANG

### 1. Работа с HTTP
Тип программы: Web-server
Порт: 3010
Маршрутизация:

**POST /redirection**  
> Выполняет запрос на указанный адрес и получает json массив вида {1:"one",3:"two",2:"four"}
> Вернуть ответ со статусом 200 (StatusOk) c этим json в отсортированном и обычном виде
```
curl -i -X POST -d '{"4":"three","1":"one","3":"two","2":"four"}' http://localhost:3010/redirection
//--------------
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 08 May 2024 16:26:34 GMT
Content-Length: 112
{ 
    "origin": {"4":"three","1":"one","3":"two","2":"four"}, 
    "sort": {"1":"one","2":"four","3":"two","4":"three"} 
}
```

**GET /get**  
> если в параметрах передается token вернуть ответ со статусом 200, в противном случае вернуть статус 400 (BadRequest)
```
curl -i -X GET "http://localhost:3010/get?first=1&token=123&test=last"
//--------------
HTTP/1.1 200 OK
Date: Fri, 10 May 2024 14:12:21 GMT
Content-Length: 0

curl -i -X GET "http://localhost:3010/get?first=1&test=last"
//--------------
HTTP/1.1 400 Bad Request
Date: Fri, 10 May 2024 14:12:03 GMT
Content-Length: 26
Content-Type: text/plain; charset=utf-8

need token in query params
```
---
### 2. Асинхронное программирование
Тип программы: скрипт

> В цикле от 1 до 10 выводить номер (i) в отдельной горутине
> Должны вывестись все значения.

Тип программы: скрипт

> Существует 2 канала, в бесконечном цикле каждые 0.5 сек в первый канал заносится значение счетчика (Считает количество итераций). Если число четное его записать во второй канал.
> При получения значения во втором канале значение выводится в консоль.
> Через 4 секунды корректо завершить работу всех горутин.


---
### 3. SQL or NoSQL
Тип программы: Web-server
Порт: 3020
Маршрутизация:

> /get GET - выполнить запрос к базе данных и вывести все полученные значения в формате json cо статусом 200 (StatusOk)
```
curl -i -X GET http://localhost:3020/get
//--------------
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 14 May 2024 01:55:21 GMT
Content-Length: 103

[{"id":1,"name":"Alexandr","age":30},{"id":2,"name":"Egor","age":25},{"id":3,"name":"Julia","age":25}]
```

---
### 4. Профилирование

> Отпрофилировать один из кодов по оперативной памяти (куче) и CPU. Полученный результат представить в виде svg файла и прокомментировать