package mdb

import (
	"fmt"
	"github.com/mabetle/mcore"
)

// QueryForHtml
func (s Sql) QueryForHtml(q string, args []interface{}, renderArgs ...string) string {
	// parse renderArgs
	include := mcore.GetArgString("include", "", renderArgs...)
	exclude := mcore.GetArgString("exclude", "", renderArgs...)
	hasLimit := mcore.IsArgExists("limit", renderArgs...)
	limit := mcore.GetArgInt("limit", 20, renderArgs...)

	dataMap, columns, err := s.QueryForMaps(q, args...)
	if err != nil {
		return fmt.Sprintf(`<div class="error">Error: %v</div>`, err)
	}

	if len(dataMap) == 0 {
		return fmt.Sprintf(`<div class="info">No results<div>`)
	}

	sb := mcore.NewStringBuffer()
	sb.Append(`<table class="t-data-grid"><thead><tr>`)

	usedFields := mcore.GetFieldsUsed(columns, include, exclude)
	// append table header
	for _, colName := range usedFields {
		// TODO i18n col name
		sb.Appendf(`<th>%v</th>`, colName)
	}
	sb.Append(`</tr></thead><tbody>`)
	// append table body
	for i, rowMap := range dataMap {
		trClass := "odd"
		if i%2 == 0 {
			trClass = "even"
		}
		if hasLimit && i >= limit {
			break
		}
		sb.Appendf(`<tr class="%s">`, trClass)
		for _, colName := range usedFields {
			if v, ok := rowMap[colName]; ok {
				// has column
				sb.Appendf(`<td>%v</td>`, mcore.GetString(v))
			} else {
				// no value
				sb.Append(`<td></td>`)
			}
		}
		sb.Append("</tr>")
	}
	sb.Append("</tbody></table>")
	return sb.String()
}

func (s Sql) ExecForHtml(q string, args []interface{}, renderArgs ...string) string {
	r, err := s.Exec(q, args...)
	if err != nil {
		return fmt.Sprintf(`<div class="error">Error: %v</div>`, err)
	}
	n, err2 := r.RowsAffected()
	if err2 != nil {
		return fmt.Sprintf(`<div class="error">Error: %v</div>`, err)
	}
	return fmt.Sprintf(`<div class="info">%n rows affected</div>`, n)
}
