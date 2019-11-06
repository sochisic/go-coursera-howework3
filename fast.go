package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"user"
)

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	r := strings.NewReplacer("@", " [at] ")
	seenBrowsers := []string{}
	uniqueBrowsers := 0

	users := make([]user.User, 0)
	for scanner.Scan() {
		// fmt.Printf("%v %v\n", err, line)
		u := user.User{}
		err := u.UnmarshalJSON([]byte(scanner.Text()))
		if err != nil {
			panic(err)
		}

		users = append(users, u)
	}

	foundUsers := make([]string, len(users), len(users)*2)

	for i, user := range users {

		isAndroid := false
		isMSIE := false

		for _, browser := range user.Browsers {
			if ok := strings.Contains(browser, "Android"); ok != false {
				isAndroid = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}

			if ok := strings.Contains(browser, "MSIE"); ok != false {
				isMSIE = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		foundUsers = append(foundUsers, fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, r.Replace(user.Email)))
	}

	fmt.Fprintln(out, "found users:\n"+strings.Join(foundUsers, ""))
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}
