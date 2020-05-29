# :recycle: Rolling Hash
Rolling Hash based file diffing

## How to run the project
- Required: ```Go 1.11+```
- Clone this director within your ```gopath```.
```
go install github.com/bariabbassi/rolling-hash
rolling-hash
```
- To run the tests use the command
```
go test
```

## What the project does
- The rolling hash algorithm is used for file diffing. To implement it I have created a data type chunker that rolls a chunk through a file. Rolling means changing the first and last byte of the chunk and incrementing the index. A chunker contains both the chunk and the file. The diffing function returns the indexes of the chunks that differ.
- I have used a very dumb hashing function where "abcd" hashes to 1+2+3+4=10. My plan was to change it later but I got lazy and left it as is.
- I have tested the code with different chunk sizes and files.


## Example
- The letter "t" was replaced by "@" in the alphabet. As a result, the 4 chunks 16, 17, 18 and 19 differ. The chunk size used here is 4.
```
16 qrst 74
16 qrs@ 22

17 rstu 78
17 rs@u 26

18 stuv 82
18 s@uv 30

19 tuvw 86
19 @uvw 34

[16 17 18 19]
```

