//	Phone tokenizer and normalization package
//	https://github.com/nkozyra/entities
//	Honors E.164: http://en.wikipedia.org/wiki/E.164
//	Respects national conventions: http://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers

/*
	Usage:
	phone := phone.New("+1 (202) 555-1234")
	phone.Normalize()

*/

package phone

import
(
		"fmt"
		"regexp"
		//"strconv"
)

var
(
	pats []PhoneComponent
)

type Phone struct {
	Raw string
	Normalized string
	CountryCode string
	AreaCode string
	SubscriberNumber string
	InternationalPrefix int
	LongDistancePrefix int
	Vanity bool
}

type PhoneComponent struct {
	pattern string
	countryCodeExists bool
	countryCodePosition string

	areaCodeExists bool
	areaCodePosition string

	subscriberNumberExists bool
	subscriberNumberPosition string
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
				countryCodePosition: `$1`,
				areaCodeExists: true,
				areaCodePosition: `$2`,
			}, 
			//
			{ pattern: `\+(\d{1,3})\s+(\d{3})\s+(\d{3})[\s\-]+(\d{4})`, 
				countryCodeExists: true,
				countryCodePosition: `$1`,
				areaCodeExists: true,
				areaCodePosition: `$2`,
				subscriberNumberExists: true,
				subscriberNumberPosition: `$3 $4`,
			}, 
			// 1(817) 569-8900	
			{ pattern: `(\d{1,3})\s*(\d{3})\s+(\d{3})[\s\-]+(\d{4})`, 
				countryCodeExists: true,
				countryCodePosition: `$1`,
				areaCodeExists: true,
				areaCodePosition: `$2`,
				subscriberNumberExists: true,
				subscriberNumberPosition: `$3 $4`,
			}, 
			// 		
	}

}

func (a *Phone) Prepare() {

	//	The only characters we'll allow are +,[0-9],-,(,),[a-z]
	//	Replace the rest with whitespace

	rg,_ := regexp.Compile(`[^\d\-\+\(\)\s]+`)
	a.Raw = rg.ReplaceAllString(a.Raw," ")
	sp,_ := regexp.Compile(`\s{2,}`)
	a.Raw = sp.ReplaceAllString(a.Raw," ")
}

func (a *Phone) Normalize() {
	Init()
	a.Prepare()
	for i,_ := range pats {

		rg,err := regexp.Compile(pats[i].pattern)
		if err != nil {
			fmt.Println(err.Error())
			return
		} else {

			addbyte := []byte(a.Raw)
			if rg.Match(addbyte) {


				// country code resolution
				if pats[i].countryCodeExists == true {
					rpos := "+" + pats[i].countryCodePosition
					a.CountryCode = rg.ReplaceAllString(a.Raw, rpos)
				} else {
					a.CountryCode = "0"
				}

				// area code resolution
				if pats[i].areaCodeExists == true {
					rpos := pats[i].areaCodePosition
					a.AreaCode = rg.ReplaceAllString(a.Raw, rpos)
				} else {
					a.AreaCode = ""
				}

				break

			} else {

			}

		}
	}
}