package main

// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
// curl 'https://customers.microsoft.com/en-us/api/search' \
// -H 'Connection: keep-alive' \
// -H 'sec-ch-ua: "Google Chrome";v="95", "Chromium";v="95", ";Not A Brand";v="99"' \
// -H 'Correlation-Id: 47f3353c-81ba-48f3-8e5b-1a37529e3c7b' \
// -H 'sec-ch-ua-mobile: ?0' \
// -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36' \
// -H 'Content-Type: application/json' \
// -H 'Cam-Language: en' \
// -H 'Accept: application/json, text/plain, */*' \
// -H 'Request-Context: appId=cid-v1:dc8f084b-9ca1-4546-a7f2-693abd4999be' \
// -H 'Request-Id: |8TkxG.bUtBw' \
// -H 'sec-ch-ua-platform: "macOS"' \
// -H 'Origin: https://customers.microsoft.com' \
// -H 'Sec-Fetch-Site: same-origin' \
// -H 'Sec-Fetch-Mode: cors' \
// -H 'Sec-Fetch-Dest: empty' \
// -H 'Referer: https://customers.microsoft.com/en-us/search?sq=&ff=story_product_categories%26%3ECloud%20Platform%26%26story_country_region%26%3ENorth%20America%26%26story_country%26%3ENorth%20America%2FUnited%20States&p=0&so=story_publish_date%20desc' \
// -H 'Accept-Language: en-US,en;q=0.9' \
// -H 'Cookie: ARRAffinity=8911fa7119f3e84cbf06dd019555fbae5d731650b2cd43ad675dbfc0b7f6b52e; ARRAffinitySameSite=8911fa7119f3e84cbf06dd019555fbae5d731650b2cd43ad675dbfc0b7f6b52e; bm_sz=9E4A955C769F6FBF139BF0824D45FFB3~YAAQXzUauPsDkzF9AQAAEnkGSQ3r3qkTaf8cicgWGGecO6J5t/UBaJcitv8ExUkKSEvmSLXrFVuzlvaIsS7cPamXrlgXGUnDImdbbLn631w4ZaIo2al4mgJS6MfoY+8I+ZkTwrO6fguqKdM+0QA1Lk/W9cd4nmYpj1zBZ/qIB2vFXVQ1MCXIzl8FPMYJG8niEHO42nOpvZqL7tLmnNKjZxm0vaaTRrQXDrmxY9SKoMwuDw08kwCi67w1vWxFV/B03oHHVlNX72gUmrCQeL+FIyKQMnfELHq6ez9aTreW+hcRYgca18s=~3291187~3752246; CustomersMicrosoftLocale=en-us; MicrosoftApplicationsTelemetryDeviceId=179e4471-8aeb-4da4-9b6c-08cbd8fa85ad; MSCC=NR; ak_bmsc=A7E4C7730F8FC034C42C2D30E6549D04~000000000000000000000000000000~YAAQXzUauAYFkzF9AQAADIMGSQ0Mifnr3aIY0G0drmHJ/Dbd9oO2M8ceCf4TOaA3A7MR4idC2gMn9dHjiVC+nwyxUpbLtjEHRRxQeIqLnaF7lsCnEU5cpOjaaxX8S8Z50PObz9tnupK00GPGus2i9qBnOKkI/xtxT9w6IYjmb8GZC07aBRpPhGwZad/Kn26UvUUc/twu26Rm/I00LxI9QSp8Uzrl5eIa+NZgVyq/8euyvMvq4gupH5wGkDe9NyVx1NfeSZwAUnLHmKIzKK6nXtvwzZwGiE6AkmUZRqbZFbmwr7XKVGKHtRcDWCWK5lepzA7ULi+dqrk2JK5ODz+qacddPHI1fnGC7DCCwkuY1Xne2lgy3UHohvsNTYp6Sgk255LYxRiAp59Jr0E=; ai_user=ZotC5|2021-11-22T19:01:43.314Z; MC1=GUID=946253c2005b4d91b553046e8c36bb95&HASH=9462&LV=202111&V=4&LU=1637607703440; MS0=e4b75ac5188844319dfcdc85268f19e6; _abck=B63C411A7A330B6F9FDD7B9A14F4765F~0~YAAQXzUauNAJkzF9AQAAFLYGSQZoY/cMjDvPvXrl/keU9/0VfOWhaXDTKs9WjtcRkQHrhjBTaxEdQjsNvAg08NcyoVMj6JKjKOTP+B72t8L3aNrsOIipkuQxrejBifTLWAsKajk1TmQKzPKu7kXEwOz/SMJ5Wrb07kXl98KvmeDbm1GwbG5vjHhYBaRGHIcfUg6kILHyLOIdunt2FCrcSaj4kVs5gdhcFeBnullBL0gXQ0i5FkwwVNhIVLXtoDpQ2MSne2nAgtkfQ2wf75ux6Js+MQ5apxnNLMpuc/NWOaoYEdvZD2fyVwC/MNfpMsa85cqn5+YhPbwIIXMJVMEd6TIC/ZLqFa1iq4HjoVpdz94kqhoOms6fXxMBgWlLMk1BBAc4QkYRzhHvxbxLAEsC+idvC+umCDVKq9ya~-1~||-1||~-1; MSFPC=GUID=946253c2005b4d91b553046e8c36bb95&HASH=9462&LV=202111&V=4&LU=1637607703440; __CT_Data=gpv=5&ckp=tld&dm=microsoft.com&apv_32260_www07=5&cpv_32260_www07=5; ai_session=/4+E8sUIzCDuYnK5wKPWYO|1637607702971|1637608847712.4' \
// --data-raw '{"text":"", "facet_filters":[{"facet":"story_product_categories","values":["Cloud Platform","Azure"]},{"facet":"story_country_region","values":["North America"]},{"facet":"story_country","values":["North America/United States"]}],"related_documents":[],"page_id":"{{{page_id}}}","featured_sections":null,"sort_mode":"story_publish_date desc"}' \
// --compressed

