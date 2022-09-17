package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/json"
	"github.com/spudtrooper/goutil/request"
)

func main() {
	flag.Parse()
	// Options
	printData := true
	printCookies := true
	printPayload := true

	// Data
	uri := request.MakeURL("https://www.opentable.com/s",
		request.Param{"dateTime", `2022-09-16T21%3A00%3A00`},
		request.Param{"term", `omakase`},
		request.Param{"originCorrelationId", `a148f506-ab44-41aa-8af7-a990ed75bbb9`},
		request.Param{"sortBy", `web_conversion`},
		request.Param{"originalTerm", `omakase`},
		request.Param{"intentModifiedTerm", `omakase`},
		request.Param{"queryUnderstandingType", `default`},
		request.Param{"corrid", `cf7a1115-801f-4eaa-8ed6-5892d8aa74a7`},
		request.Param{"covers", 2},
		request.Param{"latitude", 40.7683},
		request.Param{"longitude", -73.9802},
		request.Param{"metroId", 8},
	)
	cookie := [][2]string{
		{"_csrf", `58940096-283c-4221-bd07-4a4f9a0a6dfc`},
		{"otuvid", `2692F98A-649E-49AE-8BAA-2202F63221FE`},
		{"ha_userSession", `lastModified=2022-05-04T19%3A40%3A53.000Z&origin=prod-sc`},
		{"uCke", `lo=iEWpGfchjMfjeLPKOaZa1GmeZRuVGQhq&em=iEWpGfchjMfjeLPKOaZa1GmeZRuVGQhq&gpid=h831Ux0kPNvq3KPdmctzjg%3D%3D&gid=120054296152&l=1&t=1`},
		{"authCke", `atk=5b69a69b-6f09-4135-965f-87c2774d5844&rtk=ed0a1b78-c56e-44e8-8ffb-ef2c15f9d75d&tks=CONSUMER&tkt=bearer&eik=1209600&atke=2022-09-30T10%3A38%3A53.844Z`},
		{"OT-SessionId", `2244ba45-f8a7-415a-960d-a9cf6d3d9d19`},
		{"bm_mi", `5E8AED3015C4BD8E6AF46D7FEE6A36B1~YAAQv5cwF4r0nUODAQAAk7aaSxH7wQ/4ew6OnbVFVLEju6YWyHYX62g/LwIuY6X61EXEolEHLimfYpEFQe2vpmh/or5IJeq6i4J2yi/iMVfb/RSnrpI5pEnKCggjSd7XE15JV0zUFKsIaAPlnn31s6+PIBR53FyN5CHU7WKPO+vCh4MPCfLnJhlF+fc3TOeQ7+HncZfB0n26OvV8oZBpzLytJdV0/sRUahD2ZVenZFifsQ72hf1t4kVBQwquovVKe8og0radzNnLwRcZT71yYhioXi6iP0WzBeYPFVzRkaOu1thGaDZiJ6URMoEyLFjXR+Q=~1`},
		{"ak_bmsc", `8E6B87E6A75E1688EE782CDA4C4FEFA2~000000000000000000000000000000~YAAQv5cwFxD1nUODAQAACryaSxGUkwfsABS7u7++3Y3hHhEEu57zfthd5l6RT4RtBTUr2MAHf+6Kq1Fwjuf36UgXBtpj7LiMX2UrYrXLpX0l8j5QbDQSW4rU9K5N/LUXJYY332rWpR4GN5PEMjj2Md/zIbypfTe5e8kbkaO+kzy0uD3E/YCwgA/kSien4JXDu4d3VmV0ayTlGWY6nL+SG356VRcBOfRXuqtqvF8iirupJ51Aj6Y3E9sPjI5PpuKPSkLFyRbiwPMYZILVEKZ8YkBhBtuCIyspgMhPZnnxdi9FP9CuXUm7LyLay7JtX4WGwoBhJf9fCjnNwJh4bsD77koks/EtNF27iZwt6OB5A3pNwfrD2K496uGO+m6MVYWl//vQ20s7da8N1RwQFpedwUSvjmiZFiw45MmlNVilXrM=`},
		{"OT-Session-Update-Date", `1663420866`},
		{"bm_sv", `65C0DAD0503687D9135C84423E0E35DC~YAAQv5cwF9wXnkODAQAAVSGcSxHlDtaiiNWbEpALU8rI+RonTofZLLKlcXtJO8f92vN8ml7S/FiXUdaE0fbC+idj+Aq3DHPomqKK5u49nSkjvSjg9QZt518UgcpLNzNuhDu6kL+j3pWqfI6y5YdGUFrRD454waImJDahGDBoMdgEiui5oTvuEhSxhpnjID03B+N05LsOvsJDncJdowX/OUNp7aWW51AGSgn0W3UTPl7TYrIiHjMPFOfUDXThji5inMg5gw==~1`},
		{"ftc", url.QueryEscape(`x=2022-09-17T14:21:04&c=1&pt1=1&pt2=1&er=0&p1ca=s&p1q=dateTime=2022-09-16T21%3A00%3A00&covers=2&latitude=40.7683&longitude=-73.9802&term=omakase&originCor...`)},
		{"OptanonConsent", url.QueryEscape(`isIABGlobal=false&datestamp=Sat Sep 17 2022 09:21:05 GMT-0400 (Eastern Daylight Time)&version=6.20.0&hosts=&consentId=fcc0c9bc-633a-489c-b1b8-b4568b6ebc03&interactionCount=1&landingPath=NotLandingPage&groups=C0001:1,C0002:1,C0003:1,C0004:0,C0010:1&AwaitingReconsent=false`)},
	}
	headers := map[string]string{
		"authority":                 `www.opentable.com`,
		"accept":                    `text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9`,
		"accept-language":           `en-US,en;q=0.9`,
		"cache-control":             `no-cache`,
		"dnt":                       `1`,
		"pragma":                    `no-cache`,
		"referer":                   `https://www.opentable.com/`,
		"sec-ch-ua":                 `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`,
		"sec-ch-ua-mobile":          `?0`,
		"sec-ch-ua-platform":        `"macOS"`,
		"sec-fetch-dest":            `document`,
		"sec-fetch-mode":            `navigate`,
		"sec-fetch-site":            `same-origin`,
		"sec-fetch-user":            `?1`,
		"upgrade-insecure-requests": `1`,
		"user-agent":                `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36`,

		"cookie": request.CreateCookie(cookie),
	}

	// Make the request
	var payload interface{}
	var res *request.Response
	var err error

	res, err = request.Get(uri, &payload, request.RequestExtraHeaders(headers))

	if printData {
		fmt.Println(string(res.Data))
	}
	if printCookies {
		log.Printf("cookies: %v", res.Cookies)
	}
	if printPayload {
		log.Printf("payload: %s", json.MustColorMarshal(payload))
	}
	check.Err(err)
}
