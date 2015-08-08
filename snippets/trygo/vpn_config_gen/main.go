


package main

import (
	"os"
	"io"
	"text/template"
	"fmt"
	"archive/zip"
)

const (
	vpn_extname = ".ovpn"
	zip_extname = ".zip"	
	template_file = "client.ovpn"

)
var (
	extnames = []string{".crt", ".key", ".ovpn"}
	files_list = []string{"ca.crt"}
)


type Cfg struct {
	ClientName string
}
func main() {


	if (len(os.Args) !=2) {
		fmt.Fprintf(os.Stderr, "usage: %s cfgname", os.Args[0])
		os.Exit(1)
	}
	cfg := &Cfg{os.Args[1]}

	if t, err := template.ParseFiles(template_file); err == nil {
		if f, err := os.Create(cfg.ClientName + vpn_extname); err == nil {
			defer f.Close()
			t.Execute(f, cfg)
	
		}
	
	} else {
		fmt.Fprintf(os.Stderr, "load template %s error %s", os.Args[0], err)
		os.Exit(1)
	
	}
	
	for _,extname := range extnames {
		files_list = append(files_list, cfg.ClientName + extname)
	}
	
	if zf, err := os.Create(cfg.ClientName + zip_extname); err ==nil {
		defer zf.Close()
		if z := zip.NewWriter(zf); z != nil {
			defer z.Close()		
			for _, fn := range(files_list) {
				fmt.Fprintf(os.Stderr, "zip %s\n", fn)
				if ef, err := os.Open(fn); err == nil {
					if zef, err := z.Create(fn); err == nil {
						io.Copy(zef, ef)
						z.Flush()
					}
				}
			}

		}
	}

}

