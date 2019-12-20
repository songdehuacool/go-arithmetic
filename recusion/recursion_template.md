```
func recur(level int, param int) {
	// terminator
	if (level > MAX_LEVEL) {
		return
	}

	// process current logic
	process(level, param)

	// drill down
	recur(level, newParam)

	// restore current status
}
```