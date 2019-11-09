package config

import (
	"fmt"
)

const (
    // Current Version of LeL
    VERSION = "v0.0.1"
    // Current stage of the development
	STAGE   = "beta"
)

var (
    // BANNER is the LeL Banner 
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
