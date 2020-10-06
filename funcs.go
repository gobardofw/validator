package validator

import (
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gobardofw/utils"
	"github.com/inhies/go-bytesize"
)

// IsUsername check if string is valid username
func IsUsername(username string) bool {
	r := regexp.MustCompile(`^[0-9a-zA-Z\-\._]+$`)
	return r.MatchString(username)
}

// IsTel check if string is valid tel
func IsTel(tel string) bool {
	r := regexp.MustCompile(`^(\(0\d{2}\) \d{4}-\d{4})|(0\d{10})$`)
	return r.MatchString(tel)
}

// IsMobile check if string is valid mobile number
func IsMobile(mobile string) bool {
	r := regexp.MustCompile(`^(\(09\d{2}\) \d{3}-\d{4})|(09\d{9})$`)
	return r.MatchString(mobile)
}

// IsPostalcode check if string is valid postal code
func IsPostalcode(postalCode string) bool {
	r := regexp.MustCompile(`^(\d{5}-\d{5})|(\d{10})$`)
	return r.MatchString(postalCode)
}

// IsIdentifier check if string is valid identifier
func IsIdentifier(id string) bool {
	idf, _ := strconv.Atoi(id)
	return idf > 0
}

// IsUnsigned check if string is unsigned number
func IsUnsigned(num string) bool {
	n, _ := strconv.Atoi(num)
	return n >= 0
}

// IsIDNumber check if string is valid id number
func IsIDNumber(idNum string) bool {
	r := regexp.MustCompile(`^\d{1,10}$`)
	return r.MatchString(idNum)
}

// ISNationalCode check if string is valid id national code
func ISNationalCode(idNum string) bool {
	r := regexp.MustCompile(`^(\d{3}-\d{6}-\d{1})|(\d{10})$`)
	return r.MatchString(idNum)
}

// IsCreditCardNumber check if string is valid id credit card number
func IsCreditCardNumber(num string) bool {
	r := regexp.MustCompile(`^(\d{4}-\d{4}-\d{4}-\d{4}-\d{4})|(\d{4}-\d{4}-\d{4}-\d{4})|(\d{20})|(\d{16})$`)
	return r.MatchString(num)
}

// IsUUID check if string is valid uuid
func IsUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// IsJDate check if string is valid jalali date jalali date validator
func IsJDate(jDate string) bool {
	if _, err := utils.JalaliToTime(jDate); err == nil {
		return true
	}
	return false
}

// IsIP check if address if a valid ip
func IsIP(address string) bool {
	r := regexp.MustCompile(`^(([0–9]|[1–9][0–9]|1[0–9]{2}|2[0–4][0–9]|25[0–5])\.){3}([0–9]|[1–9][0–9]|1[0–9]{2}|2[0–4][0–9]|25[0–5])$`)
	return r.MatchString(address)
}

// IsIPPort check if address if a valid ip contains port
func IsIPPort(address string) bool {
	r := regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]):[0-9]+$`)
	return r.MatchString(address)
}

// ValidateUploadSize check if file size in range
// use B, KB, MB, GB for size string
// ex: 1KB, 3MB
// Note: not use float point!
func ValidateUploadSize(file *multipart.FileHeader, min string, max string) (bool, error) {
	minSize, err := bytesize.Parse(min)
	if err != nil {
		return false, err
	}

	maxSize, err := bytesize.Parse(max)
	if err != nil {
		return false, err
	}

	return (uint64(file.Size) >= uint64(minSize) && uint64(file.Size) <= uint64(maxSize)), nil
}

// ValidateUploadMime check if file upload mime is valid
func ValidateUploadMime(file *multipart.FileHeader, mimes ...string) (bool, error) {
	f, err := file.Open()
	if err != nil {
		return false, err
	}
	defer f.Close()

	var data []byte
	_, err = f.Read(data)
	if err != nil {
		return false, err
	}

	mime := mimetype.Detect(data)
	return mimetype.EqualsAny(mime.String(), mimes...), nil
}

// ValidateUploadExt check if file upload extension is valid
func ValidateUploadExt(file *multipart.FileHeader, exts ...string) (bool, error) {
	f, err := file.Open()
	if err != nil {
		return false, err
	}
	defer f.Close()

	var data []byte
	_, err = f.Read(data)
	if err != nil {
		return false, err
	}

	mime := mimetype.Detect(data)

	for _, ext := range exts {
		if strings.ToLower(ext) == strings.ToLower(mime.Extension()) {
			return true, nil
		}
	}
	return false, nil
}
