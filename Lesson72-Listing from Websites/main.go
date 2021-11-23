package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	loopCtr := 0
	totalCases := make([]CustomerCaseStudy, 0)

	for {
		ctrString := strconv.Itoa(loopCtr)
		customerCases, moreResults, err := setPayload(ctrString)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("Page:%v\tMoreResults:%v\n", ctrString, moreResults)

		if len(customerCases) < 1 {
			fmt.Printf("Page number: %s. %s\n", ctrString, "No more customer cases")
		}

		totalCases = append(totalCases, customerCases...)
		loopCtr++
	}

	allCustomers := &AllCustomers{
		CustomerCases: totalCases,
	}

	if output, err := json.Marshal(allCustomers); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("%s\n", output)
	}
}

func setPayload(pageID string) (customerCases []CustomerCaseStudy, moreResults bool, err error) {
	moreResults = false
	data := Payload{
		// fill struct
		Text: "",
		// "facet_filters":[
		// 		{"facet":"story_product_categories","values":["Cloud Platform","Azure"]},
		// 		{"facet":"story_country_region","values":["North America"]},
		// 		{"facet":"story_country","values":["North America/United States"]}]
		FacetFilters: []FacetFilters{
			{
				Facet:  "story_product_categories",
				Values: []string{"Cloud Platform", "Azure"},
			},
			{
				Facet:  "story_country_region",
				Values: []string{"North America"},
			},
			{
				Facet:  "story_country",
				Values: []string{"North America/United States"},
			},
		},
		// "related_documents":[],
		RelatedDocuments: []interface{}{},
		// "page_id":"{{{page_id}}}",
		// PageID: "0",
		// "featured_sections":null,
		// FeaturedSections: interface{},
		// "sort_mode":"story_publish_date desc"
		SortMode: "story_publish_date desc",
	}

	data.PageID = pageID

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return []CustomerCaseStudy{}, moreResults, err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://customers.microsoft.com/en-us/api/search", body)
	if err != nil {
		return []CustomerCaseStudy{}, moreResults, err
	}

	setReqHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []CustomerCaseStudy{}, moreResults, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []CustomerCaseStudy{}, moreResults, err
	}

	var jsonResp JSONResponse
	err = json.Unmarshal(respBody, &jsonResp)
	if err != nil {
		return []CustomerCaseStudy{}, moreResults, err
	}

	moreResults = jsonResp.SearchMetadata.MoreResults

	// Is the results array empty?
	if len(jsonResp.SearchResult.Results) < 1 {
		return []CustomerCaseStudy{}, moreResults, fmt.Errorf("%s", "Zero results fetched")
	}
	customerCases = make([]CustomerCaseStudy, 0)

	for _, result := range jsonResp.SearchResult.Results {

		// result := resultInterface.(Result)
		customerCases = append(customerCases, CustomerCaseStudy{
			Score:        fmt.Sprintf("%s", result.Score),
			ID:           result.Document.ID,
			CustomerName: result.Document.StoryCustomerName[0],
			Industry:     result.Document.StoryIndustry[0],
			Headline:     result.Document.StoryHeadline,
		})

	}

	return customerCases, moreResults, nil
}

