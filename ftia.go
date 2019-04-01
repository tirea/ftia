package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/c-bata/go-prompt"
)

const (
	//tODO: make this configurable
	language string = "eng"
)

var (
	numwords   int
	knownIDs   []string
	selected   []string
	usr, _     = user.Current()
	homeDir, _ = filepath.Abs(usr.HomeDir)
	dataDir    = filepath.Join(homeDir, ".ftia")
	fname_m    = filepath.Join(dataDir, "metaWords.txt")
	fname_d    = filepath.Join(dataDir, "localizedWords.txt")
	fname_k    = filepath.Join(dataDir, "known.txt")
)

func linecount() {
	mfile, err := os.Open(fname_m)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(mfile)
	for scanner.Scan() {
		numwords++
	}
	err = mfile.Close()
	if err != nil {
		panic(err)
	}
}

func contains(q []string, s string) bool {
	for _, v := range q {
		if s == v {
			return true
		}
	}
	return false
}

func sel(n string, a bool) {
	selected = []string{}
	in, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		panic(err)
	}
	for len(selected) < int(in) {
		mfile, err := os.Open(fname_m)
		if err != nil {
			panic(err)
		}
		ln := 0
		rn := rand.Intn(numwords)
		scanner := bufio.NewScanner(mfile)
		for scanner.Scan() {
			line := scanner.Text()
			ln++
			if ln == rn {
				fields := strings.Split(line, "\t")
				ID := fields[0]
				if contains(knownIDs, ID) && !a {
					continue
				}
				w := fields[1]
				ipa := fields[2]
				inf := ""
				if fields[3] != "NULL" {
					inf = fields[3]
				}
				pos := fields[4]
				selected = append(selected, ID)
				fmt.Printf("[%d] %s [%s] %s %s\n", len(selected), w, ipa, inf, pos)
			}
		}
		err = mfile.Close()
		if err != nil {
			panic(err)
		}
	}
}

func known(s []string) {
	if len(s) != 0 {
		for _, v := range s {
			pv, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic(err)
			}
			knownIDs = append(knownIDs, selected[int(pv)-1])
		}
	}
}

func define(s []string) {
	if len(s) != 0 {
		for _, v := range s {
			dfile, err := os.Open(fname_d)
			if err != nil {
				panic(err)
			}
			pv, err := strconv.ParseInt(v, 10, 64)
			ID := selected[int(pv)-1]
			scanner := bufio.NewScanner(dfile)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, ID+"\t") && strings.Contains(line, language) {
					fields := strings.Split(line, "\t")
					d := fields[2]
					pos := fields[3]
					fmt.Printf("%s %s\n", pos, d)
				}
			}
			err = dfile.Close()
			if err != nil {
				panic(err)
			}
		}
	}
}

func save() {
	kfile, err := os.OpenFile(fname_k, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range knownIDs {
		if v != "" {
			if _, err := kfile.Write([]byte(v + ",")); err != nil {
				log.Fatal(err)
			}
		}
	}
	if err := kfile.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("data saved")
}

func load() {
	kfile, err := os.Open(fname_k)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(kfile)
	for scanner.Scan() {
		line := scanner.Text()
		k := strings.Split(line[:len(line)], ",")
		for _, kv := range k {
			knownIDs = append(knownIDs, kv)
		}
	}
	err = kfile.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("data loaded\n")
}

func executor(cmd string) {
	s := strings.Split(cmd, " ")
	switch s[0] {
	case "/select":
		sel(s[1], false)
	case "/selectfromall":
		sel(s[1], true)
	case "/known":
		known(s[1:])
	case "/define":
		define(s[1:])
	case "/q", "/quit", "/exit":
		save()
		os.Exit(0)
	}
	fmt.Println()
}

func completer(d prompt.Document) []prompt.Suggest {
	if d.GetWordBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	s := []prompt.Suggest{
		{Text: "/select", Description: ""},
		{Text: "/selectfromall", Description: ""},
		{Text: "/known", Description: ""},
		{Text: "/define", Description: ""},
		{Text: "/quit", Description: ""},
		{Text: "/exit", Description: ""},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	head := "Ftia v0.0.1-dev by Tirea Aean"
	fmt.Println(head)
	linecount()
	rand.Seed(time.Now().UTC().UnixNano())
	load()
	p := prompt.New(executor, completer,
		prompt.OptionTitle("Ftia"),
		prompt.OptionPrefix("~~> "),
		prompt.OptionSelectedDescriptionTextColor(prompt.DarkGray),
	)
	p.Run()
}
