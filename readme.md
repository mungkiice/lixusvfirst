# Lixus-VFirst Technical Test

Developed using Go programming language and uses database MongoDB. Developed for technical test purposes only.

## Function List

Functions that contained in this application.

### 1. SMS List Table View

Endpoint
```
http://{url}/ on browser
```

### 2. JWT Auth

Endpoint
```
http://{url}/login POST
```
Request Header
```
{
    "Content-Type": "application/json"
}
```
Request Body
```
{
    "username":"{username}",
    "password":"{pass}"
}
```
Response Body
```
{
  "status": "login succeed",
  "token": "{token}",
  "username": "{username}"
}
```

### 3. List Client's SMS

Endpoint
```
http://{url}/list GET
```

Request Header
```
{
    "Authorization":"Bearer {token}"
}
```

Response Body
```
{
    "SMS Status List": [
        {
            "To": "6281289594060",
            "From": "VFIRST",
            "Message": "Selamat anda menang STEAM WALLET 10 Dolla",
            "Time": null,
            "DeliveredDate": null,
            "ClientGUID": "",
            "ClientSeqNumber": "",
            "MessageID": "",
            "Circle": "",
            "Operator": "",
            "MSGStatus": "",
            "VendorStatus": "Sent.",
            "Client": "demoindo1"
        }
    ]
}
```

### 4. Push SMS

Endpoint
```
http://{url}/push POST
```

Request Header
```
{
    "Content-type":"application/json",
    "Authorization":"Bearer {token}"
}
```

Request Body
```
{
	"to": "6281289594060",
	"from": "VFIRST",
    "text": "Update dlr ulr"
}
```

Response Body
```
{
    "response": "Sent."
}
```
