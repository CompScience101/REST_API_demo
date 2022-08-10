package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	//"text/template"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/driver/postgres"

	//"github.com/go-gorm/gorm"
	"gorm.io/gorm"
)

// Book struct (Model for other)
type Book struct {
	Isbn   string `json:"isbn" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author_name" binding:"required"`
	gorm.Model
	//DateCreated string `json:"date_created"`
}

// Book struct (Model for patch)
type BookPatch struct {
	Isbn   string `json:"isbn" `
	Title  string `json:"title"`
	Author string `json:"author_name"`
	gorm.Model
	//DateCreated string `json:"date_created"`
}

// Init books var as a slice Book struct
var books []Book
var db *gorm.DB

func health(c *gin.Context) { // monitor health of api
	c.JSON(http.StatusNoContent, nil)
	fmt.Printf("replied to client health check with success status.")
}
func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "testing connection", "test2": "github actions ci"})
}

func createbook(c *gin.Context) { //not worrying about rejecting duplicate books, since library can have duplicates
	var myjson Book
	if err := c.ShouldBindJSON(&myjson); err != nil {
		//fmt.Printf("c.ShouldBindJSON(&myjson) err:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": string(err.Error())})
		fmt.Printf("Createbook invalid request")
		return
	}
	fmt.Printf("sending book to db")
	result := db.Create(&myjson)
	if result.Error == nil {
		c.JSON(http.StatusCreated, myjson)
		fmt.Printf("created book. sending client response \n")
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		fmt.Printf("error creating book in db. sending client response \n")
	}

}
func getbooks(c *gin.Context) {
	result := db.Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(result.Error.Error())})
		fmt.Printf("db error while retrieving all books")
		return
	} else if books == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no records available"})
		fmt.Printf("db error, no books while retrieving all books")
		return
	}
	c.JSON(http.StatusOK, books)
	fmt.Printf("retrieved all books. sent client response \n")
}
func getbook(c *gin.Context) {
	var myjson Book

	//convert id
	id := c.Param("id")
	idparse, err := strconv.ParseUint(id, 10, 32)
	myjson.ID = uint(idparse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please send a int for book id, for reading up single book."})
		fmt.Println("Incorrect book id formatting, for reading single book. error ", string(err.Error()))
	}

	//find book
	var mybook Book
	result := db.First(&mybook, id)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) { //no book db error
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("No book exist with id:%v", id)})
		fmt.Println("db error while retrieving book id:", id, " to read it. Book doesn't exist. Error was: ", result.Error.Error())
		return
	} else if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) { //other db error
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(result.Error.Error())})
		fmt.Println("db error while retrieving book id:", id, " to read it. Error was: ", result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, mybook)
	fmt.Printf("getbook read book. sent client response")

}

func updatebookput(c *gin.Context) {
	var myjson Book

	if err := c.ShouldBindJSON(&myjson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": string(err.Error())})
		fmt.Printf("updatebookput invalid request error: %v", err)
		return
	}
	//convert id
	id := c.Param("id")
	idparse, err := strconv.ParseUint(id, 10, 32)
	myjson.ID = uint(idparse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please send a int for book id, for updating(PUT) book."})
		fmt.Println("Incorrect book id formatting, for updating(PUT) book. error: ", string(err.Error()))
	}

	//find book
	var mybook Book
	result := db.First(&mybook, id)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) { //no book db error
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("No book exist with id:%v", id)})
		fmt.Println("db error while retrieving book id:", id, " to update(PUT) it. Book doesn't exist. Error was: ", result.Error.Error())
		return
	} else if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) { //other db error
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(result.Error.Error())})
		fmt.Println("db error while retrieving book id:", id, " to update(PUT) it. Error was: ", result.Error.Error())
		return
	}

	result = db.Save(&myjson)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(result.Error.Error())})
		fmt.Printf("updatebookput internal error while updating book in db")
		return
	}
	c.JSON(http.StatusOK, myjson)
	fmt.Printf("updatebookput updated book. sent client response")

}

