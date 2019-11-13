package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	ERR_READPATH = errors.New("Path could not be read")
)

// FindExtension returns all files with a specified extension
func FindExtension(notePath, extension string) (files []string) {
	err := filepath.Walk(notePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		// is file ends with md
		if !info.IsDir() && strings.HasSuffix(info.Name(), extension) {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		log.Printf("FKKK %s", err)
	}
	return
}

// FindExtensionMaxdepth todo
func FindExtensionMaxdepth(notePath, extension string, maxdepth int) (files []string) {
	maxdepth -= 1
	if maxdepth < 0 {
		return
	}
	infos, err := ioutil.ReadDir(notePath)
	if err != nil {
		return
	}

	for _, info := range infos {
		newPath := filepath.Join(notePath, info.Name())
		if !info.IsDir() && strings.HasSuffix(info.Name(), extension) {
			files = append(files, newPath)
		}
		if info.IsDir() {
			recfiles := FindExtensionMaxdepth(newPath, extension, maxdepth)
			files = append(files, recfiles...)
		}
	}
	return
}

// FindJSON returns all jsonFiles on the notePath
func FindJSON(notePath string) (jsonFiles []string) {
	return FindExtension(notePath, ".json")
}

// FindMd returns an array of all md paths within the notes dir
func FindMd(notePath string) (mdFiles []string) {
	return FindExtension(notePath, ".md")
}

//
func FindMilestones(misatoPath string) []string {
	return FindExtension(misatoPath, ".milestones.json")
}

func FindActivities(misatoPath string) []string {
	return FindExtension(misatoPath, ".activities.json")
}

//
func FindMisatoLogs(misatoPath string) []string {
	return FindExtension(misatoPath, ".misato.json")
}

// 
func FindGraphs(graphDir string) []string{
	return FindExtension(graphDir, ".graph.json")
}

// SearchInFile searches a pattern and returns wether or not it was found
func SearchInFile(filename string, searchword string) (bool, error) {

	read, err := ioutil.ReadFile(filename)
	if err != nil {
		return false, ERR_READPATH
	}
	return strings.Contains(string(read), searchword), nil
}

// ReadLinesFromFiles
func ReadLinesFromFiles(filename string) ([]string, error) {
	read, err := ioutil.ReadFile(filename)
	if err != nil {
		return []string{}, ERR_READPATH
	}
	return strings.Split(string(read), "\n"), nil
}

// SearchWithPreviewInFile returns
// TODO , n int with previous and after line
func SearchWithPreviewInFile(filename, searchword string) (matches []string, err error) {

	//
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, searchword) {
			matches = append(matches, line)
		}
	}

	if err = scanner.Err(); err != nil {
		return
	}
	return
}

func TrimPart(part, prefix string) (rest string) {
	rest = strings.TrimPrefix(part, prefix)
	rest = strings.TrimPrefix(rest, "/")
	return
}

// TrimParts returns an stripped array
func TrimParts(parts []string, prefix string) []string {
	stripped := make([]string, len(parts))
	for i := range parts {
		stripped[i] = TrimPart(parts[i], prefix)
	}
	return stripped
}

// HasMDFilesChanged validates if any mdFile changed and returns the modified files
func HasMDFilesChanged(since time.Time, notePath string) (changes []string, hasChanged bool) {
	mds := FindMd(notePath)

	for _, md := range mds {
		fi, err := os.Stat(md)
		if err != nil {
			log.Printf("FKKKK %s", err)
			return
		}
		if fi.ModTime().After(since) {
			hasChanged = true
			changes = append(changes, md)
		}
	}
	return
}

// FileExist checks if a file Exist and returns false if a file does not exist or an error occur
func FileExist(p string) bool {
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	}
	return true
}

func GetLelDir()(lelDir string){
	if lelDir = os.Getenv("LEL_CONFIG_DIR");lelDir == "" {
		homeDir, _ := homedir.Dir()
		lelDir = filepath.Join(homeDir,".LeL")
	}
	return 
}

// DirIsEmpty checks if the directory contains any files
func DirIsEmpty(d string) (empty bool, err error) {
	entries, err := ioutil.ReadDir(d)
	if err != nil {
		return
	}
	return len(entries) == 0, nil
}

// GetCleanFilePath returns the abs, clean path and verifies if the file Exists
func GetAbsPath(baseDir, fileToOpen string) (abs string, err error) {
	untrustedPath := filepath.Join(baseDir, filepath.FromSlash(path.Clean("/"+fileToOpen)))
	abs, err = filepath.Abs(untrustedPath)
	return
}

// WriteTo write the bts to filename
func WriteTo(filename string, bts []byte) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(bts)
	return err
}

// OpenInEditor execute the editorPath with the filepath as an argument
func OpenInEditor(executablePath, filePath string) error {
	cmd := exec.Command(executablePath, filePath)
	return cmd.Start()
}

// FindTagsInFile returns all tags, which were found in the file
func FindTagsInFile(filename string) (tags []string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if text := scanner.Text(); strings.HasPrefix(text, TAGS_PATTERN) {
			tags = append(tags, SplitAndTrim(text, TAGS_PATTERN)...)
		}
	}
	err = scanner.Err()
	return
}

func SplitAndTrim(line, pattern string) (parts []string) {
	line = strings.TrimPrefix(line, pattern)
	for _, part := range strings.Split(line, VAL_DELIMITER) {
		parts = append(parts, strings.TrimSpace(part))
	}
	return
}

func SliceAreEqual(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
