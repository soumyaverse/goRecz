package banner

import (
	"fmt"

	"github.com/burpOverflow/goRecz/pkg/colors"
)

func PrintBanner() {
	asciiArt :=
		colors.Yellow + `

		██████╗  ██████╗ ██████╗ ███████╗ ██████╗███████╗
		██╔════╝ ██╔═══██╗██╔══██╗██╔════╝██╔════╝╚══███╔╝
		██║  ███╗██║   ██║██████╔╝█████╗  ██║       ███╔╝ 
		██║   ██║██║   ██║██╔══██╗██╔══╝  ██║      ███╔╝  
		╚██████╔╝╚██████╔╝██║  ██║███████╗╚██████╗███████╗
		╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚══════╝ ╚═════╝╚══════╝
		` + colors.Reset + "(Created By: @burpOverflow)\n\t\tGithub: https://github.com/burpOverflow/goRecz\n"

	fmt.Println(asciiArt)
}