func updatebookpatch(c *gin.Context) {
	var myjson BookPatch

	if err := c.ShouldBindJSON(&myjson); err != nil { //note: gorm will update db record even if user uses a emtpy json object
		c.JSON(http.StatusBadRequest, gin.H{"error": string(err.Error())})
		fmt.Printf("updatebookpatch invalid request error: %v", err)
		return
	}
	//fmt.Printf("myjson:", myjson)//check if json object is empty
	//convert id
	id := c.Param("id")
	idparse, err := strconv.ParseUint(id, 10, 32)
	myjson.ID = uint(idparse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please send a int for book id, for updating(PATCH) book."})
		fmt.Println("Incorrect book id formatting, for updating(PATCH) book. error ", string(err.Error()))
	}

	//find book
	var mybook Book
	result := db.First(&mybook, id)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) { //no book db error
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("No book exist with id:%v", id)})
		fmt.Println("db error while retrieving book id:", id, " to update(PATCH) it. Book doesn't exist.")
		return
	} else if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) { //other db error
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(result.Error.Error())})
		fmt.Println("db error while retrieving book id:", id, " to update(PATCH) it. Error was: ", result.Error.Error())
		return
	}

	result = db.Model(&mybook).Updates(&myjson)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(result.Error.Error())})
		fmt.Printf("updatebookpatch internal error while updating book in db")
		return
	}
	c.JSON(http.StatusOK, mybook)
	fmt.Printf("updatebookpatch updated book. sent client response")

}

func deletebook(c *gin.Context) {

	//convert id
	id := c.Param("id")
	idparse, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please send a int for book id, to delete book."})
		fmt.Println("Incorrect book id formatting, for deleting book. error ", string(err.Error()))
		return
	}
	// find book
	var mybook Book
	result := db.First(&mybook, id)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) { //no book db error
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("No book exist with id:%v", id)})
		fmt.Println("db error while retrieving book id:", id, " to delete it. Book doesn't exist.")
		return
	} else if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) { //other db error
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(result.Error.Error())})
		fmt.Println("db error while retrieving book id:", id, " to delete it. Error was: ", result.Error.Error())
		return
	}

	//delete book
	result = db.Delete(&Book{}, idparse)
	if result.Error != nil { //no book db error
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(result.Error.Error())})
		fmt.Println("db error while retrieving book id:", id, " to delete it.")
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": fmt.Sprintf("deleted book id:%v", id)})
	fmt.Printf("deletebook deleted book id:%v . sent client response", id)
}

func main() { // setup router and api

	r := gin.Default()
	r.Use(cors.Default())
	r.HEAD("/api/health/", health)
	r.GET("/", test)
	r.POST("/api/createbook", createbook)
	r.GET("/api/getbooks/", getbooks)
	r.GET("/api/getbook/:id", getbook)
	r.PUT("/api/updatebook/:id", updatebookput)
	r.PATCH("/api/updatebook/:id", updatebookpatch)
	r.DELETE("/api/deletebook/:id", deletebook)

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
	s1 := os.Getenv("DATABASE_URL")   // DATABASE_URL
	temp1 := strings.Split(s1, "://") //remove
	v1 := temp1[1]
	fmt.Println("v1 ", v1)
	temp2 := strings.Split(v1, ":")
	v2 := temp2[0]
	v3 := temp2[1]
	v4 := temp2[2]
	fmt.Println("user ", v2)
	fmt.Println("v3 ", v3)
	temp3 := strings.Split(v3, "@")
	v5 := temp3[0]
	v6 := temp3[1]
	fmt.Println("password ", v5)
	fmt.Println("host ", v6)
	temp4 := strings.Split(v4, "/")
	v7 := temp4[0]
	v8 := temp4[1]
	fmt.Println("port ", v7)
	fmt.Println("dbname ", v8)

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=require TimeZone=America/New_York", v6, v2, v5, v8, v7)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//fmt.Printf("gorm db type %T",db)
	fmt.Println("dsn ", dsn)
	fmt.Println("DB_HOST ", os.Getenv("DB_HOST"), " DB_USER ", os.Getenv("DB_USER"), " DB_PASSWORD ", os.Getenv("DB_PASSWORD"), " DB_NAME ", os.Getenv("DB_NAME"), " DB_PORT ", os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println("db connection error")
		log.Fatal("could not connect to db")
	}
	fmt.Println("db connection made")

	bool := db.Migrator().HasTable(&Book{}) //check status
	if bool == true {

		fmt.Println("already created db table(s)")
	} else {
		db.Migrator().CreateTable(&Book{})
		bool := db.Migrator().HasTable(&Book{})
		if bool == true {
			fmt.Println("created db table(s)")
		} else {
			fmt.Println("failed to create db table(s)")
		}
	}

	/* delete old table
	db.Migrator().DropTable(&Book{})
	bool := db.Migrator().HasTable(&Book{})
	if bool == true {
		fmt.Println("db table(s) still exist")
	} else {
		fmt.Println("deleted db table(s)")
	}
	*/
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
