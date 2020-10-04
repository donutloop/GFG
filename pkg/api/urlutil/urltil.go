package urlutil

import "net/url"

func BuildSelfReferenceURL(location *url.URL, endpoint string, uuid string) string {
	tmp := location.Path
	location.Path = location.Path + endpoint + "/" + uuid
	s := location.String()
	location.Path = tmp
	return s
}
