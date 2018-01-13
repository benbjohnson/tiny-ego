package tinyego

import (
	"context"
	"sort"
	"strings"

	"go4.org/strutil"
)

type Widget struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Size string `json:"size"`
}

func (w *Widget) Match(filter WidgetFilter) bool {
	if filter.Query != nil && !strutil.ContainsFold(w.Name, *filter.Query) {
		return false
	} else if filter.Size != nil && w.Size != *filter.Size {
		return false
	}
	return true
}

type WidgetUpdate struct {
	Name *string `json:"name"`
	Size *string `json:"size"`
}

type WidgetFilter struct {
	Query *string `json:"name"`
	Size  *string `json:"size"`
}

type WidgetService interface {
	FindWidget(ctx context.Context, id int) (*Widget, error)
	FindWidgets(ctx context.Context, filter WidgetFilter) ([]*Widget, error)
	CreateWidget(ctx context.Context, widget *Widget) error
	UpdateWidget(ctx context.Context, id int, upd WidgetUpdate) (*Widget, error)
}

type WidgetLessFunc func(i, j *Widget) bool

func SortWidgets(a []*Widget, fn WidgetLessFunc) {
	sort.Slice(a, func(i, j int) bool { return fn(a[i], a[j]) })
}

func WidgetLess(sortBy string) WidgetLessFunc {
	switch sortBy {
	case "size":
		return WidgetLessSize
	default:
		return WidgetLessName
	}
}

func WidgetLessName(i, j *Widget) bool {
	return strings.ToLower(i.Name) < strings.ToLower(j.Name)
}

func WidgetLessSize(i, j *Widget) bool {
	return strings.ToLower(i.Size) < strings.ToLower(j.Size)
}

const (
	WidgetSizeSmall  = "small"
	WidgetSizeMedium = "medium"
	WidgetSizeLarge  = "large"
)

func IsValidWidgetSize(v string) bool {
	switch v {
	case WidgetSizeSmall, WidgetSizeMedium, WidgetSizeLarge:
		return true
	default:
		return false
	}
}
