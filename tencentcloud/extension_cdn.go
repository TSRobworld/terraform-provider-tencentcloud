package tencentcloud

const (
	CDN_SERVICE_TYPE_WEB      = "web"
	CDN_SERVICE_TYPE_DOWNLOAD = "download"
	CDN_SERVICE_TYPE_MEDIA    = "media"

	CDN_ORIGIN_TYPE_DOMAIN  = "domain"
	CDN_ORIGIN_TYPE_COS     = "cos"
	CDN_ORIGIN_TYPE_IP      = "ip"
	CDN_ORIGIN_TYPE_IPV6    = "ipv6"
	CDN_ORIGIN_TYPE_IP_IPV6 = "ip_ipv6"

	CDN_ORIGIN_PULL_PROTOCOL_HTTP   = "http"
	CDN_ORIGIN_PULL_PROTOCOL_HTTPS  = "https"
	CDN_ORIGIN_PULL_PROTOCOL_FOLLOW = "follow"

	CDN_AREA_MAINLAND = "mainland"
	CDN_AREA_OVERSEAS = "overseas"
	CDN_AREA_GLOBAL   = "global"

	CDN_SWITCH_ON  = "on"
	CDN_SWITCH_OFF = "off"

	CDN_DOMAIN_STATUS_ONLINE     = "online"
	CDN_DOMAIN_STATUS_OFFLINE    = "offline"
	CDN_DOMAIN_STATUS_PROCESSING = "processing"

	CDN_SERVICE_NAME         = "cdn"
	CDN_RESOURCE_NAME_DOMAIN = "domain"

	CDN_HOST_NOT_FOUND      = "ResourceNotFound.CdnHostNotExists"
	CDN_DOMAIN_CONFIG_ERROE = "FailedOperation.CdnConfigError"
)

var CDN_SERVICE_TYPE = []string{
	CDN_SERVICE_TYPE_WEB,
	CDN_SERVICE_TYPE_DOWNLOAD,
	CDN_SERVICE_TYPE_MEDIA,
}

var CDN_ORIGIN_TYPE = []string{
	CDN_ORIGIN_TYPE_DOMAIN,
	CDN_ORIGIN_TYPE_COS,
	CDN_ORIGIN_TYPE_IP,
	CDN_ORIGIN_TYPE_IPV6,
	CDN_ORIGIN_TYPE_IP_IPV6,
}

var CDN_BACKUP_ORIGIN_TYPE = []string{
	CDN_ORIGIN_TYPE_IP,
	CDN_ORIGIN_TYPE_DOMAIN,
}

var CDN_ORIGIN_PULL_PROTOCOL = []string{
	CDN_ORIGIN_PULL_PROTOCOL_HTTP,
	CDN_ORIGIN_PULL_PROTOCOL_HTTPS,
	CDN_ORIGIN_PULL_PROTOCOL_FOLLOW,
}

var CDN_AREA = []string{
	CDN_AREA_MAINLAND,
	CDN_AREA_OVERSEAS,
	CDN_AREA_GLOBAL,
}

var CDN_SWITCH = []string{
	CDN_SWITCH_OFF,
	CDN_SWITCH_ON,
}
