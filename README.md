# ESC Key value store

This is a plugin for [ESC](https://github.com/projesc/esc) that addes a sync key value store.

## Usage

Download from [release page](https://github.com/projesc/esc-kv/releases) and drop the file on ESC configured folder, it will load the plugin.

It will make availabe for the lua scripts:

```lua
set("foo","bar")
log("Foo "..get("foo"))
log("Fuz "..get("fuz"))
```

## License

MIT

