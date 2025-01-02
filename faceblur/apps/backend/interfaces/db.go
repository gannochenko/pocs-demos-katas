package interfaces

type SessionHandle interface {
}

type SessionManager interface {
	Begin() (*SessionHandle, error)
}
