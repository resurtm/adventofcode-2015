package problem5

import (
	"fmt"
	"strings"
)

type testCase struct {
	input  []string
	result int
}

var commonCase = []string{"rthkunfaakmwmush", "qxlnvjguikqcyfzt", "sleaoasjspnjctqt", "lactpmehuhmzwfjl", "bvggvrdgjcspkkyj", "nwaceixfiasuzyoz", "hsapdhrxlqoiumqw", "lsitcmhlehasgejo", "hksifrqlsiqkzyex", "dfwuxtexmnvjyxqc", "iawwfwylyrcbxwak", "mamtkmvvaeeifnve", "qiqtuihvsaeebjkd", "skerkykytazvbupg", "kgnxaylpgbdzedoo", "plzkdktirhmumcuf", "pexcckdvsrahvbop", "jpocepxixeqjpigq", "vnsvxizubavwrhtc", "lqveclebkwnajppk", "ikbzllevuwxscogb", "xvfmkozbxzfuezjt", "ukeazxczeejwoxli", "tvtnlwcmhuezwney", "hoamfvwwcarfuqro", "wkvnmvqllphnsbnf", "kiggbamoppmfhmlf", "ughbudqakuskbiik", "avccmveveqwhnjdx", "llhqxueawluwmygt", "mgkgxnkunzbvakiz", "fwjbwmfxhkzmwtsq", "kzmtudrtznhutukg", "gtvnosbfetqiftmf", "aoifrnnzufvhcwuy", "cldmefgeuwlbxpof", "xdqfinwotmffynqz", "pajfvqhtlbhmyxai", "jkacnevnrxpgxqal", "esxqayxzvortsqgz", "glfoarwvkzgybqlz", "xdjcnevwhdfsnmma", "jyjktscromovdchb", "pvguwmhdvfxvapmz", "iheglsjvxmkzgdbu", "lwjioxdbyhqnwekv", "zcoguugygkwizryj", "ogvnripxxfeqpxdh", "hkvajhsbfnzsygbm", "cnjqeykecopwabpq", "wojjtbcjinoiuhsj", "kpwpvgxbyzczdzjq", "wrvhylisemlewgzk", "uiezkmnhilfzahtm", "mucteynnuxpxzmvt", "zaiwbgxefusfhmst", "apptbogpxivjwink", "qryboarjtwjhjgjb", "irehxupgyseaahzd", "fobstqxguyubggoh", "ysriumfghtxtfxwe", "auchdmasvfeliptw", "mztuhefcrnknyrdl", "tyjmkhihbwabjtaa", "yquzkdtgsljkaebw", "almvdvofjtkyzbmd", "emqftiuqqpdwwbrv", "hrrhmqfpepvbawvw", "atrkgykycvgxbpyb", "dhthetnealksbdan", "zzqafhgicubptiyo", "qdtaieaziwhbttnw", "kyskgapdgqrtrefw", "edwzlpqztpydmdlr", "awszjnlmvlyqsuvl", "kcrtmtshtsgixvcp", "jtaskgkijivbbkri", "mmggfwapsetemiuj", "itagrrnjbnmhgppd", "uqmbezechbrpbnqq", "nnyimvtascflpzsa", "knqeimypkdttyudj", "vgoiyvtvegwyxjjd", "qubzdxsbecktzrho", "zehojtvktsbbxijb", "xepmjrekwcgoxyoh", "bnptxnocbpbqbyeq", "sfvynsywscbnymos", "dsltfbpcmffbluba", "kncrlzlmkikylppa", "siwudrvmildgaozv", "jhhefbvbvneqzvtc", "lqjgztxitbuccqbp", "himmwlbhjqednltt", "vwognchyertnnfil", "eejakhapkbodrntf", "qxuijkkhhlskgrba", "aankpfxxicfpllog", "vuxykvljyqexfhrn", "epgygflbxlbwybzq", "zuxmwvetmvcszayc", "xttwhfqmemgtjnkf", "hftwldmivyfunfvl", "bejlyxfamzliilrj", "zkehazcxyyvtrxti", "dsgafehmcfpycvgz", "igremmqdojqdvwmb", "swnjzvmhcslvkmiw", "fchzbfbmtqtxmaef", "xwjmyyrlznxrcytq", "brwcwzpcvbwdrthl", "fvrlridacsiojdmb", "mhsturxdlmtxozvy", "usxvqyrwywdyvjvz", "gwazuslvmarfpnzm", "rgkbudaqsnolbcqo", "dpxvlbtavdhdedkj", "nnqmjzejhodyfgyd", "ozoazxkfhujgtzvy", "psdgvhzdiwnuaxpl", "tznkilxpogbzgijz", "wnpytcseirtborhh", "lhauurlfsmagfges", "oqfbzixnlywkzwwy", "yoehapoyjpakziom", "vtjftdcsfdzbmtrn", "zcshfnodiwixcwqj", "wapbxpaxgjvtntkm", "qfyypkyvblrtaenh", "bsxhbxkovgukhcza", "kitdmvpiwzdonoyy", "slkbhxmehzavbdsf", "dovzjouqkzkcmbkl", "qpbigdcqkfnfkxvq", "eaiaquhnesvtcdsv", "mhbezlhqojdsuryj", "dqprkkzxlghkoccx", "xqepmorryeivhrhm", "frwmrjpezwmjflvf", "gjpfgwghodfslwlf", "fzyvajisdjbhfthq", "pvzxkxdscdbilrdb", "mtaxmqcnagmplvnm", "rlyafujuuydrqwnc", "gvqvrcxwyohufehq", "lmrkircgfrfusmfd", "ovlpnkxcpimyaspb", "xhyjremmqhdqywju", "pxfczlhpzbypfarm", "utjhprzhtggausyp", "utzkkzlnyskjtlqh", "cecbcnxpazvkedic", "xwvoaggihrbhmijq", "krredhmtwlfmyagw", "lwfhxgbknhwudkzw", "vyczyvuxzmhxmdmn", "swcoaosyieqekwxx", "waohmlfdftjphpqw", "gaclbbfqtiqasijg", "ybcyaxhluxmiiagp", "xgtxadsytgaznndw", "wzqhtjqpaihyxksm", "fdwltsowtcsmsyhm", "rpoelfbsararhfja", "tswgdacgnlhzwcvz", "xjgbhdlxllgeigor", "ksgthvrewhesuvke", "whgooqirdjwsfhgi", "toztqrxzavxmjewp", "hbkayxxahipxnrtl", "lazimkmdnhrtflcu", "ndoudnupbotwqgmr", "niwuwyhnudxmnnlk", "hlmihzlrpnrtwekr", "wzkttdudlgbvhqnc", "rfyzzgytifkqlxjx", "skddrtwxcyvhmjtb", "mljspkvjxbuyhari", "xwkhozaoancnwaud", "nookruxkdffeymdz", "oiqfvpxmcplyfgoa", "qoxggshmrjlzarex", "lsroezewzkrwdchx", "nkoonmvdydgzspcl", "lygxeqztdqklabov", "jempjyzupwboieye", "hpdaqkhjiddzybly", "cvcizjlnzdjfjlbh", "vaaddsbkcgdjhbkj", "pjxmtxoyrkmpnenf", "ujqdvyqnkbusxlps", "miyvzkzqploqaceb", "gapcsbkulicvlnmo", "xqpcyriqhjhaeqlj", "ipumdjwlldzqhmgh", "swdstecnzttmehxe", "ucmqordmzgioclle", "aywgqhmqlrzcxmqx", "ptkgyitqanvjocjn", "wcesxtmzbzqedgfl", "rnetcouciqdesloe", "chpnkwfdjikqxwms", "onpyrjowcuzdtzfg", "tydnqwaqwkskcycz", "dhamguhmkjzzeduy", "oecllwyrlvsyeeuf", "gsukajpoewxhqzft", "sgdnffdixtxidkih", "pqqzjxzydcvwwkmw", "wnjltltufkgnrtgm", "hylaicyfrqwolnaq", "ovfnugjjwyfjunkm", "xknyzsebmqodvhcl", "uwfmrjzjvvzoaraw", "zaldjvlcnqbessds", "zphvjuctrsksouvz", "ceqbneqjwyshgyge", "wmelhaoylbyxcson", "nghuescieaujhgkj", "dhjmflwwnskrdpph", "exvanqpoofjgiubf", "aidkmnongrzjhsvn", "mdbtkyjzpthewycc", "izctbwnzorqwcqwz", "hrvludvulaopcbrv", "mrsjyjmjmbxyqbnz", "sjdqrffsybmijezd", "geozfiuqmentvlci", "duzieldieeomrmcg", "ehkbsecgugsulotm", "cymnfvxkxeatztuq", "bacrjsgrnbtmtmdl", "kbarcowlijtzvhfb", "uwietqeuupewbjav", "ypenynjeuhpshdxw", "fwwqvpgzquczqgso", "wjegagwkzhmxqmdi", "vocvrudgxdljwhcz", "nnytqwspstuwiqep", "axapfrlcanzgkpjs", "lklrjiszochmmepj", "gxadfpwiovjzsnpi", "qidsjxzgwoqdrfie", "wgszciclvsdxxoej", "kwewlmzxruoojlaq", "ywhahockhioribnz", "ucbqdveieawzucef", "mdyyzmfoaxmzddfv", "hsxnabxyqfzceijv", "vivruyvbrtaqeebr", "jxfeweptjtgvmcjc", "mmypqxmpurhculwd", "mpiaphksvctnryli", "xqzqnuxmuzylkkun", "fndmtefjxxcygtji", "dnorqlldvzqprird", "nutokyajmjpwjaqu", "vlupfperqyqkjcaj", "dgihjeokrphkpdnk", "nvbdyrlheqzixuku", "mhrkntnxvsmvrpka", "kvhkyanlhhymwljf", "fhipumtegqfgeqqw", "vpfjgveycdefuabu", "kzincljffncylcsf", "tsezxymwmjtyegqw", "wxhcdrqedkdcwxli", "ueihvxviirnooomi", "kfelyctfvwyovlyh", "horzapuapgtvzizz", "iiqkdpmfvhwwzmtj", "rsaclclupiicstff", "quwkkhrafypkaoum", "gyrgkgmwqfkeudfe", "noydhbqacwptyfmy", "efwwuipzgtkwffhf", "suyojcitomdxsduh", "lbcxnsykojkufkml", "zpglsvoutvzkgdep", "usgrufyvgsbsmbpr", "katrrwuhwvunjqor", "btngwrpcxoyfbgbc", "bxjscjdiowjrkpns", "nwxvnfrnlkgqxvhf", "ikhyqkvljucgdlag", "xibnxsjopmxvflkl", "mzplumcfivqcjqnz", "jqflcxoxzlbwlxry", "fcscvmfepdxrshxe", "wlpffwunffklzbuc", "emvrlqajjgwzfmle", "rhaheurtzrfoqkyq", "ifuuhpxmadaysfsx", "ncyfvleyzqntpcoo", "zeogmyaqccmtvokd", "jqppbzebppdnpurn", "xixarswxsiwjzgni", "ezruwzajsoombphs", "hmiqfeizyprielxf", "jnaoxljnftymsfey", "extgzrxzovlsixnf", "yhyfmovvlrwoezsv", "ffnybaolppuzpjym", "pqowimdiusccaagn", "jgceiosiihpjsmnu", "hkoexeaopebktngx", "njhzuvsygymejqav", "yjkgcclgtvushcfk", "gmbjxhnkkxlihups", "pdlwysadiebsidjz", "omrwmgzulfoaqros", "ofvvgdezwvcffdcy", "otytpuklhxcpxhgd", "eyfaosxdauumvlux", "mvdthjfstrlqlyuo", "mdgdchgnlxaxspdm", "bakjezmhbwqxzevd", "msakswaphdwaodhg", "vjcqscgdbnsxdllh", "jjywaovewbuzreoj", "nqvplhwacylifvwk", "lpwmpixbxysmsign", "flcvbpxrchcpbgcb", "qjpkeuenenwawlok", "bnqkflfmdmntctya", "fzsgzpoqixvpsneq", "icwfdisutoilejld", "relchofohnkwbumi", "aljalgdaqwhzhfwr", "cahkvnwnbwhodpqs", "dnrzeunxiattlvdm", "nsmkhlrpwlunppjs", "mqqsexlwfqnogwub", "tfavelkqrtndpait", "ooguafrnmprfxcnz", "ntynkiordzxtwrqa", "rkkyzlxekqqlkvym", "ofxcivdnwcmgfnme", "ywotqwbrqxlrnobh", "nrbbiypwhrqihvev", "flqsjixxtydheufs", "lcfrfzypstrqctja", "hyzbuzawuzjrynny", "exfbywcnstebnvmq", "vydzwnbmcihvqrnj", "qmwqaaylinzrdmiw", "lpxpztpvfggspeun", "lhxmqqbracsuyrfm", "zgkwsrabaseidbrw", "yjlmbhbqsqgszsun", "mqfzqtbxtuteabtd", "izomzdmcqmfrevwd", "iqijrlqurdwrkoln", "fxhqzpgoxxjkkhql", "oulwontmgrjeopnk", "edaigfydjexvzzvj", "vjhybiklxpxjqpwc", "ypxfbfnpbmqmwtte", "xzvcsgasztrxdzud", "rpulqmobptfarboo", "palacmdijxzzykrf", "jmllwukplufohiby", "dnswayomusiekfmy", "sxbrjqtqgzzwhcfo", "lylvndsgbnbqiejm", "jaxxhoulxnxnaenr", "nblissutfazbcpwn", "zmlsjszzldvbiacr", "kewojtlchfkclqwk", "eqvfjasddggvfame", "yibzqlvxtraxpdon", "dgnbxsbmdrtyvaac", "uoxrcxfimhgtxqhy", "xfdxalrwcwudlviq", "xmtbdklqptoswpwl", "zezyopzdztdjerfl", "xuzluhjsqvhytgbc", "qdjtmeckispmgzki", "phakupesplzmmmvc", "gpuoqfffumzszybn", "bhywxqkrrlwuebbw", "ibvwgoyvelzenkzl", "ncohvvbmiekbaksa", "fzuvqzvxvdbeirrp", "lshtzniokucwojjd", "punrduvlnrulkium", "gnfpikidnfobrrme", "vxkvweekmnvkzgyl", "rhydssudkcjlqgxn", "cjtqvlaahohcgumo", "jwzmfyinsfwecgcb", "blpeseqhlzfilpuf", "jvtpjkyokzcvagon", "qjomincbcobjczpe", "ugsyzkzgdhxtmsfz", "hleaqgwzqjwajcra", "coumfghptpnxvvov", "hqpnbupnzwpdvgqd", "cpouyodqxgviasem", "lljvxeyozckifhfd", "huqtnvutdyfgwtwa", "yenlveuynmlmmymu", "ojdyufkomxiwjmbf", "spjzgvcwvzgffjkk", "vxykmjhyvmhyssbp", "tazdeqggfcjfvwwn", "uumwcngwcytvpufx", "avovuzkrevloneop", "owczrtbnrvjfemkt", "hzpugcanaxyvaokj", "iishlodnxvjtgzyn", "qosdonclrnxirham", "eonqlnwevahydddg", "ryqmnuikftlxuoqy", "whqepbcwabzbthha", "vekisvnwhgpyemxr", "lrwxzoamnvpnlhap", "ywepvqthnorfswjv", "evqwvsoazmwyypjy", "bgwoojddubppmjxf", "jypkfrthzgtyeddi", "tynabbhfjzkrqsju", "adxstbfqheuqbcuk", "gqwqiocdyqoiblrx", "ybuddlyuskdlegxv", "luwynbsmpgyeqsbr", "ltyqgqoyljibqndo", "jaedpajzphfybajh", "epglnrxofptsqvmy", "zjdpxkngfkstxbxh", "ekegphcwanoickfu", "cqvhuucvejqirvfs", "uqudnnqumsqcgefo", "qnzunermlnpcfflo", "ovyxaniqaawzfuxx", "djekxcezjowdhopq", "bwtwbmdehrhpjnlk", "nilsnlacerweikfa", "hyrigsrmsrzcyaus", "gvmdmgddduylmxic", "ewzovdblhmjgjwsk", "ojjfsknlonzguzlq", "yjgfruvpjvlvrvvq", "cyoryodwyhzwprbv", "crsjclrurcquqgut", "sjhfhobwtojxcmem", "ibxfjudilmdeksea", "uqbhdbjoeupyhbcz", "uqbxigzxuxgmjgnw", "jashafmtzrhswirg", "dexiolovaucyooka", "czjbwwnlwcoqnoiu", "ojigosazigfhttjc", "zfiqtgrqbmftknzn", "dlzbmvmolssbqlzl", "sgmchcurrutdtsmw", "scdwjqsdohcdrwry", "cgtdvecqwplpprxn", "iiplenflfczaktwi", "wmgnwfxfcjhyeiqg", "giihshowtcatecvl", "nqhzfincclumvkaz", "kxstpzgdfvepionc", "agbhxcijxjxerxyi", "hmgfqevgdyvisyvs", "tthakmvpowpvhtao", "ottalcghygpaafbo", "aplvozayycremgqg", "dbjxlnaouxqtdpfz", "peeyallzjsdvpalc", "ndtdjyboixuyhfox", "llabnbcobexfoldn", "cweuvfnfyumbjvxr", "ewkhhepaosalnvkk", "pivyiwsiqpwhagyx", "auzsnwdcerfttawt", "grbfrekupciuzkrt", "byfwzadtzrbndluf", "lluypxjeljzquptk", "pskwsnhqanemtfou", "sxvrtqqjdjkfhhrm", "ulsmqgmshvijyeqh", "qigofesfhekoftkf", "zhatniakqtqcxyqa", "uuczvylgnxkenqee", "mlitvtuxknihmisc", "srrtrxdvcokpyfmz", "osispuucklxcfkeb", "vqhazlaulmnpipql", "umkiueljberqhdig", "knvpbkbvgoqzwprp", "nbsocqikhuvsbloj", "wjnpepjkzkednqbm", "agbhmytsofuyqcor", "gvogzhkkpxyfecko", "ardafguxifeipxcn", "yiajcskbgykyzzkw", "sejunbydztyibnpq", "dqrgfggwcnxeiygy", "xnqqwilzfbhcweel", "jjtifhlvmyfxajqi", "gwszrpgpmbpiwhek", "kydzftzgcidiohfd", "efprvslgkhboujic", "kecjdfwqimkzuynx", "rildnxnexlvrvxts", "dlnhjbqjrzpfgjlk", "qluoxmzyhkbyvhub", "crydevvrjfmsypbi", "dosaftwumofnjvix", "pwsqxrfwigeffvef", "nzyfmnpwqyygjvfx", "iccbckrkxlwjsjat", "bmputypderxzrwab", "bhuakynbwnlreixb", "qmrzfyqjiwaawvvk", "juvtixbkwyludftn", "zapmjxmuvhuqlfol", "paiwrqjhpjavuivm", "tsepfbiqhhkbyriz", "jpprewufiogxoygk", "mmapyxbsugcsngef", "pduhmgnepnpsshnh", "aetndoqjvqyjrwut", "fnfvlorhwpkkemhz", "gedfidpwvoeazztl", "beclvhospgtowaue", "wsclsvthxustmczm", "tjbxhnpniuikijhe", "rhetyhvfcemponeg", "mavonujurprbeexi", "argbrpomztrdyasa", "bzvtffbtygjxmkvh", "maqyqkhsqgzfzvve", "seeirbiynilkhfcr", "wxmanwnozfrlxhwr", "dieulypsobhuvswb", "nxevassztkpnvxtb", "jclxuynjsrezvlcy", "xlolzyvgmwjsbmyf", "tguzoeybelluxwxc", "fkchoysvdoaasykz", "cyynwbfcqpqapldf", "rhifmzpddjykktuy", "ndvufsyusbxcsotm", "txutnzvdsorrixgg", "qjoczhukbliojneu", "ufhwujotncovjjsz", "kclsgsdwcrxsycbr", "yscwmlrdaueniiic", "nxhivrovpkgsmugb", "fdxqfyvwwvgeuqkv", "femtamfylysohmpr", "amsyzslvyxsoribh", "nhmqxncwsonhgbcz", "uomqsvcbpthlmcue", "kxtfapcqrnjkkslj", "xtieihonlfubeync", "adpcjqxgydulchgj", "cjynnzsmmujsxxpd", "neeapmzweidordog", "szoivgqyqwnyjsnk", "uwgrtzaqezgphdcu", "ptpgttqxocjwxohi", "fhltebsizfwzpgpf", "emmsazsidspkhgnh", "dxcprkbcjeqxqzgn", "tpxzqwxbzwigdtlt", "afsmksnmzustfqyt", "xyehnftstacyfpit", "vcrfqumhjcmnurlw", "rrznpjzcjgnugoch", "gbxnzkwsjmepvgzk", "jwobshgwerborffm", "zmuvfkhohoznmifs", "buyuwgynbtujtura", "bevncenmpxfyzwtf", "hqqtcrhzfsrcutjh", "kbpzshllpiowepgc", "alspewedcukgtvso", "xvsvzzdcgjuvutrw", "pmwulqraatlbuski", "abuzsiinbueowpqn", "oedruzahyfuchijk", "avhcuhqqjuqkesoq", "azqgplkzsawkvnhb", "rjyoydogkzohhcvx", "aezxwucqvqxuqotb", "kxobnsjvzvenyhbu", "nnjoiilshoavzwly", "aijttlxjrqwaewgk", "cvsaujkqfoixarsw", "zngtoacpxcsplgal", "qhkxliqtokvepcdv", "aixihrtdmxkfvcqw", "owbgdgdymxhhnoum", "tajsagmruwzuakkd", "ckrfduwmsodeuebj", "alfdhuijuwyufnne", "xpchlkijwuftgmnm", "rwcrvgphistiihlg", "xdaksnorrnkihreq", "akeschycpnyyuiug", "rgputhzsvngfuovz", "lerknhznuxzdhvre", "mqiqmyladulbkzve", "csnmupielbbpyops", "kwgrwgmhfzjbwxxz", "npwtvbslvlxvtjsd", "zxleuskblzjfmxgf", "hexvporkmherrtrn", "rhtdhcagicfndmbm", "qhnzyuswqwoobuzz", "dpvanjuofrbueoza", "kjcqujmnhkjdmrrf", "gholddsspmxtpybg", "jihlvyqdyzkshfsi", "zuviqmuqqfmtneur", "kzexjowatvkohrtx", "wgijnfhibsiruvnl", "zevkrkmhsxmicijb", "khxrcteqourjvoxa", "ylpxlkcnenbxxtta", "zrfsvctbojjkpvtw", "nlzbudxibnmcrxbt", "cqnscphbicqmyrex", "ywvdohheukipshcw", "riwatbvjqstubssf", "idlztqqaxzjiyllu", "sdpdgzemlqtizgxn", "rjtbovqlgcgojyjx", "fnfrfwujmjwdrbdr", "osnppzzmrpxmdhtj", "ljhwngclvydkwyoe", "chwqkrkzrvjwarat", "jmydkwpibkvmqlgs", "zvhfmbxnlxtujpcz", "jsnhsphowlqupqwj", "fzhkkbpasthopdev", "jerntjdsspdstyhf", "gctwmaywbyrzwdxz", "xemeaiuzlctijykr", "xulrqevtbhplmgxc", "yfejfizzsycecqpu", "gboxrvvxyzcowtzm", "lpvhcxtchwvpgaxp", "wdiwucbdyxwnjdqf", "qgwoqazzjlvnjrwj", "prtlnkakjfqcjngn", "fagvxsvjpuvqxniz", "xacmxveueaakfbsm", "ginvtonnfbnugkpz", "qpvggsppewfzvwin", "reoqnlzruyyfraxa", "kolwtqhifjbbuzor", "vrkcywvdhdprztww", "ngdvyfmvjqhbzbxt", "rooxeoilqzqjunmp", "efxmdprtogtxgyqs", "qrhjuqndgurcmwgu", "ouitjprueefafzpl", "kirdwcksqrbwbchp", "fpumsmogojuywezo", "lgjrgykywugzjees", "xigioqcpjabpbdas", "ewkhuprpqzikmeop", "fgrgxsqeducigxvr", "bclkursnqkzmjihl", "jozidniwvnqhvsbc", "oghcilcyozrmmpta", "xbgmaungzcpasapi", "iqowypfiayzbcvhv", "opdehgwdgkocrgkf", "zfzvdjeinlegcjba", "vhakxvlcayuzukap", "xyradgyiebpevnwe", "eamhtflgedwyshkn", "igteqdgchjeulfth", "kwsfkigxzpbgdxod", "vapnpsbdboiewpzp", "wbuqhjsngxpqshen", "vxxilouxuytitwgm", "cpnwlkwnkeanqnet", "wdmbtqvvlowftvgb", "wjtmcecpyqzwpbqg", "jnxmoxdhvsphcdeg", "wabxfxpotoywwodn", "mwbsoxzlqpqobvvh", "coktshbyzjkxnwlt", "rzhnggpslwzvyqrp", "dgzuqbzarbutlkfx", "wunajaiiwgijfvjh", "uotdbcgmsvbsfqlb", "kxdtlgmqbccjqldb", "ngmjzjwvwbegehfr", "cvpsabqfpyygwncs", "wqluvqlhdhskgmzj", "rbveperybfntcfxs", "fbmoypqdyyvqyknz", "zxpgzwnvmuvkbgov", "yexcyzhyrpluxfbj", "ltqaihhstpzgyiou", "munhsdsfkjebdicd", "plecvjctydfbanep", "kjrxnnlqrpcieuwx", "zbcdtcqakhobuscf", "kgovoohchranhmsh", "llxufffkyvuxcmfx", "tgaswqyzqopfvxtw", "kojcqjkdpzvbtjtv", "xggdlkmkrsygzcfk", "vvitpsnjtdqwyzhh", "gcqjuwytlhxsecci", "vbsghygcsokphnrg", "vejqximdopiztjjm", "hudqtwmwkviiuslp", "vwswfvpcwwpxlyry", "gxmfiehdxptweweq", "qjmekjdcedfasopf", "pqyxdxtryfnihphf", "felnavctjjojdlgp", "hbimufguekgdxdac", "dhxhtnqgfczywxlr", "pssottpdjxkejjrh", "edieanguabapxyig", "sciinanyqblrbzbb", "irxpsorkpcpahiqi", "qsxecaykkmtfisei", "ivfwlvxlbnrzixff", "hqxzzfulfxpmivcw", "vvbpaepmhmvqykdg", "cetgicjasozykgje", "wuetifzdarhwmhji", "gaozwhpoickokgby", "eldnodziomvdfbuv", "favpaqktqaqgixtv", "twbcobsayaecyxvu", "lzyzjihydpfjgqev", "wnurwckqgufskuoh", "fxogtycnnmcbgvqz", "aetositiahrhzidz", "dyklsmlyvgcmtswr", "ykaxtdkjqevtttbx", "kfmnceyxyhiczzjm", "nnizopcndipffpko", "yjmznhzyfinpmvkb", "sljegcvvbnjhhwdd", "zmkeadxlwhfahpwg", "rwvcogvegcohcrmx", "aguqwrfymwbpscau", "vlusytjagzvsnbwe", "smvzhburcgvqtklh", "rfuprvjkhazrcxpv", "megqlnoqmymcrclc", "gvldhkewtmlwqvqv", "awynhvtyziemnjoa", "voprnvtnzspfvpeh", "dhlguqwmunbbekih", "goayirdhnjrfuiqi", "eoghydfykxdslohz", "chpippjykogxpbxq", "hqbycjweqczwjwgf", "pvefsrvwumrlvhmt", "eghwdovaynmctktk", "crwkxoucibumzawc", "bzbtahvhkdigvvtj", "bnbptgihhfubxhho", "ddqmbwyfmfnjjaro", "gvtswqyzazihctif", "vmqctjpgadxztqqb", "dgnndowtpeooaqqf", "sxdvctfdtalufxty", "ylgeexosibsmmckw", "sxplpyskbpqnojvw", "coarhxtsvrontyeg", "fyoaurggjupvzvlv", "jlyrkqsiwuggvjem", "uwbsjoxonreuucyi", "gihuqvwxovbgokes", "dxzaaxupbcgnxcwf", "gidrgmvyrlqqslve", "csflmlvqmonoywpx", "jkxkpixlythlacnk", "ejkarcdkdslldugv", "dbzmsusevohhjkmr", "cbrqzualjpdtworc", "kpgidqlmcbpfmmwu", "zwghjuofexfowqam", "ncdlxmcrsmsocetz", "kfprzqacefifjkbd", "swwzivrxulkhvldc", "wgqejhigbjwunscp", "rsstnwcyybfauqxu", "qhngfxyhdqopyfgk", "zrndpyyejsmqsiaj", "xxknxwpvafxiwwjc", "mmaahwgoiwbxloem", "tabacndyodmpuovp", "yriwomauudscvdce", "duvyscvfidmtcugl", "mgipxnqlfpjdilge", "imeeqcdetjuhfjnw", "dvkutrdofpulqkyh", "jefvtlktxegpmbya", "iyzudqgpvlzjfydh", "giohapxnpaqayryd", "qheqdprmnqlpztls", "rdxhijmzegxkotoq", "hdnmaspumdwnrcdz", "wafpbgehbuzdgsnc", "tbtrfztsferdmhsy", "vusndcyjngtkrtmk", "ilqblestzxebcifh", "urfgjbjgzlrfsdlv", "aptcdvpsqwleqttn", "bigczjvzokvfofiw", "zjnjeufonyqgkbpx", "trcdebioegfqrrdi", "jrdvdriujlmbqewt", "jqrcmuxpwurdhaue", "yjlermsgruublkly", "zwarvgszuqeesuwq", "xthhhqzwvqiyctvs", "mzwwaxnbdxhajyyv", "nclsozlqrjvqifyi", "gcnyqmhezcqvksqw", "deuakiskeuwdfxwp", "tclkbhqqcydlgrrl", "qbpndlfjayowkcrx", "apjhkutpoiegnxfx", "oaupiimsplsvcsie", "sdmxrufyhztxzgmt", "ukfoinnlbqrgzdeh", "azosvwtcipqzckns", "mydyeqsimocdikzn", "itfmfjrclmglcrkc", "swknpgysfscdrnop", "shyyuvvldmqheuiv", "tljrjohwhhekyhle", "dayinwzuvzimvzjw", "qgylixuuervyylur", "klqqaiemurawmaaz", "hdmzgtxxjabplxvf", "xiivzelzdjjtkhnj", "ktgplkzblgxwrnvo", "gvbpyofzodnknytd", "lqhlmnmhakqeffqw", "ltzdbngrcxwuxecy", "obxnfjeebvovjcjz", "zexpwallpocrxpvp", "tjpkkmcqbbkxaiak", "qiedfixxgvciblih", "qcxkhghosuslbyih", "gnsfidwhzaxjufgm", "xrghwgvyjakkzidw", "tftftwedtecglavz", "wquqczzkzqrlfngr", "twibtkijpvzbsfro", "bmplypdsvzuhrjxp", "zanrfmestvqpwbuh", "zonrhfqowyimcukm", "kpvajjfmqpbhrjma", "kujzluicngigjbtp", "iusguantsrwxdjal", "kwxeuylcnszswahw", "visdhnkobxnemldu", "rogeadmmaicwtabl", "pxqycifbgevqudvs", "osaiozyvlyddylqr", "vffjxrolrpuxcatx", "jbmsetccdrywssjd", "qgxyhjfpbfifmvgc", "npejgalglldxjdhs", "mbbtqgmttastrlck", "whapaqwdtpkropek", "dulbdboxazfyjgkg", "xaymnudlozbykgow", "lebvqmxeaymkkfoy", "bmicnfuubkregouj", "dieatyxxxlvhneoj", "yglaapcsnsbuvrva", "bbpjaslqpzqcwkpk", "xehuznbayagrbhnd", "ikqmeovaurmqfuvr", "ylyokwuzxltvxmgv", "hqtfinrkllhqtoiz", "pjmhtigznoaejifx", "fqdbmowkjtmvvrmx", "uvqtqfoulvzozfxv", "rpajajukuxtchrjd", "sznucejifktvxdre", "ufvibsmoushmjbne", "xirdqoshngthfvax", "iafpkddchsgdqmzl", "vmualmlduipvykzh", "fnmuahmblwyceejb", "ilsaapnswfoymiov", "lenvylifraahaclv", "cukqxlipuyxedqfh", "zgwecslpniqvtvuz", "cdcdfpsxuyrhsmag", "dszjinhantnxgqra", "ioimwotsgnjeacgt", "dqcymnvjystbynhp", "yibaudyfefbfgunx", "cabslcvunjavqkbf", "goymzvmgkvlsmugf", "zxteiitpthzskjjx", "agnxcnaqhjhlurzs", "cvmgyxhhnykuxbmb", "cgqmjexydmvgwxpp", "sygjajofieojiuna", "clpvxbrbjvqfbzvu", "cbntswqynsdqnhyv", "bztpbtwbefiotkfa", "pnxccbgajvhyeybu", "asyzrvgzumtuissa", "facjyblvcqqginxa", "rvwnucnbsvberxuv", "ghrbeykzrxclasie", "ekujtselepgjtaql", "krtrzsmduhsifyiw", "ticjswvsnyrwhpnt", "clmjhsftkfjzwyke", "lbxlcixxcztddlam", "xhfeekmxgbloguri", "azxqwlucwhahtvep", "kitdjrwmockhksow", "keznwwcusgbtvfrs", "ljvzxoywcofgwajj", "vebjnhnkcfzbhrcw", "eqfcxkavstxcuels", "ldattkyawjrvcido", "bsqqeilshcwtqyil", "foqqsxahfiozcqrw", "liswfmuhzfbyzjhf", "sulbdcyzmolapfbs", "zuggzkelwxjpsgxb", "betioxrgtnhpivcw", "xmtbixstdipibhgs", "ttvurgqmulryyaji", "viobnljznzppfmxw", "qlzabfopydtxrlet", "tusvydegfxhaxolk", "thoufvvfjferxhwp", "cfyyzppfarjiilbs", "jwmhxtgafkkgseqs", "pqwuuaxbeklodwpt", "vndyveahdiwgkjyx", "ssrjgasfhdouwyoh", "thbavfcisgvvyekf", "yjdvxmubvqadgypa", "tlbmcxaelkouhsvu", "bonohfnlboxiezzr", "rktlxcbkhewyvcjl", "rsmoutcbcssodvsc", "qszdratuxcrhsvoh", "eypyfahpuzqwzwhi", "yhkrleqmqlmwdnio", "vpnvxusvmngsobmq", "hkzyhopvxrsimzys", "dblriiwnrvnhxykl", "xkriqxkrprjwpncs", "rcymltrbszhyhqti", "mzbvneplsnpiztzn", "vkqtnptgbqefvfoc", "nwdtfiaozkcjtlax", "crximadpvdaccrsm", "lrbajafxwwnxvbei", "rbexzesrytpwwmjf", "stxwjarildpnzfpg", "btamaihdivrhhrrv", "acqbucebpaulpotl", "dkjhzghxxtxgdpvm", "rsbzwsnvlpqzyjir", "mizypbwvpgqoiams", "nvrslorjpqaasudn", "wvexcpzmconqkbvk", "rfwfumhjwzrvdzam", "eaghdaqorkhdsmth", "gtuntmpqaivosewh", "nzlsmdgjrigghrmy", "dhuvxwobpzbuwjgk", "kkcuvbezftvkhebf", "aeediumxyljbuyqu", "rfkpqeekjezejtjc", "wkzasuyckmgwddwy", "eixpkpdhsjmynxhi", "elrlnndorggmmhmx", "ayxwhkxahljoxggy", "mtzvvwmwexkberaw", "evpktriyydxvdhpx", "otznecuqsfagruls", "vrdykpyebzyblnut", "cnriedolerlhbqjy", "uajaprnrrkvggqgx", "xdlxuguloojvskjq", "mfifrjamczjncuym", "otmgvsykuuxrluky", "oiuroieurpyejuvm"}

