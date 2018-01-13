package inmem

import (
	"context"
	"errors"
	"fmt"

	"github.com/benbjohnson/tiny-ego"
)

// Ensure type implements interface.
var _ tinyego.WidgetService = &WidgetService{}

// WidgetService represents an in-memory implementation of WidgetService.
type WidgetService struct {
	db *DB
}

func (s *WidgetService) FindWidget(ctx context.Context, id int) (*tinyego.Widget, error) {
	s.db.RLock()
	defer s.db.RUnlock()
	return findWidget(ctx, s.db, id)
}

func (s *WidgetService) FindWidgets(ctx context.Context, filter tinyego.WidgetFilter, opt tinyego.FindOptions) ([]*tinyego.Widget, error) {
	s.db.RLock()
	defer s.db.RUnlock()
	return findWidgets(ctx, s.db, filter, opt)
}

func (s *WidgetService) CreateWidget(ctx context.Context, widget *tinyego.Widget) error {
	s.db.RLock()
	defer s.db.RUnlock()
	return createWidget(ctx, s.db, widget)
}

func (s *WidgetService) UpdateWidget(ctx context.Context, id int, upd tinyego.WidgetUpdate) (*tinyego.Widget, error) {
	s.db.RLock()
	defer s.db.RUnlock()
	return updateWidget(ctx, s.db, id, upd)
}

func findWidget(ctx context.Context, db *DB, id int) (*tinyego.Widget, error) {
	w := db.widgets[id]
	if w == nil {
		return nil, errors.New("widget not found")
	}
	other := *w
	return &other, nil
}

func findWidgets(ctx context.Context, db *DB, filter tinyego.WidgetFilter, opt tinyego.FindOptions) ([]*tinyego.Widget, error) {
	var a []*tinyego.Widget
	for _, w := range db.widgets {
		if w.Match(filter) {
			a = append(a, w)
		}
	}

	tinyego.SortWidgets(a, tinyego.WidgetLess(opt.SortBy))
	return a, nil
}

func createWidget(ctx context.Context, db *DB, widget *tinyego.Widget) error {
	widget.ID = db.nextSeq("widgets")
	return saveWidget(ctx, db, widget)
}

func updateWidget(ctx context.Context, db *DB, id int, upd tinyego.WidgetUpdate) (*tinyego.Widget, error) {
	widget, err := findWidget(ctx, db, id)
	if err != nil {
		return nil, err
	}

	// Update fields.
	if upd.Name != nil {
		widget.Name = *upd.Name
	}
	if upd.Size != nil {
		widget.Size = *upd.Size
	}

	if err := saveWidget(ctx, db, widget); err != nil {
		return widget, err
	}
	return widget, nil
}

func saveWidget(ctx context.Context, db *DB, widget *tinyego.Widget) error {
	if widget.Name == "" {
		return errors.New("widget name required")
	} else if widget.Size == "" {
		return errors.New("widget size required")
	} else if tinyego.IsValidWidgetSize(widget.Size) {
		return fmt.Errorf("invalid widget size: %s", widget.Size)
	}

	db.widgets[widget.ID] = widget
	return nil
}
