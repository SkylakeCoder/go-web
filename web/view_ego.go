package web

import (
	"io/ioutil"
	"regexp"
	"strings"
)

type viewEGO struct {
	labelReg *regexp.Regexp
}

func NewViewEGO() *viewEGO {
	return &viewEGO{
		labelReg: regexp.MustCompile("<%=.+?%>"),
	}
}

func (ve *viewEGO) Render(egoTemplatePath string, viewParams *ViewParams) (string, error) {
	bytes, err := ioutil.ReadFile(egoTemplatePath)
	if err != nil {
		return "", err
	}
	strTemplate := string(bytes)
	labels := ve.labelReg.FindAllString(strTemplate, -1)
	for _, v := range labels {
		label := v[3:len(v)-2]
		label = strings.TrimSpace(label)
		targetValue, err := viewParams.GetAsString(label)
		if err != nil {
			return "", err
		}
		strTemplate = strings.Replace(strTemplate, v, targetValue, -1)
	}
	return strTemplate, nil
}