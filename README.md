ReleaseTools Golang version
=========

ReleaseTools in golang version. Binary compatible with:
 - Linux   amd64
 - MacOS   amd64
 - Windows amd64

# Installation
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FDALTCORE%2FReleaseTools-go.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FDALTCORE%2FReleaseTools-go?ref=badge_shield)


## MacOS
One-line install for lazy MacOS users:

```console
$ mkdir -p ~/bin \
  && curl -L -o ~/bin/release-tool-darwin-amd64 \
  https://github.com/DALTCORE/ReleaseTools-go/releases/download/1.1.4/release-tool-darwin-amd64 \
  && chmod +x ~/bin/release-tool-darwin-amd64 \
  && ln -sfv ~/bin/release-tool-darwin-amd64 ~/bin/rt
```

Make sure that you have pointed `~/bin` in your `$PATH`.

## Linux
One-line install for lazy Linux users:

```console
$ mkdir -p ~/bin \
  && curl -L -o ~/bin/release-tool-linux-amd64 \
  https://github.com/DALTCORE/ReleaseTools-go/releases/download/1.1.4/release-tool-linux-amd64 \
  && chmod +x ~/bin/release-tool-linux-amd64 \
  && ln -sfv ~/bin/release-tool-linux-amd64 ~/bin/rt
```

Make sure that you have pointed `~/bin` in your `$PATH`.

## Windows
1. Download [release-tool-windows-amd64.exe](https://github.com/DALTCORE/ReleaseTools-go/releases/download/1.1.4/release-tool-windows-amd64.exe)
and copy it to `C:\Users\<name>\bin` (create if directory does not exist yet!)  

2. Rename the file to `release-tool.exe` or `rt.exe` 

3. Open CMD as administrator and run:
```cmd
SETX path "%path%;%userprofile%\bin"
```
This migth take a while / hang. You can just close it after a couple of seconds.

4. Restart your terminal and run `release-tool help` or `rt help`

# Usage

## Commands
### Show help
```terminal
$ rt help
NAME:
   ReleaseTools - Releasing made easy

USAGE:
   release-tool [global options] command [command options] [arguments...]

VERSION:
   {{VERSION}}

COMMANDS:
     manager:prepare, mp    Prepare a new release
     manager:setup, ms      Setup a new Repo and Intothetest/accept environment
     manager:changelog, mc  Build all changelog entries to CHANGELOG.md
     init, i                Initialize a ReleaseTools environment
     status, s              Get ReleaseTools environment status
     playbook, p            Run a playbook
     changelog, c           Make a changelog entry
     help, h                Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version 
```

### Manager prepare
```terminal
$ release-tool manager:prepare (?1.0.0)
$ rt mp (?1.0.0)
```

Prepare a new release issue into Gitlab.

### Manager setup
```terminal
$ release-tool manager:setup
$ rt ms
```
 
Create a new repository in a group with the preferred permissions

### Manager changelog list
```terminal
$ release-tool manager:changelog:list
$ rt mcl
```

List all the changelogs in `changelog/unreleased` in a table.

### Manager changelog
```terminal
$ release-tool manager:changelog (?1.0.0)
$ rt mc (?1.0.0)
```

Build the CHANGELOG.md file from all the changelogs in `changelog/unreleased`.

### Init
```terminal
$ release-tool init
$ rt i
```

Initialize a fresh ReleaseTool environment for usage with projects.

### Status
```terminal
$ release-tool status
$ rt s
```

Display the status of the current ReleaseTool environment.

### Playbook
```terminal
$ release-tool playbook <playbook-name> (?1.0.0)
$ rt p <playbook-name> (?1.0.0)
```

<!-- Run a playbook based on playbook name with optional parameter for versions.

### Auto update
```terminal
$ release-tool auto-update (?--force)
$ rt au (?--force)
``` -->

Updates the ReleaseTool environment. 

### Changelog
```terminal
$ release-tool changelog (?--dry-run) (?--force)
$ rt c (?--dry-run) (?--force)
```

Create a changelog for the current merge request.

### Help
```terminal
$ release-tool help
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
      
  gitlab:                                                   # Select subject Gitlab for the playbook parser
    merge_request:                                          # Select method MergeRequest for the playbook parser
      title: "Release v:version to :to from :from"          # Set merge request title 
      from: develop                                         # Set the branch/ref/source where this MR is coming from
      to: master                                            # Set the name of the new branch
      
    make_branch:                                            # Select method MakeBranch for the playbook parser
      from: develop                                         # Set the branch/ref/source for the new branch
      to: v:version                                         # Set the new branch name
      
    create_tag:                                             # Select method CreateTag for the playbook parser
      from: master                                          # Set the branch/ref/source for the new branch
      to: v:version                                         # Set the new tag name.
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
The [`release-tool init`](#init) command creates a local `.release-tool` file in the directory you're calling it from.  
If you fill this document with the correct information, and remove the "repo" key then you can place it in 
you're home directory. ReleaseTool will merge the information from the home file to the local file on the fly

Local file example:
```yaml
repo: namespace/repo
```

Home file example:
```yaml
gitlab_url: #https://git.server.com (no slash at the end)
group: #namespace
api_url: #Gitlab.com/api/v4
api_key: #RandomKey
mattermost_webhook: #MattermostWebhookUrl
github_token: 
```
[Create new Github Token here](https://github.com/settings/tokens/new?scopes=repo&description=ReleaseTools-Go)  
This setup minimizes the risk of you committing and sharing your private keys and tokens.


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FDALTCORE%2FReleaseTools-go.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FDALTCORE%2FReleaseTools-go?ref=badge_large)