type Payload struct {
	Text             string         `json:"text,omitempty"`
	FacetFilters     []FacetFilters `json:"facet_filters,omitempty"`
	RelatedDocuments []interface{}  `json:"related_documents,omitempty"`
	PageID           string         `json:"page_id"`
	FeaturedSections interface{}    `json:"featured_sections,omitempty"`
	SortMode         string         `json:"sort_mode"`
}

type FacetFilters struct {
	Facet  string   `json:"facet,omitempty"`
	Values []string `json:"values,omitempty"`
}

type JSONResponse struct {
	SearchResult struct {
		Count    interface{} `json:"Count"`
		Coverage interface{} `json:"Coverage"`
		Facets   struct {
			StoryIndustry []struct {
				Type  int         `json:"Type"`
				From  interface{} `json:"from"`
				To    interface{} `json:"to"`
				Value string      `json:"value"`
				Count int         `json:"count"`
			} `json:"story_industry"`
			StoryProductFamilies []struct {
				Type  int         `json:"Type"`
				From  interface{} `json:"from"`
				To    interface{} `json:"to"`
				Value string      `json:"value"`
				Count int         `json:"count"`
			} `json:"story_product_families"`
			StoryCountry []struct {
				Type  int         `json:"Type"`
				From  interface{} `json:"from"`
				To    interface{} `json:"to"`
				Value string      `json:"value"`
				Count int         `json:"count"`
			} `json:"story_country"`
			StoryCountryRegion []struct {
				Type  int         `json:"Type"`
				From  interface{} `json:"from"`
				To    interface{} `json:"to"`
				Value string      `json:"value"`
				Count int         `json:"count"`
			} `json:"story_country_region"`
			Language []struct {
				Type  int         `json:"Type"`
				From  interface{} `json:"from"`
				To    interface{} `json:"to"`
				Value string      `json:"value"`
				Count int         `json:"count"`
			} `json:"language"`
			StoryOrganizationSize []struct {
				Type  int         `json:"Type"`
				From  interface{} `json:"from"`
				To    interface{} `json:"to"`
				Value string      `json:"value"`
				Count int         `json:"count"`
			} `json:"story_organization_size"`
			StoryIndustryFriendlyname []struct {
				Type  int         `json:"Type"`
				From  interface{} `json:"from"`
				To    interface{} `json:"to"`
				Value string      `json:"value"`
				Count int         `json:"count"`
			} `json:"story_industry_friendlyname"`
			StoryProductCategories []struct {
				Type  int         `json:"Type"`
				From  interface{} `json:"from"`
				To    interface{} `json:"to"`
				Value string      `json:"value"`
				Count int         `json:"count"`
			} `json:"story_product_categories"`
			Locale []struct {
				Type  int         `json:"Type"`
				From  interface{} `json:"from"`
				To    interface{} `json:"to"`
				Value string      `json:"value"`
				Count int         `json:"count"`
			} `json:"locale"`
			Media []struct {
				Type  int         `json:"Type"`
				From  interface{} `json:"from"`
				To    interface{} `json:"to"`
				Value string      `json:"value"`
				Count int         `json:"count"`
			} `json:"media"`
		} `json:"Facets"`
		Results           []Result    `json:"Results"`
		ContinuationToken interface{} `json:"ContinuationToken"`
	} `json:"search_result"`
	CamIndex struct {
		Name   string `json:"name"`
		Fields []struct {
			Name             string        `json:"name"`
			DataType         string        `json:"data_type"`
			Key              bool          `json:"key"`
			LegacyKey        bool          `json:"legacy_key"`
			Searchable       bool          `json:"searchable"`
			Filterable       bool          `json:"filterable"`
			Facetable        bool          `json:"facetable"`
			Hierarchical     bool          `json:"hierarchical"`
			Sortable         bool          `json:"sortable"`
			InPreview        bool          `json:"in_preview"`
			Title            bool          `json:"title"`
			Feeds            []string      `json:"feeds,omitempty"`
			ChildFields      []string      `json:"child_fields,omitempty"`
			Leaf             bool          `json:"leaf,omitempty"`
			Root             bool          `json:"root,omitempty"`
			RootField        string        `json:"root_field,omitempty"`
			ParentFieldChain []interface{} `json:"parent_field_chain,omitempty"`
			ParentField      string        `json:"parent_field,omitempty"`
		} `json:"fields"`
		Suggesters []struct {
			Name         string   `json:"name"`
			SourceFields []string `json:"sourceFields"`
		} `json:"suggesters"`
		Locales struct {
			SupportedCultures []string `json:"supported_cultures"`
		} `json:"locales"`
		SearchPageSize      int `json:"search_page_size"`
		FacetListSize       int `json:"facet_list_size"`
		AutoSuggestListSize int `json:"auto_suggest_list_size"`
		ScoringProfiles     []struct {
			Name string `json:"name"`
			Text struct {
				Weights struct {
					StoryProductsServices     interface{} `json:"story_products_services"`
					StoryIndustryFriendlyname interface{} `json:"story_industry_friendlyname"`
				} `json:"weights"`
			} `json:"text"`
			Functions []struct {
				Type          string `json:"type"`
				FieldName     string `json:"field_name"`
				Boost         int    `json:"boost"`
				Interpolation string `json:"interpolation"`
			} `json:"functions"`
		} `json:"scoring_profiles"`
		FeedOptions []struct {
			FeedName string `json:"feed_name"`
			PageSize int    `json:"page_size"`
			Orderby  string `json:"orderby"`
		} `json:"feed_options"`
	} `json:"cam_index"`
	RelatedDocuments interface{} `json:"related_documents"`
	SearchMetadata   struct {
		Query struct {
			Index        string      `json:"index"`
			Text         string      `json:"text"`
			Key          interface{} `json:"key"`
			PageID       string      `json:"page_id"`
			BatchCount   int         `json:"batch_count"`
			FacetFilters []struct {
				Facet    string      `json:"facet"`
				Values   []string    `json:"values"`
				Operator interface{} `json:"operator"`
			} `json:"facet_filters"`
			RelatedDocuments []interface{} `json:"related_documents"`
			FeaturedSections interface{}   `json:"featured_sections"`
			SortMode         string        `json:"sort_mode"`
		} `json:"query"`
		MoreResults       bool        `json:"more_results"`
		CompletedIn       float64     `json:"completed_in"`
		CompletedSearchIn float64     `json:"completed_search_in"`
		CompletedIoIn     interface{} `json:"completed_io_in"`
	} `json:"search_metadata"`
}

