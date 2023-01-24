package ocrAddOn

import "github.com/otiai10/gosseract"

func makeSearchable(pdfFileName string) {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(pdfFileName)
	text, _ := client.Text()
	err := ioutil.WriteFile(pdfFileName, []byte(text), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
