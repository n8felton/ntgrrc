package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"strconv"
	"strings"
)

type PoePortSetting struct {
	PortIndex int8
	PortPwr   bool
	PwrMode   string
	PortPrio  string
	LimitType string
	PwrLimit  string
	DetecType string
}

type PoeShowSettingsCommand struct {
	Address string `required:"" help:"the Netgear switch's IP address or host name to connect to" short:"a"`
}

func (poe *PoeShowSettingsCommand) Run(args *GlobalOptions) error {
	settingsPage, err := requestPoePortConfigPage(args, poe.Address)
	if err != nil {
		return err
	}
	if checkIsLoginRequired(settingsPage) {
		return errors.New("no content. please, (re-)login first")
	}
	var settings []PoePortSetting
	settings, err = findPortSettingsInHtml(strings.NewReader(settingsPage))
	if err != nil {
		return err
	}
	prettyPrintSettings(args.OutputFormat, settings)
	return nil
}

func prettyPrintSettings(format OutputFormat, settings []PoePortSetting) {
	var header = []string{"Port ID", "Port Power", "Mode", "Priority", "Limit Type", "Limit (W)", "Type"}
	var content [][]string
	for _, setting := range settings {
		var row []string
		row = append(row, fmt.Sprintf("%d", setting.PortIndex))
		row = append(row, asTextPortPower(setting.PortPwr))
		row = append(row, bidiMapLookup(setting.PwrMode, pwrModeMap))
		row = append(row, bidiMapLookup(setting.PortPrio, portPrioMap))
		row = append(row, bidiMapLookup(setting.LimitType, limitTypeMap))
		row = append(row, setting.PwrLimit)
		row = append(row, bidiMapLookup(setting.DetecType, detecTypeMap))
		content = append(content, row)
	}
	switch format {
	case MarkdownFormat:
		printMarkdownTable(header, content)
	case JsonFormat:
		printJsonDataTable("settings", header, content)
	default:
		panic("not implemented format: " + format)
	}
}

func asTextPortPower(portPwr bool) string {
	if portPwr {
		return "enabled"
	}
	return "disabled"
}

func requestPoePortConfigPage(args *GlobalOptions, host string) (string, error) {
	url := fmt.Sprintf("http://%s/PoEPortConfig.cgi", host)
	return requestPage(args, host, url)
}

func findPortSettingsInHtml(reader io.Reader) ([]PoePortSetting, error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	var configs []PoePortSetting
	doc.Find("li.poePortSettingListItem").Each(func(i int, s *goquery.Selection) {
		config := PoePortSetting{}

		id := s.Find("span.poe-port-index span").Text()
		var id64, _ = strconv.ParseInt(id, 10, 8)
		config.PortIndex = int8(id64)

		portWr, exists := s.Find("input#hidPortPwr").Attr("value")
		config.PortPwr = exists && portWr == "1"

		config.PwrMode, _ = s.Find("input#hidPwrMode").Attr("value")

		config.PortPrio, _ = s.Find("input#hidPortPrio").Attr("value")

		config.LimitType, _ = s.Find("input#hidLimitType").Attr("value")

		config.PwrLimit, _ = s.Find("input.pwrLimit").Attr("value")

		config.DetecType, _ = s.Find("input#hidDetecType").Attr("value")

		configs = append(configs, config)
	})
	return configs, nil
}
