package main

import (
  "os"
  "fmt"
  "github.com/docker/docker/pkg/archive"
/*
  "code.google.com/p/lzma"
  "archive/tar"
  "compress/gzip"
  "io"
  "strings"
*/
)

func main() {
  tar, err_open := os.Open(os.Args[1])
  if err_open != nil {
    fmt.Printf("Error Opening File: %s\n", err_open)
  }

  err_tar := archive.Untar(tar, "/", nil)
  if err_tar != nil {
    fmt.Printf("Error Extracting: %s\n", err_tar)
    os.Exit(1)
  }
  fmt.Println("Archive Extracted.")
}
/*
  sourcefile := os.Args[1]

  file, err := os.Open(sourcefile)

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  defer file.Close()

  var fileReader io.ReadCloser = file

  // if input is a tar.gz file
  if strings.HasSuffix(sourcefile, ".gz") {
    if fileReader, err = gzip.NewReader(file); err != nil {
      fmt.Println(err)
         os.Exit(1)
      }
    defer fileReader.Close()
  }

  // if input is a tar.xz file
  if strings.HasSuffix(sourcefile, ".xz") {
   //TODO: write lzma2 reader
    fileReader := lzma.NewReader(file);
    defer fileReader.Close()
  }

  tarBallReader := tar.NewReader(fileReader)

  // Extracting tarred files

  for {
    header, err := tarBallReader.Next()
    if err != nil {
      if err == io.EOF {
        break
      }
      fmt.Println(err)
      os.Exit(1)
    }

    // get the individual filename and extract to the current directory
    filename := header.Name

    switch header.Typeflag {
      case tar.TypeDir:
        // handle directory
        fmt.Println("Creating directory :", filename)
        err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

        if err != nil {
          fmt.Println(err)
          os.Exit(1)
        }

      case tar.TypeReg:
        // handle normal file
        fmt.Println("Untarring :", filename)
        writer, err := os.Create(filename)

        if err != nil {
          fmt.Println(err)
          os.Exit(1)
        }

        io.Copy(writer, tarBallReader)

        err = os.Chmod(filename, os.FileMode(header.Mode))

        if err != nil {
          fmt.Println(err)
          os.Exit(1)
        }

        writer.Close()
      default:
        fmt.Printf("Unable to untar type : %c in file %s", header.Typeflag, filename)
    }
  }

}*/
