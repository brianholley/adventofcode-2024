package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"strconv"
)

type File struct {
	index  int
	length int
	offset int
}

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func readFilesOnDisk(stdin *bufio.Scanner) []File {
	files := []File{}
	offset := 0
	index := 0
	for stdin.Scan() {
		line := stdin.Text()

		for i := range line {
			if i%2 == 0 {
				length, _ := strconv.Atoi(string(line[i]))
				file := File{index: index, length: length, offset: offset}
				files = append(files, file)
				offset += length
				index++
			} else {
				length, _ := strconv.Atoi(string(line[i]))
				offset += length
			}
		}
	}
	return files
}

func part1(stdin *bufio.Scanner) string {
	result := 0

	files := readFilesOnDisk(stdin)

	for i := 0; len(files) > 0; i++ {
		// In active file
		if i >= files[0].offset && i < files[0].offset+files[0].length {
			result += i * files[0].index

			if i == files[0].offset+files[0].length-1 {
				// fmt.Println("Finished with file", files[0])
				files = files[1:]
			}
		} else {
			if len(files) > 0 && i == files[0].offset {
				result += i * files[0].index
			} else if len(files) > 0 {
				files[len(files)-1].length--
				result += i * files[len(files)-1].index
				if files[len(files)-1].length == 0 {
					files = files[:len(files)-1]
				}
			}
		}
	}

	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0

	files := readFilesOnDisk(stdin)

	freeList := []File{}
	for i := 1; i < len(files); i++ {
		free := files[i].offset - (files[i-1].offset + files[i-1].length)
		freeList = append(freeList, File{offset: files[i-1].offset + files[i-1].length, length: free})
	}

	for i := len(files) - 1; i >= 0; i-- {
		for f := range freeList {
			if freeList[f].length >= files[i].length && freeList[f].offset < files[i].offset {
				files[i].offset = freeList[f].offset
				freeList[f].offset += files[i].length
				freeList[f].length -= files[i].length
				break
			}
		}
	}

	for i := range files {
		for b := 0; b < files[i].length; b++ {
			result += (files[i].offset + b) * files[i].index
		}
	}

	return fmt.Sprint(result)
}
