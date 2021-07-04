# Efishery-Task #
This repo is the submission for efishery backend engineer skill test
# Table of Contents
* [Auth-App](#auth-app)
	* [Setup](#auth-app-setup)
	* [Build & Run](#auth-app-build-and-run)
	* [Summary](#auth-app-summary)
	* [Examples](#auth-app-examples)
* [Fetch-App](#fetch-app)
	* [Setup](#fetch-app-setup)
	* [Build & Run](#fetch-app-build-and-run)
	* [Summary](#fetch-app-summary)
	* [Examples](#fetch-app-examples)

## Auth-App
`auth-app` will manage new user, password, and JWT generation process using a file based database.

### Auth-App Setup
To set up the application, you need to have `docker` installed in your machine or `go1.16` if you want to run the app directly in your local machine. If you already have them, set these values in `auth-app/.env` for the app configuration.
* `PORT` to define which port that the application will listen to 
	* The default port is `8080`
* `FILE_PATH` to define the path to where the file to store user data
* `SECRET` to define which secret to use to generate JWT

### Auth-App Build and Run
There are 2 ways to run the app after the configuration values for the app is set up in `auth-app/.env`.
First method is to run the application directly in local machine by using :
```
$ cd auth-app
$ go run main.go
```
The other method is to build this app docker image and run a container based on the image. To build the image, execute these in your machine :
```
$ cd auth-app
$ docker build -t <image-name> .
```
then to run the container based on the image, execute these after the image is built (make sure that you set the docker port to `PORT` value in `auth-app/.env`) :
```
$ docker run -d -p <host port>:<docker port> <image-name>
```
### Auth-App Summary
Auth-App has 4 endpoints that can be accessed. 
* GET `/hello` 
This endpoint will return a `"Hello World!"` string and `Status Code` `200`. This endpoint can be used to verify the app is running.
* POST `/user`
This endpoint will insert user data that is received from the request body into the database and will return the 4 characters password for that user. This endpoint requires request body, failed to fulfill this requirement will make the app return an error.
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
This endpoint will receive a `phone` and `password` and return a generated JWT using the `SECRET` which is set in `auth-app/.env` with `Private Claims` that contains `name`, `phone`, `role`, and `timestamp` of the user that has the correct/matching `phone` and `password`.
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
This endpoint will check the `Authorization` header of the request, verify the JWT inside the header, and return the `Private Claims` of the JWT.
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


### Auth-App Example
```
$ cd auth-app
$ docker build -t auth-app .
$ docker run -d -p 8080:8080 auth-app
$ curl localhost:8080/hello
```
