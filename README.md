## Lve-server

CLI server for serve static files in folder

## Download

[Binary](https://github.com/ivnvMkhl/live-server/tree/master/build_bin)

## Usage

Recomanded use the alias on you terminal, `.zshrc` for example

```
alias server="~/live_server_mac_arm64"
```

Base usage on current dir

```bash
server

# Output: Starting live on port: 8080 in path: / ...
```

for help use `-h` flag

```bash
server -h

# Output: Usage of /live-server:
#  -log
#    	Logging all requests
#  -p string
#    	Server startup port (shorthand) (default "8080")
#  -port string
#    	Server startup port (default "8080")
#  -spa
#    	Use server for SPA. Server any route request returned ./index.html
#  -spa-entry string
#    	Path to SPA entry file. Default: /index.html (default "/index.html")
#  -src string
#    	Relative path to files
#  -watch
#  	  Watch mode for listen modified files in serve path
```
