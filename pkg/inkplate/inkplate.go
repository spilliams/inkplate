package inkplate

import (
	"errors"
	"fmt"
	"log"

	"github.com/albenik/go-serial/v2"
)

// device should be something like "/dev/ttyUSB0"
func New(device string) (Inkplate, error) {
	// Settings are: 115200 baud, standard parity, ending with \n\r
	port, err := serial.Open(device,
		serial.WithBaudrate(115200),
		serial.WithStopBits(serial.TwoStopBits),
	)
	if err != nil {
		log.Fatal(err)
	}
	return Inkplate{
		port: port,
	}, err
}

type Inkplate struct {
	port *serial.Port
}

/////////////////////////
// Basic functionality //
/////////////////////////

func (ip *Inkplate) tx(s string) error {
	b := []byte(s)
	b = append(b, '\n')
	b = append(b, '\r')
	i, err := ip.port.Write(b)
	if err != nil {
		return err
	}
	if i == 0 {
		return errors.New("No bytes written")
	}
	return nil
}

func (ip *Inkplate) rx() (string, error) {
	buff := make([]byte, 100)
	total := 0
	for {
		n, err := ip.port.Read(buff)
		if err != nil {
			return "", err
		}
		total += n
		if n == 0 {
			// end of file
			break
		}
	}
	return string(buff[:total]), nil
}

//////////////////////
// Device Lifecycle //
//////////////////////

func (ip *Inkplate) Close() error {
	return ip.port.Close()
}

func (ip *Inkplate) IsOK() (bool, error) {
	err := ip.tx("#?*")
	if err != nil {
		return false, err
	}
	response, err := ip.rx()
	if err != nil {
		return false, err
	}
	if response != "OK" {
		return false, fmt.Errorf("Response was '%s' instead of 'OK'", response)
	}
	return true, nil
}

func (ip *Inkplate) IsPanelOn() bool {
	// #R(?)*
	return false
}

func (ip *Inkplate) SetPanelSupply(on bool) {
	// #Q(S)*
}

/////////////
// Drawing //
/////////////

func (ip *Inkplate) Print(s string) {
	// #C("STRING")*
}

func (ip *Inkplate) Clear() {
	// #K(1)*
}

func (ip *Inkplate) Display() {
	// #L(1)*
}

func (ip *Inkplate) DrawBitmap(x, y int, path string) {
	// #H(XXX,YYY,"PATH")*
}

func (ip *Inkplate) PartialUpdate(y1, x2, y2 int) {
	// #M(YY1, XX2, YY2)*
}

func (ip *Inkplate) DrawImage(x, y int, path string) {
	// #S(XXX,YYY,"PATH")*
}

func (ip *Inkplate) DrawPixel(x, y, c int) {
	// #0(XXX,YYY,CC)*
}

func (ip *Inkplate) DrawLine(x, y, i, j, c int) {
	// #1(XXX,YYY,III,JJJ,CC)*
}

func (ip *Inkplate) DrawFastVLine(x, y, l, c int) {
	// #2(XXX,YYY,LLL,CC)*
}

func (ip *Inkplate) DrawFastHLine(x, y, l, c int) {
	// #3(XXX,YYY,LLL,CC)*
}

func (ip *Inkplate) DrawRect(x, y, w, h, c int) {
	// #4(XXX,YYY,WWW,HHH,CC)*
}

func (ip *Inkplate) DrawCircle(x, y, r, c int) {
	// #5(XXX,YYY,RRR,CC)*
}

func (ip *Inkplate) DrawTriangle(x1, y1, x2, y2, x3, y3, c int) {
	// #6(XX1,YY1,XX2,YY2,XX3,YY3,CC)*
}

func (ip *Inkplate) DrawRoundedRect(x, y, w, h, r, c int) {
	// #7(XXX,YYY,WWW,HHH,RRR,CC)*
}
func (ip *Inkplate) FillRect(x, y, w, h, c int) {
	// #8(XXX,YYY,WWW,HHH,CC)*
}

func (ip *Inkplate) FillCircle(x, y, r, c int) {
	// #9(XXX,YYY,RRR,CC)*
}

func (ip *Inkplate) FillTriangle(x1, y1, x2, y2, x3, y3, c int) {
	// #A(XX1,YY1,XX2,YY2,XX3,YY3,CC)*
}

func (ip *Inkplate) FillRoundedRect(x, y, w, h, r, c int) {
	// #B(XXX,YYY,WWW,HHH,RRR,CC)*
}

func (ip *Inkplate) DrawThickLine(x, y, i, j, t, c int) {
	// #T(XXX,YYY,III,JJJ,TT,CC)*
}

func (ip *Inkplate) DrawEllipse(x, y, rx, ry, c int) {
	// #U(XXX,YYY,RRX,RRY,CC)*
}

func (ip *Inkplate) FillEllipse(x, y, rx, ry, c int) {
	// #V(XXX,YYY,RRX,RRY,CC)*
}

////////////////////
// Other Controls //
////////////////////

func (ip *Inkplate) SetTextSize(n int) {
	// #D(NN)*
}

func (ip *Inkplate) SetCursor(x, y int) {
	// #E(XXX,YYY)*
}

func (ip *Inkplate) SetTextWrap(on bool) {
	// #F(T/F)*
}

type InkplateRotation int

const (
	InkplateRotation0 InkplateRotation = iota
	InkplateRotation90
	InkplateRotation180
	InkplateRotation270
)

func (ip *Inkplate) SetRotation(r InkplateRotation) {
	// #G(RRR)*
}

type InkplateDisplayMode int

const (
	InkplateDisplayMode3Bit InkplateDisplayMode = iota
	InkplateDisplayMode1Bit
)

func (ip *Inkplate) SetDisplayMode(m InkplateDisplayMode) {
	// #I(D)*
}

func (ip *Inkplate) GetDisplayMode() InkplateDisplayMode {
	// #J(?)*
	return InkplateDisplayMode3Bit
}

func (ip *Inkplate) ReadTempCelcius() int {
	// #N(?)*
	return 0
}

type InkplateTouchpad int

const (
	InkplateTouchpadLeft InkplateTouchpad = iota
	InkplateTouchpadCenter
	InkplateTouchpadRight
)

func (ip *Inkplate) TouchpadHigh(t InkplateTouchpad) bool {
	// #O(P)*
	return false
}

func (ip *Inkplate) ReadBatteryVoltage() float32 {
	// #P(?)*
	return 0
}
