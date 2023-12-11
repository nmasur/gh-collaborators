# gh-environments

A GitHub `gh` [CLI](https://cli.github.com/) extension to list environments and their associated metadata for an organization and/or specific repositories. 

## Installation

1. Install the `gh` CLI - see the [installation](https://github.com/cli/cli#installation) instructions.

2. Install the extension:
   ```sh
   gh extension install katiem0/gh-collaborators
   ```

For more information: [`gh extension install`](https://cli.github.com/manual/gh_extension_install).

## Usage

The `gh-collaborators` extension supports `GitHub.com` and GitHub Enterprise Server, through the use of `--hostname` and the following commands:

```sh
$ gh collaborators -h
List guest collaborators and their repos.

Usage:
  collaborators [command]

Available Commands:
  list        Generate a report of repos guest collaborators have access to.

Flags:
  -h, --help   help for collaborators

Use "collaborators [command] --help" for more information about a command. 
```

### List Collaborators

Repository permissions assigned to a Repository Collaborator can be listed and written to a `csv` file for an organization or specific user.


```sh
$ gh collaborators list -h
Generate a report of repositories guest collaborators have access to.

Usage:
  collaborators list [flags] <organization>

Flags:
  -d, --debug                To debug logging
  -h, --help                 help for list
      --hostname string      GitHub Enterprise Server hostname (default "github.com")
  -o, --output-file string   Name of file to write CSV list to (default "RepoCollaboratorsReport-20231211162953.csv")
  -t, --token string         GitHub Personal Access Token (default "gh auth token")
  -u, --username string      Username of single repo collaborator to generate report for
```

The output `csv` file contains the following information:

| Field Name | Description |
|:-----------|:------------|
|`RepositoryName` | The name of the repository where the data is extracted from. |
|`RepositoryID`| The `ID` associated with the Repository, for API usage. |
|`Visibility`| The visibility of the repository. |
|`Username`| The username of the repository collaborator. |
|`AccessLevel`| The repository access permissions granted to the repository collaborator. |