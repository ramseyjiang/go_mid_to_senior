package ginpkg

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseParams struct {
	Name   string `json:"name" uri:"name"`
	Age    int    `json:"age" uri:"age"`
	Gender string `json:"gender" uri:"gender"`
}

// gin is a go framework, response of all get and post requests usage are under below.
// In command line, go run pkgusages/gin.go.
// It will monitor all requests from browser.
// r means router

func Trigger() {
	r := gin.Default()

	// Response for all get requests
	jsonResponse(r)
	jsonResponseParamAndQuery(r)
	jsonResponseFullPath(r)
	jsonResponseBind(r)
	jsonResponseBindURI(r)
	getResponseString(r)

	// Reponse for all post requests
	postJSONResponse(r)
	multiParamsFormResponse(r)
	queryAndFormPostResponse(r)

	_ = r.Run(":8888") // All requests will access port:8888.
}

// Get Request: http://localhost:8888/ping
func jsonResponse(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin framework base usage.",
		})
	})
}

/*
	Get Response: http://localhost:8888/get/123?user=davy&pwd=admin
	c.Param is used to get arguments before "?".
	c.DefaultQuery will get key value first, if it does not have, it will use the default one.
	c.Query is used to get arguments after "?".
*/
func jsonResponseParamAndQuery(router *gin.Engine) {
	// This handler will match /user/{id} but will not match /user/ or /user
	router.GET("/get/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.DefaultQuery("user", "jeff")
		pwd := c.Query("pwd")

		// c.JSON return a json.
		c.JSON(http.StatusOK, gin.H{
			"message": "hell gyy",
			"id":      id,
			"user":    user,
			"pwd":     pwd,
		})
	})
}

// http://localhost:8888/user/23, it will have a download confirm page after you add any string after the 23.
// You should add any string after http://localhost:8888/user/23/string, it will trigger download.
// "c.String" will output string, but it always download a file.
func getResponseString(router *gin.Engine) {
	// However, this one will match /user/{id}/ and also /user/{id}/send
	// If no other routers match /user/{id}, it will redirect to /user/{id}/
	router.GET("/user/:id/*action", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		action := c.Param("action")
		message := "User ID is " + fmt.Sprint(id) + ", action is " + action

		c.String(http.StatusOK, message)
	})
}

// http://localhost:8888/user/groups, it will show the whole route path.
func jsonResponseFullPath(router *gin.Engine) {
	// For each matched request Context will hold the route definition
	router.GET("/user/groups", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": c.FullPath(),
		})
	})
}

// http://localhost:8888/testBind, it will show json response.
func jsonResponseBind(router *gin.Engine) {
	router.GET("/testBind", func(c *gin.Context) {
		p := ResponseParams{}
		err := c.ShouldBind(&p) // Here if it uses c.ShouldBindJSON(), it will have an error. So change it to c.ShouldBind() fix it.
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "Something wrong",
				"data": gin.H{},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "success",
				"data": p,
			})
		}
	})
}

// Get request: http://localhost:8888/getParamsFromUrl/davy/40/male
func jsonResponseBindURI(router *gin.Engine) {
	router.GET("/getParamsFromUrl/:name/:age/:gender", func(c *gin.Context) {
		p := ResponseParams{}
		err := c.ShouldBind(&p) // //Here if it uses c.ShouldBindUri(), it will have an error. So change it to c.ShouldBindUri() fix it.
		p.Name = c.Param("name")
		p.Age, _ = strconv.Atoi(c.Param("age"))
		p.Gender = c.Param("gender")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "Something wrong!",
				"data": gin.H{},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "success",
				"data": p,
			})
		}
	})
}

/*
	Post request: http://localhost:8888/post
	c.DefaultPostForm will get key value first, if it doesn't have, it will use the default value.
*/
func postJSONResponse(router *gin.Engine) {
	router.POST("/post", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "jeff")
		pwd := c.DefaultPostForm("pwd", "pwd")
		c.JSON(http.StatusOK, gin.H{
			"message": "hell gyy",
			"user":    user,
			"pwd":     pwd,
		})
	})
}

// Post request: http://localhost:8888/form_post
func multiParamsFormResponse(router *gin.Engine) {
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
}

/*
	Post request: http://localhost:8888/query_form_post
	c.Query is used to get arguments before "?".
	c.PostForm is used to get arguments from post form.
*/
func queryAndFormPostResponse(router *gin.Engine) {
	router.POST("/query_form_post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		// During output Using fmt.Printf does not need convert string to int.
		// Using fmt.Printf, it won't have an output on browser, but it will have the output on command line.
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)

		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})
}
