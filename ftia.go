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
	// tODO: make language configurable
	language string = "eng"
	// field indices of the data file
	idField  int = 0
	navField int = 1
	ipaField int = 2
	infField int = 3
	posField int = 4
	defField int = 5
	srcField int = 6
)

type entry struct {
	ID             string
	Navi           string
	IPA            string
	InfixLocations string
	PartOfSpeech   string
	Definition     string
	Source         string
}

var (
	numwords   int
	knownIDs   []string
	selected   map[int]entry
	reverse    bool = false
	usr, _          = user.Current()
	homeDir, _      = filepath.Abs(usr.HomeDir)
	dataDir         = filepath.Join(homeDir, ".ftia")
	fname_d         = filepath.Join(dataDir, "dictionary_"+language+".txt")
	fname_k         = filepath.Join(dataDir, "known.txt")
	fname_kr        = filepath.Join(dataDir, "known_rev.txt")
)

func linecount() {
	dfile, err := os.Open(fname_d)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(dfile)
	for scanner.Scan() {
		numwords++
	}
	err = dfile.Close()
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

func mapContains(m map[int]entry, i string) bool {
	for _, x := range m {
		if x.ID == i {
			return true
		}
	}
	return false
}

func sel(n string, k bool, a bool) {
	selected = make(map[int]entry)
	if n == "" {
		return
	}
	in, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		return
	}
	for len(selected) < int(in) {
		e := new(entry)
		dfile, err := os.Open(fname_d)
		if err != nil {
			log.Fatal(err)
		}
		ln := 0
		rn := rand.Intn(numwords)
		scanner := bufio.NewScanner(dfile)
		for scanner.Scan() {
			line := scanner.Text()
			ln++
			if ln == rn {
				fields := strings.Split(line, "\t")
				ID := fields[idField]
				kmatch := contains(knownIDs, ID)
				if a || (kmatch && k) || (!kmatch && !k) {
					e.ID = fields[idField]
					e.Navi = fields[navField]
					e.IPA = fields[ipaField]
					e.InfixLocations = ""
					if fields[infField] != "NULL" {
						e.InfixLocations = fields[infField]
					}
					e.PartOfSpeech = fields[posField]
					e.Definition = fields[defField]
					e.Source = fields[srcField]
					if mapContains(selected, e.ID) {
						continue
					}
					selected[len(selected)+1] = *e
					if reverse {
						fmt.Printf("[%d] %s %s\n", len(selected), e.PartOfSpeech, e.Definition)
					} else {
						if e.InfixLocations == "" {
							fmt.Printf("[%d] %s [%s] %s\n", len(selected), e.Navi, e.IPA, e.PartOfSpeech)
						} else {
							fmt.Printf("[%d] %s [%s] %s %s\n", len(selected), e.Navi, e.IPA, e.InfixLocations, e.PartOfSpeech)
						}
					}
				}
			}
		}
		err = dfile.Close()
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
				continue
			}
			if contains(knownIDs, selected[int(pv)].ID) {
				continue
			}
			knownIDs = append(knownIDs, selected[int(pv)].ID)
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
				continue
			}
			idx := 0
			ID := selected[int(pv)].ID
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
				continue
			}
			e := selected[int(pv)]
			if e.InfixLocations == "" {
				fmt.Printf("[%d] %s [%s] %s %s\n", pv, e.Navi, e.IPA, e.PartOfSpeech, e.Definition)
			} else {
				fmt.Printf("[%d] %s [%s] %s %s %s\n", pv, e.Navi, e.IPA, e.InfixLocations, e.PartOfSpeech, e.Definition)
			}
		}
	}
}

func source(s []string) {
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
				continue
			}
			e := selected[int(pv)]
			fmt.Printf("[%d] %s: %s\n", pv, e.Navi, e.Source)
		}
	}
}

func progress() {
	k := len(knownIDs)
	p := 100 * (float64(k) / float64(numwords))
	fmt.Printf("%.2f%% (%d / %d)\n", p, k, numwords)
}

