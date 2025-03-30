## Lve-server

CLI server for serve static files in folder

## Download

[Binary](https://github.com/ivnvMkhl/live-server/tree/master/build)

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
#  -p string
#    	Port to run the server on (shorthand) (default "8080")
#  -port string
#    	Port to run the server on (default "8080")
#  -spa
#    	Use server for SPA. Server redirects any request to the ./index.html
#  -src string
#    	Relative path to files
```
