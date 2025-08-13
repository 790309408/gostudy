package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"regexp"
	"strings"
)

func docxToHTML(filePath string) (string, error) {
	// 解压ZIP获取document.xml
	r, err := zip.OpenReader(filePath)
	if err != nil {
		return "", err
	}
	defer r.Close()

	var htmlContent strings.Builder
	for _, f := range r.File {
		if f.Name == "word/document.xml" {
			rc, _ := f.Open()
			xmlData, _ := io.ReadAll(rc)
			rc.Close()

			// 自定义XML解析（示例：提取文本）
			htmlContent.WriteString(parseXMLToHTML(xmlData))
		}
	}
	return htmlContent.String(), nil
}

func parseXMLToHTML(xmlData []byte) string {
	// 使用regexp或xml.Decoder提取文本并包裹HTML标签
	re := regexp.MustCompile(`<w:t>(.*?)</w:t>`)
	matches := re.FindAllSubmatch(xmlData, -1)
	var sb strings.Builder
	for _, m := range matches {
		sb.WriteString("<p>" + string(m[1]) + "</p>")
	}
	return sb.String()
}

func Execute(filePath string) {
	str, _ := docxToHTML(filePath)
	fmt.Println(str)
	// doc, _ := document.Open(filePath)
	// // 转换为HTML（保留基础样式）
	// htmlContent := export.ToHTML(doc, export.HTMLOptions{
	// 	PreserveStyles: true,
	// })
	// _ = os.WriteFile("output.html", []byte(htmlContent), 0644)
}
