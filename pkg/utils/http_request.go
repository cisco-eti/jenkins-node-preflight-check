package utils

import (
	"net/http"
	"net/url"
	log "sqbu-github.cisco.com/Nyota/frontline-common/goutils/fllogger"
	"sqbu-github.cisco.com/Nyota/go-template/pkg/models"
	"strconv"
)

//HTTPRequest struct
type HTTPRequest struct {
	*http.Request
}

//LogError log error level messages
func (req *HTTPRequest) LogError(err error) {
	if req == nil || err == nil {
		return
	}
	ctx := req.Context()
	log.WithContext(ctx).AddDepth().Error(err.Error())
}

//LogInfo log information level messages
func (req *HTTPRequest) LogInfo(message string) {
	if req == nil || message == "" {
		return
	}
	ctx := req.Context()
	log.WithContext(ctx).AddDepth().Info(message)
}

//GetPaginationLinks is a utility that helps paginated apis to return next page, previous page and last page links
func (req *HTTPRequest) GetPaginationLinks(responseSize int, limit int, offset int) models.Links {
	if !isPaginationParamsProper(responseSize, limit, offset) {
		return models.Links{}
	}
	path := req.URL
	values, _ := url.ParseQuery(path.RawQuery)
	values.Set("limit", strconv.Itoa(limit))
	values.Set("offset", "0")
	path.RawQuery = values.Encode()

	links := models.Links{
		First: path.String(),
	}

	nextOffset := offset + limit
	prevOffset := offset - limit

	if responseSize < limit {
		//no more pages so set last
		values.Set("offset", strconv.Itoa(offset))
		path.RawQuery = values.Encode()
		links.Last = path.String()
	} else {
		//likely more pages so set next
		values.Set("offset", strconv.Itoa(nextOffset))
		path.RawQuery = values.Encode()
		links.Next = path.String()
	}

	if prevOffset >= 0 {
		//previous page exists so set prev
		values.Set("offset", strconv.Itoa(prevOffset))
		path.RawQuery = values.Encode()
		links.Prev = path.String()
	}

	return links
}

func isPaginationParamsProper(responseSize int, limit int, offset int) bool {
	return responseSize >= 0 &&
		limit >= 0 &&
		offset >= 0
}
