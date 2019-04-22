package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func read0(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("%s\n", err)
		panic(err)
	}
	return string(f)
}

func read1(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

func read2(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 1024, 1024)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

func read3(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func main() {

	dir0, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir0)

	abs, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	log.Println(abs)

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	log.Println(exPath)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)

	paths := []string{
		"/home/arnie/amelia.jpg",
		"/mnt/photos/",
		"rabbit.jpg",
		"/usr/local//go",
		"/usr/local/go/bin/x.md",
	}
	for _, p := range paths {
		dir := filepath.Dir(p)
		parent := filepath.Dir(dir)
		base := filepath.Base(dir)
		log.Printf("input: %q\n\tdir: %q\n\tparent: %q\nbase: %q\n", p, dir, parent, base)
	}

	fp := path.Join(path.Dir("/opt/config/test"), "../config/settings.toml")
	log.Println(fp)

	//file := "./test.log"
	//
	//start := time.Now()
	//
	//read0(file)
	//t0 := time.Now()
	//log.Printf("Cost time %v\n", t0.Sub(start))
	//
	//read1(file)
	//t1 := time.Now()
	//log.Printf("Cost time %v\n", t1.Sub(t0))
	//
	//read2(file)
	//t2 := time.Now()
	//log.Printf("Cost time %v\n", t2.Sub(t1))
	//
	//read3(file)
	//t3 := time.Now()
	//log.Printf("Cost time %v\n", t3.Sub(t2))

}
