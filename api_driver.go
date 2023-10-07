package main
import(
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)
//using gin framework for simple api


// structure for data
type student struct {
	ID      string `json:"id"`
	NAME    string `json:"name"`
	MAJOR   string `json:"major"`
	GRADYEAR int   `json:"gradyear"`
} 

var students = []student{
	{ID: "1", NAME:"Gio", MAJOR: "CS", GRADYEAR:2025},
	{ID: "2", NAME:"Nicole", MAJOR: "Math", GRADYEAR:2025},
	{ID: "3", NAME:"Seb", MAJOR: "Math", GRADYEAR:2023},
	{ID: "4", NAME:"Mike", MAJOR: "Business", GRADYEAR:2024},
	{ID: "5", NAME:"Yigit", MAJOR: "CS", GRADYEAR:2030},
	{ID: "6", NAME:"Carlos", MAJOR: "CS", GRADYEAR:2030},
	{ID: "7", NAME:"Harrison", MAJOR: "CS", GRADYEAR:2030},
}

//making function to get student
func getStudents(c *gin.Context){
	//send status if this funciton can run and the information
	c.IndentedJSON(http.StatusOK, students)
}

func studentByID(c *gin.Context){

	id := c.Param("id")
	student, err := getStudentByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student ID not found"})
		return
	}

	//if there is no error then return ok stat and the value
	c.IndentedJSON(http.StatusOK, student)
}

//helper function to search student by id
//uses id as a string and returns pointer to student and err
func getStudentByID(id string) (*student, error){
	// look through all the values and if there is a val that matches input then return pointer to index and nil error
	for i, b := range students{
		if b.ID == id{
			return &students[i], nil
		}
	}
	//else if theyre not found then return nothing and an error that says ""
	return nil, errors.New("student not found")
}

func main(){
	router := gin.Default()
	router.GET("/students", getStudents)
	router.GET("/students/:id", studentByID)
	router.Run("0.0.0.0:8080")
	//use curl ip:port /location to retrive info as json
}	
