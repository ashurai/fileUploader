package handlres

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
    "path/filepath"
	"time"
	"github.com/ashurai/fileUploader/database"
	"github.com/ashurai/fileUploader/model"
	"github.com/ashurai/fileUploader/helpres"
)
var basePath = "temp/files/"

func ChangeProfileImage(w http.ResponseWriter, r *http.Request) {
	var file model.File
	 
	imageName, err := helpres.FileUpload(r)
	if err != nil {
	    http.Error(w, "Invalid Data", http.StatusBadRequest)
    	return
	    //checking whether any error occurred retrieving image
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) // get the root directory path, 
		// if need web path set the host and concatinet
    if err != nil {
            log.Fatal(err)
    }
    fmt.Println(dir)
	file.Name = dir + basePath + imageName
	file.Path = imageName
	file.CreatedAt = time.Now()
	//if no error insert uploade & save in DB
	fmt.Println(file)
    database.SaveFile(&file)
    w.Header().Set("Content-Type", "application/json")
    //we then return the new user details to update our user interface
    json.NewEncoder(w).Encode(file)
}