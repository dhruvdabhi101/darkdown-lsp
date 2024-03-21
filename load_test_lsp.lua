local client = vim.lsp.start_client {
  name = "darkdownlsp",
  cmd = { "path to the binary of lsp" }
}

if not client then
  vim.notify("Do the client thing right")
  return
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = "*",
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end,
})

vim.api.nvim_create_autocmd("LspAttach", {
  callback = function(args)
    print("Attach", vim.inspect(args))
  end,
})


