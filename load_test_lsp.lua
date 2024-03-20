local client = vim.lsp.start_client {
  name = "darkdownlsp",
  cmd = { "/Users/dhruvdabhi/Developer/projects/darkdownlsp/darkdownlsp" }
}

if not client then
  vim.notify("Do the client thing right")
  return
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = "markdown",
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end,
})
