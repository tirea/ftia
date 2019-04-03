# ftia

A simple text-based vocabulary review app

## Installation

First, make sure you have a Go compiler installed and in your PATH. Visit [https://golang.org](https://golang.org) for download.

### macOS / Linux

Run the following inside a Terminal:

```
cd $HOME
mkdir -p go/src
export GOPATH=$HOME/go
cd ~/go/src
go get github.com/c-bata/go-prompt
git clone https://github.com/tirea/ftia
cd ftia
make install
```

### Windows

Run the fellowing inside a Powershell:
```
cd $HOME
mkdir go\src
setx GOPATH $HOME\go
```
Open a new Powershell, then continue:
```
cd go\src\
go get github.com/c-bata/go-prompt
git clone https://github.com/tirea/ftia
cd ftia
go build -o $HOME\ftia.exe .\ftia.go
copy -Recurse .\.ftia $HOME\
```

## Commands & Aliases

`/select x`, `/s x`, `/ftxey x` : select x random unlearned words

`/selectfromall x`, `/sa x`, `/fratsim`: select x random words both learned and unlearned

`/known x...`, `/k x`, `/nolume` : select x random learned words

`/define n...`, `/def n...`, `/d n...`, `/ralpeng n...`: show definition / translation for given entry in selection

`/source n...`, `/src n...`, `/tsim n...` : show canon source of given entry

`/add n...`, `/a n...`, `/sung n...` : mark given entries known / learned

`/delete n...`, `/del n...`, `/'aku n...` : unmark given entries known / learned

`/switch`, `/change`, `/reverse`, `/r`, `/latem` : reverse the direction of na'vi<->local

`/progress`, `/p`, `/holpxaype`, `/polpxay` : show current progress of words learned out of words in the dictionary

`/quit`, `/exit`, `/q`, `/kä`, `/hum` : save data and quit the program.

Where

`x` is a positive integer number from 1 to the number of words you wish to review with the command,

`n` is a positive integer number from 1 to `x`, and

`...` represents optional additional space-separated `n` in a list.

## Examples

```
➜ ./ftia
Ftia v1.0.0-dev by Tirea Aean
data loaded

~~> /select 5
[1] lefpomron [lɛ.fpom.ˈɾon] adj.
[2] yemfpay [jɛm.ˈfpaj] n.
[3] pxaw [p'aw] adp.
[4] tìterkup [tɪ.ˈtɛɾ.kup̚] n.
[5] tsyosyu [ˈt͡sjo.sju] n.

~~> /add 3 4 5

~~> /add 2

~~> /define 1
[1] lefpomron [lɛ.fpom.ˈɾon] adj. healthy (mentally)

~~> /select 10
[1] netrìp [ˈnɛt.ɾɪp̚] adv.
[2] lopx [l·op'] l<1><2><3>opx vin.
[3] tsngawpay [ˈt͡sŋaw.paj] n.
[4] nìsung [nɪ.ˈsuŋ] adv.
[5] säseyto [sæ.sɛj.ˈto] n.
[6] tìnvi [ˈtɪn.vi] n.
[7] kalin [ka.ˈlin] adj.
[8] mìn [m·ɪn] m<1><2><3>ìn vin.
[9] kxa [k'a] n.
[10] 'ewll [ˈʔɛ.wḷ] n.

~~> /add 1 3 4 7 8 9 10

~~> /define 2 5 6
[2] lopx [l·op'] l<1><2><3>opx vin. panic
[5] säseyto [sæ.sɛj.ˈto] n. butchering tool
[6] tìnvi [ˈtɪn.vi] n. task, errand, step (in an instruction)

~~> /source 6
[6] tìnvi: http://naviteri.org/2016/06/mrrvola-lifyavi-amip-forty-new-expression
s/ (30 Jun 2016)

~~> /source 2 5
[2] lopx: Frommer (05 Jul 2012) http://naviteri.org/2012/07/meetings-waterfalls-
and-more/
[5] säseyto: http://naviteri.org/2015/08/aylifyavi-lereyfya-2-cultural-terms-2-a
nd-more/ (30 Aug 2015)

~~> /add 2 5

~~> /delete 2 5

~~> /progress
6.21% (145 / 2334)

~~> /exit
data saved
```

```
➜ ./ftia
Ftia v2.0.0-dev by Tirea Aean
data loaded

~~> /latem

~~> /ftxey

~~> /ftxey 5
[1] n. plate, (for food)
[2] n. space, open or borderless area
[3] n. fern
[4] intj. hello
[5] n. exception

~~> /sung 3 4

~~> /ralpeng 1 2 5
[1] yomyo [ˈjom.jo] n. plate, (for food)
[2] ngip [ŋip̚] n. space, open or borderless area
[5] tìmungwrr [tɪ.muŋ.ˈwṛ] n. exception

~~> /polpxay
0.30% (7 / 2334)

~~> /hum
data saved
```