var testCasesPartOne = []testCase{
	{input: []string{"ugknbfddgicrmopn"}, result: 1},
	{input: []string{"aaa"}, result: 1},
	{input: []string{"jchzalrnumimnmhp"}, result: 0},
	{input: []string{"haegwjzuvuyypxyu"}, result: 0},
	{input: []string{"dvszwmarrgswjxmb"}, result: 0},
	{input: []string{"ugknbfddgicrmopn", "aaa", "jchzalrnumimnmhp", "haegwjzuvuyypxyu", "dvszwmarrgswjxmb"}, result: 2},
	{input: commonCase, result: -1},
}

var testCasesPartTwo = []testCase{
	{input: []string{"qjhvhtzxzqqjkmpb"}, result: 1},
	{input: []string{"xxyxx"}, result: 1},
	{input: []string{"uurcxstgmygtbstg"}, result: 0},
	{input: []string{"ieodomkazucvgmuy"}, result: 0},
	{input: []string{"aaa"}, result: 0},
	{input: []string{"aaaa"}, result: 1},
	{input: []string{"axaa"}, result: 0},
	{input: []string{"abcdefeghiaaa"}, result: 0},
	{input: []string{"abcdefeghiaaaa"}, result: 1},
	{input: []string{"qjhvhtzxzqqjkmpb", "xxyxx", "uurcxstgmygtbstg", "ieodomkazucvgmuy", "aaa", "aaaa"}, result: 3},
	{input: commonCase, result: -1},
}

