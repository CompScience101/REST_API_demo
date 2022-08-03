package main

import (
  "net/http"
  "fmt"
  "log"
  "os"
  "strings"
  //"text/template"
  "github.com/gin-gonic/gin"
  cors "github.com/rs/cors/wrapper/gin"
  "gorm.io/driver/postgres"
  //"github.com/go-gorm/gorm"
  "gorm.io/gorm"
)

// Book struct (Model)
type Book struct {
	ID     string  `json:"id" binding: "required"`
	Isbn   string  `json:"isbn binding: "required""`
	Title  string  `json:"title" binding: "required"`
	Author string `json:"author_fname" binding: "required"`
	gorm.Model
	//DateCreated string `json:"date_created"`
}

// Init books var as a slice Book struct
var books []Book

func health (c *gin.Context) { // monitor health of api
    c.JSON(http.StatusOK, nil)  
	fmt.Printf("replied to client health check with success status.")
}
func test (c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"test": "testing connection", "test2": "redeploy"})
}
/*
func createbook (){

}
func getbooks (){

}
func getbook (){

}

func updatebook() {}

func deletebook {}

func updatebookpatch {}

func transaction (c *gin.Context){ // process customer transaction
	var myjson MyTransaction 
	if c.BindJSON(&myjson) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment, please make sure you use a integer with no decimal places."})
		fmt.Printf("invalid payment, user needs to instead use a integer with no decimal places.")
		return
	}
	fmt.Printf("sending cost to calculate function %d", myjson.Cost )
	points := calculateTransaction(myjson.Cost)
	c.JSON(http.StatusCreated, gin.H{"reward": points})
	fmt.Printf("sent reward client %d  points \n", points )
	return
	//c.Data(200, "application/json; charset=utf-8", []byte(string(dat)))	
}

func calculateTransaction (cost int64)(int64) { // calculate customer reward points
    fmt.Printf("cost before conversion %d \n", cost )
    //convertedCost, err := strconv.Atoi(cost)
	//fmt.Printf("convertedCost after conversion %d  \n", convertedCost )
	var first = ((cost - 50) * 1)
	var second =  ((cost - 100) * 1)
	if first < 0 {//account for less than 50 dollars spent
		first = 0
	}
	if second < 0{ // account for less than 100 dollars spent
		second = 0
	}
	var total = first + second
	fmt.Printf("rewarded client %d  total \n", total )
	if total >= 1 {
		return total
	}
	return 0
}

*/





func main() { // setup router and api
	
  r := gin.Default()
  r.Use(cors.Default())
  r.HEAD("/api/health/", health)
  r.GET("/", test)
  /*
  r.POST("/api/createbook", createbook)
  r.GET("/api/getbooks/", getbooks)
  r.GET("/api/getbook/{id}", getbook)
  r.PUT("/api/updatebook/{id}", updatebook)
  r.DELETE("/api/deletebook/{id}", deletebook)
  */
  //r.PATCH("/api/updatebook/patch/{id}", updatebookpatch)
  
   /* test
   s := "This,is,a,delimited,string"
   v := strings.Split(s, ",")  
    for i := range v {
        p := v[i]
        fmt.Println(p)
    }
   fmt.Println("v %v", v)     // [This is a delimited string]  
  */
  /* test
  s1 := "postgres://jsfdkljkdsfkljsljkfds:alkjdfsjlkfsdlkfjjdslkafjlksdsdfjljsdflkljdsjfkldsjklflkf@dsklfkdsf.com:111111/jihuguyrtdfgyuhujkopijuy" ; User,Password,host,port,dbname
  temp1 := strings.Split(s1, "://")//remove
  v1 := temp1[1]
  fmt.Println("v1 ",v1)
  temp2 := strings.Split(v1, ":")
  v2 := temp2[0]
  v3 := temp2[1]
  v4 := temp2[2]
  fmt.Println("user ",v2)
  fmt.Println("v3 ",v3)
  temp3 := strings.Split(v3, "@")
  v5 := temp3[0]
  v6 := temp3[1]
  fmt.Println("password ",v5)
  fmt.Println("host ",v6)
  temp4 := strings.Split(v4, "/")
  v7 := temp4[0]
  v8 := temp4[1]
  fmt.Println("port ",v7)
  fmt.Println("dbname ",v8)
  // v := strings.Split(s, "://")  
   // for i := range v {
   //     p := v[i]
	//	fmt.Println("split ", p)
    //    fmt.Printf("%v", s)
    //}
  */
   
  /* manage db */
  fmt.Println("preparing db")
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  // CLI TOOL postgres://jsfdkljkdsfkljsljkfds:alkjdfsjlkfsdlkfjjdslkafjlksdsdfjljsdflkljdsjfkldsjklflkf@dsklfkdsf.com:111111/asjkdlfjklfdsklfjsdlakjfsdjflksdjkf
  // // format for gorm connection // dsn := "dbname=asjkdlfjklfdsklfjsdlakjfsdjflksdjkf host=dsklfkdsf.com port=111111 user=jsfdkljkdsfkljsljkfds password=alkjdfsjlkfsdlkfjjdslkafjlksdsdfjljsdflkljdsjfkldsjklflkf sslmode=require TimeZone=America/New_York""
  // INFORMATION FROM https://data.heroku.com/ "DATABASE CREDENTIALS"; go https://data.heroku.com/ then to app then "settings" then credentials
   
  
  //use db url
  s1 := os.Getenv("DATABASE_URL") // DATABASE_URL
  temp1 := strings.Split(s1, "://")//remove
  v1 := temp1[1]
  fmt.Println("v1 ",v1)
  temp2 := strings.Split(v1, ":")
  v2 := temp2[0]
  v3 := temp2[1]
  v4 := temp2[2]
  fmt.Println("user ",v2)
  fmt.Println("v3 ",v3)
  temp3 := strings.Split(v3, "@")
  v5 := temp3[0]
  v6 := temp3[1]
  fmt.Println("password ",v5)
  fmt.Println("host ",v6)
  temp4 := strings.Split(v4, "/")
  v7 := temp4[0]
  v8 := temp4[1]
  fmt.Println("port ",v7)
  fmt.Println("dbname ",v8)

  dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=require TimeZone=America/New_York", v6, v2, v5, v8, v7)
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  fmt.Println("dsn ", dsn )
  fmt.Println("DB_HOST ",os.Getenv("DB_HOST")," DB_USER ",os.Getenv("DB_USER")," DB_PASSWORD ",os.Getenv("DB_PASSWORD")," DB_NAME ",os.Getenv("DB_NAME"), " DB_PORT ", os.Getenv("DB_PORT") )
  if(err != nil){ 
	fmt.Println("db connection error") 
	log.Fatal("could not connect to db");
  } 
  fmt.Println("db connection made")  
  db.Migrator().CreateTable(&Book{})
  bool := db.Migrator().HasTable(&Book{})
  if(bool == true){
	fmt.Println("created db table(s)")
  }else{
	fmt.Println("failed to create db table(s)")
  }
  fmt.Println("preparing router") 
  port := os.Getenv("PORT")
  if port == "" { //heroku assignment fail
		r.Run(":8080")
		fmt.Println("started router") 
		//log.Fatal("$PORT must be set")//let heroku log error
  }
  r.Run(":" + port)
  
  fmt.Println("started router") 
}
