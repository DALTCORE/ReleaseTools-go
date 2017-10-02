ReleaseTools Golang version
=========

ReleaseTools in golang version. Binary compatible with:
 - Linux   amd64
 - MacOS   amd64
 - Windows amd64

# Installation

## MacOS
One-line install for lazy MacOS users:

```console
$ mkdir -p ~/bin \
  && curl -L -o ~/bin/release-tool-darwin-amd64 \
  https://github.com/DALTCORE/ReleaseTools-go/releases/download/1.0.5/release-tool-darwin-amd64 \
  && chmod +x ~/bin/release-tool-darwin-amd64 \
  && ln -sfv ~/bin/release-tool-darwin-amd64 ~/bin/rt
```

Make sure that you have pointed `~/bin` in your `$PATH`.

## Linux
One-line install for lazy Linux users:

```console
$ mkdir -p ~/bin \
  && curl -L -o ~/bin/release-tool-linux-amd64 \
  https://github.com/DALTCORE/ReleaseTools-go/releases/download/1.0.5/release-tool-linux-amd64 \
  && chmod +x ~/bin/release-tool-linux-amd64 \
  && ln -sfv ~/bin/release-tool-linux-amd64 ~/bin/rt
```

Make sure that you have pointed `~/bin` in your `$PATH`.

## Windows
:information_desk_person: dunno lol! 

# Usage

## Commands
### Show help
```terminal
$ rt help
NAME:
   ReleaseTools - Releasing made easy

USAGE:
   assets.exe [global options] command [command options] [arguments...]

VERSION:
   {{VERSION}}

COMMANDS:
     manager:prepare, mp    Prepare a new release
     manager:setup, ms      Setup a new Repo and Intothetest/accept environment
     manager:changelog, mc  Build all changelog entries to CHANGELOG.md
     init, i                Initialize a ReleaseTools environment
     status, s              Get ReleaseTools environment status
     playbook, p            Run a playbook
     auto-update, au        Auto update the Release Tools
     changelog, c           Make a changelog entry
     help, h                Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version 
```

### Manager prepare
```terminal
$ rt mp 1.0.0
```

Prepare a new release issue into Gitlab.

### Manager setup
```terminal
 $ rt ms
```
 
Create a new repository in a group with the preferred permissions

### Manager changelog
```terminal
 $ rt mc 1.0.0
```

Build the CHANGELOG.md file from all the changelogs in `changelog/unreleased`.

### Init
```terminal
 $ rt i
```

Initialize a fresh ReleaseTool environment for usage with projects.

### Status
```terminal
 $ rt s
```

Display the status of the current ReleaseTool environment.

### Playbook
```terminal
 $ rt p <playbook-name> (?1.0.0)
```

Run a playbook based on playbook name with optional parameter for versions.

### Auto update
```terminal
 $ rt au (?--force)
```

Updates the ReleaseTool environment. 

### Changelog
```terminal
 $ rt c (?--dry-run) (?--force)
```

Create a changelog for the current merge request.

### Help
```terminal
 $ rt help
```

Shows help for ReleaseTool

## Playbooks
Playbooks are just simple scripts written in the YAML language to be parsed by the ReleaseTool
```yaml
playbook:                                                   # Definition for ReleaseTool 
  version: 1.1                                              # Select Playbook Parser version

  mattermost:                                               # Select subject Mattermost for the playbook parser
    notify:                                                 # Select method Notify for the playbook parser
      channel: Townsquare                                   # Set Mattermost channel
      message: "Foobar with :url and :version and :repo"    # Set Message to be send to the channel
      
  gitlab:
    merge_request:
      title: "Release v:version to :to from :from"
      from: develop
      to: master
      
    make_branch:
      from: develop
      to: v:version
      
    create_tag:
      from: master
      to: v:version
```

## Stubs
**prepare.stub**
```text
**Release `:repo` version `:version`**  
- [X] Create issue  
- [ ] Notify in Mattermost `rt p notify-upcoming-release :version`  
- [ ] Merge request *develop > releases/v:version* `rt p develop-to-release :version`  
- [ ] Checkout releases/v:version `git fetch --all; git checkout releases/v:version`  
- [ ] Generate changelog `rt mp :version`  
- [ ] Create merge request *releases/v:version > staging* `rt p release-to-staging :version`  
- [ ] Wait for merge request *releases/v:version > staging* to be merged  
- [ ] Merge request *staging > master* `rt p staging-to-master :version`  
- [ ] Wait for merge request *staging > master* to be merged  
- [ ] Create tag v:version `rt p create-tag :version`  
- [ ] Merge request *master > develop* `rt p master-backport`  
- [ ] Wait for merge request *master > develop* to be merged  
- [ ] Notify in Mattermost `rt p release-done`  
- [ ] Close issue  
```

## Useful information

### Global and local release tool information
The [`rt i`](#init) command creates a local `.release-tool` file in the directory you're calling it from.  
If you fill this document with the correct information, and remove the "repo" key then you can place it in 
you're home directory. ReleaseTool will merge the information from the home file to the local file on the fly

Local file example:
```yaml
repo: namespace/repo
```

Home file example:
```yaml
api_url: 
api_key: 
mattermost_webhook: 
github_token: 
```

This setup minimizes the risk of you committing and sharing your private keys and tokens.
