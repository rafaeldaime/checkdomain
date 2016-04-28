package main

var namesDe = []string{
	"aode",
	"badé",
	"bidê",
	"bode",
	"cadê",
	"cide",
	"codé",
	"cude",
	"gadé",
	"godê",
	"gude",
	"hode",
	"hude",
	"igde",
	"jade",
	"lide",
	"made",
	"ocde",
	"oide",
	"onde",
	"padê",
	"rede",
	"rude",
	"sade",
	"sede",
	"uade",
	"vide",
	"abedê",
	"ácade",
	"açude",
	"alodê",
	"amade",
	"ápode",
	"aredê",
	"arode",
	"balde",
	"bendé",
	"binde",
	"bonde",
	"caide",
	"cande",
	"chade",
	"chede",
	"clade",
	"clide",
	"conde",
	"credé",
	"crude",
	"dendê",
	"díade",
	"donde",
	"éfode",
	"égide",
	"eledê",
	"êmide",
	"fóide",
	"forde",
	"frade",
	"geode",
	"glide",
	"gonde",
	"grade",
	"grede",
	"grode",
	"guedé",
	"gunde",
	"gurde",
	"haidê",
	"halde",
	"heide",
	"hilde",
	"iande",
	"idade",
	"irade",
	"ixode",
	"jalde",
	"jande",
	"jeúde",
	"joede",
	"lande",
	"laúde",
	"leide",
	"lende",
	"leude",
	"linde",
	"liode",
	"lóide",
	"lorde",
	"mandê",
	"mende",
	"milde",
	"miode",
	"miúde",
	"molde",
	"mondé",
	"mundé",
	"naide",
	"neide",
	"nende",
	"nonde",
	"odede",
	"ogudê",
	"olede",
	"olodé",
	"pande",
	"piadé",
	"pundé",
	"quedê",
	"raide",
	"reide",
	"ríade",
	"rondé",
	"saide",
	"saúde",
	"siode",
	"slide",
	"suede",
	"tarde",
	"tíade",
	"trude",
	"unade",
	"verde",
	"zaide",
	"zande",
	"zende",
	"zóide",
	"abdude",
	"abside",
	"áclide",
	"acnide",
	"acorde",
	"ácrode",
	"adrede",
	"afrode",
	"ailude",
	"alaíde",
	"alarde",
	"alaúde",
	"alcade",
	"alcide",
	"alóide",
	"ambude",
	"anaide",
	"ápside",
	"árcade",
	"arolde",
	"arpade",
	"arvade",
	"asdode",
	"áspide",
	"ataíde",
	"ataúde",
	"avonde",
	"axóide",
	"babadé",
	"bajude",
	"báside",
	"bátide",
	"bedade",
	"berede",
	"beride",
	"bigode",
	"bípede",
	"bólide",
	"buende",
	"bunode",
	"burude",
	"cabide",
	"cápide",
	"cegude",
	"céride",
	"cerude",
	"chande",
	"cheide",
	"cidade",
	"cigude",
	"citode",
	"claude",
	"cleide",
	"démodé",
	"dicode",
	"dípode",
	"dóride",
	"dríade",
	"duende",
	"ebande",
	"ecóide",
	"elande",
	"elende",
	"enéade",
	"eneide",
	"equede",
	"erande",
	"êupode",
	"filode",
	"fiorde",
	"flande",
	"fólade",
	"frande",
	"fraude",
	"geóide",
	"glande",
	"gleide",
	"gourde",
	"grande",
	"graúde",
	"greide",
	"guende",
	"hadade",
	"harade",
	"harude",
	"helede",
	"hélode",
	"henide",
	"hesede",
	"híbode",
	"hilode",
	"hióide",
	"ialodê",
	"iápide",
	"iaundé",
	"ieúbde",
	"ijinde",
	"ilíade",
	"inandê",
	"iódide",
	"jagode",
	"jápide",
	"jerede",
	"lagide",
	"lâmede",
	"lápide",
	"lépade",
	"licode",
	"lípide",
	"litode",
	"loxode",
	"mabode",
	"mamede",
	"mamude",
	"matade",
	"medade",
	"mênade",
	"méride",
	"metade",
	"midade",
	"mióide",
	"mobede",
	"molide",
	"mônade",
	"monede",
	"muende",
	"nabade",
	"náiade",
	"naunde",
	"nhunde",
	"nômade",
	"nucode",
	"ogcode",
	"omóide",
	"oncode",
	"oótide",
	"oquedê",
	"oréade",
	"oróide",
	"oroode",
	"ovóide",
	"pagode",
	"pajade",
	"pálade",
	"palode",
	"palude",
	"parede",
	"peinde",
	"pevide",
	"pióide",
	"pípede",
	"píxide",
	"quende",
	"ráfide",
	"rágade",
	"rípide",
	"séride",
	"stande",
	"tágide",
	"talide",
	"talude",
	"tríade",
	"trôade",
	"validé",
	"víride",
	"zedade",
	"zoóide",
	"abecedê",
	"abujede",
	"adalide",
	"adápide",
	"adônide",
	"aeróide",
	"aftóide",
	"aganide",
	"agátide",
	"albende",
	"alcaide",
	"alcalde",
	"algóide",
	"alidade",
	"alípede",
	"amerade",
	"amizade",
	"ancrode",
	"andonde",
	"andrade",
	"anidade",
	"antíade",
	"antóide",
	"aridade",
	"astride",
	"atitude",
	"atlóide",
	"azélide",
	"bacondê",
	"bagdade",
	"balurde",
	"bapende",
	"báquide",
	"barvade",
	"batilde",
	"batóide",
	"báucide",
	"beldade",
	"benilde",
	"bicaude",
	"bilbode",
	"bisonde",
	"bondade",
	"brômide",
	"bunóide",
	"cabonde",
	"caconde",
	"cacundê",
	"calonde",
	"candide",
	"canindé",
	"canóide",
	"capande",
	"capaude",
	"cassade",
	"catende",
	"caxinde",
	"caxundé",
	"cebóide",
	"celiode",
	"cércide",
	"ceróide",
	"céspede",
	"cestode",
	"cíclade",
	"ciliode",
	"ciminde",
	"cimóide",
	"citóide",
	"cladode",
	"clâmide",
	"clóride",
	"cnêmide",
	"cobarde",
	"cocóide",
	"colóide",
	"cólpode",
	"cômpede",
	"conóide",
	"coróide",
	"covarde",
	"críside",
	"cubóide",
	"cúspide",
	"custode",
	"danaide",
	"debalde",
	"degradê",
	"deidade",
	"dinsede",
	"dioxide",
	"dípcade",
	"dípsade",
	"domóide",
	"duidade",
	"durandé",
	"duríade",
	"ebálide",
	"ecodidé",
	"edípode",
	"efátide",
	"efélide",
	"egópode",
	"embalde",
	"epípode",
	"epúlide",
	"erípede",
	"eslaide",
	"espelde",
	"estande",
	"etmóide",
	"exocade",
	"exômide",
	"exópode",
	"facóide",
	"falóide",
	"ficóide",
	"filóide",
	"fisóide",
	"fitóide",
	"flômide",
	"fucóide",
	"fusóide",
	"gadóide",
	"galeode",
	"ganóide",
	"gelóide",
	"glícide",
	"glifode",
	"glótide",
	"glúcide",
	"goniode",
	"grilode",
	"gueledé",
	"halóide",
	"hamaide",
	"hasside",
	"helóide",
	"hemóide",
	"herdade",
	"heróide",
	"hipóide",
	"homóide",
	"hóspede",
	"humilde",
	"ibéride",
	"icabode",
	"icodidé",
	"ilíride",
	"imbunde",
	"isópode",
	"issende",
	"jacundê",
	"jeguedê",
	"jequedê",
	"jupiede",
	"lacondé",
	"láctide",
	"lambude",
	"lâmpide",
	"lanzude",
	"lepóide",
	"lépride",
	"lícnide",
	"liconde",
	"liópode",
	"lipóide",
	"litóide",
	"litonde",
	"loconde",
	"lotóide",
	"lupóide",
	"maamude",
	"maconde",
	"macunde",
	"madride",
	"maldade",
	"mambude",
	"mariode",
	"matilde",
	"maulide",
	"mavilde",
	"melinde",
	"merinde",
	"mesóide",
	"mezande",
	"milóide",
	"milorde",
	"miríade",
	"místide",
	"mitonde",
	"mixóide",
	"mnióide",
	"moamede",
	"molóide",
	"monóide",
	"mormode",
	"mucaade",
	"muçande",
	"mucóide",
	"mucunde",
	"mudende",
	"mugande",
	"mulende",
	"mulóide",
	"nanóide",
	"nébride",
	"necrode",
	"nepóide",
	"nereide",
	"netóide",
	"ninrode",
	"nodóide",
	"notiode",
	"ocípode",
	"ofióide",
	"ogdóade",
	"omuande",
	"onônide",
	"oozóide",
	"opióide",
	"ormasde",
	"orneode",
	"ortóide",
	"osíride",
	"osteíde",
	"ostende",
	"ovótide",
	"oxálide",
	"ozônide",
	"pêmpade",
	"péptide",
	"perinde",
	"peróide",
	"pérside",
	"picnode",
	"picóide",
	"piedade",
	"piéride",
	"piróide",
	"pisóide",
	"pissode",
	"pitíade",
	"plêiade",
	"podóide",
	"políade",
	"poliide",
	"prótide",
	"provede",
	"psálide",
	"psécade",
	"psicode",
	"ptéride",
	"quipede",
	"rabalde",
	"ragóide",
	"ramalde",
	"ranóide",
	"rapsode",
	"rebelde",
	"ritende",
	"ritonde",
	"rizóide",
	"rodóide",
	"sameode",
	"sarcode",
	"saudade",
	"senóide",
	"sépside",
	"sifóide",
	"sílfide",
	"símpode",
	"soidade",
	"tacóide",
	"talmude",
	"talóide",
	"tanaide",
	"tanóide",
	"táquide",
	"teróide",
	"tétrade",
	"tiarode",
	"tifóide",
	"tígride",
	"tiróide",
	"tólpide",
	"toróide",
	"toxóide",
	"trenode",
	"tríbade",
	"trigode",
	"trinode",
	"trípode",
	"trófide",
	"ulcóide",
	"umidade",
	"unípede",
	"upgrade",
	"urópode",
	"vaidade",
	"valóide",
	"vedóide",
	"verdade",
	"vildade",
	"viróide",
	"virtude",
	"voborde",
	"vocóide",
	"voivode",
	"vontade",
	"xifóide",
	"xilóide",
	"zenaide",
	"zenilde",
	"zoraide",
	"abaconde",
	"abaxóide",
	"acaróide",
	"acidóide",
	"aclâmide",
	"acridade",
	"acrípede",
	"acritude",
	"actenode",
	"acuidade",
	"adelaide",
	"adenóide",
	"adonaide",
	"adoníade",
	"aerópode",
	"agamóide",
	"agatóide",
	"agráfide",
	"ajugóide",
	"albípede",
	"aleirode",
	"aleurode",
	"alfenide",
	"alisóide",
	"altitude",
	"alvaiade",
	"alvalade",
	"amaróide",
	"ambílide",
	"ambitude",
	"amebóide",
	"amibóide",
	"andróide",
	"anecóide",
	"aneróide",
	"anfípode",
	"anisóide",
	"antípode",
	"anuidade",
	"aráquide",
	"arbelode",
	"arfaxade",
	"arilóide",
	"aristide",
	"artêmide",
	"asaróide",
	"ascáride",
	"aseidade",
	"asnidade",
	"astróide",
	"atenaíde",
	"atetóide",
	"atrípede",
	"axonóide",
	"azotóide",
	"azulóide",
	"bacáride",
	"báctride",
	"balânide",
	"bantóide",
	"baridade",
	"basidade",
	"báuquide",
	"belópode",
	"belverde",
	"berebedé",
	"berlinde",
	"bigarade",
	"bilebode",
	"bilióide",
	"bissóide",
	"bolbóide",
	"bonafide",
	"bracóide",
	"budidade",
	"bugróide",
	"bulbóide",
	"butóride",
	"cacábide",
	"cachondé",
	"cachundé",
	"cactóide",
	"calátide",
	"caltóide",
	"cambonde",
	"candindé",
	"caquinde",
	"caridade",
	"carótide",
	"carpóide",
	"catópode",
	"cauliode",
	"caulóide",
	"cavidade",
	"celépode",
	"celópode",
	"cenópode",
	"cercnide",
	"cerítide",
	"cestóide",
	"cetilide",
	"chasside",
	"chefóide",
	"ciamóide",
	"cianóide",
	"ciatóide",
	"ciclóide",
	"cicnóide",
	"cilípede",
	"cipéride",
	"cissóide",
	"cistóide",
	"ciuróide",
	"cividade",
	"clenóide",
	"clinóide",
	"cloróide",
	"clotilde",
	"clotóide",
	"coanóide",
	"comidade",
	"concóide",
	"concorde",
	"confrade",
	"copépode",
	"cordóide",
	"coreóide",
	"corióide",
	"cormóide",
	"corônide",
	"corvóide",
	"cotidade",
	"coxípede",
	"coxópode",
	"cricóide",
	"crinóide",
	"cruzóide",
	"ctenóide",
	"cupróide",
	"dacriode",
	"dasípode",
	"decápode",
	"deltóide",
	"demitade",
	"dermóide",
	"desmóide",
	"diapside",
	"dicópode",
	"dilépide",
	"diplóide",
	"diquende",
	"discóide",
	"dodônide",
	"ecbólade",
	"egestade",
	"elanóide",
	"endópode",
	"enérgide",
	"enídride",
	"epulóide",
	"eqüidade",
	"equínide",
	"equióide",
	"equípede",
	"ergátide",
	"ericóide",
	"erinóide",
	"eriópode",
	"erucóide",
	"escáfide",
	"esfecode",
	"espilode",
	"estenode",
	"estidade",
	"estróide",
	"eucáride",
	"eufótide",
	"eumênide",
	"euplóide",
	"eurípede",
	"fasmóide",
	"fealdade",
	"feridade",
	"fibróide",
	"fieldade",
	"filópode",
	"fimatode",
	"finidade",
	"finitude",
	"fisálide",
	"fisápode",
	"fisópode",
	"fixidade",
	"flomóide",
	"fofidade",
	"fungóide",
	"fusípede",
	"gabróide",
	"gafidade",
	"galatode",
	"galeóide",
	"gastrode",
	"ginópode",
	"glenóide",
	"globóide",
	"gobióide",
	"gonópode",
	"graióide",
	"gramondé",
	"grecóide",
	"grupóide",
	"habitude",
	"haplóide",
	"hebetude",
	"helcóide",
	"helióide",
	"hematode",
	"henadade",
	"hendíade",
	"herpolde",
	"hexápode",
	"hialóide",
	"hidátide",
	"hidríade",
	"hidróide",
	"hienóide",
	"himenode",
	"hipenode",
	"hipnóide",
	"hipópode",
	"holópode",
	"holóside",
	"homópode",
	"icoróide",
	"icterode",
	"ictióide",
	"ignípede",
	"inacóide",
	"ivaneide",
	"ivanilde",
	"jaspóide",
	"labróide",
	"lacmóide",
	"lactóide",
	"lagópede",
	"lagópode",
	"lâmpride",
	"latípede",
	"latíride",
	"latitude",
	"lealdade",
	"lenidade",
	"lepóride",
	"lepróide",
	"leucóide",
	"levidade",
	"levípede",
	"levuride",
	"lianóide",
	"licitude",
	"licópode",
	"linfóide",
	"linópode",
	"lipáride",
	"lipitude",
	"lirióide",
	"lobópode",
	"lofópode",
	"lorípede",
	"luperode",
	"mafamede",
	"magmóide",
	"mamaindê",
	"mambunde",
	"mandande",
	"maquinde",
	"mastóide",
	"matróide",
	"mecópode",
	"megápode",
	"melitode",
	"mesópode",
	"metópide",
	"metópode",
	"mielóide",
	"milépede",
	"milharde",
	"milhorde",
	"milípede",
	"miraonde",
	"mirtóide",
	"mitauadê",
	"mixópode",
	"moçâmede",
	"mocidade",
	"mohâmede",
	"molípede",
	"monópode",
	"morfóide",
	"morópode",
	"munhande",
	"muquende",
	"muscóide",
	"mussende",
	"nanópode",
	"natabude",
	"necedade",
	"nefróide",
	"negróide",
	"nematode",
	"nemópode",
	"neuróide",
	"ninfóide",
	"niquiode",
	"nitidade",
	"notópode",
	"nudípede",
	"obovóide",
	"oceânide",
	"ocnerode",
	"ocrópode",
	"octípede",
	"octópode",
	"oficlide",
	"ofiópode",
	"ofiúride",
	"omacunde",
	"omáspide",
	"oneirode",
	"onfalode",
	"oniróide",
	"orizóide",
	"ornêmide",
	"orobóide",
	"ortálide",
	"ortópode",
	"osteóide",
	"ovalóide",
	"oxiúride",
	"parápode",
	"paridade",
	"parótide",
	"parúlide",
	"parvóide",
	"pelâmide",
	"peltóide",
	"pepônide",
	"percóide",
	"perípode",
	"perópode",
	"pigmóide",
	"pigópode",
	"pilípede",
	"pinípede",
	"pirálide",
	"pirâmide",
	"placóide",
	"pleópode",
	"pliópode",
	"poetóide",
	"polípode",
	"pomerode",
	"potâmide",
	"prasóide",
	"promonde",
	"psicóide",
	"pteróide",
	"puridade",
	"quelóide",
	"queróide",
	"quiçonde",
	"quietude",
	"quilíade",
	"quinóide",
	"quiônide",
	"quitandê",
	"quitunde",
	"rabdóide",
	"radióide",
	"ralidade",
	"ranfóide",
	"realdade",
	"remípede",
	"rendendê",
	"resinide",
	"retípede",
	"retitude",
	"rizópode",
	"rombóide",
	"ruindade",
	"sabélide",
	"sacáride",
	"sagidade",
	"sanidade",
	"sarcóide",
	"sarônide",
	"sarópode",
	"sauróide",
	"serfóide",
	"sialóide",
	"sicidade",
	"sicínide",
	"sidéride",
	"sifílide",
	"siglóide",
	"sigmóide",
	"sinápode",
	"soledade",
	"solidade",
	"solípede",
	"sonípede",
	"sujidade",
	"tablóide",
	"talcóide",
	"tamaindé",
	"tarazede",
	"tecápode",
	"tecópode",
	"telépode",
	"tenióide",
	"terópode",
	"teutóide",
	"tigróide",
	"tilópode",
	"timidade",
	"tireóide",
	"tirsóide",
	"titânide",
	"tocodede",
	"tolidade",
	"toxidade",
	"toxótide",
	"tracóide",
	"tricóide",
	"tricorde",
	"trindade",
	"trocóide",
	"tumbandé",
	"tupióide",
	"uberdade",
	"ubiedade",
	"uirafede",
	"uncípede",
	"vacínide",
	"vagípede",
	"valáride",
	"validade",
	"valsóide",
	"valverde",
	"vanidade",
	"verbóide",
	"vimbunde",
	"visconde",
	"xistóide",
	"zebróide",
	"abancíade",
	"abantíade",
	"abraxóide",
	"abruptude",
	"acantóide",
	"acididade",
	"acinópode",
	"actinóide",
	"adelópode",
	"adenópode",
	"adicidade",
	"agasílide",
	"agenóride",
	"agilidade",
	"agláspide",
	"agripnode",
	"agróstide",
	"alantóide",
	"albicaude",
	"albiverde",
	"albumóide",
	"alcalóide",
	"alcebíade",
	"alcoólide",
	"aldegunde",
	"alergóide",
	"alginóide",
	"alislóide",
	"alismóide",
	"alissóide",
	"alozoóide",
	"alviverde",
	"amalópode",
	"amarílide",
	"amaritude",
	"amazônide",
	"amblípode",
	"ambrevade",
	"amplitude",
	"ampulóide",
	"anaeróide",
	"anagálide",
	"anarcóide",
	"ancilóide",
	"anciróide",
	"anelípede",
	"anexidade",
	"angelóide",
	"anginóide",
	"anguípede",
	"angulóide",
	"anisópode",
	"anopíxide",
	"anosidade",
	"ansiedade",
	"antenóide",
	"anteróide",
	"antozóide",
	"apirópode",
	"aracnóide",
	"arcêutide",
	"arcoverde",
	"arctitude",
	"arctópode",
	"arduidade",
	"argilóide",
	"argiróide",
	"aritmóide",
	"artrópode",
	"asconóide",
	"asseidade",
	"astacóide",
	"asteróide",
	"atelópode",
	"atricaude",
	"atripóide",
	"aulétride",
	"auriverde",
	"auroíride",
	"autacóide",
	"babilarde",
	"bacaróide",
	"bactéride",
	"balanóide",
	"barbípede",
	"bassáride",
	"batecondê",
	"beatitude",
	"belonóide",
	"bentidade",
	"berilóide",
	"bestidade",
	"bibiônide",
	"bicúspide",
	"biglábide",
	"blesidade",
	"bolbítide",
	"botrióide",
	"bradípode",
	"brasílide",
	"brevípede",
	"brissóide",
	"brutidade",
	"bubonóide",
	"budeidade",
	"cabridade",
	"calamóide",
	"calanóide",
	"caláspide",
	"calicóide",
	"camacinde",
	"camelóide",
	"cancróide",
	"cantáride",
	"capiverde",
	"caprípede",
	"capulóide",
	"cardióide",
	"caretóide",
	"cariátide",
	"caricóide",
	"catenóide",
	"caucálide",
	"cecrópide",
	"cefalóide",
	"ceguidade",
	"celomolde",
	"celsitude",
	"celulóide",
	"centípede",
	"centróide",
	"ceratóide",
	"chissonde",
	"cianípede",
	"cianópode",
	"ciclidade",
	"cidaróide",
	"cifonóide",
	"cinaróide",
	"cirrípede",
	"citozóide",
	"cladópode",
	"clemátide",
	"clitóride",
	"clunípede",
	"cocleóide",
	"coevidade",
	"coleópode",
	"coleróide",
	"conchóide",
	"condróide",
	"conicóide",
	"coracóide",
	"coralóide",
	"cormópode",
	"cornípede",
	"coronóide",
	"cotilóide",
	"cranióide",
	"creosende",
	"criptóide",
	"crisálide",
	"crisópide",
	"crisótide",
	"crocálide",
	"crocípede",
	"cromátide",
	"crueldade",
	"crugidade",
	"culicóide",
	"curtípede",
	"curvidade",
	"curvípede",
	"dacitóide",
	"dacrióide",
	"debilóide",
	"deliróide",
	"dendróide",
	"dentípode",
	"desumilde",
	"diabétide",
	"dictióide",
	"didélfide",
	"diecidade",
	"dioniside",
	"discópode",
	"disfólide",
	"diteríade",
	"divindade",
	"dolicóide",
	"dubiedade",
	"ebriedade",
	"ecceidade",
	"ecomônade",
	"edacidade",
	"edilidade",
	"efeméride",
	"elasípode",
	"elipsóide",
	"elitróide",
	"elurópode",
	"embolóide",
	"embrióide",
	"emulsóide",
	"encântide",
	"endrômide",
	"enotéride",
	"ensicaude",
	"entomóide",
	"enzimóide",
	"epeolóide",
	"epicáride",
	"epimênide",
	"epipíxide",
	"equinóide",
	"ergatóide",
	"eritróide",
	"escafóide",
	"escantude",
	"escaróide",
	"esfenóide",
	"esferóide",
	"esparóide",
	"espiróide",
	"esteróide",
	"estilóide",
	"eticidade",
	"etoposide",
	"eucnemide",
	"eucolóide",
	"eufrônide",
	"euganóide",
	"euisópode",
	"eunucóide",
	"euprimode",
	"europóide",
	"evanidade",
	"evisópode",
	"exetmóide",
	"faculdade",
	"falangode",
	"falcípede",
	"faneróide",
	"fatuidade",
	"fecalóide",
	"feculóide",
	"ferecíade",
	"fimatóide",
	"firmitude",
	"fissípede",
	"flavípede",
	"flexípede",
	"fortitude",
	"fosfátide",
	"fosfóride",
	"fossípede",
	"frialdade",
	"fulvípede",
	"gametóide",
	"gamozóide",
	"gangáride",
	"garibalde",
	"gemeidade",
	"gimnópode",
	"glicéride",
	"glicóside",
	"glossóide",
	"gonozóide",
	"grafióide",
	"grapsóide",
	"gratitude",
	"grecidade",
	"hebdômade",
	"helcópode",
	"helicóide",
	"helxípode",
	"hematóide",
	"heraclide",
	"hespéride",
	"hetacorde",
	"heteróide",
	"hexacorde",
	"hibridade",
	"hidatóide",
	"hidrópede",
	"himantode",
	"himenóide",
	"hiracóide",
	"histióide",
	"histópede",
	"holáspide",
	"homalóide",
	"hominóide",
	"hoplópode",
	"humanóide",
	"humildade",
	"icteróide",
	"igualdade",
	"ilicitude",
	"imanidade",
	"inatidade",
	"indigóide",
	"indolóide",
	"inetitude",
	"inimizade",
	"ipseidade",
	"ipsilóide",
	"irmandade",
	"isidióide",
	"isolépide",
	"isoplóide",
	"ivarimode",
	"joquebede",
	"judeidade",
	"justidade",
	"juventude",
	"labelóide",
	"laicidade",
	"lambdóide",
	"lampíride",
	"larinóide",
	"larvápode",
	"lassitude",
	"lemuróide",
	"lentípede",
	"lentitude",
	"leonítide",
	"lepidóide",
	"leproside",
	"leptópode",
	"leucêmide",
	"leucinode",
	"leucópode",
	"liberdade",
	"liceidade",
	"licnítide",
	"licópside",
	"lignípede",
	"limnatide",
	"limnétide",
	"limoníade",
	"liocóride",
	"liolépide",
	"liotíride",
	"lissótide",
	"longípede",
	"longitude",
	"macozóide",
	"macrópode",
	"magnitude",
	"magridade",
	"majestade",
	"malenóide",
	"mangualde",
	"mansidade",
	"margarode",
	"margélide",
	"marmáride",
	"medusóide",
	"melâmpode",
	"melanóide",
	"menfítede",
	"meráspide",
	"mesmidade",
	"metalóide",
	"metanóide",
	"micelóide",
	"micetóide",
	"micipóide",
	"micrópode",
	"micrótide",
	"mictíride",
	"miosótide",
	"miriápode",
	"miriópode",
	"mitilóide",
	"mixinóide",
	"moliônide",
	"mondágide",
	"monlevade",
	"monocorde",
	"morbidade",
	"morulóide",
	"mugilóide",
	"mujilóide",
	"multípede",
	"multitude",
	"nectópode",
	"nefelóide",
	"negritude",
	"nematóide",
	"neomígade",
	"nescidade",
	"neurópode",
	"nigrípede",
	"niilidade",
	"nimiedade",
	"ninfálide",
	"nixonóide",
	"nocuidade",
	"nucleóide",
	"nudicaude",
	"obviedade",
	"octacorde",
	"odonáiade",
	"odontóide",
	"oficleide",
	"ofidióide",
	"ofiuróide",
	"oligópode",
	"olimpíade",
	"omalópode",
	"ombridade",
	"omunhande",
	"onfalóide",
	"onoséride",
	"ootocóide",
	"opacidade",
	"orbitóide",
	"orfandade",
	"organóide",
	"ostracode",
	"oticidade",
	"oxienóide",
	"oxozônide",
	"pagetóide",
	"palmípede",
	"pancáride",
	"pantópode",
	"papulóide",
	"paquípode",
	"paranóide",
	"parástade",
	"parcidade",
	"parvidade",
	"patetóide",
	"paucidade",
	"paurópode",
	"pedaliode",
	"pegamóide",
	"pelecóide",
	"pereópode",
	"petalóide",
	"pezizóide",
	"pinacóide",
	"pináspide",
	"pinitóide",
	"piperóide",
	"piramóide",
	"pirenóide",
	"piritóide",
	"pirrópode",
	"pitecóide",
	"placitude",
	"placópode",
	"plasmóide",
	"platípede",
	"platípode",
	"platitude",
	"plecópode",
	"plectóide",
	"plenitude",
	"plexípede",
	"plumípede",
	"podestade",
	"polaróide",
	"polipóide",
	"portlande",
	"potestade",
	"pravidade",
	"prismóide",
	"prisópode",
	"promenade",
	"prosópide",
	"protíride",
	"protópode",
	"protúlide",
	"prumidade",
	"psilópode",
	"ptenópode",
	"puberdade",
	"quelípode",
	"quelópode",
	"quenômede",
	"quenópode",
	"querópode",
	"quetópode",
	"qüididade",
	"quilópode",
	"quimbande",
	"quindunde",
	"quingandé",
	"quirópode",
	"quissende",
	"quissonde",
	"quotidade",
	"rabdópode",
	"ramiverde",
	"raninóide",
	"rectitude",
	"resinóide",
	"retinóide",
	"rináspide",
	"ripsálide",
	"rizomóide",
	"rosineide",
	"rotalóide",
	"rotulóide",
	"rubrípede",
	"sacaróide",
	"saciedade",
	"salmácide",
	"samaróide",
	"satelóide",
	"saurópode",
	"sebastode",
	"selêucide",
	"senectude",
	"sepalóide",
	"sequidade",
	"seriedade",
	"serrípede",
	"servitude",
	"sesamóide",
	"seticaude",
	"seticorde",
	"siciônide",
	"siconóide",
	"sifilóide",
	"sifonóide",
	"siluróide",
	"sinágride",
	"sincáride",
	"sinérgide",
	"sinusóide",
	"sociedade",
	"solenóide",
	"sospitade",
	"suberóide",
	"sulfínide",
	"surdidade",
	"tanatóide",
	"taquípode",
	"tarsióide",
	"tarsípede",
	"tensidade",
	"tenuidade",
	"tenuípede",
	"teorópode",
	"tepalóide",
	"teratóide",
	"tericaude",
	"tetanóide",
	"tetrápode",
	"tilacóide",
	"torpidade",
	"torpitude",
	"torulóide",
	"torvidade",
	"traqueíde",
	"trematode",
	"tricópode",
	"trilépide",
	"trinidade",
	"triplóide",
	"tritêmide",
	"tritíride",
	"trombóide",
	"trucidade",
	"tuberóide",
	"turpitude",
	"ulceróide",
	"unicidade",
	"upanixade",
	"upeneóide",
	"upsilóide",
	"utilidade",
	"vacinóide",
	"vacuidade",
	"vacuólide",
	"vaginóide",
	"vaguidade",
	"variedade",
	"veleidade",
	"versidade",
	"vetustade",
	"vicindade",
	"vilaverde",
	"vinilóide",
	"viuvidade",
	"vizindade",
	"vocalóide",
	"xantópode",
	"abaliedade",
	"abutilóide",
	"acantópode",
	"acestóride",
	"acidáspide",
	"acriópside",
	"acuticaude",
	"adiantóide",
	"aduncidade",
	"agaricóide",
	"agatélpide",
	"agelenóide",
	"ageratóide",
	"agousidade",
	"alacridade",
	"almorávide",
	"alteridade",
	"alticópode",
	"ambliópode",
	"ambulípede",
	"amiantóide",
	"amidalóide",
	"amorfópode",
	"anantópode",
	"aneuplóide",
	"anfiaraíde",
	"angelitude",
	"anguilóide",
	"antracóide",
	"antropóide",
	"aquileóide",
	"aracnópode",
	"arcanidade",
	"aritenóide",
	"arquiabade",
	"asbestóide",
	"asclepíade",
	"ascosidade",
	"asperidade",
	"assirióide",
	"astacópode",
	"atiprimode",
	"atuosidade",
	"audacidade",
	"avessidade",
	"bacarióide",
	"bacivóride",
	"bacteróide",
	"baianidade",
	"balsamóide",
	"barbulóide",
	"basaltóide",
	"basicidade",
	"batracóide",
	"belacidade",
	"belemnóide",
	"belidióide",
	"beninidade",
	"bimastóide",
	"bipirâmide",
	"botriítide",
	"bracteóide",
	"braquióide",
	"braquípode",
	"brevicaude",
	"bronzípede",
	"bucalidade",
	"bupréstide",
	"calamidade",
	"calididade",
	"calosidade",
	"camarótide",
	"canforóide",
	"cantanhede",
	"carangóide",
	"carbonóide",
	"carcinóide",
	"cardiápode",
	"casagrande",
	"cassinóide",
	"castoróide",
	"catividade",
	"caucasóide",
	"cavalidade",
	"cefalópode",
	"celeridade",
	"celerípede",
	"cencrâmide",
	"cenosidade",
	"centrópode",
	"ceratixode",
	"ceratópode",
	"cerebróide",
	"cerosidade",
	"cerulípede",
	"cicadeóide",
	"cicláspide",
	"ciclopóide",
	"cimbalóide",
	"cinctípede",
	"civilidade",
	"cloritóide",
	"colubróide",
	"comicidade",
	"comodidade",
	"compiedade",
	"completude",
	"compsótide",
	"concretude",
	"condilóide",
	"conicidade",
	"conidióide",
	"coriariade",
	"coronópode",
	"corticóide",
	"cossenóide",
	"crassidade",
	"crassípede",
	"crassitude",
	"cretinóide",
	"criptópode",
	"crisalóide",
	"criticóide",
	"crotalóide",
	"cruoridade",
	"cteniópode",
	"cuatimundé",
	"curulidade",
	"curvicaude",
	"debilidade",
	"dermatóide",
	"dessuetude",
	"destridade",
	"desunidade",
	"dialeurode",
	"dicacidade",
	"dicteríade",
	"didelfóide",
	"didemnóide",
	"dieptápode",
	"difiozóide",
	"difteróide",
	"difusidade",
	"digitálide",
	"dinocaride",
	"dioicidade",
	"docilidade",
	"dolicópode",
	"dolosidade",
	"donosidade",
	"eceticóide",
	"ectetmóide",
	"ectofleode",
	"eczemátide",
	"egosséride",
	"elatinóide",
	"eneatéride",
	"entreverde",
	"eocrinóide",
	"epicântide",
	"epidérmide",
	"epidípnide",
	"epitelóide",
	"epitoxóide",
	"equinípede",
	"equinópode",
	"eritemóide",
	"eritrópode",
	"escafópode",
	"esclereide",
	"esculóside",
	"escuridade",
	"esplenóide",
	"espongóide",
	"esponjóide",
	"esquimóide",
	"esquizóide",
	"estenópode",
	"esticópode",
	"estilópode",
	"estomápode",
	"estratóide",
	"estrofóide",
	"etiguidade",
	"eucestóide",
	"eucopépode",
	"eudrilóide",
	"euglenóide",
	"evolutóide",
	"exigüidade",
	"exisulinde",
	"facilidade",
	"falangiode",
	"faranítide",
	"fasianóide",
	"faticidade",
	"feminidade",
	"feracidade",
	"fibrinóide",
	"fidelidade",
	"filarióide",
	"filosidade",
	"fisicidade",
	"flavonóide",
	"floriverde",
	"fluididade",
	"fogosidade",
	"fonalidade",
	"fortuidade",
	"friganóide",
	"friosidade",
	"froixidade",
	"frouxidade",
	"fumosidade",
	"funarióide",
	"furacidade",
	"futilidade",
	"futuridade",
	"galactóide",
	"galateóide",
	"gastrópode",
	"gelividade",
	"gemelidade",
	"genialóide",
	"germindade",
	"gibosidade",
	"girocécide",
	"girocóride",
	"glaucópide",
	"goianidade",
	"gomosidade",
	"gongilóide",
	"gonozoóide",
	"gorilhóide",
	"goticidade",
	"graminóide",
	"granitóide",
	"granulóide",
	"gratuidade",
	"grilácride",
	"grossidade",
	"gulosidade",
	"habilidade",
	"halisióide",
	"hamadríade",
	"hamamélide",
	"hapalópode",
	"hastiverde",
	"hecceidade",
	"helicidade",
	"helicópode",
	"hematópode",
	"hemifólide",
	"hemorróide",
	"herminíade",
	"hersilóide",
	"heterópode",
	"heteróside",
	"heticidade",
	"hexalépide",
	"hexaplóide",
	"hilaridade",
	"himenópode",
	"hipoplóide",
	"hipsilóide",
	"histeróide",
	"holotúride",
	"homalópode",
	"hombridade",
	"homeridade",
	"humosidade",
	"icterópode",
	"identidade",
	"idoneidade",
	"imbacaçade",
	"imensidade",
	"imiquimode",
	"indenidade",
	"indianóide",
	"ineqüípede",
	"infimidade",
	"inimistade",
	"iniqüidade",
	"inocuidade",
	"inquietude",
	"insabidade",
	"invejidade",
	"ionicidade",
	"iscnáspide",
	"isocolóide",
	"isofitóide",
	"jabutifede",
	"jaguarande",
	"jocosidade",
	"juricidade",
	"labilidade",
	"lamelípede",
	"lanosidade",
	"lantanóide",
	"latinidade",
	"lemodípode",
	"lepidópede",
	"lepidópode",
	"leucemóide",
	"leuconóide",
	"leuprolide",
	"leveduride",
	"leviandade",
	"libripende",
	"lidimidade",
	"ligocóride",
	"limosidade",
	"lingulóide",
	"liquenóide",
	"litocaride",
	"lofoséride",
	"logicidade",
	"longicaude",
	"mabelemade",
	"machundade",
	"macotidade",
	"maduridade",
	"magnetóide",
	"malacópode",
	"malinidade",
	"mansuetude",
	"margaróide",
	"maritafede",
	"maturidade",
	"maxilípede",
	"maxilípode",
	"maxilópode",
	"megalópode",
	"megatíride",
	"melantóide",
	"meleágride",
	"melicéride",
	"melosidade",
	"meniscóide",
	"mesmeidade",
	"metaldeíde",
	"meteoróide",
	"micelióide",
	"mimusópode",
	"minimidade",
	"minoridade",
	"mixoplóide",
	"modicidade",
	"molaridade",
	"molibdóide",
	"molossóide",
	"moluscóide",
	"mongolóide",
	"monoculode",
	"monolépide",
	"monoplóide",
	"mortandade",
	"motilidade",
	"mucosidade",
	"multicaude",
	"murilaonde",
	"narcisóide",
	"nautilóide",
	"nectozóide",
	"nefritóide",
	"negritóide",
	"nematópode",
	"nitritóide",
	"nivosidade",
	"nobilidade",
	"nodosidade",
	"notobáride",
	"notogapide",
	"nubilidade",
	"nucleóside",
	"nugacidade",
	"objetidade",
	"obtusidade",
	"oceanidade",
	"oceanítide",
	"octolépide",
	"octoploíde",
	"odiosidade",
	"odocnêmide",
	"odonéspide",
	"odontópode",
	"ofiofólide",
	"ofiolépide",
	"oleosidade",
	"olistópode",
	"olivinóide",
	"oncinópode",
	"onosséride",
	"opistópode",
	"opticidade",
	"oquizópode",
	"oreocálide",
	"orgonidade",
	"orizópside",
	"ornitópode",
	"otenzepade",
	"otostílide",
	"ovomucóide",
	"oxigráfide",
	"oxisoprede",
	"pacacidade",
	"palmelóide",
	"panolópode",
	"paracônide",
	"paradóxide",
	"paralépide",
	"parasetode",
	"parilidade",
	"paronirode",
	"parvoidade",
	"pecilópode",
	"pediópside",
	"pegmatóide",
	"pelecípode",
	"pelibrande",
	"pelosidade",
	"pelucidade",
	"penfigóide",
	"peptoniode",
	"peracáride",
	"pereiópode",
	"pergamóide",
	"pergelóide",
	"perseidade",
	"petalópode",
	"petauróide",
	"piboserode",
	"pidotimode",
	"pilosidade",
	"pinguípede",
	"pintagóide",
	"piroméride",
	"pirotônide",
	"plagiópode",
	"planetóide",
	"platanóide",
	"platinóide",
	"plebeidade",
	"pleurópode",
	"podicípede",
	"pogonópode",
	"polilépide",
	"poliplóide",
	"polizoóide",
	"porfiróide",
	"porquidade",
	"pouquidade",
	"presbítide",
	"probóscide",
	"protáspide",
	"protopoíde",
	"pseudópode",
	"pterigóide",
	"pulcritude",
	"quadrúpede",
	"quatimundé",
	"queratóide",
	"quietitude",
	"quitinóide",
	"rabditóide",
	"ramosidade",
	"rapacidade",
	"reumatóide",
	"rigoridade",
	"rinelépide",
	"romanidade",
	"rugosidade",
	"rusticóide",
	"sacaróside",
	"salacidade",
	"salinidade",
	"santalóide",
	"sarcopside",
	"saurópside",
	"seborreide",
	"secantóide",
	"sedosidade",
	"selenípede",
	"selenópode",
	"seminômade",
	"septicorde",
	"serenitude",
	"serosidade",
	"sifonópode",
	"similidade",
	"similitude",
	"sinalpende",
	"sintoxóide",
	"siringóide",
	"sistemóide",
	"sobriedade",
	"solicitude",
	"sorosidade",
	"subulípede",
	"subunidade",
	"sucosidade",
	"sulfuróide",
	"sutilidade",
	"talcitóide",
	"tasmanóide",
	"taticidade",
	"tatilidade",
	"teangélide",
	"tegaserode",
	"temeridade",
	"tempestade",
	"tenebróide",
	"terpenóide",
	"tetracorde",
	"tipicidade",
	"tirilazade",
	"tonicidade",
	"topicidade",
	"tortricode",
	"tosquedade",
	"tosquidade",
	"toxicidade",
	"tragulóide",
	"trapezóide",
	"traqueóide",
	"tremelóide",
	"tricúspide",
	"tridésmide",
	"trietéride",
	"triórquide",
	"triptéride",
	"trirráfide",
	"tumididade",
	"ubiqüidade",
	"undosidade",
	"unicúspide",
	"upanissade",
	"vaporidade",
	"variolóide",
	"velocípede",
	"velosidade",
	"venosidade",
	"vericidade",
	"vetusidade",
	"vilosidade",
	"vinosidade",
	"virgindade",
	"viscolóide",
	"xistrópode",
}
