# jwt-gin
User module 

## Step 1: Initing module

Use the cmd below to create a directory and go into it.

```cmd
mkdir jwt-gin
cd jwt-gin
```

Create a go.mod file to help manage the libraries we need.

```cmd
go mod init jwt-gin
```

Install the libs we need.

```cmd
// gin framework
go get -u github.com/gin-gonic/gin
// ORM library
go get -u github.com/jinzhu/gorm
// package that we will be used to authenticate and generate our JWT
go get -u github.com/dgrijalva/jwt-go
// to help manage our environment variables
go get -u github.com/joho/godotenv
// to encrypt our user's password
go get -u golang.org/x/crypto
```

## Step 2: Starting Serv

Create a file called 'main.go' as a entry point in the base directory and input the code below.

```go 
package main
 
import (
    "net/http"
 
    "github.com/gin-gonic/gin"
)
 
func main() {
 
    r := gin.Default()
 
    public := r.Group("/api")
 
    public.POST("/register", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})
    })
 
    r.Run(":8080")
}
```

Start Serv

```cmd
go run main.go
```

The output of the process proved that the serv is listening on port 8000.

```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/register             --> main.main.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
```

Use the postman to test the "/register" route and the serv will respond with the content below

```
{
    "data": "this is the register endpoint!"
}
```

## Step 3: User Register


```
