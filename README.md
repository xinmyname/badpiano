# Bad Piano
A polyphonic piano with multiple instruments.
* Requires keyboard
* Pick from six instruments: epiano, harpsi, jazzy, midc, organ, piano
* midc is a sine wave at 261.6256Hz - "middle c"
* Press left/right arrows to select a new instrument
* Play notes with the following keys: 
```  
  [2] [3]     [5] [6] [7]
[Q] [W] [E] [R] [T] [Y] [U]
```
* Up to four notes can be played simultaneously

## Build Instructions
A Go script assembles a .p8 file from the Lua code and encodes the samples as strings. Execute `go run mkcart` to build the `badpiano.p8` file, then copy it into your cart folder.

## Caveats
* It doesn't sound all that great.
* It can't be compressed into a cart because the instruments are too large.

