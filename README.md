# Efishery-Task #
This repo is the submission for efishery backend engineer skill test
# Table of Contents
* [Auth-App](#auth-app)
	* [Setup](#auth-app-setup)
	* [Build & Run](#auth-app-build-and-run)
	* [Summary](#auth-app-summary)
    * [System Design](#auth-app-system-design)
	* [Examples](#auth-app-examples)
* [Fetch-App](#fetch-app)
	* [Setup](#fetch-app-setup)
	* [Build & Run](#fetch-app-build-and-run)
	* [Summary](#fetch-app-summary)
    * [System Design](#fetch-app-system-design)
	* [Examples](#fetch-app-examples)

## Auth-App
`auth-app` will manage new user, password, and JWT generation process using a file based database.

### Auth-App Setup
To set up the application, you need to have `docker` installed in your machine or `go1.16` if you want to run the app directly in your local machine. If you already have them, set these values in `auth-app/.env` for the app configuration.
* `SERVER_PORT` to define which port that the application will listen to 
	* The default port is `8080`
* `FILE_PATH` to define the path to where the file to store user data
* `SECRET` to define which secret to use to generate/parse JWT

### Auth-App Build and Run
There are 2 ways to run the app after the configuration values for the app is set up in `auth-app/.env`.
First method is to run the application directly in local machine by using :
```
$ cd auth-app
$ go mod tidy
// run 'go get xxx' if needed to pull the needed library/package
$ go run main.go
```
The other method is to build this app docker image and run a container based on the image. To build the image, execute these in your machine :
```
$ cd auth-app
$ docker build -t <image-name> .
```
then to run the container based on the image, execute these after the image is built (make sure that you set the docker port to `SERVER_PORT` value in `auth-app/.env`) :
```
$ docker run -d -p <host port>:<docker port> <image-name>
```
`-d` means the docker will run detached/in background and `-p` means the docker will publish and listen to those port
### Auth-App Summary
Auth-App has 4 endpoints that can be accessed. 
* GET `/hello` 
-- This endpoint will return a `"Hello World!"` string and `Status Code` `200`. This endpoint can be used to verify the app is running.
* POST `/user`
-- This endpoint will insert user data that is received from the request body into the database and will return the 4 characters password for that user. This endpoint requires request body, failed to fulfill this requirement will make the app return an error.
```
//HTTP Request Body (JSON)
{
	"name": "xxx",
	"phone": "000",
	"role": "admin"
}

//HTTP Response (Application/JSON)
{
	"status_code": 201,
	"data": "<4 characters string>"
}
```
* GET `/token`
-- This endpoint will receive a `phone` and `password` and return a generated JWT using the `SECRET` which is set in `auth-app/.env` with `Private Claims` that contains `name`, `phone`, `role`, and `timestamp` of the user that has the correct/matching `phone` and `password`.
```
//HTTP Request Body (JSON)
{
	"phone": "000",
	"password": "<4 characters string>"
}

//HTTP Response (Application/JSON)
{
	"status_code": 200,
	"data": "<JWT>"
}
```
* GET `/claims`
-- This endpoint will check the `Authorization` header of the request, verify the JWT inside the header, and return the `Private Claims` of the JWT.
```
//HTTP Request Header (Bearer Token)
Authorization: Bearer <JWT>

//HTTP Response (Application/JSON)
{
	"status_code": 200,
	"data": {
		"name": "xxx",
		"phone": "000",
		"role": "admin",
		"timestamp": "04 Jul 21 09:07 UTC"
	}
}
```

### Auth-App System Design
``` 
Using kataras/iris Golang Web Framework

              HTTP Response
                    A
                    |
                    |
HTTP Request ---> Handler <---> Logic <---> Repo <---> Database
              1            2          3         4
            (Model)     (Model)    (Model)    (Model)

1. HTTP Request is parsed and assigned into a context variable which will be handled by Handler functions.
    a. Handler will parse the request body and header and assign those values to variables/struct which is taken from the 'model' package that is shared through all levels/layers of functions.
    b. Handler will also return an HTTP Response after processing the received/parsed data. Whether it's an error/failed or successful process.
    c. Handler will only handle request parsing, response, auth, and user/client communication related task.
2. Handler will calls Logic functions that will do data processing with the parsed values from request in variables/struct from 'model' package as the functions parameters.
    a. Logic functions will only do data processing. All the data needed for the process are either given by Handler as parameters or fetched by Repo functions from Database
    b. Logic functions can also be a private functions that can only be called internally in the package. These functions function to help processing the data and will not be called nor return values to Handler.
3. Logic will calls Repo functions to fetch more data that is needed for the process
4. Repo will fetch data from all external sources, in this case a database. It will fetch the data and parse them and return them before returning the parsed data to Logic.
    a. Since Auth-App is using a csv file for the database, most of Repo functions are built to read and write a csv formatted file.                       
```
### Auth-App Example
```
$ cd auth-app
$ docker build -t auth-app .
$ docker run -d -p 8080:8080 auth-app
$ curl localhost:8080/hello
```
## Fetch-App
`fetch-app` will fetch and process data/resources
### Fetch-App Setup
To set up the application, you need to have `docker` installed in your machine or `go1.16` if you want to run the app directly in your local machine. If you already have them, set these values in `fetch-app/.env` for the app configuration.
* `SERVER_PORT` to define which port that the application will listen to 
	* The default port is `3000`
* `SECRET` to define which secret to use to generate/parse JWT
* `KEY` to be used when getting currency conversion rate from https://free.currencyconverterapi.com.
    * You need to register and verify your email first on the free plan to get the API Key
* `CACHE_DURATION` to define how long before the cache need to be updated
    * The number that is set will be counted in `minute` (ie. `CACHE_DURATION=45` means the cache will be updated every 45 minutes after the app is started)
    * Based on the documentation on https://free.currencyconverterapi.com, the currency conversion rate is updated at 60 minutes interval, therefore the default value for this configuration is `60`


### Fetch-App Build and Run
There are 2 ways to run the app after the configuration values for the app is set up in `fetch-app/.env`.
First method is to run the application directly in local machine by using :
```
$ cd fetch-app
$ go mod tidy
// run 'go get xxx' if needed to pull the needed library/package
$ go run main.go
```
The other method is to build this app docker image and run a container based on the image. To build the image, execute these in your machine :
```
$ cd fetch-app
$ docker build -t <image-name> .
```
then to run the container based on the image, execute these after the image is built (make sure that you set the docker port to `SERVER_PORT` value in `fetch-app/.env`) :
```
$ docker run -d -p <host port>:<docker port> <image-name>
```

### Fetch-App Summary
Fetch-App has 4 endpoints that can be accessed. 
* GET `/hello` 
-- This endpoint will return a `"Hello World!"` string and `Status Code` `200`. This endpoint can be used to verify the app is running.
```
All endpoints below have middlewares that will verify the JWT secret and/or the role inside JWT is 'admin'. Therefore these endpoints require a valid JWT to work correctly.

//HTTP Request Header (Bearer Token)
Authorization: Bearer <JWT>
```
* GET `/claims`
-- This endpoint will check the `Authorization` header of the request, verify the JWT inside the header, and return the `Private Claims` of the JWT.
```
//HTTP Request Header (Bearer Token)
Authorization: Bearer <JWT>

//HTTP Response (Application/JSON)
{
	"status_code": 200,
	"data": {
		"name": "xxx",
		"phone": "000",
		"role": "admin",
		"timestamp": "04 Jul 21 09:07 UTC"
	}
}
```
* GET `/komoditas`
-- This endpoint will fetch commodities data from https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list, clean the data from `nil`, `null`, and/or empty values, then returns the cleaned data with additional field of the commodity price in USD currency
```
//HTTP Request Header (Bearer Token)
Authorization: Bearer <JWT>

//HTTP Response (Application/JSON)
{
	"status_code": 200,
    "data": [
        {
            "uuid": "a9e10e7d-d385-49b5-ac55-fe6d04f300af",
            "komoditas": "Patin Albino",
            "area_provinsi": "JAWA TENGAH",
            "area_kota": "PURWOREJO",
            "size": "120",
            "price": "20000",
            "tgl_parsed": "2020-06-08T12:14:31.719Z",
            "timestamp": "1591593271719",
            "price_usd": "$1.383169"
        },
        {
            "uuid": "ff8e461c-910b-4d11-b531-8494c1c7da2e",
            "komoditas": "Gabus",
            "area_provinsi": "JAWA BARAT",
            "area_kota": "BANDUNG",
            "size": "70",
            "price": "30000",
            "tgl_parsed": "2020-06-28T14:53:22+0800",
            "timestamp": "1593327202",
            "price_usd": "$2.074753"
        }
        ...
    ]
}
```
* GET `/komoditas/compiled`
-- This endpoint will fetch commodities data from https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list, clean the data from `nil`, `null`, and/or empty values, then returns the aggregated data by the `area_provinsi` value, weekly, and returns the max, min, avg, and median profit (assuming profit is `price` * `size`). This endpoint requires 'admin' role inside the valid JWT in the header of the request.
```
//HTTP Request Header (Bearer Token)
Authorization: Bearer <JWT>

//HTTP Response (Application/JSON)
{
	"status_code": 200,
    "data": [
        {
            "area_provinsi": "BALI",
            "Profit": {
                "Tahun 2020": {
                    "Minggu ke 44": 109174680,
                    "Minggu ke 45": 720000
                },
                "Tahun 2021": {
                    "Minggu ke 10": 2250000,
                    "Minggu ke 14": 79999920,
                    "Minggu ke 16": 6000000,
                    "Minggu ke 23": 1175000,
                    "Minggu ke 24": 5000000,
                    "Minggu ke 25": 400000000,
                    "Minggu ke 3": 11100000,
                    "Minggu ke 8": 1000000
                }
            },
            "max_profit": 400000000,
            "min_profit": 720000,
            "average_profit": 61641960,
            "median_profit": 11100000
        },
        {
            "area_provinsi": "JAWA SELATAN",
            "Profit": {
                "Tahun 2021": {
                    "Minggu ke 14": 800000
                }
            },
            "max_profit": 800000,
            "min_profit": 800000,
            "average_profit": 800000,
            "median_profit": 800000
        },
        {
            "area_provinsi": "SUMATERA UTARA",
            "Profit": {
                "Tahun 2020": {
                    "Minggu ke 41": 3000000,
                    "Minggu ke 42": 90000000
                }
            },
            "max_profit": 90000000,
            "min_profit": 3000000,
            "average_profit": 46500000,
            "median_profit": 3000000
        },
        {
            "area_provinsi": "SULAWESI BARAT",
            "Profit": {
                "Tahun 2020": {
                    "Minggu ke 42": 1200000,
                    "Minggu ke 45": 2000000
                },
                "Tahun 2021": {
                    "Minggu ke 26": 50000000
                }
            },
            "max_profit": 50000000,
            "min_profit": 1200000,
            "average_profit": 17733333,
            "median_profit": 1200000
        }
        ...
    ]
}
```
### Fetch-App System Design
``` 
Using Gin Golang Web Framework
Using memory/RAM based cache
Using goroutines based worker

                        HTTP Response                 Cache
                              A                         A
                              |                         |
                              |                         V
HTTP Request ---> Middleware ---> Handler <---> Logic <---> Repo <---> External Data Sources
              1               2             3           4          5
            (Model)        (Model)       (Model)     (Model)    (Model)

1. HTTP Request will be validated and authorized by the Middleware before be forwarded to Handler. Especially the JWT inside the header request.
    a. Failure on validating or authorizing HTTP Request will make Middleware to return an HTTP Response to user/client instead of Handler.
2. HTTP Request is parsed and assigned into a context variable which will be handled by Handler functions.
    a. Handler will parse the request body and header and assign those values to variables/struct which is taken from the 'model' package that is shared through all levels/layers of functions.
    b. Handler will also return an HTTP Response after processing the received/parsed data. Whether it's an error/failed or successful process.
    c. Handler will only handle request parsing, response, auth, and user/client communication related task.
3. Handler will calls Logic functions that will do data processing with the parsed values from request in variables/struct from 'model' package as the functions parameters.
    a. Logic functions will only do data processing. All the data needed for the process are either given by Handler as parameters or fetched by Repo functions from Database.
    b. Logic functions can also be a private functions that can only be called internally in the package. These functions function to help processing the data and will not be called nor return values to Handler.
4. Logic will calls Repo functions to fetch more data that is needed for the process.
    a. Logic will try to get the data from the cache first before calling Repo functions.
    b. There is a worker that will call Repo functions to get all the needed data and store it to the cache at some interval.
5. Repo will fetch data from all external sources, in this case an external HTTP server that hosts the data. It will fetch the data and parse them and return them before returning the parsed data to Logic.
    a. Fetch-App is getting data from an external HTTP server. Therefore, most of the Repo functions are built to send HTTP requests to the server and parse the response for the data before returning the data to the function that calls it.                       
```
### Fetch-App Example
```
$ cd fetch-app
$ docker build -t fetch-app .
$ docker run -d -p 3000:3000 fetch-app
$ curl localhost:3000/hello
```
`-d` means the docker will run detached/in background and `-p` means the docker will publish and listen to those port
