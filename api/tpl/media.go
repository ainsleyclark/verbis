package tpl

// getMedia
//
// Obtains the media by ID and returns a domain.Media type
// or nil if not found or the ID parameter failed to be
// parsed.
//
// Example:
// {{ $image := media 10 }}
// {{ $image.Url }}
func (t *TemplateManager) getMedia(i interface{}) interface{} {
	var id int

	if i == nil {
		return nil
	}

	switch i.(type) {
	default:
		return nil
	case int:
		id = i.(int)
	case *int:
		p := i.(*int)
		if p != nil {
			id = *p
		}
	case float32:
		id = int(i.(float32))
	case *float32:
		p := i.(*float32)
		if p != nil {
			id = int(*p)
		}
	case float64:
		id = int(i.(float64))
	case *float64:
		p := i.(*float64)
		if p != nil {
			id = int(*p)
		}
	}

	m, err := t.store.Media.GetById(id)

	if err != nil {
		return nil
	}

	return m
}