func setReqHeaders(req *http.Request) {
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"95\", \"Chromium\";v=\"95\", \";Not A Brand\";v=\"99\"")
	req.Header.Set("Correlation-Id", "a77eacaa-3f80-4f91-a1de-a6c0b8ba7402")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cam-Language", "en")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Request-Context", "appId=cid-v1:dc8f084b-9ca1-4546-a7f2-693abd4999be")
	req.Header.Set("Request-Id", "|FUUqs.lGDkf")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
	req.Header.Set("Origin", "https://customers.microsoft.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://customers.microsoft.com/en-us/search?sq=&ff=story_country_region%26%3ENorth%20America&p=0")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cookie", "MC1=GUID=190fd34c9a6a4a72889f1334bd1f7b23&HASH=190f&LV=202111&V=4&LU=1636494599703; MUID=3DCDE67FD13A60592D6EF695D0BE613F; _mkto_trk=id:157-GQE-382&token:_mch-microsoft.com-1636494631678-79705; fptctx2=H3ihr9e92IdW6yd1ZgQ9SxLzXxHcL2CcU%252fZDGCdp0wEB6GmzCVvTzqkwWOBV0NcoZ3QRjZDxNv1f0Q6tZ%252fxtKA09dcDAQjEPWm6K790ZxwFLc2zAwhIVsHmlk3QhRr60jHKHkn3yS6xjcP2CenghxboMEf0T1%252fJRZDHa2dGwIxSi5Qww4UWnRaA5nAgJYXRab04CW1NFOBJa2SiqMxrrUz2CNVcqJpL5yc40hb4plIIN24whDL2asUz6SqlDx6teQL5ukiiYZdKQFxsTa72U9q62Z3VRZv0orfnGXYjv1Hw%253d; MSCC=NR; at_check=true; AMCVS_EA76ADE95776D2EC7F000101%40AdobeOrg=1; AMCV_EA76ADE95776D2EC7F000101%40AdobeOrg=1585540135%7CMCIDTS%7C18954%7CMCMID%7C44065314491019843263426956050288906803%7CMCAAMLH-1638210289%7C9%7CMCAAMB-1638210289%7CRKhpRz8krg2tLO6pguXWp5olkAcUniQYPHaMWWgdJ3xzPWQmdj0y%7CMCCIDH%7C-1367273948%7CMCOPTOUT-1637612689s%7CNONE%7CMCAID%7C30C566119E3848F5-60001FDA8E9E7889%7CvVersion%7C4.4.0; _cs_c=0; mbox=session#c018a3fa87ce4ef4b9e4c18aa255244d#1637607349|PC#c018a3fa87ce4ef4b9e4c18aa255244d.35_0#1671792188; _cs_id=d7ea1e62-948e-a324-9c55-3a0e67c02430.1637605489.1.1637605489.1637605489.1613561419.1671769489249; _CT_RS_=Recording; ARRAffinity=d4d1188bacc68bf5d612c7eb0d13ac0b5e816d9ed517bcfe96844c0e8a0bdc3e; ARRAffinitySameSite=d4d1188bacc68bf5d612c7eb0d13ac0b5e816d9ed517bcfe96844c0e8a0bdc3e; CustomersMicrosoftLocale=en-us; MicrosoftApplicationsTelemetryDeviceId=e938651c-ae8f-4245-a188-5982429bb229; ai_user=UzzaF|2021-11-22T18:24:54.471Z; MSFPC=GUID=190fd34c9a6a4a72889f1334bd1f7b23&HASH=190f&LV=202111&V=4&LU=1636494599703; _abck=ACCA2AD097A6F2D7365A018928A2F51E~0~YAAQXzUauCx9mjF9AQAAnpT5SQaoto8GK4tZgdeKaX/o5OJ3jeUJgS9BiEjMiRvBj9JfvGwiJRx9p+BhDqj+9UdXKNdANJDRp1uaomP+3S3gVygiGVkThg4kC0GQnKFfncNGFIkOT9Uy27muV+Tff5j2sJAfyVpLcL5f9HacmSgv1YeFCX6+bNqG+tC0d2/dJu2VkHHhsbmQt5FKut3shyzrQuRd67CAA4cHBKJSD6gGcir40j7RWBri8QrGHeww+ljz6KDH3eGPDHmgUlahMeHlrA59kwTa+DUAEMPjqt3guGJugxf8KJJdtRtB+FEZx7yjibSctqdXDbxVFWZm5F7OY6sBIUoXoBiYEaaiEoNbA6c4LxInpaItMSkEYm4NGirsRpyqFxbhekanmTAt8n2yLODt8YBxktys~-1~-1~-1; ak_bmsc=C2E44D1D4F2C222036586E9FAF5D27D0~000000000000000000000000000000~YAAQXzUauC19mjF9AQAAnpT5SQ29jvtYKkhDY2F/oYxKtjqgdCfprekI9qwSdraw8bM/wgC1wXXQvtq2cVB+RYo/TPW4Eb3EmK627bJYk42wLv+mEHurwW/9UlXFwkhQlKv3Oy8rJdFJT4Ww8Tiv1u6NbXazSZxr29W+E0TZ12lM9oc26TI45ng1CalO5mQL/4N65mv9A5f1xw0WtCeS24/Z+YENcjd7wGcA0XzXUBrWb9ZXt51FHrlbKKjiB6LXSYNLKwYXXwi0n8LQwpin7PDlUT0+icK2VHobkGd+E+XNJuZMcCjwiTIxqKibGLQPVC01TxRUo9lZaulajo5Q73XNaksfYcS3AhClwu1TW8f/eBeQ5HY3xjF831PgEbJSx9P4sWFJXYAAM5E=; bm_sz=E65CF8F4A9A4A9FE32F5A16053362E4D~YAAQXzUauC59mjF9AQAAnpT5SQ1WBuMaPAfnE1/2UDsQWry8AMctfFIUqm71adAz1NtT4txhY81LuPytUQRyff0k1EJfq9Y5dFidUFuUZqnXw+FMmwwL531669yadg8AMAq3V2QLD0+VDcJkQhoBYpvxeK7SZ5+vGNrWb2Gwc0Sd9gvtZ8xaYTR8Gcg5rjIKS1COSuOT1D8Gb/+Wy5MQK40Wq7qVPVs4kuYYTEkjXeyti25ZO5CAF2a8NFxCkOxBhGFElYFrVvXPUVVBp2jlLSpDkIr4EA6uQ65eoC8FbXKxcNxyRUs=~3682869~4536129; __CT_Data=gpv=3&ckp=tld&dm=microsoft.com&apv_1067_www32=1&cpv_1067_www32=1&rpv_1067_www32=1&apv_32260_www07=2&cpv_32260_www07=2; MS0=2bcc4428d5ff4981bc77d946a16730cf; ai_session=mfYCFod+Gfs30zsOyMRVkK|1637623632251|1637624015552.7")

}
