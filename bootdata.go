package main

import (
	"fmt"
	"os"
	"unicode"
	"strconv"
)

var data [512]byte

func ReadWrite(file string, dest string) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println("bootdata: could not open '" + file + "':", err)
		os.Exit(1)
	}
	defer f.Close()
	
	n, err := f.ReadAt(data[0:512], 0)
	if err != nil {
		if err.Error() != "EOF" {
			fmt.Println("bootdata: could not read '" + file + "'", err)
			os.Exit(1)
		}
	}

	fmt.Println("\033[93mBootloader data in \"" + file + "\"\033[0m\n")

	if dest == "" {
		fmt.Println("Byte:")
		fmt.Printf("\033[93m Column \033[0m ")
		for i := 1; i <= 32; i++ {
			fmt.Printf("\033[33m")
			if i < 10 {
				fmt.Printf("0")
			}
			fmt.Printf(strconv.Itoa(i))
			fmt.Printf(" \033[0m")
		}
		fmt.Printf("\n")

		for i := 0; i < 16; i++ {
			fmt.Printf(" \033[33m0x%04X:\033[0m ", i * 0x20)
			for j := 0; j < 32; j++ {
				current := (i * 0x20) + j
				fmt.Printf("%02X ", data[current])	
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")

		fmt.Println("ANSI:")
		for i := 0; i < 16; i++ {
			fmt.Printf(" \033[33m0x%04X:\033[0m ", i * 0x20)
			for j := 0; j < 32; j++ {
				current := (i * 0x20) + j
				if unicode.IsPrint(rune(data[current])) && rune(data[current]) != 0x0A && rune(data[current]) != 0x0D {
					fmt.Printf(string(rune(data[current])))
				} else {
					fmt.Printf("\033[31m.\033[0m")
				}	
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
		
		fmt.Println("Bytes read:", n)
	} else {
		_f, err := os.OpenFile(dest, os.O_RDWR | os.O_SYNC, 0)
		if err != nil {
			fmt.Println("bootdata: could not open '" + dest + "':", err)
			os.Exit(1)
		}
		defer _f.Close()

		_n, err := _f.WriteAt(data[0:512], int64(0))
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("bootdata: could not write to '" + dest + "'", err)
				os.Exit(1)
			}
		}
		
		fmt.Println("Bytes written:", _n)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: bootdata <file/disk> [-w <file/disk>]")
		os.Exit(1)
	}

	file := ""
	dest := ""
	
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch arg {
		case "-w":
			if len(os.Args) < i + 2 {
				fmt.Println("Usage: bootdata <file/disk> [-w <file/disk>]")
				os.Exit(1)
			}
			dest = os.Args[i + 1]
			i++
		default:
			if file == "" {
				file = arg
			} else {
				fmt.Println("\033[31m!!\033[0m \033[33mUnknown argument '" + arg + "'\033[0m \033[31m!!\033[0m")
			}
		}
	}

	if file == "" {
		fmt.Println("Usage: bootdata <file/disk> [-w <file/disk>]")
		os.Exit(1)
	}	
	
	ReadWrite(file, dest)	
}
