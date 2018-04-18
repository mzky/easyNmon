package main
import ( 
        "github.com/gin-gonic/gin"
	"flag"
        "os/exec"
        "fmt"
)
func main() {
	port := flag.String("port","8080","http listen port")
	flag.Parse()
        r := gin.Default()
        r.GET("/:value", func(c *gin.Context) {
                value := c.Param("value") //取value值
		lsCmd := exec.Command("/bin/sh", "-c", "./nmonCTL.sh "+value)
		err := lsCmd.Start()  
                if err!=nil{
                       	fmt.Println(err)
                }	
                //fmt.Print(string(out))
                c.JSON(200, gin.H{
                      	"message": value,
                })
        })
	sport := ":"
	sport += *port
        r.Run(sport) // listen and serve on 0.0.0.0:8080
}
