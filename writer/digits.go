package writer

func ui64dc(n uint64) (bytes int) {
	if n < 1000000000 {
		if n < 100000 {
			if n < 100 {
				if n < 10 {
					return 1
				} else {
					return 2
				}
			} else if n < 10000 {
				if n < 1000 {
					return 3
				} else {
					return 4
				}
			} else {
				return 5
			}
		} else if n < 10000000 {
			if n < 1000000 {
				return 6
			} else {
				return 7
			}
		} else if n < 100000000 {
			return 8
		} else {
			return 9
		}
	} else if n < 100000000000000 {
		if n < 100000000000 {
			if n < 10000000000 {
				return 10
			} else {
				return 11
			}
		} else if n < 10000000000000 {
			if n < 1000000000000 {
				return 12
			} else {
				return 13
			}
		} else {
			return 14
		}
	} else if n < 10000000000000000 {
		if n < 1000000000000000 {
			return 15
		} else {
			return 16
		}
	} else if n < 1000000000000000000 {
		if n < 100000000000000000 {
			return 17
		} else {
			return 18
		}
	} else if n < 10000000000000000000 {
		return 19
	} else {
		return 20
	}
}

func ui32dc(n uint32) (bytes int) {
	if n < 100000 {
		if n < 100 {
			if n < 10 {
				return 1
			} else {
				return 2
			}
		} else if n < 10000 {
			if n < 1000 {
				return 3
			} else {
				return 4
			}
		} else {
			return 5
		}
	} else if n < 10000000 {
		if n < 1000000 {
			return 6
		} else {
			return 7
		}
	} else if n < 1000000000 {
		if n < 100000000 {
			return 8
		} else {
			return 9
		}
	} else {
		return 10
	}
}

func ui16dc(n uint16) (bytes int) {
	if n < 100 {
		if n < 10 {
			return 1
		} else {
			return 2
		}
	} else if n < 10000 {
		if n < 1000 {
			return 3
		} else {
			return 4
		}
	} else {
		return 5
	}
}

func ui8dc(n uint8) (bytes int) {
	if n < 10 {
		return 1
	} else if n < 100 {
		return 2
	} else {
		return 3
	}
}
