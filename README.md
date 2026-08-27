# regroup-cli

CLI client for the [Regroup](https://www.regroup.com/) Mass Notification API.

## Install

```
go install github.com/Folger-Shakespeare-Library/regroup-cli/cmd/regroup@latest
```

## Setup

```
regroup config init
```

Prompts for API key and secret. Credentials are stored in `~/.config/regroup/config.json` (mode 0600).

## Usage

```
regroup contacts list
regroup contacts list --email user@example.com
regroup contacts list --group technology
regroup contacts add --email user@example.com --group technology --first-name Jane --last-name Doe
regroup contacts remove --email user@example.com --group technology
regroup groups list
regroup groups contacts <group-slug>
regroup channels list
regroup channels contacts <channel-slug>
regroup config show
regroup config path
```

## Output

Default output is JSON. Use `--output table` for tabular output. Pipe to `jq` for filtering:

```
regroup contacts list | jq '.[] | select(.groups | index("Technology"))'
```

Exit code 0 on success, no output. Non-zero on error, message to stderr.

## Flags

Groups and channels are specified by slug (lowercase, hyphens for spaces). Use `regroup groups list` or `regroup channels list` to see available slugs.

`--group` accepts repeated flags or comma-separated values:

```
regroup contacts add --email user@example.com --group technology --group staff
regroup contacts add --email user@example.com --group technology,staff
```

