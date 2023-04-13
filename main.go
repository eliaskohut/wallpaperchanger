package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/reujab/wallpaper"
)

func main() {
	go func() {
		// Get the absolute path to the "wallpapers" directory
		dir, err := filepath.Abs("./wallpapers")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for {
			// Get the current date and time
			now := time.Now()

			// Get the day of the week (Sunday = 0, Monday = 1, etc.)
			dayOfWeek := int(now.Weekday())

			// Define the wallpaper filenames for each day of the week
			wallpapers := map[int]string{
				0: "dota2.jpg",
				1: "philosophy.jpg",
				2: "developing.png",
				3: "music.jpg",
				4: "history.jpg",
				5: "culture.jpg",
				6: "dota1.jpg",
			}

			// Get the filename for the current day of the week
			filename, ok := wallpapers[dayOfWeek]
			if !ok {
				fmt.Println("Error: Could not determine wallpaper filename.")
				os.Exit(1)
			}

			// Get the absolute path to the wallpaper file
			filepath := filepath.Join(dir, filename)

			// Set the desktop wallpaper
			err := wallpaper.SetFromFile(filepath)
			if err != nil {
				fmt.Println("Error: Could not set desktop wallpaper.")
				os.Exit(1)
			}
			err = wallpaper.SetMode(wallpaper.Crop)
			if err != nil {
				fmt.Println("Error: Could not set desktop wallpaper mode.")
				os.Exit(1)
			}

			// Print a success message
			fmt.Printf("Desktop wallpaper set to %s\n", filepath)

			// Sleep for 12 hours before changing the wallpaper again
			time.Sleep(12 * time.Hour)
		}
	}()

	// Wait forever (or until the program is terminated)
	select {}
}
