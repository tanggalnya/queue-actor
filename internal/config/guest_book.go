package config

type GuestBookCfg struct {
	Username string
	Password string
}

var guestBook GuestBookCfg

func GuestBookConfig() GuestBookCfg {
	return guestBook
}

func initGuestBookConfig() {
	guestBook = GuestBookCfg{
		Username: getStringOrPanic("GUEST_BOOK_USERNAME"),
		Password: getStringOrPanic("GUEST_BOOK_PASSWORD"),
	}
}