func save(fname string) {
	sfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range knownIDs {
		if v != "" {
			if _, err := sfile.Write([]byte(v + ",")); err != nil {
				log.Fatal(err)
			}
		}
	}
	if err := sfile.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("data saved")
}

func load(fname string) {
	knownIDs = []string{}
	lfile, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(lfile)
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
	err = lfile.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data loaded")
}

func rev() {
	if reverse {
		save(fname_kr)
		load(fname_k)
	} else {
		save(fname_k)
		load(fname_kr)
	}
	reverse = !reverse
}

func executor(cmd string) {
	s := strings.Split(cmd, " ")
	if len(s) > 0 {
		if contains([]string{"/q", "/quit", "/exit", "/kä", "/hum"}, s[0]) {
			if reverse {
				save(fname_kr)
			} else {
				save(fname_k)
			}
			os.Exit(0)
		} else if contains([]string{"/progress", "/p", "/holpxaype", "/polpxay"}, s[0]) {
			progress()
		} else if contains([]string{"/switch", "/change", "/reverse", "/r", "/latem"}, s[0]) {
			rev()
		}
	}
	if len(s) > 1 {
		switch s[0] {
		case "/select", "/s", "/ftxey":
			sel(s[1], false, false)
		case "/known", "/k", "/nolume":
			sel(s[1], true, false)
		case "/selectfromall", "/sa", "/fratsim":
			sel(s[1], true, true)
		case "/add", "/a", "/sung":
			add(s[1:])
		case "/delete", "/del", "/'aku":
			del(s[1:])
		case "/define", "/def", "/d", "/ralpeng":
			define(s[1:])
		case "/source", "/src", "/tsim":
			source(s[1:])
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
		{Text: "/ftxey", Description: "ftxey kawnomuma aylì'ut nìfya'o arenulke"},
		{Text: "/selectfromall", Description: "select n random words both learned and unlearned"},
		{Text: "/fratsim", Description: "ftxey frafnelì'ut nìfya'o arenulke"},
		{Text: "/known", Description: "select n random learned words"},
		{Text: "/nolume", Description: "ftxey awnomuma aylì'ut nìfya'o arenulke"},
		{Text: "/define", Description: "show definition / translation for given entry in selection"},
		{Text: "/ralpeng", Description: "wìntxu ralit aylì'uä"},
		{Text: "/source", Description: "show canon source of given entry"},
		{Text: "/tsim", Description: "wìntxu tsimit aylì'uä"},
		{Text: "/add", Description: "mark given entries known / learned"},
		{Text: "/sung", Description: "sung sna'or aylì'uä awnomum"},
		{Text: "/delete", Description: "unmark given entries known / learned"},
		{Text: "/'aku", Description: "'aku ta sna'o aylì'uä awnomum"},
		{Text: "/reverse", Description: "reverse the direction of na'vi<->local"},
		{Text: "/latem", Description: "sar lahea lì'fyati tup sar lì'fyati leNa'vi"},
		{Text: "/progress", Description: "show current progress of words learned out of words in the dictionary"},
		{Text: "/polpxay", Description: "wìntxu holpxayt aylì'uä awnomum"},
		{Text: "/quit", Description: "save and quit program"},
		{Text: "/exit", Description: "save and quit program"},
		{Text: "/hum", Description: "pamrel si fte ziverok ulte ftang fìkem sivi"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	head := "Ftia v2.1.0-dev by Tirea Aean"
	fmt.Println(head)
	linecount()
	rand.Seed(time.Now().UTC().UnixNano())
	load(fname_k)
	fmt.Println()
	p := prompt.New(executor, completer,
		prompt.OptionTitle("Ftia"),
		prompt.OptionPrefix("~~> "),
		prompt.OptionSelectedDescriptionTextColor(prompt.DarkGray),
	)
	p.Run()
}
