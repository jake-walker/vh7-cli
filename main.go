package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/jake-walker/vh7-cli/vh7"
	"github.com/pterm/pterm"
	"github.com/teris-io/cli"
	"golang.org/x/tools/godoc/util"
)

const apiUrl string = "https://vh7.uk/"
const defaultPasteLanguage string = "plaintext"

func getApi() *vh7.ApiService {
	return vh7.NewApiService(http.DefaultClient, apiUrl)
}

func actionShorten(args []string, options map[string]string) int {
	if !validateUrl(args[0]) {
		pterm.Error.Println("Whoops! That isn't a valid URL.")
		return 1
	}

	api := getApi()
	spinner, _ := pterm.DefaultSpinner.Start("Shortening...")

	shorten, err := api.CreateShorten(args[0])

	if err != nil {
		spinner.Fail("Something went wrong!")
		pterm.Error.Println("There was a problem shortening!", err)
		return 1
	}

	spinner.Success("Shortened!")
	pterm.Success.Println("Your URL has been shortened to:\n" +
		fullUrl(shorten.Link))
	// saveToHistory(*shorten)
	return 0
}

func actionPaste(args []string, options map[string]string) int {
	text := ""
	pasteType := ""

	if len(args) < 1 {
		pasteType = "clipboard"

		if clipboard.Unsupported {
			pterm.Error.Println("Cannot read your clipboard. Please specify a file to paste instead.")
			return 1
		}

		text, _ = clipboard.ReadAll()
	} else {
		pasteType = "file"

		if !fileExists(args[0]) {
			pterm.Error.Printf("Cannot open file %v!", args[0])
			return 1
		}

		data, err := ioutil.ReadFile(args[0])
		if err != nil {
			pterm.Error.Println("There was a problem reading the file!", err)
			return 1
		}

		if !util.IsText(data) {
			pterm.Error.Println("The file does not look like text. Try using the 'upload' command instead.")
			return 1
		}

		text = string(data)
	}

	text = strings.TrimSpace(text)

	if text == "" {
		pterm.Error.Printf("Your %v contains no text.", pasteType)
		return 1
	}

	language, found := options["language"]

	if !found {
		language = defaultPasteLanguage
	}

	api := getApi()
	spinner, _ := pterm.DefaultSpinner.Start("Pasting...")

	paste, err := api.CreatePaste(text, language)

	if err != nil {
		spinner.Fail("Something went wrong!")
		pterm.Error.Println("There was a problem pasting!", err)
		return 1
	}

	spinner.Success("Pasted!")
	pterm.Success.Println("Your paste has been shortened to:\n" +
		fullUrl(paste.Link))
	// saveToHistory(*paste)
	return 0
}

func actionUpload(args []string, options map[string]string) int {
	if !fileExists(args[0]) {
		pterm.Error.Printf("Cannot open file %v!", args[0])
		return 1
	}

	file, err := os.Open(args[0])
	if err != nil {
		pterm.Error.Println("There was a problem opening the file!", err)
		return 1
	}
	defer file.Close()

	api := getApi()
	spinner, _ := pterm.DefaultSpinner.Start("Uploading...")

	upload, err := api.CreateUpload(file)

	if err != nil {
		spinner.Fail("Something went wrong!")
		pterm.Error.Println("There was a problem pasting!", err)
		return 1
	}

	spinner.Success("Uploaded!")
	pterm.Success.Println("Your upload has been been created!\n" +
		"It will expire on " + prettyDate(upload.Expiry) + "\n" +
		fullUrl(upload.Link))
	// saveToHistory(*upload)
	return 0
}

func actionInfo(args []string, options map[string]string) int {
	link := cleanLink(args[0])
	api := getApi()

	spinner, _ := pterm.DefaultSpinner.Start("Fetching info...")

	info, err := api.GetInfo(link)

	if err != nil {
		spinner.Fail("Something went wrong!")
		pterm.Error.Println("There was a problem shortening!", err)
		return 1
	}

	spinner.Success("Fetched info!")

	pterm.DefaultSection.Println(link)

	metadata := fmt.Sprintf("Created:  %v\n"+
		"Updated:  %v\n"+
		"Expires:  %v", prettyDate(info.Created), prettyDate(info.Updated), prettyDate(info.Expiry))

	infoType := info.GetType()

	if infoType == "url" {
		pterm.Info.Println(fmt.Sprintf("Type:     Short URL\n"+
			"URL:      %v\n", info.Url.Url) + metadata)
	} else if infoType == "paste" {
		pterm.Info.Println(fmt.Sprintf("Type:     Paste\n"+
			"Language: %v\n"+
			"Hash:     %v\n", info.Paste.Language, info.Paste.Hash) + metadata)
	} else if infoType == "upload" {
		pterm.Info.Println(fmt.Sprintf("Type:     Upload\n"+
			"Filename: %v\n"+
			"Mimetype: %v\n"+
			"Hash:     %v\n", info.Upload.OriginalFilename, info.Upload.Mimetype, info.Upload.Hash) + metadata)
	} else {
		pterm.Error.Println("This link is an unknown type!")
	}

	return 0
}

func actionHistory(args []string, options map[string]string) int {
	history := loadHistory()
	tableData := pterm.TableData{
		{"Date", "Link", "Type", "Summary"},
	}

	for i := len(history) - 1; i >= 0; i-- {
		item := history[i]
		tableData = append(tableData, []string{
			prettyTime(item.Date),
			fullUrl(item.Data.Link),
			item.Data.GetType(),
			item.Data.GetSummary(),
		})
	}

	pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()

	return 0
}

func main() {
	shorten := cli.NewCommand("shorten", "shorten a url").
		WithShortcut("s").
		WithArg(cli.NewArg("url", "the url to shorten")).
		WithAction(actionShorten)

	paste := cli.NewCommand("paste", "paste some code").
		WithShortcut("p").
		WithOption(cli.NewOption("language", "the language to highlight the paste as (default: plaintext)").WithType(cli.TypeString)).
		WithArg(cli.NewArg("path", "the file to paste").AsOptional()).
		WithAction(actionPaste)

	upload := cli.NewCommand("upload", "upload a file").
		WithShortcut("u").
		WithArg(cli.NewArg("path", "the file to upload")).
		WithAction(actionUpload)

	info := cli.NewCommand("info", "get information about a link").
		WithShortcut("i").
		WithArg(cli.NewArg("link", "the link to get info about")).
		WithAction(actionInfo)

	// history := cli.NewCommand("history", "get history of created short links").
	// 	WithAction(actionHistory)

	app := cli.New("VH7 URL shortener, pastebin and temporary file storage.\n" +
		fmt.Sprintf("    Version: %s, Commit: %s, Built at: %s.", BuildVersion, commit, date)).
		// WithOption(cli.NewOption("silent", "silent execution (just the output)").WithType(cli.TypeBool)).
		WithCommand(shorten).
		WithCommand(paste).
		WithCommand(upload).
		WithCommand(info) // .
		// WithCommand(history)

	os.Exit(app.Run(os.Args, os.Stdout))
}
