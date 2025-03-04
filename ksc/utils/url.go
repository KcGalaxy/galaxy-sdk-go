package utils

import (
	"fmt"
)

type UrlInfo struct {
	UseSSL      bool
	Locate      bool
	UseInternal bool
	Domain      string
}

type ServiceInfo struct {
	Service string
	Region  string
}

const (
	defaultDomain  = "example.com"
	urlTpl         = "%s://%s.api.%s"
	internalPrefix = "internal"
	http           = "http"
	https          = "https"
)

func Url(urlInfo *UrlInfo, info ServiceInfo) string {
	protocol := Protocol(urlInfo.UseSSL)
	reqDomain := urlInfo.Domain
	if reqDomain == "" {
		reqDomain = defaultDomain
	}
	var reqUrlPrefix string
	if urlInfo.Locate && &info.Region != nil {
		reqUrlPrefix = info.Service + "." + info.Region
	} else if info.Service != "" {
		reqUrlPrefix = info.Service
	}
	if reqUrlPrefix == "" || urlInfo.UseInternal {
		reqUrlPrefix = internalPrefix
	}
	return fmt.Sprintf(urlTpl, protocol, reqUrlPrefix, reqDomain)
}

func Protocol(useSSL bool) string {
	if useSSL {
		return https
	}
	return http
}
