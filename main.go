package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kennygrant/sanitize"
)

var client = &http.Client{}

func main() {

	req, err := http.NewRequest("POST", "https://www.register.ca/", nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	for i, name := range names {
		domain := sanitize.Path(name)
		ok := check(domain, resp.Cookies())
		if ok {
			fmt.Printf("\n[%d] %s: %v\n", i, domain, ok)
		}
	}
}

func check(domain string, cookies []*http.Cookie) bool {

	fmt.Printf(".")
	domain = domain[0 : len(domain)-2]
	//fmt.Printf("%s.ca: ", domain)

	host := "https://www.register.ca/en/register.htm"

	form := url.Values{}

	form.Add("domain", domain)
	form.Add("Submit", "+SEARCH+NOW+")
	form.Add("check", "check")

	req, err := http.NewRequest("POST", host, strings.NewReader(form.Encode()))

	req.Header.Add("Host", "www.register.ca")
	req.Header.Add("Refer", "https://www.register.ca/")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	for _, c := range cookies {
		req.AddCookie(c)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// use utfBody using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//html, _ := doc.Html()
	//fmt.Printf("RESP: %s\n", html)

	ok := false
	doc.Find(".greenPriceColor").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if text == "available" {
			ok = true
		}
	})

	//fmt.Printf("%v\n", ok)

	return ok
}

