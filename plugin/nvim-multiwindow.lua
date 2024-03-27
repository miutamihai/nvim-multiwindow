function os.capture(cmd, raw)
  local handle = assert(io.popen(cmd, 'r'))
  local output = assert(handle:read('*a'))

  handle:close()

  if raw then
    return output
  end

  output = string.gsub(
    string.gsub(
      string.gsub(output, '^%s+', ''),
      '%s+$',
      ''
    ),
    '[\n\r]+',
    ' '
  )

  return output
end

vim.keymap.set('n', '<leader>N', function()
  local cwd = vim.fn['getcwd']()
  local output = os.capture(string.format("%s/nvim-multiwindow", cwd))

  print(output)
end, { noremap = true, silent = true, desc = "[N]vimMultiwindow: New Window" })
