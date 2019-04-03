# ftia

A simple text-based vocabulary review app

## Installation


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

## Usage

`/select x` : select x random unlearned words

`/selectfromall x`: select x random words both learned and unlearned

`/known x...` : select x random learned words

`/define n...` : show definition / translation for given entry in selection

`/source n...` : show canon source of given entry

`/add n...` : mark given entries known / learned

`/delete n...` : unmark given entries known / learned

`/progress` : show current progress of words learned out of words in the dictionary

`/quit`, `/exit`, `/q` : save data and quit the program.

`x` is a positive integer number from 1 to the number of words you wish to review with the command.

`n` is a positive integer number from 1 to the `x`

`...` represents optional additional space-separated `n` in a list.

## Example

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