# lottomatic
Generate the winning lotto numbers

## Download
Binaries for Windows and Linux can be downloaded [here](https://github.com/bwks/lottomatic/releases/latest).

## Usage
```
./lottomatic -h
Usage: ./lottomatic
  -M int
        numbers in the main pool of numbers (default 35)
  -S int
        numbers in the pool of supplementary numbers (default 20)
  -g int
        number of games to play (default 4)
  -m int
        number of main numbers (default 7)
  -s int
        number of supplementary numbers (default 1)
```

```
./lottomatic -M 35 -m 7 -S 20 -s 1
2  3  6  12 19 29 34 | 1 
2  17 18 21 26 29 33 | 8 
10 18 22 23 24 33 34 | 14 
10 15 17 23 26 32 33 | 2 
```
