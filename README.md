# ftia

A simple text-based vocabulary review app

## Installation


### macOS / Linux

Run the following inside a Terminal:

```
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

`/select x` : show untranslated entries for given number of random words, excluding words marked as known

`/selectfromall x`: show untranslated entries for given number of random words, including words marked as known

`/known n...` : mark the given entry as known so it won't show up in the future

`/define n...` : get the definition / translation of the given entry

`/quit`, `/exit`, `/q` : save data and quit the program.

`x` is a positive integer number from 1 to the number of words you wish to review with the command.

`n` is a positive integer number from 1 to the `x` 

`...` represents optional additional space-separated `n` in a list.

## Example

```
➜ ./ftia
Ftia v0.0.1-dev by Tirea Aean
data loaded
~~> /select 5
[1] tsenga [ˈt͡sɛ.ŋa]  conj.
[2] olo'eyktan [o.lo.ˈʔɛjk.tan]  n.
[3] toktor [ˈtok.toɾ]  n.
[4] lemweypey [lɛm.ˈwɛj.pɛj]  adj.
[5] fuke [fu.ˈkɛ]  conj.
~~> /known 1
~~> /known 2 3
~~> /define 4
adj. patient
~~> /define 5
conj. or not
~~> /select 10
[1] alu [ˈa.lu]  conj.
[2] snotipx [sno.ˈtip']  n.
[3] 'ewan [ˈʔɛ.wan]  adj.
[4] ioi [i.ˈo.i]  n.
[5] nìmweypey [nɪm.ˈwɛj.pɛj]  adv.
[6] fìtxan [fɪ.ˈt'an]  adv.
[7] vospxìtsìng [vo.sp'ɪ.ˈt͡sɪŋ]  n.
[8] sa'sem [ˈsaʔ.sɛm]  n.
[9] nguway [ˈŋu.waj]  n.
[10] pamtseo si [ˈpam.t͡sɛ.o ˈs·i] pamtseo s<1><2><3>i vin.
~~> /known 1 3 4 5 6 8 9 10
~~> /define 2 7
n. selfcontrol
n. April
~~> /exit
data saved
```