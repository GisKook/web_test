package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"
	"fmt"
	"sync"
	"math/rand"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/client"
)

const(
	DST_REMOTE_ADDR_FMT string = "http://222.222.218.50:%s/json"
	DST_SCREEN_SHOT_FMT string = "%s.png"
)

func main(){
	var wg sync.WaitGroup

	ports := []string{"9222", "9223"}
	for _,port:= range ports{
		log.Println("--------------chrome port-------------",port)
		wg.Add(1)
		p := port
		go func(){ 
			dst_remote_addr := fmt.Sprintf(DST_REMOTE_ADDR_FMT, p)
			dst_screen_shot := fmt.Sprintf(DST_SCREEN_SHOT_FMT, p)
			attach_chrome(dst_remote_addr, dst_screen_shot)
		}()
	}
	wg.Wait()
}

func attach_chrome(remote_addr, file_name string) {
	log.Println(remote_addr, file_name)
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithTargets(client.New(client.URL(remote_addr)).WatchPageTargets(ctxt)),chromedp.WithLog(log.Printf))
	// c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	for{
	// run task list
	var buf []byte
	err = c.Run(ctxt, screenshot(`http://222.222.218.52:8023/web/user/main`, `#map_layer0`, &buf))
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(file_name, buf, 0644)
	if err != nil {
		log.Println(err)
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
			a = append(a,chromedp.Sleep(1000*time.Millisecond))
		case 1:
			a = append(a,chromedp.Click(`#right_click`))
			a = append(a,chromedp.Sleep(1000*time.Millisecond))
		case 2:
			a = append(a,chromedp.Click(`#up_click`))
			a = append(a,chromedp.Sleep(1000*time.Millisecond))
		case 3:
			a = append(a,chromedp.Click(`#down_click`))
			a = append(a,chromedp.Sleep(1000*time.Millisecond))
		}
	}

	return a
}

func screenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	log.Println(urlstr)
	log.Println(sel)
	tasks := chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
		chromedp.Sleep(1000 * time.Millisecond),
		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
		chromedp.Sleep(1000 * time.Millisecond),
		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
		chromedp.Sleep(1000 * time.Millisecond),
	}
	// a := gen_action(12,12,29,29)
	a := gen_action(4,4,4,4)
	tasks = append(tasks, a...)
	tasks = append(tasks, chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByID))
	tasks = append(tasks, chromedp.Sleep(1000 * time.Millisecond))
	log.Println("task list ---------------", len(tasks))

	return tasks
	

//	return chromedp.Tasks{
//		chromedp.Navigate(urlstr),
//		chromedp.WaitVisible(sel, chromedp.ByID),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(100 * time.Millisecond),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(100 * time.Millisecond),
//		chromedp.Click(`#map_zoom_slider > div.esriSimpleSliderIncrementButton`, chromedp.NodeVisible),
//		chromedp.Sleep(100 * time.Millisecond),
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
//		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByID),
//	}
}