package javahelper

import (
	"errors"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"io"
	"net/http"
	"os"
	"strconv"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Download the file at url to the filepath,
// return an error if something goes wrong otherwise return nil
func downloadFile(url string, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("status code: " + strconv.Itoa(resp.StatusCode))
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	a := app.New()
	w := a.NewWindow("Downloading Java")
	w.SetCloseIntercept(func() {})

	progress := widget.NewProgressBar()

	label := widget.NewLabel("Downloading Java...")
	w.SetContent(
		container.NewVBox(
			label,
			progress,
		),
	)
	w.SetFixedSize(true)
	w.CenterOnScreen()

	// launch a goroutine to download the file in parts and update the progress bar
	go func() {
		var downloaded int64
		var total int64 = resp.ContentLength
		progress.Max = float64(total)
		for {
			// read a chunk
			n, err := io.CopyN(out, resp.Body, 1024*1024)
			if err != nil {
				break
			}
			downloaded += n
			//set the value of the progress bar on a scale from 0.0 to 1.0
			progress.SetValue(float64(downloaded))
			println("Downloaded", downloaded, "bytes of", total, "bytes ("+strconv.Itoa(int(progress.Value*100))+"%)")
		}
		if err != nil && err != io.EOF {
			println(err.Error())
		}
		w.Close()
	}()

	w.ShowAndRun()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// append a text at the end of a file, returns an error if something goes wrong otherwise nil
func FileAppend(filename string, text string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		println("Couldn't export path")
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		return err
	}

	return nil
}
