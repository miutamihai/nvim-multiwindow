local BIN_PATH = "~/.local/share/nvim/nvim-multiwindow"
local MAC_URL =
"https://gitlab.com/api/v4/projects/56250262/packages/generic/nvim-multiwindow/latest/nvim-multiwindow-darwin-arm64"
local LINUX_URL =
"https://gitlab.com/api/v4/projects/56250262/packages/generic/nvim-multiwindow/latest/nvim-multiwindow-linux-amd64"
local WINDOWS_URL =
"https://gitlab.com/api/v4/projects/56250262/packages/generic/nvim-multiwindow/latest/nvim-multiwindow-windows-amd64"

local function relative_to_absolute(relative_path)
  return vim.fn.fnamemodify(relative_path, ":p")
end

local function bin_exists()
  local f = io.open(relative_to_absolute(BIN_PATH), "rb")

  if f ~= nil then
    io.close(f)
    return true
  else
    return false
  end
end

local function ensure_bin()
  local exists = bin_exists()

  if not exists then
    local os_name = vim.loop.os_uname().sysname
    local url = MAC_URL

    if os_name == "Linux" then
      url = LINUX_URL
    elseif os_name == "Windows" then
      url = WINDOWS_URL
    end

    -- TODO: Remove dependency on curl
    os.execute(string.format("curl %s -o %s", url, BIN_PATH))
    os.execute(string.format("chmod 777 %s", BIN_PATH))
    return
  end
end

vim.keymap.set('n', '<leader>N', function()
  ensure_bin()
  local file = vim.fn.expand('%')
  os.execute(string.format("%s %s", BIN_PATH, file))
end, { noremap = true, silent = true, desc = "[N]vimMultiwindow: New Window" })
