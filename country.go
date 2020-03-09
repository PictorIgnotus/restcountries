package restcountries

type Country struct {
	Name           	string
	Capital        	string
	AltSpellings   	[]string
	Relevance      	string
	Region         	string
	Subregion      	string
	Translations   	map[string]string
	Population     	int32
	LatLng         	[]float32
	Demonym        	string
	Area           	float32
	Gini           	float32
	Timezones      	[]string
	Borders        	[]string
	NativeName     	string
	CallingCodes   	[]string
	TopLevelDomain 	[]string
	Alpha2Code     	string
	Alpha3Code     	string
	NumericCode	   	string
	Currencies		[]Currency
	Languages		[]Language
	Flag			string
	RegionalBlocks	[]RegionalBlock
	Cioc			string

}

type Currency struct {
	Code	string
	Name	string
	Symbol	string
}

type Language struct {
	ISO639_1	string
	ISO639_2	string
	Name		string
	NativeName	string
}

type RegionalBlock struct {
	Acronym 		string
	Name 			string
	OtherAcronyms 	[]string
	OtherNames 		[]string
}
