// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validate

var commonPasswords = map[string]bool{
	"000000":      true,
	"111111":      true,
	"11111111":    true,
	"112233":      true,
	"121212":      true,
	"123123":      true,
	"123456":      true,
	"1234567":     true,
	"12345678":    true,
	"123456789":   true,
	"131313":      true,
	"232323":      true,
	"654321":      true,
	"666666":      true,
	"696969":      true,
	"777777":      true,
	"7777777":     true,
	"8675309":     true,
	"987654":      true,
	"aaaaaa":      true,
	"abc123":      true,
	"abcdef":      true,
	"abgrtyu":     true,
	"access":      true,
	"access14":    true,
	"action":      true,
	"albert":      true,
	"alberto":     true,
	"alexis":      true,
	"alejandra":   true,
	"alejandro":   true,
	"amanda":      true,
	"amateur":     true,
	"america":     true,
	"andrea":      true,
	"andrew":      true,
	"angela":      true,
	"angels":      true,
	"animal":      true,
	"anthony":     true,
	"apollo":      true,
	"apples":      true,
	"arsenal":     true,
	"arthur":      true,
	"asdfgh":      true,
	"ashley":      true,
	"asshole":     true,
	"august":      true,
	"austin":      true,
	"badboy":      true,
	"bailey":      true,
	"banana":      true,
	"barney":      true,
	"baseball":    true,
	"batman":      true,
	"beatriz":     true,
	"beaver":      true,
	"beavis":      true,
	"bigcock":     true,
	"bigdaddy":    true,
	"bigdick":     true,
	"bigdog":      true,
	"bigtits":     true,
	"birdie":      true,
	"bitches":     true,
	"biteme":      true,
	"blazer":      true,
	"blonde":      true,
	"blondes":     true,
	"blowjob":     true,
	"blowme":      true,
	"bond007":     true,
	"bonita":      true,
	"bonnie":      true,
	"booboo":      true,
	"booger":      true,
	"boomer":      true,
	"boston":      true,
	"brandon":     true,
	"brandy":      true,
	"braves":      true,
	"brazil":      true,
	"bronco":      true,
	"broncos":     true,
	"bulldog":     true,
	"buster":      true,
	"butter":      true,
	"butthead":    true,
	"calvin":      true,
	"camaro":      true,
	"cameron":     true,
	"canada":      true,
	"captain":     true,
	"carlos":      true,
	"carter":      true,
	"casper":      true,
	"charles":     true,
	"charlie":     true,
	"cheese":      true,
	"chelsea":     true,
	"chester":     true,
	"chicago":     true,
	"chicken":     true,
	"cocacola":    true,
	"coffee":      true,
	"college":     true,
	"compaq":      true,
	"computer":    true,
	"consumer":    true,
	"cookie":      true,
	"cooper":      true,
	"corvette":    true,
	"cowboy":      true,
	"cowboys":     true,
	"crystal":     true,
	"cumming":     true,
	"cumshot":     true,
	"dakota":      true,
	"dallas":      true,
	"daniel":      true,
	"danielle":    true,
	"debbie":      true,
	"dennis":      true,
	"diablo":      true,
	"diamond":     true,
	"doctor":      true,
	"doggie":      true,
	"dolphin":     true,
	"dolphins":    true,
	"donald":      true,
	"dragon":      true,
	"dreams":      true,
	"driver":      true,
	"eagle1":      true,
	"eagles":      true,
	"edward":      true,
	"einstein":    true,
	"erotic":      true,
	"estrella":    true,
	"extreme":     true,
	"falcon":      true,
	"fender":      true,
	"ferrari":     true,
	"firebird":    true,
	"fishing":     true,
	"florida":     true,
	"flower":      true,
	"flyers":      true,
	"football":    true,
	"forever":     true,
	"freddy":      true,
	"freedom":     true,
	"fucked":      true,
	"fucker":      true,
	"fucking":     true,
	"fuckme":      true,
	"fuckyou":     true,
	"gandalf":     true,
	"gateway":     true,
	"gators":      true,
	"gemini":      true,
	"george":      true,
	"giants":      true,
	"ginger":      true,
	"gizmodo":     true,
	"golden":      true,
	"golfer":      true,
	"gordon":      true,
	"gregory":     true,
	"guitar":      true,
	"gunner":      true,
	"hammer":      true,
	"hannah":      true,
	"hardcore":    true,
	"harley":      true,
	"heather":     true,
	"helpme":      true,
	"hentai":      true,
	"hockey":      true,
	"hooters":     true,
	"horney":      true,
	"hotdog":      true,
	"hunter":      true,
	"hunting":     true,
	"iceman":      true,
	"iloveyou":    true,
	"internet":    true,
	"iwantu":      true,
	"jackie":      true,
	"jackson":     true,
	"jaguar":      true,
	"jasmine":     true,
	"jasper":      true,
	"jennifer":    true,
	"jeremy":      true,
	"jessica":     true,
	"johnny":      true,
	"johnson":     true,
	"jordan":      true,
	"joseph":      true,
	"joshua":      true,
	"junior":      true,
	"justin":      true,
	"killer":      true,
	"knight":      true,
	"ladies":      true,
	"lakers":      true,
	"lauren":      true,
	"leather":     true,
	"legend":      true,
	"letmein":     true,
	"little":      true,
	"london":      true,
	"lovers":      true,
	"maddog":      true,
	"madison":     true,
	"maggie":      true,
	"magnum":      true,
	"marine":      true,
	"mariposa":    true,
	"marlboro":    true,
	"martin":      true,
	"marvin":      true,
	"master":      true,
	"matrix":      true,
	"matthew":     true,
	"maverick":    true,
	"maxwell":     true,
	"melissa":     true,
	"member":      true,
	"mercedes":    true,
	"merlin":      true,
	"michael":     true,
	"michelle":    true,
	"mickey":      true,
	"midnight":    true,
	"miller":      true,
	"mistress":    true,
	"monica":      true,
	"monkey":      true,
	"monster":     true,
	"morgan":      true,
	"mother":      true,
	"mountain":    true,
	"muffin":      true,
	"murphy":      true,
	"mustang":     true,
	"naked":       true,
	"nascar":      true,
	"nathan":      true,
	"naughty":     true,
	"ncc1701":     true,
	"newyork":     true,
	"nicholas":    true,
	"nicole":      true,
	"nipple":      true,
	"nipples":     true,
	"oliver":      true,
	"orange":      true,
	"packers":     true,
	"panther":     true,
	"panties":     true,
	"parker":      true,
	"password":    true,
	"password1":   true,
	"password12":  true,
	"password123": true,
	"patrick":     true,
	"peaches":     true,
	"peanut":      true,
	"pepper":      true,
	"phantom":     true,
	"phoenix":     true,
	"player":      true,
	"please":      true,
	"pookie":      true,
	"porsche":     true,
	"prince":      true,
	"princess":    true,
	"private":     true,
	"purple":      true,
	"pussies":     true,
	"qazwsx":      true,
	"qwerty":      true,
	"qwertyui":    true,
	"rabbit":      true,
	"rachel":      true,
	"racing":      true,
	"raiders":     true,
	"rainbow":     true,
	"ranger":      true,
	"rangers":     true,
	"rebecca":     true,
	"redskins":    true,
	"redsox":      true,
	"redwings":    true,
	"richard":     true,
	"robert":      true,
	"roberto":     true,
	"rocket":      true,
	"rosebud":     true,
	"runner":      true,
	"rush2112":    true,
	"russia":      true,
	"samantha":    true,
	"sammy":       true,
	"samson":      true,
	"sandra":      true,
	"saturn":      true,
	"scooby":      true,
	"scooter":     true,
	"scorpio":     true,
	"scorpion":    true,
	"sebastian":   true,
	"secret":      true,
	"sexsex":      true,
	"shadow":      true,
	"shannon":     true,
	"shaved":      true,
	"sierra":      true,
	"silver":      true,
	"skippy":      true,
	"slayer":      true,
	"smokey":      true,
	"snoopy":      true,
	"soccer":      true,
	"sophie":      true,
	"spanky":      true,
	"sparky":      true,
	"spider":      true,
	"squirt":      true,
	"srinivas":    true,
	"startrek":    true,
	"starwars":    true,
	"steelers":    true,
	"steven":      true,
	"sticky":      true,
	"stupid":      true,
	"success":     true,
	"suckit":      true,
	"summer":      true,
	"sunshine":    true,
	"superman":    true,
	"surfer":      true,
	"swimming":    true,
	"sydney":      true,
	"tequiero":    true,
	"taylor":      true,
	"tennis":      true,
	"teresa":      true,
	"tester":      true,
	"testing":     true,
	"theman":      true,
	"thomas":      true,
	"thunder":     true,
	"thx1138":     true,
	"tiffany":     true,
	"tigers":      true,
	"tigger":      true,
	"tomcat":      true,
	"topgun":      true,
	"toyota":      true,
	"travis":      true,
	"trouble":     true,
	"trustno1":    true,
	"tucker":      true,
	"turtle":      true,
	"twitter":     true,
	"united":      true,
	"vagina":      true,
	"victor":      true,
	"victoria":    true,
	"viking":      true,
	"voodoo":      true,
	"voyager":     true,
	"walter":      true,
	"warrior":     true,
	"welcome":     true,
	"whatever":    true,
	"william":     true,
	"willie":      true,
	"wilson":      true,
	"winner":      true,
	"winston":     true,
	"winter":      true,
	"wizard":      true,
	"xavier":      true,
	"xxxxxx":      true,
	"xxxxxxxx":    true,
	"yamaha":      true,
	"yankee":      true,
	"yankees":     true,
	"yellow":      true,
	"zxcvbn":      true,
	"zxcvbnm":     true,
	"zzzzzz":      true}
