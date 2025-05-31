package api

import (
	"embed"

	"github.com/gin-gonic/gin"
	db "github.com/louischering/simplebank/db/sqlc"
)

// Server serves HTTP requests for baking app.
type Server struct {
	store      db.Store
	router     *gin.Engine
	templateFS *embed.FS
	count      int
}

var id = 2

type Contact struct {
	Id    int
	Name  string
	Email string
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

type Page struct {
	Data Data
	Form FormData
}

func NewServer(store db.Store, templateFS *embed.FS) *Server {
	// data := Data{Contacts{{Id: 0, Name: "Louis", Email: "email@email.com"}, {Id: 1, Name: "SomeoneElse", Email: "other@email.com"}}}
	server := &Server{store: store, templateFS: templateFS, count: 0}
	router := gin.Default()

	// router.LoadHTMLGlob("views/*")

	// router.Static("/css", "./css")

	router.POST("api/accounts", server.createAccount)
	router.GET("api/accounts/:ID", server.getAccount)
	router.GET("api/accounts/", server.listAccounts)

	// router.GET("/", func(c *gin.Context) {
	// 	formData := FormData{
	// 		Values: make(map[string]string),
	// 		Errors: make(map[string]string),
	// 	}
	// 	page := Page{
	// 		Data: data,
	// 		Form: formData,
	// 	}
	// 	formData.Values["name"] = ""
	// 	formData.Values["email"] = ""
	// 	c.HTML(http.StatusOK, "index", page)
	// })

	// router.DELETE("/contacts/:id", func(c *gin.Context) {
	// 	time.Sleep(time.Second * 2)

	// 	id := c.Param("id")
	// 	var updatedPeople []Contact
	// 	strId, _ := strconv.Atoi(id)
	// 	for _, person := range data.Contacts {
	// 		if person.Id != strId {
	// 			updatedPeople = append(updatedPeople, person)
	// 		}
	// 	}
	// 	data.Contacts = updatedPeople
	// 	c.Status(http.StatusNoContent)
	// })

	// router.POST("/contacts", func(c *gin.Context) {
	// 	var newContact Contact
	// 	id++
	// 	newContact.Id = id
	// 	newContact.Name = c.Request.FormValue("name")
	// 	newContact.Email = c.Request.FormValue("email")
	// 	if checkEmailExists(data.Contacts, newContact.Email) {
	// 		formData := FormData{
	// 			Errors: make(map[string]string),
	// 			Values: make(map[string]string),
	// 		}
	// 		formData.Values["name"] = newContact.Name
	// 		formData.Values["email"] = newContact.Email
	// 		formData.Errors["email"] = "email already exists"
	// 		c.HTML(http.StatusUnprocessableEntity, "form", formData)
	// 		return
	// 	}
	// 	data.Contacts = append(data.Contacts, newContact)
	// 	c.HTML(http.StatusOK, "oob-contact", newContact)
	// 	c.HTML(http.StatusOK, "form", FormData{
	// 		Errors: make(map[string]string),
	// 		Values: make(map[string]string),
	// 	})
	// })

	server.router = router
	return server
}

func checkEmailExists(emails []Contact, email string) bool {
	for _, contact := range emails {
		if contact.Email == email {
			return true
		}
	}
	return false
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// loadTemplate loads templates embedded by go-assets-builder
// func (s *Server) loadTemplate() error {
// 	entries, err := s.templateFS.ReadDir("templates")
// 	if err != nil {
// 		return err
// 	}
// 	for _, entry := range entries {
// 		print(entry)
// 	}
// 	return nil, nil
// }
