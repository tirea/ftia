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
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(mfile)
	for scanner.Scan() {
		numwords++
	}
	err = mfile.Close()
	if err != nil {
		log.Fatal(err)
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

func sel(n string, k bool, a bool) {
	selected = []string{}
	if n == "" {
		return
	}
	in, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	for len(selected) < int(in) {
		mfile, err := os.Open(fname_m)
		if err != nil {
			log.Fatal(err)
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
				if contains(knownIDs, ID) {
					if k || a {
						w := fields[1]
						ipa := fields[2]
						inf := ""
						if fields[3] != "NULL" {
							inf = fields[3]
						}
						pos := fields[4]
						if contains(selected, ID) {
							continue
						}
						selected = append(selected, ID)
						fmt.Printf("[%d] %s [%s] %s %s\n", len(selected), w, ipa, inf, pos)
					}
				} else {
					if !k || a {
						w := fields[1]
						ipa := fields[2]
						inf := ""
						if fields[3] != "NULL" {
							inf = fields[3]
						}
						pos := fields[4]
						if contains(selected, ID) {
							continue
						}
						selected = append(selected, ID)
						fmt.Printf("[%d] %s [%s] %s %s\n", len(selected), w, ipa, inf, pos)
					}
				}
			}
		}
		err = mfile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func add(s []string) {
	if len(selected) == 0 {
		return
	}
	if len(s) != 0 {
		for _, v := range s {
			if v == "" {
				continue
			}
			pv, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			if contains(knownIDs, selected[int(pv)-1]) {
				continue
			}
			knownIDs = append(knownIDs, selected[int(pv)-1])
		}
	}
}

func del(s []string) {
	if len(selected) == 0 {
		return
	}
	if len(s) != 0 {
		for _, v := range s {
			if v == "" {
				continue
			}
			pv, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			idx := 0
			ID := selected[int(pv)-1]
			for i, x := range knownIDs {
				if x == ID {
					idx = i
				}
			}
			knownIDs[idx] = knownIDs[len(knownIDs)-1]
			knownIDs[len(knownIDs)-1] = ""
			knownIDs = knownIDs[:len(knownIDs)-1]
		}
	}
}

func define(s []string) {
	if len(s) != 0 {
		for _, v := range s {
			if v == "" {
				continue
			}
			dfile, err := os.Open(fname_d)
			if err != nil {
				log.Fatal(err)
			}
			pv, err := strconv.ParseInt(v, 10, 64)
			ID := selected[int(pv)-1]
			scanner := bufio.NewScanner(dfile)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, ID+"\t") && strings.Contains(line, language+"\t") {
					fields := strings.Split(line, "\t")
					d := fields[2]
					pos := fields[3]
					fmt.Printf("%s %s\n", pos, d)
				}
			}
			err = dfile.Close()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func progress() {
	k := len(knownIDs)
	p := 100 * (float64(k) / float64(numwords))
	fmt.Printf("%.2f%% (%d / %d)\n", p, k, numwords)
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
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(kfile)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		k := strings.Split(line[:len(line)-1], ",")
		for _, kv := range k {
			knownIDs = append(knownIDs, kv)
		}
	}
	err = kfile.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data loaded\n")
}

func executor(cmd string) {
	s := strings.Split(cmd, " ")
	if len(s) > 0 {
		if contains([]string{"/q", "/quit", "/exit"}, s[0]) {
			save()
			os.Exit(0)
		} else if s[0] == "/progress" {
			progress()
		}
	}
	if len(s) > 1 {
		switch s[0] {
		case "/select":
			sel(s[1], false, false)
		case "/known":
			sel(s[1], true, false)
		case "/selectfromall":
			sel(s[1], true, true)
		case "/add":
			add(s[1:])
		case "/delete":
			del(s[1:])
		case "/define":
			define(s[1:])
		}
	}
	fmt.Println()
}

func completer(d prompt.Document) []prompt.Suggest {
	if d.GetWordBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	s := []prompt.Suggest{
		{Text: "/select", Description: "select n random unlearned words"},
		{Text: "/selectfromall", Description: "select n random words both learned and unlearned"},
		{Text: "/known", Description: "select n random learned words"},
		{Text: "/define", Description: "show definition / translation for given entry in selection"},
		{Text: "/add", Description: "mark given entries known / learned"},
		{Text: "/delete", Description: "unmark given entries known / learned"},
		{Text: "/progress", Description: "show current progress of words learned out of words in the dictionary"},
		{Text: "/quit", Description: "save and quit program"},
		{Text: "/exit", Description: "save and quit program"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	head := "Ftia v0.0.2-dev by Tirea Aean"
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
