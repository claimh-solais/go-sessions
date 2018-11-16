package session

// SetSession comment
func (ctx *MemoryStore) SetSession(sess *Session) {
	ctx.Store.SetSession(sess)
}
