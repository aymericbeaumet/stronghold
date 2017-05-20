# Specifications

## Goals

- Detect when illegitimate changes were made to the configuration.

- Restrict read and write accesses to sensible files

## Non-goals

- The **listing** of your encrypted files will always be visible to all the
  people which can access your vault. Stronghold only restricts which people
  can read or write those files.

## Permissions

- There are three permissions, which can be used to restrict the files:

  * forbidden (`---`)
  * _read-only_ (`r--`)
  * _read-write_ (`rw-`)

Note: the unix-style permissions are only here to help.

## Users

A user is the most granular entity to which you can attribute a permission. It
is denoted by a PGP key id.

## Groups

Groups allow to group users to configure their permissions in batch.

- A group can be empty
- A user can belong to several groups

### super-admin

The `root` group is special, its members:

- have read and write access to all the files stored in the vault
- are the only ones which can create, populate and delete the other groups

The legitimacy of the members in the root group is ensured by crawling up the chain
of signatures until the initialization of the Stronghold vault. This asserts no
one can be added by a third-party to the root group.

## Principles

Once a file is registered in the system, it is immediately encrypted with the
ser's private key, giving read and write access to this specific user only.

If inside the context of a VCS, the newly tracked file will be added to the
ignore file for your convenience. Note we don't do any commit for you though.
