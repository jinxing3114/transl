package messages

func init(){
Text["sr"] = []string{
		"(без вредности)",
		"Дошло је до интерне грешке на серверу.",
		"Обриши",
		"Грешка",
		"Постављање фајла није успело.",
		"Почетна",
		"Неисправан податак добијен за параметар \"{param}.\"",
		"Пријава на систем је обавезна",
		"Недостају обавезни аргументи: {params}",
		"Недостају обавезни параметри: {params}",
		"Не",
		"Не постоји помоћ за непознату команду \"{command}\".",
		"Не постоји помоћ за непознату под-команду \"{command}\".",
		"Нема резултата.",
		"Само фајлови са следећим екстензијама су дозвољени: {extensions}.",
		"Само следећи MIME типови су дозвољени: {mimeTypes}.",
		"Страница није пронађена.",
		"Молимо вас исправите следеће грешке:",
		"Молимо вас поставите фајл.",
		"Приказано <b>{begin, number}-{end, number}</b> од <b>{totalCount, number}</b> {totalCount, plural, =1{# ставке} one{# ставкe} few{# ставке} many{# ставки} other{# ставки}}.",
		"Фајл \"{file}\" није слика.",
		"Фајл \"{file}\" је превелик. Величина не може бити већа од {formattedLimit}.",
		"Фајл \"{file}\" је премали. Величина не може бити мања од {formattedLimit}.",
		"Формат атрибута \"{attribute}\" је неисправан.",
		"Слика \"{file}\" је превелика. Висина не сме бити већа од {limit, number} {limit, plural, one{пиксел} other{пиксела}}.",
		"Слика \"{file}\" је превелика. Ширина не сме бити већа од {limit, number} {limit, plural, one{пиксел} other{пиксела}}.",
		"Слика \"{file}\" је премала. Висина не сме бити мања од {limit, number} {limit, plural, one{пиксел} other{пиксела}}.",
		"Слика \"{file}\" је премала. Ширина не сме бити мања од {limit, number} {limit, plural, one{пиксел} other{пиксела}}.",
		"Код за потврду није исправан.",
		"Укупно <b>{count, number}</b> {count, plural, one{ставка} few{ставке} other{ставки}}.",
		"Није могуће верификовати ваше послате податке.",
		"Непозната команда \"{command}\".",
		"Непозната опција: --{name}",
		"Исправи",
		"Приказ",
		"Да",
		"Немате права да извршите ову акцију.",
		"Можете поставити највише {limit, number} {limit, plural, one{фајл} few{фајлa} other{фајлова}}.",
		"улазна вредност",
		"{attribute} \"{value}\" је већ заузет.",
		"{attribute} не сме бити празан.",
		"{attribute} је неисправан.",
		"{attribute} није исправан URL.",
		"{attribute} није исправна email адреса.",
		"{attribute} мора бити \"{requiredValue}\".",
		"{attribute} мора бити број.",
		"{attribute} мора бити текст.",
		"{attribute} мора бити цели број.",
		"{attribute} мора бити \"{true}\" или \"{false}\".",
		"",
		"{attribute} мора бити већи од \"{compareValue}\".",
		"{attribute} мора бити већи или једнак \"{compareValue}\".",
		"{attribute} мора бити мањи од \"{compareValue}\".",
		"{attribute} мора бити мањи или једнак \"{compareValue}\".",
		"{attribute} не сме бити већи од \"{max}\".",
		"{attribute} не сме бити мањи од {min}.",
		"{attribute} мора бити исправно поновљен.",
		"{attribute} не сме бити једнак \"{compareValue}\".",
		"{attribute} треба да садржи барем {min, number} {min, plural, one{карактер} other{карактера}}.",
		"{attribute} треба да садржи највише {max, number} {max, plural, one{карактер} other{карактера}}.",
		"{attribute} треба да садржи {length, number} {length, plural, one{карактер} other{карактера}}.",
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
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"Да ли сте сигурни да желите да обришете ову ставку?",
		"{n, plural, one{# бајт} few{# бајта} many{# бајтова} other{# бајта}}",
		"{n, plural, one{# гигабајт} few{# гигабајта} many{# гигабајта} other{# гигабајта}}",
		"{n, plural, one{# килобајт} few{# килобајта} many{# килобајта} other{# килобајта}}",
		"{n, plural, one{# мегабајт} few{# мегабајта} many{# мегабајта} other{# мегабајта}}",
		"{n, plural, one{# петабајт} few{# петабајта} many{# петабајта} other{# петабајта}}",
		"{n, plural, one{# терабајт} few{# терабајта} many{# терабајта} other{# терабајта}}",
		"{n} Б",
		"{n} ГБ",
		"{n} КБ",
		"{n} МБ",
		"{n} ПБ",
		"{n} ТБ",
		"Тражени приказ \"{name}\" није пронађен.",
		"за {delta, plural, =1{дан} one{# дан} few{# дана} many{# дана} other{# дана}}",
		"за {delta, plural, =1{минут} one{# минут} few{# минута} many{# минута} other{# минута}}",
		"за {delta, plural, =1{месец} one{# месец} few{# месеца} many{# месеци} other{# месеци}}",
		"за {delta, plural, =1{секунду} one{# секунду} few{# секундe} many{# секунди} other{# секунди}}",
		"за {delta, plural, =1{годину} one{# годину} few{# године} many{# година} other{# година}}",
		"за {delta, plural, =1{сат} one{# сат} few{# сата} many{# сати} other{# сати}}",
		"пре {delta, plural, =1{дан} one{дан} few{# дана} many{# дана} other{# дана}}",
		"пре {delta, plural, =1{минут} one{# минут} few{# минута} many{# минута} other{# минута}}",
		"пре {delta, plural, =1{месец} one{# месец} few{# месеца} many{# месеци} other{# месеци}}",
		"пре {delta, plural, =1{секунде} one{# секунде} few{# секунде} many{# секунди} other{# секунди}}",
		"пре {delta, plural, =1{године} one{# године} few{# године} many{# година} other{# година}}",
		"пре {delta, plural, =1{сат} one{# сат} few{# сата} many{# сати} other{# сати}}",
		"управо сада",
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