package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type Directory struct {
	name           string
	size           int
	parent         *Directory
	subdirectories []*Directory
	files          []*File
}

func build_filesystem(root *Directory, lines []string) {
	var reading_files bool
	var files []*File
	current_dir := root

	for _, line := range lines {
		if reading_files && strings.HasPrefix(line, "$") {
			for _, f := range files {
				current_dir.files = append(current_dir.files, f)
			}
			fmt.Println("Added files to dir " + current_dir.name + " " + string(len(current_dir.files)))
			files = nil
			reading_files = false
		}

		if line == "$ cd /" {
			fmt.Println("Moving to root directory", line)
			continue
		} else if line == "$ ls" {
			fmt.Println("Listing files")
			reading_files = true
			continue
		} else if line == "$ cd .." {
			fmt.Println("Moving to parent directory " + current_dir.parent.name)
			current_dir = current_dir.parent
		} else if strings.HasPrefix(line, "$ cd") {
			fmt.Println("Moving down to directory")
			split := strings.Split(line, " ")
			name := split[2]
			new_dir := Directory{name: name, parent: current_dir}
			current_dir.subdirectories = append(current_dir.subdirectories, &new_dir)
			current_dir = &new_dir
		} else {
			split := strings.Split(line, " ")
			if split[0] == "dir" {
				fmt.Println("Skipping directory in list")
				continue
			}
			size, _ := strconv.Atoi(split[0])
			name := split[1]

			f := File{name: name, size: size}
			fmt.Println("Appending file", f)
			files = append(files, &f)
		}
	}

	if reading_files && len(files) > 0 {
		for _, f := range files {
			current_dir.files = append(current_dir.files, f)
		}
		fmt.Println("Added files to dir " + current_dir.name + " " + string(len(current_dir.files)))
	}
}

func print_filesystem(dir *Directory, indent int) {
	fmt.Println(strings.Repeat("/", indent), dir.name, dir.size)
	for _, f := range dir.files {
		prefix := strings.Repeat("-", indent+2)
		fmt.Println(prefix, f.name, f.size)
	}

	for _, sub_directory := range dir.subdirectories {
		print_filesystem(sub_directory, indent+2)
	}
}

func calculate_directory_size(dir *Directory) {
	var file_sizes int
	for _, file := range dir.files {
		file_sizes += file.size
	}

	var sub_directory_sizes int
	for _, sub_directory := range dir.subdirectories {
		calculate_directory_size(sub_directory)
		sub_directory_sizes += sub_directory.size
	}
	dir.size = file_sizes + sub_directory_sizes
}

func list_directory_sizes(root *Directory) []int {
	var queue []*Directory
	var dir *Directory
	var directory_sizes []int

	queue = append(queue, root)

	for len(queue) > 0 {
		dir, queue = queue[0], queue[1:]

		directory_sizes = append(directory_sizes, dir.size)

		for _, sub_directory := range dir.subdirectories {
			queue = append(queue, sub_directory)
		}
	}

	return directory_sizes
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// start with root directory
	root := Directory{name: "/"}

	build_filesystem(&root, lines)

	calculate_directory_size(&root)

	// print_filesystem(&root, 0)

	directory_sizes := list_directory_sizes(&root)

	var root_size int
	root_size, directory_sizes = directory_sizes[0], directory_sizes[1:]
	unused_space := 70000000 - root_size

	sort.Ints(directory_sizes)

	for _, size := range directory_sizes {
		if unused_space+size >= 30000000 {
			fmt.Println(size)
			break
		}
	}
}
