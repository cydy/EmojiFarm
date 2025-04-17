# EmojiFarm

Generates small farms out of emojis, formatted to fit in a tweet.

Can be seen on 
- https://bsky.app/profile/emojifarm.bsky.social
- https://twitter.com/EmojiFarm

## Examples

🐥🦃🌱🌱🌼🌱🌺🌺🌼🌺🌱🌼🌱    
🌱🌱🦃🦃🌱🌺🌺🌱🌱🌹🌹🌹🌹    
🐣🐥🦃🐥🌸🌱🌹🌱🌹🌼🌱🌹🌼    
🌱🌱🌱🌱🌱🌼🌺🌺🌼🌱🌱🌱🌱    
🌱🦃🌱🌱🌸🌺🌱🌱🌹🌸🌺🌸🌸    
🦃🐥🌱🌱🌺🌹🌹🌸🌼🌼🌼🌹🌱    
⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓    
🌱🌱🌱🌱🌱🐑🐑🐑🌱🐑🌱🐑🌱    
🌱🌱🌱🐥🌱🌱🌱🌱🐑🐑🌱🐏🌱    
🌱🥚🌱🥚🐑🌱🐑🐑🐑🌱🌱🐏🌱    

🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈    
🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈    
🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈    
🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈🍈    
🍑🍑🍑🍑🍑🍑🌲🌹🌱🌱💮🌹🌱    
🍑🍑🍑🍑🍑🍑🌲🌻💮🌻🌱🌻🌱    
🍑🍑🍑🍑🍑🍑🌲🌱💮🌻🌱🌻🌱    
🍑🍑🍑🍑🍑🍑🌲🌱🌱🌹🌹🌹🌼    
🍑🍑🍑🍑🍑🍑🌲🌸🌸🌼🌹🌱💮    
🍑🍑🍑🍑🍑🍑🌲🌹🌱💮🌱🌹🌱    
 
🌊🌊🐚🐟🌊🌊🌊🐟☘️☘️🍃🍃🌿    
🌊🌊🐚🌊🐟🌊🐚🌊🌿🍃🍃🌱🌱    
🌊🐚🌊🐚🌊🌊🌊🌊🌱🌿☘️🍃☘️    
🐚🌊🐚🌊🌊🌊🐟🐚🌱☘️☘️🍃🌱    
🌲🌲🌲🌲🌲🌲🌲🌲🌲🏔🏔🏔🏔    
🍏🍏🍏🍏🍏🍏🍏🍏🍏🍇🍇🍇🍇    
🍏🍏🍏🍏🍏🍏🍏🍏🍏🍇🍇🍇🍇    
🍏🍏🍏🍏🍏🍏🍏🍏🍏🍇🍇🍇🍇    
🍏🍏🍏🍏🍏🍏🍏🍏🍏🍇🍇🍇🍇    
🍏🍏🍏🍏🍏🍏🍏🍏🍏🍇🍇🍇🍇    

🌱🐏🌱🌱🌲🐖🐃🌱🌱🐃🐖🌱🌱    
🌱🌱🐏🐏🌲🌱🐄🐎🌱🌱🐖🐎🌱    
🌱🐏🌱🌱🌲🐄🐖🐑🌱🌱🐄🌱🌱    
🌱🐏🐏🌱🌲🐎🌱🐄🌱🐖🐄🐑🌱    
🐏🐏🐏🐏🌲🌱🌱🐑🐎🌱🌱🌱🐑    
🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲    
🐣🥚🥚🐣🦃🌱🦆🌱🐣🦆🐓🐣🐓    
🌱🦃🌱🌱🌱🦆🦆🌱🦃🐓🌱🌱🦆    
🦃🐣🦆🦆🌱🌱🦃🌱🌱🌱🌱🌱🐣    
🥚🦃🐣🥚🌱🐣🌱🌱🥚🌱🌱🐣🌱    

## Instructions

### Command Line Usage
```
# Clone the repository
git clone https://github.com/cydy/EmojiFarm.git
cd EmojiFarm

# Run the program
go run .
```

#### Method 2: Install and Run
```
# Install the program
go install github.com/cydy/EmojiFarm@latest

# Run the program
emojifarm
```

Extra small release with [tinygo](https://tinygo.org/) (reduce file size by ~80%, 1.4MB -> 295KB)
```
tinygo build -no-debug -panic=trap -scheduler=none -gc=leaking -opt=s -o emojifarm .
./emojifarm
```

### WebAssembly Usage
To use the WebAssembly version:
1. Download the following files from the latest release:
   - `emojifarm-wasm.wasm`
   - `wasm_exec.js`
   - `wasm.html` from repo
2. Place all files in the same directory
3. Serve the files using a local web server (required due to CORS restrictions)
   ex. `python3 -m http.server 8000`
4. Open `http://localhost:8000/wasm.html` in your web browser

### Command Line Options
```
-seed string
        Seed for generation. The seed can be provided in two ways:
        - As a flag: "-seed hello" or "-seed 12345"
        - As the first positional argument: "hello" or "12345"
        - Empty: Creates a unique grid each time (default)
        Any text will create a reproducible grid based on the text.

-v, -verbose
        Enable verbose output, showing additional information about the generation process
```

## Credit
Thanks to [Joe Sandow](https://github.com/joesondow) for his emoji twitter bots which inspired me to work on this.
- https://twitter.com/EmojiAquarium
- https://twitter.com/EmojiMeadow