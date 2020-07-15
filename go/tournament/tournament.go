package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	WinScore = 3
	TieScore = 1
)

type Team struct {
	name    string
	matches int
	wins    int
	loses   int
	ties    int
}

func (this Team) GetPoints() int {
	return this.wins*WinScore + this.ties*TieScore
}
func (this *Team) RegisterWin() {
	this.matches++
	this.wins++
}
func (this *Team) RegisterLoss() {
	this.matches++
	this.loses++
}
func (this *Team) RegisterTie() {
	this.matches++
	this.ties++
}

func Tally(reader io.Reader, writer io.Writer) error {
	teams := map[string]*Team{}
	scanner := bufio.NewScanner(reader)

	// scan lines
	for scanner.Scan() {
		// split into fields by ;
		// [home];[guest];[result]
		text := strings.TrimSpace(scanner.Text())
		if len(text) == 0 || strings.HasPrefix(text, "#") {
			continue
		}
		fields := strings.FieldsFunc(text, func(r rune) bool {
			return r == ';'
		})
		if len(fields) != 3 {
			return errors.New("Invalid row, less than 3 fields")
		}
		homeTeamName := fields[0]
		guestTeamName := fields[1]
		result := fields[2]
		// get or create Team objects
		home, ok := teams[homeTeamName]
		if !ok {
			home = &Team{name: homeTeamName}
			teams[homeTeamName] = home
		}
		guest, ok := teams[guestTeamName]
		if !ok {
			guest = &Team{name: guestTeamName}
			teams[guestTeamName] = guest
		}
		// handle result
		switch result {
		case "win":
			home.RegisterWin()
			guest.RegisterLoss()
		case "loss":
			guest.RegisterWin()
			home.RegisterLoss()
		case "draw":
			home.RegisterTie()
			guest.RegisterTie()
		default:
			return errors.New("Invalid game state")
		}
	}

	// map to slice
	arr := make([]*Team, 0, len(teams))
	for _, entry := range teams {
		arr = append(arr, entry)
	}

	// sort teams by points (desc), name if tied
	sort.Slice(arr, func(i, j int) bool {
		a := arr[i]
		b := arr[j]
		if a.GetPoints() == b.GetPoints() {
			return a.name < b.name
		} else {
			return a.GetPoints() > b.GetPoints()
		}
	})

	// output
	writer.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _, team := range arr {
		outputLine := fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", team.name, team.matches, team.wins, team.ties, team.loses, team.GetPoints())
		writer.Write([]byte(outputLine))
	}
	return nil
}
