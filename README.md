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

```
go get github.com/cydy/EmojiFarm
go run .
```

Extra small release with [tinygo](https://tinygo.org/) (reduce file size by ~80%, 1.4MB -> 295KB)

```
tinygo build -no-debug -panic=trap -scheduler=none -gc=leaking -opt=s -o emojifarm .
```

Options
```
-seed string
        Seed for generation. 
        - Empty: Creates a unique grid each time (default)
        - Any text: Creates reproducible grids based on the text
        - Examples: "-seed hello" or "-seed 12345" or "-seed my-favorite-grid"
```


## Credit
Thanks to [Joe Sandow](https://github.com/joesondow) for his emoji twitter bots which inspired me to work on this.
- https://twitter.com/EmojiAquarium
- https://twitter.com/EmojiMeadow
