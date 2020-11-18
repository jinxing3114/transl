package messages

func init(){
Text["pl"] = []string{
		"(brak wartości)",
		"Wystąpił wewnętrzny błąd serwera.",
		"Usuń",
		"Błąd",
		"Wgrywanie pliku nie powiodło się.",
		"Strona domowa",
		"Otrzymano nieprawidłowe dane dla parametru \"{param}\".",
		"Wymagane zalogowanie się",
		"Brak wymaganych argumentów: {params}",
		"Brak wymaganych parametrów: {params}",
		"Nie",
		"",
		"",
		"Brak wyników.",
		"Dozwolone są tylko pliki z następującymi rozszerzeniami: {extensions}.",
		"Dozwolone są tylko pliki z następującymi typami MIME: {mimeTypes}.",
		"Nie odnaleziono strony.",
		"Proszę poprawić następujące błędy:",
		"Proszę wgrać plik.",
		"Wyświetlone <b>{begin, number}-{end, number}</b> z <b>{totalCount, number}</b> {totalCount, plural, one{rekordu} other{rekordów}}.",
		"Plik \"{file}\" nie jest obrazem.",
		"Plik \"{file}\" jest zbyt duży. Jego rozmiar nie może przekraczać {formattedLimit}.",
		"Plik \"{file}\" jest za mały. Jego rozmiar nie może być mniejszy niż {formattedLimit}.",
		"Format {attribute} jest nieprawidłowy.",
		"Obraz \"{file}\" jest zbyt duży. Wysokość nie może być większa niż {limit, number} {limit, plural, one{piksela} few{pikseli} many{pikseli} other{piksela}}.",
		"Obraz \"{file}\" jest zbyt duży. Szerokość nie może być większa niż {limit, number} {limit, plural, one{piksela} few{pikseli} many{pikseli} other{piksela}}.",
		"Obraz \"{file}\" jest za mały. Wysokość nie może być mniejsza niż {limit, number} {limit, plural, one{piksela} few{pikseli} many{pikseli} other{piksela}}.",
		"Obraz \"{file}\" jest za mały. Szerokość nie może być mniejsza niż {limit, number} {limit, plural, one{piksela} few{pikseli} many{pikseli} other{piksela}}.",
		"Kod weryfikacyjny jest nieprawidłowy.",
		"Razem <b>{count, number}</b> {count, plural, one{rekord} few{rekordy} many{rekordów} other{rekordu}}.",
		"Nie udało się zweryfikować przesłanych danych.",
		"",
		"Nieznana opcja: --{name}",
		"Aktualizuj",
		"Zobacz szczegóły",
		"Tak",
		"Brak upoważnienia do wykonania tej czynności.",
		"Możliwe wgranie najwyżej {limit, number} {limit, plural, one{pliku} few{plików} many{plików} other{pliku}}.",
		"wartość wejściowa",
		"{attribute} \"{value}\" jest już w użyciu.",
		"{attribute} nie może pozostać bez wartości.",
		"{attribute} zawiera nieprawidłową wartość.",
		"{attribute} nie zawiera prawidłowego adresu URL.",
		"{attribute} nie zawiera prawidłowego adresu email.",
		"{attribute} musi mieć wartość \"{requiredValue}\".",
		"{attribute} musi być liczbą.",
		"{attribute} musi być tekstem.",
		"{attribute} musi być liczbą całkowitą.",
		"{attribute} musi mieć wartość \"{true}\" lub \"{false}\".",
		"{attribute} musi mieć tę samą wartość co \"{compareValueOrAttribute}\".",
		"",
		"",
		"",
		"",
		"{attribute} musi wynosić nie więcej niż {max}.",
		"{attribute} musi wynosić nie mniej niż {min}.",
		"",
		"",
		"{attribute} musi zawierać co najmniej {min, number} {min, plural, one{znak} few{znaki} many{znaków} other{znaku}}.",
		"{attribute} musi zawierać nie więcej niż {max, number} {max, plural, one{znak} few{znaki} many{znaków} other{znaku}}.",
		"{attribute} musi zawierać dokładnie {length, number} {length, plural, one{znak} few{znaki} many{znaków} other{znaku}}.",
		"{nFormatted} B",
		"{nFormatted} GB",
		"{nFormatted} GiB",
		"{nFormatted} KB",
		"{nFormatted} KiB",
		"{nFormatted} MB",
		"{nFormatted} MiB",
		"{nFormatted} PB",
		"{nFormatted} PiB",
		"{nFormatted} TB",
		"{nFormatted} TiB",
		"{nFormatted} {n, plural, =1{bajt} few{bajty} many{bajtów} other{bajta}}",
		"{nFormatted} {n, plural, =1{gibibajt} few{gigabajty} many{gibiajtów} other{gibiajta}}",
		"{nFormatted} {n, plural, =1{gigabajt} few{gigabajty} many{gigabajtów} other{gigabaja}}",
		"{nFormatted} {n, plural, =1{kibibajt} few{kibibajty} many{kibibajtów} other{kibibajtów}}",
		"{nFormatted} {n, plural, =1{kilobajt} few{kilobajty} many{kilobajtów} other{kilobajtów}}",
		"{nFormatted} {n, plural, =1{mebibajt} few{mebibajty} many{mebibajtów} other{mebibajta}}",
		"{nFormatted} {n, plural, =1{megabajt} few{megabajty} many{megabajtów} other{megabajta}}",
		"{nFormatted} {n, plural, =1{pebibajt} few{pebibajty} many{pebibajtów} other{pebibajta}}",
		"{nFormatted} {n, plural, =1{petabajt} few{petabajty} many{petabajtów} other{petabajta}}",
		"{nFormatted} {n, plural, =1{tebibajt} few{tebibajty} many{tebibajtów} other{tebibajta}}",
		"{nFormatted} {n, plural, =1{terabajt} few{terabajty} many{terabajtów} other{terabajta}}",
		"Czy na pewno usunąć ten element?",
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
		"Żądany widok \"{name}\" nie został odnaleziony.",
		"za {delta, plural, =1{jeden dzień} other{# dni}}",
		"za {delta, plural, =1{minutę} few{# minuty} many{# minut} other{# minuty}}",
		"za {delta, plural, =1{miesiąc} few{# miesiące} many{# miesięcy} other{# miesiąca}}",
		"za {delta, plural, =1{sekundę} few{# sekundy} many{# sekund} other{# sekundy}}",
		"za {delta, plural, =1{rok} few{# lata} many{# lat} other{# roku}}",
		"za {delta, plural, =1{godzinę} few{# godziny} many{# godzin} other{# godziny}}",
		"{delta, plural, =1{jeden dzień} other{# dni} other{# dnia}} temu",
		"{delta, plural, =1{minutę} few{# minuty} many{# minut} other{# minuty}} temu",
		"{delta, plural, =1{miesiąc} few{# miesiące} many{# miesięcy} other{# miesiąca}} temu",
		"{delta, plural, =1{sekundę} few{# sekundy} many{# sekund} other{# sekundy}} temu",
		"{delta, plural, =1{rok} few{# lata} many{# lat} other{# roku}} temu",
		"{delta, plural, =1{godzinę} few{# godziny} many{# godzin} other{# godziny}} temu",
		"przed chwilą",
		"",
		"",
		"Zestawienie {values} dla {attributes} jest już w użyciu.",
		"Nieznany alias: -{name}",
		"",
		"{attribute} posiada złą maskę podsieci.",
		"{attribute} nie jest w dozwolonym zakresie.",
		"{attribute} musi być poprawnym adresem IP.",
		"{attribute} musi być adresem IP w określonej podsieci.",
		"{attribute} musi mieć wartość większą od \"{compareValueOrAttribute}\".",
		"{attribute} musi mieć wartość większą lub równą \"{compareValueOrAttribute}\".",
		"{attribute} musi mieć wartość mniejszą od \"{compareValueOrAttribute}\".",
		"{attribute} musi mieć wartość mniejszą lub równą \"{compareValueOrAttribute}\".",
		"{attribute} nie może być podsiecią.",
		"{attribute} nie może być adresem IPv4.",
		"{attribute} nie może być adresem IPv6.",
		"{attribute} musi mieć wartość różną od \"{compareValueOrAttribute}\".",
		"{delta, plural, =1{1 dzień} other{# dni} other{# dnia}}",
		"{delta, plural, =1{1 godzina} few{# godziny} many{# godzin} other{# godziny}}",
		"{delta, plural, =1{1 minuta} few{# minuty} many{# minut} other{# minuty}}",
		"{delta, plural, =1{1 miesiąc} few{# miesiące} many{# miesięcy} other{# miesiąca}}",
		"{delta, plural, =1{1 sekunda} few{# sekundy} many{# sekund} other{# sekundy}}",
		"{delta, plural, =1{1 rok} few{# lata} many{# lat} other{# roku}}",
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