var names = []string{

	"abacá",
	"açoca",
	"adaca",
	"áfaca",
	"afeca",
	"aiacá",
	"aiocá",
	"aiucá",
	"álica",
	"amaca",
	"ameça",
	"amuca",
	"anacã",
	"ápoca",
	"araçá",
	"areca",
	"aricá",
	"aruca",
	"atacã",
	"avacá",
	"azaca",
	"azeca",
	"balça",
	"bançá",
	"barça",
	"bença",
	"bosca",
	"brica",
	"broça",
	"burca",
	"caiçá",
	"carca",
	"cesca",
	"chica",
	"chuca",
	"conca",
	"corca",
	"couça",
	"craca",
	"cruca",
	"cuaca",
	"cueca",
	"curca",
	"dinca",
	"dirca",
	"driça",
	"dueça",
	"época",
	"erica",
	"eroca",
	"eruca",
	"fanca",
	"fasca",
	"fôrça",
	"funca",
	"gorca",
	"graça",
	"guacá",
	"gurca",
	"haúça",
	"icica",
	"iluca",
	"ipaca",
	"ipeca",
	"ipuca",
	"isoca",
	"ítaca",
	"iveca",
	"jaacã",
	"joiça",
	"jouça",
	"lenca",
	"lerca",
	"liaça",
	"loiça",
	"lonca",
	"lorçá",
	"louça",
	"maica",
	"marcá",
	"maúça",
	"melca",
	"mfeca",
	"milca",
	"miúça",
	"moacã",
	"moeca",
	"moóca",
	"murça",
	"musca",
	"narça",
	"nhaca",
	"nhoca",
	"nooca",
	"norça",
	"oiacá",
	"oincá",
	"oraçá",
	"orica",
	"osaca",
	"oteca",
	"paica",
	"pança",
	"peaça",
	"piaçá",
	"pioca",
	"pirca",
	"piúca",
	"placá",
	"poiçá",
	"poncã",
	"posca",
	"praça",
	"procá",
	"pruca",
	"pseca",
	"queca",
	"quiçá",
	"ranca",
	"roncã",
	"ruaça",
	"saicã",
	"sancã",
	"sarça",
	"seiça",
	"sorça",
	"surça",
	"taacã",
	"tanca",
	"taoca",
	"tarca",
	"teaça",
	"tença",
	"terçã",
	"tinca",
	"toaca",
	"tonca",
	"torca",
	"traca",
	"trica",
	"tuaca",
	"uaicá",
	"uasca",
	"ujica",
	"uraca",
	"úrica",
	"uvaça",
	"verça",
	"vodca",
	"xorca",
	"zurca",
	"abisca",
	"ablaca",
	"acelca",
	"aciocá",
	"aivaca",
	"alasca",
	"alfaça",
	"alfeça",
	"áltica",
	"alvacá",
	"amança",
	"ambaca",
	"amioca",
	"amurca",
	"andaca",
	"ângica",
	"anteca",
	"apiacá",
	"araicá",
	"arauca",
	"arnecã",
	"arouca",
	"arreçã",
	"asteca",
	"avenca",
	"aviaca",
	"avusca",
	"azteca",
	"babaca",
	"baboca",
	"bacaca",
	"bacoca",
	"baduca",
	"baiúca",
	"bajeca",
	"balaca",
	"banaca",
	"bataca",
	"bateca",
	"bazuca",
	"bebica",
	"beluca",
	"beraca",
	"biboca",
	"bijuca",
	"biloca",
	"birica",
	"biteca",
	"bituca",
	"boança",
	"bocaça",
	"boenca",
	"boiaca",
	"botica",
	"brasca",
	"bresca",
	"brisca",
	"broaça",
	"bruaca",
	"bunica",
	"buraca",
	"buricá",
	"butaca",
	"cabeça",
	"cafuca",
	"cainca",
	"cajaça",
	"calaça",
	"camacã",
	"camoca",
	"camucá",
	"canaca",
	"canica",
	"capaça",
	"capoca",
	"caraça",
	"careca",
	"caroca",
	"caruca",
	"catacá",
	"chasca",
	"chilca",
	"chirca",
	"chouca",
	"chulca",
	"cicica",
	"cicoca",
	"cloaca",
	"colaça",
	"conaça",
	"corica",
	"coroca",
	"cotecá",
	"cotica",
	"covoca",
	"crauçá",
	"creaca",
	"crença",
	"crosca",
	"cuenca",
	"cuiaca",
	"culaça",
	"cumacá",
	"curaçá",
	"curica",
	"curuçá",
	"dânaca",
	"dédica",
	"dijiça",
	"dixiça",
	"doença",
	"duarca",
	"eiveca",
	"embaca",
	"embeca",
	"endaca",
	"enjaca",
	"entaca",
	"enxaca",
	"eparca",
	"ervaca",
	"espica",
	"estaca",
	"eureca",
	"exorca",
	"fajeca",
	"falaca",
	"fataça",
	"favica",
	"fedoca",
	"feleca",
	"feloca",
	"fetuca",
	"filaça",
	"fílica",
	"fitaça",
	"fiunça",
	"foboca",
	"fogaça",
	"foneca",
	"foreca",
	"frança",
	"frosca",
	"fubaca",
	"fubica",
	"fuboca",
	"fugeca",
	"fúlica",
	"funeca",
	"galeca",
	"genica",
	"geóica",
	"goiaca",
	"grança",
	"grauçá",
	"gresca",
	"guaicá",
	"guasca",
	"guaucá",
	"guiaca",
	"guriçá",
	"guruçá",
	"halaca",
	"hanucá",
	"hérica",
	"homaça",
	"iapuçá",
	"ibiaçá",
	"ibioca",
	"ilarca",
	"impaca",
	"incoca",
	"incuça",
	"indaca",
	"indicã",
	"inhaca",
	"inveca",
	"inzica",
	"ioioca",
	"itaóca",
	"jabeca",
	"jacicá",
	"janaca",
	"janeca",
	"japuçá",
	"játaca",
	"jateca",
	"jaticá",
	"jetuca",
	"jijoca",
	"jitica",
	"judoca",
	"labaça",
	"lacacã",
	"lamuca",
	"lanuça",
	"larica",
	"laroca",
	"laruça",
	"lediça",
	"liança",
	"limaça",
	"liraça",
	"lodaça",
	"lojeca",
	"lojica",
	"loteca",
	"lufica",
	"lupaça",
	"lupoça",
	"lutoca",
	"luzica",
	"maboca",
	"macaça",
	"maduca",
	"magaça",
	"maiacá",
	"mainça",
	"maioca",
	"malaca",
	"mamica",
	"manaça",
	"manicá",
	"mapuca",
	"maracá",
	"mareca",
	"maricá",
	"maroca",
	"maruca",
	"masacá",
	"mataca",
	"matuca",
	"maunça",
	"maxaca",
	"mebuca",
	"meleca",
	"meroça",
	"merucá",
	"meteca",
	"metoca",
	"mijoca",
	"miloca",
	"mimaça",
	"mirica",
	"moaica",
	"mobica",
	"mocica",
	"mococa",
	"modeca",
	"mogica",
	"moloca",
	"monoca",
	"motoca",
	"motuca",
	"mubica",
	"muçaça",
	"mucica",
	"mumuca",
	"mupeça",
	"muracá",
	"muruca",
	"musaça",
	"mutaca",
	"mutuca",
	"muvuca",
	"muxeca",
	"nabiça",
	"nacaca",
	"nadica",
	"narica",
	"nasica",
	"negaça",
	"neguça",
	"niroca",
	"octeca",
	"olmeca",
	"ombaca",
	"ondaca",
	"onhoca",
	"ooteca",
	"oureça",
	"ovença",
	"paçoca",
	"paduca",
	"paliça",
	"panaca",
	"pânica",
	"panoca",
	"pareca",
	"paricá",
	"patica",
	"patuca",
	"paxicá",
	"peança",
	"pearça",
	"peasca",
	"peliça",
	"pepeca",
	"pepiça",
	"peruca",
	"petica",
	"péxica",
	"piarça",
	"picica",
	"pigaça",
	"pijuca",
	"piloca",
	"pinaça",
	"piracá",
	"pitaca",
	"pivoca",
	"poçoca",
	"poçuca",
	"podica",
	"ponaca",
	"popuca",
	"poracá",
	"poruca",
	"priaca",
	"puçuca",
	"pujacá",
	"punaca",
	"púraca",
	"puruca",
	"putuca",
	"quiaça",
	"quinca",
	"rebaçã",
	"rebeca",
	"riosca",
	"rolaça",
	"rosaça",
	"sacaca",
	"sacoca",
	"sadaca",
	"sajica",
	"sanica",
	"saraçá",
	"sárica",
	"savica",
	"sebaça",
	"sedeca",
	"sêneca",
	"sênica",
	"sésica",
	"síreca",
	"sirica",
	"sodica",
	"soneca",
	"soroca",
	"sumaca",
	"tabeca",
	"taboca",
	"tacacá",
	"taioca",
	"taluca",
	"tamiça",
	"taniça",
	"tanoca",
	"tauoca",
	"tavoca",
	"tibaca",
	"ticaca",
	"ticuca",
	"tijuca",
	"timaca",
	"tinoca",
	"tipuca",
	"tiraca",
	"tiriça",
	"tobaca",
	"tococa",
	"tovaca",
	"trença",
	"triaca",
	"trouça",
	"truaca",
	"tubaca",
	"tudaca",
	"tuiúca",
	"túnica",
	"turica",
	"uanica",
	"uapaca",
	"uapuçá",
	"unhaca",
	"urraca",
	"uruoca",
	"usança",
	"valoca",
	"vidoca",
	"vilaça",
	"viloca",
	"vocaca",
	"xereca",
	"xicaca",
	"xitaca",
	"xixica",
	"xoroca",
	"zalaca",
	"zaruca",
	"zoteca",
	"abitica",
	"acaiaca",
	"adiciça",
	"afavaca",
	"albarca",
	"alcorça",
	"aldonça",
	"alenica",
	"alferça",
	"alharca",
	"alhorca",
	"alítica",
	"almança",
	"alparca",
	"alverca",
	"amálaca",
	"amanaca",
	"amauaca",
	"amigaça",
	"anclaca",
	"andança",
	"angracá",
	"arabaca",
	"aramaçã",
	"arapoca",
	"arapuçá",
	"araricá",
	"arataca",
	"araveca",
	"areosca",
	"aritica",
	"arjunça",
	"arumaçá",
	"asiarca",
	"astracã",
	"atérica",
	"aularca",
	"axaraca",
	"babanca",
	"babaúca",
	"baguaça",
	"baianca",
	"bainica",
	"bajunça",
	"balancá",
	"balzaca",
	"barbacã",
	"barbica",
	"barcaça",
	"batanca",
	"bebinca",
	"beiçoca",
	"belfaça",
	"benfica",
	"benteca",
	"bernica",
	"bespiça",
	"betouca",
	"bicanca",
	"bilhoca",
	"bilosca",
	"biqueca",
	"birosca",
	"bispiça",
	"bisteca",
	"bitruca",
	"bochaca",
	"bojança",
	"bonsuça",
	"bordaça",
	"borraça",
	"boudica",
	"brasuca",
	"brumaça",
	"bundaça",
	"burjaca",
	"caacica",
	"caapucá",
	"caatica",
	"cachiça",
	"cachoça",
	"cachuça",
	"cadurça",
	"cagança",
	"caiança",
	"caiçaca",
	"caitoca",
	"calanca",
	"caldaça",
	"calmiça",
	"cambacã",
	"cambica",
	"cambucá",
	"camoeca",
	"camuacá",
	"camueca",
	"camurça",
	"canaica",
	"canjica",
	"carcaça",
	"cardoça",
	"carioca",
	"carnaça",
	"carniça",
	"carnuça",
	"carraca",
	"carroça",
	"carruca",
	"casarca",
	"catança",
	"catraca",
	"chainça",
	"chamiça",
	"chapeca",
	"chianca",
	"chicaca",
	"chinoca",
	"chorica",
	"colança",
	"comarca",
	"combaca",
	"comboca",
	"cordeca",
	"cornaca",
	"corriça",
	"cousica",
	"covanca",
	"crejica",
	"crifeca",
	"cruciça",
	"cualicá",
	"cúltica",
	"cumbaca",
	"cumbeca",
	"cumbica",
	"cunarca",
	"curiaca",
	"curriça",
	"curuaca",
	"curvaça",
	"cussuca",
	"datisca",
	"desmuca",
	"dialaca",
	"dicerca",
	"dombica",
	"dondoca",
	"dorasca",
	"durança",
	"enleaça",
	"enoteca",
	"enxarca",
	"ericica",
	"esmiúça",
	"estança",
	"etnarca",
	"exoteca",
	"fachaca",
	"fachoca",
	"faiança",
	"faiença",
	"falança",
	"fanzoca",
	"fatiaça",
	"fedença",
	"ferraça",
	"festuca",
	"fetusca",
	"figança",
	"filança",
	"finança",
	"fitança",
	"folerca",
	"folheca",
	"fondiça",
	"fonisca",
	"fonseca",
	"fonteca",
	"forleca",
	"formica",
	"forreca",
	"fugança",
	"fumbeca",
	"furreca",
	"furrica",
	"fuzarca",
	"galeaça",
	"galhaça",
	"galhuça",
	"garauçá",
	"genarca",
	"gentaça",
	"ginarca",
	"goitacá",
	"gonarca",
	"grimaça",
	"guaiacá",
	"guajicá",
	"guanacá",
	"guapeca",
	"guaruça",
	"guaxica",
	"guiraca",
	"háltica",
	"herança",
	"heureca",
	"holerca",
	"ibiboca",
	"iboboca",
	"inética",
	"ipojuca",
	"irapoca",
	"iriceca",
	"irideca",
	"itamaca",
	"itapuca",
	"jagonça",
	"jamaica",
	"jaupoca",
	"jerarca",
	"jéssica",
	"jipioca",
	"jipooca",
	"juapoca",
	"jurueca",
	"juvenca",
	"lactuca",
	"lagasca",
	"lagoaça",
	"lamisca",
	"lampaça",
	"landoca",
	"laqueca",
	"lectica",
	"léplica",
	"lequeça",
	"licença",
	"limpaça",
	"linhaça",
	"loiraça",
	"luxança",
	"luzença",
	"machacá",
	"magaíça",
	"magarça",
	"maigaça",
	"maiorca",
	"maipoca",
	"maitaca",
	"maituca",
	"malança",
	"malpica",
	"manduça",
	"mangaca",
	"mangoça",
	"manjica",
	"manjuca",
	"manteca",
	"mântica",
	"mantuca",
	"manusca",
	"marasca",
	"margaça",
	"marioca",
	"marosca",
	"massacá",
	"massoca",
	"mástica",
	"mastuca",
	"matança",
	"matinca",
	"matusca",
	"mazorca",
	"mazurca",
	"medrica",
	"melança",
	"membeca",
	"menaica",
	"menarca",
	"mendoca",
	"merarca",
	"merreca",
	"meruoca",
	"micreca",
	"milhaça",
	"minorca",
	"mirmica",
	"moalaca",
	"mocheca",
	"molhaça",
	"mombaça",
	"mombuca",
	"monarca",
	"morança",
	"morraça",
	"movença",
	"mualaca",
	"muçuíca",
	"mucuoca",
	"mudança",
	"mumbaca",
	"mumbica",
	"mumbuca",
	"mundiça",
	"munduca",
	"muqueca",
	"murraça",
	"musgaça",
	"mussacá",
	"mussuca",
	"mutança",
	"nagaica",
	"naiveca",
	"narosca",
	"negarça",
	"nerusca",
	"nevasca",
	"nevoaça",
	"nhaneca",
	"nhanica",
	"nianeca",
	"nomarca",
	"notolca",
	"novença",
	"opórica",
	"oreóica",
	"ouvença",
	"oxiteca",
	"pagarca",
	"palhoça",
	"pantaça",
	"papança",
	"páprica",
	"parança",
	"parasca",
	"parbiça",
	"pardoca",
	"patença",
	"patruça",
	"pedauca",
	"peidoca",
	"peitaca",
	"peitica",
	"pepoaça",
	"pepuaça",
	"periaca",
	"perisca",
	"pernoca",
	"pernuca",
	"perruca",
	"pértica",
	"perunca",
	"piaçoca",
	"picanca",
	"pinhoca",
	"piraaca",
	"pirisca",
	"pirosca",
	"pitaica",
	"pitança",
	"polieca",
	"pomboca",
	"poqueca",
	"porcoça",
	"possuca",
	"potança",
	"prédica",
	"préfica",
	"prexeca",
	"proença",
	"protoca",
	"provica",
	"pubarca",
	"pujança",
	"quereca",
	"quibaca",
	"quiboça",
	"quiçaça",
	"quicuca",
	"quifaça",
	"quifuça",
	"quinica",
	"quirica",
	"quitaca",
	"quizaca",
	"rabanca",
	"ranhoca",
	"rapança",
	"rapioca",
	"redoiça",
	"redouça",
	"réplica",
	"ressoca",
	"rodoiça",
	"rodouça",
	"rolança",
	"rubisca",
	"ruivaca",
	"sabença",
	"sabraca",
	"sanisca",
	"sanjica",
	"sarguça",
	"satarca",
	"secança",
	"selisca",
	"silerca",
	"sinarca",
	"siríaca",
	"sirsaca",
	"sitarca",
	"sitioca",
	"sobesca",
	"súnhaca",
	"tabanca",
	"taipoca",
	"tairoca",
	"talisca",
	"tambica",
	"tambuca",
	"tapiaca",
	"tapioca",
	"tapuoca",
	"tarruca",
	"teipoca",
	"telarca",
	"terarca",
	"terluca",
	"testaça",
	"tibraca",
	"timarca",
	"tipisca",
	"tolteca",
	"toparca",
	"trapuca",
	"treliça",
	"tuapoca",
	"tubança",
	"tuiaíca",
	"tuiroca",
	"uaitacá",
	"uramaçá",
	"urapuca",
	"uruçuca",
	"valença",
	"valesca",
	"valinca",
	"varasca",
	"varunca",
	"velança",
	"velisca",
	"verruca",
	"vespiça",
	"vinheca",
	"vinhuça",
	"virosca",
	"vivença",
	"xanduca",
	"zarasca",
	"zíncica",
	"zoética",
	"zondoca",
	"zooteca",
	"zuleica",
}
