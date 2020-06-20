package log

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

const (
	defaultFileSize = 10 * 1024 * 1024 // 10МБ, возможно в будущем стоит поправить.
	defaultMaxFiles = 25               // Максимальное количество файлов, пока посетителей не много, должно хватить.

	defaultTimeFormat = "2006-01-02T15:04:05"
	defaultGzipFormat = "gz"
	defaultFile       = "./logs/fast-shop.log"

	dot = "."
)

type FileWriter struct {
	mu       sync.Mutex
	file     *os.File
	dir      string
	filepath string
}

func NewWriter(filepath string) (w *FileWriter, err error) {
	if filepath == "" {
		filepath = defaultFile
	}

	w = &FileWriter{filepath: filepath}

	w.dir, _ = path.Split(filepath)

	if _, err = os.Stat(w.dir); os.IsNotExist(err) {
		if err := os.MkdirAll(w.dir, 0750); err != nil {
			return nil, err
		}
	}

	w.file, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (w *FileWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	stat, err := w.file.Stat()
	if err != nil {
		return 0, err
	}

	if stat.Size() > defaultFileSize {
		if err := w.rotateFile(); err != nil {
			return 0, err
		}
	}

	return w.file.Write(p)
}

func (w *FileWriter) Sync() (err error) {
	w.mu.Lock()
	err = w.file.Sync()
	w.mu.Unlock()

	return err
}

func (w *FileWriter) Close() (err error) {
	w.mu.Lock()
	err = w.file.Close()
	w.mu.Unlock()

	return err
}

func (w *FileWriter) rotateFile() (err error) {
	oldName := w.file.Name()

	// Добавляем время сжатия к текущему файлу.
	dir, file := path.Split(oldName)

	newName := buildName(dir, time.Now().Format(defaultTimeFormat), dot, file)

	if err := w.file.Close(); err != nil {
		return err
	}

	if err := os.Rename(oldName, newName); err != nil {
		return err
	}

	w.file, err = os.Create(w.filepath)
	if err != nil {
		return err
	}

	go w.compressAndRemoveOld(newName)

	return nil
}

func (w *FileWriter) compressAndRemoveOld(src string) {
	if err := w.compressLogFile(src); err != nil {
		log.Println("[error] compress log file failed")
	}

	if err := w.dropOldFiles(); err != nil {
		log.Println("[error] drop old log files failed")
	}
}

func buildName(args ...string) string {
	return strings.Join(args, "")
}

func (w *FileWriter) compressLogFile(src string) (err error) {
	file, err := os.Open(src)
	if err != nil {
		return err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// Добавляем формат сжатия к имени.
	dst := buildName(src, dot, defaultGzipFormat)

	gzFile, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fileInfo.Mode())
	if err != nil {
		return err
	}

	defer func() { _ = gzFile.Close() }()

	gzWriter := gzip.NewWriter(gzFile)

	defer func() {
		if err != nil {
			_ = os.Remove(dst)
		}
	}()

	if _, err := io.Copy(gzWriter, file); err != nil {
		return err
	}

	if err := gzWriter.Close(); err != nil {
		return err
	}

	if err := gzFile.Close(); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	if err := os.Remove(src); err != nil {
		return err
	}

	return nil
}

func (w *FileWriter) dropOldFiles() (err error) {
	files, err := ioutil.ReadDir(w.dir)
	if err != nil {
		return err
	}

	for i := 1; i < len(files)-defaultMaxFiles; i++ {
		if err := os.Remove(buildName(w.dir, files[i].Name())); err != nil {
			return err
		}
	}

	return nil
}
