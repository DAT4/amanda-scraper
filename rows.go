package main

type row struct {
	rev         string
	startTime   string
	endTime     string
	expTime     string
	target      string
	ra          string
	dec         string
	pattern     string
	pi          string
	proposal    string
	observation string
	notes       string
}

func rowsFromTable(table [][]string) []row {
	var out []row
	for _, data := range table {
		out = append(
			out,
			row{
				rev:         data[0],
				startTime:   data[1],
				endTime:     data[2],
				expTime:     data[3],
				target:      data[4],
				ra:          data[5],
				dec:         data[6],
				pattern:     data[7],
				pi:          data[8],
				proposal:    data[9],
				observation: data[10],
				notes:       data[11],
			},
		)
	}
	return out
}
