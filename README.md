<div align="center">
	<img src="assets/img/cedar.png" width="100" height="100">
	<h1>Cedar</h1>
	<p>
		<b>A personal log, directly from your command line</b>
	</p>
	<br>
	<br>
</div>

Cedar is a command line tool that allows you to easily capture a personal log while you're working. Logs captured via the CLI are appended to a daily .txt file that can be viewed anytime. Connect your GitHub account to automatically log your daily .txt files.

## Installation

Download binary: ![cedar](https://github.com/robertjkeck2/cedar)

## Usage

### Simple personal logging
Cedar allows you to quickly capture logs through the command line. Simply use the CLI to type a log entry and the entry will automatically save to the daily .txt file.

#### Basic log entry
Cedar is all about simplicity. Type any log message you'd like after the `cedar` command and it will be logged with the current time.

```
cedar This is a basic entry. No " or other special characters needed.
```

#### Read today's log
See the current day's log entries by simply typing the `cedar` command and hitting Enter.

```
cedar
```

If you need to see previous day's entries, use the synced GitHub repo or your favorite search tool on the `~/.cedar` directory.

### Automatic syncing
Daily .txt files are stored in the `~/.cedar` directory by default. Cedar also allows for automatic syncing with a GitHub account so .txt files can be publically accessible. 

#### Setting up GitHub syncing
Create a new repository (either private or public will do) on GitHub and copy the .git URL.

```
cedar <repo-url>
```

All entries will be committed directly to `master` each time you log with `cedar`.

## Maintainers

- [John Keck](https://github.com/robertjkeck2)

## License

MIT