func RunPartOne() {
	fmt.Printf(">>> PART ONE <<<\n\n")
	for k, v := range testCasesPartOne {
		fmt.Printf("test case : %d\n", k)
		if len(v.input) > 250 {
			fmt.Printf("input     : [TOO LONG VALUE]\n")
		} else {
			fmt.Printf("input     : %#v\n", v.input)
		}
		fmt.Printf("expected  : %d\n", v.result)
		res := v.solvePartOne()
		fmt.Printf("result    : %d\n\n", res)
		//if v.result != -1 && v.result != res {
		//	panic("INVALID CALCS!!!")
		//}
	}
}

func RunPartTwo() {
	fmt.Printf(">>> PART TWO <<<\n\n")
	for k, v := range testCasesPartTwo {
		fmt.Printf("test case : %d\n", k)
		if len(v.input) > 250 {
			fmt.Printf("input     : [TOO LONG VALUE]\n")
		} else {
			fmt.Printf("input     : %#v\n", v.input)
		}
		fmt.Printf("expected  : %d\n", v.result)
		res := v.solvePartTwo()
		fmt.Printf("result    : %d\n\n", res)
		//if v.result != -1 && v.result != res {
		//	panic("INVALID CALCS!!!")
		//}
	}
}

func (tc *testCase) solvePartOne() int {
	goodWords := 0
	for _, word := range tc.input {
		if hasDisallowedThings(word) {
			continue
		}
		if has3Vowels(word) && hasTwiceLetters(word) {
			goodWords += 1
		}
	}
	return goodWords
}

