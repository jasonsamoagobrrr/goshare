package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	// proxy flag
	var proxy string
	flag.StringVar(&proxy, "p", "127.0.0.1:9050", "Set a socks5 proxy (defaults to self embedded tor)")

	// key flag
	var key string
	flag.StringVar(&key, "k", "", "Enter a passphrase that will become a 16 bit key")

	// encrypt flag
	var encryption bool
	flag.BoolVar(&encryption, "e", false, "encrypt the file(s), must pass key flag ")

	// decrypt flag
	var decrypt bool
	flag.BoolVar(&decrypt, "d", false, "decrypt a file or directory")

	// path flag, if not specified grab first arg after flags processed
	var path string
	flag.StringVar(&path, "f", "", "specify the file or directory")

	// concurrency flag
	var concurrency int
	flag.IntVar(&concurrency, "c", 20, "set the concurrency level")

	// recursive flag
	var recursive bool
	flag.BoolVar(&recursive, "r", false, "recursive directory upload")

	// parse flags
	flag.Parse()

	// if file was not provided in the parsed flags, then first parsed arg is assumed as the file
	if path == "" {
		path = flag.Arg(0)
	}

	// check if encryption is on and if the pass was given
	if encryption == true {
		if key == "" {
			log.Fatal("No pass phrase provided")
		} else {
			println("your hash: " + createHash(key))
		}
	}

	if decrypt == true {
		if key == "" {
			log.Fatal("No pass phrase provided")
		}
	}

	// Define work channels
	files := make(chan string)
	output := make(chan string)

	// logic to check if file or directory
	dirOrnoDir, err := isDir(path)
	if err != nil {
		log.Print(err)
	}

	// processing work group
	var procWG sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		procWG.Add(1)
		go func() {
			for i := range files {
				if decrypt != true {
					output <- callUpload("https://pomf.cat/upload.php", proxy, i, encryption, key)
				} else {
					output <- decryptFile(i, key)
				}
			}
			procWG.Done()
		}()
	}

	// Output worker
	var outputWG sync.WaitGroup
	outputWG.Add(1)
	go func() {
		for o := range output {

			fmt.Println(o)
		}
		outputWG.Done()
	}()

	// Close the output channel when the HTTP workers are done
	go func() {
		procWG.Wait()
		close(output)
	}()

	// logic for loading channels if user specified a dir
	if dirOrnoDir == true {
		directory, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		// non-recursive file add to file channel
		if recursive == false {

			// add all the listings that are files to file channel
			for _, file := range directory {
				if file.IsDir() == false {
					files <- path + "/" + file.Name()
				}
			}

			// recursively add to file channel
		} else {
			err := filepath.Walk(path, func(pathBoi string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Println(err)
					return err
				}
				if info.IsDir() == false {

					// make sure under 75 mebibyte
					files <- pathBoi
				}
				return nil
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		files <- path

		//test := <-files
		time.Sleep(10)
	}
	close(files)

	// results work group (outputs all the URLs and keys)
	time.Sleep(1000000000)

}

func callUpload(urlPath string, proxy string, filePath string, encrypt bool, key string) string {

	//Go struct for the JSON response
	type Response struct {
		Success bool `json:"success"`
		Files   []struct {
			Hash string `json:"hash"`
			Name string `json:"name"`
			URL  string `json:"url"`
			Size int    `json:"size"`
		} `json:"files"`
	}

	// Create var to store response in
	var result Response

	// New HTTP Client
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// New multipart writer (simulate -F of curl)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create file field and a form file writer for the field
	fw, err := writer.CreateFormFile("files[]", filePath)
	if err != nil {
	}

	// create file var
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
		log.Print(err)
	}

	// make encrypted file var
	var encryptedFile string

	// make file data var
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
		log.Print(err)
	}

	// Open file
	if encrypt == true {
		// get encrypted file name/location
		encryptedFile = encryptFile(fileData, key)
		// set file var to encrypted file
		file, err = os.Open(encryptedFile)
		if err != nil {
			panic(err)
			log.Print(err)
		}
	}

	// Copy file into the multipart file writer
	_, err = io.Copy(fw, file)
	if err != nil {
		log.Print(err)
	}

	// delete temp encrypted files
	os.Remove(encryptedFile)

	// Close multipart writer
	writer.Close()

	// Build the request
	req, err := http.NewRequest("POST", urlPath, bytes.NewReader(body.Bytes()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Do the request
	rsp, _ := client.Do(req)
	// Get body of response
	rspBody, err := ioutil.ReadAll(rsp.Body)
	// Unmarshal the response body and convert it to our matching go struct
	if err := json.Unmarshal(rspBody, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	// Loop through the go struct and get our URL/Filename
	var fileUrl string
	for _, rec := range result.Files {
		fileUrl = "https://a.pomf.cat/" + string(rec.URL)
	}
	// Let us know if the Request failed
	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}
	return fileUrl
}

func isDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func createHash(key string) string {
	// hash the key/passphrase
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	// create new AES cypher
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	// create nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func encryptFile(data []byte, passphrase string) string {
	f, err := os.CreateTemp(os.TempDir(), "tempPre")
	if err != nil {
		log.Fatal(err.Error())
	}
	f.Write(encrypt(data, passphrase))
	return f.Name()
	//defer f.Close()
}

func decryptFile(filename string, passphrase string) string {
	data, _ := ioutil.ReadFile(filename)
	file, _ := os.Open(filename)
	file.Write(decrypt(data, passphrase))
	return filename + "decrypted"
}
