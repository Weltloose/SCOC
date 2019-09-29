package main

import (
	"io"
	"log"
	"os"
	"os/exec"

	flag "github.com/spf13/pflag"
)

const BUF_SIZE = 2048

var start = flag.IntP("start", "s", -1, "the start page num to fetch.")
var end = flag.IntP("end", "e", -1, "the end page num to fetch.")
var lines = flag.IntP("lines", "l", 72, "fixed line number of a page.")
var dest = flag.StringP("dest", "d", "", "the dest printer num.")
var f_mark = flag.BoolP("f_mark", "f", false, "pages end at form-feed.")

type read_decision func(int, byte) bool

func main() {
	flag.Parse()

	var to_read io.Reader = os.Stdin

	if flag.NArg() == 1 {
		var err error
		to_read, err = os.OpenFile(flag.Arg(0), os.O_RDONLY, 0000)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *start <= 0 || *end <= 0 || *end < *start {
		log.Fatal("wrong style of start and end")
	}

	var to_write io.Writer = os.Stdout

	if *dest != "" {
		cmd := exec.Command("lp", "-d", *dest)
		to_write, _ = cmd.StdinPipe()
		cmd.Start()
	}

	var fc read_decision

	if *f_mark {
		fc = read_with_fmark
	} else {
		fc = read_with_lines
	}

	read_pages(to_read, fc, to_write)

}

func read_with_fmark(l int, c byte) bool {
	return c == '\f'
}

func read_with_lines(l int, c byte) bool {
	return c == '\n' && l == *lines
}

func read_pages(f io.Reader, fc read_decision, w io.Writer) {
	buf := make([]byte, BUF_SIZE)
	ln := 0
	pgs := 1
	s := -1
	if *start == 1 { // 如果是第一页需要额外判定
		s = 0
	}
	for n, _ := f.Read(buf); n > 0; n, _ = f.Read(buf) {
		if s >= 0 {      // 如果从之前的buf就开始打印了，那么这里要把它置为0
			s = 0
		}
		e := n           // 默认到buf读入的末尾，如果在buf中出现了终止的情况再修改e
		for i := 0; i < n; i++ {
			if buf[i] == '\n' {
				ln++
			}
			if fc(ln, buf[i]) {     // 判断是否换页
				ln = 0
				pgs++
				if pgs == *start {
					s = i + 1
				} else if pgs == *end+1 {
					if i == 0 {
						return // 如果它是在下一个buf的第一个字节，则说明该有的已经全都打印了
					}
					e = i + 1
				}

			}

		}
		if s != -1 {
			w.Write(buf[s:e])
		}
		if e != n {
			return
		}

	}
}
