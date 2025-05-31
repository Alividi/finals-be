package helper

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetQueryInt(r *http.Request, key string, def int) int {
	val := r.URL.Query().Get(key)
	if val == "" {
		return def
	}
	return StringToInt(val, def)
}

func GetQueryInt64(r *http.Request, key string, def int64) int64 {
	val := r.URL.Query().Get(key)
	if val == "" {
		return def
	}
	num, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return def
	}
	return num
}

func GetQueryFloat32(r *http.Request, key string, def float32) float32 {
	val := r.URL.Query().Get(key)
	if val == "" {
		return def
	}
	f, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return def
	}
	return float32(f)
}

func GetQueryFloat64(r *http.Request, key string, def float64) float64 {
	val := r.URL.Query().Get(key)
	if val == "" {
		return def
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return def
	}
	return f
}

func GetQueryString(r *http.Request, key, def string) string {
	val := r.URL.Query().Get(key)
	if val == "" {
		return def
	}
	return val
}

func StringToInt(s string, def int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return i
}

func GetURLParamInt(r *http.Request, key string) int {
	vars := mux.Vars(r)

	return StringToInt(vars[key], 0)
}

func GetURLParamInt64(r *http.Request, key string) int64 {
	vars := mux.Vars(r)
	num, err := strconv.ParseInt(vars[key], 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func GetURLParamString(r *http.Request, key string) string {
	vars := mux.Vars(r)

	return vars[key]
}

func GetQueryBool(r *http.Request, key string, def bool) bool {
	val := r.URL.Query().Get(key)
	if val == "" {
		return def
	}
	return StringToBool(val, def)
}

func StringToBool(s string, def bool) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return def
	}
	return b
}

func GetFormValue(r *http.Request, key string) string {
	return r.FormValue(key)
}

func GetFormFile(r *http.Request, key string) (multipart.File, *multipart.FileHeader, error) {
	file, header, err := r.FormFile(key)
	if err != nil {
		return nil, nil, err
	}

	if header.Size > 10*1024*1024 { // 10 MB limit
		return nil, nil, NewErrBadRequest(fmt.Sprintf("File size exceeds limit: %d mb", 10*1024*1024/1024/1024))
	}

	return file, header, nil
}

func GetFormFiles(r *http.Request, key string) ([]multipart.File, []*multipart.FileHeader, error) {
	if err := r.ParseMultipartForm(10 * 1024 * 1024); err != nil {
		return nil, nil, err
	}

	headers := r.MultipartForm.File[key]
	if len(headers) == 0 {
		return nil, nil, http.ErrMissingFile
	}

	files := make([]multipart.File, len(headers))
	for i, header := range headers {
		if header.Size > 10*1024*1024 {
			return nil, nil, NewErrBadRequest(fmt.Sprintf("File size exceeds limit: %d MB", 10*1024*1024/1024/1024))
		}

		file, err := header.Open()
		if err != nil {
			for j := 0; j < i; j++ {
				files[j].Close()
			}
			return nil, nil, err
		}
		files[i] = file
	}

	return files, headers, nil
}

func DerefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
