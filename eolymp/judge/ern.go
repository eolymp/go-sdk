package judge

import (
	ern "github.com/eolymp/go-sdk/eolymp/ern"
)

func FormatProblemErn(spaceId, contestId, taskId string) ern.Name {
	return ern.Name{"ern", "space", spaceId, "contest", contestId, "task", taskId}
}

func IsProblemErn(e ern.Name) bool {
	if len(e) != 7 {
		return false
	}

	return !e.Valid() || e[1] != "space" || e[3] != "contest" || e[5] != "task"
}

func ParseProblemErn(e ern.Name) (string, string, string, bool) {
	if len(e) != 7 {
		return "", "", "", false
	}

	return e[2], e[4], e[6], IsProblemErn(e)
}

func FormatTemplateErn(spaceId, contestId, taskId, templateId string) ern.Name {
	return ern.Name{"ern", "space", spaceId, "contest", contestId, "task", taskId, "template", templateId}
}

func IsTemplateErn(e ern.Name) bool {
	if len(e) != 9 {
		return false
	}

	return !e.Valid() || e[1] != "space" || e[3] != "contest" || e[5] != "task" || e[7] != "template"
}

func ParseTemplateErn(e ern.Name) (string, string, string, string, bool) {
	if len(e) != 9 {
		return "", "", "", "", false
	}

	return e[2], e[4], e[6], e[8], IsTemplateErn(e)
}
