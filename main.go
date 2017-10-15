package main
import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	bReadError    := false
	settings      := ""
	slicer        := "N/A"
	layers        := "N/A"
	quality       := "N/A"
	profile       := "N/A"
	supports      := "N/A"
	retraction    := "N/A"
	jerk          := "N/A"
	speed1        := "N/A"
	speedprint    := "N/A"
	speedtravel   := "N/A"
	filamentsize  := "N/A"
	hotendtemp    := "N/A"
	bedtemp       := "N/A"
	infillpattern := "N/A"
	if len(os.Args) != 2 {
		fmt.Printf("Syntax: SlicingInfo GCodeFilePath\n\n")
		return
	}
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			bReadError = true;
			fmt.Fprintf(os.Stderr, "SlicingInfo:\n  %v\n\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if strings.Contains(line, ";Generated with") { slicer = line[16:] }
			if strings.Contains(line, ";LAYER_COUNT:")   { layers = line[13:] }
			if strings.Contains(line, ";SETTING_3 ") {
				settings += line[11:]
			}
			//fmt.Println(line)
		}
	}
	if !bReadError {
		settings = strings.Replace(settings, "\\\\n", "\n", -1)
		for _, settingsLine := range strings.Split(string(settings), "\n") {
			if strings.Contains(settingsLine, "quality_type = ") { quality = settingsLine[15:] }
			if strings.Contains(settingsLine, "name = ")         { profile = settingsLine[7:] }
			if strings.Contains(settingsLine, "support_enable = ") {
				if settingsLine[17:] == "True"  { supports = "True" }
				if settingsLine[17:] == "False" { supports = "False" }
			}
			if strings.Contains(settingsLine, "retraction_enable = ") {
				if settingsLine[20:] == "True"  { retraction = "True" }
				if settingsLine[20:] == "False" { retraction = "False" }
			}
			if strings.Contains(settingsLine, "jerk_enabled = ") {
				if settingsLine[15:] == "True"  { jerk = "True" }
				if settingsLine[15:] == "False" { jerk = "False" }
			}
			if strings.Contains(settingsLine, "speed_layer_0 = ")              { speed1 = settingsLine[16:] }
			if strings.Contains(settingsLine, "speed_print = ")                { speedprint = settingsLine[14:] }
			if strings.Contains(settingsLine, "speed_travel = ")               { speedtravel = settingsLine[15:] }
			if strings.Contains(settingsLine, "material_diameter = ")          { filamentsize = settingsLine[20:] }
			if strings.Contains(settingsLine, "material_print_temperature = ") { hotendtemp = settingsLine[29:] }
			if strings.Contains(settingsLine, "material_bed_temperature = ")   { bedtemp = settingsLine[27:] }
			if strings.Contains(settingsLine, "infill_pattern = ")             { infillpattern = settingsLine[17:] }
		}
		if (slicer != "N/A")        { fmt.Printf("Slicer:          %s\n", slicer) }
		if (layers != "N/A")        { fmt.Printf("Layers:          %s\n", layers) }
		if (quality != "N/A")       { fmt.Printf("Quality:         %s\n", quality) }
		if (profile != "N/A")       { fmt.Printf("Profile:         %s\n", profile) }
		if (filamentsize != "N/A")  { fmt.Printf("Filament size:   %s\n", filamentsize) }
		if (hotendtemp != "N/A")    { fmt.Printf("Hotend temp:     %s\n", hotendtemp) }
		if (bedtemp != "N/A")       { fmt.Printf("Bed temp:        %s\n", bedtemp) }		
		if (supports != "N/A")      { fmt.Printf("Supports:        %s\n", supports) }
		if (retraction != "N/A")    { fmt.Printf("Retraction:      %s\n", retraction) }
		if (jerk != "N/A")          { fmt.Printf("Jerk:            %s\n", jerk) }
		if (speed1 != "N/A")        { fmt.Printf("Speed 1st layer: %s\n", speed1) }
		if (speedprint != "N/A")    { fmt.Printf("Print speed:     %s\n", speedprint) }
		if (speedtravel != "N/A")   { fmt.Printf("Travel speed:    %s\n", speedtravel) }
		if (infillpattern != "N/A") { fmt.Printf("Infill pattern:  %s\n", infillpattern) }
		//if (settings != "N/A")      { fmt.Printf("Settings:\n\n  %s\n", settings) }

		fmt.Printf("\nFinished.\n\n")
	}
}