package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"./bundler"
	"./config"
)

var lubuLogo = `																						
	LLLLLLLLLLL                               BBBBBBBBBBBBBBBBB                     
	L:::::::::L                               B::::::::::::::::B                    
	L:::::::::L                               B::::::BBBBBB:::::B                   
	LL:::::::LL                               BB:::::B     B:::::B                  
	L:::::L               uuuuuu    uuuuuu    B::::B     B:::::Buuuuuu    uuuuuu  
	L:::::L               u::::u    u::::u    B::::B     B:::::Bu::::u    u::::u  
	L:::::L               u::::u    u::::u    B::::BBBBBB:::::B u::::u    u::::u  
	L:::::L               u::::u    u::::u    B:::::::::::::BB  u::::u    u::::u  
	L:::::L               u::::u    u::::u    B::::BBBBBB:::::B u::::u    u::::u  
	L:::::L               u::::u    u::::u    B::::B     B:::::Bu::::u    u::::u  
	L:::::L               u::::u    u::::u    B::::B     B:::::Bu::::u    u::::u  
	L:::::L         LLLLLLu:::::uuuu:::::u    B::::B     B:::::Bu:::::uuuu:::::u  
	LL:::::::LLLLLLLLL:::::Lu:::::::::::::::uuBB:::::BBBBBB::::::Bu:::::::::::::::uu
	L::::::::::::::::::::::L u:::::::::::::::uB:::::::::::::::::B  u:::::::::::::::u
	L::::::::::::::::::::::L  uu::::::::uu:::uB::::::::::::::::B    uu::::::::uu:::u
	LLLLLLLLLLLLLLLLLLLLLLLL    uuuuuuuu  uuuuBBBBBBBBBBBBBBBBB       uuuuuuuu  uuuu
`

func main() {
	fmt.Println(lubuLogo)
	dir := filepath.Dir(os.Args[1])
	config := config.ReadConfig(dir + "\\")
	err := os.MkdirAll(strings.Join(strings.Split(config.Output, "/")[:1], "/"), 0755)
	if err != nil {
		panic(fmt.Sprintf("Error creating output directory: %s", err.Error()))
	}

	err = os.WriteFile(dir+"/"+config.Output, []byte(fmt.Sprintf("--[[\n%s\n\tBundled using LuBu - https://github.com/chaposcripts/lubu\n]]\n\n%s", lubuLogo, bundler.Generate(config, dir))), 0644)
	if err != nil {
		panic(fmt.Sprintf("Error creating output file: %s", err.Error()))
	}
	fmt.Println("[LuBu] Done, saved to", dir+"/"+config.Output)
}
