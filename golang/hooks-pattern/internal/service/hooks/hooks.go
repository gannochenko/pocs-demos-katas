package hooks

import "context"

type Handler = func(ctx context.Context, args interface{})

type Service struct {
	handlers map[string][]Handler
}

func (s *Service) On(eventName string, handler Handler) {
	if s.handlers == nil {
		s.handlers = make(map[string][]Handler)
	}

	if _, ok := s.handlers[eventName]; !ok {
		s.handlers[eventName] = make([]Handler, 0)
	}

	s.handlers[eventName] = append(s.handlers[eventName], handler)
}

func (s *Service) Trigger(ctx context.Context, eventName string, args interface{}) {
	if s.handlers == nil {
		return
	}

	if handlers, ok := s.handlers[eventName]; ok {
		for _, handler := range handlers {
			handler(ctx, args)
		}
	}
}
