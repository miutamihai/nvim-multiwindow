# Nvim Multiwindow
Open current file in new window, directly within neovim!

## Installation 
Using Lazy:
```lua
{
  "miutamihai/nvim-multiwindow"
}
```

## Usage
Hit `<leader> + N` and you'll see the current file open in a new window

## Support
In order to open a new window in the _same_ neovim "client" (i.e. terminal or GUI application), `nvim-multiwindow`
has to lookup the process hierarchy to find the appropriate client to run the 
new window in. Because of that, clients have to be manually handled. The currently
implemented clients are:
- [x] [Wezterm](https://wezfurlong.org/wezterm/index.html)
- [x] [Neovide](https://neovide.dev)
