package web

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

const _PARTIAL_FLAG string = "partial"
const _LIST_ITEM_FLAG string = "listitem"

// viewEGO is the implementation of the ego template.
// ego template is just like the ejs in node.js Express framework.
type viewEGO struct {
	settings        *AppSettings
	varLabelReg     *regexp.Regexp
	htmlLabelReg    *regexp.Regexp
	partialLabelReg *regexp.Regexp
}

func newViewEGO(settings *AppSettings) *viewEGO {
	return &viewEGO{
		settings:        settings,
		varLabelReg:     regexp.MustCompile("<%=.+?%>"),
		htmlLabelReg:    regexp.MustCompile("<%-.+?%>"),
		partialLabelReg: regexp.MustCompile("<%=.+?" + _LIST_ITEM_FLAG + ".+?%>"),
	}
}

// Render parses the ego template and return the result.
func (ve *viewEGO) Render(templateRelativePath string, viewParams *KeyValues) (string, error) {
	bytes, err := ioutil.ReadFile(getTemplatePath(ve.settings.ViewDir, templateRelativePath))
	if err != nil {
		return "", err
	}
	strTemplate := string(bytes)
	varLabels := ve.varLabelReg.FindAllString(strTemplate, -1)
	for _, v := range varLabels {
		label := v[3 : len(v)-2]
		label = strings.Replace(label, " ", "", -1)
		targetValue, err := viewParams.GetAsString(label)
		if err != nil {
			return "", err
		}
		strTemplate = strings.Replace(strTemplate, v, targetValue, -1)
	}

	htmlMap := make(map[string]string)
	htmlLabels := ve.htmlLabelReg.FindAllString(strTemplate, -1)
	for _, v := range htmlLabels {
		label := v[3 : len(v)-2]
		label = strings.Replace(label, " ", "", -1)
		if !strings.Contains(label, _PARTIAL_FLAG) {
			_, exist := htmlMap[label]
			if !exist {
				bytes, err := ioutil.ReadFile(getTemplatePath(ve.settings.ViewDir, label))
				if err != nil {
					return "", err
				}
				htmlMap[label] = string(bytes)
			}
			html, _ := htmlMap[label]
			strTemplate = strings.Replace(strTemplate, v, html, -1)
		} else {
			path, items, err := ve.parsePartial(label)
			if err != nil {
				return "", err
			}
			_, exist := htmlMap[path]
			if !exist {
				bytes, err := ioutil.ReadFile(getTemplatePath(ve.settings.ViewDir, path))
				if err != nil {
					return "", err
				}
				htmlMap[path] = string(bytes)
			}
			html, _ := htmlMap[path]
			itemLabels := ve.partialLabelReg.FindAllString(html, -1)
			if len(itemLabels) < 1 {
				return "", fmt.Errorf("invalid partial template: %s", templateRelativePath)
			}
			l, err := viewParams.GetAsList(items)
			if err != nil {
				return "", err
			}
			replacedHtml := ""
			for e := l.Front(); e != nil; e = e.Next() {
				strValue := fmt.Sprintf("%v", e.Value)
				replacedHtml += strings.Replace(html, itemLabels[0], strValue, -1)
			}
			strTemplate = strings.Replace(strTemplate, v, replacedHtml, -1)
		}
	}

	return strTemplate, nil
}

func (ve *viewEGO) parsePartial(exp string) (path, items string, err error) {
	exp = strings.Replace(exp, _PARTIAL_FLAG+"(", "", -1)
	exp = strings.Replace(exp, ")", "", -1)
	split := strings.Split(exp, ",")
	if len(split) != 2 {
		err = fmt.Errorf("invalid partial expression: %s", exp)
		return
	}
	path = strings.Replace(split[0], "\"", "", -1)
	path = strings.Replace(path, "'", "", -1)
	items = split[1]
	err = nil
	return
}
