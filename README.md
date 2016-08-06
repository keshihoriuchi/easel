# easel

easel is a tool for write down your lean-canvas using markdown.

## Install

Please download a binary from [release](https://github.com/tanksuzuki/easel/releases).

Distributions are available for the...

* Windows
* Mac OS X(10.7-)
* Linux

## Usage

```
$ easel help

easel is a tool for write down your lean-canvas using markdown.
https://github.com/tanksuzuki/easel

Usage:
	easel command [args]

Commands:
	init        Generate markdown template
	watch       Run the preview server
	write       Convert the markdown to HTML

Use "easel help [command]" for more information about a command.
```

### Generate base markdown

```
$ easel init > sample.md
```

sample.md:

```
# PROBLEM



# CUSTOMER SEGMENTS



# UNIQUE VALUE PROPOSITION



# SOLUTION



# CHANNELS



# REVENUE STREAMS



# COST STRUCTURE



# KEY METRICS



# UNFAIR ADVANTAGE



```

### Write down your canvas using markdown

Write your lean-canvas to above markdown.

**NOTE**: Do not add/modify h1(`# heading`).

### Convert the canvas to HTML with style

```
$ easel write sample.md > sample.html
```

### Preview the canvas on your browser

```
$ easel watch sample.md
```

Please access `localhost:3000`.

Saving the markdown will reload the preview automatically.

## License

easel is licensed under the Apache License, Version 2.0.

## Author

[Asuka Suzuki](https://github.com/tanksuzuki)