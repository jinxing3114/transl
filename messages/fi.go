package messages

func init(){
Text["fi"] = []string{
		"(ei asetettu)",
		"Sisäinen palvelinvirhe.",
		"Poista",
		"Virhe",
		"Tiedoston lähetys epäonnistui.",
		"Koti",
		"Parametri \"{param}\" vastaanotti virheellistä dataa.",
		"Kirjautuminen vaaditaan",
		"Pakolliset argumentit puuttuu: {params}",
		"Pakolliset parametrit puuttuu: {params}",
		"Ei",
		"",
		"",
		"Ei tuloksia.",
		"Sallittuja ovat vain tiedostot, joiden tiedostopääte on: {extensions}.",
		"Sallittuja ovat vain tiedostot, joiden MIME-tyyppi on: {mimeTypes}.",
		"Sivua ei löytynyt.",
		"Korjaa seuraavat virheet:",
		"Lähetä tiedosto.",
		"Näytetään <b>{begin, number}-{end, number}</b> kaikkiaan <b>{totalCount, number}</b> {totalCount, plural, one{tuloksesta} other{tuloksesta}}.",
		"Tiedosto \"{file}\" ei ole kuva.",
		"Tiedosto \"{file}\" on liian iso. Sen koko ei voi olla suurempi kuin {formattedLimit}.",
		"Tiedosto \"{file}\" on liian pieni. Sen koko ei voi olla pienempi kuin {formattedLimit}.",
		"Attribuutin {attribute} formaatti on virheellinen.",
		"Kuva \"{file}\" on liian suuri. Korkeus ei voi olla suurempi kuin {limit, number} {limit, plural, one{pikseli} other{pikseliä}}.",
		"Kuva \"{file}\" on liian suuri. Leveys ei voi olla suurempi kuin {limit, number} {limit, plural, one{pikseli} other{pikseliä}}.",
		"Kuva \"{file}\" on liian pieni. Korkeus ei voi olla pienempi kuin {limit, number} {limit, plural, one{pikseli} other{pikseliä}}.",
		"Kuva \"{file}\" on liian pieni. Leveys ei voi olla pienempi kuin {limit, number} {limit, plural, one{pikseli} other{pikseliä}}.",
		"Vahvistuskoodi on virheellinen.",
		"Yhteensä <b>{count, number}</b> {count, plural, one{tulos} other{tulosta}}.",
		"Tietojen lähetystä ei voida varmistaa.",
		"",
		"Tuntematon valinta: --{name}",
		"Päivitä",
		"Näytä",
		"Kyllä",
		"Sinulla ei ole tarvittavia oikeuksia toiminnon suorittamiseen.",
		"Voit lähettää enintään {limit, number} {limit, plural, one{tiedoston} other{tiedostoa}}.",
		"syötetty arvo",
		"{attribute} \"{value}\" on jo käytössä.",
		"{attribute} ei voi olla tyhjä.",
		"{attribute} on virheellinen.",
		"{attribute} on virheellinen URL.",
		"{attribute} on virheellinen sähköpostiosoite.",
		"{attribute} täytyy olla \"{requiredValue}\".",
		"{attribute} täytyy olla luku.",
		"{attribute} täytyy olla merkkijono.",
		"{attribute} täytyy olla kokonaisluku.",
		"{attribute} täytyy olla joko {true} tai {false}.",
		"{attribute} täytyy olla yhtä suuri kuin \"{compareValueOrAttribute}\".",
		"",
		"",
		"",
		"",
		"{attribute} ei saa olla suurempi kuin \"{max}\".",
		"{attribute} ei saa olla pienempi kuin \"{min}\".",
		"",
		"",
		"{attribute} tulisi sisältää vähintään {min, number} {min, plural, one{merkki} other{merkkiä}}.",
		"{attribute} tulisi sisältää enintään {max, number} {max, plural, one{merkki} other{merkkiä}}.",
		"{attribute} tulisi sisältää {length, number} {length, plural, one{merkki} other{merkkiä}}.",
		"{nFormatted} t",
		"{nFormatted} Gt",
		"{nFormatted} GiB",
		"{nFormatted} kt",
		"{nFormatted} KiB",
		"{nFormatted} Mt",
		"{nFormatted} MiB",
		"{nFormatted} Pt",
		"{nFormatted} PiB",
		"{nFormatted} Tt",
		"{nFormatted} TiB",
		"{nFormatted} {n, plural, =1{tavu} other{tavua}}",
		"{nFormatted} {n, plural, =1{gibitavu} other{gibitavua}}",
		"{nFormatted} {n, plural, =1{gigatavu} other{gigatavua}}",
		"{nFormatted} {n, plural, =1{kibitavu} other{kibitavua}}",
		"{nFormatted} {n, plural, =1{kilotavu} other{kilotavua}}",
		"{nFormatted} {n, plural, =1{mebitavu} other{mebitavua}}",
		"{nFormatted} {n, plural, =1{megatavu} other{megatavua}}",
		"{nFormatted} {n, plural, =1{pebitavu} other{pebitavua}}",
		"{nFormatted} {n, plural, =1{petatavu} other{petatavua}}",
		"{nFormatted} {n, plural, =1{tebitavu} other{tebitavua}}",
		"{nFormatted} {n, plural, =1{teratavu} other{teratavua}}",
		"Haluatko varmasti poistaa tämän?",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"Pyydettyä näkymää \"{name}\" ei löytynyt.",
		"{delta, plural, =1{päivässä} other{# päivässä}}",
		"{delta, plural, =1{minuutissa} other{# minuutissa}}",
		"{delta, plural, =1{kuukaudessa} other{# kuukaudessa}}",
		"{delta, plural, =1{sekunnissa} other{# sekunnissa}}",
		"{delta, plural, =1{vuodessa} other{# vuodessa}}",
		"{delta, plural, =1{tunnissa} other{# tunnissa}}",
		"{delta, plural, =1{päivä} other{# päivää}} sitten",
		"{delta, plural, =1{minuutti} other{# minuuttia}} sitten",
		"{delta, plural, =1{kuukausi} other{# kuukautta}} sitten",
		"{delta, plural, =1{sekunti} other{# sekuntia}} sitten",
		"{delta, plural, =1{vuosi} other{# vuotta}} sitten",
		"{delta, plural, =1{tunti} other{# tuntia}} sitten",
		"juuri nyt",
		"",
		"Powered by {yii}",
		"",
		"",
		"Yii Framework",
		"{attribute} sisältää väärän aliverkkopeitteen.",
		"{attribute} ei ole sallitulla alueella.",
		"{attribute} täytyy olla kelvollinen IP-osoite.",
		"{attribute} täytyy olla määritetyllä aliverkolla oleva IP-osoite.",
		"{attribute} täytyy olla suurempi kuin \"{compareValueOrAttribute}\".",
		"{attribute} täytyy olla suurempi tai yhtä suuri kuin \"{compareValueOrAttribute}\".",
		"{attribute} täytyy olla pienempi kuin \"{compareValueOrAttribute}\".",
		"{attribute} täytyy olla pienempi tai yhtä suuri kuin \"{compareValueOrAttribute}\".",
		"{attribute} ei saa olla aliverkko.",
		"{attribute} ei saa olla IPv4-osoite.",
		"{attribute} ei saa olla IPv6-osoite.",
		"{attribute} ei saa olla yhtä suuri kuin \"{compareValueOrAttribute}\".",
		"{delta, plural, =1{1 päivä} other{# päivää}}",
		"{delta, plural, =1{1 tunti} other{# tuntia}}",
		"{delta, plural, =1{1 minuutti} other{# minuuttia}}",
		"{delta, plural, =1{1 kuukausi} other{# kuukautta}}",
		"{delta, plural, =1{1 sekunti} other{# sekuntia}}",
		"{delta, plural, =1{1 vuosi} other{# vuotta}}",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}
}