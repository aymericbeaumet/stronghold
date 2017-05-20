# Stronghold

Secretly store your secrets, wherever, whenever.

## Development

To run the program:

```bash
go run main.go
```

Whenever the dependencies are updated:

```bash
dep ensure --update
```

## Goals

- simple and meaningful CLI
- GPG1 and GPG2 can be interchanged at anytime
- it has been carefully crafted to not get in the way of VCS

## Commands

### Setup

- [ ] `stronghold init` # Initialize in the current directory, or at the VCS root (if any)
- [ ] `stronghold deinit` # Decrypt all files you can access, delete the .stronghold folder

### Authenticating the current user

The first user to authenticate will also see its key added to the root group,
this group has full read/write access to the secrets and the configuration.

- [ ] `stronghold auth email1`
- [ ] `stronghold deauth email1`

### Groups

- [ ] `stronghold group create groupName`
- [ ] `stronghold group delete groupName`
- [ ] `stronghold group add groupName email1 (keyring|mit|pgp|sks):email2 keybase:username`
- [ ] `stronghold group remove groupName email1`
- [ ] `stronghold group list groupName [--files]`

### Files

- [ ] `stronghold [file] add file/path/1 directory/1 [--group group1:rw]`
- [ ] `stronghold [file] remove file/path/1`
- [ ] `stronghold [file] list [--groups]`
- [ ] `stronghold [file] restrict file/path/1 [--group group1:rw]`
- [ ] `stronghold [file] encrypt [file/path/1] [--all]`
- [ ] `stronghold [file] decrypt [file/path/1] [--all]`
```

## License

[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png)](http://creativecommons.org/publicdomain/zero/1.0/)

To the extent possible under law, [Aymeric Beaumet](https://aymericbeaumet.com)
has waived all copyright and related or neighboring rights to this work.
