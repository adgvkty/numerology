package version1

//DestinyText глобальный типа данных для текстов
type DestinyText struct {
	destinyChar    string
	destinyPros    string
	destinyCons    string
	destinyPurpose string
}

//LoadResources загружает ресурсы (тексты)
//в картуmap[int]DestinyText
func LoadResources() map[int]DestinyText {

	//создание карты с 9 значениями для возврата текстов
	DestinyTexts := make(map[int]DestinyText, 9)

	//присвоения текстов композитными литералами
	//обращение к карте идёт по индексу,
	//где i+1 равен числу судьбы в тексте
	DestinyTexts[1] = DestinyText{
		destinyChar:    "\nОбщая характеристика. Люди, чье число судьбы по дате\n рождения равно единице, — прирожденные лидеры. Стремление быть в первых\n рядах, выделяться из серой массы — врожденная их особенность. Это яркие\n индивидуалисты, для которых собственное “Я” всегда стоит на первом месте.\n Это невероятно активные, амбициозные, оригинальные и независимые\n личности, способные ради достижения своих целей идти по головам.\n",
		destinyPros:    "\nДостоинства. Уверенность в себе, небывалая\n целеустремленность, решительность и отвага, энергичность и позитивное\n видение жизни, честность и благородство — качества, которыми наделило\n число 1 своих подопечных. Люди-единицы слывут неунывающими оптимистами,\n отличаются хорошим чувством юмора. Они из тех, кто предпочитает работать\n своим собственным умом. От природы они талантливы, и этот свой талант\n умеют направлять в созидательное русло.\n",
		destinyCons:    "\nНедостатки. Личности, управляемые числом судьбы 1,\n отличаются завидным упрямством, доходящим до упертости, чрезмерной\n прямолинейностью, высокомерностью и надменностью. Они эгоистичны,\n нетерпеливы, циничны, часто излишне агрессивны. Единицы любят власть и\n стремятся подогнуть под себя окружающих, характеризуются диктаторскими\n замашками. Тяжело воспринимают критику в свой адрес, не умеют признавать\n свою неправоту. Не лишены тщеславия, склонны к лени.\n",
		destinyPurpose: "\nПредназначение: быть руководителем и вдохновителем,\n заряжать окружающих своим энтузиазмом, призывать их к деятельности.\n",
	}
	DestinyTexts[2] = DestinyText{
		destinyChar:    "\nОбщая характеристика. Сотрудничество и партнерство — два\n слова, которые как нельзя точно характеризуют людей, числом судьбы\n которых является 2. Двойки — это прирожденные дипломаты и миротворцы.\n Стремятся к гармонии, без труда идут на компромисс.\n",
		destinyPros:    "\nДостоинства. Подопечные двойки — личности\n уравновешенные, внимательные, мягкие и терпеливые. Они добры, скромны,\n тактичны, не конфликтны, умеют делиться своим теплом. Хладнокровные и\n спокойные, благоразумные и осмотрительные, наделены аналитическим складом\n ума. Люди-двойки хорошо умеют разруливать чужие конфликты,\n создают вокруг себя обстановку покоя и гармонии. Они наделены врожденными\n способностями психолога и проницательностью.\n",
		destinyCons:    "\nНедостатки. Люди-двойки — неисправимые мечтатели,\n которые часто летают в облаках. Они чрезмерно застенчивы и нерешительны,\n абсолютно бесхитростны и непрактичны. Легко идут на поводу у окружающих\n людей. Склонны к быстрой смене настроения. Двойкам не хватает инициативы и\n напористости, малейшие неудачи заставляют их впадать в пессимизм и\n уныние. Всё это в реальной жизни часто оборачивается для людей с числом\n судьбы 2 комплексом неполноценности.\n",
		destinyPurpose: "\nПредназначение: принимать жизнь в естественном её виде,\n адаптироваться к ней, избегать крайностей.\n",
	}
	DestinyTexts[3] = DestinyText{
		destinyChar:    "\nОбщая характеристика. Талант и самовыражение — слова,\n характеризующие наилучшим образом людей, чье число судьбы представлено\n тройкой. Природная талантливость троек часто способствует тому, что эти\n люди свою жизнь связывают со сферой творчества. Эти личности наделены\n острым умом и любознательностью, изобретательностью и развитой интуицией,\n динамичностью и оптимизмом. У них огромная внутренняя энергия, которая\n просто не позволяет им усидеть на одном месте.\n",
		destinyPros:    "\nДостоинства. У троек хороший характер, который\n притягивает окружающих, подобно магниту. Они буквально излучают обаяние и\n очарование, любят быть в центре внимания, красноречивы и\n привлекательны. Очень дружелюбны и коммуникативны. Тройки ценят\n искренность и честность, дружба для них — не пустой звук. Люди с числом\n судьбы 3 отличаются необычайной щедростью, зачастую даже во вред себе.\n",
		destinyCons:    "\nНедостатки. Люди-тройки характеризуются нетерпеливостью\n и избыточной эмоциональностью, вспыльчивостью и\n гордостью. Склонны к зависти и хвастовству. Довольно болтливы, не прочь\n поучаствовать в распространении сплетен. Зачастую впустую расходуют свой\n талант, силы и время. Любят разбрасываться деньгами.\n",
		destinyPurpose: "\nПредназначение: дарить и принимать любовь, служить\n вдохновением окружающим людям.\n",
	}
	DestinyTexts[4] = DestinyText{
		destinyChar:    "\nОбщая характеристика. Человек, число судьбы по дате\n рождения которого равно 4, — это человек-крепость: серьезный,\n основательный, надежный и ответственный. Четверки отличаются\n уравновешенностью и трудолюбием, самостоятельностью и осторожностью, любят\n порядок и стабильность.\n",
		destinyPros:    "\nДостоинства. Люди-четверки характеризуются отвагой и\nрешительностью, при этом не любят рисковать. Порядочность, сдержанность\n и серьезность, высоконравственность, кропотливость и\n аккуратность, умение выполнять рутинную работу — вот что выделяет четверок\n на фоне окружающих. Эти люди обладают хорошими деловыми\n качествами. Мужчины с числом судьбы 4 — это личности, у которых руки\n растут из нужного места: их них получаются хорошие механики.\n",
		destinyCons:    "\nНедостатки. Подопечные числа судьбы 4 часто бывают\n нетерпеливыми и упрямыми, ленивыми и неповоротливыми. Склонны к унынию и\n пессимизму. Иногда им не достает уверенности в себе. В отношении своих\n родных и близких люди-четверки нередко проявляют властность и жесткость.\n",
		destinyPurpose: "\nПредназначение: “внедрять” в окружающих людей чувство\n постоянства и безопасности.\n",
	}
	DestinyTexts[5] = DestinyText{
		destinyChar:    "\nОбщая характеристика. Авантюризм и любовь к риску,\n импульсивность и подвижность, жизнерадостность и энергичность,\n непредсказуемость и неординарность, свободолюбие и независимость  — слова,\n которые хорошо характеризуют людей, чье число судьбы по дате рождения\n равно 5.",
		destinyPros:    "\nДостоинства. Пятерки — это люди, совершенно не способные\n усидеть на одном месте. Они очень активны, динамичны,\n постоянно движутся вперед и просто обожают всякого рода перемены. Тяга ко\n всему новому руководит их жизненным путем. Стать душой любой компании\n не составляет для них никакого труда.\n",
		destinyCons:    "\nНедостатки. Люди числа 5 не любят рутинную деятельность,\n поэтому часто берутся за несколько дел сразу и выполняют их довольно\n небрежно. Пятерки отличаются непостоянством и легкомысленностью. На этих\n личностей не всегда можно положиться, так как подопечные числа судьбы 5\n весьма безответственны.Эти люди могут вести себя эксцентрично, любят\n азарт, часто бывают ревнивы и завистливы.\n",
		destinyPurpose: "\nПредназначение: побуждать окружающих к действию.\n",
	}
	DestinyTexts[6] = DestinyText{
		destinyChar:    "\nОбщая характеристика. Надежность и стабильность, доброта\n и искренность, невероятный магнетизм — качества, характеризующие\n личность человека, числом судьбы которого является 6.\n Они стремятся к гармонии. Способны к состраданию.\n",
		destinyPros:    "\nДостоинства. Благородные и дружелюбные, спокойные и\n деликатные, добрые и гуманные, шестерки пользуются доверием окружающих. От\n этих личностей буквально веет теплом, с ними рядом надежно и весело.\n Они скромны и миролюбивы. В целом, люди очень милые и привлекательные.\n",
		destinyCons:    "\nНедостатки. Несмотря на весь набор положительных\n качества, люди числа 6 злопамятны и мстительны. Иногда они слишком\n придирчивы и ворчливы, часто необязательны и непостоянны. Прямолинейные,\n навязчивые и себялюбивые. Могут быть ленивыми и пассивными.\n",
		destinyPurpose: "\nПредназначение: забота о близких, помощь нуждающимся,\n формирование теплой, уютной, семейной обстановки.\n",
	}
	DestinyTexts[7] = DestinyText{
		destinyChar:    "\nОбщая характеристика. Число судьбы 7 — это число\n одиноких и молчаливых мыслителей, число-загадка. Его обладатели тяготеют к\n непрерывному духовному развитию, к получению новых знаний, к познанию\n мира и раскрытию его тайн. Люди числа 7 — интеллектуалы,\n наделенные оригинальностью и многими талантами. Это философы, для которых\n материальный аспект жизни занимает одно из последних мест.\n",
		destinyPros:    "\nДостоинства. У семерок пытливый ум аналитического\n склада, развитая интуиция. Они проницательны — от их внимания ничего не\n скрыть. Физическому труду подопечные числа 7 предпочитают умственный.\n Люди-семерки очень терпеливы и независимы, яркие индивидуалисты. Замкнуты,\n любят одиночество.\n",
		destinyCons:    "\nНедостатки. Так как семерки предпочитают одиночество,\n они могут быть угрюмыми и эмоционально холодными, отчужденными и\n неприступными. Частыми их спутниками являются пессимизм и уныние, что со\n временем может привести к депрессивным состояниям. Подопечные числа судьбы\n 7 склонны к неискренности, расчетливости и жестокости, к\n коварству и даже предательству. Могут стать алкоголиками. Зачастую их них\n получаются ярые фанатики.\n",
		destinyPurpose: "\nПредназначение: передавать свои знания, применять их во\n благо человечества; достигнуть внутреннего совершенства и являться\n примером для людей, стремящихся к интеллектуальной и духовной эволюции.\n",
	}
	DestinyTexts[8] = DestinyText{
		destinyChar:    "\nОбщая характеристика. 8 одаривает своего обладателя\n стремлением к достижениям, связанным с материальным благами, успехом,\n властью и деньгами. Люди-восьмерки предприимчивы, материальное\n предпочитают духовному. Обладают качествами лидера, умеют реализовать\n возникшие в голове идеи.\n",
		destinyPros:    "\nДостоинства. Восьмерки — люди сильные и волевые. Они\n изобретательны и оригинальны, независимы и честолюбивы, отважны и стойки,\n надежны и практичны. Обладают хорошими исполнительскими способностями.\n Характер у восьмерок энергичный и боевой. Это невозмутимые, уверенные в\n себе реалисты. Трудолюбивы, не страшатся изнурительного труда,\n отличаются терпением, упорством и целенаправленностью. Умеют отвечать за\n свои действия и слова.\n",
		destinyCons:    "\nНедостатки. Подопечные числа судьбы 8 характеризуются\n избыточной тягой к власти и богатству. Это помешанные на себе эгоисты,\n самовлюбленные и непредсказуемые циники, которые не любят подчиняться,\n часто играют другими людьми. Они очень упрямы и своенравны, склонны к\n тирании. Довольно безалаберны. Наделены способностью делать деньги, но\nсовершенно не умеют обращаться с заработанными финансами.\n",
		destinyPurpose: "\nПредназначение: научиться применять свою энергию и\n положительные качества на пользу окружающим, стараться не только для себя,\n а во благо других.\n",
	}
	DestinyTexts[9] = DestinyText{
		destinyChar:    "Общая характеристика. Сострадание и гуманизм — ключевые\n качества людей, находящихся в управлении числа судьбы 9\n. Девятки — это романтики и мечтатели, пришедшие в этот мир, чтобы отдавать\n другим людям свою любовь и сопереживание. Это интеллектуальные\n личности, наделенные сильной волей, живым умом и\n врожденной наблюдательностью. Подопечные числа 9 — люди свободолюбивые, не\n терпящие принуждения и ограничения своей независимости.\n",
		destinyPros:    "Достоинства. Девятки умеют создавать о себе хорошее\n впечатление, притягивают к себе людей, легко завоевывают чужое доверие и\n симпатию. Характеризуются тактичностью, деликатностью и вежливостью. В\n глазах окружающих выглядят общительными, дружелюбными и веселыми. У\n девяток стойкий, боевой и твердый характер, доминирующая натура, большой\n внутренний потенциал.\n",
		destinyCons:    "Недостатки. Люди-девятки могут быть излишне\n чувствительными и капризными, вспыльчивыми и агрессивными. Они\n нетерпеливы, себялюбивы, подвержены унынию и нерешительности. Уязвимы\n перед пагубными привычками (курение, алкоголизм, наркомания, токсикомани).\n",
		destinyPurpose: "Предназначение: озарять и облагораживать жизнь\n окружающих людей, помогать им находить правильный смысл жизни, служить\n примером настоящей любви и великодушия.\n",
	}

	return DestinyTexts
}
