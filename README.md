<div align="center">
	<img src="cedar.png" width="100" height="100">
	<h1>Cedar</h1>
	<p>
		<b>Log personal notes directly from your command line</b>
	</p>
	<br>
	<br>
</div>

Cedar is a command line tool that allows you to easily capture personal notes while you're working. Notes captured via the CLI are appended to a daily .txt file that can be stored anywhere. Connect your GitHub account to automatically log your daily .txt files.

## Installation

Download binary: ![cedar](https://github.com/robertjkeck2/cedar)

## Usage

### Simple note logging
Cedar allows you to quickly capture notes through the command line. Simply use the CLI to type a note with any desired flags and the note will automatically save to the daily .txt file.

#### Basic Note
```
cedar "This is a basic note."
```

### Automatic syncing
Daily .txt files are stored in the `~/.cedar` directory by default. Cedar also allows for automatic syncing with a GitHub account so .txt files can be publically accessible. 

#### Setting up GitHub access
```
cedar login
```

#### Setting up GitHub syncing
```
cedar sync
```

## Maintainers

- [John Keck](https://github.com/robertjkeck2)

## License

MIT
