package httpGo
import(
	"net/http"
	"fmt"
	"os"
	"io"
	"crypto/md5"
	"encoding/hex"
	"github.com/alruiz12/simpleBT/src/conf"
	"time"
	"sync"
	"strings"
)

func PutObjAPI(w http.ResponseWriter, r *http.Request){
	var startPUT time.Time
	startPUT = time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		//name := md5String(r.URL.Path)
		results:= strings.Split(r.URL.Path, "/")	// ["",account, container, object]
		addedResults:=results[1]+results[2]+results[3]
		file, err := os.Create(os.Getenv("GOPATH") + "/src/github.com/alruiz12/simpleBT/src/" + addedResults )
		if err != nil {
			fmt.Println(err)
		}
		_, err = io.Copy(file, r.Body)
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
		putOK := make(chan bool)

		go PutObjProxy(os.Getenv("GOPATH") + "/src/github.com/alruiz12/simpleBT/src/" + addedResults, conf.TrackerAddr, conf.NumNodes, putOK ,results[1], results[2], results[3])
		success := <-putOK
		if success == true {
			fmt.Println("put success ", time.Since(startPUT))
			w.WriteHeader(http.StatusCreated)
		} else {
			fmt.Println("put fail")
			w.WriteHeader(http.StatusBadRequest)
		}
		currentKey := md5sum(os.Getenv("GOPATH") + "/src/github.com/alruiz12/simpleBT/src/" + addedResults)
		fmt.Println(CheckPiecesObj(currentKey, "NEW.xml", conf.FilePath, conf.NumNodes))
	}()
	wg.Wait()
	fmt.Println("API: ",time.Since(startPUT))
}


func md5String(str string) string{
	hasher:=md5.New()
	_, err:= hasher.Write([]byte(str))
	if err != nil {
		fmt.Println(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

func PutAccAPI(w http.ResponseWriter, r *http.Request){
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		//accountName:=r.Header["Name"][0]
		accountName:=r.URL.Path[1:]
		if accountName==""{
			fmt.Println("create fail")
			w.WriteHeader(http.StatusBadRequest)
		} else{
			createOK := make(chan bool)
			go CreateAccountProxy(accountName, createOK)


			//go Put(os.Getenv("GOPATH") + "/src/github.com/alruiz12/simpleBT/src/" + name, conf.TrackerAddr, conf.NumNodes, putOK)
			success := <-createOK
			if success == true {
				fmt.Println("create success ")
				w.WriteHeader(http.StatusCreated)
			} else {
				fmt.Println("create fail")
				w.WriteHeader(http.StatusBadRequest)
			}
		}

	}()
	wg.Wait()
}


func PutContAPI(w http.ResponseWriter, r *http.Request){
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		//accountName:=r.Header["Name"][0]
		accountName:=r.URL.Path[1:]
		results:= strings.Split(accountName, "/")
		fmt.Println("PutContAPI: ",results[1])
		if accountName==""{
			fmt.Println("put fail")
			w.WriteHeader(http.StatusBadRequest)
		} else{
			createOK := make(chan bool)
			fmt.Println("PutContAPI")
			go putContProxy(results[0], results[1], createOK)


			success := <-createOK
			if success == true {
				fmt.Println("put success ")
				w.WriteHeader(http.StatusCreated)
			} else {
				fmt.Println("put fail")
				w.WriteHeader(http.StatusBadRequest)
			}
		}

	}()
	wg.Wait()
}