type Result struct {
	Score      interface{} `json:"Score"`
	Highlights interface{} `json:"Highlights"`
	Document   struct {
		ID                        string        `json:"id"`
		Locale                    string        `json:"locale"`
		TemplateName              string        `json:"template_name"`
		Media                     []interface{} `json:"media"`
		StorySearchResultsImage   string        `json:"story_search_results_image"`
		StoryHeadline             string        `json:"story_headline"`
		StoryIndustry             []string      `json:"story_industry"`
		StoryIndustryFriendlyname []string      `json:"story_industry_friendlyname"`
		StoryCustomerName         []string      `json:"story_customer_name"`
	} `json:"Document"`
}

type CustomerCaseStudy struct {
	Score        string `json:"score,omitempty"`
	ID           string `json:"id,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
	Headline     string `json:"headline,omitempty"`
	Industry     string `json:"industry,omitempty"`
}

type AllCustomers struct {
	CustomerCases []CustomerCaseStudy `json:"customerCases,omitempty"`
}

// Refer to {{{page_id}}} in --data-raw as a filter to the post method
const curlStatement = `curl 'https://customers.microsoft.com/en-us/api/search' \
-H 'Connection: keep-alive' \
-H 'sec-ch-ua: "Google Chrome";v="95", "Chromium";v="95", ";Not A Brand";v="99"' \
-H 'Correlation-Id: 47f3353c-81ba-48f3-8e5b-1a37529e3c7b' \
-H 'sec-ch-ua-mobile: ?0' \
-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36' \
-H 'Content-Type: application/json' \
-H 'Cam-Language: en' \
-H 'Accept: application/json, text/plain, */*' \
-H 'Request-Context: appId=cid-v1:dc8f084b-9ca1-4546-a7f2-693abd4999be' \
-H 'Request-Id: |8TkxG.bUtBw' \
-H 'sec-ch-ua-platform: "macOS"' \
-H 'Origin: https://customers.microsoft.com' \
-H 'Sec-Fetch-Site: same-origin' \
-H 'Sec-Fetch-Mode: cors' \
-H 'Sec-Fetch-Dest: empty' \
-H 'Referer: https://customers.microsoft.com/en-us/search?sq=&ff=story_product_categories%26%3ECloud%20Platform%26%26story_country_region%26%3ENorth%20America%26%26story_country%26%3ENorth%20America%2FUnited%20States&p=0&so=story_publish_date%20desc' \
-H 'Accept-Language: en-US,en;q=0.9' \
-H 'Cookie: ARRAffinity=8911fa7119f3e84cbf06dd019555fbae5d731650b2cd43ad675dbfc0b7f6b52e; ARRAffinitySameSite=8911fa7119f3e84cbf06dd019555fbae5d731650b2cd43ad675dbfc0b7f6b52e; bm_sz=9E4A955C769F6FBF139BF0824D45FFB3~YAAQXzUauPsDkzF9AQAAEnkGSQ3r3qkTaf8cicgWGGecO6J5t/UBaJcitv8ExUkKSEvmSLXrFVuzlvaIsS7cPamXrlgXGUnDImdbbLn631w4ZaIo2al4mgJS6MfoY+8I+ZkTwrO6fguqKdM+0QA1Lk/W9cd4nmYpj1zBZ/qIB2vFXVQ1MCXIzl8FPMYJG8niEHO42nOpvZqL7tLmnNKjZxm0vaaTRrQXDrmxY9SKoMwuDw08kwCi67w1vWxFV/B03oHHVlNX72gUmrCQeL+FIyKQMnfELHq6ez9aTreW+hcRYgca18s=~3291187~3752246; CustomersMicrosoftLocale=en-us; MicrosoftApplicationsTelemetryDeviceId=179e4471-8aeb-4da4-9b6c-08cbd8fa85ad; MSCC=NR; ak_bmsc=A7E4C7730F8FC034C42C2D30E6549D04~000000000000000000000000000000~YAAQXzUauAYFkzF9AQAADIMGSQ0Mifnr3aIY0G0drmHJ/Dbd9oO2M8ceCf4TOaA3A7MR4idC2gMn9dHjiVC+nwyxUpbLtjEHRRxQeIqLnaF7lsCnEU5cpOjaaxX8S8Z50PObz9tnupK00GPGus2i9qBnOKkI/xtxT9w6IYjmb8GZC07aBRpPhGwZad/Kn26UvUUc/twu26Rm/I00LxI9QSp8Uzrl5eIa+NZgVyq/8euyvMvq4gupH5wGkDe9NyVx1NfeSZwAUnLHmKIzKK6nXtvwzZwGiE6AkmUZRqbZFbmwr7XKVGKHtRcDWCWK5lepzA7ULi+dqrk2JK5ODz+qacddPHI1fnGC7DCCwkuY1Xne2lgy3UHohvsNTYp6Sgk255LYxRiAp59Jr0E=; ai_user=ZotC5|2021-11-22T19:01:43.314Z; MC1=GUID=946253c2005b4d91b553046e8c36bb95&HASH=9462&LV=202111&V=4&LU=1637607703440; MS0=e4b75ac5188844319dfcdc85268f19e6; _abck=B63C411A7A330B6F9FDD7B9A14F4765F~0~YAAQXzUauNAJkzF9AQAAFLYGSQZoY/cMjDvPvXrl/keU9/0VfOWhaXDTKs9WjtcRkQHrhjBTaxEdQjsNvAg08NcyoVMj6JKjKOTP+B72t8L3aNrsOIipkuQxrejBifTLWAsKajk1TmQKzPKu7kXEwOz/SMJ5Wrb07kXl98KvmeDbm1GwbG5vjHhYBaRGHIcfUg6kILHyLOIdunt2FCrcSaj4kVs5gdhcFeBnullBL0gXQ0i5FkwwVNhIVLXtoDpQ2MSne2nAgtkfQ2wf75ux6Js+MQ5apxnNLMpuc/NWOaoYEdvZD2fyVwC/MNfpMsa85cqn5+YhPbwIIXMJVMEd6TIC/ZLqFa1iq4HjoVpdz94kqhoOms6fXxMBgWlLMk1BBAc4QkYRzhHvxbxLAEsC+idvC+umCDVKq9ya~-1~||-1||~-1; MSFPC=GUID=946253c2005b4d91b553046e8c36bb95&HASH=9462&LV=202111&V=4&LU=1637607703440; __CT_Data=gpv=5&ckp=tld&dm=microsoft.com&apv_32260_www07=5&cpv_32260_www07=5; ai_session=/4+E8sUIzCDuYnK5wKPWYO|1637607702971|1637608847712.4' \
--data-raw '{"text":"","facet_filters":[{"facet":"story_product_categories","values":["Cloud Platform","Azure"]},{"facet":"story_country_region","values":["North America"]},{"facet":"story_country","values":["North America/United States"]}],"related_documents":[],"page_id":"0","featured_sections":null,"sort_mode":"story_publish_date desc"}' \
--compressed`
