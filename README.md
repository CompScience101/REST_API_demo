# Assessment Test REST API

> My REST API to implement Book Library program.

## Running golang REST API without go docker container (option #1)
1. Copy "assessment_test_5_9_2022.exe" to local hard drive.
2. locate the "assessment_test_5_9_2022.exe" and double click that executable and make sure this window opens and stays open while you use the client program
   file, NOTE: you may need to allow this file/program in your OS's firewall to run it successfully, in Windows OS you may get a "Windows Security Alert" prompt,
   please make sure the public network checkbox is selected, then choose "Allow Access" button, then click "yes" if you are prompted to allow the program to 
   make changes to the computer.


## Unit Testing
1. copy my code(all files) to your local hard drive.
2. make sure go lang environment is installed on your local machine
3. open your OS's command line prompt and change directory to the directory that you put my code
    files into. Then run the command "go test" to run my unit test. If all is setup up correctly 
    all my test will pass and you will see the word "PASS" on the command prompt screen. 



## Using Endpoints
- When using the "Create transaction" endpoint, please make sure to provide a integer value for the "cost" field of your json.
  Use http://localhost:8080/api/transaction to communicate from the client to the api.
- When using the "Get Health" endpoint you don't need to pass any data, just use the url and path specified and if it succeeds
  you will get a 200 status response along with content body of the running status, otherwise you will get no response from 
  server if it is not available. Use http://localhost:8080/api/health to communicate from the client to the api.

## Endpoints

### Get Health
``` bash
# URL: api/health
# Method: HEAD
# HTTP/1.1
# Response sample
# code: 200
# content: n/a
```

### Create Transaction
``` bash
# URL: api/transaction 
# Method:POST 
# HTTP/1.1

# Request sample
# {
#   "cost": 120,
# }

# Response sample
# Code: 201
# Content: {
#   reward: 90
# }
```

## Errors and questions
if you run into any errors, please contact me as soon as possible by email, and I will try to help or rewrite my code.



```


### Writer

Nicholas Donald
