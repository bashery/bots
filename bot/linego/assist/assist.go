package main

import (
	"fmt"
	"strings"
	"runtime"
	"time"
	"io/ioutil"


	"./linego/LineThrift"
	"./linego/auth"
	"./linego/helper"
	"./linego/service"
	"./linego/talk"
)



func checkEqual(list1 []string, list2 []string) bool {
	for _, v := range list1 {
		if helper.InArray(list2, v) {
			return true
		}
	}
	return false
}

func canceling(group string, korban []string) {
					runtime.GOMAXPROCS(10)
					for _, vo := range korban {
						cancel := []string{vo}
						if helper.InArray(service.Banned, vo) {
							go func() {
								talk.CancelGroupInvitation(group, cancel)
							}()
						}
					}
}

func addbl(pelaku string) {
		if !helper.IsBanned(pelaku) {
				service.Banned = append(service.Banned, pelaku)
		}
}

func cancelall(group string, korban []string) {
					runtime.GOMAXPROCS(10)
					for _, vo := range korban {
						    cancel := []string{vo}
							go func() {
								talk.CancelGroupInvitation(group, cancel)
							}()
					}
}

func checkurl(group string, pelaku string) {
			runtime.GOMAXPROCS(10)
			go func() {
				res, _ := talk.GetGroup(group)
				cek := res.PreventedJoinByTicket
				if !cek {
					go func() {
						res.PreventedJoinByTicket = true
						talk.UpdateGroup(res)
					}()
					go func() {
						talk.KickoutFromGroup(group, []string{pelaku})
					}()
				}
			}()
}


func bot(op *LineThrift.Operation) {
	var Mid string = service.MID

	if op.Type == 26 {
		msg := op.Message
		sender := msg.From_
		if sender != "" && helper.IsAccess(msg.From_) {
			if strings.ToLower(msg.Text) == "speed" {
				start := time.Now()
				talk.SendMessage(msg.To, "es....", 2)
				elapsed := time.Since(start)
				stringTime := elapsed.String()
				talk.SendMessage(msg.To, stringTime, 2)
			} else if strings.ToLower(msg.Text) == ".res" {
				basename := "bot9"
				talk.SendMessage(msg.To, basename, 2)
			} else if strings.ToLower(msg.Text) == ".kick" {
				runtime.GOMAXPROCS(100)
				res, _ := talk.GetGroup(msg.To)
				memlist := res.Members
				for _, v := range memlist {
					if !helper.IsAccess(v.Mid) {
						cancel := []string{v.Mid}
						go func() {
							talk.KickoutFromGroup(msg.To, cancel)
						}()
					}
				}
			} else if strings.ToLower(msg.Text) == ".cancelall" {
				runtime.GOMAXPROCS(100)
				res, _ := talk.GetGroup(msg.To)
				memlist := res.Invitee
				for _, v := range memlist {
					if !helper.IsAccess(v.Mid) {
						cancel := []string{v.Mid}
						go func() {
							talk.CancelGroupInvitation(msg.To, cancel)
						}()
					}
				}
			} else if strings.ToLower(msg.Text) == "bye" {
				talk.LeaveGroup(msg.To)
			} else if strings.ToLower(msg.Text) == "banned" {
				fmt.Println("banned")
				if service.Banned != nil {
					for _, i:= range service.Banned {
						talk.SendMessage(msg.To, i, 2)
						fmt.Println(i)
					}
				} else {
					talk.SendMessage(msg.To, "empty", 2)
				}
			} else if strings.HasPrefix("addban", msg.Text) {
				fmt.Println("prefix")
			} else if strings.ToLower(msg.Text) == "clearban" {
				service.Banned = []string{}
				talk.SendMessage(msg.To, "Done", 2)
			} else if strings.ToLower(msg.Text) == ".invitebot" {
				    talk.InviteIntoGroup(msg.To, service.Squad)
			} else if strings.ToLower(msg.Text) == ".open" {
				res, _ := talk.GetGroup(msg.To)
				cek := res.PreventedJoinByTicket
				if cek {
					res.PreventedJoinByTicket = false
					talk.UpdateGroup(res)
					fmt.Println("done")
				} else {
					res.PreventedJoinByTicket = true
					talk.UpdateGroup(res)
					fmt.Println("done")
			    }
            }
		}
	} else if op.Type == 19 {
		korban := op.Param3
		kicker := op.Param2
		group := op.Param1

		if korban == Mid {
            go func() {
            	addbl(kicker)
            }()
		} else if helper.IsAccess(korban) {
			runtime.GOMAXPROCS(10)
			go func() {
				talk.KickoutFromGroup(group, []string{kicker})
			}()
			go func() {
				talk.InviteIntoGroup(group, service.Squad)
			}()
			go func() {
            	addbl(kicker)
            }()
		}

	} else if op.Type == 13 {
		runtime.GOMAXPROCS(30)
	    korban := strings.Split(op.Param3, "\x1e")
		inviter := op.Param2
		group := op.Param1

		if helper.InArray(korban, Mid) {
			if helper.IsAccess(inviter) {
				go func() {
					talk.AcceptGroupInvitation(group)
				}()
			}

		} else if helper.IsBanned(inviter) {
					go func() {
						cancelall(group, korban)
					}()
					go func() {
						talk.KickoutFromGroup(group, []string{inviter})
					}()
		}  else {
			if !helper.IsAccess(inviter) {
					go func() {
						canceling(group, korban)
						}()
					}
					if checkEqual(korban, service.Banned) {
						go func() {
							talk.KickoutFromGroup(group, []string{inviter})
						}()
						go func() {
            				addbl(inviter)
            			}()
					}
		}

	} else if op.Type == 32 {
		runtime.GOMAXPROCS(20)
		korban := op.Param3
		kicker := op.Param2
		group := op.Param1

		if helper.IsAccess(korban) && !helper.IsAccess(kicker) {
			go func(){
				talk.KickoutFromGroup(group, []string{kicker})
			}()
			go func(){
				talk.InviteIntoGroup(group, service.Squad)
			}()
			go func() {
            	addbl(kicker)
            }()
		}
	} else if op.Type == 17 {
		runtime.GOMAXPROCS(10)
		kicker := op.Param2
		group := op.Param1

		if helper.IsBanned(kicker) {
			go func(){
				talk.KickoutFromGroup(group, []string{kicker})
			}()
		}
	} else if op.Type == 11 {
		runtime.GOMAXPROCS(10)
		changer := op.Param2
		group := op.Param1
		//par := op.Param3

		if !helper.IsAccess(changer) {
			go func(){
				checkurl(group, changer)
			}()
		}
	}
}


func main() {

	pc, file, line, ok := runtime.Caller(1)
	if ok {
	 fmt.Printf("Called from %s, line #%d, func: %v\n",
        file, line, runtime.FuncForPC(pc).Name())
	}

	basename := "bot7"
	filepath := fmt.Sprintf("/root/golang/token/%s.txt", basename)

    b, err := ioutil.ReadFile(filepath) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }
    token := string(b)
    
	auth.LoginWithAuthToken(token)
	//auth.LoginWithQrCode(true)
	for {
		fetch, _ := talk.FetchOperations(service.Revision, 1)
		if len(fetch) > 0 {
			rev := fetch[0].Revision
			service.Revision = helper.MaxRevision(service.Revision, rev)
			bot(fetch[0])
		}
	}
	return
}