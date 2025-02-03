Configure Global Environment Variables in GoLand:
PATH=/usr/bin:$PATH

export PATH="$PATH:/home/gwk/go/go1.23.5/bin"

cross compile example

reload
source ~/.bashrc


export CGO_ENABLED=1
export GOOS=windows
export GOARCH=amd64          # or 386 if you need 32-bit Windows
export CC=x86_64-w64-mingw32-gcc

which x86_64-w64-mingw32-gcc

go build -o myapp.exe