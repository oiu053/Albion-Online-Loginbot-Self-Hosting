package environment

func Token() (Token string) {
	return "Bot Token"
}
func ApplicationID() (AplicationID string) {
	return "ApplicationID"
}
func BotOwner() (BotOwnerID string) {
	return "Possible BotownerID"
}
func Delcommands() (delcommands bool) {
	return true
}

func BlacklistedGuildIDs() (BlacklistedGuildIDs []string) {

	return []string{"-"}
}
