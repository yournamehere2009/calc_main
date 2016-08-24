package main

import (
	"fmt"
	"strconv"

	"net/http"

	"github.com/yournamehere2009/calc"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var html string

	html = "<html><head><title>Mike's Go Calculator - Result</title></head>"

	html += "<form action\"/calculate/\" method=\"POST\">" +
		"<input type=\"text\" name=\"formula\" >" +
		"<input type=\"checkbox\" name=\"show_work\" value=\"1\"> Show your work<br>" +
		"<input type=\"submit\" value=\"Go!\" >" +
		"</form>"

	if r.FormValue("formula") != "" {
		if result, workSteps, err := calc.ComputeFormula(r.FormValue("formula")); err == nil {
			if r.FormValue("show_work") == "1" {
				html += "<ul>"
				for i := 0; i < len(workSteps); i++ {
					html += "<li>" + workSteps[i] + "</li>"
				}
				html += "</ul>"
			} else {
				html += "<div>Answer: " + strconv.FormatFloat(float64(result), 'f', -1, 64) + "</div"
			}
		}
	}

	html += "</body></html>"

	fmt.Fprintf(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
