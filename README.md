# el
CLI wrapper application for Elhuyar Hiztegia

## English
The main reason why I am writing this application is because
whenever I want to look up a translation to basque, I have to use
a web browser and find the word in the
[Elhuyar translator](https://hiztegiak.elhuyar.eus/) page. This is
too much effort for me, so this application is intended to lighten that process.
Specifically, it fetches, saves and prints arbitrary entries from the command line.
And by the way, I am writing it to learn the Go programming language.

The application is a web scraper, as Elhuyar does not provide any public API
to use it's dictionaries.

### Basic usage
To build the application (`go` is required):
```
$ go build -o el cmd/el/main.go
```

To run the program:
```
$ ./el <word>
```
*eg.* to search word *esan*:
```
$ ./el esan
```

## Euskara
Aplikazio hau idazteko arrazoi nagusia honakoa da, euskarara itzulpen
bat bilatu nahi dudan bakoitzean, nabigatzailea erabiliz [Elhuyar hiztegian](https://hiztegiak.elhuyar.eus/)
bilatu behar dut hitza. Hau, lan gehiegi da niretzat, eta aplikazio honek
prozesu hori arintzeko balio du. Zehazki, bilaketa egin, gorde eta pantailaratzeko
balio du komando lerrotik. Horretaz gain, Go programazio lengoaia ikasteko egin dut.

Aplikazioa web arakatzaile bat da, izan ere Elhuyarrek ez du eskaintzen inolako API-rik
bere hiztegiak erabili ahal izateko.

### Erabilera basikoa
Aplikazioa konpilatzeko (`go` behar da):
```
$ go build -o el cmd/el/main.go
```

Programa exekutatzeko:
```
$ ./el <word>
```
Adibidez, *esan* hitza bilatzeko:
```
$ ./el esan
```
