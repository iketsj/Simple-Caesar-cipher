# Simple-Caesar-cipher
This is a simple terminal Caesar cipher that I wrote in Golang

## Example Use
Encode:
```
go run main.go -c "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" -t "Ike T. Sanglay Jr." -s 1 -m e
```
Decoding:
```
go run main.go -c "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" -t "Jlf U. Tbohmbz Ks." -s 1 -m d
```

## Output
If a character is not present in the character table. It will be retained as is.

Encode:
```
Mode is to encode with a shift value of 1.
Jlf U. Tbohmbz Ks.
```
Decode:
```
Mode is to decode with a shift value of 1.
Ike T. Sanglay Jr.
```

## Program Arguments
| Parameter		| Required		| Description		|
| ---		    | ---		    | ---		|
| -c		    | Yes		    | This will be the character lookup table	|
| -t		    | Yes		    | The text to be ciphered or deciphered		|
| -s		    | Yes		    | The shift value for the cipher	|
| -m	        | Yes		    | e or d for mode(encode or decode)		|
| -h	        | No		    | Show the help message		|