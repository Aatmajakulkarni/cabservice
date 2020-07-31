# cabservice

A cab service subset is implemented in Golang

A backend RESTful API available -
    /v1/new/login - new user login (POST)
    /v1/user/book - books nearest available cab(POST)
    /v1/user/rides - lists all user's past rides (GET)
    /v1/cab//endride/:id - ends ride for given cab id (PUT)
 
Gin framework is used for REST apis.
Models folder is for declaring structure.
Mysql is used for database.

The API service can be run by 'go run main.go' in the root folder. It will run in default 9000 port of the localhost. Can use CURL for testing.

For more info, contact kulkarni.aatmaja@gmail.com
