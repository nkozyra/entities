//	Phone tokenizer and normalization package
//	https://github.com/nkozyra/entities

package phone

import
(
		"fmt"
		"regexp"
		"strconv"
)

var
(
	pats []PhoneComponent
)

type Phone struct {
	Raw string
	Normalized string
	CountryCode string
	AreaCode int
	InternationalPrefix int
	LongDistancePrefix int
	Vanity bool
}

type PhoneComponent struct {
	pattern string
	countryCodeExists bool
	countryCodePosition int
}

func New(raw string) (Phone) {
	na := Phone{Raw: raw}
	return na
}

func Init() {
 
	pats = []PhoneComponent{
			//+1 (123) 456-7890, +1 (123) 456 7890, 
			{ pattern: `\+(\d{1,3})\s+\((\d{3})\)\s+(\d{3})[\s\-]+(\d{4})`, 
				countryCodeExists: true,
				countryCodePosition: 1,
			}, 
			{ pattern: "what" },
	}

}

func (a *Phone) Prepare() {
	//	The only characters we'll allow are +,[0-9],-,(,),[a-z]
	//	Replace the rest with whitespace
}

func (a *Phone) Normalize() {
	Init()

	for i,_ := range pats {

		rg,err := regexp.Compile(pats[i].pattern)
		if err != nil {
			fmt.Println(err.Error())
			return
		} else {

			addbyte := []byte(a.Raw)
			if rg.Match(addbyte) {
				fmt.Println("YES",pats[i])

				// country code resolution
				if pats[i].countryCodeExists == true {
					rpos := "+$" + strconv.FormatInt(int64(pats[i].countryCodePosition), 10)
					a.CountryCode = rg.ReplaceAllString(a.Raw, rpos)
				} else {
					a.CountryCode = ""
				}

				
			} else {
				fmt.Println("No match",pats[i])
			}

		}
	}
}