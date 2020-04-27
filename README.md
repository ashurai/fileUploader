# fileUploader
A Go / Golang demo application to upload files using bytes and FormFile with Mongo implementation. With Helper handler and model structure design

### Checkout on your local using clone or just ``` go get https://github.com/ashurai/fileUploader```

##### Before running application make sure on your local go V1.12.* or above is availabel and mongoDB setup is running with default configuration.

##### To start running application just hit ``` go run main.go``` from cli

### Avaialable Endpoints 

#### POST localhost:8090/upload
Using postman or brower post method with the help of curl / javascript 

#### GET localhost:8090/files/1
To list out all the files with local directory, and id's along with pagination, per page only 5 records are supported for now

