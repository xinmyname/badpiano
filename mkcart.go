package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Writing badpiano.p8...")

	p8, err := os.Create("badpiano.p8")

	if err != nil {
		log.Fatal(err)
	}

	defer p8.Close()

	p8.WriteString("pico-8 cartridge // http://www.pico-8.com\n")
	p8.WriteString("version 41\n")
	p8.WriteString("__lua__\n")
	writeFileAsText("main.lua", p8)
	p8.WriteString("-->8\n")
	writeFileAsText("audio.lua", p8)
	p8.WriteString("-->8\n")
	writeFileAsText("keys.lua", p8)
	p8.WriteString("-->8\n")
	p8.WriteString("--data\n")
	p8.WriteString("data={}\n")
	p8.WriteString("data.samples={}\n")
	writeSampleData("Samples/epiano.u8", "data.samples.epiano", p8)
	writeSampleData("Samples/harpsi.u8", "data.samples.harpsi", p8)
	writeSampleData("Samples/jazzy.u8", "data.samples.jazzy", p8)
	writeSampleData("Samples/midc.u8", "data.samples.midc", p8)
	writeSampleData("Samples/organ.u8", "data.samples.organ", p8)
	writeSampleData("Samples/piano.u8", "data.samples.piano", p8)
	p8.WriteString("__gfx__\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111102222222220333333333304444444440555555555500666666666607777777770888888888809999999990aaaaaaaaaa0bbbbbbbbb0cccccccccc\n")
	p8.WriteString("1111111111100000000000333333333300000000000555555555500666666666600000000000888888888800000000000aaaaaaaaaa00000000000cccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("11111111111111110033333333333333333005555555555555555006666666666666666600888888888888888800aaaaaaaaaaaaaaaaa00ccccccccccccccccc\n")
	p8.WriteString("__map__\n")
	p8.WriteString("000102030405060708090a0b0c0d0e0f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\n")
	p8.WriteString("101112131415161718191a1b1c1d1e1f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\n")
	p8.WriteString("202122232425262728292a2b2c2d2e2f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\n")
	p8.WriteString("303132333435363738393a3b3c3d3e3f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\n")
	p8.WriteString("404142434445464748494a4b4c4d4e4f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\n")
	p8.WriteString("505152535455565758595a5b5c5d5e5f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\n")
	p8.WriteString("606162636465666768696a6b6c6d6e6f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\n")
	p8.WriteString("707172737475767778797a7b7c7d7e7f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\n")

	fmt.Println("OK.")
}

type writer interface {
	io.StringWriter
	io.Writer
}

func writeBytesAsText(bytes []byte, writer writer) {

	for i := 0; i < len(bytes); i += 1 {
		switch bytes[i] {
		case 0:
			if i < len(bytes)-1 {
				nextCh := bytes[i+1]
				if nextCh >= 48 && nextCh <= 57 {
					writer.WriteString("\\000")
				} else {
					writer.WriteString("\\0")
				}
			} else {
				writer.WriteString("\\0")
			}
		case 10:
			writer.WriteString("\\n")
		case 13:
			writer.WriteString("\\r")
		case 92:
			writer.WriteString("\\\\")
		case 34:
			writer.WriteString("\\\"")
		default:
			writer.Write(bytes[i : i+1])
		}
	}
}

func writeFileAsText(filename string, writer writer) error {

	bytes, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	_, err = writer.Write(bytes)

	if err != nil {
		return err
	}

	return nil
}

func writeSampleData(filename string, sampleName string, writer writer) {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err.Error())
	}

	writer.WriteString(sampleName)

	writer.WriteString("=\"")
	writeBytesAsText(bytes, writer)
	writer.WriteString("\"\n")
}
