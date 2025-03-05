package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Desktop File Maker")

	// Entry widgets for each field
	Type_Label := widget.NewLabel("Type*")
	Type := widget.NewEntry()
	Type.SetPlaceHolder("Enter App Type Here...")

	Version_Label := widget.NewLabel("App Version")
	Version := widget.NewEntry()
	Version.SetPlaceHolder("Enter App Version Here...")

	Name_Label := widget.NewLabel("App Name*")
	Name := widget.NewEntry()
	Name.SetPlaceHolder("Enter App Name Here...")

	Comment_Label := widget.NewLabel("App Comment")
	Comment := widget.NewEntry()
	Comment.SetPlaceHolder("Enter App Comment Here...")

	Exec_Label := widget.NewLabel("App Exec*")
	Exec := widget.NewEntry()
	Exec.SetPlaceHolder("Enter App Exec Path Here...")

	Icon_Label := widget.NewLabel("App Icon")
	Icon := widget.NewEntry()
	Icon.SetPlaceHolder("Enter App Icon Path Here...")

	Terminal_Label := widget.NewLabel("Does the app run in a Terminal?")
	Terminal := widget.NewCheck("Terminal?", func(b bool) {})

	Category_Label := widget.NewLabel("App's Categories")
	Category := widget.NewEntry()
	Category.SetPlaceHolder("Enter App Categories Here...")

	StartupWMClass_Label := widget.NewLabel("App's StartupWMClass (Recommended)")
	StartupWMClass := widget.NewEntry()
	StartupWMClass.SetPlaceHolder("Enter App's StartupWMClass Here...")

	// Set window size
	myWindow.Resize(fyne.NewSize(900, 600))

	// Create the UI layout
	content := container.NewVBox(
		Type_Label, Type,
		Version_Label, Version,
		Name_Label, Name,
		Comment_Label, Comment,
		Exec_Label, Exec,
		Icon_Label, Icon,
		Terminal_Label, Terminal,
		Category_Label, Category,
		StartupWMClass_Label, StartupWMClass,

		widget.NewButton("Make Desktop", func() {
			// Get data from user input
			appName := Name.Text
			execPath := Exec.Text
			iconPath := Icon.Text
			comment := Comment.Text
			terminal := Terminal.Checked
			categories := Category.Text
			startupWMClass := StartupWMClass.Text

			// Create desktop file content
			desktopFileContent := fmt.Sprintf(`
			
[Desktop Entry]
Name=%s
Exec=%s
Icon=%s
Type=Application
Categories=%s
Terminal=%t
Comment=%s
StartupWMClass=%s
`, appName, execPath, iconPath, categories, terminal, comment, startupWMClass)

			// Save the desktop file
			err := os.WriteFile(appName+".desktop", []byte(desktopFileContent), 0644)
			if err != nil {
				log.Println("Error writing desktop file:", err)
			} else {
				log.Println("Desktop file created:", appName+".desktop")
			}
			// Mark the file as executable
			err = os.Chmod(appName+".desktop", 0755)
			if err != nil {
				log.Println("Error setting file as executable:", err)
			} else {
				log.Println("Desktop file is now executable")
			}

			// Copy the .desktop file to the trusted applications directory
			desktopDir := os.Getenv("HOME") + "/.local/share/applications/"
			err = ioutil.WriteFile(desktopDir+appName+".desktop", []byte(desktopFileContent), 0755)
			if err != nil {
				log.Println("Error copying desktop file to trusted directory:", err)
			} else {
				log.Println("Desktop file added to trusted directory")
			}
			//update the desktop data base
			cmd := exec.Command("update-desktop-database", "~/.local/share/applications")
			err = cmd.Run()
			if err != nil {
				log.Println("Error updating desktop database:", err)
			} else {
				log.Println("Desktop database updated")
			}
		}),
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
