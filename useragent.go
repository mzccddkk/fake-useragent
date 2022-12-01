package useragent

func Random() string {
	return ua.getAll().random()
}

func Chrome() string {
	return ua.get(chrome).random()
}

func Edge() string {
	return ua.get(edge).random()
}

func IE() string {
	return ua.get(ie).random()
}

func Firefox() string {
	return ua.get(firefox).random()
}

func Safari() string {
	return ua.get(safari).random()
}

func UC() string {
	return ua.get(uc).random()
}

func Opera() string {
	return ua.get(opera).random()
}

func Android() string {
	return ua.get(android).random()
}

func IOS() string {
	return ua.get(ios).random()
}

func Windows() string {
	return ua.get(windows).random()
}

func Linux() string {
	return ua.get(linux).random()
}

func MacOS() string {
	return ua.get(macOs).random()
}

func MacOSX() string {
	return ua.get(macOsX).random()
}

func IPhone() string {
	return ua.get(iphone).random()
}

func IPad() string {
	return ua.get(ipad).random()
}

func Mobile() string {
	return ua.get(mobile).random()
}

func Computer() string {
	return ua.get(computer).random()
}
