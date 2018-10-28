package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/rfaulhaber/fdate"
)

func TestCreateCalendar(t *testing.T) {
	testCases := []struct{
		input fdate.Date
		expected string
	}{
		{
			fdate.NewDate(226, fdate.Thermidor, 18),
`            Thermidor CCXXVI
Pri Duo Tri Qua Qui Sex Sep Oct Non Déc 
  1   2   3   4   5   6   7   8   9  10
 11  12  13  14  15  16  17 ` + " \033[7m18\033[27m " + ` 19  20
 21  22  23  24  25  26  27  28  29  30
`		},
		{
			fdate.NewDate(226, fdate.Complémentaires, 2),
			`         Complémentaires CCXXVI
 1 - La Fête de la Vertu
 2 - La Fête du Génie
` + "\033[7m 3\033[27m" +  ` - La Fête du Travail
 4 - La Fête de l'Opinion
 5 - La Fête des Récompenses
`,
		},

	}

	for _, tc := range testCases {
		result := CreateCalendar(tc.input)

		assert.Equal(t, tc.expected, result)
	}
}

