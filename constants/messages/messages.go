package messages

const (
	HelloMessage = `<b>Приветствую вас Товарищи!</b> 🐉

Сегодня мы снова в строю, чтобы гонять ящеров из наших славных земель!
Кто здесь командир, кто разведчик, а кто просто знаток ящерской подлянки?

Участвуй в игре "Ящер дня", зови союзников и, главное, не доверяй тем, у кого чешуя блестит слишком сильно — тут явно ящер шпионит!`
	AlreadyStartedMessage      = "🛡️ Я уже в деле и готов к охоте на ящеров! Помните: ящер может быть ближе, чем кажется... 🐉"
	RegisteredMessage          = "@%s теперь ты на поле боя в поисках ящера дня. Но помни: любой может оказаться хвостатым предателем... даже ты!"
	AlreadyRegisteredMessage   = "Эй, @%s, ты уже участвуешь в игре! 🐉 Не пытайся отмазаться!"
	NoParticipantsMessage      = "Нет учатсников, так что все признаны ящерами! Участвуйте в игре что бы защитить свою честь!"
	TooFewParticipantsMessage  = "Все кроме @%s признаны ящерами!"
	TopWinnersMessage          = "<b>Топ 10 ящеров за все время:</b>\n\n"
	ParticipantsListMessage    = "<b>Список участников в игре :</b>\n\n"
	ParticipantsListMessageEnd = "\n\n<b>Тот кого нет в этом списке воняет ящером, гасите его!</b>"
	TopWinnersEmpty            = "Список топ-10 ящеров за всё время пуст, будто их и не бывало! Похоже, хвостатые отлично шифруются, или мы пока не поймали ни одного заслуживающего места в этом почётном рейтинге."
	BotIsNotStarted            = "Бот не запущен, запустите его товарищи!"
	SpecialWinnerMessage       = "А главный ящер всея чата @YohoCX! И не отмоешся ты от этого позора никогда!!!"
)

var SearchInProgressMessages = map[int]string{
	0:  "Товарищи, ящер пока прячется, но мы его найдём!",
	1:  "Он где-то тут, я это чувствую. Шерстим всё вокруг!",
	2:  "Пока что тихо... Но не расслабляйтесь, ящер хитёр!",
	3:  "Где этот шипящий змей? Мы должны его найти!",
	4:  "Ящер затаился, но мы на его следу. Продолжаем поиски!",
	5:  "Проверяйте кусты, подвалы и соседи — ящер любит прятаться!",
	6:  "Кажется, он забился под диван. Или это опять кошка?",
	7:  "Наши агенты уже в пути. Ящер, выходи сам, или мы тебя выкурим!",
	8:  "Тишина... Но не расслабляйтесь, он точно рядом!",
	9:  "Ящер хитрый, но у нас глаза на затылке. Скоро найдём!",
	10: "Небо чистое, но запах чешуи в воздухе. Где-то он есть!",
	11: "Ящер ещё не пойман, но разведка сообщает, что он близко!",
	12: "Ящер, мы знаем, что ты нас слышишь. Вылезай сам, пока не поздно!",
	13: "Где этот хвостатый? Признавайтесь, кто-нибудь видел ящера?",
	14: "Поисковая группа продолжает осмотр. Ящер, тебе не спрятаться!",
}

var WinnerMessages = map[int]string{
	0:  "Тревога, товарищи! @%s — ящер! Готовьте лопаты и солёные огурцы! Но если ящерка держите руки на столе!",
	1:  "Ящер засветился! Это @%s! Вперёд, товарищи, в бой!",
	2:  "Он тут, он здесь, и это @%s! Воняет чешуёй, ловим его!",
	3:  "Всё ясно: @%s оказался ящером! Покажем ему, кто тут главный!",
	4:  "Вы слышите это шипение? Ага, это @%s — наш ящер дня!",
	5:  "Внимание, внимание! Ящер найден, и это @%s! Режим охоты активирован!",
	6:  "Наши разведчики спалили @%s! Этот ящер теперь никуда не денется!",
	7:  "@%s, ты попался! Давай на выход, пока по-хорошему!",
	8:  "Горячий след привёл нас к @%s! Ящер обнаружен!",
	9:  "Кто-то слишком хорошо шифровался, но мы вычислили @%s. Это ящер!",
	10: "@%s, ты наш ящер! Готовься к атаке!",
	11: "Шутки в сторону, @%s! Ты оказался ящером, и теперь тебе не спрятаться!",
	12: "Граждане, внимание! @%s оказался тем самым ящером, которого мы искали!",
	13: "А вот и он, скользкий враг! @%s, твоя чешуя блестит слишком ярко!",
	14: "@%s — это ящер! Товарищи, время объединиться и показать, кто здесь хозяин!",
}

var WinnerAlreadyFoundMessages = map[int]string{
	0: "Отбой, бойцы! Ящер уже найден, и это <b>%s</b>. На сегодня все, идём чистить лопаты до завтра!",
	1: "Спокойнее, товарищ! Мы уже споймали ящера — это <b>%s</b>. Хватит геройствовать на сегодня!",
	2: "Чего ты шумишь? Ящер дня — <b>%s</b>, а я уже всё... устал.",
	3: "Товарищ, не гони лошадей! Ящер найден, и это <b>%s</b>. Пора передохнуть до следующего дня!",
	4: "Бой окончен: <b>%s</b> — ящер дня. Продолжим наше веселье завтра!",
}

var HealthCheckMessages = map[int]string{
	0: "Всё путём, товарищ!",
	1: "На связи, товарищи! Я жив!",
	2: "Я в строю, всё под контролем. Ящеры могут прятаться, но я всё замечу!",
	3: "Отлично себя чувствую, уже стою на часах, жду появления хвостатых.",
	4: "Докладываю: бот здоров, силён и в полной боевой готовности.",
	5: "Я здесь! Живее всех живых, а ящеров пока не видно.",
	6: "Здоров как бык, бодр как огурец, готов снова гонять скользких врагов.",
	7: "В порядке, товарищ! Продолжаю наблюдение за чешуйчатыми.",
	8: "К бою готов! Всё функционирует на ура, давайте искать ящеров.",
	9: "Шшш… Всё спокойно, я на посту. Если ящер появится, я дам знать!",
}
