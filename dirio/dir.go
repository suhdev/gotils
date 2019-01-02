// Package dirio used to read directory
package dirio

import (
	"io/ioutil"
	"path"
)

func readDirectory(dir string, fc chan *Request, filter func(string) bool) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fc <- &Request{
			reqType: reqFinish,
			payload: nil,
		}
		return
	}
	for _, f := range files {
		name := path.Join(dir, f.Name())
		if f.IsDir() {
			fc <- &Request{
				reqType: reqDir,
				payload: name,
			}
		} else {
			if filter != nil {
				if filter(name) {
					fc <- &Request{
						reqType: reqFile,
						payload: name,
					}
				}
			} else {
				fc <- &Request{
					reqType: reqFile,
					payload: name,
				}
			}
		}
	}

	fc <- &Request{
		reqType: reqFinish,
		payload: nil,
	}
}

const (
	reqDir = iota
	reqFile
	reqFinish
	reqExit
)

type Request struct {
	reqType  int
	payload  interface{}
	response chan interface{}
}

// ReadDir reads files within a directory as a string array
// dir the directory to read
// maxCon the no. of concurrent goroutines to use to read directory recursively
// filter a filter function to filter files (can be nil)
func ReadDir(dir string, maxCon int, filter func(string) bool) []string {
	files := make([]string, 0)
	f := ReadDirChan(dir, maxCon, filter)
	for x := range f {
		files = append(files, x)
	}
	return files
}

// ReadDirChan reads files within a directory as a string channel
// dir the directory to read
// maxCon the no. of concurrent goroutines to use to read directory recursively
// filter a filter function to filter files (can be nil)
func ReadDirChan(dir string, maxCon int, filter func(string) bool) chan string {
	files := make(chan string, 10000)
	reqs := make(chan *Request, 1000)
	dirs := make([]string, 0)
	count := 0

	go func() {
		for r := range reqs {
			if r.reqType == reqExit {
				break
			} else if r.reqType == reqFile {
				if x, ok := r.payload.(string); ok {
					files <- x
				}
			} else if r.reqType == reqDir {
				if x, ok := r.payload.(string); ok {
					if count < maxCon {
						count++
						go readDirectory(x, reqs, filter)
					} else {
						dirs = append(dirs, x)
					}
				}
			} else if r.reqType == reqFinish {
				count--
				if len(dirs) > 0 {
					x, a := dirs[0], dirs[1:]
					dirs = a
					count++
					go readDirectory(x, reqs, filter)
				} else if count == 0 {
					break
				}
			}
		}
		close(reqs)
		close(files)
	}()
	reqs <- &Request{
		reqType: reqDir,
		payload: dir,
	}
	return files
}
