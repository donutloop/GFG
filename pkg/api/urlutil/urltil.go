package urlutil

import "net/url"

func BuildSelfReferenceURL(location *url.URL, endpoint string, uuid string) string {
	location.Path = location.Path + endpoint + "?id=" + uuid
	return location.String()
}
