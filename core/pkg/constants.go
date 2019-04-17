package adbt

var
(
	SupportedDatabases = []string{"MongoDB", "MySQL"}
	SupportedPeriods  = []string{"Weekly", "Daily"}
)

// helper functions

func isDatabaseSupported(db string) bool {
	for _, d := range SupportedDatabases {
		if d == db {
			return true
		}
	}
	return false
}