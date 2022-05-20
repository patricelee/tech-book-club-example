package sproutclass

// ----- legacy code -----
type QuarterlyRepoetGenerator struct {
	Data []Department
}

func (q QuarterlyRepoetGenerator) Generate() string {
	pageText := "<html><header><title>Quarterly Report</title></head></body><table>"
	for _, d := range q.Data {
		pageText += "<tr>"
		pageText += "<td>" + d.Manager + "</td>"
		pageText += "<td>" + d.Manager + "</td>"
		// a log code...
	}

	return pageText
}

type Department struct {
	Name    string
	Manager string
}

// ----- legacy code -----