func (tc *testCase) solvePartTwo() int {
	goodWords := 0
	for _, word := range tc.input {
		//if !hasTwiceLetters(word) {
		//	continue
		//}
		if checkRepeatedTriples(word) && checkConsecutiveLetters(word) {
			goodWords++
		}
	}
	return goodWords
}

func has3Vowels(word string) bool {
	cnt := 0
	for _, ch := range word {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			cnt += 1
		}
		if cnt == 3 {
			return true
		}
	}
	return false
}

func hasTwiceLetters(word string) bool {
	var prevCh int32
	for i, ch := range word {
		if i != 0 && prevCh == ch {
			return true
		}
		prevCh = ch
	}
	return false
}

func hasDisallowedThings(word string) bool {
	return strings.Contains(word, "ab") ||
		strings.Contains(word, "cd") ||
		strings.Contains(word, "pq") ||
		strings.Contains(word, "xy")
}

func checkConsecutiveLetters(word string) bool {
	for pos1, pos2 := 0, 1; pos2 < len(word); pos1, pos2 = pos1+1, pos2+1 {
		cha1, cha2 := word[pos1], word[pos2]

		for ps1, ps2 := pos1+2, pos2+2; ps2 < len(word); ps1, ps2 = ps1+1, ps2+1 {
			ch1, ch2 := word[ps1], word[ps2]

			if ch1 == cha1 && ch2 == cha2 {
				return true
			}
		}
	}
	return false
}

func checkRepeatedTriples(word string) bool {
	for pos1, pos2, pos3 := 0, 1, 2; pos3 < len(word); pos1, pos2, pos3 = pos1+1, pos2+1, pos3+1 {
		cha1, cha3 := word[pos1], word[pos3]

		if cha1 == cha3 {
			return true
		}
	}
	return false
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
