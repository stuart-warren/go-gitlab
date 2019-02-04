package table

import (
	"fmt"
	"strconv"

	"github.com/makkes/gitlab-cli/api"
)

func pad(s string, width int) string {
	if width < 0 {
		return s
	}
	return fmt.Sprintf(fmt.Sprintf("%%-%ds", width), s)
}

func calcProjectColumnWidths(ps []api.Project) map[string]int {
	res := make(map[string]int)
	res["id"] = 15
	res["name"] = 40
	res["url"] = 50
	for _, p := range ps {
		w := len(strconv.Itoa(p.ID))
		if w > res["id"] {
			res["id"] = w
		}

		w = len(p.Name)
		if w > res["name"] {
			res["name"] = w
		}

		w = len(p.URL)
		if w > res["url"] {
			res["url"] = w
		}
	}
	return res
}

func calcPipelineColumnWidths(pipelines []api.PipelineDetails) map[string]int {
	res := make(map[string]int)
	res["id"] = 20
	res["status"] = 20
	res["duration"] = 10
	res["url"] = 50
	for _, p := range pipelines {
		w := len(fmt.Sprintf("%d:%d", p.ProjectID, p.ID))
		if w > res["id"] {
			res["id"] = w
		}

		w = len(p.Status)
		if w > res["status"] {
			res["status"] = w
		}

		w = len(strconv.Itoa(p.Duration))
		if w > res["duration"] {
			res["duration"] = w
		}

		w = len(p.URL)
		if w > res["url"] {
			res["url"] = w
		}
	}
	return res
}

func PrintPipelines(ps []api.PipelineDetails) {
	widths := calcPipelineColumnWidths(ps)
	fmt.Printf("%s %s %s %s\n",
		pad("ID", widths["id"]),
		pad("STATUS", widths["status"]),
		pad("DURATION", widths["duration"]),
		pad("URL", widths["url"]))
	for _, p := range ps {
		fmt.Printf("%s %s %s %s\n",
			pad(fmt.Sprintf("%d:%d", p.ProjectID, p.ID), widths["id"]),
			pad(p.Status, widths["status"]),
			pad(strconv.Itoa(p.Duration), widths["duration"]),
			pad(p.URL, widths["url"]))
	}
}

func PrintProjects(ps []api.Project) {
	widths := calcProjectColumnWidths(ps)
	fmt.Printf("%s %s %s\n",
		pad("ID", widths["id"]),
		pad("NAME", widths["name"]),
		pad("URL", widths["url"]))
	for _, p := range ps {
		fmt.Printf("%s %s %s\n",
			pad(strconv.Itoa(p.ID), widths["id"]),
			pad(p.Name, widths["name"]),
			pad(p.URL, widths["url"]))

	}
}