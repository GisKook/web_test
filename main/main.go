//// Command click is a chromedp example demonstrating how to use a selector to
//// click on an element.
//package main
//
//import (
//	"context"
//	"log"
//	"time"
//
//	"github.com/chromedp/chromedp"
//	cdpclient "github.com/chromedp/chromedp/client"
//)
//
//func main() {
//	var err error
//
//	// create context
//	ctxt, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	// create chrome instance
//	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
//	if err != nil {
//		log.Fatal(err)
//	}
//	// test
//	for i:=0 ; i<0; i++ {
//	client := cdpclient.New()
//t, err := client.NewPageTargetWithURL(ctxt, "http://127.0.0.1:8080/web/user/main")
//if err != nil {
//		log.Fatal(err)
//}
//
//h, err := chromedp.NewTargetHandler(t, log.Printf, log.Printf, log.Printf)
//if err != nil {
//		log.Fatal(err)
//}
//
//if err := h.Run(ctxt); err != nil {
//		log.Fatal(err)
//}
//click().Do(ctxt,h)
//	}
//	// test
//
//	// run task list
//	err = c.Run(ctxt, click())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// shutdown chrome
//	err = c.Shutdown(ctxt)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// wait for chrome to finish
//	err = c.Wait()
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func click() chromedp.Tasks {
//	return chromedp.Tasks{
//		chromedp.Navigate(`http://127.0.0.1:8080/web/user/main`),
//		chromedp.WaitVisible(`#map_layer0`),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(5 * time.Second),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderDecrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(5 * time.Second),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(5 * time.Second),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderDecrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(5 * time.Second),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(5 * time.Second),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderDecrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(300 * time.Second),
//	}
//}

package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/client"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithTargets(client.New().WatchPageTargets(ctxt)),chromedp.WithLog(log.Printf))
	// c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var buf []byte
	err = c.Run(ctxt, screenshot(`http://222.222.218.52:8023/web/user/main`, `#map_layer0`, &buf))
	if err != nil {
		log.Println("YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY ....")
		log.Fatal(err)
	}
	log.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX ....")

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("contact-form.png", buf, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func screenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	log.Println("------------------------------------")
	log.Println(urlstr)
	log.Println(sel)

	return chromedp.Tasks{
	//	chromedp.Navigate(`http://127.0.0.1:8080/web/user/main`),
	//	chromedp.WaitVisible(`#map_layer0`),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.Sleep(2 * time.Second),
		//chromedp.WaitVisible(`#map_layer0`),
	//	chromedp.WaitNotVisible(`div.v-middle > div.la-ball-clip-rotate`, chromedp.ByQuery),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByID),
	}
}