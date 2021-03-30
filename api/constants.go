package api

func MavenSource() func(string) string {
	sources := map[string]string{
		"mavenCentral": "https://repo1.maven.org/maven2",
		"mavenApache":  "https://repo.maven.apache.org/maven2/",
		"google":       "https://dl.google.com/dl/android/maven2",
	}

	return func(key string) string {
		if source, contain := sources[key]; contain {
			return source
		}
		return key
	}
}
