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
///////////////////////////////////////////////////////example2////////////////////////////////////////////////////////
//package main
//
//import (
//	"context"
//	"io/ioutil"
//	"log"
//	"time"
//
//	"github.com/chromedp/chromedp"
//	"github.com/chromedp/chromedp/client"
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
//	c, err := chromedp.New(ctxt, chromedp.WithTargets(client.New().WatchPageTargets(ctxt)),chromedp.WithLog(log.Printf))
//	// c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// run task list
//	var buf []byte
//	err = c.Run(ctxt, screenshot(`http://222.222.218.52:8023/web/user/main`, `#map_layer0`, &buf))
//	if err != nil {
//		log.Println("YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY ....")
//		log.Fatal(err)
//	}
//	log.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX ....")
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
//
//	err = ioutil.WriteFile("contact-form.png", buf, 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func screenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
//	log.Println("------------------------------------")
//	log.Println(urlstr)
//	log.Println(sel)
//
//	return chromedp.Tasks{
//	//	chromedp.Navigate(`http://127.0.0.1:8080/web/user/main`),
//	//	chromedp.WaitVisible(`#map_layer0`),
////		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
//		chromedp.Navigate(urlstr),
//		chromedp.WaitVisible(sel, chromedp.ByID),
//		chromedp.Sleep(2 * time.Second),
//		//chromedp.WaitVisible(`#map_layer0`),
//	//	chromedp.WaitNotVisible(`div.v-middle > div.la-ball-clip-rotate`, chromedp.ByQuery),
//		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByID),
//	}
//}
///////////////////////////////////////////////////////example3////////////////////////////////////////////////////////
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create pool
	pool, err := chromedp.NewPool( chromedp.PortRange(9223,9223)/*chromedp.PoolLog(log.Printf, log.Printf, log.Printf)*/ )
	if err != nil {
		log.Fatal(err)
	}

	// loop over the URLs
	var wg sync.WaitGroup
	for i, urlstr := range []string{
		"http://222.222.218.52:8023/web/user/main",
	} {
		wg.Add(1)
		go takeScreenshot(ctxt, &wg, pool, i, urlstr)
	}

	// wait for to finish
	wg.Wait()

	// shutdown pool
	err = pool.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}

func takeScreenshot(ctxt context.Context, wg *sync.WaitGroup, pool *chromedp.Pool, id int, urlstr string) {
	defer wg.Done()

	// allocate
	c, err := pool.Allocate(ctxt,runner.HeadlessPathPort("/headless_shell/headless_shell", 9223))
	if err != nil {
		log.Printf("url (%d) `%s` error: %v", id, urlstr, err)
		return
	}
	defer c.Release()

	// run tasks
	var buf []byte
	err = c.Run(ctxt, screenshot(urlstr, &buf))
	if err != nil {
		log.Printf("url (%d) `%s` error: %v", id, urlstr, err)
		return
	}

	// write to disk
	err = ioutil.WriteFile(fmt.Sprintf("%d.png", id), buf, 0644)
	if err != nil {
		log.Printf("url (%d) `%s` error: %v", id, urlstr, err)
		return
	}
}

func screenshot(urlstr string, picbuf *[]byte) chromedp.Action {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Sleep(2 * time.Second),
		chromedp.WaitVisible(`#map_layer0`),
//		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByID),
		chromedp.ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
			buf, err := page.CaptureScreenshot().Do(ctxt, h)
			if err != nil {
				return err
			}
			*picbuf = buf
			return nil
		}),
	}
}