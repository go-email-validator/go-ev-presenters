package utils

import "github.com/go-email-validator/go-email-validator/pkg/ev/evmail"

var EmailStrings = []string{
	"l-serderov@mail.ru",
	"ra.habusov@mail.ru",
	"borcov1337@mail.ru",
	"isaykanast@mail.ru",
	"zobrilovich@mail.ru",
	"porofert@mail.ru",
	"volodixina.agata@mail.ru",
	"poki9797@mail.ru",
	"ostap.promise.meshalkin@mail.ru",
	"roma.xayatov@mail.ru",
	"hger547fds@mail.ru",
	"maksim-sokolov-01@mail.ru",
	"Xsmiley2002@mail.ru",
	"yampolskaya.oksana.1975@mail.ru",
	"lost_team-96@mail.ru",
	"sokov_y@mail.ru",
	"larri.somov.91@mail.ru",
	"sasha.abdyllaevgame@mail.ru",
	"vakon00@mail.ru",
	"crosexspacod2008@mail.ru",
	"arnidipe1963@mail.ru",
	"stockchina@mail.ru",
	"alexeypotapov123@mail.ru",
	"sergey-0613@mail.ru",
	"andrey.taganov.2000@mail.ru",
	"anna_popova1969@mail.ru",
	"sergei.popov8080@mail.ru",
	"edik.somov.80@mail.ru",
	"olgapetr.82@mail.ru",
	"zaza.baturiya@mail.ru",
	"bekmek1985@mail.ru",
	"aidaralmenov@mail.ru",
	"vajvancevazhanna@mail.ru",
	"alya.perfurowa@mail.ru",
	"bukinnikita82@mail.ru",
	"narcissovplaton@mail.ru",
	"2ttt00@mail.ru",
	"vityanexlebaev8523@mail.ru",
	"annapavlov1983@mail.ru",
	"kirill_88_95@mail.ru",
	"vanechka.somov.1999@mail.ru",
	"sofya_petryashova@mail.ru",
	"annanos.nos@mail.ru",
	"anna.iova.88@mail.ru",
	"nana.boden@mail.ru",
	"lenkowkirill501@mail.ru",
	"voroshkov20032@mail.ru",
	"klya.moro.99@mail.ru",
	"kira.svetlova.1976@mail.ru",
	"gastbest@mail.ru",
	"y-sukova@mail.ru",
	"pitov.egor@mail.ru",
	"kralupysomov@mail.ru",
	"juriy.99@mail.ru",
	"pir.gomov@mail.ru",
	"yurij.loza.90@mail.ru",
	"juliasuhareva78@mail.ru",
	"dmitry_flor@mail.ru",
	"mileshin_06@mail.ru",
	"boden99@mail.ru",
	"postroy-novoe@mail.ru",
	"karimov.andreysaraiki@mail.ru",
	"galya.ilyasova.1985@mail.ru",
	"mrterrr13@mail.ru",
	"teyanaperlin@mail.ru",
	"kipilll93@mail.ru",
	"ribeachf@mail.ru",
	"qmanzb2011@mail.ru",
	"kakkupit@mail.ru",
	"goga.mr.egorka@mail.ru",
	"garretkor@mail.ru",
	"pofignafig123@mail.ru",
	"osmiusmiha75@mail.ru",
	"vladislav2003@MAIL.RU",
	"zver-code@mail.ru",
	"okay-d2@mail.ru",
	"ivannlaptevv16@mail.ru",
	"ivan.kim.9696@mail.ru",
	"rtyui.rtyu.2019@mail.ru",
	"5608013@mail.ru",
	"natuly.sid@mail.ru",
	"khaydaralikas9@mail.ru",
	"adam.boldyrev.77@mail.ru",
	"emaleevoleg@mail.ru",
	"gurshenko968@mail.ru",
	"zubachevaalla19851@mail.ru",
	"sergi.popov.70@mail.ru",
	"vip.wtwert@mail.ru",
	"vtope_ru@mail.ru",
	"andreypavlovxd@Mail.ru",
	"metcon-2002@mail.ru",
	"megostop@mail.ru",
	"a.sloeva@mail.ru",
	"gunasheva1981@mail.ru",
	"halinh@mail.ru",
	"vlad-nn92@mail.ru",
	"junglepin@mail.ru",
	"danilca0@mail.ru",
	"mylikespochta@mail.ru",
	"anairs@mail.ru",
	"podoba.007@mail.ru",
	"wb.all@mail.ru",
	"ranelchik@mail.ru",
	"SkaderFoxy@mail.ru",
	"isupov01394@mail.ru",
	"reeeop90@mail.ru",
	"elechka2208@mail.ru",
	"super.pasha.pasha@mail.ru",
	"susics@mail.ru",
	"Grachly@mail.ru",
	"mitronin.davidka@mail.ru",
	"astahin@mail.ru",
	"l.a.w.m@mail.ru",
	"jennifer199211@mail.ru",
	"monsh-sergei@mail.ru",
	"adisrep2018@mail.ru",
}

var EmailAddresses = getEmailAddresses(EmailStrings)

func getEmailAddresses(emails []string) []evmail.Address {
	addrs := make([]evmail.Address, len(emails))
	for i, email := range emails {
		addrs[i] = evmail.FromString(email)
	}

	return addrs
}
