package naming // import "project-draco.io/naming"

import (
	"fmt"
	"regexp"
	"strings"
)

// JavaToHR converts a qualified java method/constructor/field name to HR file name.
// Returns an empty string if the qualified name is not supported
func JavaToHR(qn string) string {
	if qn == "" || strings.Contains(qn, "static {}") {
		return ""
	}
	kind := "[FE]"
	args := ""
	if idxpar := strings.Index(qn, "("); idxpar != -1 {
		kind = "[MT]"
		args = qn[idxpar+1 : len(qn)-1]
		args = RemoveGenerics(args)
		arr := strings.Split(args, ",")
		for i := range arr {
			arr[i] = lastSegment(strings.TrimSpace(arr[i]))
		}
		args = fmt.Sprintf("(%v)", strings.Join(arr, ","))
		qn = qn[0:idxpar]
	}
	entityName, qclassName := lastAndRemainingSegments(qn)
	className := lastSegment(qclassName)
	qclassName = strings.Replace(qclassName, ".", "_", -1)
	if kind == "[MT]" && className == entityName {
		kind = "[CS]"
	}
	return fmt.Sprintf("%v.java/[CN]/%v/%v/%v%v",
		qclassName, className, kind, entityName, args)
}

// JavaClassToHR converts a qualified java class name to a partial HR file name.
func JavaClassToHR(cn string) string {
	if cn == "" {
		return ""
	}
	return strings.Replace(cn, ".", "_", -1) + ".java/[CN]/"
}

// FileFromHR returns the file portion from a HR file name
func FileFromHR(hr string) string {
	idx := strings.Index(hr, "/[CN]/")
	if idx == -1 {
		return ""
	}
	return hr[:idx+6]
}

// HRToJava converts a HR file name to qualified java name
func HRToJava(hr string) string {
	panic("not implemented")
}

// RemoveGenerics removes from a given string any generic type specification
// of the form '<...>'
func RemoveGenerics(str string) string {
	re := regexp.MustCompile("<[^<>]*>")
	for {
		mm := re.FindAllString(str, -1)
		if len(mm) == 0 {
			break
		}
		for _, m := range mm {
			str = strings.Replace(str, m, "", -1)
		}
	}
	return str
}

func lastAndRemainingSegments(str string) (string, string) {
	idx := strings.LastIndex(str, ".")
	if idx == -1 {
		return str, ""
	}
	return str[idx+1:], str[0:idx]
}

func lastSegment(str string) string {
	last, _ := lastAndRemainingSegments(str)
	return last
}
