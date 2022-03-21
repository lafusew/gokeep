# Gokeep

Gokeep is a simple implementation of a CLI password manager, written in Go.

## ⚠️ Warning ⚠️

Currently the app isn't safe. pwds arent encrypted before storage. 
This is a naive implementation of the solution and my first application I write in Go (it's also the first application running outside a Javascript Engine).
Use it for testing purpose only (this might change if I'm able to make a first release).

If you notice any security issue or bad pratices, please create an issue pointing me the correct/safer solution. 

## Behavior

`gokeep init` setup the local sqlite database and trigger the master key management.

The encryption key should be safely stored by the user. If you forget it, your credentials wont be recoverable as gokeep doesn't store your encryption key.

`gokeep start` starts the application. A prompt asking for your password will be dispayed and the app should continue to run after you provided it. 

Once the app is started, we should be able to run command through promptui. 

Credentials cmd list: 

- create
- read
- (list) readAll
- update
- delete

Application cmd list:

- stop
- key (Master Key management)
- gui (tbd)

You can also use the app without letting it open. But you'll have to enter your masterkey for each command.

## Todo 

- [x] Be able to keep the app up and running ⚙️
- [ ] Crypto 🔐
  - [x] Encryption / Decryption 🔐
    - [ ] Encryption error only on update to be fixed 
  - [ ] Key management 🔑
- [ ] Add handy flags to all commands (flags tbd) ⛳️
- [ ] Static web GUI served on a localhost server 💅
- [ ] Refactoring 🏭
  - [ ] Refactor cmd and methods associated
  - [ ] Refactor Error handling
- [ ] Generate random and safe pwd ⁉️

## Security concern

I don't know anything about security 🤕, so please don't use this application with your real credentials until someone who know what he's doing read my code. 

Currently the sqlite db isn't protected at all. I'm just AES encrypting right before SQL Insert and decrypting right after SQL Select. 

As it will be local only, I really don't know the required security level and good pratices for this kind of local application, if you do please create an issue and tell me. 🙏
