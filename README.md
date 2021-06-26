# lottomatic
Generate the winning lotto numbers

## Download
Binaries for Windows and Linux can be downloaded [here](https://github.com/bwks/lottomatic/releases/latest).

## Usage
```
./lottomatic -h
Usage: ./lottomatic
  -M int
        amount of numbers in the main pool (default 35)
  -S int
        amount of numbers in the supplementary pool (default 20)
  -g int
        number of games to play (default 4)
  -m int
        amount of numbers to select from the main pool (default 7)
  -s int
        amount of numbers to select from the supplementary pool (default 1)
```

## Examples
With the defaults
```
./lottomatic
2  3  6  12 19 29 34 | 1 
2  17 18 21 26 29 33 | 8 
10 18 22 23 24 33 34 | 14 
10 15 17 23 26 32 33 | 2 
```

With `40` numbers in the main pool and `6` numbers in the draw. `0` supplementry or powerball numbers and `10` games
```
./lottomatic -M 40 -m 6 -S 0 -g 10
1  12 22 23 29 36 
5  24 27 31 34 39 
13 15 19 20 30 31 
5  8  12 14 19 27 
4  13 15 34 36 40 
8  9  10 28 36 37 
12 14 15 17 33 40 
8  13 28 29 32 39 
14 15 27 30 36 40 
2  13 22 29 34 39 
```
