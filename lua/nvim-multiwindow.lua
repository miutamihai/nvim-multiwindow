vim.keymap.set('n', '<leader>N', function()
  os.execute("./nvim-multiwindow")
end, { noremap = true, silent = true, desc = "[N]vimMultiwindow: New Window" })
