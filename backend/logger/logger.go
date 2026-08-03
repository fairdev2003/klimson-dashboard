package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"io"

	"github.com/fatih/color"
	"github.com/zgierz/klimson/backend/config"
)

func GetMethodColor(method string) func(a ...interface{}) string {

	switch method {
	case "GET":
		{
			return color.New(color.FgGreen).SprintFunc()
		}
	case "POST":
		{
			return color.New(color.FgYellow).SprintFunc()
		}
	case "PUT":
		{
			return color.New(color.FgBlue).SprintFunc()
		}
	case "DELETE":
		{
			return color.New(color.FgBlue).SprintFunc()
		}
	case "OPTIONS":
		{
			return color.New(color.FgCyan).SprintFunc()
		}
	}
	return color.New(color.FgWhite).SprintFunc()
}

func GreenServerLog(log ...any) {

	file, err := os.Open("./" + "server-config" + "/config.json")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Print(err.Error())
	}

	var cfg = config.PreGeneratedConfig
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		fmt.Print(err.Error())
	}

	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgHiCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Println(blue("● server/"+cfg.InternalName) + " " + yellow("("+cfg.Api.Version+")") + red(" % ") + green(log...))
}

func BlueServerLog(log ...any) {

	file, err := os.Open("./" + "server-config" + "/config.json")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Print(err.Error())
	}

	var cfg = config.PreGeneratedConfig
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		fmt.Print(err.Error())
	}

	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Println(blue("● server/"+cfg.InternalName) + " " + yellow("("+cfg.Api.Version+")") + blue(" % ") + blue(log...))
}

func GetFileLocation() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		file = "???"
	}

	_, filename := filepath.Split(file)

	return filename
}

func ServerLog(log ...any) {

	file, err := os.Open("./" + "server-config" + "/config.json")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Print(err.Error())
	}

	var cfg = config.PreGeneratedConfig
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		fmt.Print(err.Error())
	}

	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgHiCyan).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	purple := color.New(color.FgMagenta).SprintFunc()

	fmt.Println(blue("● server/"+cfg.InternalName) + " " + yellow("("+cfg.Api.Version+")") + fmt.Sprintf(" [ %s ]", purple(GetFileLocation())) + red(" % ") + white(fmt.Sprint(log...)))
}

func WarnLog(log ...any) {

	black := color.New(color.FgBlack).SprintFunc()
	yellow := color.New(color.BgYellow).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	fmt.Println(yellow(black(" WARN ") + " " + white(fmt.Sprint(log...))))
}

func ErrorLog(log ...any) {
	black := color.New(color.FgGreen).SprintFunc()
	redBg := color.New(color.BgRed).SprintFunc()
	redText := color.New(color.FgRed).SprintFunc()
	fmt.Println(redBg(black(" ERROR ") + " " + redText(fmt.Sprint(log...))))
}
