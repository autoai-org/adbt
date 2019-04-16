package adbt

var
(
	SupportedDatabases = []string{"MongoDB", "MySQL"}
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