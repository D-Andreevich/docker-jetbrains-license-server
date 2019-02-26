package main
 
import (
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "fmt"
    "log"
    "net/http"
)
 
const (
    cert = `-----BEGIN RSA PRIVATE KEY-----
MIIBOwIBAAJBALecq3BwAI4YJZwhJ+snnDFj3lF3DMqNPorV6y5ZKXCiCMqj8OeO
mxk4YZW9aaV9ckl/zlAOI0mpB3pDT+Xlj2sCAwEAAQJAW6/aVD05qbsZHMvZuS2A
a5FpNNj0BDlf38hOtkhDzz/hkYb+EBYLLvldhgsD0OvRNy8yhz7EjaUqLCB0juIN
4QIhAMsJQ3xiJemnJ2pD65iRNCC/Kr7jtxbbBwa6ZFLjp12pAiEA54JCn41fF8GZ
90b9L5dtFQB2/yIcGX4Xo7bCvl8DaPMCIBgOZ+2T33QYtwXTOFXiVm/O1qy5ZFcT
6ng0m3BqwsjJAiEAqna/l7wAyP1E4U7kHqbhKxWsiTAUgLDXtzRbMNHFMQECIQCA
xlpXEPqnC3P8if0G9xHomqJ531rOJuzB8fNtRFmxnA==
-----END RSA PRIVATE KEY-----`
)
 
var privateKey *rsa.PrivateKey
 
func init() {
    block, _ := pem.Decode([]byte(cert))
 
    key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        log.Panicln(err)
    }
    privateKey = key
}
 
func main() {
    http.HandleFunc("/rpc/releaseTicket.action", releaseTicketHandler)
    http.HandleFunc("/rpc/obtainTicket.action", obtainTicketHandler)
    http.HandleFunc("/rpc/ping.action", pingTicketHandler)

    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)

    log.Println("IntelliJIdeaLicense Server Start listen on http://0.0.0.0:1234")
 
    log.Fatal(http.ListenAndServe(":1234", nil))
}
 
func releaseTicketHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Release Ticket request")
    http.Error(w, "not found", 404)
}
 
func obtainTicketHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Obtain Ticket request")
    salt := r.URL.Query().Get("salt")
//    userName := r.URL.Query().Get("userName")
    userName := "JetBrains"
 
    str := fmt.Sprintf("<ObtainTicketResponse>"+
        "<message></message>"+
        "<prolongationPeriod>999999999</prolongationPeriod>"+
        "<responseCode>OK</responseCode>"+
        "<salt>%s</salt>"+
        "<ticketId>1</ticketId>"+
        "<ticketProperties>licensee=%s    licenseType=0   </ticketProperties>"+
        "</ObtainTicketResponse>", salt, userName)
 
    writeAnswer(w, str)
}
 
func pingTicketHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Ping Ticket request")
    salt := r.URL.Query().Get("salt")
 
    str := fmt.Sprintf("<PingResponse>"+
        "<message></message>"+
        "<responseCode>OK</responseCode>"+
        "<salt>%s</salt>"+
        "</PingResponse>", salt)
 
    writeAnswer(w, str)
}
 
func writeAnswer(w http.ResponseWriter, str string) {
    h := crypto.MD5.New()
    h.Write([]byte(str))
    hashed := h.Sum(nil)
 
    bs, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.MD5, hashed)
 
    str = fmt.Sprintf("<!-- %x -->\n%s", string(bs), str)
    w.Write([]byte(str))
}
