# Assessment Test REST API

> My REST API to implement a Book Library program.

## Testing golang REST API on my Heroku account
1. Visit url https://rest-api-my-demo.herokuapp.com . If you are testing the API directly then add the path that you wish to test via your client. If your client is Postman, then you can use and import my collection that I uploaded titled "REST-API-DEMO.postman_collection.json" to your Postman client.

## Testing golang REST API locally (option #1)
1. Two options for running the REST API locally;
   - Copy "_rest_api_for_resume.exe" from this repository to your local hard drive, or 
   - Download Golang environment locally and download my repository(title the local project folder "_rest_api_for_resume") locally. Then from command line run: go get && go build. local port is 8080.
2. locate the "_rest_api_for_resume.exe" and double click that executable and make sure this window opens and stays open while you use the client program. NOTE: you may need to allow this file/program in your OS's firewall to run it successfully, in Windows OS you may get a "Windows Security Alert" prompt,
   please make sure the public network checkbox is selected, then choose "Allow Access" button, then click "yes" if you are prompted to allow the program to 
   make changes to the computer.
3. Use http://localhost:8080 as the application address when consuming the REST API endpoints.


## Unit Testing
1. copy my code(all files) to your local hard drive.
2. make sure go lang environment is installed on your local machine
3. open your OS's command line prompt and change directory to the directory that you put my code
    files into. Then run the command "go test" to run my unit test. If all is setup up correctly 
    all my test will pass and you will see the word "PASS" on the command prompt screen. 



## Using Endpoints
- Please use the json formatting referenced in the examples below for your data entries. If you test the API
  
## Endpoints
- This API uses the seven main HTTP methods; GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS
- Also this API implements input validation in the API.

### Get Health
``` bash
# URL: api/health
# Method: HEAD
# HTTP/1.1

# Response sample
# code: 204
# content: n/a
```

### Create Book
``` bash
# URL: api/createbook
# Method: POST 
# HTTP/1.1

# Request sample
# {
#    "isbn": "ajdkljklajdljlsdlkfj",
#    "title": "Greatest Cook on Earth",
#    "author_name": "Stanley Lee",
# }

# Response sample
# Code: 201
# Content: {
#    "isbn": "ajdkljklajdljlsdlkfj",
#    "title": "Greatest Cook on Earth",
#    "author_name": "Stanley Lee",
#    "ID": 2,
#    "CreatedAt": "2022-08-07T21:19:14.8213596-04:00",
#    "UpdatedAt": "2022-08-07T21:19:14.8213596-04:00",
#    "DeletedAt": null
# }
```

### Get BookList
``` bash
# URL: api/getbooks 
# Method: READ 
# HTTP/1.1

# Response sample
# Code: 200
# [
#     {
#         "isbn": "ajdkljklajdljlsdlkfj",
#         "title": "Greatest Cook on Earth",
#         "author_name": "Tony Stark",
#         "ID": 1,
#         "CreatedAt": "2022-08-07T21:18:51.994608-04:00",
#         "UpdatedAt": "2022-08-07T21:18:51.994608-04:00",
#         "DeletedAt": null
#     },
#     {
#         "isbn": "ajdkljklajdljlsdlkfj",
#         "title": "Greatest Cook on Earth",
#         "author_name": "Stanley Lee",
#         "ID": 2,
#         "CreatedAt": "2022-08-07T21:19:14.821359-04:00",
#         "UpdatedAt": "2022-08-07T21:19:14.821359-04:00",
#         "DeletedAt": null
#     }
# ]
```

### Get Single Book
``` bash
# URL: api/getbook/:id
# Method: GET 
# HTTP/1.1

# Response sample
# Code: 200
# Content: {
#     "isbn": "ahuoeiojio",
#     "title": "The Incredible Hulk",
#     "author_name": "Silver Surfer",
#     "ID": 13,
#     "CreatedAt": "2022-08-11T12:05:49.019444-04:00",
#     "UpdatedAt": "2022-08-11T12:05:49.019444-04:00",
#     "DeletedAt": null
# }
```


### Update Whole Book
``` bash
# URL: api/updatebook/:id
# Method: PUT 
# HTTP/1.1

# Request sample
# {
#    "isbn": "ajdkljklajdljlsdlkfj", 
#    "title": "Greatest Cartoon Book",
#    "author_name": "Stanley Lee"
#  }

# Response sample
# Code: 200
# Content: {
#    "isbn": "ajdkljklajdljlsdlkfj",
#    "title": "Greatest Cook on Earth",
#    "author_name": "Stanley Lee",
#    "ID": 2,
#    "CreatedAt": "2022-08-07T21:19:14.8213596-04:00",
#    "UpdatedAt": "2022-08-07T21:19:14.8213596-04:00",
#    "DeletedAt": null
# }
```

### Update Part of Book
``` bash
# URL: api/updatebook/:id
# Method: PATCH 
# HTTP/1.1

# Request sample
# {
#   "title": "The Incredible Hulk",
#    "author_name": "Silver Surfer"
#  }

# Response sample
# Code: 200
# Content: {
#     "isbn": "ahuoeiojio",
#     "title": "The Incredible Hulk",
#     "author_name": "Silver Surfer",
#     "ID": 12,
#     "CreatedAt": "2022-08-11T11:50:55.931572-04:00",
#     "UpdatedAt": "2022-08-11T11:53:36.3829322-04:00",
#     "DeletedAt": null
# }
```

### Delete Book
``` bash
# URL: api/deletebook/:id
# Method: DELETE 
# HTTP/1.1

# Response sample
# Code: 200
# Content: {
#     "success": "deleted book id:12"
# }
```


## Errors and questions
if you run into any errors, please contact me as soon as possible by email, and I will try to help or rewrite my code.



```


### Writer

Nicholas Donald
