# Gokeep

Gokeep is a simple implementation of a CLI password manager, written in Go.

## ⚠️ Warning ⚠️

Currently the app isn't safe. pwds arent encrypted before storage. 
This is a naive implementation of the solution and my first application I write in Go (it's also the first application running outside a Javascript Engine).
Use it for testing purpose only (this might change if I'm able to make a first release).

If you notice any security issue or bad pratices, please create an issue pointing me the correct/safer solution. 

## Behavior

`gokeep init` setup the local sqlite database and ask for an encryption password. 

The encryption key should be safely stored by the user. If you forget it, your credentials wont be recoverable as gokeep doesn't store your encryption key.

`gokeep start` starts the application. A prompt asking for your password will be dispayed and the app should continue to run after you provided it. 

Once the app is started, we should be able to run command through promptui. 

Credentials cmd list: 

- new
- delete
- find
- findAll
- update

Application cmd list:

- stop
- forget (remove encryption key from memory)

You'll also be able to use the 1 command at the time (requiring your master keyworld for each cmd).

## Todo 

- [x] Be able to keep the app up and running
- [ ] Encryption / Decryption 
- [ ] Add handy flags to all commands (flags tbd)
- [ ] Refactor cmd and their assiociated method
- [ ] Static web GUI served on a localhost server