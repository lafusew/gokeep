# Gokeep

Gokeep is a simple implementation of a CLI password manager, written in Go.

This is the first application I write in Go (it's also the first application running outside a Javascript Engine).

If you notice any security issue or bad pratices, please create an issue pointing me the correct/safer solution. 

## Behavior

`gokeep init` setup the local sqlite database and ask for an encryption password. 

The encryption key should be safely stored by the user. If you forget it, your credentials wont be recoverable as gokeep doesn't store your encryption key.

`gokeep` should start the application. A prompt asking for your password will be dispayed and the app should continue to run after you provided it. 

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

## Todo 

- [x] CRUD SQL Method mapped to /data public functions
- [ ] All prompt UI use case abstracted.
- [ ] Encryption / Decryption
- [ ] Keep the app up and running (currently shutting down after cmd execution)
- [ ] Static web GUI served on a localhost server