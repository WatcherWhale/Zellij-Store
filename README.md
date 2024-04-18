# Zellij Store

A way to share environment variables between Zellij Sessions

## Usage

1. In your config.fish add:

```fish
# config.fish

zellij-store | source
```

2. To add environment variables

```console
$ zellij-store store <var1> <var2> ... <varN>
```

3. To clean stores

```console
$ zellij-store clean
```
