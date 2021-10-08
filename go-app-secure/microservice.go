package main

import (
	//"crypto/tls"
	//"crypto/x509"
	//"io"
	//"io/ioutil"
	//"log"
	"fmt"
	"os"
	"net/http"
    "github.com/rashika/go-app-secure/pkg/api"
	//"github.com/PacktPublishing/Cloud-Native-Go/api"
)

func main() {
	fmt.Println("server started!")
	http.HandleFunc("/", index)
	//http.HandleFunc("/api/echo", api.EchoHandleFunc)
	//http.HandleFunc("/api/hello", api.HelloHandleFunc)

	http.HandleFunc("/api/song", api.SongsHandleFunc) // store and retrieve all songs
	http.HandleFunc("/api/song/", api.SongHandleFunc) // retrieve songs by IDs, update nad delete
	//http.HandleFunc("/api/books/", api.BookHandleFunc)



	// Create a CA certificate pool and add cert.pem to it
	/*caCert, err := ioutil.ReadFile("goappsecure.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create the TLS Config with the CA pool and enable Client certificate validation
	tlsConfig := &tls.Config{
		ClientCAs: caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	// Create a Server instance to listen on port 8443 with the TLS config
	server := &http.Server{
		Addr:      port(),
		TLSConfig: tlsConfig,
	}

	// Listen to HTTPS connections with the server certificate and wait
	log.Fatal(server.ListenAndServeTLS("goappsecure.crt", "goappsecure.key"))
}


*/











	err:= http.ListenAndServeTLS(port(), "gsd.crt", "gsd.key", nil)
	if err!=nil{
		fmt.Println(err.Error())
	}
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Cloud Native Go!")
}
