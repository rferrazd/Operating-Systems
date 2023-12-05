// vim go.mod
// go mod init example.com/dinner
// binary file knows how to run the other files...
module example.com/dinner // giving main a name

replace example.com/dinner => ../dining-philosophers // module we are defining is all the files in this folder

go 1.21.1