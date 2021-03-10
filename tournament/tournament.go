package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type record struct {
	name           string
	MP, W, D, L, P int
}

// Tally adds stuff up
func Tally(reader io.Reader, writer io.Writer) error {

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	records := map[string]record{}

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" || strings.Contains(s, "#") {
			continue
		}
		r := strings.Split(s, ";")
		if len(r) != 3 {
			return fmt.Errorf("bad record: <%s>", r)
		}
		if len(r) == 3 {
			name1 := r[0]
			name2 := r[1]
			result := r[2]
			team1 := records[name1]
			team2 := records[name2]
			team1.name = name1
			team2.name = name2
			team1.MP++
			team2.MP++
			switch result {
			case "win":
				team1.W++
				team2.L++
				team1.P += 3
			case "loss":
				team1.L++
				team2.W++
				team2.P += 3
			case "draw":
				team1.D++
				team1.P++
				team2.D++
				team2.P++
			default:
				return errors.New("bad input")
			}
			records[name1] = team1
			records[name2] = team2
		}
	}

	// sort teams
	teams := make([]record, 0)
	for _, value := range records {
		teams = append(teams, value)
	}
	sort.SliceStable(teams, func(i, j int) bool {
		if teams[i].P != teams[j].P {
			return teams[i].P > teams[j].P
		}
		return teams[i].name < teams[j].name
	})

	fmt.Fprintf(writer, "%-30s | MP |  W |  D |  L |  P\n", "Team")
	for _, r := range teams {
		fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n", r.name, r.MP, r.W, r.D, r.L, r.P)
	}
	return nil
}
