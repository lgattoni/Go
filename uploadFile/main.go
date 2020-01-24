package main

import (
        "fmt"
        "io/ioutil"
        "os/exec"
        "net/http"
)



func uploadmyfile(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Upload File")

    r.ParseMultipartForm(10 << 20)
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    fmt.Fprintf(w, "Uploaded File:", handler.Filename, "\n")
    fmt.Fprintf(w, "\n")
    fmt.Fprintf(w, "File Size:", handler.Size, "\n")
    fmt.Fprintf(w, "\n")
    fmt.Fprintf(w, "MIME Header:", handler.Header, "\n")
    fmt.Fprintf(w, "\n")

    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)


    tempFile, err := ioutil.TempFile("/opt/upload", handler.Filename)
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    // tempFile.Write(fileBytes)

    exec.Command("rm", "-f", "/opt/upload/*[0-9]")
    ioutil.WriteFile("/opt/upload/" + handler.Filename, fileBytes, 0666)
    exec.Command("rm", "-f", "/opt/upload/" + handler.Filename + "[0-9]*")
    fmt.Fprintf(w, "\n")
    fmt.Fprintf(w, "=================================\n")
    fmt.Fprintf(w, "UPLOAD AREA\n")
    fmt.Fprintf(w, "=================================\n")
    fmt.Fprintf(w, "Successfully Uploaded File\n")

    out, err := exec.Command("rm", "-f", "/opt/upload/" + handler.Filename + "[0-9]*").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println("Command rm  Successfully Executed")
    output := string(out[:])
    fmt.Println(output)

}

func myweb() {
    http.HandleFunc("/opt/upload", uploadmyfile)
    http.ListenAndServe(":8085", nil)
}

func main() {
    fmt.Println("Upload Files into /opt/upload/")
    myweb()
}