package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func Klaatu() cli.Command {
	return cli.Command{
		Name:  "klaatu",
		Usage: "barada nikto",
		Action: func(c *cli.Context) {
			if len(c.Args()) == 2 {
				if c.Args()[0] == "barada" && c.Args()[1] == "nikto" {
					fmt.Println(`tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt11111111111111111111111111t1ttttttttttttttt111111ttttttt11t11111111111111111111111111i
tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt111111111111111111111ttttttttttttttttttttttttttttttt11111111111111111111111111111i
tttttttftttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt11111tttttttttttttttttttttttttttttttttttttttttttttttttt111111111111111111111
ffffffffffffffffffftttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt1111111111111111111
ffffffffffffffffffffffffffftttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt1111111111111111111
ffffffffffffffffffffffffffffftttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt1111111111111111111
fffffffffffffffffffffffffffffffttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt111111111111111111
fffffffffffffffffffffffffffffffffffttttftttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt1111111111111111
ffffffffffffffffffffffffffffffffffffffttttttttttttttttttttttftttt1ii;;;;;;;iiii11tttttttttttttttttttttttttttttttttttttttttttttttttttt11111111111111111
fffffffffffffffffffffffffffffffffffffffffftttttttttttttttttt1iii;;;;:::;;;;;;;;;iiii111tttttttttttttttttttttttttttttttttttttttttttttttt111111111111111
ffffffffffffffffffffffffffffffffffffffffffftttttttttttttt11iiii;;;::::::::;;;;;;;iii11111ttttttttttttttttttttttttttttttttttttttttttttttt11111111111111
ffffffffffffffffffffffffffffffffffffffffffftttttttttttt1111iiiii;;;;;;;;;;;iiiiii1111ttttttttttttttttttttttttttttttttttttttttttttttttt1111111111111111
fffffffffffffffffffffffffffffffffffffffffffffftftttttttttt11111iiiiiiiiii1111111tttttttffffftttttttttttttttttttttttttttttttttttttttttttt11111111111111
fffffffffffffffffffffffffffffffffffffffffffffffffftttttttttttttttttttttttfffffffffffffffffffffttttttttttttttttttttttttttttttttttttttttttttt11111111111
fffffffffffffffffffffffffffffffffffffffffffffffftfffffffffffffffffLLLLLLLLLLLLLLLLLLLLLLLLLLLLLfttttttttttttttttttttttttttttttttttttttttttt11111111111
fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffLLLCCCCCCCCGGGGGCCCCCCCCCCCCCCCLLLLffttttttttttttttttttttttttttttttttttttttttt11111111111
fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffLLCCCGGGGGGGGGGGGGGGGGGGGGGGCCCCCCLLffttttttttttttttttttttttttttttttttttttttt1t1111111111
fffffffffffffffffffffffffffffffffffffffffffffffLLLLfffftttttffLLCCCGGGGGGGGGGGG0G0GGGGGGGGCCCCCLLLfttttttfftttttttttttttttttttttttttttttttt11111111111
fffffffffffffffffffffffffffffffffffffffffffffffLLLLfftt1ii1ttfLLCCCCGGGGGGGG0000GGGGGGGGGGCCCCLLLLLfttttttfttttttttttttttttttttttttttttttttt1111111111
ffffffffffffffffffffffffffffffffffffffffffffffLLLLLft11iii11tfLLCCCCGGGGGGG00000GGG00GGGGGCCCCLLfffLtftfffftttttttttttttttttttttttttttttttt11111111111
ffffffffffffffffffffffffffffffffffffffffffffffLCLLLft1iiii11tfLLLCCCGGGGGGG00000000000GGGGCCLLfffffLfffffffftttttttttttttttttttttttttttttttt1111111111
ffffffffffffffffffffffffffffffffffffffffffffffCCCLLft11ii11ttffLCCCGGGGGGG0000000000000GGGCCLft1ttffftfffffftttttttttttttttttttttttttttttttt1111111111
fffffffffffffffffffffffffffffffffffffffffffffLCCCLLfft111ttttfLLCCGGGGG0000000000000000GGGCLft11tttL;1ftffffttttttttfttttttttttttttttttttttt1111111111
fffffffffffffffffffffffffffffffffffffffffffffLCCCLLffttt1ttttfLLCGCCCGGGGGGGGG000000GGGGGGCLfti1ttt0;L1tfftftttttttttttttttttttttttttttttttt1111111111
ffffffffffffffffffffffffffffffffffffffffffffffLLLft1111111tttfLLCCGGGGG00000000000000000GGCf1i;i1tL0CCiLfftffffttttttttttttttttttttttttttttttt11111111
ffffffffffffffffffffffffffffffffffffffffffffLGCCLfLLCG0888@@@@88888888@@@88@888888888880GGLLf1i1ttG8G0Gt1ffffffftttttttttttttttttttttttttttt1111111111
ffffffffffffffffffffffffffffffffffffffffffffL8@@@@@8@88@G1LCt;:::;fC00888888888888888888888C1tfGL11ff8G8Lfffttftttttttttttttttttttttttttttt11111111111
ffffffffffffffffffffffffffffffffffffffffffffLC8@@888@8G1i,....,....      ...;G0888888888888G1tttffG8L808tffttttttttttttttttttttttttttttttt1t1111111111
ffffffffffffffffffffffffffffffffffffffffffff0G8@@88@@88@8L1;.,,.....,,.:itL1f888888888888888ttttftG000G8fffttttttttttttttttttttttttttttttt1t1111111111
ffffffffffffffffffffffffffffffffffffffffffffLC8@@@@8888@C t8GL11tLG088888@8888888888888888881tt1ft0008CGLtfttttttttttttttttttttttttttttttt111111111111
ffffffffffffffffffffffffffffffffffffffffffffft0@@008@t.f0C08888888888888888@88888880000000CLftttttG0G@Cttfftttttttttttttttttttttttttttttt111111111111i
fffffffffffffffffffffffffffffffffffffffffffffftf0G;.fftttttttfLLCCCCCCGGGG00000000000000GGCLffttffC8Gfttttttttttttttttttttttttttttttttttttt1111111111i
fffffffffffffffffffffffffffffffffffffffffffffft,,CGCCLfffffLLLCCCGGGGGGGG000000000000000GGCfftfffLCf1tttttttttttttttttttttttttttttttttttt1111111111111
ffffffffffffffffffffffffffffffffffffffffffft,;ftCGGGCCLfffLLLLCCCCGGGGGGG000000000000000GGLffttfLLGttttttttttttttttttttttttttttttttttttt1111111111111i
fffffffffffffffffffffffffffffffffffffffL1,ifffffLGGGCCLfffLLLLCCCCGGGGG0000000000000000GGCLfftfLfGttttttttttttttttttttttttttttttttttttttt111111111111i
fffffffffffffffffffffffffffffffffffffi,1LffffffffGGGGCCLLLLLLLLCCCGGGG00000000000000000GGCLffffLCtttttfttttttttttttttttttttttttttttttttttt111111111111
ffffffffffffffffffffffffffffffffff;:tLfffffffffftGGGGCCLfffffLLLCCGGGGG000000000000000GGGLffffLCttttttttttttttttttttttttttttttttttttttttt1111111111111
fffffffffffffffffffffffffffffft::tLffffffffffffftLGGGCCLLffffLLLCCGGGGGG00000000000000GGCLLLLLGtttttttttttttttttttttttttttttttttttttttttt1111111111111
ffffffffffffffffffffffffffLt:;fffffffffffffffffftfGGGGCLLLffLLLLCCCGGGGG0000000000000GGGCLLLLCftttttttttttttttttttttttttttttttttttttttttt11111111111ii
ffffffffffffffffffffffff1:ifffffffffffffffffffffftCGGGCLLffLLLLLCCGGGGGG0000000000000GGCLLLCCCtttttttttttttttttttttttttttttttttttttttttt111111111111ii
fffffffffffffffffffffi:ifLffffffffffffffffffffffttL0GGCLLLLffLLLCCCGGGGGG00000000000GGGCLLCGttLttttttttttttttttttttttttttttttttttttttt1111111111111iii
ffffffffffffffffft;:1fffffffffffffffffffffffffftfttGGGCCLffffLLLCCCGGGGGG00000000000GGCCCCC11tfttttttttttttttttttttttttttttttttttt11111111111111111iii
fffffffffffffft;:1ffffffffffffffffffffffffffffffftttGGCCLfffLLLLCCCGGGGG00000000000GGGGGGLLftffLttttttttttt1111111111111111111tttt111t11111111111111ii
fffffffffff1;;1fffffffffffffffffffffffffffffffffftttt1CGGGGCCCLCGGGGGGG0000000000000000GGGCLfLCCCLi;iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii111111111iiiii
fffffffti;;1fffffffffffffffffffffffffffffffffffffft1itCGGGG0880000008000000000000000000GGGGCCGGCLLLLLfft1iiiii;iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii
ffffti;;1ffffffffffffffffffffffffffffffffffft1iiiiii1LCCGGCCCCG008880000000000000000000GGGGGGGCCCLLLLffffi::;;iiiii;;;;;iiiii;;;;;;;;;;;;iiiiii;;;ii;i
f1;;;tffffffffffffffffffffffffffffffft11iiiiiiiiii1itLCCCCCCLLLLCGG00000000000000000GGGGGGGGGCCGGGGCCLLfftti::;::;;;;;iiiii;;;;;;;;;;;;;;;;;;;;;i;;;;;
;itfffffffffffffffffffffffffffft11iiiiiiiiiiii111i1LffLLCCCLffLLLLCCGGGG000000000GGGGGGGGGGGGGGGGGGGCCLLLfffi:;;;;;;;;;;;;;;;;iiii;;;;;;;;;;;;;;;;;;;;
ffffffffffffffffffffffffftt11111iiiiiiiiiiiii11i1itfffffffftttttttfLCCCGGGGGGGGGGGGGGGGGGGGGGGGGGGGGCCCLLLLfff;;;;;;;;;;;;;;;;;;;;;;;;;i;;;;;;;;;;;;;;
fffffffffffffffffffftt1111111111111111111iii1111tLffffffft1iiiiii1ttffLCCCCCCGGGGGGGGGGGGGGGGGGGGGGCCCLLLLLfff1iiiiiiii;;;;;;;;;;;;;;;;;;;1;:;;;;;;;;;
ffffffffffffffffttt1111111111111111111111111111itffffffffffftt1111ttffLLCCCCGGGGGGGGGGGGGGGGGGCCCCCCLLLLLLLLfffiiiiiiiiiiiiiiiiiiiiiiiiii1tti:;;;;;;;;
ffffffffffffffffftttttttt1111111111111111111111iffffffffffLLLLLLLCCLCCCGGGGGGGGGGGGGGGGGGCCCCCCCCCLLLLLLLLfLfffi111111111111111111111111tfffti:;;;;;;;
fffffffffffttffffffttttttttttttttttt1111111111i1ffffffffffLLLCCCCCCCCCGGGGGGGGGGGGGGGCCCCCCCCCLLLLLLLLLLLLLLffL1111111tttttttttttttttttffLLLfti;;;;;;;
fffffffffffttfffffffffttttttttttttttt1111111111tfftffffffffLLLCCCCCCCCCCGGGGGGGGGGCCCCCCCCCLLLLLLLLLLLLLLLLLLLLftttttttfffffffffffttffLLLLLLftt;;;;;;;
ffffffffttttffffffffftttttttttttttttt1111111111ttttttfffffffLLLLLCCCCCCCCCCCCCCCCCCCCCCLLLLLLLLLLLLLLLLLLLLLLLLLfffffffffffffffffffffLLLLLLLLft1:;;;;;
fftttttttttfLfffffttttttttttttttttttttt11111111tttttttttfffffffLLLLLLLLLLLLCCCCCLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLfffffLLLLLLLLLLLLLftti:;;;;
tttttttttttLLffffttttttt11tttttttttttttt11111111tttttttttffffffffffLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLCCCLLLLLLLLLLLLLLLLLLLLLLLLLLLfftt;;;;;
ttttttttttfLffffttttttt11ttttttttttttttt11t11111ttttttttttttttfffffffffLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLCCCCCLLLLLLLLLLLLLLLLLLLLCCLLLLLff11:::;
ttttttttttLLfffttttttttttttffftttttttttt1111111111tttttttttttttffffffffffLLLLLLLLLffffffLLLLLLLLLLLLLLLLLLLCCCCCCCCLLLLLLLLLLLLLLLLLLLLLLLLLLLfftti::;
tttttttttLLLfffftttttttttffffffttttttt11111111111111tttttttttttttttffffffffffffLfffffffffffLLLLLLLLLLLLLLLCCCCCCCCCCLLLLLLLLLLLLLLLLLLLLLLLLLLLft11:::`)
				}
			} else {
				fmt.Println("Did you mean Klaatu barada nikto?")
			}
		},
	}
}
