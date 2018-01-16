// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"log"
	"time"
	"math/rand"

	"github.com/chromedp/chromedp"
	// cdpclient "github.com/chromedp/chromedp/client"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}
	// test
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
	// test

	// run task list
	for{
	err = c.Run(ctxt, click())
	if err != nil {
		log.Fatal(err)
	}
}

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
}


type action struct{
	used int; 
	value int;
}
func gen_action(lm,rm, um, dm int) []chromedp.Action{ 
	rand.Seed(time.Now().Unix())
	la := rand.Intn(lm)
	ra := rand.Intn(rm)
	ua := rand.Intn(um)
	da := rand.Intn(dm) 
	t := la+ra+ua+da
	actions := make([]*action, t)
	tt := []int{la, ra, ua, da}
	for i, va := range tt{
		v := 0
		log.Println(va)
		for v < va{
			d:=rand.Intn(t)
			if actions[d] == nil{
				actions[d] = new(action)
				actions[d].used = 1 
				actions[d].value = i // 0 for left 1 right 2 up 3 down
				v++
			}
		}
	}
	a := make([]chromedp.Action,0)
	for _, v := range actions{ 
		switch v.value{
		case 0:
			a = append(a,chromedp.Click(`#left_click`))
			a = append(a,chromedp.Sleep(500*time.Millisecond))
		case 1:
			a = append(a,chromedp.Click(`#right_click`))
			a = append(a,chromedp.Sleep(500*time.Millisecond))
		case 2:
			a = append(a,chromedp.Click(`#up_click`))
			a = append(a,chromedp.Sleep(500*time.Millisecond))
		case 3:
			a = append(a,chromedp.Click(`#down_click`))
			a = append(a,chromedp.Sleep(500*time.Millisecond))
		}
	}

	return a
}

func click() chromedp.Tasks {
	tasks := chromedp.Tasks{
		chromedp.Navigate(`http://127.0.0.1:8080/web/user/main`),
		chromedp.WaitVisible(`#map_layer0`),
		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
		chromedp.Sleep(100 * time.Millisecond),
		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
		chromedp.Sleep(100 * time.Millisecond),
		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
		chromedp.Sleep(100 * time.Millisecond),
	}
	a := gen_action(12,12,29,29)
	tasks = append(tasks, a...)

	return tasks
}

//func click() chromedp.Tasks {
//	return chromedp.Tasks{
//		chromedp.Navigate(`http://127.0.0.1:8080/web/user/main`),
//		chromedp.WaitVisible(`#map_layer0`),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#left_click`),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#left_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#left_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#up_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#up_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#up_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#right_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#right_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#right_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#right_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#right_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#right_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#down_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#down_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#down_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#left_click`),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#left_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#left_click`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderDecrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderDecrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderDecrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(1000 * time.Millisecond),
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