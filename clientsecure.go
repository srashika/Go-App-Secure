package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("client started")
	// Read the key pair to create certificate
	/*cert, err := tls.LoadX509KeyPair("go-app-secure/goappsecure.crt", "go-app-secure/goappsecure.key")
	if err != nil {
		log.Fatal(err)
	}
    */
	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile("go-app-secure/gsd.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	req, err := http.NewRequest("GET", "https://go-app-secure.default:8080/api/song", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Create a HTTPS client and supply the created CA pool and certificate
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
				//Certificates: []tls.Certificate{cert},
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()


	b, err := ioutil.ReadAll(resp.Body)

	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier

	if err != nil {

		log.Fatalln(err)

	}



	fmt.Println(string(b))

    
	
	






	// Request /hello via the created HTTPS client over port 8443 via GET
/*	r, err := client.Get("https://10.190.212.195:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body to stdout
	fmt.Printf("%s\n", body)
	*/
}
