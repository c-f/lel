package config

import (
	"fmt"
)

const (
	VERSION = "v0.0.1"
	STAGE   = "beta"
)

var (
	BANNER = fmt.Sprintf(`
      __        __            __
     /\__\     /\  \         /\__\
    /:/  /    /::\  \       /:/  /
   /:/  /    /:/\:\  \     /:/  / 
  /:/  /    /::\ \:\  \   /:/  /  
 /:/__/    /:/\:\_\:\__\ /:/__/   @iNvist
 \:\  \    \:\ \/_/\/__/ \:\  \   https://editor.l3l.lol
  \:\  \    \:\  \        \:\  \  version: %s
   \:\__\    \:\__\        \:\__\ 
    \/__/     \/__/         \/__/	

`, VERSION)
)
