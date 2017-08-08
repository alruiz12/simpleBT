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
	"encoding/json"
)

func PutObjAPI(w http.ResponseWriter, r *http.Request){
	var startPUT time.Time
	startPUT = time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
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

		go PutObjProxy(os.Getenv("GOPATH") + "/src/github.com/alruiz12/simpleBT/src/" + addedResults, conf.TrackerAddr, conf.NumNodes, putOK ,results[1], results[2], results[3], addedResults)
		success := <-putOK
		if success == true {
			fmt.Println("put success ", time.Since(startPUT))
			w.WriteHeader(http.StatusCreated)
		} else {
			fmt.Println("put fail")
			w.WriteHeader(http.StatusBadRequest)
		}

		os.Remove(os.Getenv("GOPATH") + "/src/github.com/alruiz12/simpleBT/src/" + addedResults )
		if err != nil {
			fmt.Println(err)
		}
	}()
	wg.Wait()
	fmt.Println("PUT: ",time.Since(startPUT))
}

func GetObjAPI(w http.ResponseWriter, r *http.Request){
	var startGET time.Time
	startGET = time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		results:= strings.Split(r.URL.Path, "/")	// ["",account, container, object]
		addedResults:=results[1]+results[2]+results[3]

		GetOK := make(chan bool)
		go GetObjProxy(addedResults, conf.ProxyAddr, conf.TrackerAddr, GetOK,results[1], results[2], results[3])
		success := <-GetOK
		if success == true {
			fmt.Println("get success ", time.Since(startGET))
			w.WriteHeader(http.StatusOK)
		} else {
			fmt.Println("get fail")
			w.WriteHeader(http.StatusBadRequest)
		}

		//Todo put it in Get
		//currentHash := md5sum(os.Getenv("GOPATH") + "/src/github.com/alruiz12/simpleBT/src/" + addedResults)
		//fmt.Println(CheckPiecesObj(addedResults, "NEW.xml", conf.FilePath, conf.NumNodes, currentHash))

	}()
	wg.Wait()
	fmt.Println("GET API: ",time.Since(startGET))
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
		accountName:=r.URL.Path[1:]
		if accountName==""{
			fmt.Println("create fail")
			w.WriteHeader(http.StatusBadRequest)
		} else{
			createOK := make(chan bool)
			go CreateAccountProxy(accountName, createOK)


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


func GetAccAPI(w http.ResponseWriter, r *http.Request){
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		accountName:=r.URL.Path[1:]
		if accountName==""{
			fmt.Println("create fail")
			w.WriteHeader(http.StatusBadRequest)
		} else{
			getOK := make(chan bool)
			account:= GetAccountProxy(accountName, getOK)

			success := <-getOK
			if success == true {
				fmt.Println("get success ")

				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				if err := json.NewEncoder(w).Encode(account); err != nil {
					fmt.Println("GetNodes: error encoding response: ",err.Error())
				}
				w.WriteHeader(http.StatusOK)
			} else {
				fmt.Println("get fail")
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



