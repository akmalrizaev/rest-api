package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simpleapi/internal/models"
	"strconv"
	"strings"
	"sync"
)

var teachers = make(map[int]models.Teacher)

var mutex = &sync.Mutex{}

var nextID = 1

// Initialize some dummy data
func init() {
	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Math",
	}
	nextID++
	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextID++
	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Doe",
		Class:     "11A",
		Subject:   "Biology",
	}
	nextID++
}

func TeachersHandler(w http.ResponseWriter, r *http.Request) {

	// teachers/{id}
	// teachers/9

	// Query Params
	// teachers/?key=value&query=value2

	switch r.Method {
	case http.MethodGet:

		/*
			fmt.Println(r.URL.Path)
			path := strings.TrimPrefix(r.URL.Path, "/teachers/")
			userID := strings.TrimSuffix(path, "/")
			fmt.Println("The ID is:", userID)

			fmt.Println("Query Params", r.URL.Query())
			queryParams := r.URL.Query()
			key := queryParams.Get("key")
			fmt.Printf("Key: %v", key)

		*/

		// w.Write([]byte("Hello GET Method on Teachers Route"))
		getTeachersHandler(w, r)

	case http.MethodPost:

		addTeacherHandler(w, r)

		/*

			Parse form data (necessary for x-www-form-urlencoded)
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}

			fmt.Println("Form", r.Form)

			Prepare response data
			response := make(map[string]interface{})
			for key, value := range r.Form {
				response[key] = value
				response[key] = value[0]

			}

			fmt.Println("Processed Response Map:", response)

			RAW Body
			body, err := io.ReadAll(r.Body)
			if err != nil {
				return
			}
			defer r.Body.Close()

			fmt.Println("RAW Body without converting to string:", body)
			fmt.Println("RAW Body", string(body))

			if we expect json data, then unmarshal it
			var userInstance User
			err = json.Unmarshal(body, &userInstance)
			if err != nil {
				return
			}

			fmt.Println("JSON Unmarshaling:", userInstance)
			fmt.Println("Received user name as:", userInstance.Name)

			Access the request details
			fmt.Println("Access the request details:")
			fmt.Println("Body:", r.Body)
			fmt.Println("Form", r.Form)
			fmt.Println("Header", r.Header)
			fmt.Println("Context", r.Context())
			fmt.Println("ContentLength", r.ContentLength)
			fmt.Println("Host", r.Host)
			fmt.Println("Method", r.Method)
			fmt.Println("Protocol", r.Proto)
			fmt.Println("Remote Addr", r.RemoteAddr)
			fmt.Println("Request URI", r.RequestURI)
			fmt.Println("TLS", r.TLS)
			fmt.Println("Trailers", r.Trailer)
			fmt.Println("Transfer Encoding", r.TransferEncoding)
			fmt.Println("URL", r.URL)
			fmt.Println("User Agent", r.UserAgent())
			fmt.Println("Port", r.URL.Port())
			fmt.Println("URL", r.URL.Scheme)

			w.Write([]byte("Hello POST Method on Teachers Route"))
			return

		*/

	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Teachers Route"))
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Teachers Route"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Teachers Route"))
		return
	}

	// if r.Method == http.MethodGet {
	// 	w.Write([]byte("Hello GET Method on Teachers Route"))
	// 	return
	// }
	// w.Write([]byte("Hello from Teachers Route"))
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")

	if idStr == "" {
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		teacherList := make([]models.Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {

				teacherList = append(teacherList, teacher)
			}
		}

		response := struct {
			Status string           `json:"status"`
			Count  int              `json:"count"`
			Data   []models.Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teacherList),
			Data:   teacherList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

	// Handle Path parameter
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	teacher, exists := teachers[id]
	if !exists {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(teacher)

}

func addTeacherHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var newTeachers []models.Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return

	}

	addedTeachers := make([]models.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		newTeacher.ID = nextID
		teachers[nextID] = newTeacher
		addedTeachers[i] = newTeacher
		nextID++
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string           `json: "status"`
		Count  int              `json: "count"`
		Data   []models.Teacher `json: "data"`
	}{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}
	json.NewEncoder(w).Encode(response)

}
