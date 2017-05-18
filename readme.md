# Stronghold

Secretly store your secrets, wherever, whenever.

## Goals

- consistant CLI
- use a local GPG to guarantee consistency
- compatible with git-secret and blackbox

## Commands

```bash
# Initialize in the current directory, or at the VCS root (if any)
stronghold init
# Deinitialize
stronghold deinit
# Add a new user to the pool
stronghold allow hi@aymericbeaumet.com
stronghold disallow hi@aymericbeaumet.com
# Files can be added/removed like so
stronghold add file/path/1
stronghold rm file/path/1
# Files can be added/updated for a specific user with --only and --except
stronghold add file/path/1 --only hi@aymericbeaumet.com
stronghold update file/path/1 --except hi@aymericbeaumet.com
# List known files and their permissions
```

## Alternatives

- [blackbox](https://github.com/StackExchange/blackbox)
- [git-secret](https://github.com/sobolevn/git-secret)

## License

[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png)](http://creativecommons.org/publicdomain/zero/1.0/)

To the extent possible under law, [Aymeric Beaumet](https://aymericbeaumet.com)
has waived all copyright and related or neighboring rights to this work.
